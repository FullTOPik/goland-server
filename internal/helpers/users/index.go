package userHelpers

import (
	"log"
	userservice "server/internal/services/userService"
)

func PrintUser(user userservice.User) {
	log.Printf("%d %s %s %s %s %s %v %d", user.Id, user.Username, user.Password, user.Role, user.UpdatedAt, user.CreatedAt, user.IsBlocked, user.LocationId)
}
