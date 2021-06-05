package api

import (
	"os"
	"testing"

	"github.com/adl3879/simple_bank/db/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store controllers.Store) *Server {
	server, err := NewServer(store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
