/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6135-f341d0cf30dd
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTEventParams - struct for BTEventParams
type BTEventParams struct {
	implBTEventParams interface{}
}

// BTDocumentOpenEventParamsAsBTEventParams is a convenience function that returns BTDocumentOpenEventParams wrapped in BTEventParams
func (o *BTDocumentOpenEventParams) AsBTEventParams() *BTEventParams {
	return &BTEventParams{o}
}

// NewBTEventParams instantiates a new BTEventParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEventParams() *BTEventParams {
	this := BTEventParams{Newbase_BTEventParams()}
	return &this
}

// NewBTEventParamsWithDefaults instantiates a new BTEventParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEventParamsWithDefaults() *BTEventParams {
	this := BTEventParams{Newbase_BTEventParamsWithDefaults()}
	return &this
}

// GetJsonType returns the JsonType field value if set, zero value otherwise.
func (o *BTEventParams) GetJsonType() string {
	type getResult interface {
		GetJsonType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetJsonType()
	} else {
		var de string
		return de
	}
}

// GetJsonTypeOk returns a tuple with the JsonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEventParams) GetJsonTypeOk() (*string, bool) {
	type getResult interface {
		GetJsonTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetJsonTypeOk()
	} else {
		return nil, false
	}
}

// HasJsonType returns a boolean if a field has been set.
func (o *BTEventParams) HasJsonType() bool {
	type getResult interface {
		HasJsonType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasJsonType()
	} else {
		return false
	}
}

// SetJsonType gets a reference to the given string and assigns it to the JsonType field.
func (o *BTEventParams) SetJsonType(v string) {
	type getResult interface {
		SetJsonType(v string)
	}

	o.GetActualInstance().(getResult).SetJsonType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTEventParams) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTDocumentOpenEventParams'
	if jsonDict["jsonType"] == "BTDocumentOpenEventParams" {
		// try to unmarshal JSON data into BTDocumentOpenEventParams
		var qr *BTDocumentOpenEventParams
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEventParams = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEventParams = nil
			return fmt.Errorf("Failed to unmarshal BTEventParams as BTDocumentOpenEventParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DocumentOpenEventInfo'
	if jsonDict["jsonType"] == "DocumentOpenEventInfo" {
		// try to unmarshal JSON data into BTDocumentOpenEventParams
		var qr *BTDocumentOpenEventParams
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTEventParams = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTEventParams = nil
			return fmt.Errorf("Failed to unmarshal BTEventParams as BTDocumentOpenEventParams: %s", err.Error())
		}
	}

	var qtx *base_BTEventParams
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTEventParams = qtx
		return nil // data stored in dst.base_BTEventParams, return on the first match
	} else {
		dst.implBTEventParams = nil
		return fmt.Errorf("Failed to unmarshal BTEventParams as base_BTEventParams: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTEventParams) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTEventParams) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTEventParams
}

type NullableBTEventParams struct {
	value *BTEventParams
	isSet bool
}

func (v NullableBTEventParams) Get() *BTEventParams {
	return v.value
}

func (v *NullableBTEventParams) Set(val *BTEventParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEventParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEventParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEventParams(val *BTEventParams) *NullableBTEventParams {
	return &NullableBTEventParams{value: val, isSet: true}
}

func (v NullableBTEventParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEventParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTEventParams struct {
	JsonType *string `json:"jsonType,omitempty"`
}

// Newbase_BTEventParams instantiates a new base_BTEventParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTEventParams() *base_BTEventParams {
	this := base_BTEventParams{}
	return &this
}

// Newbase_BTEventParamsWithDefaults instantiates a new base_BTEventParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTEventParamsWithDefaults() *base_BTEventParams {
	this := base_BTEventParams{}
	return &this
}

// GetJsonType returns the JsonType field value if set, zero value otherwise.
func (o *base_BTEventParams) GetJsonType() string {
	if o == nil || o.JsonType == nil {
		var ret string
		return ret
	}
	return *o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTEventParams) GetJsonTypeOk() (*string, bool) {
	if o == nil || o.JsonType == nil {
		return nil, false
	}
	return o.JsonType, true
}

// HasJsonType returns a boolean if a field has been set.
func (o *base_BTEventParams) HasJsonType() bool {
	if o != nil && o.JsonType != nil {
		return true
	}

	return false
}

// SetJsonType gets a reference to the given string and assigns it to the JsonType field.
func (o *base_BTEventParams) SetJsonType(v string) {
	o.JsonType = &v
}

func (o base_BTEventParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JsonType != nil {
		toSerialize["jsonType"] = o.JsonType
	}
	return json.Marshal(toSerialize)
}
