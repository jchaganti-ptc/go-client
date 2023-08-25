/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.21206-13add30fbba2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCoordinateSystem387 struct for BTCoordinateSystem387
type BTCoordinateSystem387 struct {
	BtType *string        `json:"btType,omitempty"`
	Matrix *BTBSMatrix386 `json:"matrix,omitempty"`
	Origin *BTVector3d389 `json:"origin,omitempty"`
	Xaxis  *BTVector3d389 `json:"xaxis,omitempty"`
	Yaxis  *BTVector3d389 `json:"yaxis,omitempty"`
	Zaxis  *BTVector3d389 `json:"zaxis,omitempty"`
}

// NewBTCoordinateSystem387 instantiates a new BTCoordinateSystem387 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCoordinateSystem387() *BTCoordinateSystem387 {
	this := BTCoordinateSystem387{}
	return &this
}

// NewBTCoordinateSystem387WithDefaults instantiates a new BTCoordinateSystem387 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCoordinateSystem387WithDefaults() *BTCoordinateSystem387 {
	this := BTCoordinateSystem387{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCoordinateSystem387) SetBtType(v string) {
	o.BtType = &v
}

// GetMatrix returns the Matrix field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetMatrix() BTBSMatrix386 {
	if o == nil || o.Matrix == nil {
		var ret BTBSMatrix386
		return ret
	}
	return *o.Matrix
}

// GetMatrixOk returns a tuple with the Matrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetMatrixOk() (*BTBSMatrix386, bool) {
	if o == nil || o.Matrix == nil {
		return nil, false
	}
	return o.Matrix, true
}

// HasMatrix returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasMatrix() bool {
	if o != nil && o.Matrix != nil {
		return true
	}

	return false
}

// SetMatrix gets a reference to the given BTBSMatrix386 and assigns it to the Matrix field.
func (o *BTCoordinateSystem387) SetMatrix(v BTBSMatrix386) {
	o.Matrix = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTCoordinateSystem387) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetXaxis returns the Xaxis field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetXaxis() BTVector3d389 {
	if o == nil || o.Xaxis == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Xaxis
}

// GetXaxisOk returns a tuple with the Xaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetXaxisOk() (*BTVector3d389, bool) {
	if o == nil || o.Xaxis == nil {
		return nil, false
	}
	return o.Xaxis, true
}

// HasXaxis returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasXaxis() bool {
	if o != nil && o.Xaxis != nil {
		return true
	}

	return false
}

// SetXaxis gets a reference to the given BTVector3d389 and assigns it to the Xaxis field.
func (o *BTCoordinateSystem387) SetXaxis(v BTVector3d389) {
	o.Xaxis = &v
}

// GetYaxis returns the Yaxis field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetYaxis() BTVector3d389 {
	if o == nil || o.Yaxis == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Yaxis
}

// GetYaxisOk returns a tuple with the Yaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetYaxisOk() (*BTVector3d389, bool) {
	if o == nil || o.Yaxis == nil {
		return nil, false
	}
	return o.Yaxis, true
}

// HasYaxis returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasYaxis() bool {
	if o != nil && o.Yaxis != nil {
		return true
	}

	return false
}

// SetYaxis gets a reference to the given BTVector3d389 and assigns it to the Yaxis field.
func (o *BTCoordinateSystem387) SetYaxis(v BTVector3d389) {
	o.Yaxis = &v
}

// GetZaxis returns the Zaxis field value if set, zero value otherwise.
func (o *BTCoordinateSystem387) GetZaxis() BTVector3d389 {
	if o == nil || o.Zaxis == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Zaxis
}

// GetZaxisOk returns a tuple with the Zaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCoordinateSystem387) GetZaxisOk() (*BTVector3d389, bool) {
	if o == nil || o.Zaxis == nil {
		return nil, false
	}
	return o.Zaxis, true
}

// HasZaxis returns a boolean if a field has been set.
func (o *BTCoordinateSystem387) HasZaxis() bool {
	if o != nil && o.Zaxis != nil {
		return true
	}

	return false
}

// SetZaxis gets a reference to the given BTVector3d389 and assigns it to the Zaxis field.
func (o *BTCoordinateSystem387) SetZaxis(v BTVector3d389) {
	o.Zaxis = &v
}

func (o BTCoordinateSystem387) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Matrix != nil {
		toSerialize["matrix"] = o.Matrix
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Xaxis != nil {
		toSerialize["xaxis"] = o.Xaxis
	}
	if o.Yaxis != nil {
		toSerialize["yaxis"] = o.Yaxis
	}
	if o.Zaxis != nil {
		toSerialize["zaxis"] = o.Zaxis
	}
	return json.Marshal(toSerialize)
}

type NullableBTCoordinateSystem387 struct {
	value *BTCoordinateSystem387
	isSet bool
}

func (v NullableBTCoordinateSystem387) Get() *BTCoordinateSystem387 {
	return v.value
}

func (v *NullableBTCoordinateSystem387) Set(val *BTCoordinateSystem387) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCoordinateSystem387) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCoordinateSystem387) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCoordinateSystem387(val *BTCoordinateSystem387) *NullableBTCoordinateSystem387 {
	return &NullableBTCoordinateSystem387{value: val, isSet: true}
}

func (v NullableBTCoordinateSystem387) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCoordinateSystem387) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
