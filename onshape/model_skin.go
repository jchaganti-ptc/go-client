/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28126-700645382199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Skin struct for Skin
type Skin struct {
	Extensions          map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras              *map[string]interface{}           `json:"extras,omitempty"`
	InverseBindMatrices *int32                            `json:"inverseBindMatrices,omitempty"`
	Joints              []int32                           `json:"joints,omitempty"`
	Name                *string                           `json:"name,omitempty"`
	Skeleton            *int32                            `json:"skeleton,omitempty"`
}

// NewSkin instantiates a new Skin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkin() *Skin {
	this := Skin{}
	return &this
}

// NewSkinWithDefaults instantiates a new Skin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkinWithDefaults() *Skin {
	this := Skin{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Skin) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skin) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Skin) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Skin) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Skin) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skin) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Skin) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Skin) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetInverseBindMatrices returns the InverseBindMatrices field value if set, zero value otherwise.
func (o *Skin) GetInverseBindMatrices() int32 {
	if o == nil || o.InverseBindMatrices == nil {
		var ret int32
		return ret
	}
	return *o.InverseBindMatrices
}

// GetInverseBindMatricesOk returns a tuple with the InverseBindMatrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skin) GetInverseBindMatricesOk() (*int32, bool) {
	if o == nil || o.InverseBindMatrices == nil {
		return nil, false
	}
	return o.InverseBindMatrices, true
}

// HasInverseBindMatrices returns a boolean if a field has been set.
func (o *Skin) HasInverseBindMatrices() bool {
	if o != nil && o.InverseBindMatrices != nil {
		return true
	}

	return false
}

// SetInverseBindMatrices gets a reference to the given int32 and assigns it to the InverseBindMatrices field.
func (o *Skin) SetInverseBindMatrices(v int32) {
	o.InverseBindMatrices = &v
}

// GetJoints returns the Joints field value if set, zero value otherwise.
func (o *Skin) GetJoints() []int32 {
	if o == nil || o.Joints == nil {
		var ret []int32
		return ret
	}
	return o.Joints
}

// GetJointsOk returns a tuple with the Joints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skin) GetJointsOk() ([]int32, bool) {
	if o == nil || o.Joints == nil {
		return nil, false
	}
	return o.Joints, true
}

// HasJoints returns a boolean if a field has been set.
func (o *Skin) HasJoints() bool {
	if o != nil && o.Joints != nil {
		return true
	}

	return false
}

// SetJoints gets a reference to the given []int32 and assigns it to the Joints field.
func (o *Skin) SetJoints(v []int32) {
	o.Joints = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Skin) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skin) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Skin) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Skin) SetName(v string) {
	o.Name = &v
}

// GetSkeleton returns the Skeleton field value if set, zero value otherwise.
func (o *Skin) GetSkeleton() int32 {
	if o == nil || o.Skeleton == nil {
		var ret int32
		return ret
	}
	return *o.Skeleton
}

// GetSkeletonOk returns a tuple with the Skeleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skin) GetSkeletonOk() (*int32, bool) {
	if o == nil || o.Skeleton == nil {
		return nil, false
	}
	return o.Skeleton, true
}

// HasSkeleton returns a boolean if a field has been set.
func (o *Skin) HasSkeleton() bool {
	if o != nil && o.Skeleton != nil {
		return true
	}

	return false
}

// SetSkeleton gets a reference to the given int32 and assigns it to the Skeleton field.
func (o *Skin) SetSkeleton(v int32) {
	o.Skeleton = &v
}

func (o Skin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.InverseBindMatrices != nil {
		toSerialize["inverseBindMatrices"] = o.InverseBindMatrices
	}
	if o.Joints != nil {
		toSerialize["joints"] = o.Joints
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Skeleton != nil {
		toSerialize["skeleton"] = o.Skeleton
	}
	return json.Marshal(toSerialize)
}

type NullableSkin struct {
	value *Skin
	isSet bool
}

func (v NullableSkin) Get() *Skin {
	return v.value
}

func (v *NullableSkin) Set(val *Skin) {
	v.value = val
	v.isSet = true
}

func (v NullableSkin) IsSet() bool {
	return v.isSet
}

func (v *NullableSkin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkin(val *Skin) *NullableSkin {
	return &NullableSkin{value: val, isSet: true}
}

func (v NullableSkin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
