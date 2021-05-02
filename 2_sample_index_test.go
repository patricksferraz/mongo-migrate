// +build integration

package migrate

import (
	"gopkg.in/mgo.v2"
)

const globalTestIndexName = "test_idx_2"

func init() {
	Register(func(db *mgo.Database) error {
		return db.C(globalTestCollection).EnsureIndex(mgo.Index{Name: globalTestIndexName, Key: []string{"a"}})
	}, func(db *mgo.Database) error {
		return db.C(globalTestCollection).DropIndexName(globalTestIndexName)
	})
}
