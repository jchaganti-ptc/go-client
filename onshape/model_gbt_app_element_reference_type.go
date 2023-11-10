/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25652-944cf1af37c9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTAppElementReferenceType the model 'GBTAppElementReferenceType'
type GBTAppElementReferenceType string

// List of GBTAppElementReferenceType
const (
	GBTAppElementReferenceTypeUnknown       GBTAppElementReferenceType = "UNKNOWN"
	GBTAppElementReferenceTypePartstudio    GBTAppElementReferenceType = "PARTSTUDIO"
	GBTAppElementReferenceTypeAssembly      GBTAppElementReferenceType = "ASSEMBLY"
	GBTAppElementReferenceTypePart          GBTAppElementReferenceType = "PART"
	GBTAppElementReferenceTypeFlattenedPart GBTAppElementReferenceType = "FLATTENED_PART"
	GBTAppElementReferenceTypeCompositePart GBTAppElementReferenceType = "COMPOSITE_PART"
	GBTAppElementReferenceTypeMeshPart      GBTAppElementReferenceType = "MESH_PART"
	GBTAppElementReferenceTypeSurface       GBTAppElementReferenceType = "SURFACE"
	GBTAppElementReferenceTypeSketch        GBTAppElementReferenceType = "SKETCH"
	GBTAppElementReferenceTypeCurve         GBTAppElementReferenceType = "CURVE"
)

// All allowed values of GBTAppElementReferenceType enum
var AllowedGBTAppElementReferenceTypeEnumValues = []GBTAppElementReferenceType{
	"UNKNOWN",
	"PARTSTUDIO",
	"ASSEMBLY",
	"PART",
	"FLATTENED_PART",
	"COMPOSITE_PART",
	"MESH_PART",
	"SURFACE",
	"SKETCH",
	"CURVE",
}

func (v *GBTAppElementReferenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTAppElementReferenceType(value)
	for _, existing := range AllowedGBTAppElementReferenceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTAppElementReferenceType", value)
}

// NewGBTAppElementReferenceTypeFromValue returns a pointer to a valid GBTAppElementReferenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTAppElementReferenceTypeFromValue(v string) (*GBTAppElementReferenceType, error) {
	ev := GBTAppElementReferenceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTAppElementReferenceType: valid values are %v", v, AllowedGBTAppElementReferenceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTAppElementReferenceType) IsValid() bool {
	for _, existing := range AllowedGBTAppElementReferenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTAppElementReferenceType value
func (v GBTAppElementReferenceType) Ptr() *GBTAppElementReferenceType {
	return &v
}

type NullableGBTAppElementReferenceType struct {
	value *GBTAppElementReferenceType
	isSet bool
}

func (v NullableGBTAppElementReferenceType) Get() *GBTAppElementReferenceType {
	return v.value
}

func (v *NullableGBTAppElementReferenceType) Set(val *GBTAppElementReferenceType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTAppElementReferenceType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTAppElementReferenceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTAppElementReferenceType(val *GBTAppElementReferenceType) *NullableGBTAppElementReferenceType {
	return &NullableGBTAppElementReferenceType{value: val, isSet: true}
}

func (v NullableGBTAppElementReferenceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTAppElementReferenceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
