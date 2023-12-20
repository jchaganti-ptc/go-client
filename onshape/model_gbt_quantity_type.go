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

// GBTQuantityType the model 'GBTQuantityType'
type GBTQuantityType string

// List of GBTQuantityType
const (
	GBTQuantityTypeUnknown           GBTQuantityType = "UNKNOWN"
	GBTQuantityTypeInteger           GBTQuantityType = "INTEGER"
	GBTQuantityTypeReal              GBTQuantityType = "REAL"
	GBTQuantityTypeLength            GBTQuantityType = "LENGTH"
	GBTQuantityTypeAngle             GBTQuantityType = "ANGLE"
	GBTQuantityTypeMass              GBTQuantityType = "MASS"
	GBTQuantityTypeTime              GBTQuantityType = "TIME"
	GBTQuantityTypeTemperature       GBTQuantityType = "TEMPERATURE"
	GBTQuantityTypeCurrent           GBTQuantityType = "CURRENT"
	GBTQuantityTypeAnything          GBTQuantityType = "ANYTHING"
	GBTQuantityTypeAnythingWithUnits GBTQuantityType = "ANYTHING_WITH_UNITS"
	GBTQuantityTypeForce             GBTQuantityType = "FORCE"
	GBTQuantityTypePressure          GBTQuantityType = "PRESSURE"
	GBTQuantityTypeMoment            GBTQuantityType = "MOMENT"
	GBTQuantityTypeAcceleration      GBTQuantityType = "ACCELERATION"
	GBTQuantityTypeAngularVelocity   GBTQuantityType = "ANGULAR_VELOCITY"
	GBTQuantityTypeEnergy            GBTQuantityType = "ENERGY"
	GBTQuantityTypeArea              GBTQuantityType = "AREA"
	GBTQuantityTypeVolume            GBTQuantityType = "VOLUME"
	GBTQuantityTypeBoolean           GBTQuantityType = "BOOLEAN"
	GBTQuantityTypeString            GBTQuantityType = "STRING"
)

// All allowed values of GBTQuantityType enum
var AllowedGBTQuantityTypeEnumValues = []GBTQuantityType{
	"UNKNOWN",
	"INTEGER",
	"REAL",
	"LENGTH",
	"ANGLE",
	"MASS",
	"TIME",
	"TEMPERATURE",
	"CURRENT",
	"ANYTHING",
	"ANYTHING_WITH_UNITS",
	"FORCE",
	"PRESSURE",
	"MOMENT",
	"ACCELERATION",
	"ANGULAR_VELOCITY",
	"ENERGY",
	"AREA",
	"VOLUME",
	"BOOLEAN",
	"STRING",
}

func (v *GBTQuantityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTQuantityType(value)
	for _, existing := range AllowedGBTQuantityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTQuantityType", value)
}

// NewGBTQuantityTypeFromValue returns a pointer to a valid GBTQuantityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTQuantityTypeFromValue(v string) (*GBTQuantityType, error) {
	ev := GBTQuantityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTQuantityType: valid values are %v", v, AllowedGBTQuantityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTQuantityType) IsValid() bool {
	for _, existing := range AllowedGBTQuantityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTQuantityType value
func (v GBTQuantityType) Ptr() *GBTQuantityType {
	return &v
}

type NullableGBTQuantityType struct {
	value *GBTQuantityType
	isSet bool
}

func (v NullableGBTQuantityType) Get() *GBTQuantityType {
	return v.value
}

func (v *NullableGBTQuantityType) Set(val *GBTQuantityType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTQuantityType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTQuantityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTQuantityType(val *GBTQuantityType) *NullableGBTQuantityType {
	return &NullableGBTQuantityType{value: val, isSet: true}
}

func (v NullableGBTQuantityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTQuantityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
