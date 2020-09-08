package routes

import (
	"net/http"

	"github.com/chandanghosh/blog/api/controllers"
)

var (
	userRoutes = []Route{
		{
			URI:     "/users",
			Method:  http.MethodGet,
			Handler: controllers.GetUsers,
		},
		{
			URI:     "/users/{id}",
			Method:  http.MethodGet,
			Handler: controllers.GetUser,
		},
		{
			URI:     "/users",
			Method:  http.MethodPost,
			Handler: controllers.CreateUser,
		},
		{
			URI:     "/users/{id}",
			Method:  http.MethodPut,
			Handler: controllers.UpdateUser,
		},
		{
			URI:     "/users/{id}",
			Method:  http.MethodDelete,
			Handler: controllers.DeleteUser,
		},
	}
)
