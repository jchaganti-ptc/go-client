/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.8692-e40d034a55a7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementContentDeltaInfo struct for BTAppElementContentDeltaInfo
type BTAppElementContentDeltaInfo struct {
	Delta *string `json:"delta,omitempty"`
}

// NewBTAppElementContentDeltaInfo instantiates a new BTAppElementContentDeltaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementContentDeltaInfo() *BTAppElementContentDeltaInfo {
	this := BTAppElementContentDeltaInfo{}
	return &this
}

// NewBTAppElementContentDeltaInfoWithDefaults instantiates a new BTAppElementContentDeltaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementContentDeltaInfoWithDefaults() *BTAppElementContentDeltaInfo {
	this := BTAppElementContentDeltaInfo{}
	return &this
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *BTAppElementContentDeltaInfo) GetDelta() string {
	if o == nil || o.Delta == nil {
		var ret string
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentDeltaInfo) GetDeltaOk() (*string, bool) {
	if o == nil || o.Delta == nil {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *BTAppElementContentDeltaInfo) HasDelta() bool {
	if o != nil && o.Delta != nil {
		return true
	}

	return false
}

// SetDelta gets a reference to the given string and assigns it to the Delta field.
func (o *BTAppElementContentDeltaInfo) SetDelta(v string) {
	o.Delta = &v
}

func (o BTAppElementContentDeltaInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Delta != nil {
		toSerialize["delta"] = o.Delta
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementContentDeltaInfo struct {
	value *BTAppElementContentDeltaInfo
	isSet bool
}

func (v NullableBTAppElementContentDeltaInfo) Get() *BTAppElementContentDeltaInfo {
	return v.value
}

func (v *NullableBTAppElementContentDeltaInfo) Set(val *BTAppElementContentDeltaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementContentDeltaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementContentDeltaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementContentDeltaInfo(val *BTAppElementContentDeltaInfo) *NullableBTAppElementContentDeltaInfo {
	return &NullableBTAppElementContentDeltaInfo{value: val, isSet: true}
}

func (v NullableBTAppElementContentDeltaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementContentDeltaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
