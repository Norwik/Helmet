// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spacewalkio/helmet/core/component"
	"github.com/spacewalkio/helmet/core/model"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// CreateToken controller
func CreateToken(c echo.Context, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(c.Request().Body)

	if !strings.Contains(string(data), "grant_type=client_credentials") {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	oauthMethod := &component.OAuthAuthMethod{Database: gc.GetDatabase()}

	oauthRecord, err := oauthMethod.ValidateClientCredentials(c.Request().Header.Get("authorization"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	item := gc.GetDatabase().CreateOAuthAccessData(&model.OAuthAccessData{
		AccessToken: component.NewCorrelation().UUIDv4(),
		Meta:        "",
		OAuthDataID: oauthRecord.ID,
		ExpireAt:    time.Now().Add(time.Second * 3600).UTC(),
	})

	log.WithFields(log.Fields{
		"time": time.Now().UTC(),
	}).Info(`Cleanup stale access tokens`)

	gc.GetDatabase().CleanupExpiredTokens()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": item.AccessToken,
		"token_type":   "access_token",
		"expires_in":   3600,
	})
}
