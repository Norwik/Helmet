// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"fmt"
	"strings"

	"github.com/clivern/drifter/core/model"
	"github.com/clivern/drifter/core/util"
)

// Authorization type
type Authorization struct {
}

// Authorize validates http method
func (k *KeyBasedAuthMethod) Authorize(endpoint model.Endpoint, httpMethod string) (bool, error) {
	if util.InArray("ANY", endpoint.Proxy.HTTPMethods) {
		return true, nil
	}

	httpMethod = strings.ToUpper(httpMethod)

	if !util.InArray(httpMethod, endpoint.Proxy.HTTPMethods) {
		return false, fmt.Errorf("HTTP method %s not allowed for endpoint %s", httpMethod, endpoint.Name)
	}

	return true, nil
}
