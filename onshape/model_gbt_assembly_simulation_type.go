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

// GBTAssemblySimulationType the model 'GBTAssemblySimulationType'
type GBTAssemblySimulationType string

// List of GBTAssemblySimulationType
const (
	GBTAssemblySimulationTypeLinearStatic    GBTAssemblySimulationType = "LINEAR_STATIC"
	GBTAssemblySimulationTypeModal           GBTAssemblySimulationType = "MODAL"
	GBTAssemblySimulationTypeContactAnalysis GBTAssemblySimulationType = "CONTACT_ANALYSIS"
	GBTAssemblySimulationTypeUnknown         GBTAssemblySimulationType = "UNKNOWN"
)

// All allowed values of GBTAssemblySimulationType enum
var AllowedGBTAssemblySimulationTypeEnumValues = []GBTAssemblySimulationType{
	"LINEAR_STATIC",
	"MODAL",
	"CONTACT_ANALYSIS",
	"UNKNOWN",
}

func (v *GBTAssemblySimulationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTAssemblySimulationType(value)
	for _, existing := range AllowedGBTAssemblySimulationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTAssemblySimulationType", value)
}

// NewGBTAssemblySimulationTypeFromValue returns a pointer to a valid GBTAssemblySimulationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTAssemblySimulationTypeFromValue(v string) (*GBTAssemblySimulationType, error) {
	ev := GBTAssemblySimulationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTAssemblySimulationType: valid values are %v", v, AllowedGBTAssemblySimulationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTAssemblySimulationType) IsValid() bool {
	for _, existing := range AllowedGBTAssemblySimulationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTAssemblySimulationType value
func (v GBTAssemblySimulationType) Ptr() *GBTAssemblySimulationType {
	return &v
}

type NullableGBTAssemblySimulationType struct {
	value *GBTAssemblySimulationType
	isSet bool
}

func (v NullableGBTAssemblySimulationType) Get() *GBTAssemblySimulationType {
	return v.value
}

func (v *NullableGBTAssemblySimulationType) Set(val *GBTAssemblySimulationType) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTAssemblySimulationType) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTAssemblySimulationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTAssemblySimulationType(val *GBTAssemblySimulationType) *NullableGBTAssemblySimulationType {
	return &NullableGBTAssemblySimulationType{value: val, isSet: true}
}

func (v NullableGBTAssemblySimulationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTAssemblySimulationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
