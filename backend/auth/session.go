package auth

import (
	"time"

	"github.com/zgordan-vv/robofunding/backend/db"
	"github.com/zgordan-vv/robofunding/backend/models"
	"github.com/zgordan-vv/robofunding/backend/utils"
)

func NewSession(userId string) (*models.Session, error) {
	session := &models.Session{
		Id:      utils.NewId(),
		UserId:  userId,
		Expires: time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	err := db.GetDB().CreateSession(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func GetSession(sessionId string) (*models.Session, error) {
	return db.GetDB().GetSession(sessionId)
}
