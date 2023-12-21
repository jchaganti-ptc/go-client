/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27960-e195de6ac56c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Callback struct for Callback
type Callback struct {
	Empty      *bool                             `json:"empty,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Getref     *string                           `json:"get$ref,omitempty"`
}

// NewCallback instantiates a new Callback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallback() *Callback {
	this := Callback{}
	return &this
}

// NewCallbackWithDefaults instantiates a new Callback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallbackWithDefaults() *Callback {
	this := Callback{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *Callback) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Callback) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *Callback) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *Callback) SetEmpty(v bool) {
	o.Empty = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Callback) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Callback) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Callback) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Callback) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *Callback) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Callback) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *Callback) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *Callback) SetGetref(v string) {
	o.Getref = &v
}

func (o Callback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Getref != nil {
		toSerialize["get$ref"] = o.Getref
	}
	return json.Marshal(toSerialize)
}

type NullableCallback struct {
	value *Callback
	isSet bool
}

func (v NullableCallback) Get() *Callback {
	return v.value
}

func (v *NullableCallback) Set(val *Callback) {
	v.value = val
	v.isSet = true
}

func (v NullableCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallback(val *Callback) *NullableCallback {
	return &NullableCallback{value: val, isSet: true}
}

func (v NullableCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
