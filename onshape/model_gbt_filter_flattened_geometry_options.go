/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27881-5dbbe8053cdf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTFilterFlattenedGeometryOptions the model 'GBTFilterFlattenedGeometryOptions'
type GBTFilterFlattenedGeometryOptions string

// List of GBTFilterFlattenedGeometryOptions
const (
	GBTFilterFlattenedGeometryOptionsModelOnly         GBTFilterFlattenedGeometryOptions = "MODEL_ONLY"
	GBTFilterFlattenedGeometryOptionsFlattenedOnly     GBTFilterFlattenedGeometryOptions = "FLATTENED_ONLY"
	GBTFilterFlattenedGeometryOptionsModelAndFlattened GBTFilterFlattenedGeometryOptions = "MODEL_AND_FLATTENED"
	GBTFilterFlattenedGeometryOptionsUnknown           GBTFilterFlattenedGeometryOptions = "UNKNOWN"
)

// All allowed values of GBTFilterFlattenedGeometryOptions enum
var AllowedGBTFilterFlattenedGeometryOptionsEnumValues = []GBTFilterFlattenedGeometryOptions{
	"MODEL_ONLY",
	"FLATTENED_ONLY",
	"MODEL_AND_FLATTENED",
	"UNKNOWN",
}

func (v *GBTFilterFlattenedGeometryOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTFilterFlattenedGeometryOptions(value)
	for _, existing := range AllowedGBTFilterFlattenedGeometryOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTFilterFlattenedGeometryOptions", value)
}

// NewGBTFilterFlattenedGeometryOptionsFromValue returns a pointer to a valid GBTFilterFlattenedGeometryOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTFilterFlattenedGeometryOptionsFromValue(v string) (*GBTFilterFlattenedGeometryOptions, error) {
	ev := GBTFilterFlattenedGeometryOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTFilterFlattenedGeometryOptions: valid values are %v", v, AllowedGBTFilterFlattenedGeometryOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTFilterFlattenedGeometryOptions) IsValid() bool {
	for _, existing := range AllowedGBTFilterFlattenedGeometryOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTFilterFlattenedGeometryOptions value
func (v GBTFilterFlattenedGeometryOptions) Ptr() *GBTFilterFlattenedGeometryOptions {
	return &v
}

type NullableGBTFilterFlattenedGeometryOptions struct {
	value *GBTFilterFlattenedGeometryOptions
	isSet bool
}

func (v NullableGBTFilterFlattenedGeometryOptions) Get() *GBTFilterFlattenedGeometryOptions {
	return v.value
}

func (v *NullableGBTFilterFlattenedGeometryOptions) Set(val *GBTFilterFlattenedGeometryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTFilterFlattenedGeometryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTFilterFlattenedGeometryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTFilterFlattenedGeometryOptions(val *GBTFilterFlattenedGeometryOptions) *NullableGBTFilterFlattenedGeometryOptions {
	return &NullableGBTFilterFlattenedGeometryOptions{value: val, isSet: true}
}

func (v NullableGBTFilterFlattenedGeometryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTFilterFlattenedGeometryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
