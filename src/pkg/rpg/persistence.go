package rpg

import (
  "launchpad.net/mgo"
  "os"
)

type M map[string]interface{}

type Database interface {
  StoreEntity(Entity) os.Error
  Entities(...M) ([]Entity, os.Error)
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
  return m.db.C("entities").Insert(entity.Serialize())
}

func (m *MongoConn) Entities(args ...M) ([]Entity, os.Error) {
  var err os.Error
  var count int
  var params, result M

  if len(args) == 0 {
    params = M{}
  } else {
    params = args[0]
  }

  qry := m.db.C("entities").Find(params)
  count, err = qry.Count()   // FIXME: count is probably too expensive
  if err != nil {
    return nil, err
  }

  entities := make([]Entity, count)
  i := 0
  err = qry.For(&result, func() os.Error {
    entities[i] = Unserialize(result)
    return nil
  })
  return entities, err
}

func (m *MongoConn) Close() {
  m.session.Close()
}
