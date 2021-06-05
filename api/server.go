package api

import (
	"fmt"

	"github.com/adl3879/simple_bank/db/controllers"
	"github.com/adl3879/simple_bank/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for the backend service
type Server struct {
	store      controllers.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store controllers.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("rrrrrrrrrrrrrrr")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.SetupRouter()
	return server, nil
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}

func errorReponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
