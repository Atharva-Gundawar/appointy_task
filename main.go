package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	
	"github.com/Atharva-Gundawar/appointy_task/controllers"
)
	
func main(){

	r := httprouter.New()

	uc := userControllers.NewUserController(getSession())
	pc := postControllers.NewPostController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.GET("/post/:id", pc.GetPost)
	r.GET("/post/users/:id", pc.GetPosts)
	r.POST("/post", pc.CreatePost)

	http.ListenAndServe("localhost:9000", r)
}


func getSession() *mgo.Session{

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s
}
