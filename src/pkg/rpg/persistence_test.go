package rpg

import (
  "testing"
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
  err = mdb.C("entities").RemoveAll(M{})
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

  player := NewPlayer("foo")
  err = db.StoreEntity(player)
  if err != nil {
    t.Error(err)
  }

  var result M
  err = mdb.C("entities").Find(M{"name": "foo"}).One(&result)
  if err != nil {
    t.Error(err)
  } else if result["name"] != "foo" || result["xp"] != 0 || result["hp"] != 10 {
    t.Error("player values were not correct")
  }

  teardown(t)
}

func TestEntities(t *testing.T) {
  var db *MongoConn
  var entities []Entity
  setup(t)

  db, err = NewMongoConn("localhost", "go-rpg-test")
  if err != nil {
    die(t, err.String())
  }
  defer db.Close()

  player := NewPlayer("foo")
  err = db.StoreEntity(player)
  if err != nil {
    t.Error(err)
  }

  entities, err = db.Entities()
  if err != nil {
    t.Error(err)
  } else if len(entities) != 1 {
    t.Error("expected 1 entities, found", len(entities))
  }

  teardown(t)
}

func TestEntitiesWithFilter(t *testing.T) {
  var db *MongoConn
  var entities []Entity
  setup(t)

  db, err = NewMongoConn("localhost", "go-rpg-test")
  if err != nil {
    die(t, err.String())
  }
  defer db.Close()

  player_1 := NewPlayer("foo")
  err = db.StoreEntity(player_1)
  if err != nil {
    t.Error(err)
  }

  player_2 := NewPlayer("bar")
  err = db.StoreEntity(player_2)
  if err != nil {
    t.Error(err)
  }

  entities, err = db.Entities(M{"name": "bar"})
  if err != nil {
    t.Error(err)
  } else if len(entities) != 1 {
    t.Error("expected 1 entity, found", len(entities))
  } else if entities[0].Name() != "bar" {
    t.Error("found wrong player:", entities[0])
  }

  teardown(t)
}
