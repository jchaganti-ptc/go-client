/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.21206-13add30fbba2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// ShippingDetails struct for ShippingDetails
type ShippingDetails struct {
	Address *Address `json:"address,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Phone   *string  `json:"phone,omitempty"`
}

// NewShippingDetails instantiates a new ShippingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingDetails() *ShippingDetails {
	this := ShippingDetails{}
	return &this
}

// NewShippingDetailsWithDefaults instantiates a new ShippingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingDetailsWithDefaults() *ShippingDetails {
	this := ShippingDetails{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ShippingDetails) GetAddress() Address {
	if o == nil || o.Address == nil {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingDetails) GetAddressOk() (*Address, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ShippingDetails) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ShippingDetails) SetAddress(v Address) {
	o.Address = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShippingDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShippingDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShippingDetails) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ShippingDetails) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingDetails) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ShippingDetails) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ShippingDetails) SetPhone(v string) {
	o.Phone = &v
}

func (o ShippingDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableShippingDetails struct {
	value *ShippingDetails
	isSet bool
}

func (v NullableShippingDetails) Get() *ShippingDetails {
	return v.value
}

func (v *NullableShippingDetails) Set(val *ShippingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingDetails(val *ShippingDetails) *NullableShippingDetails {
	return &NullableShippingDetails{value: val, isSet: true}
}

func (v NullableShippingDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
