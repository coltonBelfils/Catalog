package controler

import (
	"Catalog/model/dataTypes"
	"Catalog/model/sql/calls/execs"
	"Catalog/model/sql/calls/queries"
	"Catalog/model/sql/connector"
	"Catalog/niceErrors"
)

func CreateUser(username string, email string) (dataTypes.User, *niceErrors.NiceErrors) {
	//TODO need to validate email by sending them an email that they respond to before they can use their account at all

	//TODO make sure that user doesn't already exist. SQL does that now but it's not super elegant

	CUExec, nErr := execs.CreateUser(username, email)
	if nErr != nil {
		return dataTypes.User{}, nErr
	}

	nErr = connector.SendExec(CUExec)
	if nErr != nil {
		return dataTypes.User{}, nErr
	}

	testQuery := queries.UserQueryByUsername(username)

	nErr = connector.SendQuery(testQuery)
	if nErr != nil {
		return dataTypes.User{}, niceErrors.FromErrorFull(nErr, "user: " + username + ", " + email + " was successfully created, but could not be queried", "Error creating user", niceErrors.ERROR)
	}

	return testQuery.Results[0], nil
}
