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
)

// BufferViewModel struct for BufferViewModel
type BufferViewModel struct {
	BufferModel    *BufferModel                   `json:"bufferModel,omitempty"`
	BufferViewData *BufferViewModelBufferViewData `json:"bufferViewData,omitempty"`
	ByteLength     *int32                         `json:"byteLength,omitempty"`
	ByteOffset     *int32                         `json:"byteOffset,omitempty"`
	ByteStride     *int32                         `json:"byteStride,omitempty"`
	Name           *string                        `json:"name,omitempty"`
	Target         *int32                         `json:"target,omitempty"`
}

// NewBufferViewModel instantiates a new BufferViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBufferViewModel() *BufferViewModel {
	this := BufferViewModel{}
	return &this
}

// NewBufferViewModelWithDefaults instantiates a new BufferViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBufferViewModelWithDefaults() *BufferViewModel {
	this := BufferViewModel{}
	return &this
}

// GetBufferModel returns the BufferModel field value if set, zero value otherwise.
func (o *BufferViewModel) GetBufferModel() BufferModel {
	if o == nil || o.BufferModel == nil {
		var ret BufferModel
		return ret
	}
	return *o.BufferModel
}

// GetBufferModelOk returns a tuple with the BufferModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModel) GetBufferModelOk() (*BufferModel, bool) {
	if o == nil || o.BufferModel == nil {
		return nil, false
	}
	return o.BufferModel, true
}

// HasBufferModel returns a boolean if a field has been set.
func (o *BufferViewModel) HasBufferModel() bool {
	if o != nil && o.BufferModel != nil {
		return true
	}

	return false
}

// SetBufferModel gets a reference to the given BufferModel and assigns it to the BufferModel field.
func (o *BufferViewModel) SetBufferModel(v BufferModel) {
	o.BufferModel = &v
}

// GetBufferViewData returns the BufferViewData field value if set, zero value otherwise.
func (o *BufferViewModel) GetBufferViewData() BufferViewModelBufferViewData {
	if o == nil || o.BufferViewData == nil {
		var ret BufferViewModelBufferViewData
		return ret
	}
	return *o.BufferViewData
}

// GetBufferViewDataOk returns a tuple with the BufferViewData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModel) GetBufferViewDataOk() (*BufferViewModelBufferViewData, bool) {
	if o == nil || o.BufferViewData == nil {
		return nil, false
	}
	return o.BufferViewData, true
}

// HasBufferViewData returns a boolean if a field has been set.
func (o *BufferViewModel) HasBufferViewData() bool {
	if o != nil && o.BufferViewData != nil {
		return true
	}

	return false
}

// SetBufferViewData gets a reference to the given BufferViewModelBufferViewData and assigns it to the BufferViewData field.
func (o *BufferViewModel) SetBufferViewData(v BufferViewModelBufferViewData) {
	o.BufferViewData = &v
}

// GetByteLength returns the ByteLength field value if set, zero value otherwise.
func (o *BufferViewModel) GetByteLength() int32 {
	if o == nil || o.ByteLength == nil {
		var ret int32
		return ret
	}
	return *o.ByteLength
}

// GetByteLengthOk returns a tuple with the ByteLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModel) GetByteLengthOk() (*int32, bool) {
	if o == nil || o.ByteLength == nil {
		return nil, false
	}
	return o.ByteLength, true
}

// HasByteLength returns a boolean if a field has been set.
func (o *BufferViewModel) HasByteLength() bool {
	if o != nil && o.ByteLength != nil {
		return true
	}

	return false
}

// SetByteLength gets a reference to the given int32 and assigns it to the ByteLength field.
func (o *BufferViewModel) SetByteLength(v int32) {
	o.ByteLength = &v
}

// GetByteOffset returns the ByteOffset field value if set, zero value otherwise.
func (o *BufferViewModel) GetByteOffset() int32 {
	if o == nil || o.ByteOffset == nil {
		var ret int32
		return ret
	}
	return *o.ByteOffset
}

// GetByteOffsetOk returns a tuple with the ByteOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModel) GetByteOffsetOk() (*int32, bool) {
	if o == nil || o.ByteOffset == nil {
		return nil, false
	}
	return o.ByteOffset, true
}

// HasByteOffset returns a boolean if a field has been set.
func (o *BufferViewModel) HasByteOffset() bool {
	if o != nil && o.ByteOffset != nil {
		return true
	}

	return false
}

// SetByteOffset gets a reference to the given int32 and assigns it to the ByteOffset field.
func (o *BufferViewModel) SetByteOffset(v int32) {
	o.ByteOffset = &v
}

// GetByteStride returns the ByteStride field value if set, zero value otherwise.
func (o *BufferViewModel) GetByteStride() int32 {
	if o == nil || o.ByteStride == nil {
		var ret int32
		return ret
	}
	return *o.ByteStride
}

// GetByteStrideOk returns a tuple with the ByteStride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModel) GetByteStrideOk() (*int32, bool) {
	if o == nil || o.ByteStride == nil {
		return nil, false
	}
	return o.ByteStride, true
}

// HasByteStride returns a boolean if a field has been set.
func (o *BufferViewModel) HasByteStride() bool {
	if o != nil && o.ByteStride != nil {
		return true
	}

	return false
}

// SetByteStride gets a reference to the given int32 and assigns it to the ByteStride field.
func (o *BufferViewModel) SetByteStride(v int32) {
	o.ByteStride = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BufferViewModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BufferViewModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BufferViewModel) SetName(v string) {
	o.Name = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *BufferViewModel) GetTarget() int32 {
	if o == nil || o.Target == nil {
		var ret int32
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModel) GetTargetOk() (*int32, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *BufferViewModel) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given int32 and assigns it to the Target field.
func (o *BufferViewModel) SetTarget(v int32) {
	o.Target = &v
}

func (o BufferViewModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BufferModel != nil {
		toSerialize["bufferModel"] = o.BufferModel
	}
	if o.BufferViewData != nil {
		toSerialize["bufferViewData"] = o.BufferViewData
	}
	if o.ByteLength != nil {
		toSerialize["byteLength"] = o.ByteLength
	}
	if o.ByteOffset != nil {
		toSerialize["byteOffset"] = o.ByteOffset
	}
	if o.ByteStride != nil {
		toSerialize["byteStride"] = o.ByteStride
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableBufferViewModel struct {
	value *BufferViewModel
	isSet bool
}

func (v NullableBufferViewModel) Get() *BufferViewModel {
	return v.value
}

func (v *NullableBufferViewModel) Set(val *BufferViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBufferViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBufferViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBufferViewModel(val *BufferViewModel) *NullableBufferViewModel {
	return &NullableBufferViewModel{value: val, isSet: true}
}

func (v NullableBufferViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBufferViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
