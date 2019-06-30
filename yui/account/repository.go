package account

import (
	"database/sql"
	"fmt"
)

type (
	UserRepo interface {
		FindByUsername(email string) (*User, error)
		Save(user *User) error
	}
)

// UserRepo SQL Database Implementation
type UserRepoDB struct {
	DB *sql.DB
}

// UserRepo FindByUsername DB implementation
func (r *UserRepoDB) FindByUsername(username string) (*User, error) {
	user := User{}
	result := r.DB.QueryRow(fmt.Sprintf("select password from users where username='%s'", username))
	err := result.Scan(&user.Password)
	return &user, err
}

// UserRepo Save DB implementation
func (r *UserRepoDB) Save(user *User) error {
	_, err := r.DB.Query(fmt.Sprintf("insert into users values ('%s', '%s', '%s', '%s', '%s', '%s')",
		user.ID,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password))
	return err
}
