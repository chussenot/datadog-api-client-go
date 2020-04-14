/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// IncidentFacetSchemaDataOneOf struct for IncidentFacetSchemaDataOneOf
type IncidentFacetSchemaDataOneOf struct {
	IncidentFacetSchemaDataOneOfInterface interface{ GetType() string }
}

func (s IncidentFacetSchemaDataOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.IncidentFacetSchemaDataOneOfInterface)
}

func (s *IncidentFacetSchemaDataOneOf) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaled map[string]interface{}
	err = json.Unmarshal(src, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["type"]; ok {
		switch v {
		case "aggregation_percentiles":
			var result *IncidentFacetPercentilesAggregationDataSchema = &IncidentFacetPercentilesAggregationDataSchema{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.IncidentFacetSchemaDataOneOfInterface = result
			return nil
		case "aggregation_stats":
			var result *IncidentFacetStatsAggregationDataSchema = &IncidentFacetStatsAggregationDataSchema{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.IncidentFacetSchemaDataOneOfInterface = result
			return nil
		case "aggregation_terms":
			var result *IncidentFacetTermsAggregationDataSchema = &IncidentFacetTermsAggregationDataSchema{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.IncidentFacetSchemaDataOneOfInterface = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'type' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'type' not found in unmarshaled payload: %+v", unmarshaled)
	}
}

type NullableIncidentFacetSchemaDataOneOf struct {
	value *IncidentFacetSchemaDataOneOf
	isSet bool
}

func (v NullableIncidentFacetSchemaDataOneOf) Get() *IncidentFacetSchemaDataOneOf {
	return v.value
}

func (v *NullableIncidentFacetSchemaDataOneOf) Set(val *IncidentFacetSchemaDataOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentFacetSchemaDataOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentFacetSchemaDataOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentFacetSchemaDataOneOf(val *IncidentFacetSchemaDataOneOf) *NullableIncidentFacetSchemaDataOneOf {
	return &NullableIncidentFacetSchemaDataOneOf{value: val, isSet: true}
}

func (v NullableIncidentFacetSchemaDataOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentFacetSchemaDataOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}