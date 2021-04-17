package api

import "net/http"

func GetDocumentation(w http.ResponseWriter, req *http.Request) {
	HtmlRequestResponder(w, `
<html>
<head>
<title>Placeholder</title>
</head>
<body>
<h2>Placeholder</h2>
</body>
</html>
`, 200)
}
