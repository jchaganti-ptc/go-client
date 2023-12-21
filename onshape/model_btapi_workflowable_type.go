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

// BTAPIWorkflowableType All workflowable types that can be enumerated.
type BTAPIWorkflowableType string

// List of BTAPIWorkflowableType
const (
	BTAPIWorkflowableTypeRelease    BTAPIWorkflowableType = "RELEASE"
	BTAPIWorkflowableTypeTask       BTAPIWorkflowableType = "TASK"
	BTAPIWorkflowableTypeAssignment BTAPIWorkflowableType = "ASSIGNMENT"
	BTAPIWorkflowableTypeObsoletion BTAPIWorkflowableType = "OBSOLETION"
)

// All allowed values of BTAPIWorkflowableType enum
var AllowedBTAPIWorkflowableTypeEnumValues = []BTAPIWorkflowableType{
	"RELEASE",
	"TASK",
	"ASSIGNMENT",
	"OBSOLETION",
}

func (v *BTAPIWorkflowableType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTAPIWorkflowableType(value)
	for _, existing := range AllowedBTAPIWorkflowableTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTAPIWorkflowableType", value)
}

// NewBTAPIWorkflowableTypeFromValue returns a pointer to a valid BTAPIWorkflowableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTAPIWorkflowableTypeFromValue(v string) (*BTAPIWorkflowableType, error) {
	ev := BTAPIWorkflowableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTAPIWorkflowableType: valid values are %v", v, AllowedBTAPIWorkflowableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTAPIWorkflowableType) IsValid() bool {
	for _, existing := range AllowedBTAPIWorkflowableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTAPIWorkflowableType value
func (v BTAPIWorkflowableType) Ptr() *BTAPIWorkflowableType {
	return &v
}

type NullableBTAPIWorkflowableType struct {
	value *BTAPIWorkflowableType
	isSet bool
}

func (v NullableBTAPIWorkflowableType) Get() *BTAPIWorkflowableType {
	return v.value
}

func (v *NullableBTAPIWorkflowableType) Set(val *BTAPIWorkflowableType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAPIWorkflowableType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAPIWorkflowableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAPIWorkflowableType(val *BTAPIWorkflowableType) *NullableBTAPIWorkflowableType {
	return &NullableBTAPIWorkflowableType{value: val, isSet: true}
}

func (v NullableBTAPIWorkflowableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAPIWorkflowableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
