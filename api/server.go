package api

import (
	"github.com/adl3879/simple_bank/db/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTp requests for the backend service
type Server struct {
	store  controllers.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store controllers.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}

func errorReponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
