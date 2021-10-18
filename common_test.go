package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	router := gin.Default()
	if withTemplates {
		router.LoadHTMLGlob("templates/*")
	}
	return router
}

func testHTTPResponse(t *testing.T, router *gin.Engine, req *http.Request, f func(recorder *httptest.ResponseRecorder) bool) {
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if !f(recorder) {
		t.Fail()
	}
}

func saveLists() {
	tmpArticleList = articleList
}

func restoreLists() {
	articleList = tmpArticleList
}
