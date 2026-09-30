// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/norwik/helmet/db"
	"github.com/norwik/helmet/pkg/util"
	log "github.com/sirupsen/logrus"
)

// CreateAuthMethod controller
func CreateAuthMethod(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(r.Body)

	method := &db.AuthMethod{}

	err := method.LoadFromJSON(data)

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	err = method.Validate()

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method = gc.GetDatabase().CreateAuthMethod(method)

	return util.WriteJSON(w, http.StatusCreated, method)
}

// GetAuthMethod controller
func GetAuthMethod(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	method := gc.GetDatabase().GetAuthMethodByID(id)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get an auth method`)

	return util.WriteJSON(w, http.StatusOK, method)
}

// GetAuthMethods controller
func GetAuthMethods(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	log.Info(`Get auth methods`)

	methods := gc.GetDatabase().GetAuthMethods()

	return util.WriteJSON(w, http.StatusOK, methods)
}

// DeleteAuthMethod controller
func DeleteAuthMethod(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	method := gc.GetDatabase().GetAuthMethodByID(id)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an auth method`)

	gc.GetDatabase().DeleteAuthMethodByID(method.ID)

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// UpdateAuthMethod controller
func UpdateAuthMethod(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	method := gc.GetDatabase().GetAuthMethodByID(id)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	data, _ := ioutil.ReadAll(r.Body)

	err := method.LoadFromJSON(data)

	method.ID = id

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	err = method.Validate()

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Update an auth method`)

	gc.GetDatabase().UpdateAuthMethodByID(&method)

	return util.WriteJSON(w, http.StatusOK, method)
}
