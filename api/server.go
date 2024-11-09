package api

import (
	"github.com/SohamKanji/task-management-system-go/api/validators"
	db "github.com/SohamKanji/task-management-system-go/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	router *gin.Engine
	store  *db.Store
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		router: gin.Default(),
		store:  store,
	}
	server.setup()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validstatus", validators.ValidStatus)
	}

	return server
}

func (server *Server) setup() {
	server.router.GET("/tasks", server.listTasks)
	server.router.POST("/tasks", server.createTask)
	server.router.GET("/tasks/:id", server.getTask)
	server.router.PUT("/tasks/:id", server.updateTask)
	server.router.DELETE("/tasks/:id", server.deleteTask)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
