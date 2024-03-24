package snippets

import "github.com/dgraph-io/badger/v4"

type Snippets struct {
	db *badger.DB
}
