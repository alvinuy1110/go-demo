package main
import (
	"fmt"
	//"bytes"
	//"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	//"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)
func SetUpRouter() *gin.Engine{
	router := gin.Default()
	return router
}

func TestGetAlbums(t *testing.T) {
	mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	r := SetUpRouter()
	r.GET("/", getAlbums)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	fmt.Println(mockResponse)
	fmt.Println(string(responseData))
	//assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}