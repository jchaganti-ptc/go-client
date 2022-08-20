/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6126-5c3a878ad24b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTTableCellParameter2399 - struct for BTTableCellParameter2399
type BTTableCellParameter2399 struct {
	implBTTableCellParameter2399 interface{}
}

// BTTableCellParameterWithValue2122AsBTTableCellParameter2399 is a convenience function that returns BTTableCellParameterWithValue2122 wrapped in BTTableCellParameter2399
func (o *BTTableCellParameterWithValue2122) AsBTTableCellParameter2399() *BTTableCellParameter2399 {
	return &BTTableCellParameter2399{o}
}

// BTTableCellPropertyParameter2983AsBTTableCellParameter2399 is a convenience function that returns BTTableCellPropertyParameter2983 wrapped in BTTableCellParameter2399
func (o *BTTableCellPropertyParameter2983) AsBTTableCellParameter2399() *BTTableCellParameter2399 {
	return &BTTableCellParameter2399{o}
}

// NewBTTableCellParameter2399 instantiates a new BTTableCellParameter2399 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableCellParameter2399() *BTTableCellParameter2399 {
	this := BTTableCellParameter2399{Newbase_BTTableCellParameter2399()}
	return &this
}

// NewBTTableCellParameter2399WithDefaults instantiates a new BTTableCellParameter2399 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableCellParameter2399WithDefaults() *BTTableCellParameter2399 {
	this := BTTableCellParameter2399{Newbase_BTTableCellParameter2399WithDefaults()}
	return &this
}

// GetIsEverVisible returns the IsEverVisible field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetIsEverVisible() bool {
	type getResult interface {
		GetIsEverVisible() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsEverVisible()
	} else {
		var de bool
		return de
	}
}

// GetIsEverVisibleOk returns a tuple with the IsEverVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetIsEverVisibleOk() (*bool, bool) {
	type getResult interface {
		GetIsEverVisibleOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsEverVisibleOk()
	} else {
		return nil, false
	}
}

// HasIsEverVisible returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasIsEverVisible() bool {
	type getResult interface {
		HasIsEverVisible() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsEverVisible()
	} else {
		return false
	}
}

