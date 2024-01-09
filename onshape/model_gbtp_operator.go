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

// GBTPOperator the model 'GBTPOperator'
type GBTPOperator string

// List of GBTPOperator
const (
	GBTPOperatorNone           GBTPOperator = "NONE"
	GBTPOperatorPlus           GBTPOperator = "PLUS"
	GBTPOperatorMinus          GBTPOperator = "MINUS"
	GBTPOperatorTimes          GBTPOperator = "TIMES"
	GBTPOperatorDivide         GBTPOperator = "DIVIDE"
	GBTPOperatorModulus        GBTPOperator = "MODULUS"
	GBTPOperatorPower          GBTPOperator = "POWER"
	GBTPOperatorNegate         GBTPOperator = "NEGATE"
	GBTPOperatorOr             GBTPOperator = "OR"
	GBTPOperatorAnd            GBTPOperator = "AND"
	GBTPOperatorNot            GBTPOperator = "NOT"
	GBTPOperatorEqualTo        GBTPOperator = "EQUAL_TO"
	GBTPOperatorNotEqualTo     GBTPOperator = "NOT_EQUAL_TO"
	GBTPOperatorGreater        GBTPOperator = "GREATER"
	GBTPOperatorLess           GBTPOperator = "LESS"
	GBTPOperatorGreaterOrEqual GBTPOperator = "GREATER_OR_EQUAL"
	GBTPOperatorLessOrEqual    GBTPOperator = "LESS_OR_EQUAL"
	GBTPOperatorConcatenate    GBTPOperator = "CONCATENATE"
	GBTPOperatorConditional    GBTPOperator = "CONDITIONAL"
)

// All allowed values of GBTPOperator enum
var AllowedGBTPOperatorEnumValues = []GBTPOperator{
	"NONE",
	"PLUS",
	"MINUS",
	"TIMES",
	"DIVIDE",
	"MODULUS",
	"POWER",
	"NEGATE",
	"OR",
	"AND",
	"NOT",
	"EQUAL_TO",
	"NOT_EQUAL_TO",
	"GREATER",
	"LESS",
	"GREATER_OR_EQUAL",
	"LESS_OR_EQUAL",
	"CONCATENATE",
	"CONDITIONAL",
}

func (v *GBTPOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTPOperator(value)
	for _, existing := range AllowedGBTPOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTPOperator", value)
}

// NewGBTPOperatorFromValue returns a pointer to a valid GBTPOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTPOperatorFromValue(v string) (*GBTPOperator, error) {
	ev := GBTPOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTPOperator: valid values are %v", v, AllowedGBTPOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTPOperator) IsValid() bool {
	for _, existing := range AllowedGBTPOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTPOperator value
func (v GBTPOperator) Ptr() *GBTPOperator {
	return &v
}

type NullableGBTPOperator struct {
	value *GBTPOperator
	isSet bool
}

func (v NullableGBTPOperator) Get() *GBTPOperator {
	return v.value
}

func (v *NullableGBTPOperator) Set(val *GBTPOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTPOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTPOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTPOperator(val *GBTPOperator) *NullableGBTPOperator {
	return &NullableGBTPOperator{value: val, isSet: true}
}

func (v NullableGBTPOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTPOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
