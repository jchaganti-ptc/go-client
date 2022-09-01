/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6309-06d9e62c38f0
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMicroversionInfo struct for BTMicroversionInfo
type BTMicroversionInfo struct {
	Microversion *string `json:"microversion,omitempty"`
}

// NewBTMicroversionInfo instantiates a new BTMicroversionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMicroversionInfo() *BTMicroversionInfo {
	this := BTMicroversionInfo{}
	return &this
}

// NewBTMicroversionInfoWithDefaults instantiates a new BTMicroversionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMicroversionInfoWithDefaults() *BTMicroversionInfo {
	this := BTMicroversionInfo{}
	return &this
}

// GetMicroversion returns the Microversion field value if set, zero value otherwise.
func (o *BTMicroversionInfo) GetMicroversion() string {
	if o == nil || o.Microversion == nil {
		var ret string
		return ret
	}
	return *o.Microversion
}

// GetMicroversionOk returns a tuple with the Microversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionInfo) GetMicroversionOk() (*string, bool) {
	if o == nil || o.Microversion == nil {
		return nil, false
	}
	return o.Microversion, true
}

// HasMicroversion returns a boolean if a field has been set.
func (o *BTMicroversionInfo) HasMicroversion() bool {
	if o != nil && o.Microversion != nil {
		return true
	}

	return false
}

// SetMicroversion gets a reference to the given string and assigns it to the Microversion field.
func (o *BTMicroversionInfo) SetMicroversion(v string) {
	o.Microversion = &v
}

func (o BTMicroversionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Microversion != nil {
		toSerialize["microversion"] = o.Microversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTMicroversionInfo struct {
	value *BTMicroversionInfo
	isSet bool
}

func (v NullableBTMicroversionInfo) Get() *BTMicroversionInfo {
	return v.value
}

func (v *NullableBTMicroversionInfo) Set(val *BTMicroversionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMicroversionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMicroversionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMicroversionInfo(val *BTMicroversionInfo) *NullableBTMicroversionInfo {
	return &NullableBTMicroversionInfo{value: val, isSet: true}
}

func (v NullableBTMicroversionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMicroversionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
