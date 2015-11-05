package oauth2

import (
	"errors"
	"time"
)

// Client ...
type Client struct {
	ID          int
	ClientID    string `sql:"type:varchar(254);unique;not null"`
	Password    string `sql:"type:varchar(60);not null"`
	RedirectURI string `sql:"type:varchar(200)"`
}

// Scope ...
type Scope struct {
	ID          int
	Scope       string `sql:"type:varchar(200);unique;not null"`
	Description string
	IsDefault   bool `sql:"default:false"`
}

// User ...
type User struct {
	ID        int    `json:"id"`
	Username  string `sql:"type:varchar(254);unique;not null" json:"username"`
	Password  string `sql:"type:varchar(60);not null" json:"password"`
	FirstName string `sql:"type:varchar(254);unique;not null" json:"first_name"`
	LastName  string `sql:"type:varchar(254);unique;not null" json:"last_name"`
}

// RefreshToken ...
type RefreshToken struct {
	ID           int
	RefreshToken string    `sql:"type:varchar(40);unique;not null"`
	ExpiresAt    time.Time `sql:"not null"`
}

// AccessToken ...
type AccessToken struct {
	ID             int
	AccessToken    string    `sql:"type:varchar(40);unique;not null"`
	ExpiresAt      time.Time `sql:"not null"`
	ClientID       int       `sql:"index"`
	UserID         int       `sql:"index"`
	RefreshTokenID int       `sql:"index"`
	Scopes         []Scope   `gorm:"many2many:access_token_scopes"`
	Client         Client
	User           User
	RefreshToken   RefreshToken
}

// AuthCode ...
type AuthCode struct {
	ID          int
	Code        string    `sql:"type:varchar(40);unique;not null"`
	RedirectURI string    `sql:"type:varchar(200)"`
	ExpiresAt   time.Time `sql:"not null"`
	ClientID    int       `sql:"index;not null"`
	UserID      int       `sql:"index"`
	Scopes      []Scope   `gorm:"many2many:auth_code_scopes"`
	Client      Client
	User        User
}

// Validate validates user data
func (user *User) Validate() error {
	if user.Username == "" {
		return errors.New("username required")
	}

	if user.Password == "" {
		return errors.New("password required")
	}

	if user.FirstName == "" {
		return errors.New("first_name required")
	}

	if user.LastName == "" {
		return errors.New("last_name required")
	}

	return nil
}