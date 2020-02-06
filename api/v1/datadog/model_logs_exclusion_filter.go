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

// LogsExclusionFilter Exclusion filter is defined by a query, a sampling rule, and a active/inactive toggle.
type LogsExclusionFilter struct {
	// Default query is '*', meaning all logs flowing in the index would be excluded. Scope down exclusion filter to only a subset of logs with a log query.
	Query      *string `json:"query,omitempty"`
	SampleRate float64 `json:"sample_rate"`
}

// NewLogsExclusionFilter instantiates a new LogsExclusionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsExclusionFilter(sampleRate float64) *LogsExclusionFilter {
	this := LogsExclusionFilter{}
	this.SampleRate = sampleRate
	return &this
}

// NewLogsExclusionFilterWithDefaults instantiates a new LogsExclusionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsExclusionFilterWithDefaults() *LogsExclusionFilter {
	this := LogsExclusionFilter{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LogsExclusionFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogsExclusionFilter) GetQueryOk() (string, bool) {
	if o == nil || o.Query == nil {
		var ret string
		return ret, false
	}
	return *o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LogsExclusionFilter) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LogsExclusionFilter) SetQuery(v string) {
	o.Query = &v
}

// GetSampleRate returns the SampleRate field value
func (o *LogsExclusionFilter) GetSampleRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SampleRate
}

// SetSampleRate sets field value
func (o *LogsExclusionFilter) SetSampleRate(v float64) {
	o.SampleRate = v
}

type NullableLogsExclusionFilter struct {
	Value        LogsExclusionFilter
	ExplicitNull bool
}

func (v NullableLogsExclusionFilter) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLogsExclusionFilter) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}