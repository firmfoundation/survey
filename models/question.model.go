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
