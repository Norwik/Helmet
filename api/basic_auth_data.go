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
	"github.com/norwik/helmet/migration"
	"github.com/norwik/helmet/module"
	"github.com/norwik/helmet/pkg/util"
	log "github.com/sirupsen/logrus"
)

// CreateBasicAuthData controller
func CreateBasicAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(r.Body)

	item := &db.BasicAuthData{}

	err := item.LoadFromJSON(data)

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.Username == "" {
		item.Username = module.NewCorrelation().UUIDv4()
	}

	if item.Password == "" {
		item.Password = module.NewCorrelation().UUIDv4()
	}

	err = item.Validate()

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method := gc.GetDatabase().GetAuthMethodByID(item.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": method.ID,
		}).Info(`Auth method not found`)

		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Auth method not found",
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if !util.InArray(method.Type, []string{migration.BasicAuthentication, migration.AnyAuthentication}) {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	item = gc.GetDatabase().CreateBasicAuthData(item)

	return util.WriteJSON(w, http.StatusCreated, item)
}

// GetBasicAuthData controller
func GetBasicAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetBasicAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth item not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get a basic auth item`)

	return util.WriteJSON(w, http.StatusOK, item)
}

// GetBasicAuthItems controller
func GetBasicAuthItems(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	items := gc.GetDatabase().GetBasicAuthItems()

	log.Info(`Get basic auth items`)

	return util.WriteJSON(w, http.StatusOK, items)
}

// DeleteBasicAuthData controller
func DeleteBasicAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetBasicAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth item not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting a basic auth item`)

	gc.GetDatabase().DeleteBasicAuthDataByID(item.ID)

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// UpdateBasicAuthData controller
func UpdateBasicAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetBasicAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth item not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	data, _ := ioutil.ReadAll(r.Body)

	err := item.LoadFromJSON(data)

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.Username == "" {
		item.Username = module.NewCorrelation().UUIDv4()
	}

	if item.Password == "" {
		item.Password = module.NewCorrelation().UUIDv4()
	}

	err = item.Validate()

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method := gc.GetDatabase().GetAuthMethodByID(item.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Auth method not found",
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if !util.InArray(method.Type, []string{migration.BasicAuthentication, migration.AnyAuthentication}) {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	gc.GetDatabase().UpdateBasicAuthDataByID(&item)

	return util.WriteJSON(w, http.StatusOK, item)
}
