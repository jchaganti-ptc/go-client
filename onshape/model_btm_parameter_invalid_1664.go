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

// BTMParameterInvalid1664 struct for BTMParameterInvalid1664
type BTMParameterInvalid1664 struct {
	BtType             *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
	ParameterId        *string `json:"parameterId,omitempty"`
}

// NewBTMParameterInvalid1664 instantiates a new BTMParameterInvalid1664 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterInvalid1664() *BTMParameterInvalid1664 {
	this := BTMParameterInvalid1664{}
	return &this
}

// NewBTMParameterInvalid1664WithDefaults instantiates a new BTMParameterInvalid1664 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterInvalid1664WithDefaults() *BTMParameterInvalid1664 {
	this := BTMParameterInvalid1664{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterInvalid1664) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterInvalid1664) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterInvalid1664) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterInvalid1664) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterInvalid1664) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterInvalid1664) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterInvalid1664) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterInvalid1664) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterInvalid1664) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterInvalid1664) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterInvalid1664) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterInvalid1664) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterInvalid1664) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterInvalid1664) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterInvalid1664) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterInvalid1664) SetParameterId(v string) {
	o.ParameterId = &v
}

func (o BTMParameterInvalid1664) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterInvalid1664 struct {
	value *BTMParameterInvalid1664
	isSet bool
}

func (v NullableBTMParameterInvalid1664) Get() *BTMParameterInvalid1664 {
	return v.value
}

func (v *NullableBTMParameterInvalid1664) Set(val *BTMParameterInvalid1664) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterInvalid1664) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterInvalid1664) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterInvalid1664(val *BTMParameterInvalid1664) *NullableBTMParameterInvalid1664 {
	return &NullableBTMParameterInvalid1664{value: val, isSet: true}
}

func (v NullableBTMParameterInvalid1664) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterInvalid1664) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
