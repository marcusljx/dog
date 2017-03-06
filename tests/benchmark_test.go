package tests

import (
	"testing"

	"fmt"

	"os"

	"github.com/marcusljx/dog"
)

func BenchmarkStore(b *testing.B) {
	db, err := dog.Open(filepath)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		db.Set(fmt.Sprintf("%d", i), s1A)
	}

	if err := db.Close(); err != nil {
		b.FailNow()
	}
	if err := os.Remove(filepath); err != nil {
		b.FailNow()
	}
}
