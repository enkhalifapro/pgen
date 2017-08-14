package server

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// prepare structural logging.
var log = logrus.WithFields(logrus.Fields{"component": "server"})

// Server which handles API requests.
type Server struct {
	Engine     *gin.Engine `inject:""`
	Controller *Controller `inject:""`
	Middleware *Middleware `inject:""`
}

// Run server
func (s *Server) Run() error {
	// users APIs
	s.Engine.GET("/", s.Controller.User.Root)
	s.Engine.GET("/api/v1/users/id/:id", s.Controller.User.GetById)
	s.Engine.GET("/api/v1/users/email/:email", s.Controller.User.GetByEmail)
	s.Engine.GET("/api/v1/users/name/:name", s.Controller.User.GetBySlug)
	return s.Engine.Run(fmt.Sprintf("%v:%v", viper.GetString("host"), viper.GetString("port")))
}
