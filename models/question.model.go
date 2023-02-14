package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Question    string    `json:"question" gorm:"type:varchar(255);not null"`
	IndicatorID uuid.UUID `json:"indicator_id" gorm:"type:uuid;default:uuid_generate_v4();not null"`
	Indicator   Indicator `json:"indicator"`
	SurveyID    uuid.UUID `json:"survey_id" gorm:"type:uuid;default:uuid_generate_v4();not null"`
	Survey      Survey    `json:"survey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Question) CreateQuestion(db *gorm.DB) (*Question, error) {
	var err error
	err = db.Debug().Create(&q).Error
	if err != nil {
		return &Question{}, err
	}
	return q, nil
}

func (q *Question) GetQuestionBySurveyID(db *gorm.DB, uid string) (*[]Question, error) {
	var err error
	questions := []Question{}
	err = db.Preload("Indicator").Preload("Survey").Where("survey_id = ?", uid).Limit(100).Find(&questions).Error
	if err != nil {
		return &[]Question{}, err
	}
	return &questions, err
}

func (q *Question) GetAllUserSurveyIndicatorQuestions(db *gorm.DB, survey_id string, user_id string) []map[string]interface{} {
	var r []map[string]interface{}
	sql := `select u.id as user_id,u.email,q.question as question,i.weight as total_weight, a.answer_point as total_indicator
			from survey_journals as a
			inner join questions as q on a.question_id=q.id
			inner join indicators as i on i.id=q.indicator_id
			inner join users as u on u.id=a.user_id
	 		where a.survey_id=? and a.user_id=?`
	db.Debug().Raw(sql, survey_id, user_id).Limit(100).Find(&r)
	return r
}
