package router

import (
	"html/template"
	"net/http"
	"os"

	"github.com/BankEx/ebh/handlers"
	globalHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func New(handlers *handlers.HandlersService) http.Handler {
	r := mux.NewRouter()
	defineRoutes(r, handlers)

	return defineMiddlewares(r)
}

func defineMiddlewares(r *mux.Router) http.Handler {
	router := globalHandlers.LoggingHandler(os.Stdout, r)
	router = globalHandlers.RecoveryHandler()(router)

	return router
}

func defineRoutes(r *mux.Router, handlers *handlers.HandlersService) {
	apiSubRouter := r.PathPrefix("/api/v1").Subrouter()

	addUnprotectedAPISubrouter(apiSubRouter, handlers)
	//addProtectedAPISubrouter(apiSubRouter)
}

func addUnprotectedAPISubrouter(apiSubRouter *mux.Router, handlers *handlers.HandlersService) {

	apiSubRouter.HandleFunc("/sertificates", handlers.CreateSert).Methods("POST")

	apiSubRouter.HandleFunc("/sertificates", handlers.GetAll).Methods("GET")

	apiSubRouter.HandleFunc("/help", getHelp).Methods("GET")
	apiSubRouter.HandleFunc("/swagger.yml", swagger).Methods("GET")

	//apiSubRouter.HandleFunc("/healthcheck", handlers.HealthCheck).Methods("GET")
}

func getHelp(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./docs/index.html")
	t.Execute(w, nil)
}

func swagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "./docs/swagger.yml")
}
