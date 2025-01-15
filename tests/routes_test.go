package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		c.JSON(201, gin.H{"message": "User created successfully!"})
	})

	req, _ := http.NewRequest("POST", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
