/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9191-43c781405890
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUnitsMaximumDisplayPrecisionInfo struct for BTUnitsMaximumDisplayPrecisionInfo
type BTUnitsMaximumDisplayPrecisionInfo struct {
	UnitsDisplayPrecision *map[string]int32 `json:"unitsDisplayPrecision,omitempty"`
}

// NewBTUnitsMaximumDisplayPrecisionInfo instantiates a new BTUnitsMaximumDisplayPrecisionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUnitsMaximumDisplayPrecisionInfo() *BTUnitsMaximumDisplayPrecisionInfo {
	this := BTUnitsMaximumDisplayPrecisionInfo{}
	return &this
}

// NewBTUnitsMaximumDisplayPrecisionInfoWithDefaults instantiates a new BTUnitsMaximumDisplayPrecisionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUnitsMaximumDisplayPrecisionInfoWithDefaults() *BTUnitsMaximumDisplayPrecisionInfo {
	this := BTUnitsMaximumDisplayPrecisionInfo{}
	return &this
}

// GetUnitsDisplayPrecision returns the UnitsDisplayPrecision field value if set, zero value otherwise.
func (o *BTUnitsMaximumDisplayPrecisionInfo) GetUnitsDisplayPrecision() map[string]int32 {
	if o == nil || o.UnitsDisplayPrecision == nil {
		var ret map[string]int32
		return ret
	}
	return *o.UnitsDisplayPrecision
}

// GetUnitsDisplayPrecisionOk returns a tuple with the UnitsDisplayPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnitsMaximumDisplayPrecisionInfo) GetUnitsDisplayPrecisionOk() (*map[string]int32, bool) {
	if o == nil || o.UnitsDisplayPrecision == nil {
		return nil, false
	}
	return o.UnitsDisplayPrecision, true
}

// HasUnitsDisplayPrecision returns a boolean if a field has been set.
func (o *BTUnitsMaximumDisplayPrecisionInfo) HasUnitsDisplayPrecision() bool {
	if o != nil && o.UnitsDisplayPrecision != nil {
		return true
	}

	return false
}

// SetUnitsDisplayPrecision gets a reference to the given map[string]int32 and assigns it to the UnitsDisplayPrecision field.
func (o *BTUnitsMaximumDisplayPrecisionInfo) SetUnitsDisplayPrecision(v map[string]int32) {
	o.UnitsDisplayPrecision = &v
}

func (o BTUnitsMaximumDisplayPrecisionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnitsDisplayPrecision != nil {
		toSerialize["unitsDisplayPrecision"] = o.UnitsDisplayPrecision
	}
	return json.Marshal(toSerialize)
}

type NullableBTUnitsMaximumDisplayPrecisionInfo struct {
	value *BTUnitsMaximumDisplayPrecisionInfo
	isSet bool
}

func (v NullableBTUnitsMaximumDisplayPrecisionInfo) Get() *BTUnitsMaximumDisplayPrecisionInfo {
	return v.value
}

func (v *NullableBTUnitsMaximumDisplayPrecisionInfo) Set(val *BTUnitsMaximumDisplayPrecisionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUnitsMaximumDisplayPrecisionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUnitsMaximumDisplayPrecisionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUnitsMaximumDisplayPrecisionInfo(val *BTUnitsMaximumDisplayPrecisionInfo) *NullableBTUnitsMaximumDisplayPrecisionInfo {
	return &NullableBTUnitsMaximumDisplayPrecisionInfo{value: val, isSet: true}
}

func (v NullableBTUnitsMaximumDisplayPrecisionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUnitsMaximumDisplayPrecisionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
