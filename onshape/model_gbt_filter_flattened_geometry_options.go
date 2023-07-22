/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19303-3cbf47a47fe4
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
