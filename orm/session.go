package orm

import (
	"database/sql"
	"goden/log"
)

// defer RollBack

// SqlType
const (
	None = iota
	Insert
	Delete
	Update
	Select
)

type Session struct {
	db 			*sql.DB
	sql 		string
	sqlArgs 	[]any
	dialect		Dialect
	refTable	*Schema
}

type SessionResult struct {
	session *Session
	rows 	*sql.Rows
	result  *sql.Result
	// sqlAffected int
	// sqlType		string
}

func NewSession(db *sql.DB, dialect Dialect) *Session {
	return &Session{
		db:	db,
		dialect: dialect,
	}
}

func (s *Session) Reset() {
	s.sql = ""
	s.sqlArgs = nil
}

func (s *Session) RawSql(sql string, args ...any) *Session{
	s.Reset()
	s.sql = sql
	// s.sqlArgs = append(s.sqlArgs, args...)
	s.sqlArgs = args
	return s
}

func (s *Session) Execute() (sql.Result,error) {
	log.Info(s.sql, s.sqlArgs)

	result,err := s.db.Exec(s.sql, s.sqlArgs...)
	if err!=nil {
		log.Error(err)
	}
	return result,err
}

func (s *Session) Query() (rows *sql.Rows,err error) {
	log.Info(s.sql, s.sqlArgs)

	rows, err = s.db.Query(s.sql, s.sqlArgs...)
	// defer rows.Close()
	if err!=nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryOne() *sql.Row {
	log.Info(s.sql, s.sqlArgs)

	return s.db.QueryRow(s.sql, s.sqlArgs...)
}

// not work
func (s *Session) Test() {
	log.Info(s.sql, s.sqlArgs)

	rows, err := s.db.Query(s.sql, s.sqlArgs...)
	defer rows.Close()
	resultColumns,_ := rows.Columns()
	var results []any
	for i:=0;i<len(resultColumns);i++ {
		results = append(results, (*string)(nil))
	}
	for rows.Next() {
		_ = rows.Scan(results...)
	}
	for _,r := range results {
		log.Warn(r)
	}
	if err!=nil {
		log.Error(err)
	}
	return
}


// 	resultColumns,_ := rows.Columns()
// 	var results []any -> init
// 	for rows.Next() {
//	not work
// 		_ = rows.Scan(results...)
// }





