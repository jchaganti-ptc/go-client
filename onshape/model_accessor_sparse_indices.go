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
)

// AccessorSparseIndices struct for AccessorSparseIndices
type AccessorSparseIndices struct {
	BufferView    *int32                            `json:"bufferView,omitempty"`
	ByteOffset    *int32                            `json:"byteOffset,omitempty"`
	ComponentType *int32                            `json:"componentType,omitempty"`
	Extensions    map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras        *map[string]interface{}           `json:"extras,omitempty"`
}

// NewAccessorSparseIndices instantiates a new AccessorSparseIndices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessorSparseIndices() *AccessorSparseIndices {
	this := AccessorSparseIndices{}
	return &this
}

// NewAccessorSparseIndicesWithDefaults instantiates a new AccessorSparseIndices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessorSparseIndicesWithDefaults() *AccessorSparseIndices {
	this := AccessorSparseIndices{}
	return &this
}

// GetBufferView returns the BufferView field value if set, zero value otherwise.
func (o *AccessorSparseIndices) GetBufferView() int32 {
	if o == nil || o.BufferView == nil {
		var ret int32
		return ret
	}
	return *o.BufferView
}

// GetBufferViewOk returns a tuple with the BufferView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseIndices) GetBufferViewOk() (*int32, bool) {
	if o == nil || o.BufferView == nil {
		return nil, false
	}
	return o.BufferView, true
}

// HasBufferView returns a boolean if a field has been set.
func (o *AccessorSparseIndices) HasBufferView() bool {
	if o != nil && o.BufferView != nil {
		return true
	}

	return false
}

// SetBufferView gets a reference to the given int32 and assigns it to the BufferView field.
func (o *AccessorSparseIndices) SetBufferView(v int32) {
	o.BufferView = &v
}

// GetByteOffset returns the ByteOffset field value if set, zero value otherwise.
func (o *AccessorSparseIndices) GetByteOffset() int32 {
	if o == nil || o.ByteOffset == nil {
		var ret int32
		return ret
	}
	return *o.ByteOffset
}

// GetByteOffsetOk returns a tuple with the ByteOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseIndices) GetByteOffsetOk() (*int32, bool) {
	if o == nil || o.ByteOffset == nil {
		return nil, false
	}
	return o.ByteOffset, true
}

// HasByteOffset returns a boolean if a field has been set.
func (o *AccessorSparseIndices) HasByteOffset() bool {
	if o != nil && o.ByteOffset != nil {
		return true
	}

	return false
}

// SetByteOffset gets a reference to the given int32 and assigns it to the ByteOffset field.
func (o *AccessorSparseIndices) SetByteOffset(v int32) {
	o.ByteOffset = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *AccessorSparseIndices) GetComponentType() int32 {
	if o == nil || o.ComponentType == nil {
		var ret int32
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseIndices) GetComponentTypeOk() (*int32, bool) {
	if o == nil || o.ComponentType == nil {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *AccessorSparseIndices) HasComponentType() bool {
	if o != nil && o.ComponentType != nil {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given int32 and assigns it to the ComponentType field.
func (o *AccessorSparseIndices) SetComponentType(v int32) {
	o.ComponentType = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *AccessorSparseIndices) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseIndices) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *AccessorSparseIndices) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *AccessorSparseIndices) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *AccessorSparseIndices) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparseIndices) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *AccessorSparseIndices) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *AccessorSparseIndices) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

func (o AccessorSparseIndices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BufferView != nil {
		toSerialize["bufferView"] = o.BufferView
	}
	if o.ByteOffset != nil {
		toSerialize["byteOffset"] = o.ByteOffset
	}
	if o.ComponentType != nil {
		toSerialize["componentType"] = o.ComponentType
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	return json.Marshal(toSerialize)
}

type NullableAccessorSparseIndices struct {
	value *AccessorSparseIndices
	isSet bool
}

func (v NullableAccessorSparseIndices) Get() *AccessorSparseIndices {
	return v.value
}

func (v *NullableAccessorSparseIndices) Set(val *AccessorSparseIndices) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessorSparseIndices) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessorSparseIndices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessorSparseIndices(val *AccessorSparseIndices) *NullableAccessorSparseIndices {
	return &NullableAccessorSparseIndices{value: val, isSet: true}
}

func (v NullableAccessorSparseIndices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessorSparseIndices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
