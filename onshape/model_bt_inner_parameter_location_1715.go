/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5931-1b7b413e7f30
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTInnerParameterLocation1715 - struct for BTInnerParameterLocation1715
type BTInnerParameterLocation1715 struct {
	implBTInnerParameterLocation1715 interface{}
}

// BTInnerDerivedParameterLocation591AsBTInnerParameterLocation1715 is a convenience function that returns BTInnerDerivedParameterLocation591 wrapped in BTInnerParameterLocation1715
func (o *BTInnerDerivedParameterLocation591) AsBTInnerParameterLocation1715() *BTInnerParameterLocation1715 {
	return &BTInnerParameterLocation1715{o}
}

// BTInnerArrayParameterLocation2368AsBTInnerParameterLocation1715 is a convenience function that returns BTInnerArrayParameterLocation2368 wrapped in BTInnerParameterLocation1715
func (o *BTInnerArrayParameterLocation2368) AsBTInnerParameterLocation1715() *BTInnerParameterLocation1715 {
	return &BTInnerParameterLocation1715{o}
}

// NewBTInnerParameterLocation1715 instantiates a new BTInnerParameterLocation1715 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInnerParameterLocation1715() *BTInnerParameterLocation1715 {
	this := BTInnerParameterLocation1715{Newbase_BTInnerParameterLocation1715()}
	return &this
}

// NewBTInnerParameterLocation1715WithDefaults instantiates a new BTInnerParameterLocation1715 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInnerParameterLocation1715WithDefaults() *BTInnerParameterLocation1715 {
	this := BTInnerParameterLocation1715{Newbase_BTInnerParameterLocation1715WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInnerParameterLocation1715) GetBtType() string {
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
func (o *BTInnerParameterLocation1715) GetBtTypeOk() (*string, bool) {
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
func (o *BTInnerParameterLocation1715) HasBtType() bool {
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
func (o *BTInnerParameterLocation1715) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTInnerParameterLocation1715) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTInnerArrayParameterLocation-2368'
	if jsonDict["btType"] == "BTInnerArrayParameterLocation-2368" {
		// try to unmarshal JSON data into BTInnerArrayParameterLocation2368
		var qr *BTInnerArrayParameterLocation2368
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTInnerParameterLocation1715 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTInnerParameterLocation1715 = nil
			return fmt.Errorf("Failed to unmarshal BTInnerParameterLocation1715 as BTInnerArrayParameterLocation2368: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTInnerDerivedParameterLocation-591'
	if jsonDict["btType"] == "BTInnerDerivedParameterLocation-591" {
		// try to unmarshal JSON data into BTInnerDerivedParameterLocation591
		var qr *BTInnerDerivedParameterLocation591
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTInnerParameterLocation1715 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTInnerParameterLocation1715 = nil
			return fmt.Errorf("Failed to unmarshal BTInnerParameterLocation1715 as BTInnerDerivedParameterLocation591: %s", err.Error())
		}
	}

	var qtx *base_BTInnerParameterLocation1715
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTInnerParameterLocation1715 = qtx
		return nil // data stored in dst.base_BTInnerParameterLocation1715, return on the first match
	} else {
		dst.implBTInnerParameterLocation1715 = nil
		return fmt.Errorf("Failed to unmarshal BTInnerParameterLocation1715 as base_BTInnerParameterLocation1715: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTInnerParameterLocation1715) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTInnerParameterLocation1715) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTInnerParameterLocation1715
}

type NullableBTInnerParameterLocation1715 struct {
	value *BTInnerParameterLocation1715
	isSet bool
}

func (v NullableBTInnerParameterLocation1715) Get() *BTInnerParameterLocation1715 {
	return v.value
}

func (v *NullableBTInnerParameterLocation1715) Set(val *BTInnerParameterLocation1715) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInnerParameterLocation1715) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInnerParameterLocation1715) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInnerParameterLocation1715(val *BTInnerParameterLocation1715) *NullableBTInnerParameterLocation1715 {
	return &NullableBTInnerParameterLocation1715{value: val, isSet: true}
}

func (v NullableBTInnerParameterLocation1715) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInnerParameterLocation1715) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTInnerParameterLocation1715 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTInnerParameterLocation1715 instantiates a new base_BTInnerParameterLocation1715 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTInnerParameterLocation1715() *base_BTInnerParameterLocation1715 {
	this := base_BTInnerParameterLocation1715{}
	return &this
}

// Newbase_BTInnerParameterLocation1715WithDefaults instantiates a new base_BTInnerParameterLocation1715 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTInnerParameterLocation1715WithDefaults() *base_BTInnerParameterLocation1715 {
	this := base_BTInnerParameterLocation1715{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTInnerParameterLocation1715) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTInnerParameterLocation1715) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTInnerParameterLocation1715) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTInnerParameterLocation1715) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTInnerParameterLocation1715) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
