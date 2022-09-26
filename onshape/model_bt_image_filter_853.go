/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6621-03f431879e4b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTImageFilter853 struct for BTImageFilter853
type BTImageFilter853 struct {
	BtType  *string `json:"btType,omitempty"`
	IsImage *bool   `json:"isImage,omitempty"`
}

// NewBTImageFilter853 instantiates a new BTImageFilter853 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTImageFilter853() *BTImageFilter853 {
	this := BTImageFilter853{}
	return &this
}

// NewBTImageFilter853WithDefaults instantiates a new BTImageFilter853 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTImageFilter853WithDefaults() *BTImageFilter853 {
	this := BTImageFilter853{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTImageFilter853) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTImageFilter853) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTImageFilter853) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTImageFilter853) SetBtType(v string) {
	o.BtType = &v
}

// GetIsImage returns the IsImage field value if set, zero value otherwise.
func (o *BTImageFilter853) GetIsImage() bool {
	if o == nil || o.IsImage == nil {
		var ret bool
		return ret
	}
	return *o.IsImage
}

// GetIsImageOk returns a tuple with the IsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTImageFilter853) GetIsImageOk() (*bool, bool) {
	if o == nil || o.IsImage == nil {
		return nil, false
	}
	return o.IsImage, true
}

// HasIsImage returns a boolean if a field has been set.
func (o *BTImageFilter853) HasIsImage() bool {
	if o != nil && o.IsImage != nil {
		return true
	}

	return false
}

// SetIsImage gets a reference to the given bool and assigns it to the IsImage field.
func (o *BTImageFilter853) SetIsImage(v bool) {
	o.IsImage = &v
}

func (o BTImageFilter853) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsImage != nil {
		toSerialize["isImage"] = o.IsImage
	}
	return json.Marshal(toSerialize)
}

type NullableBTImageFilter853 struct {
	value *BTImageFilter853
	isSet bool
}

func (v NullableBTImageFilter853) Get() *BTImageFilter853 {
	return v.value
}

func (v *NullableBTImageFilter853) Set(val *BTImageFilter853) {
	v.value = val
	v.isSet = true
}

func (v NullableBTImageFilter853) IsSet() bool {
	return v.isSet
}

func (v *NullableBTImageFilter853) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTImageFilter853(val *BTImageFilter853) *NullableBTImageFilter853 {
	return &NullableBTImageFilter853{value: val, isSet: true}
}

func (v NullableBTImageFilter853) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTImageFilter853) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
