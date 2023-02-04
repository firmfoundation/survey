package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/firmfoundation/survey/initdb"
	"github.com/firmfoundation/survey/models"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

const tempDir = "templates"

func Index(w http.ResponseWriter, r *http.Request) error {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)

	if r.Method != http.MethodGet {
		return CustomeError(nil, 405, "Error: Method not allowed.")
	}

	t, _ := template.ParseFiles(exPath + "/templates/index.html")
	t.Execute(w, "")
	return nil
}

func HandleCreateSurvey(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return CustomeError(nil, 405, "Error: Method not allowed.")
	}

	survey := &models.Survey{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(survey)
	if err != nil {
		return CustomeError(nil, 500, "Error: parsing in one or more submitted body fields.")
	}
	s, err := survey.SaveSurvey(initdb.DB)
	if err != nil {
		return CustomeError(nil, 500, "Error: unable to save survey data.")
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/octet-stream")
	json.NewEncoder(w).Encode(s)

	return nil
}

func HandleCreateIndicator(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return CustomeError(nil, 405, "Error: Method not allowed.")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return CustomeError(nil, http.StatusUnprocessableEntity, "Error: parsing in one or more submitted body fields.")
	}

	indicator := &models.Indicator{}
	err = json.Unmarshal(body, &indicator)
	if err != nil {
		return CustomeError(nil, 500, "Error: decoding in one or more submitted body fields.")
	}

	result, err := indicator.CreateIndicator(initdb.DB)
	if err != nil {
		return CustomeError(nil, 500, "Error: unable to create indicator data.")
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}

func HandleCreateQuestion(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return CustomeError(nil, 405, "Error: Method not allowed.")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return CustomeError(nil, http.StatusUnprocessableEntity, "Error: parsing in one or more submitted body fields.")
	}

	question := &models.Question{}
	err = json.Unmarshal(body, &question)
	if err != nil {
		return CustomeError(nil, 500, "Error: decoding in one or more submitted body fields.")
	}

	result, err := question.CreateQuestion(initdb.DB)
	if err != nil {
		return CustomeError(nil, 500, "Error: unable to create questionnaire data.")
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}

func HandleCreateSurveyJournal(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return CustomeError(nil, 405, "Error: Method not allowed.")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return CustomeError(nil, http.StatusUnprocessableEntity, "Error: parsing in one or more submitted body fields.")
	}

	journal := &models.SurveyJournal{}
	err = json.Unmarshal(body, &journal)
	if err != nil {
		return CustomeError(nil, 500, "Error: decoding in one or more submitted body fields.")
	}

	result, err := journal.CreateSurveyJournal(initdb.DB)
	if err != nil {
		return CustomeError(nil, 500, "Error: unable to create survey journal data.")
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return CustomeError(nil, 405, "Error: Method not allowed.")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return CustomeError(nil, http.StatusUnprocessableEntity, "Error: parsing in one or more submitted body fields.")
	}

	user := &models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return CustomeError(nil, 500, "Error: decoding in one or more submitted body fields.")
	}

	result, err := user.CreateUser(initdb.DB)
	if err != nil {
		return CustomeError(nil, 500, "Error: unable to create user data.")
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}
