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

// BTAssemblyFeatureDataInfo struct for BTAssemblyFeatureDataInfo
type BTAssemblyFeatureDataInfo struct {
	Name *string `json:"name,omitempty"`
}

// NewBTAssemblyFeatureDataInfo instantiates a new BTAssemblyFeatureDataInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyFeatureDataInfo() *BTAssemblyFeatureDataInfo {
	this := BTAssemblyFeatureDataInfo{}
	return &this
}

// NewBTAssemblyFeatureDataInfoWithDefaults instantiates a new BTAssemblyFeatureDataInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyFeatureDataInfoWithDefaults() *BTAssemblyFeatureDataInfo {
	this := BTAssemblyFeatureDataInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAssemblyFeatureDataInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureDataInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAssemblyFeatureDataInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAssemblyFeatureDataInfo) SetName(v string) {
	o.Name = &v
}

func (o BTAssemblyFeatureDataInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyFeatureDataInfo struct {
	value *BTAssemblyFeatureDataInfo
	isSet bool
}

func (v NullableBTAssemblyFeatureDataInfo) Get() *BTAssemblyFeatureDataInfo {
	return v.value
}

func (v *NullableBTAssemblyFeatureDataInfo) Set(val *BTAssemblyFeatureDataInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyFeatureDataInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyFeatureDataInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyFeatureDataInfo(val *BTAssemblyFeatureDataInfo) *NullableBTAssemblyFeatureDataInfo {
	return &NullableBTAssemblyFeatureDataInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyFeatureDataInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyFeatureDataInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
