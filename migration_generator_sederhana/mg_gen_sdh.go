package migrationgeneratorsederhana

import (
	"fmt"
	"reflect"
	"strings"

	sf "github.com/adityarizkyramadhan/sdk-golang/struct_field"
)

type MigratorSederhana struct {
	Data     map[string]interface{}
	DataSave interface{}
}

func NewMigratorSegerhana(data interface{}) *MigratorSederhana {
	convData := sf.NewGetTagByField(data)
	return &MigratorSederhana{
		Data:     convData.GetFieldNameAndType(),
		DataSave: data,
	}
}

func (m *MigratorSederhana) converterDataType() map[string]interface{} {
	for i := range m.Data {
		if m.Data[i] == "string" {
			m.Data[i] = "text"
		}
	}
	return m.Data
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (m *MigratorSederhana) GenerateSQL() string {
	tableName := getType(m.DataSave)
	tableName = strings.ToLower(tableName)
	tableName = strings.ReplaceAll(tableName, "*", "")
	data := m.converterDataType()
	sqlQuery := fmt.Sprintf(`CREATE TABLE %s (`, tableName)
	for i := range data {
		sqlQuery += i
		sqlQuery += " "
		sqlQuery += data[i].(string)
		sqlQuery += ", "
	}
	sqlQuery += ");"

	return sqlQuery
}
