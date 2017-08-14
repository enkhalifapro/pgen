package server

import (
	"github.com/enkhalifapro/pgen/controllers"
)

// Controller collection struct for server.
type Controller struct {
	User *controllers.UserController                 `inject:""`
}
