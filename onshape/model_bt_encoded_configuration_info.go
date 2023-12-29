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

// BTEncodedConfigurationInfo struct for BTEncodedConfigurationInfo
type BTEncodedConfigurationInfo struct {
	EncodedId  *string `json:"encodedId,omitempty"`
	QueryParam *string `json:"queryParam,omitempty"`
}

// NewBTEncodedConfigurationInfo instantiates a new BTEncodedConfigurationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEncodedConfigurationInfo() *BTEncodedConfigurationInfo {
	this := BTEncodedConfigurationInfo{}
	return &this
}

// NewBTEncodedConfigurationInfoWithDefaults instantiates a new BTEncodedConfigurationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEncodedConfigurationInfoWithDefaults() *BTEncodedConfigurationInfo {
	this := BTEncodedConfigurationInfo{}
	return &this
}

// GetEncodedId returns the EncodedId field value if set, zero value otherwise.
func (o *BTEncodedConfigurationInfo) GetEncodedId() string {
	if o == nil || o.EncodedId == nil {
		var ret string
		return ret
	}
	return *o.EncodedId
}

// GetEncodedIdOk returns a tuple with the EncodedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEncodedConfigurationInfo) GetEncodedIdOk() (*string, bool) {
	if o == nil || o.EncodedId == nil {
		return nil, false
	}
	return o.EncodedId, true
}

// HasEncodedId returns a boolean if a field has been set.
func (o *BTEncodedConfigurationInfo) HasEncodedId() bool {
	if o != nil && o.EncodedId != nil {
		return true
	}

	return false
}

// SetEncodedId gets a reference to the given string and assigns it to the EncodedId field.
func (o *BTEncodedConfigurationInfo) SetEncodedId(v string) {
	o.EncodedId = &v
}

// GetQueryParam returns the QueryParam field value if set, zero value otherwise.
func (o *BTEncodedConfigurationInfo) GetQueryParam() string {
	if o == nil || o.QueryParam == nil {
		var ret string
		return ret
	}
	return *o.QueryParam
}

// GetQueryParamOk returns a tuple with the QueryParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEncodedConfigurationInfo) GetQueryParamOk() (*string, bool) {
	if o == nil || o.QueryParam == nil {
		return nil, false
	}
	return o.QueryParam, true
}

// HasQueryParam returns a boolean if a field has been set.
func (o *BTEncodedConfigurationInfo) HasQueryParam() bool {
	if o != nil && o.QueryParam != nil {
		return true
	}

	return false
}

// SetQueryParam gets a reference to the given string and assigns it to the QueryParam field.
func (o *BTEncodedConfigurationInfo) SetQueryParam(v string) {
	o.QueryParam = &v
}

func (o BTEncodedConfigurationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncodedId != nil {
		toSerialize["encodedId"] = o.EncodedId
	}
	if o.QueryParam != nil {
		toSerialize["queryParam"] = o.QueryParam
	}
	return json.Marshal(toSerialize)
}

type NullableBTEncodedConfigurationInfo struct {
	value *BTEncodedConfigurationInfo
	isSet bool
}

func (v NullableBTEncodedConfigurationInfo) Get() *BTEncodedConfigurationInfo {
	return v.value
}

func (v *NullableBTEncodedConfigurationInfo) Set(val *BTEncodedConfigurationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEncodedConfigurationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEncodedConfigurationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEncodedConfigurationInfo(val *BTEncodedConfigurationInfo) *NullableBTEncodedConfigurationInfo {
	return &NullableBTEncodedConfigurationInfo{value: val, isSet: true}
}

func (v NullableBTEncodedConfigurationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEncodedConfigurationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
