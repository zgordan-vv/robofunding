package resolvers

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"

	//"github.com/zgordan-vv/robofunding/backend/auth"
	"github.com/zgordan-vv/robofunding/backend/auth"
	"github.com/zgordan-vv/robofunding/backend/db"
	"github.com/zgordan-vv/robofunding/backend/models"
	"github.com/zgordan-vv/robofunding/backend/utils"
)

type Resolver struct {
	db *db.DB
}

func NewResolver() *Resolver {
	return &Resolver{
		db: db.GetDB(),
	}
}

func (r *Resolver) CreateUser(ctx context.Context, args struct{ Input *models.CreateUserInput }) (*UserResolver, error) {
	fmt.Println("Creating user:", args.Input)
	input := args.Input
	if input.Email == "" {
		return nil, errors.New("email shoud not be empty")
	}
	if len(input.Password) < 6 {
		return nil, errors.New("password should be 6 chars or more")
	}
	if input.Password != input.PasswordConfirmation {
		return nil, errors.New("password should be confirmed")
	}
	password := hashPwd(input.Password)
	user := &models.User{
		Id:       utils.NewId(),
		Email:    input.Email,
		Password: password,
	}
	err := r.db.CreateUser(user)
	if err != nil {
		return nil, err
	}
	userResolver := &UserResolver{
		db: r.db,
		m:  user,
	}
	return userResolver, nil
}

var authErr = errors.New("Password or email is incorrect")

func (r *Resolver) DoLogin(ctx context.Context, args struct{ Input *models.UserLoginInput }) (string, error) {
	if sessionId := ctx.Value("ctx").(models.ContextValue).SessionId; sessionId != "" {
		return sessionId, nil
	}
	email := args.Input.Email
	password := args.Input.Password
	user, err := r.db.GetUserByEmail(email)
	if err != nil {
		return "", authErr
	}
	hashedPassword := hashPwd(password)
	if hashedPassword != user.Password {
		return "", authErr
	}
	user.Password = ""
	session, err := auth.NewSession(user.Id)
	if err != nil {
		return "", err
	}
	return session.Id, nil
}

func (r *Resolver) GetMe(ctx context.Context) (*UserResolver, error) {
	me := ctx.Value("ctx").(models.ContextValue).CurrentUser
	fmt.Println("ME:", me)
	if me == nil {
		me = &models.User{}
	}
	return &UserResolver{
		db: r.db,
		m:  me,
	}, nil
}

func hashPwd(pwd string) string {
	h := fnv.New64a()
	h.Write([]byte(pwd))
	return fmt.Sprintf("%x", h.Sum64())
}
