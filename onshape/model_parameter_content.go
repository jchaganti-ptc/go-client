/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.26043-b28d7068bd76
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// ParameterContent struct for ParameterContent
type ParameterContent struct {
	Empty *bool `json:"empty,omitempty"`
}

// NewParameterContent instantiates a new ParameterContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterContent() *ParameterContent {
	this := ParameterContent{}
	return &this
}

// NewParameterContentWithDefaults instantiates a new ParameterContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterContentWithDefaults() *ParameterContent {
	this := ParameterContent{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *ParameterContent) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContent) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *ParameterContent) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *ParameterContent) SetEmpty(v bool) {
	o.Empty = &v
}

func (o ParameterContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	return json.Marshal(toSerialize)
}

type NullableParameterContent struct {
	value *ParameterContent
	isSet bool
}

func (v NullableParameterContent) Get() *ParameterContent {
	return v.value
}

func (v *NullableParameterContent) Set(val *ParameterContent) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterContent) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterContent(val *ParameterContent) *NullableParameterContent {
	return &NullableParameterContent{value: val, isSet: true}
}

func (v NullableParameterContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}