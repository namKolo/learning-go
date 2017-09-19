package handler

import (
	"net/http"

	model "learning-go/to-do-app/model"
	formatter "learning-go/to-do-app/util"

	db "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"
)

type (
	ItemHandler struct {
		session *db.Session
	}
)

func NewItemHandler(s *db.Session) *ItemHandler {
	return &ItemHandler{s}
}

func (ih ItemHandler) CreateNewItem(w http.ResponseWriter, r *http.Request) {
	text := r.PostFormValue("text")
	item := model.NewItem(text)

	res, err := db.Table("items").Insert(item).RunWrite(ih.session)
	if err != nil {
		formatter.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id := res.GeneratedKeys[0]
	insertedItem, err := db.Table("items").Get(id).Run(ih.session)
	if err != nil {
		formatter.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var finalItem model.Item
	err = insertedItem.One(&finalItem)
	if err != nil {
		formatter.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	formatter.RespondJSON(w, http.StatusAccepted, finalItem)
}

func (ih ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := db.Table("items").Get(id).Run(ih.session)
	if err != nil {
		formatter.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if res.IsNil() {
		formatter.RespondError(w, http.StatusNotFound, "Item not found")
		return
	}

	defer res.Close()
	var item model.Item
	err = res.One(&item)
	if err != nil {
		formatter.RespondError(w, http.StatusInternalServerError, "Something wrong")
		return
	}
	formatter.RespondJSON(w, http.StatusAccepted, item)

}
