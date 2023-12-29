/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28126-700645382199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTBillOfMaterialsExpansionStatus the model 'GBTBillOfMaterialsExpansionStatus'
type GBTBillOfMaterialsExpansionStatus string

// List of GBTBillOfMaterialsExpansionStatus
const (
	GBTBillOfMaterialsExpansionStatusNotExpandable GBTBillOfMaterialsExpansionStatus = "NOT_EXPANDABLE"
	GBTBillOfMaterialsExpansionStatusExpanded      GBTBillOfMaterialsExpansionStatus = "EXPANDED"
	GBTBillOfMaterialsExpansionStatusCollapsed     GBTBillOfMaterialsExpansionStatus = "COLLAPSED"
	GBTBillOfMaterialsExpansionStatusUnknown       GBTBillOfMaterialsExpansionStatus = "UNKNOWN"
)

// All allowed values of GBTBillOfMaterialsExpansionStatus enum
var AllowedGBTBillOfMaterialsExpansionStatusEnumValues = []GBTBillOfMaterialsExpansionStatus{
	"NOT_EXPANDABLE",
	"EXPANDED",
	"COLLAPSED",
	"UNKNOWN",
}

func (v *GBTBillOfMaterialsExpansionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTBillOfMaterialsExpansionStatus(value)
	for _, existing := range AllowedGBTBillOfMaterialsExpansionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTBillOfMaterialsExpansionStatus", value)
}

// NewGBTBillOfMaterialsExpansionStatusFromValue returns a pointer to a valid GBTBillOfMaterialsExpansionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTBillOfMaterialsExpansionStatusFromValue(v string) (*GBTBillOfMaterialsExpansionStatus, error) {
	ev := GBTBillOfMaterialsExpansionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTBillOfMaterialsExpansionStatus: valid values are %v", v, AllowedGBTBillOfMaterialsExpansionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTBillOfMaterialsExpansionStatus) IsValid() bool {
	for _, existing := range AllowedGBTBillOfMaterialsExpansionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTBillOfMaterialsExpansionStatus value
func (v GBTBillOfMaterialsExpansionStatus) Ptr() *GBTBillOfMaterialsExpansionStatus {
	return &v
}

type NullableGBTBillOfMaterialsExpansionStatus struct {
	value *GBTBillOfMaterialsExpansionStatus
	isSet bool
}

func (v NullableGBTBillOfMaterialsExpansionStatus) Get() *GBTBillOfMaterialsExpansionStatus {
	return v.value
}

func (v *NullableGBTBillOfMaterialsExpansionStatus) Set(val *GBTBillOfMaterialsExpansionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTBillOfMaterialsExpansionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTBillOfMaterialsExpansionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTBillOfMaterialsExpansionStatus(val *GBTBillOfMaterialsExpansionStatus) *NullableGBTBillOfMaterialsExpansionStatus {
	return &NullableGBTBillOfMaterialsExpansionStatus{value: val, isSet: true}
}

func (v NullableGBTBillOfMaterialsExpansionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTBillOfMaterialsExpansionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
