package connector

import (
	"Catalog/model/configuration"
	"Catalog/niceErrors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type SqlQuery interface {
	QueryString() string
	QueryParameters() []interface{}
	RowsProcessor(rows *sql.Rows) *niceErrors.NiceErrors
}

type SqlExec interface {
	ExecString() string
	ExecParameters() []interface{}
}

func SendQuery(query SqlQuery) *niceErrors.NiceErrors {
	conf, nErr := configuration.GetConfiguration()
	if nErr != nil {
		return niceErrors.FromErrorFull(nErr, "Cannot read configuration", "-", niceErrors.FATAL)
	}

	connString := conf.Sql.Username + ":" + conf.Sql.Password + "@tcp(" + conf.Sql.Address + ":" + conf.Sql.Port + ")/" + conf.Sql.DbName + "?parseTime=true"

	conn, err := sql.Open("mysql", connString)
	defer conn.Close()
	if err != nil {
		return niceErrors.FromErrorFull(err, "Cannot connect to mysql", "-", niceErrors.FATAL)
	}

	rows, err := conn.Query(query.QueryString(), query.QueryParameters()...)
	if err != nil {
		paramsStr := "{ "
		for i, v := range query.QueryParameters() {
			var conv string
			if convRaw, ok := v.(string); ok {
				conv = convRaw
			}
			paramsStr += strconv.Itoa(i) + ":" + conv + ", "
		}
		paramsStr += " }"
		return niceErrors.FromErrorFull(err, "Error from query execution, query: " + query.QueryString() + ", params: " + paramsStr, "-", niceErrors.ERROR)
	}

	nErr = query.RowsProcessor(rows)
	if nErr != nil {
		return nErr
	}

	rows.Close()

	return nil
}

func SendExec(exec SqlExec) *niceErrors.NiceErrors {
	conf, nErr := configuration.GetConfiguration()
	if nErr != nil {
		return niceErrors.FromErrorFull(nErr, "Cannot read configuration", "-", niceErrors.FATAL)
	}

	connString := conf.Sql.Username + ":" + conf.Sql.Password + "@tcp(" + conf.Sql.Address + ":" + conf.Sql.Port + ")/" + conf.Sql.DbName + "?parseTime=true"

	conn, err := sql.Open("mysql", connString)
	defer conn.Close()
	if err != nil {
		return niceErrors.FromErrorFull(err, "Cannot connect to mysql", "-", niceErrors.FATAL)
	}

	_, err = conn.Exec(exec.ExecString(), exec.ExecParameters()...)
	if err != nil {
		paramsStr := "{ "
		for i, v := range exec.ExecParameters() {
			var conv string
			if convRaw, ok := v.(string); ok {
				conv = convRaw
			}
			paramsStr += strconv.Itoa(i) + ":" + conv + ", "
		}
		paramsStr += " }"
		return niceErrors.FromErrorFull(err, "Error from exec execution, exec: " + exec.ExecString() + ", params: " + paramsStr, "-", niceErrors.ERROR)
	}

	return nil
}