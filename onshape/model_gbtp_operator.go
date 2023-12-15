/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27678-64d64396ca66
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
