package controllers

import (
	"net/http"

	"github.com/skwongg/jackt/api/responses"
)

//Home is the homepage
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
