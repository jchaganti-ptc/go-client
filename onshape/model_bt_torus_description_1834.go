/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.5998-d3227e94fd7e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTTorusDescription1834 struct for BTTorusDescription1834
type BTTorusDescription1834 struct {
	Direction                 *BTVector3d389 `json:"direction,omitempty"`
	DirectionOrientedWithFace *BTVector3d389 `json:"directionOrientedWithFace,omitempty"`
	Origin                    *BTVector3d389 `json:"origin,omitempty"`
	Type                      *string        `json:"type,omitempty"`
	Axis                      *BTVector3d389 `json:"axis,omitempty"`
	BtType                    *string        `json:"btType,omitempty"`
	MajorRadius               *float64       `json:"majorRadius,omitempty"`
	MinorRadius               *float64       `json:"minorRadius,omitempty"`
}

// NewBTTorusDescription1834 instantiates a new BTTorusDescription1834 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTorusDescription1834() *BTTorusDescription1834 {
	this := BTTorusDescription1834{}
	return &this
}

// NewBTTorusDescription1834WithDefaults instantiates a new BTTorusDescription1834 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTorusDescription1834WithDefaults() *BTTorusDescription1834 {
	this := BTTorusDescription1834{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *BTTorusDescription1834) GetDirection() BTVector3d389 {
	if o == nil || o.Direction == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTorusDescription1834) GetDirectionOk() (*BTVector3d389, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *BTTorusDescription1834) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given BTVector3d389 and assigns it to the Direction field.
func (o *BTTorusDescription1834) SetDirection(v BTVector3d389) {
	o.Direction = &v
}

// GetDirectionOrientedWithFace returns the DirectionOrientedWithFace field value if set, zero value otherwise.
func (o *BTTorusDescription1834) GetDirectionOrientedWithFace() BTVector3d389 {
	if o == nil || o.DirectionOrientedWithFace == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.DirectionOrientedWithFace
}

// GetDirectionOrientedWithFaceOk returns a tuple with the DirectionOrientedWithFace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTorusDescription1834) GetDirectionOrientedWithFaceOk() (*BTVector3d389, bool) {
	if o == nil || o.DirectionOrientedWithFace == nil {
		return nil, false
	}
	return o.DirectionOrientedWithFace, true
}

// HasDirectionOrientedWithFace returns a boolean if a field has been set.
func (o *BTTorusDescription1834) HasDirectionOrientedWithFace() bool {
	if o != nil && o.DirectionOrientedWithFace != nil {
		return true
	}

	return false
}

// SetDirectionOrientedWithFace gets a reference to the given BTVector3d389 and assigns it to the DirectionOrientedWithFace field.
func (o *BTTorusDescription1834) SetDirectionOrientedWithFace(v BTVector3d389) {
	o.DirectionOrientedWithFace = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTTorusDescription1834) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTorusDescription1834) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTTorusDescription1834) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTTorusDescription1834) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTTorusDescription1834) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTorusDescription1834) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTTorusDescription1834) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTTorusDescription1834) SetType(v string) {
	o.Type = &v
}

// GetAxis returns the Axis field value if set, zero value otherwise.
func (o *BTTorusDescription1834) GetAxis() BTVector3d389 {
	if o == nil || o.Axis == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Axis
}

// GetAxisOk returns a tuple with the Axis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTorusDescription1834) GetAxisOk() (*BTVector3d389, bool) {
	if o == nil || o.Axis == nil {
		return nil, false
	}
	return o.Axis, true
}

// HasAxis returns a boolean if a field has been set.
func (o *BTTorusDescription1834) HasAxis() bool {
	if o != nil && o.Axis != nil {
		return true
	}

	return false
}

// SetAxis gets a reference to the given BTVector3d389 and assigns it to the Axis field.
func (o *BTTorusDescription1834) SetAxis(v BTVector3d389) {
	o.Axis = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTorusDescription1834) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTorusDescription1834) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTorusDescription1834) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTorusDescription1834) SetBtType(v string) {
	o.BtType = &v
}

// GetMajorRadius returns the MajorRadius field value if set, zero value otherwise.
func (o *BTTorusDescription1834) GetMajorRadius() float64 {
	if o == nil || o.MajorRadius == nil {
		var ret float64
		return ret
	}
	return *o.MajorRadius
}

// GetMajorRadiusOk returns a tuple with the MajorRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTorusDescription1834) GetMajorRadiusOk() (*float64, bool) {
	if o == nil || o.MajorRadius == nil {
		return nil, false
	}
	return o.MajorRadius, true
}

// HasMajorRadius returns a boolean if a field has been set.
func (o *BTTorusDescription1834) HasMajorRadius() bool {
	if o != nil && o.MajorRadius != nil {
		return true
	}

	return false
}

// SetMajorRadius gets a reference to the given float64 and assigns it to the MajorRadius field.
func (o *BTTorusDescription1834) SetMajorRadius(v float64) {
	o.MajorRadius = &v
}

// GetMinorRadius returns the MinorRadius field value if set, zero value otherwise.
func (o *BTTorusDescription1834) GetMinorRadius() float64 {
	if o == nil || o.MinorRadius == nil {
		var ret float64
		return ret
	}
	return *o.MinorRadius
}

// GetMinorRadiusOk returns a tuple with the MinorRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTorusDescription1834) GetMinorRadiusOk() (*float64, bool) {
	if o == nil || o.MinorRadius == nil {
		return nil, false
	}
	return o.MinorRadius, true
}

// HasMinorRadius returns a boolean if a field has been set.
func (o *BTTorusDescription1834) HasMinorRadius() bool {
	if o != nil && o.MinorRadius != nil {
		return true
	}

	return false
}

// SetMinorRadius gets a reference to the given float64 and assigns it to the MinorRadius field.
func (o *BTTorusDescription1834) SetMinorRadius(v float64) {
	o.MinorRadius = &v
}

func (o BTTorusDescription1834) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Axis != nil {
		toSerialize["axis"] = o.Axis
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.MajorRadius != nil {
		toSerialize["majorRadius"] = o.MajorRadius
	}
	if o.MinorRadius != nil {
		toSerialize["minorRadius"] = o.MinorRadius
	}
	return json.Marshal(toSerialize)
}

type NullableBTTorusDescription1834 struct {
	value *BTTorusDescription1834
	isSet bool
}

func (v NullableBTTorusDescription1834) Get() *BTTorusDescription1834 {
	return v.value
}

func (v *NullableBTTorusDescription1834) Set(val *BTTorusDescription1834) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTorusDescription1834) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTorusDescription1834) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTorusDescription1834(val *BTTorusDescription1834) *NullableBTTorusDescription1834 {
	return &NullableBTTorusDescription1834{value: val, isSet: true}
}

func (v NullableBTTorusDescription1834) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTorusDescription1834) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
