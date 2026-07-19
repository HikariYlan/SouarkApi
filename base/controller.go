package base

import (
	"net/http"
	"slices"
	"souark/api/services/send"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateFilterFunction func(*http.Request) bson.D

type ValidateDataFunction[Entity any] func(*http.Request) (bool, Entity)

type Routes[Entity any] []string

type Controller[Entity any] struct {
	Repository   Repository[Entity]
	CreateFilter CreateFilterFunction
	ValidateData ValidateDataFunction[Entity]
	Routes       Routes[Entity]
}

func (controller Controller[Entity]) NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	if slices.Contains(controller.Routes, "get") {
		router.HandleFunc("GET /", controller.ReadAll)
	}
	if slices.Contains(controller.Routes, "post") {
		router.HandleFunc("POST /", controller.Create)
	}
	if slices.Contains(controller.Routes, "getId") {
		router.HandleFunc("GET /{id}", controller.Read)
	}
	if slices.Contains(controller.Routes, "delete") {
		router.HandleFunc("DELETE /{id}", controller.Delete)
	}

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

	validation, document := controller.ValidateData(req)

	if validation == false {
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
