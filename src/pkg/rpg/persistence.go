package rpg

import (
  "launchpad.net/gobson/bson"
  "launchpad.net/mgo"
  "os"
)

type Database interface {
  StoreEntity(Entity) os.Error
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

func (m *MongoConn) StoreEntity(entity Entity) os.Error {
  data := bson.M{"name": entity.Name(), "xp": entity.XP(), "hp": entity.HP()}
  return m.db.C("entities").Insert(data)
}

func (m *MongoConn) Close() {
  m.session.Close()
}
