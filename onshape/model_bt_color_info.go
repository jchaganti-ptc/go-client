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
)

// BTColorInfo struct for BTColorInfo
type BTColorInfo struct {
	Blue  *int32 `json:"blue,omitempty"`
	Green *int32 `json:"green,omitempty"`
	Red   *int32 `json:"red,omitempty"`
}

// NewBTColorInfo instantiates a new BTColorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTColorInfo() *BTColorInfo {
	this := BTColorInfo{}
	return &this
}

// NewBTColorInfoWithDefaults instantiates a new BTColorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTColorInfoWithDefaults() *BTColorInfo {
	this := BTColorInfo{}
	return &this
}

// GetBlue returns the Blue field value if set, zero value otherwise.
func (o *BTColorInfo) GetBlue() int32 {
	if o == nil || o.Blue == nil {
		var ret int32
		return ret
	}
	return *o.Blue
}

// GetBlueOk returns a tuple with the Blue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTColorInfo) GetBlueOk() (*int32, bool) {
	if o == nil || o.Blue == nil {
		return nil, false
	}
	return o.Blue, true
}

// HasBlue returns a boolean if a field has been set.
func (o *BTColorInfo) HasBlue() bool {
	if o != nil && o.Blue != nil {
		return true
	}

	return false
}

// SetBlue gets a reference to the given int32 and assigns it to the Blue field.
func (o *BTColorInfo) SetBlue(v int32) {
	o.Blue = &v
}

// GetGreen returns the Green field value if set, zero value otherwise.
func (o *BTColorInfo) GetGreen() int32 {
	if o == nil || o.Green == nil {
		var ret int32
		return ret
	}
	return *o.Green
}

// GetGreenOk returns a tuple with the Green field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTColorInfo) GetGreenOk() (*int32, bool) {
	if o == nil || o.Green == nil {
		return nil, false
	}
	return o.Green, true
}

// HasGreen returns a boolean if a field has been set.
func (o *BTColorInfo) HasGreen() bool {
	if o != nil && o.Green != nil {
		return true
	}

	return false
}

// SetGreen gets a reference to the given int32 and assigns it to the Green field.
func (o *BTColorInfo) SetGreen(v int32) {
	o.Green = &v
}

// GetRed returns the Red field value if set, zero value otherwise.
func (o *BTColorInfo) GetRed() int32 {
	if o == nil || o.Red == nil {
		var ret int32
		return ret
	}
	return *o.Red
}

// GetRedOk returns a tuple with the Red field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTColorInfo) GetRedOk() (*int32, bool) {
	if o == nil || o.Red == nil {
		return nil, false
	}
	return o.Red, true
}

// HasRed returns a boolean if a field has been set.
func (o *BTColorInfo) HasRed() bool {
	if o != nil && o.Red != nil {
		return true
	}

	return false
}

// SetRed gets a reference to the given int32 and assigns it to the Red field.
func (o *BTColorInfo) SetRed(v int32) {
	o.Red = &v
}

func (o BTColorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Blue != nil {
		toSerialize["blue"] = o.Blue
	}
	if o.Green != nil {
		toSerialize["green"] = o.Green
	}
	if o.Red != nil {
		toSerialize["red"] = o.Red
	}
	return json.Marshal(toSerialize)
}

type NullableBTColorInfo struct {
	value *BTColorInfo
	isSet bool
}

func (v NullableBTColorInfo) Get() *BTColorInfo {
	return v.value
}

func (v *NullableBTColorInfo) Set(val *BTColorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTColorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTColorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTColorInfo(val *BTColorInfo) *NullableBTColorInfo {
	return &NullableBTColorInfo{value: val, isSet: true}
}

func (v NullableBTColorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTColorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
