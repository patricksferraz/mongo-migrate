// +build integration

package migrate

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const globalTestCollection = "test-global"

func init() {
	Register(func(db *mgo.Database) error {
		return db.C(globalTestCollection).Insert(bson.M{"a": "b"})
	}, func(db *mgo.Database) error {
		return db.C(globalTestCollection).Remove(bson.M{"a": "b"})
	})
}
