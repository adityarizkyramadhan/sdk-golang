package struct_field_test

import (
	"fmt"
	"testing"

	sf "github.com/adityarizkyramadhan/sdk-golang/struct_field"
	"github.com/stretchr/testify/assert"
)

type Mahasiswa struct {
	Nama string `field_tag:"nama"`
	Umur int    `field_tag:"umur"`
	Nim  string `field_tag:"nim"`
}

func Test_Method_Satu(t *testing.T) {

	mhs := new(Mahasiswa)

	structField := sf.NewGetTagByField(mhs)

	testMethodSatu := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range testMethodSatu {
		t.Run(tt.name, func(t *testing.T) {
			panjangMahasiswa := structField.SizeField()
			assert.Equal(t, 2, panjangMahasiswa, panjangMahasiswa)
		})
	}
}

func Test_Method_Dua(t *testing.T) {
	mhs := new(Mahasiswa)

	structField := sf.NewGetTagByField(mhs)
	testMethodDua := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range testMethodDua {
		t.Run(tt.name, func(t *testing.T) {
			field, ok := structField.GetByFieldName("Nama")
			assert.True(t, ok)
			tagName := structField.GetStructTag(field, "field_tag")
			assert.Equal(t, "nama", tagName, tagName)
		})
	}
}

func Test_Array(t *testing.T) {
	mhs := new(Mahasiswa)

	structField := sf.NewGetTagByField(mhs)
	arr := structField.ToArray()
	fmt.Println(arr)
}

func Test_To_String(t *testing.T) {
	mhs := new(Mahasiswa)

	structField := sf.NewGetTagByField(mhs)
	str := structField.ToString()
	fmt.Println(str)
}

func Test_Type_Data(t *testing.T) {
	mhs := new(Mahasiswa)

	structField := sf.NewGetTagByField(mhs)

	data := structField.GetFieldNameAndType()

	for _, v := range structField.ToArray() {
		fmt.Println(data[v])
	}
}
