package initdb

import (
	"fmt"

	"github.com/firmfoundation/survey/models"
)

func Migrate() {
	DB.AutoMigrate(
		&models.Survey{},
		&models.Indicator{},
		&models.Question{},
		&models.SurveyJournal{},
	)

	fmt.Println(" Migration complete")
}
