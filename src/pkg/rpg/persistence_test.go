package rpg

import (
  "testing"
  "launchpad.net/gobson/bson"
  "launchpad.net/mgo"
  "os"
)

var msession *mgo.Session
var mdb mgo.Database
var err os.Error

func setup(t *testing.T) {
  msession, err = mgo.Mongo("localhost")
  if err != nil {
    t.Fatal("Can't connect to mongodb")
  }
  mdb = msession.DB("go-rpg-test")
}

func teardown(t *testing.T) {
  err = mdb.C("entities").RemoveAll(bson.M{})
  msession.Close()
}

func die(t *testing.T, args ...interface{}) {
  teardown(t)
  t.Fatal(args)
}

func TestStoreEntity(t *testing.T) {
  var db *MongoConn
  setup(t)

  db, err = NewMongoConn("localhost", "go-rpg-test")
  if err != nil {
    die(t, err.String())
  }
  defer db.Close()

  player := NewPlayer("foo", 0)
  err = db.StoreEntity(player)
  if err != nil {
    t.Error(err)
  }

  var result bson.M
  err = mdb.C("entities").Find(bson.M{"name": "foo"}).One(&result)
  if err != nil {
    t.Error(err)
  }
  if result["name"] != "foo" || result["xp"] != 0 || result["hp"] != 10 {
    t.Error("player values were not correct")
  }

  teardown(t)
}
