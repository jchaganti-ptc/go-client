/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6415-48a6b2252b8c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAllowFlattenedGeometryFilter2140 struct for BTAllowFlattenedGeometryFilter2140
type BTAllowFlattenedGeometryFilter2140 struct {
	BtType                  *string `json:"btType,omitempty"`
	AllowsFlattenedGeometry *bool   `json:"allowsFlattenedGeometry,omitempty"`
}

// NewBTAllowFlattenedGeometryFilter2140 instantiates a new BTAllowFlattenedGeometryFilter2140 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAllowFlattenedGeometryFilter2140() *BTAllowFlattenedGeometryFilter2140 {
	this := BTAllowFlattenedGeometryFilter2140{}
	return &this
}

// NewBTAllowFlattenedGeometryFilter2140WithDefaults instantiates a new BTAllowFlattenedGeometryFilter2140 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAllowFlattenedGeometryFilter2140WithDefaults() *BTAllowFlattenedGeometryFilter2140 {
	this := BTAllowFlattenedGeometryFilter2140{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAllowFlattenedGeometryFilter2140) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowFlattenedGeometryFilter2140) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAllowFlattenedGeometryFilter2140) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAllowFlattenedGeometryFilter2140) SetBtType(v string) {
	o.BtType = &v
}

// GetAllowsFlattenedGeometry returns the AllowsFlattenedGeometry field value if set, zero value otherwise.
func (o *BTAllowFlattenedGeometryFilter2140) GetAllowsFlattenedGeometry() bool {
	if o == nil || o.AllowsFlattenedGeometry == nil {
		var ret bool
		return ret
	}
	return *o.AllowsFlattenedGeometry
}

// GetAllowsFlattenedGeometryOk returns a tuple with the AllowsFlattenedGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowFlattenedGeometryFilter2140) GetAllowsFlattenedGeometryOk() (*bool, bool) {
	if o == nil || o.AllowsFlattenedGeometry == nil {
		return nil, false
	}
	return o.AllowsFlattenedGeometry, true
}

// HasAllowsFlattenedGeometry returns a boolean if a field has been set.
func (o *BTAllowFlattenedGeometryFilter2140) HasAllowsFlattenedGeometry() bool {
	if o != nil && o.AllowsFlattenedGeometry != nil {
		return true
	}

	return false
}

// SetAllowsFlattenedGeometry gets a reference to the given bool and assigns it to the AllowsFlattenedGeometry field.
func (o *BTAllowFlattenedGeometryFilter2140) SetAllowsFlattenedGeometry(v bool) {
	o.AllowsFlattenedGeometry = &v
}

func (o BTAllowFlattenedGeometryFilter2140) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.AllowsFlattenedGeometry != nil {
		toSerialize["allowsFlattenedGeometry"] = o.AllowsFlattenedGeometry
	}
	return json.Marshal(toSerialize)
}

type NullableBTAllowFlattenedGeometryFilter2140 struct {
	value *BTAllowFlattenedGeometryFilter2140
	isSet bool
}

func (v NullableBTAllowFlattenedGeometryFilter2140) Get() *BTAllowFlattenedGeometryFilter2140 {
	return v.value
}

func (v *NullableBTAllowFlattenedGeometryFilter2140) Set(val *BTAllowFlattenedGeometryFilter2140) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAllowFlattenedGeometryFilter2140) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAllowFlattenedGeometryFilter2140) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAllowFlattenedGeometryFilter2140(val *BTAllowFlattenedGeometryFilter2140) *NullableBTAllowFlattenedGeometryFilter2140 {
	return &NullableBTAllowFlattenedGeometryFilter2140{value: val, isSet: true}
}

func (v NullableBTAllowFlattenedGeometryFilter2140) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAllowFlattenedGeometryFilter2140) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
