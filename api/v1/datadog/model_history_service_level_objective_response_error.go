/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// HistoryServiceLevelObjectiveResponseError A service level objective response containing the requested history.
type HistoryServiceLevelObjectiveResponseError struct {
	// human readable error
	Error *string `json:"error,omitempty"`
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveResponseError) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveResponseError) GetErrorOk() (string, bool) {
	if o == nil || o.Error == nil {
		var ret string
		return ret, false
	}
	return *o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveResponseError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *HistoryServiceLevelObjectiveResponseError) SetError(v string) {
	o.Error = &v
}

type NullableHistoryServiceLevelObjectiveResponseError struct {
	Value        HistoryServiceLevelObjectiveResponseError
	ExplicitNull bool
}

func (v NullableHistoryServiceLevelObjectiveResponseError) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableHistoryServiceLevelObjectiveResponseError) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
