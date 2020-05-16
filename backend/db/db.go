package db

import (
	"errors"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/zgordan-vv/robofunding/backend/models"
)

type DB struct {
	*pg.DB
}

var dbInstance *DB

func init() {
	db := pg.Connect(&pg.Options{
		User:     "robouser",
		Password: "robopwd",
		Addr:     "localhost:5432",
		Database: "robofunding",
	})
	if db == nil {
		panic("no database")
	}
	dbInstance = &DB{db}
	must(dbInstance.createUsersTable())
	must(dbInstance.createSessionsTable())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDB() *DB {
	return dbInstance
}

var createOpts = &orm.CreateTableOptions{
	IfNotExists: true,
}

func (db *DB) createUsersTable() error {
	conn := db.Conn()
	defer conn.Close()
	return conn.CreateTable(&models.User{}, createOpts)
}

func (db *DB) CreateUser(user *models.User) error {
	conn := db.Conn()
	defer conn.Close()
	if _, err := db.GetUserByEmail(user.Email); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return conn.Insert(user)
		} else {
			return err
		}
	}
	return errors.New("User already exists")
}

func (db *DB) createSessionsTable() error {
	conn := db.Conn()
	defer conn.Close()
	return conn.CreateTable(&models.Session{}, createOpts)
}

func (db *DB) CreateSession(session *models.Session) error {
	conn := db.Conn()
	defer conn.Close()
	return conn.Insert(session)
}

func (db *DB) GetSession(sessionId string) (*models.Session, error) {
	conn := db.Conn()
	defer conn.Close()
	session := &models.Session{
		Id: sessionId,
	}
	err := db.Select(session)
	return session, err
}

func (db *DB) GetUser(userId string) (*models.User, error) {
	conn := db.Conn()
	defer conn.Close()
	user := &models.User{Id: userId}
	err := db.Select(user)
	return user, err
}

func (db *DB) GetUserByEmail(email string) (*models.User, error) {
	conn := db.Conn()
	defer conn.Close()
	user := &models.User{}
	err := db.Model(user).Where("email = ?", email).Select()
	return user, err
}

func (db *DB) DeleteSession(sessionId string) error {
	conn := db.Conn()
	defer conn.Close()
	return conn.Delete(&models.Session{Id: sessionId})
}
