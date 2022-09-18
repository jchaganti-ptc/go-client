/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6510-4a0bae47e25c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVector3d389 struct for BTVector3d389
type BTVector3d389 struct {
	BtType *string  `json:"btType,omitempty"`
	X      *float64 `json:"x,omitempty"`
	Y      *float64 `json:"y,omitempty"`
	Z      *float64 `json:"z,omitempty"`
}

// NewBTVector3d389 instantiates a new BTVector3d389 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVector3d389() *BTVector3d389 {
	this := BTVector3d389{}
	return &this
}

// NewBTVector3d389WithDefaults instantiates a new BTVector3d389 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVector3d389WithDefaults() *BTVector3d389 {
	this := BTVector3d389{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTVector3d389) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVector3d389) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTVector3d389) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTVector3d389) SetBtType(v string) {
	o.BtType = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *BTVector3d389) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVector3d389) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *BTVector3d389) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *BTVector3d389) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *BTVector3d389) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVector3d389) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *BTVector3d389) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *BTVector3d389) SetY(v float64) {
	o.Y = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *BTVector3d389) GetZ() float64 {
	if o == nil || o.Z == nil {
		var ret float64
		return ret
	}
	return *o.Z
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVector3d389) GetZOk() (*float64, bool) {
	if o == nil || o.Z == nil {
		return nil, false
	}
	return o.Z, true
}

// HasZ returns a boolean if a field has been set.
func (o *BTVector3d389) HasZ() bool {
	if o != nil && o.Z != nil {
		return true
	}

	return false
}

// SetZ gets a reference to the given float64 and assigns it to the Z field.
func (o *BTVector3d389) SetZ(v float64) {
	o.Z = &v
}

func (o BTVector3d389) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	if o.Z != nil {
		toSerialize["z"] = o.Z
	}
	return json.Marshal(toSerialize)
}

type NullableBTVector3d389 struct {
	value *BTVector3d389
	isSet bool
}

func (v NullableBTVector3d389) Get() *BTVector3d389 {
	return v.value
}

func (v *NullableBTVector3d389) Set(val *BTVector3d389) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVector3d389) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVector3d389) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVector3d389(val *BTVector3d389) *NullableBTVector3d389 {
	return &NullableBTVector3d389{value: val, isSet: true}
}

func (v NullableBTVector3d389) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVector3d389) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
