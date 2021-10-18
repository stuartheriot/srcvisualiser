package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	router := getRouter(true)

	router.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		page, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(page), "<title>SrcViz</title>") > 0

		return statusOK && pageOK
	})
}

func TestRenderJSON(t *testing.T) {
	router := getRouter(true)

	router.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/2", nil)
	req.Header.Add("Accept", "application/json")

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		page, err := ioutil.ReadAll(w.Body)

		log.Println(string(page))

		pageOK := err == nil && strings.Index(string(page), "Go is Cool") > 0

		return statusOK && pageOK
	})
}

func TestRenderXML(t *testing.T) {
	router := getRouter(true)

	router.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/2", nil)
	req.Header.Add("Accept", "application/xml")

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		page, err := ioutil.ReadAll(w.Body)

		log.Println(string(page))

		pageOK := err == nil && strings.Index(string(page), "<Title>Go is Cool") > 0

		return statusOK && pageOK
	})
}
