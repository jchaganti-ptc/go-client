/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6228-2857ab16a033
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementStartTransactionParams struct for BTAppElementStartTransactionParams
type BTAppElementStartTransactionParams struct {
	Description    *string `json:"description,omitempty"`
	ParentChangeId *string `json:"parentChangeId,omitempty"`
	ReturnError    *bool   `json:"returnError,omitempty"`
}

// NewBTAppElementStartTransactionParams instantiates a new BTAppElementStartTransactionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementStartTransactionParams() *BTAppElementStartTransactionParams {
	this := BTAppElementStartTransactionParams{}
	return &this
}

// NewBTAppElementStartTransactionParamsWithDefaults instantiates a new BTAppElementStartTransactionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementStartTransactionParamsWithDefaults() *BTAppElementStartTransactionParams {
	this := BTAppElementStartTransactionParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAppElementStartTransactionParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementStartTransactionParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAppElementStartTransactionParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAppElementStartTransactionParams) SetDescription(v string) {
	o.Description = &v
}

// GetParentChangeId returns the ParentChangeId field value if set, zero value otherwise.
func (o *BTAppElementStartTransactionParams) GetParentChangeId() string {
	if o == nil || o.ParentChangeId == nil {
		var ret string
		return ret
	}
	return *o.ParentChangeId
}

// GetParentChangeIdOk returns a tuple with the ParentChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementStartTransactionParams) GetParentChangeIdOk() (*string, bool) {
	if o == nil || o.ParentChangeId == nil {
		return nil, false
	}
	return o.ParentChangeId, true
}

// HasParentChangeId returns a boolean if a field has been set.
func (o *BTAppElementStartTransactionParams) HasParentChangeId() bool {
	if o != nil && o.ParentChangeId != nil {
		return true
	}

	return false
}

// SetParentChangeId gets a reference to the given string and assigns it to the ParentChangeId field.
func (o *BTAppElementStartTransactionParams) SetParentChangeId(v string) {
	o.ParentChangeId = &v
}

// GetReturnError returns the ReturnError field value if set, zero value otherwise.
func (o *BTAppElementStartTransactionParams) GetReturnError() bool {
	if o == nil || o.ReturnError == nil {
		var ret bool
		return ret
	}
	return *o.ReturnError
}

// GetReturnErrorOk returns a tuple with the ReturnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementStartTransactionParams) GetReturnErrorOk() (*bool, bool) {
	if o == nil || o.ReturnError == nil {
		return nil, false
	}
	return o.ReturnError, true
}

// HasReturnError returns a boolean if a field has been set.
func (o *BTAppElementStartTransactionParams) HasReturnError() bool {
	if o != nil && o.ReturnError != nil {
		return true
	}

	return false
}

// SetReturnError gets a reference to the given bool and assigns it to the ReturnError field.
func (o *BTAppElementStartTransactionParams) SetReturnError(v bool) {
	o.ReturnError = &v
}

func (o BTAppElementStartTransactionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ParentChangeId != nil {
		toSerialize["parentChangeId"] = o.ParentChangeId
	}
	if o.ReturnError != nil {
		toSerialize["returnError"] = o.ReturnError
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementStartTransactionParams struct {
	value *BTAppElementStartTransactionParams
	isSet bool
}

func (v NullableBTAppElementStartTransactionParams) Get() *BTAppElementStartTransactionParams {
	return v.value
}

func (v *NullableBTAppElementStartTransactionParams) Set(val *BTAppElementStartTransactionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementStartTransactionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementStartTransactionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementStartTransactionParams(val *BTAppElementStartTransactionParams) *NullableBTAppElementStartTransactionParams {
	return &NullableBTAppElementStartTransactionParams{value: val, isSet: true}
}

func (v NullableBTAppElementStartTransactionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementStartTransactionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
