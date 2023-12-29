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
	"fmt"
)

// BTMConfiguredValue1341 - struct for BTMConfiguredValue1341
type BTMConfiguredValue1341 struct {
	implBTMConfiguredValue1341 interface{}
}

// BTMConfiguredValueByEnum1923AsBTMConfiguredValue1341 is a convenience function that returns BTMConfiguredValueByEnum1923 wrapped in BTMConfiguredValue1341
func (o *BTMConfiguredValueByEnum1923) AsBTMConfiguredValue1341() *BTMConfiguredValue1341 {
	return &BTMConfiguredValue1341{o}
}

// BTMConfiguredValueByBoolean1501AsBTMConfiguredValue1341 is a convenience function that returns BTMConfiguredValueByBoolean1501 wrapped in BTMConfiguredValue1341
func (o *BTMConfiguredValueByBoolean1501) AsBTMConfiguredValue1341() *BTMConfiguredValue1341 {
	return &BTMConfiguredValue1341{o}
}

// NewBTMConfiguredValue1341 instantiates a new BTMConfiguredValue1341 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfiguredValue1341() *BTMConfiguredValue1341 {
	this := BTMConfiguredValue1341{Newbase_BTMConfiguredValue1341()}
	return &this
}

// NewBTMConfiguredValue1341WithDefaults instantiates a new BTMConfiguredValue1341 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfiguredValue1341WithDefaults() *BTMConfiguredValue1341 {
	this := BTMConfiguredValue1341{Newbase_BTMConfiguredValue1341WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfiguredValue1341) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetConfigurationValueString returns the ConfigurationValueString field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetConfigurationValueString() string {
	type getResult interface {
		GetConfigurationValueString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConfigurationValueString()
	} else {
		var de string
		return de
	}
}

// GetConfigurationValueStringOk returns a tuple with the ConfigurationValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetConfigurationValueStringOk() (*string, bool) {
	type getResult interface {
		GetConfigurationValueStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConfigurationValueStringOk()
	} else {
		return nil, false
	}
}

// HasConfigurationValueString returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasConfigurationValueString() bool {
	type getResult interface {
		HasConfigurationValueString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasConfigurationValueString()
	} else {
		return false
	}
}

// SetConfigurationValueString gets a reference to the given string and assigns it to the ConfigurationValueString field.
func (o *BTMConfiguredValue1341) SetConfigurationValueString(v string) {
	type getResult interface {
		SetConfigurationValueString(v string)
	}

	o.GetActualInstance().(getResult).SetConfigurationValueString(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfiguredValue1341) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfiguredValue1341) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMConfiguredValue1341) GetValue() BTMParameter1 {
	type getResult interface {
		GetValue() BTMParameter1
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValue()
	} else {
		var de BTMParameter1
		return de
	}
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValue1341) GetValueOk() (*BTMParameter1, bool) {
	type getResult interface {
		GetValueOk() (*BTMParameter1, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetValueOk()
	} else {
		return nil, false
	}
}

// HasValue returns a boolean if a field has been set.
func (o *BTMConfiguredValue1341) HasValue() bool {
	type getResult interface {
		HasValue() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasValue()
	} else {
		return false
	}
}

// SetValue gets a reference to the given BTMParameter1 and assigns it to the Value field.
func (o *BTMConfiguredValue1341) SetValue(v BTMParameter1) {
	type getResult interface {
		SetValue(v BTMParameter1)
	}

	o.GetActualInstance().(getResult).SetValue(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMConfiguredValue1341) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMConfiguredValueByBoolean-1501'
	if jsonDict["btType"] == "BTMConfiguredValueByBoolean-1501" {
		// try to unmarshal JSON data into BTMConfiguredValueByBoolean1501
		var qr *BTMConfiguredValueByBoolean1501
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMConfiguredValue1341 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMConfiguredValue1341 = nil
			return fmt.Errorf("Failed to unmarshal BTMConfiguredValue1341 as BTMConfiguredValueByBoolean1501: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMConfiguredValueByEnum-1923'
	if jsonDict["btType"] == "BTMConfiguredValueByEnum-1923" {
		// try to unmarshal JSON data into BTMConfiguredValueByEnum1923
		var qr *BTMConfiguredValueByEnum1923
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMConfiguredValue1341 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMConfiguredValue1341 = nil
			return fmt.Errorf("Failed to unmarshal BTMConfiguredValue1341 as BTMConfiguredValueByEnum1923: %s", err.Error())
		}
	}

	var qtx *base_BTMConfiguredValue1341
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMConfiguredValue1341 = qtx
		return nil // data stored in dst.base_BTMConfiguredValue1341, return on the first match
	} else {
		dst.implBTMConfiguredValue1341 = nil
		return fmt.Errorf("Failed to unmarshal BTMConfiguredValue1341 as base_BTMConfiguredValue1341: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMConfiguredValue1341) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMConfiguredValue1341) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMConfiguredValue1341
}

type NullableBTMConfiguredValue1341 struct {
	value *BTMConfiguredValue1341
	isSet bool
}

func (v NullableBTMConfiguredValue1341) Get() *BTMConfiguredValue1341 {
	return v.value
}

func (v *NullableBTMConfiguredValue1341) Set(val *BTMConfiguredValue1341) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfiguredValue1341) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfiguredValue1341) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfiguredValue1341(val *BTMConfiguredValue1341) *NullableBTMConfiguredValue1341 {
	return &NullableBTMConfiguredValue1341{value: val, isSet: true}
}

func (v NullableBTMConfiguredValue1341) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfiguredValue1341) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMConfiguredValue1341 struct {
	BtType                   *string        `json:"btType,omitempty"`
	ConfigurationValueString *string        `json:"configurationValueString,omitempty"`
	ImportMicroversion       *string        `json:"importMicroversion,omitempty"`
	NodeId                   *string        `json:"nodeId,omitempty"`
	Value                    *BTMParameter1 `json:"value,omitempty"`
}

