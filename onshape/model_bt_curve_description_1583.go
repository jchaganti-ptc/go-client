/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.26882-0482adeaa8aa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTCurveDescription1583 - struct for BTCurveDescription1583
type BTCurveDescription1583 struct {
	implBTCurveDescription1583 interface{}
}

// BTCircleDescription1145AsBTCurveDescription1583 is a convenience function that returns BTCircleDescription1145 wrapped in BTCurveDescription1583
func (o *BTCircleDescription1145) AsBTCurveDescription1583() *BTCurveDescription1583 {
	return &BTCurveDescription1583{o}
}

// BTEllipseDescription866AsBTCurveDescription1583 is a convenience function that returns BTEllipseDescription866 wrapped in BTCurveDescription1583
func (o *BTEllipseDescription866) AsBTCurveDescription1583() *BTCurveDescription1583 {
	return &BTCurveDescription1583{o}
}

// BTSplineDescription2118AsBTCurveDescription1583 is a convenience function that returns BTSplineDescription2118 wrapped in BTCurveDescription1583
func (o *BTSplineDescription2118) AsBTCurveDescription1583() *BTCurveDescription1583 {
	return &BTCurveDescription1583{o}
}

// BTLineDescription1559AsBTCurveDescription1583 is a convenience function that returns BTLineDescription1559 wrapped in BTCurveDescription1583
func (o *BTLineDescription1559) AsBTCurveDescription1583() *BTCurveDescription1583 {
	return &BTCurveDescription1583{o}
}

// NewBTCurveDescription1583 instantiates a new BTCurveDescription1583 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveDescription1583() *BTCurveDescription1583 {
	this := BTCurveDescription1583{Newbase_BTCurveDescription1583()}
	return &this
}

