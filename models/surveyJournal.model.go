package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SurveyJournal struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	QuestionID  uuid.UUID `json:"question_id" gorm:"type:varchar(255);not null"`
	Question    Question  `json:"question"`
	SurveyID    uuid.UUID `json:"survey_id" gorm:"type:uuid;default:uuid_generate_v4();not null"`
	Survey      Survey    `json:"survey"`
	AnswerPoint int       `json:"answer_point"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4();not null"`
	User        User      `json:"user"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SurveyResult struct {
	SurveyID uuid.UUID `json:"survey_id"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
	Result   []Result  `json:"result"`
}

type Result struct {
	QuestionID uuid.UUID `json:"question_id"`
	Answer     int       `json:"answer"`
}

func (q *SurveyJournal) CreateSurveyJournal(db *gorm.DB) (*SurveyJournal, error) {
	var err error
	err = db.Debug().Create(&q).Error
	if err != nil {
		return &SurveyJournal{}, err
	}
	return q, nil
}

func (q *SurveyJournal) CreateSurveyResult(db *gorm.DB, batch []SurveyJournal, u *User) (*SurveyJournal, error) {

	// err = db.Debug().Create(&batch).Error
	// if err != nil {
	// 	return &SurveyJournal{}, err
	// }
	// return q, nil
	err := db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		var err error

		//tx 1 check if user exists
		user := User{}
		if err = tx.Where("email = ?", u.Email).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {

				if err = tx.Debug().Create(u).Error; err != nil {
					return err
				}

				//update batch with the newly created userid
				for i := range batch {
					batch[i].UserID = u.ID
				}
				//reset user to empty
				user.ID = uuid.Nil
			} else {
				return err
			}

		}

		if user.ID.String() != uuid.Nil.String() {
			//update batch with existing userid
			for i := range batch {
				batch[i].UserID = user.ID
			}
		}

		//tx - 2
		if err = tx.Debug().Create(&batch).Error; err != nil {
			// return any error will rollback
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return &SurveyJournal{}, err
	}

	return q, nil
}

func (q *SurveyJournal) GetAllSurveyJournalUsers(db *gorm.DB, survey_id string) []map[string]interface{} {
	var r []map[string]interface{}
	sql := `select a.survey_id as survey_id, s.name as survey_name, a.user_id as user_id, u.email, u.full_name, 
	sum(i.weight) as total_weight, sum(a.answer_point) as total_indicator
	from survey_journals as a
	inner join questions as q on a.question_id=q.id
	inner join indicators as i on i.id=q.indicator_id
	inner join users as u on u.id=a.user_id
	inner join surveys as s on s.id=a.survey_id
	where a.survey_id=?
	group by a.survey_id, s.name, a.user_id, u.email, u.full_name`
	db.Debug().Raw(sql, survey_id).Limit(100).Find(&r)
	return r
}

func GetUserSurveyIndicators(db *gorm.DB, survey_id string, user_id string) []map[string]interface{} {
	var r []map[string]interface{}
	sql := `select i.name as indicator,sum(i.weight) as total_weight, sum(a.answer_point) as total_indicator
	 from survey_journals as a
	 inner join questions as q on a.question_id=q.id
	 inner join indicators as i on i.id=q.indicator_id
	 where a.survey_id=? and a.user_id=?
	 group by i.name`
	db.Debug().Raw(sql, survey_id, user_id).Limit(100).Find(&r)
	return r
}

/*
func (q *SurveyJournal) GetAllSurveyJournalUsers(db *gorm.DB, survey_id string) (*[]SurveyJournal, error) {
	var err error
	surveyJournal := []SurveyJournal{}
	err = db.Debug().Preload("User").Where("survey_id = ?", survey_id).Limit(100).Find(&surveyJournal).Error
	if err != nil {
		return &[]SurveyJournal{}, err
	}
	return &surveyJournal, err
}
*/
