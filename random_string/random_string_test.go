package randomstring_test

import (
	"fmt"
	"testing"

	randomstring "github.com/adityarizkyramadhan/sdk-golang/random_string"
)

func TestGenerate(t *testing.T) {
	rs := randomstring.Default()

	fmt.Println(rs.Generate(30))
}
