package cache_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/adityarizkyramadhan/sdk-golang/cache"
	"github.com/stretchr/testify/assert"
)

func Test_Name(t *testing.T) {
	type Mahasiswa struct {
		NIM       string
		Mahasiswa string
	}
	tests := []struct {
		name      string
		Mahasiswa *Mahasiswa
		Key       string
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			Mahasiswa: &Mahasiswa{
				NIM:       "215150201111007",
				Mahasiswa: "Aditya Rizky R",
			},
			Key: "data-1",
		},
		{
			name: "Test 2",
			Mahasiswa: &Mahasiswa{
				NIM:       "215150201111008",
				Mahasiswa: "Mahasiswa 2",
			},
			Key: "data-2",
		},
		{
			name: "Test 3",
			Mahasiswa: &Mahasiswa{
				NIM:       "215150201111009",
				Mahasiswa: "Mahasiswa 3",
			},
			Key: "data-3",
		},
	}
	cs := cache.NewCacheService("localhost:6379", 100000, time.Minute)

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cs.AddData(ctx, tt.Key, tt.Mahasiswa, time.Minute)
			assert.NoError(t, err, "Failed get data")
			temp := new(Mahasiswa)
			err = cs.GetData(ctx, tt.Key, temp)
			assert.NoError(t, err, "Failed get data")
		})
	}

	err := cs.Close()
	if err != nil {
		log.Println(err.Error())
	}

	/*
		=== RUN   Test_Name
		=== RUN   Test_Name/Test_1
		=== RUN   Test_Name/Test_2
		=== RUN   Test_Name/Test_3
		--- PASS: Test_Name (0.01s)
		    --- PASS: Test_Name/Test_1 (0.01s)
		    --- PASS: Test_Name/Test_2 (0.00s)
		    --- PASS: Test_Name/Test_3 (0.00s)
		PASS
		ok      github.com/adityarizkyramadhan/sdk-golang/cache 0.361s

	*/
}
