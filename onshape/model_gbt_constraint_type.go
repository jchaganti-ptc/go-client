/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19458-7ff87863110f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTConstraintType the model 'GBTConstraintType'
type GBTConstraintType string

// List of GBTConstraintType
const (
	GBTConstraintTypeNone                GBTConstraintType = "NONE"
	GBTConstraintTypeCoincident          GBTConstraintType = "COINCIDENT"
	GBTConstraintTypeParallel            GBTConstraintType = "PARALLEL"
	GBTConstraintTypeVertical            GBTConstraintType = "VERTICAL"
	GBTConstraintTypeHorizontal          GBTConstraintType = "HORIZONTAL"
	GBTConstraintTypePerpendicular       GBTConstraintType = "PERPENDICULAR"
	GBTConstraintTypeConcentric          GBTConstraintType = "CONCENTRIC"
	GBTConstraintTypeMirror              GBTConstraintType = "MIRROR"
	GBTConstraintTypeMidpoint            GBTConstraintType = "MIDPOINT"
	GBTConstraintTypeTangent             GBTConstraintType = "TANGENT"
	GBTConstraintTypeEqual               GBTConstraintType = "EQUAL"
	GBTConstraintTypeLength              GBTConstraintType = "LENGTH"
	GBTConstraintTypeDistance            GBTConstraintType = "DISTANCE"
	GBTConstraintTypeAngle               GBTConstraintType = "ANGLE"
	GBTConstraintTypeRadius              GBTConstraintType = "RADIUS"
	GBTConstraintTypeNormal              GBTConstraintType = "NORMAL"
	GBTConstraintTypeFix                 GBTConstraintType = "FIX"
	GBTConstraintTypeProjected           GBTConstraintType = "PROJECTED"
	GBTConstraintTypeOffset              GBTConstraintType = "OFFSET"
	GBTConstraintTypeCircularPattern     GBTConstraintType = "CIRCULAR_PATTERN"
	GBTConstraintTypePierce              GBTConstraintType = "PIERCE"
	GBTConstraintTypeLinearPattern       GBTConstraintType = "LINEAR_PATTERN"
	GBTConstraintTypeMajorDiameter       GBTConstraintType = "MAJOR_DIAMETER"
	GBTConstraintTypeMinorDiameter       GBTConstraintType = "MINOR_DIAMETER"
	GBTConstraintTypeQuadrant            GBTConstraintType = "QUADRANT"
	GBTConstraintTypeDiameter            GBTConstraintType = "DIAMETER"
	GBTConstraintTypeSilhouetted         GBTConstraintType = "SILHOUETTED"
	GBTConstraintTypeCenterlineDimension GBTConstraintType = "CENTERLINE_DIMENSION"
	GBTConstraintTypeIntersected         GBTConstraintType = "INTERSECTED"
	GBTConstraintTypeRho                 GBTConstraintType = "RHO"
	GBTConstraintTypeEqualCurvature      GBTConstraintType = "EQUAL_CURVATURE"
	GBTConstraintTypeUnknown             GBTConstraintType = "UNKNOWN"
)

// All allowed values of GBTConstraintType enum
var AllowedGBTConstraintTypeEnumValues = []GBTConstraintType{
	"NONE",
	"COINCIDENT",
	"PARALLEL",
	"VERTICAL",
	"HORIZONTAL",
	"PERPENDICULAR",
	"CONCENTRIC",
	"MIRROR",
	"MIDPOINT",
	"TANGENT",
	"EQUAL",
	"LENGTH",
	"DISTANCE",
	"ANGLE",
	"RADIUS",
	"NORMAL",
	"FIX",
	"PROJECTED",
	"OFFSET",
	"CIRCULAR_PATTERN",
	"PIERCE",
	"LINEAR_PATTERN",
	"MAJOR_DIAMETER",
	"MINOR_DIAMETER",
	"QUADRANT",
	"DIAMETER",
	"SILHOUETTED",
	"CENTERLINE_DIMENSION",
	"INTERSECTED",
	"RHO",
	"EQUAL_CURVATURE",
	"UNKNOWN",
}

func (v *GBTConstraintType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTConstraintType(value)
	for _, existing := range AllowedGBTConstraintTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTConstraintType", value)
}

// NewGBTConstraintTypeFromValue returns a pointer to a valid GBTConstraintType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTConstraintTypeFromValue(v string) (*GBTConstraintType, error) {
	ev := GBTConstraintType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTConstraintType: valid values are %v", v, AllowedGBTConstraintTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTConstraintType) IsValid() bool {
	for _, existing := range AllowedGBTConstraintTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTConstraintType value
func (v GBTConstraintType) Ptr() *GBTConstraintType {
	return &v
}

type NullableGBTConstraintType struct {
	value *GBTConstraintType
	isSet bool
}

func (v NullableGBTConstraintType) Get() *GBTConstraintType {
	return v.value
}

func (v *NullableGBTConstraintType) Set(val *GBTConstraintType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTConstraintType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTConstraintType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTConstraintType(val *GBTConstraintType) *NullableGBTConstraintType {
	return &NullableGBTConstraintType{value: val, isSet: true}
}

func (v NullableGBTConstraintType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTConstraintType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
