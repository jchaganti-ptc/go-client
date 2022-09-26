/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6621-03f431879e4b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTFSValue1888 - struct for BTFSValue1888
type BTFSValue1888 struct {
	implBTFSValue1888 interface{}
}

// BTFSValueArray1499AsBTFSValue1888 is a convenience function that returns BTFSValueArray1499 wrapped in BTFSValue1888
func (o *BTFSValueArray1499) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// BTFSValueString1422AsBTFSValue1888 is a convenience function that returns BTFSValueString1422 wrapped in BTFSValue1888
func (o *BTFSValueString1422) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// BTFSValueUndefined2003AsBTFSValue1888 is a convenience function that returns BTFSValueUndefined2003 wrapped in BTFSValue1888
func (o *BTFSValueUndefined2003) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// BTFSValueMap2062AsBTFSValue1888 is a convenience function that returns BTFSValueMap2062 wrapped in BTFSValue1888
func (o *BTFSValueMap2062) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// BTFSValueNumber772AsBTFSValue1888 is a convenience function that returns BTFSValueNumber772 wrapped in BTFSValue1888
func (o *BTFSValueNumber772) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// BTFSValueBoolean1195AsBTFSValue1888 is a convenience function that returns BTFSValueBoolean1195 wrapped in BTFSValue1888
func (o *BTFSValueBoolean1195) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// BTFSValueOther1124AsBTFSValue1888 is a convenience function that returns BTFSValueOther1124 wrapped in BTFSValue1888
func (o *BTFSValueOther1124) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// BTFSValueTooBig1247AsBTFSValue1888 is a convenience function that returns BTFSValueTooBig1247 wrapped in BTFSValue1888
func (o *BTFSValueTooBig1247) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// BTFSValueWithUnits1817AsBTFSValue1888 is a convenience function that returns BTFSValueWithUnits1817 wrapped in BTFSValue1888
func (o *BTFSValueWithUnits1817) AsBTFSValue1888() *BTFSValue1888 {
	return &BTFSValue1888{o}
}

// NewBTFSValue1888 instantiates a new BTFSValue1888 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValue1888(btType string) *BTFSValue1888 {
	this := BTFSValue1888{Newbase_BTFSValue1888(btType)}
	return &this
}

// NewBTFSValue1888WithDefaults instantiates a new BTFSValue1888 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValue1888WithDefaults() *BTFSValue1888 {
	this := BTFSValue1888{Newbase_BTFSValue1888WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value
func (o *BTFSValue1888) GetBtType() string {
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

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *BTFSValue1888) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// SetBtType sets field value
func (o *BTFSValue1888) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *BTFSValue1888) GetTypeTag() string {
	type getResult interface {
		GetTypeTag() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTypeTag()
	} else {
		var de string
		return de
	}
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValue1888) GetTypeTagOk() (*string, bool) {
	type getResult interface {
		GetTypeTagOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTypeTagOk()
	} else {
		return nil, false
	}
}

