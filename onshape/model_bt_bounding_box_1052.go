/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5928-bd774e9c9f97
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBoundingBox1052 struct for BTBoundingBox1052
type BTBoundingBox1052 struct {
	MaxCorner *BTVector3d389 `json:"maxCorner,omitempty"`
	MinCorner *BTVector3d389 `json:"minCorner,omitempty"`
	Valid     *bool          `json:"valid,omitempty"`
}

// NewBTBoundingBox1052 instantiates a new BTBoundingBox1052 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBoundingBox1052() *BTBoundingBox1052 {
	this := BTBoundingBox1052{}
	return &this
}

// NewBTBoundingBox1052WithDefaults instantiates a new BTBoundingBox1052 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBoundingBox1052WithDefaults() *BTBoundingBox1052 {
	this := BTBoundingBox1052{}
	return &this
}

// GetMaxCorner returns the MaxCorner field value if set, zero value otherwise.
func (o *BTBoundingBox1052) GetMaxCorner() BTVector3d389 {
	if o == nil || o.MaxCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MaxCorner
}

// GetMaxCornerOk returns a tuple with the MaxCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBox1052) GetMaxCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.MaxCorner == nil {
		return nil, false
	}
	return o.MaxCorner, true
}

// HasMaxCorner returns a boolean if a field has been set.
func (o *BTBoundingBox1052) HasMaxCorner() bool {
	if o != nil && o.MaxCorner != nil {
		return true
	}

	return false
}

// SetMaxCorner gets a reference to the given BTVector3d389 and assigns it to the MaxCorner field.
func (o *BTBoundingBox1052) SetMaxCorner(v BTVector3d389) {
	o.MaxCorner = &v
}

// GetMinCorner returns the MinCorner field value if set, zero value otherwise.
func (o *BTBoundingBox1052) GetMinCorner() BTVector3d389 {
	if o == nil || o.MinCorner == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MinCorner
}

// GetMinCornerOk returns a tuple with the MinCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBox1052) GetMinCornerOk() (*BTVector3d389, bool) {
	if o == nil || o.MinCorner == nil {
		return nil, false
	}
	return o.MinCorner, true
}

// HasMinCorner returns a boolean if a field has been set.
func (o *BTBoundingBox1052) HasMinCorner() bool {
	if o != nil && o.MinCorner != nil {
		return true
	}

	return false
}

// SetMinCorner gets a reference to the given BTVector3d389 and assigns it to the MinCorner field.
func (o *BTBoundingBox1052) SetMinCorner(v BTVector3d389) {
	o.MinCorner = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BTBoundingBox1052) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBoundingBox1052) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *BTBoundingBox1052) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BTBoundingBox1052) SetValid(v bool) {
	o.Valid = &v
}

func (o BTBoundingBox1052) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxCorner != nil {
		toSerialize["maxCorner"] = o.MaxCorner
	}
	if o.MinCorner != nil {
		toSerialize["minCorner"] = o.MinCorner
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}

type NullableBTBoundingBox1052 struct {
	value *BTBoundingBox1052
	isSet bool
}

func (v NullableBTBoundingBox1052) Get() *BTBoundingBox1052 {
	return v.value
}

func (v *NullableBTBoundingBox1052) Set(val *BTBoundingBox1052) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBoundingBox1052) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBoundingBox1052) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBoundingBox1052(val *BTBoundingBox1052) *NullableBTBoundingBox1052 {
	return &NullableBTBoundingBox1052{value: val, isSet: true}
}

func (v NullableBTBoundingBox1052) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBoundingBox1052) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
