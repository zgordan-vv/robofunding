package models

type User struct {
	Id       string    `json:"_id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Fundings []Funding `json:"fundings"`
}

type Session struct {
	Id      string
	UserId  string
	Expires int64
}

type AuthInfo struct {
	User      User
	SessionId string
}

type CreateUserInput struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Funding struct {
	Id          string  `json:"_id"`
	RobotId     string  `json:"robot_id"`
	Sum         float64 `json:"sum"`
	Description string  `json:"description"`
}

type ContextValue struct {
	SessionId   string
	CurrentUser *User
}
