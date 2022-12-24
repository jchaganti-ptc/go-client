/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9191-43c781405890
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTJEditDelete1992 Deletes the specified node.
type BTJEditDelete1992 struct {
	BtType *string      `json:"btType,omitempty"`
	Path   *BTJPath3073 `json:"path,omitempty"`
}

// NewBTJEditDelete1992 instantiates a new BTJEditDelete1992 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJEditDelete1992() *BTJEditDelete1992 {
	this := BTJEditDelete1992{}
	return &this
}

// NewBTJEditDelete1992WithDefaults instantiates a new BTJEditDelete1992 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJEditDelete1992WithDefaults() *BTJEditDelete1992 {
	this := BTJEditDelete1992{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJEditDelete1992) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditDelete1992) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJEditDelete1992) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJEditDelete1992) SetBtType(v string) {
	o.BtType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTJEditDelete1992) GetPath() BTJPath3073 {
	if o == nil || o.Path == nil {
		var ret BTJPath3073
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJEditDelete1992) GetPathOk() (*BTJPath3073, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTJEditDelete1992) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given BTJPath3073 and assigns it to the Path field.
func (o *BTJEditDelete1992) SetPath(v BTJPath3073) {
	o.Path = &v
}

func (o BTJEditDelete1992) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableBTJEditDelete1992 struct {
	value *BTJEditDelete1992
	isSet bool
}

func (v NullableBTJEditDelete1992) Get() *BTJEditDelete1992 {
	return v.value
}

func (v *NullableBTJEditDelete1992) Set(val *BTJEditDelete1992) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJEditDelete1992) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJEditDelete1992) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJEditDelete1992(val *BTJEditDelete1992) *NullableBTJEditDelete1992 {
	return &NullableBTJEditDelete1992{value: val, isSet: true}
}

func (v NullableBTJEditDelete1992) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJEditDelete1992) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
