package resolvers

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/zgordan-vv/robofunding/backend/db"
	"github.com/zgordan-vv/robofunding/backend/models"
)

type UserResolver struct {
	db *db.DB
	m  *models.User
}

func (u *UserResolver) Id() *graphql.ID {
	return strToId(u.m.Id)
}

func (u *UserResolver) Email() string {
	return u.m.Email
}