// Newbase_BTMConfiguredValue1341 instantiates a new base_BTMConfiguredValue1341 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMConfiguredValue1341() *base_BTMConfiguredValue1341 {
	this := base_BTMConfiguredValue1341{}
	return &this
}

// Newbase_BTMConfiguredValue1341WithDefaults instantiates a new base_BTMConfiguredValue1341 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMConfiguredValue1341WithDefaults() *base_BTMConfiguredValue1341 {
	this := base_BTMConfiguredValue1341{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMConfiguredValue1341) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfiguredValue1341) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMConfiguredValue1341) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMConfiguredValue1341) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationValueString returns the ConfigurationValueString field value if set, zero value otherwise.
func (o *base_BTMConfiguredValue1341) GetConfigurationValueString() string {
	if o == nil || o.ConfigurationValueString == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationValueString
}

// GetConfigurationValueStringOk returns a tuple with the ConfigurationValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfiguredValue1341) GetConfigurationValueStringOk() (*string, bool) {
	if o == nil || o.ConfigurationValueString == nil {
		return nil, false
	}
	return o.ConfigurationValueString, true
}

// HasConfigurationValueString returns a boolean if a field has been set.
func (o *base_BTMConfiguredValue1341) HasConfigurationValueString() bool {
	if o != nil && o.ConfigurationValueString != nil {
		return true
	}

	return false
}

// SetConfigurationValueString gets a reference to the given string and assigns it to the ConfigurationValueString field.
func (o *base_BTMConfiguredValue1341) SetConfigurationValueString(v string) {
	o.ConfigurationValueString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMConfiguredValue1341) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfiguredValue1341) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMConfiguredValue1341) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMConfiguredValue1341) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMConfiguredValue1341) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfiguredValue1341) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMConfiguredValue1341) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMConfiguredValue1341) SetNodeId(v string) {
	o.NodeId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *base_BTMConfiguredValue1341) GetValue() BTMParameter1 {
	if o == nil || o.Value == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMConfiguredValue1341) GetValueOk() (*BTMParameter1, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *base_BTMConfiguredValue1341) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTMParameter1 and assigns it to the Value field.
func (o *base_BTMConfiguredValue1341) SetValue(v BTMParameter1) {
	o.Value = &v
}

func (o base_BTMConfiguredValue1341) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationValueString != nil {
		toSerialize["configurationValueString"] = o.ConfigurationValueString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}
