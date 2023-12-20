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

// ElementType the model 'ElementType'
type ElementType string

// List of ElementType
const (
	ElementTypeScalar ElementType = "SCALAR"
	ElementTypeVec2   ElementType = "VEC2"
	ElementTypeVec3   ElementType = "VEC3"
	ElementTypeVec4   ElementType = "VEC4"
	ElementTypeMat2   ElementType = "MAT2"
	ElementTypeMat3   ElementType = "MAT3"
	ElementTypeMat4   ElementType = "MAT4"
)

// All allowed values of ElementType enum
var AllowedElementTypeEnumValues = []ElementType{
	"SCALAR",
	"VEC2",
	"VEC3",
	"VEC4",
	"MAT2",
	"MAT3",
	"MAT4",
}

func (v *ElementType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ElementType(value)
	for _, existing := range AllowedElementTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ElementType", value)
}

// NewElementTypeFromValue returns a pointer to a valid ElementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewElementTypeFromValue(v string) (*ElementType, error) {
	ev := ElementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ElementType: valid values are %v", v, AllowedElementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ElementType) IsValid() bool {
	for _, existing := range AllowedElementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElementType value
func (v ElementType) Ptr() *ElementType {
	return &v
}

type NullableElementType struct {
	value *ElementType
	isSet bool
}

func (v NullableElementType) Get() *ElementType {
	return v.value
}

func (v *NullableElementType) Set(val *ElementType) {
	v.value = val
	v.isSet = true
}

func (v NullableElementType) IsSet() bool {
	return v.isSet
}

func (v *NullableElementType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableElementType(val *ElementType) *NullableElementType {
	return &NullableElementType{value: val, isSet: true}
}

func (v NullableElementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableElementType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
