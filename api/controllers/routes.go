package controllers

import "github.com/skwongg/jackt/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Lift routes
	s.Router.HandleFunc("/lifts", middlewares.SetMiddlewareJSON(s.CreateLift)).Methods("POST")
	s.Router.HandleFunc("/lifts", middlewares.SetMiddlewareJSON(s.GetLifts)).Methods("GET")
	s.Router.HandleFunc("/lifts/{id}", middlewares.SetMiddlewareJSON(s.GetLift)).Methods("GET")
	s.Router.HandleFunc("/lifts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateLift))).Methods("PUT")
	s.Router.HandleFunc("/lifts/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteLift)).Methods("DELETE")
}
