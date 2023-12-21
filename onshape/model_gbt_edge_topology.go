/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27960-e195de6ac56c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTEdgeTopology the model 'GBTEdgeTopology'
type GBTEdgeTopology string

// List of GBTEdgeTopology
const (
	GBTEdgeTopologyWire     GBTEdgeTopology = "WIRE"
	GBTEdgeTopologyOneSided GBTEdgeTopology = "ONE_SIDED"
	GBTEdgeTopologyTwoSided GBTEdgeTopology = "TWO_SIDED"
	GBTEdgeTopologyLaminar  GBTEdgeTopology = "LAMINAR"
	GBTEdgeTopologyUnknown  GBTEdgeTopology = "UNKNOWN"
)

// All allowed values of GBTEdgeTopology enum
var AllowedGBTEdgeTopologyEnumValues = []GBTEdgeTopology{
	"WIRE",
	"ONE_SIDED",
	"TWO_SIDED",
	"LAMINAR",
	"UNKNOWN",
}

func (v *GBTEdgeTopology) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTEdgeTopology(value)
	for _, existing := range AllowedGBTEdgeTopologyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTEdgeTopology", value)
}

// NewGBTEdgeTopologyFromValue returns a pointer to a valid GBTEdgeTopology
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTEdgeTopologyFromValue(v string) (*GBTEdgeTopology, error) {
	ev := GBTEdgeTopology(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTEdgeTopology: valid values are %v", v, AllowedGBTEdgeTopologyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTEdgeTopology) IsValid() bool {
	for _, existing := range AllowedGBTEdgeTopologyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTEdgeTopology value
func (v GBTEdgeTopology) Ptr() *GBTEdgeTopology {
	return &v
}

type NullableGBTEdgeTopology struct {
	value *GBTEdgeTopology
	isSet bool
}

func (v NullableGBTEdgeTopology) Get() *GBTEdgeTopology {
	return v.value
}

func (v *NullableGBTEdgeTopology) Set(val *GBTEdgeTopology) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTEdgeTopology) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTEdgeTopology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTEdgeTopology(val *GBTEdgeTopology) *NullableGBTEdgeTopology {
	return &NullableGBTEdgeTopology{value: val, isSet: true}
}

func (v NullableGBTEdgeTopology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTEdgeTopology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
