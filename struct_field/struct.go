package struct_field

import "reflect"

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
	return reflect.TypeOf(f.Data).Elem().Field(0)
}

func (f *FieldTag) GetByFieldName(name string) (reflect.StructField, bool) {
	return reflect.TypeOf(f.Data).Elem().FieldByName(name)
}

func (f *FieldTag) GetStructTag(rsf reflect.StructField, tagName string) string {
	return string(rsf.Tag.Get(tagName))
}
