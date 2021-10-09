package usersModel

import "gopkg.in/mgo.v2/bson"

type User struct{
	Id				bson.ObjectId				`json:"id" bson:"_id"`
	Name			string						`json:"name" bson:"name"`
	Email			string						`json:"email" bson:"email"`
	Password		int							`json:"password" bson:"password"`
	
	Posts			Pets []struct {
					Post_ID 		string 				`json:"id" bson:"_id"`
					} 							`json:"posts" bson:"posts"`

}
