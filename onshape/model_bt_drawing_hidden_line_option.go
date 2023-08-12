/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.20454-7718daa9749d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTDrawingHiddenLineOption the model 'BTDrawingHiddenLineOption'
type BTDrawingHiddenLineOption string

// List of BTDrawingHiddenLineOption
const (
	BTDrawingHiddenLineOptionDrafting BTDrawingHiddenLineOption = "drafting"
	BTDrawingHiddenLineOptionExcluded BTDrawingHiddenLineOption = "excluded"
	BTDrawingHiddenLineOptionMarked   BTDrawingHiddenLineOption = "marked"
)

// All allowed values of BTDrawingHiddenLineOption enum
var AllowedBTDrawingHiddenLineOptionEnumValues = []BTDrawingHiddenLineOption{
	"drafting",
	"excluded",
	"marked",
}

func (v *BTDrawingHiddenLineOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTDrawingHiddenLineOption(value)
	for _, existing := range AllowedBTDrawingHiddenLineOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTDrawingHiddenLineOption", value)
}

// NewBTDrawingHiddenLineOptionFromValue returns a pointer to a valid BTDrawingHiddenLineOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTDrawingHiddenLineOptionFromValue(v string) (*BTDrawingHiddenLineOption, error) {
	ev := BTDrawingHiddenLineOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTDrawingHiddenLineOption: valid values are %v", v, AllowedBTDrawingHiddenLineOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTDrawingHiddenLineOption) IsValid() bool {
	for _, existing := range AllowedBTDrawingHiddenLineOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTDrawingHiddenLineOption value
func (v BTDrawingHiddenLineOption) Ptr() *BTDrawingHiddenLineOption {
	return &v
}

type NullableBTDrawingHiddenLineOption struct {
	value *BTDrawingHiddenLineOption
	isSet bool
}

func (v NullableBTDrawingHiddenLineOption) Get() *BTDrawingHiddenLineOption {
	return v.value
}

func (v *NullableBTDrawingHiddenLineOption) Set(val *BTDrawingHiddenLineOption) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDrawingHiddenLineOption) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDrawingHiddenLineOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDrawingHiddenLineOption(val *BTDrawingHiddenLineOption) *NullableBTDrawingHiddenLineOption {
	return &NullableBTDrawingHiddenLineOption{value: val, isSet: true}
}

func (v NullableBTDrawingHiddenLineOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDrawingHiddenLineOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
