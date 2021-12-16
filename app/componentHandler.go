package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qastack-components/dto"
	"qastack-components/service"
	"strconv"

	"github.com/gorilla/mux"
)

type ComponentHandler struct {
	service service.ComponentService
}

func (u ComponentHandler) AddComponent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	var request dto.AddComponentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {

		userId, appError := u.service.AddComponent(request)
		if appError != nil {
			WriteResponse(w, appError.Code, appError.AsMessage())
		} else {
			WriteResponse(w, http.StatusCreated, userId)
		}
	}
}

func (u ComponentHandler) AllComponent(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	projectKey := r.URL.Query().Get("projectKey")

	pageId, _ := strconv.Atoi(page)
	// projectKeyId, _ := strconv.Atoi(projectKey)
	components, err := u.service.AllComponent(projectKey,pageId)

	if err != nil {
		fmt.Println("Inside error" + err.Message)

		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		fmt.Println("Inside error")
		WriteResponse(w, http.StatusOK, components)
	}
}

func (u ComponentHandler) DeleteComponent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// convert the id type from string to int
	id, _ := strconv.Atoi(params["id"])

	result,error := u.service.DeleteComponent(id)
	if error != nil {
		fmt.Println("Inside error" + error.Message)

		WriteResponse(w, error.Code, error.AsMessage())
	} else {
		fmt.Println("Inside error")
		WriteResponse(w, http.StatusOK, result)
	}

}

func (u ComponentHandler) UpdateComponent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// convert the id type from string to int
	id, _ := strconv.Atoi(params["id"])
	var request dto.UpdateComponentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {

		userId, appError := u.service.UpdateComponent(id,request)
		if appError != nil {
			WriteResponse(w, appError.Code, appError.AsMessage())
		} else {
			WriteResponse(w, http.StatusCreated, userId)
		}
	}
}