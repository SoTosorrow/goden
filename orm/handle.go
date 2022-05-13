package orm

import (
	"database/sql"
	"goden/log"
)

type Butler struct {
	db *sql.DB
	dialect Dialect
	// dbType string
}

func (b *Butler) GetDB() *sql.DB {
	return b.db
}

func NewButler(driverName string, dataSourceName string) (b *Butler, err error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err!= nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err!=nil {
		log.Error(err)
		return
	}
	b = &Butler{db : db}
	log.Success("Success: Database connected")
	return
}

// example : b.RawSql("xxxx ?",val) or b.RawSql("xxxx ?",&val)
// []any 和 []*any
// sqlQuoteArgs []&any
// 对其他变量的引用会不会导致其他变量声明周期延长？当sqlQuoteArgs=nil时，其他变量会不会被释放？GC问题
// 考虑gc压力
func (b *Butler) RawSql(sql string, args ...any) *Session {
	return &Session{
		db : b.db,
		sql: sql,
		sqlArgs: args,
		dialect: b.dialect,
	}
}

func (b *Butler) Close() {
	if err := b.db.Close(); err!=nil {
		log.Error("Failed: Database close")
	}
	log.Success("Success: Database closed")
}

func (b *Butler) NewSession() *Session {
	return &Session{db: b.db}
}