// SetIsEverVisible gets a reference to the given bool and assigns it to the IsEverVisible field.
func (o *BTTableCellParameter2399) SetIsEverVisible(v bool) {
	type getResult interface {
		SetIsEverVisible(v bool)
	}

	o.GetActualInstance().(getResult).SetIsEverVisible(v)
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetIsReadOnly() bool {
	type getResult interface {
		GetIsReadOnly() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsReadOnly()
	} else {
		var de bool
		return de
	}
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetIsReadOnlyOk() (*bool, bool) {
	type getResult interface {
		GetIsReadOnlyOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIsReadOnlyOk()
	} else {
		return nil, false
	}
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasIsReadOnly() bool {
	type getResult interface {
		HasIsReadOnly() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIsReadOnly()
	} else {
		return false
	}
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *BTTableCellParameter2399) SetIsReadOnly(v bool) {
	type getResult interface {
		SetIsReadOnly(v bool)
	}

	o.GetActualInstance().(getResult).SetIsReadOnly(v)
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetBtType() string {
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
func (o *BTTableCellParameter2399) GetBtTypeOk() (*string, bool) {
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
func (o *BTTableCellParameter2399) HasBtType() bool {
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
func (o *BTTableCellParameter2399) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetError() string {
	type getResult interface {
		GetError() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetError()
	} else {
		var de string
		return de
	}
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetErrorOk() (*string, bool) {
	type getResult interface {
		GetErrorOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorOk()
	} else {
		return nil, false
	}
}

// HasError returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasError() bool {
	type getResult interface {
		HasError() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasError()
	} else {
		return false
	}
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BTTableCellParameter2399) SetError(v string) {
	type getResult interface {
		SetError(v string)
	}

	o.GetActualInstance().(getResult).SetError(v)
}

// GetOverrideSpec returns the OverrideSpec field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetOverrideSpec() BTParameterSpec6 {
	type getResult interface {
		GetOverrideSpec() BTParameterSpec6
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOverrideSpec()
	} else {
		var de BTParameterSpec6
		return de
	}
}

// GetOverrideSpecOk returns a tuple with the OverrideSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetOverrideSpecOk() (*BTParameterSpec6, bool) {
	type getResult interface {
		GetOverrideSpecOk() (*BTParameterSpec6, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOverrideSpecOk()
	} else {
		return nil, false
	}
}

// HasOverrideSpec returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasOverrideSpec() bool {
	type getResult interface {
		HasOverrideSpec() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasOverrideSpec()
	} else {
		return false
	}
}

// SetOverrideSpec gets a reference to the given BTParameterSpec6 and assigns it to the OverrideSpec field.
func (o *BTTableCellParameter2399) SetOverrideSpec(v BTParameterSpec6) {
	type getResult interface {
		SetOverrideSpec(v BTParameterSpec6)
	}

	o.GetActualInstance().(getResult).SetOverrideSpec(v)
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetParameter() BTMParameter1 {
	type getResult interface {
		GetParameter() BTMParameter1
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameter()
	} else {
		var de BTMParameter1
		return de
	}
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetParameterOk() (*BTMParameter1, bool) {
	type getResult interface {
		GetParameterOk() (*BTMParameter1, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterOk()
	} else {
		return nil, false
	}
}

// HasParameter returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasParameter() bool {
	type getResult interface {
		HasParameter() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParameter()
	} else {
		return false
	}
}

// SetParameter gets a reference to the given BTMParameter1 and assigns it to the Parameter field.
func (o *BTTableCellParameter2399) SetParameter(v BTMParameter1) {
	type getResult interface {
		SetParameter(v BTMParameter1)
	}

	o.GetActualInstance().(getResult).SetParameter(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTTableCellParameter2399) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTTableCellParameterWithValue-2122'
	if jsonDict["btType"] == "BTTableCellParameterWithValue-2122" {
		// try to unmarshal JSON data into BTTableCellParameterWithValue2122
		var qr *BTTableCellParameterWithValue2122
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableCellParameter2399 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableCellParameter2399 = nil
			return fmt.Errorf("Failed to unmarshal BTTableCellParameter2399 as BTTableCellParameterWithValue2122: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTableCellPropertyParameter-2983'
	if jsonDict["btType"] == "BTTableCellPropertyParameter-2983" {
		// try to unmarshal JSON data into BTTableCellPropertyParameter2983
		var qr *BTTableCellPropertyParameter2983
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableCellParameter2399 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableCellParameter2399 = nil
			return fmt.Errorf("Failed to unmarshal BTTableCellParameter2399 as BTTableCellPropertyParameter2983: %s", err.Error())
		}
	}

	var qtx *base_BTTableCellParameter2399
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTTableCellParameter2399 = qtx
		return nil // data stored in dst.base_BTTableCellParameter2399, return on the first match
	} else {
		dst.implBTTableCellParameter2399 = nil
		return fmt.Errorf("Failed to unmarshal BTTableCellParameter2399 as base_BTTableCellParameter2399: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTTableCellParameter2399) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTTableCellParameter2399) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTTableCellParameter2399
}

type NullableBTTableCellParameter2399 struct {
	value *BTTableCellParameter2399
	isSet bool
}

func (v NullableBTTableCellParameter2399) Get() *BTTableCellParameter2399 {
	return v.value
}

func (v *NullableBTTableCellParameter2399) Set(val *BTTableCellParameter2399) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableCellParameter2399) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableCellParameter2399) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableCellParameter2399(val *BTTableCellParameter2399) *NullableBTTableCellParameter2399 {
	return &NullableBTTableCellParameter2399{value: val, isSet: true}
}

func (v NullableBTTableCellParameter2399) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableCellParameter2399) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTTableCellParameter2399 struct {
	IsEverVisible *bool             `json:"isEverVisible,omitempty"`
	IsReadOnly    *bool             `json:"isReadOnly,omitempty"`
	BtType        *string           `json:"btType,omitempty"`
	Error         *string           `json:"error,omitempty"`
	OverrideSpec  *BTParameterSpec6 `json:"overrideSpec,omitempty"`
	Parameter     *BTMParameter1    `json:"parameter,omitempty"`
}

// Newbase_BTTableCellParameter2399 instantiates a new base_BTTableCellParameter2399 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTTableCellParameter2399() *base_BTTableCellParameter2399 {
	this := base_BTTableCellParameter2399{}
	return &this
}

// Newbase_BTTableCellParameter2399WithDefaults instantiates a new base_BTTableCellParameter2399 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTTableCellParameter2399WithDefaults() *base_BTTableCellParameter2399 {
	this := base_BTTableCellParameter2399{}
	return &this
}

// GetIsEverVisible returns the IsEverVisible field value if set, zero value otherwise.
func (o *base_BTTableCellParameter2399) GetIsEverVisible() bool {
	if o == nil || o.IsEverVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsEverVisible
}

// GetIsEverVisibleOk returns a tuple with the IsEverVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCellParameter2399) GetIsEverVisibleOk() (*bool, bool) {
	if o == nil || o.IsEverVisible == nil {
		return nil, false
	}
	return o.IsEverVisible, true
}

// HasIsEverVisible returns a boolean if a field has been set.
func (o *base_BTTableCellParameter2399) HasIsEverVisible() bool {
	if o != nil && o.IsEverVisible != nil {
		return true
	}

	return false
}

// SetIsEverVisible gets a reference to the given bool and assigns it to the IsEverVisible field.
func (o *base_BTTableCellParameter2399) SetIsEverVisible(v bool) {
	o.IsEverVisible = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *base_BTTableCellParameter2399) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCellParameter2399) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *base_BTTableCellParameter2399) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *base_BTTableCellParameter2399) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTTableCellParameter2399) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCellParameter2399) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTTableCellParameter2399) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTTableCellParameter2399) SetBtType(v string) {
	o.BtType = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *base_BTTableCellParameter2399) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCellParameter2399) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *base_BTTableCellParameter2399) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *base_BTTableCellParameter2399) SetError(v string) {
	o.Error = &v
}

