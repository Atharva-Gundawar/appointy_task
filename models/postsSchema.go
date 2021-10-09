package postsModel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct{
	Id					bson.ObjectId	`json:"id" bson:"_id"`
	User_ID				bson.ObjectId	`json:"user_id" bson:"user_id"`
	Caption				string			`json:"caption" bson:"caption"`
	Image_URL			string			`json:"image_url" bson:"image_url"`
	Posted_Timestamp	time.time		`json:"posted_time" bson:"posted_time"`
}
