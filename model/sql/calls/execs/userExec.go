package execs

import (
	"Catalog/niceErrors"
	"regexp"
)

type UserExec struct {
	Results        interface{} //TODO change to accurate type
	execString     string
	execParameters []interface{}
}

func (u *UserExec) ExecString() string {
	return u.execString
}

func (u *UserExec) ExecParameters() []interface{} {
	return u.execParameters
}

func CreateUser(username string, email string) (*UserExec, *niceErrors.NiceErrors) {
	if len(username) > 20 && len(username) < 0 {
		return &UserExec{}, niceErrors.New("invalid username: "+username, "invalid username: "+username, niceErrors.INFO)
	}

	valid, err := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if err != nil {
		return &UserExec{}, niceErrors.FromErrorFull(err, "invalid email: "+email, "Invalid email", niceErrors.INFO)
	}

	if !valid {
		return &UserExec{}, niceErrors.New("invalid email: "+email, "invalid email: "+email, niceErrors.INFO)
	}

	return &UserExec{
		execString: `
INSERT INTO user(user_id, email, username)
values(null, ?, ?);
`,
		execParameters: []interface{}{email, username},
	}, nil
}
