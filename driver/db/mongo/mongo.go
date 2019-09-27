package mongo

import (
	"cortex3/conf"
	"fmt"
	"gopkg.in/mgo.v2"
	"sync"
	"time"
)

const (
	D_CORTEX_CONFIG = "cortex3"
	C_CORTEX_RECORD = "bauHistory"
	C_CORTEX_TEST = "cortex_request_log"
)


var (
	once       sync.Once
	MgoSession *mgo.Session
	mDatabase  *mgo.Database
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitMongo() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{conf.Yaml.Conf.Mongo.Host},
		Timeout:   3 * time.Second,
		Username:  conf.Yaml.Conf.Mongo.User,
		Password:  conf.Yaml.Conf.Mongo.Pwd,
		PoolLimit: 4096,
		Source:    conf.Yaml.Conf.Mongo.AuthDb,
	}
	session, err := mgo.DialWithInfo(dialInfo)
	CheckErr(err)
	session.SetMode(mgo.Monotonic, true)
	//mDatabase = session.DB(D_chat)
	MgoSession = session

	fmt.Println("Mongo 初始化成功")
}
func Connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	return connect(db, collection)
}
func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	if MgoSession == nil {
		InitMongo()
	}

	//GetSession()
	cSession := MgoSession.Copy()

	c := cSession.DB(db).C(collection)

	//cSession.SetMode(mgo.Monotonic, true)
	return cSession, c
}

/**具体操作*/
func Save(db, collection string, docs interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Insert(docs)
}

func SaveAll(db, collection string, docs []interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Insert(docs...)
}

func Pipeline(db, collection string, pipe, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Pipe(pipe).All(result)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Find(query).Select(selector).One(result)
}

func FindOneSortMost(db, collection, sort string, query, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	count, err := c.Find(query).Count()
	if err != nil {
		return err
	}
	if sort != "" {
		return c.Find(query).Sort(sort).Skip(count - 1).One(result)
	}
	return c.Find(query).Skip(count - 1).One(result)
}

func FindAll(db, collection, sort string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	if sort != "" {
		return c.Find(query).Sort(sort).Select(selector).All(result)
	}
	return c.Find(query).Select(selector).All(result)
}

func FindCount(db, collection, sort string, query, selector interface{}) (int, error) {
	session, c := connect(db, collection)
	defer session.Close()

	if sort != "" {
		return c.Find(query).Sort(sort).Select(selector).Count()
	}
	return c.Find(query).Select(selector).Count()
}

func FindAllWithLimit(db, collection, sort string, query, selector, result interface{}, limit int) error {
	session, c := connect(db, collection)
	defer session.Close()

	if sort != "" {
		return c.Find(query).Sort(sort).Select(selector).Limit(limit).All(result)
	}
	return c.Find(query).Sort("-update_time").Select(selector).Limit(limit).All(result)
}

func FindAllWithLimitAndSkip(db, collection, sort string, query, selector, result interface{}, limit int, skip int) error {
	session, c := connect(db, collection)
	defer session.Close()

	if sort != "" {
		return c.Find(query).Sort(sort).Select(selector).Skip(skip).Limit(limit).All(result)
	}
	return c.Find(query).Sort("-update_time").Select(selector).Skip(skip).Limit(limit).All(result)
}

func Update(db, collection string, selector, update interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Update(selector, update)

}

func UpdateAll(db, collection string, selector, update interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	_, err := c.RemoveAll(selector)
	return err
}

func MultiFind(db, collection, sorted string, query, result interface{}, page, limit int) error {
	session, c := connect(db, collection)
	defer session.Close()

	if sorted != "" {
		return c.Find(query).Sort(sorted).Skip(page * limit).Limit(limit).All(result)
	}
	return c.Find(query).Skip(page * limit).Limit(limit).All(result)
}
