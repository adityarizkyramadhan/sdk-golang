package migrationgeneratorsederhana_test

import (
	"fmt"
	"testing"

	mig "github.com/adityarizkyramadhan/sdk-golang/migration_generator_sederhana"
)

type Mahasiswa struct {
	Nama string `field_tag:"nama"`
	Umur int    `field_tag:"umur"`
	Nim  string `field_tag:"nim"`
}

func TestGenerateQuery(t *testing.T) {
	mhs := new(Mahasiswa)
	sqlMigrator := mig.NewMigratorSegerhana(mhs)
	fmt.Println(sqlMigrator.GenerateSQL())
}
