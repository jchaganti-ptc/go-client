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

// GBTSurfaceType the model 'GBTSurfaceType'
type GBTSurfaceType string

// List of GBTSurfaceType
const (
	GBTSurfaceTypePlane    GBTSurfaceType = "PLANE"
	GBTSurfaceTypeCylinder GBTSurfaceType = "CYLINDER"
	GBTSurfaceTypeCone     GBTSurfaceType = "CONE"
	GBTSurfaceTypeSphere   GBTSurfaceType = "SPHERE"
	GBTSurfaceTypeTorus    GBTSurfaceType = "TORUS"
	GBTSurfaceTypeOther    GBTSurfaceType = "OTHER"
	GBTSurfaceTypeRevolved GBTSurfaceType = "REVOLVED"
	GBTSurfaceTypeExtruded GBTSurfaceType = "EXTRUDED"
	GBTSurfaceTypeMesh     GBTSurfaceType = "MESH"
	GBTSurfaceTypeSpline   GBTSurfaceType = "SPLINE"
	GBTSurfaceTypeUnknown  GBTSurfaceType = "UNKNOWN"
)

// All allowed values of GBTSurfaceType enum
var AllowedGBTSurfaceTypeEnumValues = []GBTSurfaceType{
	"PLANE",
	"CYLINDER",
	"CONE",
	"SPHERE",
	"TORUS",
	"OTHER",
	"REVOLVED",
	"EXTRUDED",
	"MESH",
	"SPLINE",
	"UNKNOWN",
}

func (v *GBTSurfaceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTSurfaceType(value)
	for _, existing := range AllowedGBTSurfaceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTSurfaceType", value)
}

// NewGBTSurfaceTypeFromValue returns a pointer to a valid GBTSurfaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTSurfaceTypeFromValue(v string) (*GBTSurfaceType, error) {
	ev := GBTSurfaceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTSurfaceType: valid values are %v", v, AllowedGBTSurfaceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTSurfaceType) IsValid() bool {
	for _, existing := range AllowedGBTSurfaceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTSurfaceType value
func (v GBTSurfaceType) Ptr() *GBTSurfaceType {
	return &v
}

type NullableGBTSurfaceType struct {
	value *GBTSurfaceType
	isSet bool
}

func (v NullableGBTSurfaceType) Get() *GBTSurfaceType {
	return v.value
}

func (v *NullableGBTSurfaceType) Set(val *GBTSurfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTSurfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTSurfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTSurfaceType(val *GBTSurfaceType) *NullableGBTSurfaceType {
	return &NullableGBTSurfaceType{value: val, isSet: true}
}

func (v NullableGBTSurfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTSurfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
