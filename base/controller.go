package base

import (
	"encoding/json"
	"net/http"
	"souark/api/services/send"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateFilterFunction func(*http.Request) bson.D

type Controller[Entity any] struct {
	Repository   Repository[Entity]
	CreateFilter CreateFilterFunction
}

func (controller Controller[Entity]) NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", controller.ReadAll)
	router.HandleFunc("POST /", controller.Create)
	router.HandleFunc("GET /{id}", controller.Read)
	router.HandleFunc("DELETE /{id}", controller.Delete)

	return router
}

func (controller Controller[Entity]) ReadAll(res http.ResponseWriter, req *http.Request) {

	filter := controller.CreateFilter(req)

	data, err := controller.Repository.FindAll(filter)

	if err != nil {
		send.Json(nil, res, http.StatusInternalServerError)
		return
	}

	send.Json(data, res, http.StatusOK)
}

func (controller Controller[Entity]) Create(res http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var document Entity
	if err := decoder.Decode(&document); err != nil {
		send.Json(map[string]string{"error": "Invalid data"}, res, http.StatusBadRequest)
		return
	}

	newDocument, err := controller.Repository.InsertOne(document)

	if err != nil {
		send.Json(err, res, http.StatusBadRequest)
		return
	}

	send.Json(newDocument, res, http.StatusCreated)

}

func (controller Controller[Entity]) Read(res http.ResponseWriter, req *http.Request) {

	id := req.PathValue("id")
	document, err := controller.Repository.FindById(id)

	if err != nil {
		send.Json(map[string]string{"error": "Data not found"}, res, http.StatusNotFound)
		return
	}

	send.Json(document, res, http.StatusOK)
}

func (controller Controller[Entity]) Delete(res http.ResponseWriter, req *http.Request) {

	id := req.PathValue("id")

	acknowledged, err := controller.Repository.Delete(id)

	if acknowledged == false || err != nil {
		send.Json(map[string]string{"error": "Data not found"}, res, http.StatusNotFound)
		return
	}

	send.Json(nil, res, http.StatusNoContent)
}
