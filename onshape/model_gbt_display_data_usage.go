/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28658-06d4d4923fc7
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
