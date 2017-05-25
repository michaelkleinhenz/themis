package workitems;

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type WorkItem struct {
    Id       string    `bson:"id"`
		Type     string    `bson:"type"`
}

func NewWorkItem(id, typeStr string) WorkItem {
	return WorkItem {
		Id:   id,
		Type: typeStr,
	}
}

func WriteToDatabase(database *mgo.Database) {
	workItem := NewWorkItem("newId0", "newType0");
	coll := database.C("workitems")
	if err := coll.Insert(workItem); err != nil {
			panic(err)
	}
	fmt.Println("WorkItem inserted successfully!")
}

func ReadFromDatabase(database *mgo.Database, id string) WorkItem {
	var workItem WorkItem
	coll := database.C("workitems")
	err := coll.Find(bson.M{"id": id}).One(&workItem)
	if err != nil {
			panic(err)
	}
	return workItem
}

func UpdateDatabase(database *mgo.Database) {
	coll := database.C("workitems")
	workItemId := bson.ObjectIdHex("55da804ea5b2a779329ceb8e") // FIXME
	newValue := "Some Value"
	update := bson.M{"$set": bson.M{"type": newValue}}
	if err := coll.UpdateId(workItemId, update); err != nil {
			panic(err)
	}
}

func DeleteDatabase(database *mgo.Database) {
	coll := database.C("workitems")
	info, err := coll.RemoveAll(bson.M{"type": "some Type"})
	if err != nil {
			panic(err)
	}
	fmt.Printf("%d workItem(s) removed!\n", info.Removed)
}

/*
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
*/