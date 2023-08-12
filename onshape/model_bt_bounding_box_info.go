/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.20454-7718daa9749d
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBoundingBoxInfo struct for BTBoundingBoxInfo
type BTBoundingBoxInfo struct {
	HighX *float64 `json:"highX,omitempty"`
	HighY *float64 `json:"highY,omitempty"`
	HighZ *float64 `json:"highZ,omitempty"`
	LowX  *float64 `json:"lowX,omitempty"`
	LowY  *float64 `json:"lowY,omitempty"`
	LowZ  *float64 `json:"lowZ,omitempty"`
}

// NewBTBoundingBoxInfo instantiates a new BTBoundingBoxInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBoundingBoxInfo() *BTBoundingBoxInfo {
	this := BTBoundingBoxInfo{}
	return &this
}

// NewBTBoundingBoxInfoWithDefaults instantiates a new BTBoundingBoxInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBoundingBoxInfoWithDefaults() *BTBoundingBoxInfo {
	this := BTBoundingBoxInfo{}
	return &this
}

// GetHighX returns the HighX field value if set, zero value otherwise.
func (o *BTBoundingBoxInfo) GetHighX() float64 {
	if o == nil || o.HighX == nil {
		var ret float64
		return ret
	}
	return *o.HighX
}

// GetHighXOk returns a tuple with the HighX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBoxInfo) GetHighXOk() (*float64, bool) {
	if o == nil || o.HighX == nil {
		return nil, false
	}
	return o.HighX, true
}

// HasHighX returns a boolean if a field has been set.
func (o *BTBoundingBoxInfo) HasHighX() bool {
	if o != nil && o.HighX != nil {
		return true
	}

	return false
}

// SetHighX gets a reference to the given float64 and assigns it to the HighX field.
func (o *BTBoundingBoxInfo) SetHighX(v float64) {
	o.HighX = &v
}

// GetHighY returns the HighY field value if set, zero value otherwise.
func (o *BTBoundingBoxInfo) GetHighY() float64 {
	if o == nil || o.HighY == nil {
		var ret float64
		return ret
	}
	return *o.HighY
}

// GetHighYOk returns a tuple with the HighY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBoxInfo) GetHighYOk() (*float64, bool) {
	if o == nil || o.HighY == nil {
		return nil, false
	}
	return o.HighY, true
}

// HasHighY returns a boolean if a field has been set.
func (o *BTBoundingBoxInfo) HasHighY() bool {
	if o != nil && o.HighY != nil {
		return true
	}

	return false
}

// SetHighY gets a reference to the given float64 and assigns it to the HighY field.
func (o *BTBoundingBoxInfo) SetHighY(v float64) {
	o.HighY = &v
}

// GetHighZ returns the HighZ field value if set, zero value otherwise.
func (o *BTBoundingBoxInfo) GetHighZ() float64 {
	if o == nil || o.HighZ == nil {
		var ret float64
		return ret
	}
	return *o.HighZ
}

// GetHighZOk returns a tuple with the HighZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBoxInfo) GetHighZOk() (*float64, bool) {
	if o == nil || o.HighZ == nil {
		return nil, false
	}
	return o.HighZ, true
}

// HasHighZ returns a boolean if a field has been set.
func (o *BTBoundingBoxInfo) HasHighZ() bool {
	if o != nil && o.HighZ != nil {
		return true
	}

	return false
}

// SetHighZ gets a reference to the given float64 and assigns it to the HighZ field.
func (o *BTBoundingBoxInfo) SetHighZ(v float64) {
	o.HighZ = &v
}

// GetLowX returns the LowX field value if set, zero value otherwise.
func (o *BTBoundingBoxInfo) GetLowX() float64 {
	if o == nil || o.LowX == nil {
		var ret float64
		return ret
	}
	return *o.LowX
}

// GetLowXOk returns a tuple with the LowX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBoxInfo) GetLowXOk() (*float64, bool) {
	if o == nil || o.LowX == nil {
		return nil, false
	}
	return o.LowX, true
}

// HasLowX returns a boolean if a field has been set.
func (o *BTBoundingBoxInfo) HasLowX() bool {
	if o != nil && o.LowX != nil {
		return true
	}

	return false
}

// SetLowX gets a reference to the given float64 and assigns it to the LowX field.
func (o *BTBoundingBoxInfo) SetLowX(v float64) {
	o.LowX = &v
}

// GetLowY returns the LowY field value if set, zero value otherwise.
func (o *BTBoundingBoxInfo) GetLowY() float64 {
	if o == nil || o.LowY == nil {
		var ret float64
		return ret
	}
	return *o.LowY
}

// GetLowYOk returns a tuple with the LowY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBoxInfo) GetLowYOk() (*float64, bool) {
	if o == nil || o.LowY == nil {
		return nil, false
	}
	return o.LowY, true
}

// HasLowY returns a boolean if a field has been set.
func (o *BTBoundingBoxInfo) HasLowY() bool {
	if o != nil && o.LowY != nil {
		return true
	}

	return false
}

// SetLowY gets a reference to the given float64 and assigns it to the LowY field.
func (o *BTBoundingBoxInfo) SetLowY(v float64) {
	o.LowY = &v
}

// GetLowZ returns the LowZ field value if set, zero value otherwise.
func (o *BTBoundingBoxInfo) GetLowZ() float64 {
	if o == nil || o.LowZ == nil {
		var ret float64
		return ret
	}
	return *o.LowZ
}

// GetLowZOk returns a tuple with the LowZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBoxInfo) GetLowZOk() (*float64, bool) {
	if o == nil || o.LowZ == nil {
		return nil, false
	}
	return o.LowZ, true
}

// HasLowZ returns a boolean if a field has been set.
func (o *BTBoundingBoxInfo) HasLowZ() bool {
	if o != nil && o.LowZ != nil {
		return true
	}

	return false
}

// SetLowZ gets a reference to the given float64 and assigns it to the LowZ field.
func (o *BTBoundingBoxInfo) SetLowZ(v float64) {
	o.LowZ = &v
}

func (o BTBoundingBoxInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HighX != nil {
		toSerialize["highX"] = o.HighX
	}
	if o.HighY != nil {
		toSerialize["highY"] = o.HighY
	}
	if o.HighZ != nil {
		toSerialize["highZ"] = o.HighZ
	}
	if o.LowX != nil {
		toSerialize["lowX"] = o.LowX
	}
	if o.LowY != nil {
		toSerialize["lowY"] = o.LowY
	}
	if o.LowZ != nil {
		toSerialize["lowZ"] = o.LowZ
	}
	return json.Marshal(toSerialize)
}

type NullableBTBoundingBoxInfo struct {
	value *BTBoundingBoxInfo
	isSet bool
}

func (v NullableBTBoundingBoxInfo) Get() *BTBoundingBoxInfo {
	return v.value
}

func (v *NullableBTBoundingBoxInfo) Set(val *BTBoundingBoxInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBoundingBoxInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBoundingBoxInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBoundingBoxInfo(val *BTBoundingBoxInfo) *NullableBTBoundingBoxInfo {
	return &NullableBTBoundingBoxInfo{value: val, isSet: true}
}

func (v NullableBTBoundingBoxInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBoundingBoxInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
