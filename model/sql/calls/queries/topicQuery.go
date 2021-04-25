package queries

import (
	"Catalog/model/dataTypes"
	"Catalog/niceErrors"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

//Query structure

type TopicQuery struct {
	Results         []dataTypes.Topic
	queryString     string
	queryParameters []interface{}
}

func (q *TopicQuery) QueryString() string {
	return q.queryString
}

func (q *TopicQuery) QueryParameters() []interface{} {
	return q.queryParameters
}

func (q *TopicQuery) RowsProcessor(rows *sql.Rows) *niceErrors.NiceErrors {
	q.Results = []dataTypes.Topic{}

	for rows.Next() {
		var (
			rawTopicId     string
			title          string
			description    string
			imageUrl       string
			rawDateCreated sql.NullTime
		)

		err := rows.Scan(&rawTopicId, &title, &description, &imageUrl, &rawDateCreated)
		if err != nil {
			return niceErrors.FromErrorFull(err, "Error scanning sql row", "-", niceErrors.FATAL)
		}

		var topicId uuid.UUID = uuid.UUID{}
		if topicId, err = uuid.Parse(rawTopicId); err != nil {
			return niceErrors.FromErrorFull(err, "Error parsing uuid", "-", niceErrors.FATAL)
		}

		var dateCreated time.Time
		if rawDateCreated.Valid {
			dateCreated = rawDateCreated.Time
		}

		q.Results = append(q.Results, dataTypes.Topic{
			TopicId:     topicId,
			Title:       title,
			Description: description,
			ImageUrl:    imageUrl,
			DateCreated: dateCreated,
		})
	}
	return nil
}

//Queries

func TopicQueryByAdminUserId(userId uuid.UUID) *TopicQuery {
	return &TopicQuery{
		queryString: `
SELECT t.topic_id, t.title, t.description, t.image_url,  t.date_created
FROM topic_admin as ta join topic as t on ta.topic_id = t.topic_id
where ta.user_id = ?;
`,
		queryParameters: []interface{}{userId},
	}
}

func TopicQueryByAdminUserUsername(username string) *TopicQuery {
	return &TopicQuery{
		queryString: `
SELECT t.topic_id, t.title, t.description, t.image_url,  t.date_created
FROM user as u join topic_admin as ta on u.user_id = ta.user_id
join topic as t on ta.topic_id = t.topic_id
where u.username = ?;
`,
		queryParameters: []interface{}{username},
	}
}
