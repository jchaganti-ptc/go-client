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

// BTImmutableIntArray struct for BTImmutableIntArray
type BTImmutableIntArray struct {
	Empty *bool `json:"empty,omitempty"`
}

// NewBTImmutableIntArray instantiates a new BTImmutableIntArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTImmutableIntArray() *BTImmutableIntArray {
	this := BTImmutableIntArray{}
	return &this
}

// NewBTImmutableIntArrayWithDefaults instantiates a new BTImmutableIntArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTImmutableIntArrayWithDefaults() *BTImmutableIntArray {
	this := BTImmutableIntArray{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *BTImmutableIntArray) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTImmutableIntArray) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *BTImmutableIntArray) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *BTImmutableIntArray) SetEmpty(v bool) {
	o.Empty = &v
}

func (o BTImmutableIntArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	return json.Marshal(toSerialize)
}

type NullableBTImmutableIntArray struct {
	value *BTImmutableIntArray
	isSet bool
}

func (v NullableBTImmutableIntArray) Get() *BTImmutableIntArray {
	return v.value
}

func (v *NullableBTImmutableIntArray) Set(val *BTImmutableIntArray) {
	v.value = val
	v.isSet = true
}

func (v NullableBTImmutableIntArray) IsSet() bool {
	return v.isSet
}

func (v *NullableBTImmutableIntArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTImmutableIntArray(val *BTImmutableIntArray) *NullableBTImmutableIntArray {
	return &NullableBTImmutableIntArray{value: val, isSet: true}
}

func (v NullableBTImmutableIntArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTImmutableIntArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
