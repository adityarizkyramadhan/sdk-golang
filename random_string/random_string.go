package randomstring

import (
	"errors"
	"math/rand"
	"time"
)

var (
	errNegativeLength = errors.New("length must be positive")
	seedRand          = rand.New(rand.NewSource(time.Now().Unix()))
)

type StringRandom struct {
	charset string
}

func Default() *StringRandom {
	return &StringRandom{
		charset: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
}

func New(charset string) *StringRandom {
	return &StringRandom{
		charset: charset,
	}
}

func (s *StringRandom) Generate(length int) (string, error) {
	if length < 0 {
		return "", errNegativeLength
	}
	random := make([]byte, length)
	for i := range random {
		random[i] = s.charset[seedRand.Intn(len(s.charset))]
	}
	return string(random), nil
}
