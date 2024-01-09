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

// MaterialNormalTextureInfo struct for MaterialNormalTextureInfo
type MaterialNormalTextureInfo struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Index      *int32                            `json:"index,omitempty"`
	Scale      *float32                          `json:"scale,omitempty"`
	TexCoord   *int32                            `json:"texCoord,omitempty"`
}

// NewMaterialNormalTextureInfo instantiates a new MaterialNormalTextureInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaterialNormalTextureInfo() *MaterialNormalTextureInfo {
	this := MaterialNormalTextureInfo{}
	return &this
}

// NewMaterialNormalTextureInfoWithDefaults instantiates a new MaterialNormalTextureInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaterialNormalTextureInfoWithDefaults() *MaterialNormalTextureInfo {
	this := MaterialNormalTextureInfo{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *MaterialNormalTextureInfo) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialNormalTextureInfo) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *MaterialNormalTextureInfo) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *MaterialNormalTextureInfo) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *MaterialNormalTextureInfo) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialNormalTextureInfo) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *MaterialNormalTextureInfo) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *MaterialNormalTextureInfo) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *MaterialNormalTextureInfo) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialNormalTextureInfo) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *MaterialNormalTextureInfo) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *MaterialNormalTextureInfo) SetIndex(v int32) {
	o.Index = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *MaterialNormalTextureInfo) GetScale() float32 {
	if o == nil || o.Scale == nil {
		var ret float32
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialNormalTextureInfo) GetScaleOk() (*float32, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *MaterialNormalTextureInfo) HasScale() bool {
	if o != nil && o.Scale != nil {
		return true
	}

	return false
}

// SetScale gets a reference to the given float32 and assigns it to the Scale field.
func (o *MaterialNormalTextureInfo) SetScale(v float32) {
	o.Scale = &v
}

// GetTexCoord returns the TexCoord field value if set, zero value otherwise.
func (o *MaterialNormalTextureInfo) GetTexCoord() int32 {
	if o == nil || o.TexCoord == nil {
		var ret int32
		return ret
	}
	return *o.TexCoord
}

// GetTexCoordOk returns a tuple with the TexCoord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialNormalTextureInfo) GetTexCoordOk() (*int32, bool) {
	if o == nil || o.TexCoord == nil {
		return nil, false
	}
	return o.TexCoord, true
}

// HasTexCoord returns a boolean if a field has been set.
func (o *MaterialNormalTextureInfo) HasTexCoord() bool {
	if o != nil && o.TexCoord != nil {
		return true
	}

	return false
}

// SetTexCoord gets a reference to the given int32 and assigns it to the TexCoord field.
func (o *MaterialNormalTextureInfo) SetTexCoord(v int32) {
	o.TexCoord = &v
}

func (o MaterialNormalTextureInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.TexCoord != nil {
		toSerialize["texCoord"] = o.TexCoord
	}
	return json.Marshal(toSerialize)
}

type NullableMaterialNormalTextureInfo struct {
	value *MaterialNormalTextureInfo
	isSet bool
}

func (v NullableMaterialNormalTextureInfo) Get() *MaterialNormalTextureInfo {
	return v.value
}

func (v *NullableMaterialNormalTextureInfo) Set(val *MaterialNormalTextureInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMaterialNormalTextureInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMaterialNormalTextureInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaterialNormalTextureInfo(val *MaterialNormalTextureInfo) *NullableMaterialNormalTextureInfo {
	return &NullableMaterialNormalTextureInfo{value: val, isSet: true}
}

func (v NullableMaterialNormalTextureInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaterialNormalTextureInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
