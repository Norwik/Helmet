// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spacewalkio/helmet/core/model"
	"github.com/spacewalkio/helmet/core/service"
	"github.com/spacewalkio/helmet/core/util"
)

// API type
type API struct {
	httpClient *service.HTTPClient
}

// NewAPI gets a new API instance
func NewAPI(httpClient *service.HTTPClient) *API {
	return &API{
		httpClient: httpClient,
	}
}

// GetEndpoints gets a list of endpoints
func (a *API) GetEndpoints(ctx context.Context, helmetURL, apiKey string) ([]model.Endpoint, error) {
	app := &model.App{}

	response, err := a.httpClient.Get(
		ctx,
		fmt.Sprintf("%s/_api/v1/endpoint", helmetURL),
		map[string]string{},
		map[string]string{"X-Api-Key": apiKey},
	)

	if err != nil {
		return app.Endpoint, err
	}

	if a.httpClient.GetStatusCode(response) != http.StatusOK {
		return app.Endpoint, fmt.Errorf(
			"Invalid response code %d",
			a.httpClient.GetStatusCode(response),
		)
	}

	body, err := a.httpClient.ToString(response)

	if err != nil {
		return app.Endpoint, err
	}

	util.LoadFromJSON(app, []byte(body))

	return app.Endpoint, nil
}
