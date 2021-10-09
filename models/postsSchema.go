package postsModel

import "gopkg.in/mgo.v2/bson"
import "time"

type User struct{
	Id					bson.ObjectId	`json:"id" bson:"_id"`
	Caption				string			`json:"caption" bson:"caption"`
	Image_URL			string			`json:"image_url" bson:"image_url"`
	Posted_Timestamp	int				`json:"posted_time" bson:"posted_time"`
}
