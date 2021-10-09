package userControllers

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
type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil{
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


func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	u := models.User{}

	userEmail := p.ByName("email")
	userPassword := p.ByName("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), 8)

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	u.Email = userEmail
	u.Password = hashedPassword

	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

