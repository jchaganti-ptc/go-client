/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17983-3c4f6e999748
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTPType the model 'GBTPType'
type GBTPType string

// List of GBTPType
const (
	GBTPTypeUndefined GBTPType = "UNDEFINED"
	GBTPTypeBoolean   GBTPType = "BOOLEAN"
	GBTPTypeNumber    GBTPType = "NUMBER"
	GBTPTypeString    GBTPType = "STRING"
	GBTPTypeArray     GBTPType = "ARRAY"
	GBTPTypeMap       GBTPType = "MAP"
	GBTPTypeBox       GBTPType = "BOX"
	GBTPTypeBuiltin   GBTPType = "BUILTIN"
	GBTPTypeFunction  GBTPType = "FUNCTION"
	GBTPTypeUnknown   GBTPType = "UNKNOWN"
)

// All allowed values of GBTPType enum
var AllowedGBTPTypeEnumValues = []GBTPType{
	"UNDEFINED",
	"BOOLEAN",
	"NUMBER",
	"STRING",
	"ARRAY",
	"MAP",
	"BOX",
	"BUILTIN",
	"FUNCTION",
	"UNKNOWN",
}

func (v *GBTPType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTPType(value)
	for _, existing := range AllowedGBTPTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTPType", value)
}

// NewGBTPTypeFromValue returns a pointer to a valid GBTPType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTPTypeFromValue(v string) (*GBTPType, error) {
	ev := GBTPType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTPType: valid values are %v", v, AllowedGBTPTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTPType) IsValid() bool {
	for _, existing := range AllowedGBTPTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTPType value
func (v GBTPType) Ptr() *GBTPType {
	return &v
}

type NullableGBTPType struct {
	value *GBTPType
	isSet bool
}

func (v NullableGBTPType) Get() *GBTPType {
	return v.value
}

func (v *NullableGBTPType) Set(val *GBTPType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTPType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTPType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTPType(val *GBTPType) *NullableGBTPType {
	return &NullableGBTPType{value: val, isSet: true}
}

func (v NullableGBTPType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTPType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
