package router

import (
	"go-login/cmd/api/handlers/createuser"
	"go-login/cmd/api/handlers/getuser"
	"go-login/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users/:id", getuser.Do(app))
	mux.POST("/users", createuser.Do(app))
	return mux
}
