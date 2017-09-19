// App has router and db instances
package main

import (
	config "learning-go/to-do-app/config"
	"learning-go/to-do-app/handler"
	"log"
	"net/http"

	db "github.com/dancannon/gorethink"

	"github.com/gorilla/mux"
)

type App struct {
	Router  *mux.Router
	Session *db.Session
}

func (a *App) Initialize(config *config.Config) {
	session, err := db.Connect(db.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	a.Session = session
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	itemHandler := handler.NewItemHandler(a.Session)
	a.Get("/items/{id}", itemHandler.GetItem)
	a.Post("/items", itemHandler.CreateNewItem)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
