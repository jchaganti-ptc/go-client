/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23604-b59b123004e9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// CompanyRole struct for CompanyRole
type CompanyRole struct {
	Admin       *bool   `json:"admin,omitempty"`
	CompanyId   *string `json:"companyId,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	Guest       *bool   `json:"guest,omitempty"`
	Light       *bool   `json:"light,omitempty"`
}

// NewCompanyRole instantiates a new CompanyRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyRole() *CompanyRole {
	this := CompanyRole{}
	return &this
}

// NewCompanyRoleWithDefaults instantiates a new CompanyRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyRoleWithDefaults() *CompanyRole {
	this := CompanyRole{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *CompanyRole) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyRole) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *CompanyRole) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *CompanyRole) SetAdmin(v bool) {
	o.Admin = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *CompanyRole) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyRole) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *CompanyRole) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *CompanyRole) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *CompanyRole) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyRole) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *CompanyRole) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *CompanyRole) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetGuest returns the Guest field value if set, zero value otherwise.
func (o *CompanyRole) GetGuest() bool {
	if o == nil || o.Guest == nil {
		var ret bool
		return ret
	}
	return *o.Guest
}

// GetGuestOk returns a tuple with the Guest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyRole) GetGuestOk() (*bool, bool) {
	if o == nil || o.Guest == nil {
		return nil, false
	}
	return o.Guest, true
}

// HasGuest returns a boolean if a field has been set.
func (o *CompanyRole) HasGuest() bool {
	if o != nil && o.Guest != nil {
		return true
	}

	return false
}

// SetGuest gets a reference to the given bool and assigns it to the Guest field.
func (o *CompanyRole) SetGuest(v bool) {
	o.Guest = &v
}

// GetLight returns the Light field value if set, zero value otherwise.
func (o *CompanyRole) GetLight() bool {
	if o == nil || o.Light == nil {
		var ret bool
		return ret
	}
	return *o.Light
}

// GetLightOk returns a tuple with the Light field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyRole) GetLightOk() (*bool, bool) {
	if o == nil || o.Light == nil {
		return nil, false
	}
	return o.Light, true
}

// HasLight returns a boolean if a field has been set.
func (o *CompanyRole) HasLight() bool {
	if o != nil && o.Light != nil {
		return true
	}

	return false
}

// SetLight gets a reference to the given bool and assigns it to the Light field.
func (o *CompanyRole) SetLight(v bool) {
	o.Light = &v
}

func (o CompanyRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.CompanyName != nil {
		toSerialize["companyName"] = o.CompanyName
	}
	if o.Guest != nil {
		toSerialize["guest"] = o.Guest
	}
	if o.Light != nil {
		toSerialize["light"] = o.Light
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyRole struct {
	value *CompanyRole
	isSet bool
}

func (v NullableCompanyRole) Get() *CompanyRole {
	return v.value
}

func (v *NullableCompanyRole) Set(val *CompanyRole) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyRole) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyRole(val *CompanyRole) *NullableCompanyRole {
	return &NullableCompanyRole{value: val, isSet: true}
}

func (v NullableCompanyRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
