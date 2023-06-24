/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17666-197c9d1638c5
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTDisplayDataUsage the model 'GBTDisplayDataUsage'
type GBTDisplayDataUsage string

// List of GBTDisplayDataUsage
const (
	GBTDisplayDataUsageBase          GBTDisplayDataUsage = "BASE"
	GBTDisplayDataUsagePreviewBefore GBTDisplayDataUsage = "PREVIEW_BEFORE"
	GBTDisplayDataUsagePreviewAfter  GBTDisplayDataUsage = "PREVIEW_AFTER"
	GBTDisplayDataUsagePreviewFinal  GBTDisplayDataUsage = "PREVIEW_FINAL"
	GBTDisplayDataUsageCompareTarget GBTDisplayDataUsage = "COMPARE_TARGET"
	GBTDisplayDataUsageUnknown       GBTDisplayDataUsage = "UNKNOWN"
)

// All allowed values of GBTDisplayDataUsage enum
var AllowedGBTDisplayDataUsageEnumValues = []GBTDisplayDataUsage{
	"BASE",
	"PREVIEW_BEFORE",
	"PREVIEW_AFTER",
	"PREVIEW_FINAL",
	"COMPARE_TARGET",
	"UNKNOWN",
}

func (v *GBTDisplayDataUsage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTDisplayDataUsage(value)
	for _, existing := range AllowedGBTDisplayDataUsageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTDisplayDataUsage", value)
}

// NewGBTDisplayDataUsageFromValue returns a pointer to a valid GBTDisplayDataUsage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTDisplayDataUsageFromValue(v string) (*GBTDisplayDataUsage, error) {
	ev := GBTDisplayDataUsage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTDisplayDataUsage: valid values are %v", v, AllowedGBTDisplayDataUsageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTDisplayDataUsage) IsValid() bool {
	for _, existing := range AllowedGBTDisplayDataUsageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTDisplayDataUsage value
func (v GBTDisplayDataUsage) Ptr() *GBTDisplayDataUsage {
	return &v
}

type NullableGBTDisplayDataUsage struct {
	value *GBTDisplayDataUsage
	isSet bool
}

func (v NullableGBTDisplayDataUsage) Get() *GBTDisplayDataUsage {
	return v.value
}

func (v *NullableGBTDisplayDataUsage) Set(val *GBTDisplayDataUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTDisplayDataUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTDisplayDataUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTDisplayDataUsage(val *GBTDisplayDataUsage) *NullableGBTDisplayDataUsage {
	return &NullableGBTDisplayDataUsage{value: val, isSet: true}
}

func (v NullableGBTDisplayDataUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTDisplayDataUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}