// NewBTCurveDescription1583WithDefaults instantiates a new BTCurveDescription1583 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveDescription1583WithDefaults() *BTCurveDescription1583 {
	this := BTCurveDescription1583{Newbase_BTCurveDescription1583WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveDescription1583) GetBtType() string {
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
func (o *BTCurveDescription1583) GetBtTypeOk() (*string, bool) {
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
func (o *BTCurveDescription1583) HasBtType() bool {
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
func (o *BTCurveDescription1583) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTCurveDescription1583) GetDirection() BTVector3d389 {
	type getResult interface {
		GetDirection() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDirection()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveDescription1583) GetDirectionOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetDirectionOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDirectionOk()
	} else {
		return nil, false
	}
}

// HasDirection returns a boolean if a field has been set.
func (o *BTCurveDescription1583) HasDirection() bool {
	type getResult interface {
		HasDirection() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDirection()
	} else {
		return false
	}
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTCurveDescription1583) SetDirection(v BTVector3d389) {
	type getResult interface {
		SetDirection(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetDirection(v)
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *BTCurveDescription1583) GetDirectionOrientedWithFace() BTVector3d389 {
	type getResult interface {
		GetDirectionOrientedWithFace() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDirectionOrientedWithFace()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveDescription1583) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDirectionOrientedWithFaceOk()
	} else {
		return nil, false
	}
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *BTCurveDescription1583) HasDirectionOrientedWithFace() bool {
	type getResult interface {
		HasDirectionOrientedWithFace() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDirectionOrientedWithFace()
	} else {
		return false
	}
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *BTCurveDescription1583) SetDirectionOrientedWithFace(v BTVector3d389) {
	type getResult interface {
		SetDirectionOrientedWithFace(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetDirectionOrientedWithFace(v)
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTCurveDescription1583) GetOrigin() BTVector3d389 {
	type getResult interface {
		GetOrigin() BTVector3d389
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOrigin()
	} else {
		var de BTVector3d389
		return de
	}
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveDescription1583) GetOriginOk() (*BTVector3d389, bool) {
	type getResult interface {
		GetOriginOk() (*BTVector3d389, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOriginOk()
	} else {
		return nil, false
	}
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTCurveDescription1583) HasOrigin() bool {
	type getResult interface {
		HasOrigin() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasOrigin()
	} else {
		return false
	}
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTCurveDescription1583) SetOrigin(v BTVector3d389) {
	type getResult interface {
		SetOrigin(v BTVector3d389)
	}

	o.GetActualInstance().(getResult).SetOrigin(v)
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTCurveDescription1583) GetType() GBTCurveTypeEnum {
	type getResult interface {
		GetType() GBTCurveTypeEnum
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetType()
	} else {
		var de GBTCurveTypeEnum
		return de
	}
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveDescription1583) GetTypeOk() (*GBTCurveTypeEnum, bool) {
	type getResult interface {
		GetTypeOk() (*GBTCurveTypeEnum, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTypeOk()
	} else {
		return nil, false
	}
}

// HasType returns a boolean if a field has been set.
func (o *BTCurveDescription1583) HasType() bool {
	type getResult interface {
		HasType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasType()
	} else {
		return false
	}
}

// SetType gets a reference to the given GBTCurveTypeEnum and assigns it to the Type field.
func (o *BTCurveDescription1583) SetType(v GBTCurveTypeEnum) {
	type getResult interface {
		SetType(v GBTCurveTypeEnum)
	}

	o.GetActualInstance().(getResult).SetType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTCurveDescription1583) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTCircleDescription-1145'
	if jsonDict["btType"] == "BTCircleDescription-1145" {
		// try to unmarshal JSON data into BTCircleDescription1145
		var qr *BTCircleDescription1145
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveDescription1583 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveDescription1583 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveDescription1583 as BTCircleDescription1145: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEllipseDescription-866'
	if jsonDict["btType"] == "BTEllipseDescription-866" {
		// try to unmarshal JSON data into BTEllipseDescription866
		var qr *BTEllipseDescription866
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveDescription1583 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveDescription1583 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveDescription1583 as BTEllipseDescription866: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTLineDescription-1559'
	if jsonDict["btType"] == "BTLineDescription-1559" {
		// try to unmarshal JSON data into BTLineDescription1559
		var qr *BTLineDescription1559
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveDescription1583 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveDescription1583 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveDescription1583 as BTLineDescription1559: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSplineDescription-2118'
	if jsonDict["btType"] == "BTSplineDescription-2118" {
		// try to unmarshal JSON data into BTSplineDescription2118
		var qr *BTSplineDescription2118
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTCurveDescription1583 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTCurveDescription1583 = nil
			return fmt.Errorf("Failed to unmarshal BTCurveDescription1583 as BTSplineDescription2118: %s", err.Error())
		}
	}

	var qtx *base_BTCurveDescription1583
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTCurveDescription1583 = qtx
		return nil // data stored in dst.base_BTCurveDescription1583, return on the first match
	} else {
		dst.implBTCurveDescription1583 = nil
		return fmt.Errorf("Failed to unmarshal BTCurveDescription1583 as base_BTCurveDescription1583: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTCurveDescription1583) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTCurveDescription1583) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTCurveDescription1583
}

type NullableBTCurveDescription1583 struct {
	value *BTCurveDescription1583
	isSet bool
}

func (v NullableBTCurveDescription1583) Get() *BTCurveDescription1583 {
	return v.value
}

func (v *NullableBTCurveDescription1583) Set(val *BTCurveDescription1583) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveDescription1583) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveDescription1583) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveDescription1583(val *BTCurveDescription1583) *NullableBTCurveDescription1583 {
	return &NullableBTCurveDescription1583{value: val, isSet: true}
}

func (v NullableBTCurveDescription1583) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveDescription1583) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTCurveDescription1583 struct {
	BtType                    *string           `json:"btType,omitempty"`
	Direction                 *BTVector3d389    `json:"direction,omitempty"`
	DirectionOrientedWithFace *BTVector3d389    `json:"directionOrientedWithFace,omitempty"`
	Origin                    *BTVector3d389    `json:"origin,omitempty"`
	Type                      *GBTCurveTypeEnum `json:"type,omitempty"`
}

// Newbase_BTCurveDescription1583 instantiates a new base_BTCurveDescription1583 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTCurveDescription1583() *base_BTCurveDescription1583 {
	this := base_BTCurveDescription1583{}
	return &this
}

// Newbase_BTCurveDescription1583WithDefaults instantiates a new base_BTCurveDescription1583 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTCurveDescription1583WithDefaults() *base_BTCurveDescription1583 {
	this := base_BTCurveDescription1583{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTCurveDescription1583) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveDescription1583) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTCurveDescription1583) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTCurveDescription1583) SetBtType(v string) {
	o.BtType = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *base_BTCurveDescription1583) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveDescription1583) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *base_BTCurveDescription1583) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *base_BTCurveDescription1583) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *base_BTCurveDescription1583) GetDirectionOrientedWithFace() BTVector3d389 {
	if o == nil || o.DirectionOrientedWithFace == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.DirectionOrientedWithFace
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveDescription1583) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	if o == nil || o.DirectionOrientedWithFace == nil {
		return nil, false
	}
	return o.DirectionOrientedWithFace, true
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *base_BTCurveDescription1583) HasDirectionOrientedWithFace() bool {
	if o != nil && o.DirectionOrientedWithFace != nil {
		return true
	}

	return false
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *base_BTCurveDescription1583) SetDirectionOrientedWithFace(v BTVector3d389) {
	o.DirectionOrientedWithFace = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *base_BTCurveDescription1583) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveDescription1583) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *base_BTCurveDescription1583) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *base_BTCurveDescription1583) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *base_BTCurveDescription1583) GetType() GBTCurveTypeEnum {
	if o == nil || o.Type == nil {
		var ret GBTCurveTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTCurveDescription1583) GetTypeOk() (*GBTCurveTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *base_BTCurveDescription1583) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GBTCurveTypeEnum and assigns it to the Type field.
func (o *base_BTCurveDescription1583) SetType(v GBTCurveTypeEnum) {
	o.Type = &v
}

func (o base_BTCurveDescription1583) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.DirectionOrientedWithFace != nil {
		toSerialize["directionOrientedWithFace"] = o.DirectionOrientedWithFace
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}
