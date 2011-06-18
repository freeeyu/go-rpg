package rpg

import (
  "launchpad.net/mgo"
  "os"
)

type M map[string]interface{}

type Database interface {
  StorePlayer(*Player) os.Error
  Players(...M) ([]*Player, os.Error)
  Close()
}

type MongoConn struct {
  session *mgo.Session
  db mgo.Database
}

func NewMongoConn(host string, databaseName string) (*MongoConn, os.Error) {
  session, err := mgo.Mongo(host)
  if err != nil {
    return nil, err
  }
  return &MongoConn{session, session.DB(databaseName)}, nil
}

func (m *MongoConn) StorePlayer(player *Player) os.Error {
  data := M{"name": player.Name(), "xp": player.XP(), "hp": player.HP()}
  return m.db.C("players").Insert(data)
}

func (m *MongoConn) Players(args ...M) ([]*Player, os.Error) {
  var result *Player
  var err os.Error
  var count int

  qry := m.db.C("players")
  count, err = qry.Count()   // FIXME: count is probably too expensive
  if err != nil {
    return nil, err
  }

  players := make([]*Player, count)
  i := 0
  err = qry.Find(M{}).For(&result, func() os.Error {
    players[i] = result
    return nil
  })
  return players, err
}

func (m *MongoConn) Close() {
  m.session.Close()
}
