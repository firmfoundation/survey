package handles

import (
	"encoding/json"
	"net/http"

	"github.com/firmfoundation/survey/initdb"
	"github.com/firmfoundation/survey/models"
	"github.com/firmfoundation/survey/util"
)

func GetAllSurveys(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return util.CustomeError(nil, 405, "Error: Method not allowed.")
	}

	survey := &models.Survey{}
	result, err := survey.GetAllSurveys(initdb.DB)
	if err != nil {
		return util.CustomeError(nil, 500, "Error: unable to Get survey data.")
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}

func GetAllSurveyUsers(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return util.CustomeError(nil, 405, "Error: Method not allowed.")
	}

	survey_id := r.URL.Query().Get("survey_id")
	surveyJournal := &models.SurveyJournal{}
	result := surveyJournal.GetAllSurveyJournalUsers(initdb.DB, survey_id)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}

// func GetAllSurveyUsers(w http.ResponseWriter, r *http.Request) error {
// 	if r.Method != http.MethodGet {
// 		return util.CustomeError(nil, 405, "Error: Method not allowed.")
// 	}
// 	survey_id := r.URL.Query().Get("survey_id")
// 	surveyJournal := &models.SurveyJournal{}
// 	result, err := surveyJournal.GetAllSurveyJournalUsers(initdb.DB, survey_id)
// 	if err != nil {
// 		return util.CustomeError(nil, 500, "Error: unable to Get survey journal data.")
// 	}

// 	w.WriteHeader(http.StatusAccepted)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(result)
// 	return nil
// }
