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

// BTFSValueOther1124 struct for BTFSValueOther1124
type BTFSValueOther1124 struct {
	BtType  string    `json:"btType"`
	TypeTag *string   `json:"typeTag,omitempty"`
	Type    *GBTPType `json:"type,omitempty"`
}

// NewBTFSValueOther1124 instantiates a new BTFSValueOther1124 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueOther1124(btType string) *BTFSValueOther1124 {
	this := BTFSValueOther1124{}
	this.BtType = btType
	return &this
}

// NewBTFSValueOther1124WithDefaults instantiates a new BTFSValueOther1124 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueOther1124WithDefaults() *BTFSValueOther1124 {
	this := BTFSValueOther1124{}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValueOther1124) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValueOther1124) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *BTFSValueOther1124) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValueOther1124) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueOther1124) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValueOther1124) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValueOther1124) SetTypeTag(v string) {
	o.TypeTag = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTFSValueOther1124) GetType() GBTPType {
	if o == nil || o.Type == nil {
		var ret GBTPType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueOther1124) GetTypeOk() (*GBTPType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTFSValueOther1124) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GBTPType and assigns it to the Type field.
func (o *BTFSValueOther1124) SetType(v GBTPType) {
	o.Type = &v
}

func (o BTFSValueOther1124) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["btType"] = o.BtType
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueOther1124 struct {
	value *BTFSValueOther1124
	isSet bool
}

func (v NullableBTFSValueOther1124) Get() *BTFSValueOther1124 {
	return v.value
}

func (v *NullableBTFSValueOther1124) Set(val *BTFSValueOther1124) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueOther1124) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueOther1124) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueOther1124(val *BTFSValueOther1124) *NullableBTFSValueOther1124 {
	return &NullableBTFSValueOther1124{value: val, isSet: true}
}

func (v NullableBTFSValueOther1124) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueOther1124) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
