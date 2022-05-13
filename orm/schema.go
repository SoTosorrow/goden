package orm

import (
	"go/ast"
	"reflect"
)

type Field struct {  		// table element
	Name 	string
	Type 	string
	Tag  	string
}

type Schema struct {		// table
	RefModel		any			
	Name		 	string		// table name
	Fields 			[]*Field	// table elements
	FieldNames 		[]string
	fieldMap 		map[string]*Field
}

func (s *Schema) GetField(name string) *Field{
	return s.fieldMap[name]
}

// take a Struct-Object Parse To A Relational Mapping Table 
func Parse(modelPtr any, d Dialect) *Schema {
	// get (value of ptr) by reflect.Indirect
	modelType := reflect.Indirect(reflect.ValueOf(modelPtr)).Type()

	schema := &Schema{
		RefModel:		modelPtr,
		Name:			modelType.Name(),
		fieldMap: 		make(map[string]*Field),
	}

	for i:=0; i<modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.ConvertType(reflect.Indirect(reflect.New(p.Type))),
			}
			if v,ok := p.Tag.Lookup("orm");ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields,field)
			schema.FieldNames = append(schema.FieldNames,p.Name)
			schema.fieldMap[p.Name] =field
		}
	}
	return schema
}