// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/norwik/helmet/pkg/util"
	log "github.com/sirupsen/logrus"
)

// Me controller
func Me(w http.ResponseWriter, r *http.Request, _ *GlobalContext) error {
	log.Info(`Get access status and permissions`)

	return util.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"id":         "1",
		"type":       "ok",
		"authMethod": "",
		"services":   "",
	})
}
