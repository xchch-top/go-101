package v0

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestDbUseReflectValuer(t *testing.T) {
	Open("sqlite3", "file:test.db?cache=shared&mode=memory", DbUseReflectValuer())
}
