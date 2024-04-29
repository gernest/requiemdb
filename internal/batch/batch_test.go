package batch

import (
	"os"
	"testing"
	"time"

	"github.com/gernest/requiemdb/internal/bitmaps"
)

func TestBatch(t *testing.T) {
	b := New()
	now, _ := time.Parse(time.RFC822, time.RFC822)
	now = now.UTC()

	t.Run("first column is timestamp", func(t *testing.T) {
		b.Reset()
		r := bitmaps.New()
		defer r.Release()
		r.Add(1)
		r.Add(2)
		b.Add(12, now, r)
		f, err := b.Build()
		if err != nil {
			t.Fatal(err)
		}
		got := f.String()

		os.WriteFile("testdata/batch.txt", []byte(got), 0600)
		w, err := os.ReadFile("testdata/batch.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := string(w)
		if want != got {
			t.Errorf("expected =>\n%s \ngot=> %v", want, got)
		}
	})
}