// GetOverrideSpec returns the OverrideSpec field value if set, zero value otherwise.
func (o *base_BTTableCellParameter2399) GetOverrideSpec() BTParameterSpec6 {
	if o == nil || o.OverrideSpec == nil {
		var ret BTParameterSpec6
		return ret
	}
	return *o.OverrideSpec
}

// GetOverrideSpecOk returns a tuple with the OverrideSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCellParameter2399) GetOverrideSpecOk() (*BTParameterSpec6, bool) {
	if o == nil || o.OverrideSpec == nil {
		return nil, false
	}
	return o.OverrideSpec, true
}

// HasOverrideSpec returns a boolean if a field has been set.
func (o *base_BTTableCellParameter2399) HasOverrideSpec() bool {
	if o != nil && o.OverrideSpec != nil {
		return true
	}

	return false
}

// SetOverrideSpec gets a reference to the given BTParameterSpec6 and assigns it to the OverrideSpec field.
func (o *base_BTTableCellParameter2399) SetOverrideSpec(v BTParameterSpec6) {
	o.OverrideSpec = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *base_BTTableCellParameter2399) GetParameter() BTMParameter1 {
	if o == nil || o.Parameter == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableCellParameter2399) GetParameterOk() (*BTMParameter1, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *base_BTTableCellParameter2399) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given BTMParameter1 and assigns it to the Parameter field.
func (o *base_BTTableCellParameter2399) SetParameter(v BTMParameter1) {
	o.Parameter = &v
}

func (o base_BTTableCellParameter2399) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEverVisible != nil {
		toSerialize["isEverVisible"] = o.IsEverVisible
	}
	if o.IsReadOnly != nil {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.OverrideSpec != nil {
		toSerialize["overrideSpec"] = o.OverrideSpec
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	return json.Marshal(toSerialize)
}
