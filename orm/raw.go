package orm

import (
	"database/sql"
	"goden/log"
	"strings"
)

// defer RollBack

type Session struct {
	db 			*sql.DB
	sql 		strings.Builder
	sqlParams 	[]any
	// sqlPaNums	int
	// sqlType		string
	// sqlAffected int
	// sqlRows		*sql.Rows
}
type Rows struct {
	values []any
}

// func NewSession(db *sql.DB) *Session {
// 	return &Session{db: db}
// }

func (s *Session) Reset() {
	s.sql.Reset()
	s.sqlParams = nil
	// s.sqlPaNums = 0
}

func (s *Session) RawSql(sql string, params ...any) *Session {
	s.sql.WriteString(sql)
	s.sqlParams = append(s.sqlParams, params...)
	// s.sqlPaNums = len(s.sqlParams)
	return s
}

// execute once order No execute result return
func (s *Session) Execute() (sql.Result) {
	defer s.Reset()
	log.Info(s.sql.String(), s.sqlParams)

	result,err := s.db.Exec(s.sql.String(), s.sqlParams...)
	if err!=nil {
		log.Error(err)
	}
	return result
}

func (s *Session) Query() (rows *sql.Rows, err error) {
	defer s.Reset()
	log.Info(s.sql.String(), s.sqlParams)

	rows, err = s.db.Query(s.sql.String(), s.sqlParams...)
	if err!=nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryOne() *sql.Row {
	defer s.Reset()
	log.Info(s.sql.String(), s.sqlParams)

	return s.db.QueryRow(s.sql.String(), s.sqlParams...)
}


// 	resultColumns,_ := rows.Columns()
// 	var results []any -> init
// 	for rows.Next() {
//	not work
// 		_ = rows.Scan(results...)
// }





