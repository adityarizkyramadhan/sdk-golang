package struct_field_test

import (
	"testing"

	sf "github.com/adityarizkyramadhan/sdk-golang/struct_field"
	"github.com/stretchr/testify/assert"
)

type Mahasiswa struct {
	Nama string `field_tag:"nama"`
	Umur int    `field_tag:"umur"`
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
