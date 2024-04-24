package translate

import (
	"testing"

	"github.com/dgraph-io/badger/v4"
)

func TestTranslate_TranslateKey(t *testing.T) {
	type T struct {
		k string
		w uint64
	}

	s := []T{
		{
			k: "0",
			w: 0,
		},
		{
			k: "1",
			w: 1,
		},
		{
			k: "2",
			w: 2,
		},
		{
			k: "1",
			w: 1,
		},
	}
	db, err := badger.Open(badger.DefaultOptions("").
		WithInMemory(true).WithLogger(nil))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	tx, err := New(db, []byte("2"))
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Close()

	for _, v := range s {
		id, err := tx.TranslateKey([]byte(v.k))
		if err != nil {
			t.Fatal(err)
		}
		key, err := tx.TranslateID(id)
		if err != nil {
			t.Fatal(err)
		}
		if id != v.w {
			t.Errorf("expected id  %d got %d", v.w, id)
		}
		if key != v.k {
			t.Errorf("expected key %q got %q", v.k, key)
		}
	}
}
