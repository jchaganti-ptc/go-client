/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7232-a44b68534e12
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCurveGeometryConic2284 struct for BTCurveGeometryConic2284
type BTCurveGeometryConic2284 struct {
	BtType *string   `json:"btType,omitempty"`
	Points []float64 `json:"points,omitempty"`
	Rho    *float64  `json:"rho,omitempty"`
}

// NewBTCurveGeometryConic2284 instantiates a new BTCurveGeometryConic2284 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCurveGeometryConic2284() *BTCurveGeometryConic2284 {
	this := BTCurveGeometryConic2284{}
	return &this
}

// NewBTCurveGeometryConic2284WithDefaults instantiates a new BTCurveGeometryConic2284 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCurveGeometryConic2284WithDefaults() *BTCurveGeometryConic2284 {
	this := BTCurveGeometryConic2284{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCurveGeometryConic2284) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryConic2284) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCurveGeometryConic2284) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCurveGeometryConic2284) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTCurveGeometryConic2284) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryConic2284) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTCurveGeometryConic2284) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTCurveGeometryConic2284) SetPoints(v []float64) {
	o.Points = v
}

// GetRho returns the Rho field value if set, zero value otherwise.
func (o *BTCurveGeometryConic2284) GetRho() float64 {
	if o == nil || o.Rho == nil {
		var ret float64
		return ret
	}
	return *o.Rho
}

// GetRhoOk returns a tuple with the Rho field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCurveGeometryConic2284) GetRhoOk() (*float64, bool) {
	if o == nil || o.Rho == nil {
		return nil, false
	}
	return o.Rho, true
}

// HasRho returns a boolean if a field has been set.
func (o *BTCurveGeometryConic2284) HasRho() bool {
	if o != nil && o.Rho != nil {
		return true
	}

	return false
}

// SetRho gets a reference to the given float64 and assigns it to the Rho field.
func (o *BTCurveGeometryConic2284) SetRho(v float64) {
	o.Rho = &v
}

func (o BTCurveGeometryConic2284) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.Rho != nil {
		toSerialize["rho"] = o.Rho
	}
	return json.Marshal(toSerialize)
}

type NullableBTCurveGeometryConic2284 struct {
	value *BTCurveGeometryConic2284
	isSet bool
}

func (v NullableBTCurveGeometryConic2284) Get() *BTCurveGeometryConic2284 {
	return v.value
}

func (v *NullableBTCurveGeometryConic2284) Set(val *BTCurveGeometryConic2284) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCurveGeometryConic2284) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCurveGeometryConic2284) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCurveGeometryConic2284(val *BTCurveGeometryConic2284) *NullableBTCurveGeometryConic2284 {
	return &NullableBTCurveGeometryConic2284{value: val, isSet: true}
}

func (v NullableBTCurveGeometryConic2284) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCurveGeometryConic2284) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
