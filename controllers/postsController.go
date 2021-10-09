package postControllers

import (

"fmt"
"encoding/json"
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"github.com/Atharva-Gundawar/appointy_task/models"
"golang.org/x/crypto/bcrypt"
"net/http"

)

// Getting mongo Session 
type PostController struct{
	session *mgo.Session
}

func NewPostController(s *mgo.Session) *PostController{
return &PostController{s}
}

func (uc PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	p := models.Post{}

	if err := uc.session.DB("mongo-golang").C("posts").FindId(oid).One(&p); err != nil{
		w.WriteHeader(404)
		return
	}

	uj, err :=json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc PostController) GetPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}
	p := models.Post{}
  	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil{
	w.WriteHeader(404)
	return
   	}

	posts :=   u.FieldByName("Posts")
	posts_ids := [len(posts)]bson.ObjectId

	for i := 0; i < len(posts); i++ {
        if err := uc.session.DB("mongo-golang").C("posts").FindId(oid).One(&p); err != nil{
			posts_ids[i] := u.FieldByName("id")
			return
			}
    }
	
	uj, err :=json.Marshal(posts_ids)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc PostController) CreatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	
	p := models.Post{}

	json.NewDecoder(r.Body).Decode(&p)

	p.Id = bson.NewObjectId()

	pc.session.DB("mongo-golang").C("posts").Insert(p)

	uj, err := json.Marshal(p)

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

