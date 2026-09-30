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

// CreateOAuthData controller
func CreateOAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(r.Body)

	item := &db.OAuthData{}

	err := item.LoadFromJSON(data)

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.ClientID == "" {
		item.ClientID = module.NewCorrelation().UUIDv4()
	}

	if item.ClientSecret == "" {
		item.ClientSecret = module.NewCorrelation().UUIDv4()
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

	if !util.InArray(method.Type, []string{migration.OAuthAuthentication, migration.AnyAuthentication}) {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	item = gc.GetDatabase().CreateOAuthData(item)

	return util.WriteJSON(w, http.StatusCreated, item)
}

// GetOAuthData controller
func GetOAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetOAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Oauth item not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get oauth item data`)

	return util.WriteJSON(w, http.StatusOK, item)
}

// GetOAuthItems controller
func GetOAuthItems(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	items := gc.GetDatabase().GetOAuthDataItems()

	log.Info(`Get oauth items`)

	return util.WriteJSON(w, http.StatusOK, items)
}

// DeleteOAuthData controller
func DeleteOAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetOAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Oauth item not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting oauth item`)

	gc.GetDatabase().DeleteOAuthDataByID(item.ID)

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// UpdateOAuthData controller
func UpdateOAuthData(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item := gc.GetDatabase().GetOAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Oauth item not found`)

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

	if item.ClientID == "" {
		item.ClientID = module.NewCorrelation().UUIDv4()
	}

	if item.ClientSecret == "" {
		item.ClientSecret = module.NewCorrelation().UUIDv4()
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

	if !util.InArray(method.Type, []string{migration.OAuthAuthentication, migration.AnyAuthentication}) {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	gc.GetDatabase().UpdateOAuthDataByID(&item)

	return util.WriteJSON(w, http.StatusOK, item)
}
