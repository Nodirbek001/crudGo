package datafile

import (
	"crut/crud/models"
	"encoding/json"
	"io/ioutil"
	"log"
)

func SaveData(user models.User) models.Successfull {
	file, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Println("Error at SaveDate")
		return models.Successfull{
			Success: "",
			Err:     err,
		}
	}
	err = ioutil.WriteFile("/home/nodirbek/git.file/crudGo/crud/Base/base.json", file, 0644)
	if err != nil {
		log.Println("error at WriteFile")
		return models.Successfull{
			Success: "",
			Err:     err,
		}
	}
	return models.Successfull{
		Success: "successFull",
		Err:     nil,
	}
}
