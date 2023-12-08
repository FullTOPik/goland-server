package userService

import (
	"errors"
	database "server/internal/config"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserCreating
	Id        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
	IsBlocked bool   `json:"isBlocked"`
	UpdatedAt string `json:"updatedAt"`
}

type UserCreating struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	LocationId int    `json:"locationId"`
}

func SelectAllUsers() ([]User, error) {
	userData, err := database.DB.Query("select id, username, password, role, createdAt, updatedAt, isBlocked, locationId from user")
	if err != nil {
		return []User{}, errors.New("fatal read users")
	}

	var users []User

	for userData.Next() {
		var user User
		err := userData.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.IsBlocked, &user.LocationId)
		if err != nil {
			return []User{}, errors.New("fatal read users")
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserById(userId int) (User, error) {
	userRow := database.DB.QueryRow("select id, username, password, role, createdAt, updatedAt, isBlocked, locationId from user where id = ?", userId)

	var user User
	err := userRow.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.IsBlocked, &user.LocationId)
	if err != nil {
		return User{}, errors.New("failed to search user")
	}

	return user, nil
}

func AddUser(user UserCreating) (User, error) {
	execUserRow := database.DB.QueryRow("select username from user where username = ?", user.Username)

	var username string
	err := execUserRow.Scan(&username)
	if err == nil {
		return User{}, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return User{}, errors.New("failed to create a user")
	}

	data, err := database.DB.Exec("insert into user (username, password, role, updatedAt, locationId) values (?, ?, ?, now(), ?)", user.Username, hashedPassword, user.Role, user.LocationId)
	if err != nil {
		return User{}, errors.New("failed to create a user")
	}

	_, err = data.RowsAffected()
	if err != nil {
		return User{}, errors.New("failed to create a user")
	}

	createdUserRow := database.DB.QueryRow("select id, username, password, role, createdAt, updatedAt, isBlocked, locationId from user where username = ?", user.Username)

	var createdUser User
	err = createdUserRow.Scan(&createdUser.Id, &createdUser.Username, &createdUser.Password, &createdUser.Role, &createdUser.CreatedAt, &createdUser.UpdatedAt, &createdUser.IsBlocked, &createdUser.LocationId)
	if err != nil {
		return User{}, errors.New("the user was created but could not be return")
	}

	return createdUser, nil
}

func FindUserByUsernameAndPassword(username string, password string) (User, error) {
	userRow := database.DB.QueryRow("select id, username, password, role, createdAt, updatedAt, isBlocked, locationId from user where username = ?", username)

	var user User
	err := userRow.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.IsBlocked, &user.LocationId)
	if err != nil {
		return User{}, errors.New("failed to search user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return User{}, errors.New("invalid login or password")
	}

	return user, nil
}
