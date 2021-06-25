package controler

import (
	"Catalog/model/dataTypes"
	"Catalog/model/sqlDatabase/calls/queries"
	"Catalog/model/sqlDatabase/connector"
	"Catalog/niceErrors"
	"github.com/google/uuid"
)

func GetUserByUserId(userId uuid.UUID) (dataTypes.User, *niceErrors.NiceErrors) {
	query := queries.UserQueryById(userId)
	connErr := connector.SendQuery(query)
	if connErr != nil {
		nConnErr := niceErrors.FromError(connErr)
		return dataTypes.User{}, nConnErr
	}

	if len(query.Results) == 0 {
		return dataTypes.User{}, niceErrors.New("userError", "Unable to find user with id: "+userId.String(), niceErrors.UnexpectedResultError, niceErrors.INFO)
	} else if len(query.Results) > 1 {
		return dataTypes.User{}, niceErrors.New("Multiple users with id: "+userId.String(), "Unable to find user with id: "+userId.String(), niceErrors.UnexpectedResultError, niceErrors.ERROR)
	}
	return query.Results[0], nil
}

func GetUserByUsername(username string) (dataTypes.User, *niceErrors.NiceErrors) {
	query := queries.UserQueryByUsername(username)
	connErr := connector.SendQuery(query)
	if connErr != nil {
		nConnErr := niceErrors.FromError(connErr)
		return dataTypes.User{}, nConnErr
	}

	if len(query.Results) == 0 {
		return dataTypes.User{}, niceErrors.New("userError", "Unable to find user with username: "+username, niceErrors.UnexpectedResultError, niceErrors.INFO)
	} else if len(query.Results) > 1 {
		return dataTypes.User{}, niceErrors.New("Multiple users with username: "+username, "Unable to find user with username: "+username, niceErrors.UnexpectedResultError, niceErrors.ERROR)
	}
	return query.Results[0], nil
}
