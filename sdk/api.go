// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sdk

import (
	"github.com/spacemanio/helmet/core/service"
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
