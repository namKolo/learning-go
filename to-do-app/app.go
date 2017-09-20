// App has router and db instances
package main

import (
	config "learning-go/to-do-app/config"
	"learning-go/to-do-app/handler"
	"learning-go/to-do-app/model"
	"log"
	"net/http"

	db "github.com/dancannon/gorethink"

	"github.com/gorilla/mux"
)

type App struct {
	Router  *mux.Router
	Session *db.Session
	Hub     *model.Hub
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

	hub := model.NewHub()
	go hub.Run()
	a.Hub = hub

	a.setRouters()
}

func (a *App) setRouters() {
	itemHandler := handler.NewItemHandler(a.Session)
	socketHandler := handler.NewSocketHandler(a.Hub, a.Session)
	a.Delete("/items/{id}", itemHandler.DeleteItem)
	a.Get("/items/{id}", itemHandler.GetItem)
	a.Put("/items/{id}", itemHandler.UpdateItem)
	a.Get("/items", itemHandler.GetItems)
	a.Post("/items", itemHandler.CreateItem)

	a.Router.Handle("/ws/all", socketHandler.GetAllItemsChanges())
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

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
