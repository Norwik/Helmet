// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/norwik/helmet/pkg/util"
	log "github.com/sirupsen/logrus"
)

// Health controller
func Health(w http.ResponseWriter, r *http.Request, _ *GlobalContext) error {
	log.WithFields(log.Fields{
		"status": "i am ok",
	}).Info(`Health check`)

	return util.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"status": "i am ok",
	})
}
