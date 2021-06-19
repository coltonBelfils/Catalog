package middleware

import (
	"Catalog/niceErrors"
	re "Catalog/view/api/responder"
	"context"
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"net/http"
	"strings"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			nErr := niceErrors.New("incorrect jwt format, check frontend", "you are not authorized to use this endpoint", niceErrors.WARN)
			re.JsonRequestErrorResponder(w, nErr, 401)
			return
		}

		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, ValidationKeyGetter)
		nErr := niceErrors.FromError(err)
		if nErr != nil {
			re.JsonRequestErrorResponder(w, nErr, 401)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			// Access context values in handlers like this
			// props, _ := r.Context().Value("props").(jwt.MapClaims)
			fmt.Println(ctx.Value("props"))
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			nErr := niceErrors.New("incorrect jwt format, check frontend", "you are not authorized to use this endpoint", niceErrors.WARN)
			re.JsonRequestErrorResponder(w, nErr, 401)
			return
		}
	})

}

func ValidationKeyGetter(token *jwt.Token) (interface{}, error) {
	// Verify 'aud' claim
	aud := "https://rankingcatalog.com/api"
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		return token, niceErrors.New("invalid audience", "you are not authorized to use this endpoint", niceErrors.INFO)
	}
	// Verify 'iss' claim
	iss := "https://rankingcatalog.us.auth0.com/"
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
	if !checkIss {
		return token, niceErrors.New("invalid issuer", "you are not authorized to use this endpoint", niceErrors.INFO)
	}

	cert, nErr := getPemCert(token)
	if nErr != nil {
		return nil, nErr
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func getPemCert(token *jwt.Token) (string, *niceErrors.NiceErrors) {
	cert := ""
	resp, err := http.Get("https://rankingcatalog.us.auth0.com/.well-known/jwks.json")

	if err != nil {
		nErr := niceErrors.FromErrorFull(err, "could not reach: https://rankingcatalog.us.auth0.com/.well-known/jwks.json", "something went wrong", niceErrors.FATAL)
		return cert, nErr
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		nErr := niceErrors.FromError(err)
		return cert, nErr
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		nErr := niceErrors.New("unable to find appropriate key", "something went wrong", niceErrors.INFO)
		return cert, nErr
	}

	return cert, nil
}
