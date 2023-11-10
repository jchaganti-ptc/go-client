/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25652-944cf1af37c9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Buffer struct for Buffer
type Buffer struct {
	ByteLength *int32                            `json:"byteLength,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Uri        *string                           `json:"uri,omitempty"`
}

// NewBuffer instantiates a new Buffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuffer() *Buffer {
	this := Buffer{}
	return &this
}

// NewBufferWithDefaults instantiates a new Buffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBufferWithDefaults() *Buffer {
	this := Buffer{}
	return &this
}

// GetByteLength returns the ByteLength field value if set, zero value otherwise.
func (o *Buffer) GetByteLength() int32 {
	if o == nil || o.ByteLength == nil {
		var ret int32
		return ret
	}
	return *o.ByteLength
}

// GetByteLengthOk returns a tuple with the ByteLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buffer) GetByteLengthOk() (*int32, bool) {
	if o == nil || o.ByteLength == nil {
		return nil, false
	}
	return o.ByteLength, true
}

// HasByteLength returns a boolean if a field has been set.
func (o *Buffer) HasByteLength() bool {
	if o != nil && o.ByteLength != nil {
		return true
	}

	return false
}

// SetByteLength gets a reference to the given int32 and assigns it to the ByteLength field.
func (o *Buffer) SetByteLength(v int32) {
	o.ByteLength = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Buffer) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buffer) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Buffer) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Buffer) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Buffer) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buffer) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Buffer) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Buffer) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Buffer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buffer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Buffer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Buffer) SetName(v string) {
	o.Name = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Buffer) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buffer) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Buffer) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Buffer) SetUri(v string) {
	o.Uri = &v
}

func (o Buffer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ByteLength != nil {
		toSerialize["byteLength"] = o.ByteLength
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableBuffer struct {
	value *Buffer
	isSet bool
}

func (v NullableBuffer) Get() *Buffer {
	return v.value
}

func (v *NullableBuffer) Set(val *Buffer) {
	v.value = val
	v.isSet = true
}

func (v NullableBuffer) IsSet() bool {
	return v.isSet
}

func (v *NullableBuffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuffer(val *Buffer) *NullableBuffer {
	return &NullableBuffer{value: val, isSet: true}
}

func (v NullableBuffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
