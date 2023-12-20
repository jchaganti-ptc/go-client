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

// BTUiFeatureStudioChecksum2438 struct for BTUiFeatureStudioChecksum2438
type BTUiFeatureStudioChecksum2438 struct {
	BtType       *string              `json:"btType,omitempty"`
	Crc32        *int32               `json:"crc32,omitempty"`
	Microversion *BTMicroversionId366 `json:"microversion,omitempty"`
}

// NewBTUiFeatureStudioChecksum2438 instantiates a new BTUiFeatureStudioChecksum2438 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUiFeatureStudioChecksum2438() *BTUiFeatureStudioChecksum2438 {
	this := BTUiFeatureStudioChecksum2438{}
	return &this
}

// NewBTUiFeatureStudioChecksum2438WithDefaults instantiates a new BTUiFeatureStudioChecksum2438 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUiFeatureStudioChecksum2438WithDefaults() *BTUiFeatureStudioChecksum2438 {
	this := BTUiFeatureStudioChecksum2438{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTUiFeatureStudioChecksum2438) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUiFeatureStudioChecksum2438) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTUiFeatureStudioChecksum2438) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTUiFeatureStudioChecksum2438) SetBtType(v string) {
	o.BtType = &v
}

// GetCrc32 returns the Crc32 field value if set, zero value otherwise.
func (o *BTUiFeatureStudioChecksum2438) GetCrc32() int32 {
	if o == nil || o.Crc32 == nil {
		var ret int32
		return ret
	}
	return *o.Crc32
}

// GetCrc32Ok returns a tuple with the Crc32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUiFeatureStudioChecksum2438) GetCrc32Ok() (*int32, bool) {
	if o == nil || o.Crc32 == nil {
		return nil, false
	}
	return o.Crc32, true
}

// HasCrc32 returns a boolean if a field has been set.
func (o *BTUiFeatureStudioChecksum2438) HasCrc32() bool {
	if o != nil && o.Crc32 != nil {
		return true
	}

	return false
}

// SetCrc32 gets a reference to the given int32 and assigns it to the Crc32 field.
func (o *BTUiFeatureStudioChecksum2438) SetCrc32(v int32) {
	o.Crc32 = &v
}

// GetMicroversion returns the Microversion field value if set, zero value otherwise.
func (o *BTUiFeatureStudioChecksum2438) GetMicroversion() BTMicroversionId366 {
	if o == nil || o.Microversion == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.Microversion
}

// GetMicroversionOk returns a tuple with the Microversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUiFeatureStudioChecksum2438) GetMicroversionOk() (*BTMicroversionId366, bool) {
	if o == nil || o.Microversion == nil {
		return nil, false
	}
	return o.Microversion, true
}

// HasMicroversion returns a boolean if a field has been set.
func (o *BTUiFeatureStudioChecksum2438) HasMicroversion() bool {
	if o != nil && o.Microversion != nil {
		return true
	}

	return false
}

// SetMicroversion gets a reference to the given BTMicroversionId366 and assigns it to the Microversion field.
func (o *BTUiFeatureStudioChecksum2438) SetMicroversion(v BTMicroversionId366) {
	o.Microversion = &v
}

func (o BTUiFeatureStudioChecksum2438) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Crc32 != nil {
		toSerialize["crc32"] = o.Crc32
	}
	if o.Microversion != nil {
		toSerialize["microversion"] = o.Microversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTUiFeatureStudioChecksum2438 struct {
	value *BTUiFeatureStudioChecksum2438
	isSet bool
}

func (v NullableBTUiFeatureStudioChecksum2438) Get() *BTUiFeatureStudioChecksum2438 {
	return v.value
}

func (v *NullableBTUiFeatureStudioChecksum2438) Set(val *BTUiFeatureStudioChecksum2438) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUiFeatureStudioChecksum2438) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUiFeatureStudioChecksum2438) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUiFeatureStudioChecksum2438(val *BTUiFeatureStudioChecksum2438) *NullableBTUiFeatureStudioChecksum2438 {
	return &NullableBTUiFeatureStudioChecksum2438{value: val, isSet: true}
}

func (v NullableBTUiFeatureStudioChecksum2438) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUiFeatureStudioChecksum2438) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
