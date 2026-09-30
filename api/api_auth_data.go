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

// CreateKeyBasedAuthData controller
func CreateKeyBasedAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(r.Body)

	item := &db.KeyBasedAuthData{}

	err := item.LoadFromJSON(data)

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.APIKey == "" {
		item.APIKey = module.NewCorrelation().UUIDv4()
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

	if !util.InArray(method.Type, []string{migration.KeyAuthentication, migration.AnyAuthentication}) {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	item = gc.GetDatabase().CreateKeyBasedAuthData(item)

	return util.WriteJSON(w, http.StatusCreated, item)
}

// GetKeyBasedAuthData controller
func GetKeyBasedAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetKeyBasedAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key item not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get API key item`)

	return util.WriteJSON(w, http.StatusOK, item)
}

// GetKeysBasedAuthData controller
func GetKeysBasedAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	keys := gc.GetDatabase().GetKeyBasedAuthItems()

	log.Info(`Get api keys`)

	return util.WriteJSON(w, http.StatusOK, keys)
}

// DeleteKeyBasedAuthData controller
func DeleteKeyBasedAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetKeyBasedAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an API key`)

	gc.GetDatabase().DeleteKeyBasedAuthDataByID(item.ID)

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// UpdateKeyBasedAuthData controller
func UpdateKeyBasedAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetKeyBasedAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	data, _ := ioutil.ReadAll(r.Body)

	err := item.LoadFromJSON(data)

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.APIKey == "" {
		item.APIKey = module.NewCorrelation().UUIDv4()
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

	if !util.InArray(method.Type, []string{migration.KeyAuthentication, migration.AnyAuthentication}) {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	gc.GetDatabase().UpdateKeyBasedAuthDataByID(&item)

	return util.WriteJSON(w, http.StatusOK, item)
}
