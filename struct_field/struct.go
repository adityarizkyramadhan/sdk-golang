package struct_field

import (
	"reflect"
)

type FieldTag struct {
	Data interface{}
}

func NewGetTagByField(data interface{}) *FieldTag {
	return &FieldTag{
		Data: data,
	}
}

// Banyak jumlah field di struct
func (f *FieldTag) SizeField() int {
	return reflect.TypeOf(f.Data).Elem().NumField()
}

func (f *FieldTag) GetByIndex(index int) reflect.StructField {
	return reflect.TypeOf(f.Data).Elem().Field(index)
}

func (f *FieldTag) GetByFieldName(name string) (reflect.StructField, bool) {
	return reflect.TypeOf(f.Data).Elem().FieldByName(name)
}

func (f *FieldTag) GetStructTag(rsf reflect.StructField, tagName string) string {
	return string(rsf.Tag.Get(tagName))
}

func (f *FieldTag) ToArray() []string {
	var data []string
	for i := 0; i < f.SizeField(); i++ {
		tag := f.GetByIndex(i)
		tagNama := f.GetStructTag(tag, "field_tag")
		data = append(data, tagNama)
	}
	return data
}
