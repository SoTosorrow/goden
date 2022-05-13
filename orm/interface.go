package orm

import (
	"reflect"
)

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	// turn go type -> database field type
	ConvertType(typ reflect.Value) string
	TableCheckSql(tableName string) (string, []any)
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}