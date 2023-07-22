/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19303-3cbf47a47fe4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTReleasePackageParams struct for BTReleasePackageParams
type BTReleasePackageParams struct {
	ChangeOrderId *string                      `json:"changeOrderId,omitempty"`
	Items         []BTReleasePackageItemParams `json:"items,omitempty"`
}

// NewBTReleasePackageParams instantiates a new BTReleasePackageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleasePackageParams() *BTReleasePackageParams {
	this := BTReleasePackageParams{}
	return &this
}

// NewBTReleasePackageParamsWithDefaults instantiates a new BTReleasePackageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleasePackageParamsWithDefaults() *BTReleasePackageParams {
	this := BTReleasePackageParams{}
	return &this
}

// GetChangeOrderId returns the ChangeOrderId field value if set, zero value otherwise.
func (o *BTReleasePackageParams) GetChangeOrderId() string {
	if o == nil || o.ChangeOrderId == nil {
		var ret string
		return ret
	}
	return *o.ChangeOrderId
}

// GetChangeOrderIdOk returns a tuple with the ChangeOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageParams) GetChangeOrderIdOk() (*string, bool) {
	if o == nil || o.ChangeOrderId == nil {
		return nil, false
	}
	return o.ChangeOrderId, true
}

// HasChangeOrderId returns a boolean if a field has been set.
func (o *BTReleasePackageParams) HasChangeOrderId() bool {
	if o != nil && o.ChangeOrderId != nil {
		return true
	}

	return false
}

// SetChangeOrderId gets a reference to the given string and assigns it to the ChangeOrderId field.
func (o *BTReleasePackageParams) SetChangeOrderId(v string) {
	o.ChangeOrderId = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTReleasePackageParams) GetItems() []BTReleasePackageItemParams {
	if o == nil || o.Items == nil {
		var ret []BTReleasePackageItemParams
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleasePackageParams) GetItemsOk() ([]BTReleasePackageItemParams, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTReleasePackageParams) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTReleasePackageItemParams and assigns it to the Items field.
func (o *BTReleasePackageParams) SetItems(v []BTReleasePackageItemParams) {
	o.Items = v
}

func (o BTReleasePackageParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeOrderId != nil {
		toSerialize["changeOrderId"] = o.ChangeOrderId
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleasePackageParams struct {
	value *BTReleasePackageParams
	isSet bool
}

func (v NullableBTReleasePackageParams) Get() *BTReleasePackageParams {
	return v.value
}

func (v *NullableBTReleasePackageParams) Set(val *BTReleasePackageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleasePackageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleasePackageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleasePackageParams(val *BTReleasePackageParams) *NullableBTReleasePackageParams {
	return &NullableBTReleasePackageParams{value: val, isSet: true}
}

func (v NullableBTReleasePackageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleasePackageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
