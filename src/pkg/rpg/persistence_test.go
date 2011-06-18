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
  err = mdb.C("players").RemoveAll(M{})
  msession.Close()
}

func die(t *testing.T, args ...interface{}) {
  teardown(t)
  t.Fatal(args)
}

func TestStorePlayer(t *testing.T) {
  var db *MongoConn
  setup(t)

  db, err = NewMongoConn("localhost", "go-rpg-test")
  if err != nil {
    die(t, err.String())
  }
  defer db.Close()

  player := NewPlayer("foo", 0)
  err = db.StorePlayer(player)
  if err != nil {
    t.Error(err)
  }

  var result M
  err = mdb.C("players").Find(M{"name": "foo"}).One(&result)
  if err != nil {
    t.Error(err)
  }
  t.Log(result)
  if result["name"] != "foo" || result["xp"] != 0 || result["hp"] != 10 {
    t.Error("player values were not correct")
  }

  teardown(t)
}

func TestPlayers(t *testing.T) {
  var db *MongoConn
  var players []*Player
  setup(t)

  db, err = NewMongoConn("localhost", "go-rpg-test")
  if err != nil {
    die(t, err.String())
  }
  defer db.Close()

  player := NewPlayer("foo", 0)
  err = db.StorePlayer(player)
  if err != nil {
    t.Error(err)
  }

  players, err = db.Players()
  if err != nil {
    t.Error(err)
  } else if len(players) != 1 {
    t.Error("expected 1 player, found", len(players))
  }

  teardown(t)
}
