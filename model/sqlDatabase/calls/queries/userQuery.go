package queries

import (
	"Catalog/model/dataTypes"
	"Catalog/niceErrors"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

//Query structure

type UserQuery struct {
	Results         []dataTypes.User
	queryString     string
	queryParameters []interface{}
}

func (q *UserQuery) QueryString() string {
	return q.queryString
}

func (q *UserQuery) QueryParameters() []interface{} {
	return q.queryParameters
}

func (q *UserQuery) RowsProcessor(rows *sql.Rows) *niceErrors.NiceErrors {
	q.Results = []dataTypes.User{}

	for rows.Next() {
		var (
			rawUserId      string
			email          string
			emailVerified  bool
			username       string
			rawDateCreated sql.NullTime
		)

		err := rows.Scan(&rawUserId, &email, &emailVerified, &username, &rawDateCreated)
		if err != nil {
			return niceErrors.FromErrorFull(err, "Error scanning sqlDatabase row", "-", niceErrors.SqlError, niceErrors.FATAL)
		}

		var userId uuid.UUID = uuid.UUID{}
		if userId, err = uuid.Parse(rawUserId); err != nil {
			return niceErrors.FromErrorFull(err, "Error parsing uuid", "-", niceErrors.UnexpectedResultError, niceErrors.FATAL)
		}

		var dateCreated time.Time
		if rawDateCreated.Valid {
			dateCreated = rawDateCreated.Time
		}

		q.Results = append(q.Results, dataTypes.User{
			UserId:        userId,
			Email:         email,
			EmailVerified: emailVerified,
			Username:      username,
			DateCreated:   dateCreated,
		})
	}
	return nil
}

//Queries

func UserQueryById(userId uuid.UUID) *UserQuery {
	return &UserQuery{
		queryString: `
SELECT u.user_id, u.email, u.email_verified, u.username, u.date_created
FROM user as u
WHERE u.user_id = ?;
`,
		queryParameters: []interface{}{userId},
	}
}

func UserQueryByUsername(username string) *UserQuery {
	return &UserQuery{
		queryString: `
SELECT u.user_id, u.email, u.email_verified, u.username, u.date_created
FROM user as u
WHERE u.username = ?;
`,
		queryParameters: []interface{}{username},
	}
}
