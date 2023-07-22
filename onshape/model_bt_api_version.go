/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19303-3cbf47a47fe4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTApiVersion the model 'BTApiVersion'
type BTApiVersion string

// List of BTApiVersion
const (
	BTApiVersionUndefined                       BTApiVersion = "UNDEFINED"
	BTApiVersionV1Start                         BTApiVersion = "V1_START"
	BTApiVersionV2SerializationUnification      BTApiVersion = "V2_SERIALIZATION_UNIFICATION"
	BTApiVersionV3NewBomFormat                  BTApiVersion = "V3_NEW_BOM_FORMAT"
	BTApiVersionV4TransactionsNoNew             BTApiVersion = "V4_TRANSACTIONS_NO_NEW"
	BTApiVersionV5BodyDetailsCompositeReference BTApiVersion = "V5_BODY_DETAILS_COMPOSITE_REFERENCE"
	BTApiVersionV6JsonEditResponseBugfix        BTApiVersion = "V6_JSON_EDIT_RESPONSE_BUGFIX"
)

// All allowed values of BTApiVersion enum
var AllowedBTApiVersionEnumValues = []BTApiVersion{
	"UNDEFINED",
	"V1_START",
	"V2_SERIALIZATION_UNIFICATION",
	"V3_NEW_BOM_FORMAT",
	"V4_TRANSACTIONS_NO_NEW",
	"V5_BODY_DETAILS_COMPOSITE_REFERENCE",
	"V6_JSON_EDIT_RESPONSE_BUGFIX",
}

func (v *BTApiVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTApiVersion(value)
	for _, existing := range AllowedBTApiVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTApiVersion", value)
}

// NewBTApiVersionFromValue returns a pointer to a valid BTApiVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTApiVersionFromValue(v string) (*BTApiVersion, error) {
	ev := BTApiVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTApiVersion: valid values are %v", v, AllowedBTApiVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTApiVersion) IsValid() bool {
	for _, existing := range AllowedBTApiVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTApiVersion value
func (v BTApiVersion) Ptr() *BTApiVersion {
	return &v
}

type NullableBTApiVersion struct {
	value *BTApiVersion
	isSet bool
}

func (v NullableBTApiVersion) Get() *BTApiVersion {
	return v.value
}

func (v *NullableBTApiVersion) Set(val *BTApiVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApiVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApiVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApiVersion(val *BTApiVersion) *NullableBTApiVersion {
	return &NullableBTApiVersion{value: val, isSet: true}
}

func (v NullableBTApiVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApiVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
