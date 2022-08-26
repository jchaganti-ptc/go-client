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

// BTAppElementParamsArrayBTCopyViewAssociativeDataParams struct for BTAppElementParamsArrayBTCopyViewAssociativeDataParams
type BTAppElementParamsArrayBTCopyViewAssociativeDataParams struct {
	Description    *string                           `json:"description,omitempty"`
	Items          []BTCopyViewAssociativeDataParams `json:"items,omitempty"`
	ParentChangeId *string                           `json:"parentChangeId,omitempty"`
	TransactionId  *string                           `json:"transactionId,omitempty"`
}

// NewBTAppElementParamsArrayBTCopyViewAssociativeDataParams instantiates a new BTAppElementParamsArrayBTCopyViewAssociativeDataParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementParamsArrayBTCopyViewAssociativeDataParams() *BTAppElementParamsArrayBTCopyViewAssociativeDataParams {
	this := BTAppElementParamsArrayBTCopyViewAssociativeDataParams{}
	return &this
}

// NewBTAppElementParamsArrayBTCopyViewAssociativeDataParamsWithDefaults instantiates a new BTAppElementParamsArrayBTCopyViewAssociativeDataParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementParamsArrayBTCopyViewAssociativeDataParamsWithDefaults() *BTAppElementParamsArrayBTCopyViewAssociativeDataParams {
	this := BTAppElementParamsArrayBTCopyViewAssociativeDataParams{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) SetDescription(v string) {
	o.Description = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) GetItems() []BTCopyViewAssociativeDataParams {
	if o == nil || o.Items == nil {
		var ret []BTCopyViewAssociativeDataParams
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) GetItemsOk() ([]BTCopyViewAssociativeDataParams, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTCopyViewAssociativeDataParams and assigns it to the Items field.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) SetItems(v []BTCopyViewAssociativeDataParams) {
	o.Items = v
}

// GetParentChangeId returns the ParentChangeId field value if set, zero value otherwise.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) GetParentChangeId() string {
	if o == nil || o.ParentChangeId == nil {
		var ret string
		return ret
	}
	return *o.ParentChangeId
}

// GetParentChangeIdOk returns a tuple with the ParentChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) GetParentChangeIdOk() (*string, bool) {
	if o == nil || o.ParentChangeId == nil {
		return nil, false
	}
	return o.ParentChangeId, true
}

// HasParentChangeId returns a boolean if a field has been set.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) HasParentChangeId() bool {
	if o != nil && o.ParentChangeId != nil {
		return true
	}

	return false
}

// SetParentChangeId gets a reference to the given string and assigns it to the ParentChangeId field.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) SetParentChangeId(v string) {
	o.ParentChangeId = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o BTAppElementParamsArrayBTCopyViewAssociativeDataParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ParentChangeId != nil {
		toSerialize["parentChangeId"] = o.ParentChangeId
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams struct {
	value *BTAppElementParamsArrayBTCopyViewAssociativeDataParams
	isSet bool
}

func (v NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams) Get() *BTAppElementParamsArrayBTCopyViewAssociativeDataParams {
	return v.value
}

func (v *NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams) Set(val *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams(val *BTAppElementParamsArrayBTCopyViewAssociativeDataParams) *NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams {
	return &NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams{value: val, isSet: true}
}

func (v NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementParamsArrayBTCopyViewAssociativeDataParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
