package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"osapp/handlers"
)

const (
	adminAPIPathPrefix = "/admin"
)



// registerAdminRouter - Add handler functions for each service REST API routes.
func RegisterAdminRouter(router *mux.Router) {

	adminAPI := handlers.AdminAPIHandlers{}
	// Admin router
	adminRouter := router.PathPrefix(adminAPIPathPrefix).Subrouter()
	// Version handler
	adminRouter.Methods(http.MethodGet).Path("/version").HandlerFunc(adminAPI.VersionHandler)

	// v1
	adminV1Router := adminRouter.PathPrefix("/v1").Subrouter()

	/// Service operations
	// Service info
	adminV1Router.Methods(http.MethodGet).Path("/service").HandlerFunc(adminAPI.ServiceInfoHandler)
	// Service restart and stop
	adminV1Router.Methods(http.MethodPost).Path("/service").HandlerFunc(adminAPI.ServiceStopOrRestartHandler)


	/// Config operations
	// Get config
	adminV1Router.Methods(http.MethodGet).Path("/config").HandlerFunc(adminAPI.GetConfigHandler)
	// Set config
	adminV1Router.Methods(http.MethodPut).Path("/config").HandlerFunc(adminAPI.SetConfigHandler)


	// -- user APIs --
	// Add user IAM
	adminV1Router.Methods(http.MethodPut).Path("/add-user").HandlerFunc(adminAPI.AddUser)
	// Get user IAM
	adminV1Router.Methods(http.MethodPut).Path("/get-user").HandlerFunc(adminAPI.GetUser)
	// Remove user IAM
	adminV1Router.Methods(http.MethodDelete).Path("/remove-user").HandlerFunc(adminAPI.RemoveUser)
	// List users
	adminV1Router.Methods(http.MethodGet).Path("/list-users").HandlerFunc(adminAPI.ListUsers)

	// -- Top APIs --
	// If none of the routes match, return error.
	adminV1Router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
}
