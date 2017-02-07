package dailynumber

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewMongo(col *mgo.Collection) *Mongo {
	return &Mongo{col}
}

type Mongo struct {
	col *mgo.Collection
}

type Document struct {
	Id     string `bson:"_id"`
	Number int    `bson:"number"`
}

func (m *Mongo) Generate() string {
	var doc Document

	bucket := fmt.Sprintf("%c", 65+rand.Intn(8))
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"number": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	now := time.Now()
	if _, err := m.col.Find(bson.M{"_id": now.Format("2006-01-02-") + bucket}).Apply(change, &doc); err != nil {
		// try to recover with something good enough, collisions would happen per second frame
		// this is on a "best-effort" basis but shouldn't usually happen
		year, month, day := now.In(time.UTC).Date()
		midnight := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		number := int(time.Since(midnight).Seconds())

		return fmt.Sprintf("%c%d", 73+number%18, number/18)
	}

	return fmt.Sprintf("%s%d", bucket, doc.Number)
}
