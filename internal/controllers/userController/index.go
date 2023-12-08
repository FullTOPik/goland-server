package userController

import (
	"encoding/json"
	"net/http"
	userService "server/internal/services/userService"
	"server/pkg/server"
	"strconv"
)

func GetRouters() server.RouterType {
	var Router server.RouterType

	Router.Get("", func(w http.ResponseWriter, r *http.Request, params server.Params) {
		users, err := userService.SelectAllUsers()
		if err != nil {
			server.BadRequestAnswer(w, err.Error())
		}

		json.NewEncoder(w).Encode(users)
	})

	Router.Get("/:id", func(w http.ResponseWriter, r *http.Request, params server.Params) {
		var id = params["id"].(string)
		
		userIdNumber, err := strconv.Atoi(id)
		if err != nil {
			server.BadRequestAnswer(w, "Cannot to find user")
		}

		user, err := userService.GetUserById(userIdNumber)
		if err != nil {
			server.BadRequestAnswer(w, err.Error())
		}

		json.NewEncoder(w).Encode(user)
	})

	Router.Post("/one", func(w http.ResponseWriter, r *http.Request, params server.Params) {
		var userBody userService.UserCreating
		json.NewDecoder(r.Body).Decode(&userBody)

		user, err := userService.AddUser(userBody)
		if err != nil {
			server.BadRequestAnswer(w, err.Error())
		}

		json.NewEncoder(w).Encode(user)
	})

	return Router
}
