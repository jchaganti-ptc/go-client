/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23604-b59b123004e9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTEntityType the model 'GBTEntityType'
type GBTEntityType string

// List of GBTEntityType
const (
	GBTEntityTypeVertex         GBTEntityType = "VERTEX"
	GBTEntityTypeEdge           GBTEntityType = "EDGE"
	GBTEntityTypeFace           GBTEntityType = "FACE"
	GBTEntityTypeBody           GBTEntityType = "BODY"
	GBTEntityTypeDegenerateEdge GBTEntityType = "DEGENERATE_EDGE"
	GBTEntityTypeUnknown        GBTEntityType = "UNKNOWN"
)

// All allowed values of GBTEntityType enum
var AllowedGBTEntityTypeEnumValues = []GBTEntityType{
	"VERTEX",
	"EDGE",
	"FACE",
	"BODY",
	"DEGENERATE_EDGE",
	"UNKNOWN",
}

func (v *GBTEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTEntityType(value)
	for _, existing := range AllowedGBTEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTEntityType", value)
}

// NewGBTEntityTypeFromValue returns a pointer to a valid GBTEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTEntityTypeFromValue(v string) (*GBTEntityType, error) {
	ev := GBTEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTEntityType: valid values are %v", v, AllowedGBTEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTEntityType) IsValid() bool {
	for _, existing := range AllowedGBTEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTEntityType value
func (v GBTEntityType) Ptr() *GBTEntityType {
	return &v
}

type NullableGBTEntityType struct {
	value *GBTEntityType
	isSet bool
}

func (v NullableGBTEntityType) Get() *GBTEntityType {
	return v.value
}

func (v *NullableGBTEntityType) Set(val *GBTEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTEntityType(val *GBTEntityType) *NullableGBTEntityType {
	return &NullableGBTEntityType{value: val, isSet: true}
}

func (v NullableGBTEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
