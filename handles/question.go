package handles

import (
	"encoding/json"
	"net/http"

	"github.com/firmfoundation/survey/initdb"
	"github.com/firmfoundation/survey/models"
	"github.com/firmfoundation/survey/util"
)

func GetUserSurveyIndicatorQuestions(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return util.CustomeError(nil, 405, "Error: Method not allowed.")
	}

	survey_id := r.URL.Query().Get("survey_id")
	user_id := r.URL.Query().Get("user_id")

	question := &models.Question{}
	result := question.GetAllUserSurveyIndicatorQuestions(initdb.DB, survey_id, user_id)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}
