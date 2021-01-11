// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"reflect"
	"strings"
)

// InArray check if value is on array
func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

// Unset remove element at position i
func Unset(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}

// LoadFromJSON update object from json
func LoadFromJSON(item interface{}, data []byte) error {
	err := json.Unmarshal(data, &item)
	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func ConvertToJSON(item interface{}) (string, error) {
	data, err := json.Marshal(&item)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// IsEmpty validate if string is empty or not
func IsEmpty(item string) bool {
	if strings.TrimSpace(item) == "" {
		return true
	}
	return false
}

// Rand gets a random number
func Rand(min, max int) int {
	return rand.Intn(max-min) + min
}

// GetVal gets a value from a hash map
func GetVal(hash map[string]string, key, def string) string {
	if val, ok := hash[key]; ok && val != "" {
		return val
	}

	return def
}

// MergeMaps merges two maps
func MergeMaps(m1, m2 map[string]string) map[string]string {
	for k, v := range m2 {
		m1[k] = v
	}

	return m1
}

// Base64Decode decodes 64 encoded string
func Base64Decode(encoded string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(encoded)
	return string(result), err
}

// Base64Encode encodes a string
func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
