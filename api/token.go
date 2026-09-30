// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/norwik/helmet/db"
	"github.com/norwik/helmet/module"
	"github.com/norwik/helmet/pkg/util"
	log "github.com/sirupsen/logrus"
)

// CreateToken controller
func CreateToken(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(r.Body)

	if !strings.Contains(string(data), "grant_type=client_credentials") {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	oauthMethod := &module.OAuthAuthMethod{Database: gc.GetDatabase()}

	oauthRecord, err := oauthMethod.ValidateClientCredentials(r.Header.Get("authorization"))

	if err != nil {
		return util.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	item := gc.GetDatabase().CreateOAuthAccessData(&db.OAuthAccessData{
		AccessToken: module.NewCorrelation().UUIDv4(),
		Meta:        "",
		OAuthDataID: oauthRecord.ID,
		ExpireAt:    time.Now().Add(time.Second * 3600).UTC(),
	})

	log.WithFields(log.Fields{
		"time": time.Now().UTC(),
	}).Info(`Cleanup stale access tokens`)

	gc.GetDatabase().CleanupExpiredTokens()

	return util.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"access_token": item.AccessToken,
		"token_type":   "access_token",
		"expires_in":   3600,
	})
}
