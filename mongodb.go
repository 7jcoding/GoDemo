package main

import(
	"fmt"
	"log"
	"gopgk.in/mgo.v2"
	"gopgk.in/mgo.v2/bson"
)

type Persion struct  {
	Id bson.ObjectId `bson:"_id"`
	Name string `bson:"tname"`
	Phone string `bson:"tphone"`
}

const URL = "mongodb:user:pass@192.168.1.18:27019,192.168.1.18:27020,192.168.1.18:27021"

var(
	mgoSession *mgo.Session
	dataBase = "gotest"
)
func main()  {

}