// HasTypeTag returns a boolean if a field has been set.
func (o *BTFSValue1888) HasTypeTag() bool {
	type getResult interface {
		HasTypeTag() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTypeTag()
	} else {
		return false
	}
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *BTFSValue1888) SetTypeTag(v string) {
	type getResult interface {
		SetTypeTag(v string)
	}

	o.GetActualInstance().(getResult).SetTypeTag(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTFSValue1888) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTFSValueArray-1499'
	if jsonDict["btType"] == "BTFSValueArray-1499" {
		// try to unmarshal JSON data into BTFSValueArray1499
		var qr *BTFSValueArray1499
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueArray1499: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSValueBoolean-1195'
	if jsonDict["btType"] == "BTFSValueBoolean-1195" {
		// try to unmarshal JSON data into BTFSValueBoolean1195
		var qr *BTFSValueBoolean1195
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueBoolean1195: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSValueMap-2062'
	if jsonDict["btType"] == "BTFSValueMap-2062" {
		// try to unmarshal JSON data into BTFSValueMap2062
		var qr *BTFSValueMap2062
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueMap2062: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSValueNumber-772'
	if jsonDict["btType"] == "BTFSValueNumber-772" {
		// try to unmarshal JSON data into BTFSValueNumber772
		var qr *BTFSValueNumber772
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueNumber772: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSValueOther-1124'
	if jsonDict["btType"] == "BTFSValueOther-1124" {
		// try to unmarshal JSON data into BTFSValueOther1124
		var qr *BTFSValueOther1124
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueOther1124: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSValueString-1422'
	if jsonDict["btType"] == "BTFSValueString-1422" {
		// try to unmarshal JSON data into BTFSValueString1422
		var qr *BTFSValueString1422
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueString1422: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSValueTooBig-1247'
	if jsonDict["btType"] == "BTFSValueTooBig-1247" {
		// try to unmarshal JSON data into BTFSValueTooBig1247
		var qr *BTFSValueTooBig1247
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueTooBig1247: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSValueUndefined-2003'
	if jsonDict["btType"] == "BTFSValueUndefined-2003" {
		// try to unmarshal JSON data into BTFSValueUndefined2003
		var qr *BTFSValueUndefined2003
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueUndefined2003: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSValueWithUnits-1817'
	if jsonDict["btType"] == "BTFSValueWithUnits-1817" {
		// try to unmarshal JSON data into BTFSValueWithUnits1817
		var qr *BTFSValueWithUnits1817
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFSValue1888 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFSValue1888 = nil
			return fmt.Errorf("Failed to unmarshal BTFSValue1888 as BTFSValueWithUnits1817: %s", err.Error())
		}
	}

	var qtx *base_BTFSValue1888
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTFSValue1888 = qtx
		return nil // data stored in dst.base_BTFSValue1888, return on the first match
	} else {
		dst.implBTFSValue1888 = nil
		return fmt.Errorf("Failed to unmarshal BTFSValue1888 as base_BTFSValue1888: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTFSValue1888) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTFSValue1888) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTFSValue1888
}

type NullableBTFSValue1888 struct {
	value *BTFSValue1888
	isSet bool
}

func (v NullableBTFSValue1888) Get() *BTFSValue1888 {
	return v.value
}

func (v *NullableBTFSValue1888) Set(val *BTFSValue1888) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValue1888) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValue1888) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValue1888(val *BTFSValue1888) *NullableBTFSValue1888 {
	return &NullableBTFSValue1888{value: val, isSet: true}
}

func (v NullableBTFSValue1888) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValue1888) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTFSValue1888 struct {
	BtType  string  `json:"btType"`
	TypeTag *string `json:"typeTag,omitempty"`
}

// Newbase_BTFSValue1888 instantiates a new base_BTFSValue1888 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTFSValue1888(btType string) *base_BTFSValue1888 {
	this := base_BTFSValue1888{}
	this.BtType = btType
	return &this
}

// Newbase_BTFSValue1888WithDefaults instantiates a new base_BTFSValue1888 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTFSValue1888WithDefaults() *base_BTFSValue1888 {
	this := base_BTFSValue1888{}
	return &this
}

// GetBtType returns the BtType field value
func (o *base_BTFSValue1888) GetBtType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value
// and a boolean to check if the value has been set.
func (o *base_BTFSValue1888) GetBtTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtType, true
}

// SetBtType sets field value
func (o *base_BTFSValue1888) SetBtType(v string) {
	o.BtType = v
}

// GetTypeTag returns the TypeTag field value if set, zero value otherwise.
func (o *base_BTFSValue1888) GetTypeTag() string {
	if o == nil || o.TypeTag == nil {
		var ret string
		return ret
	}
	return *o.TypeTag
}

// GetTypeTagOk returns a tuple with the TypeTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFSValue1888) GetTypeTagOk() (*string, bool) {
	if o == nil || o.TypeTag == nil {
		return nil, false
	}
	return o.TypeTag, true
}

// HasTypeTag returns a boolean if a field has been set.
func (o *base_BTFSValue1888) HasTypeTag() bool {
	if o != nil && o.TypeTag != nil {
		return true
	}

	return false
}

// SetTypeTag gets a reference to the given string and assigns it to the TypeTag field.
func (o *base_BTFSValue1888) SetTypeTag(v string) {
	o.TypeTag = &v
}

func (o base_BTFSValue1888) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["btType"] = o.BtType
	}
	if o.TypeTag != nil {
		toSerialize["typeTag"] = o.TypeTag
	}
	return json.Marshal(toSerialize)
}
