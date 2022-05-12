package orm

import (
	"database/sql"
	"goden/log"
)

type Butler struct {
	db *sql.DB
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
	log.Info("Success: Database connected")
	return
}

func (b *Butler) Close() {
	if err := b.db.Close(); err!=nil {
		log.Error("Failed: Database close")
	}
	log.Info("Success: Database closed")
}

func (b *Butler) NewSession() *Session {
	return &Session{db: b.db}
}