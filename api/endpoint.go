// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/norwik/helmet/pkg/util"
	log "github.com/sirupsen/logrus"
)

// GetEndpoints controller
func GetEndpoints(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	configs, err := gc.GetConfigs()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while decoding configs`)

		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	return util.WriteJSON(w, http.StatusOK, configs.App)
}
