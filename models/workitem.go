package models

import (
	"gopkg.in/mgo.v2/bson"
)

type WorkItem struct {
  Id          bson.ObjectId       `bson:"_id,omitempty" json:"id"`
  Type        string              `bson:"type"`
  Attributes  map[string]string   `bson:"attributes"`
}

func (workItem *WorkItem) getCollectionName() string {
  return "workitems"
}

func (workItem *WorkItem) Create() *WorkItem {
    m := map[string]string{
        "key1": "value1",
        "key2": "value2",
    }
    w := new(WorkItem)
    w.Type = "workitems"
    w.Attributes = m
    return w
}

/*
WorkItem
  id
  type
  hasChildren
  attributes (map)
  relationships
  relationalData
  links
    self
    sourceLinkTypes
    targetLinkTypes

WorkItemRelations
  area
    data (AreaModel)
  assignees
    data (User[])
  baseType
    data (WorkItemType)
  children
    links
      related
    meta
      hasChildren (boolean)
  comments
    data (Comment[])
    links
      self
      related
    meta
      totalCount
  creator
    data (User)
  iteration
    data (IterationModel)
  codebase
    links
      meta
        edit

RelationalData
  area (AreaModel)
  creator (User)
  comments (Comment[])
  assignees (User[])
  linkDicts (LinkDict[])
  iteration (IterationModel)
  totalLinkCount
  wiType (WorkItemType)

LinkDict
  linkName
  links (Link[])
  count
*/


/*
func WriteToDatabase(database *mgo.Database) {
	coll := database.C("workitems")
	if err := coll.Insert(NewWorkItem("newId0", "newType0")); err != nil {
			panic(err)
	}
	fmt.Println("WorkItem inserted successfully!")
}

func ReadFromDatabase(database *mgo.Database, id string) WorkItem {
	var workItem WorkItem
	coll := database.C("workitems")
	err := coll.Find(bson.M{}).One(&workItem) // add "key":"value"" in {} for selectors
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
*/
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