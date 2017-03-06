package tests

import (
	"testing"

	"os"

	"github.com/marcusljx/dog"
	"github.com/stretchr/testify/assert"
)

const (
	filepath = "sample.dog"
)

type Sample1 struct {
	ABC   string
	XYZ   int
	Array []string
}

type Sample2 struct {
	A []string
	B []int
}

var (
	s1A = Sample1{
		ABC:   "HELLO WORLD",
		XYZ:   12345,
		Array: []string{"Hello", "World", "Marcus"},
	}

	s1B = Sample1{
		ABC:   "Go Away",
		XYZ:   54321,
		Array: []string{"No", "No", ""},
	}

	s2A = Sample2{
		A: []string{"No", "No", ""},
		B: []int{1, 2, 3, 4},
	}

	s2B = Sample2{
		A: []string{"", "", "Yes"},
		B: []int{1, 2, 3, 4},
	}

	db *dog.DB
)

func TestStoreStruct(t *testing.T) {
	db, err := dog.Open(filepath)
	assert.NoError(t, err)

	db.Set("s1A", s1A)
	db.Set("s1B", s1B)
	db.Set("s2A", s2A)
	db.Set("s2B", s2B)

	assert.NoError(t, db.Close())
}

func TestGetStruct(t *testing.T) {
	db, err := dog.Open(filepath)
	assert.NoError(t, err)

	db.Set("s1A", s1A)
	db.Set("s1B", s1B)
	db.Set("s2A", s2A)
	db.Set("s2B", s2B)

	assert.NoError(t, db.Close())
	assert.NoError(t, os.Remove(filepath))
}
