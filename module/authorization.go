// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"github.com/norwik/helmet/conf"
	"github.com/norwik/helmet/pkg/util"
	"strings"
)

// Authorization type
type Authorization struct {
}

// Authorize validates http method
func (a *Authorization) Authorize(endpoint conf.Endpoint, httpMethod string) error {
	if util.InArray("ANY", endpoint.Proxy.HTTPMethods) {
		return nil
	}

	httpMethod = strings.ToUpper(httpMethod)

	if !util.InArray(httpMethod, endpoint.Proxy.HTTPMethods) {
		return fmt.Errorf("HTTP method %s not allowed for endpoint %s", httpMethod, endpoint.Name)
	}

	return nil
}
