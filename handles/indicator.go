package handles

import (
	"encoding/json"
	"net/http"

	"github.com/firmfoundation/survey/initdb"
	"github.com/firmfoundation/survey/models"
	"github.com/firmfoundation/survey/util"
)

func GetAllIndicators(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return util.CustomeError(nil, 405, "Error: Method not allowed.")
	}

	indicator := &models.Indicator{}
	result, err := indicator.GetAllIndicators(initdb.DB)
	if err != nil {
		return util.CustomeError(nil, 500, "Error: unable to Get indicator data.")
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}
