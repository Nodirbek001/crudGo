package controller

import (
	"crut/crud/models"
	datafile "crut/crud/storage/dataFile"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (api *api) userSignUp(w http.ResponseWriter, r *http.Request) {
	var (
		body models.User
		res  models.Successfull
	)
	id := uuid.New().String()
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.Write([]byte("error at UserSignUp"))
		return
	}
	res=datafile.SaveData(models.User{
		Id: id,
		Name: body.Name,
		Books: body.Books,
	})
	writeJson(w, res)
}

func Sign(w http.ResponseWriter, r *http.Request){
	var (
		body models.User
		res models.Successfull
	) 
	id:=uuid.New().String()
	if err:=json.NewDecoder(r.Body).Decode(&body);err!=nil {
	w.Write([]byte("error at signin"))
	return	
	}
	res=datafile.SaveData(models.User{
		Id: id,
		Name: body.Name,
		Books: body.Books,
	})
	writeJson(w,res)
}