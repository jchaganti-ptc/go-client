/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27960-e195de6ac56c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCompanyUserParams struct for BTCompanyUserParams
type BTCompanyUserParams struct {
	// Indicates the user is an admin if true.
	Admin *bool `json:"admin,omitempty"`
	// Company ID of the user.
	CompanyId *string `json:"companyId,omitempty"`
	// String to override documentation name.
	DocumentationNameOverride *string `json:"documentationNameOverride,omitempty"`
	// Email ID of the company user.
	Email *string `json:"email,omitempty"`
	// List of global permissions to grant.
	GlobalPermissions []int32 `json:"globalPermissions,omitempty"`
	// Indicates the user is a guest user if true.
	Guest *bool `json:"guest,omitempty"`
	// Indicates the user is a light user if true.
	Light *bool `json:"light,omitempty"`
}

// NewBTCompanyUserParams instantiates a new BTCompanyUserParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCompanyUserParams() *BTCompanyUserParams {
	this := BTCompanyUserParams{}
	return &this
}

// NewBTCompanyUserParamsWithDefaults instantiates a new BTCompanyUserParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCompanyUserParamsWithDefaults() *BTCompanyUserParams {
	this := BTCompanyUserParams{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *BTCompanyUserParams) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserParams) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *BTCompanyUserParams) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *BTCompanyUserParams) SetAdmin(v bool) {
	o.Admin = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTCompanyUserParams) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTCompanyUserParams) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTCompanyUserParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetDocumentationNameOverride returns the DocumentationNameOverride field value if set, zero value otherwise.
func (o *BTCompanyUserParams) GetDocumentationNameOverride() string {
	if o == nil || o.DocumentationNameOverride == nil {
		var ret string
		return ret
	}
	return *o.DocumentationNameOverride
}

// GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserParams) GetDocumentationNameOverrideOk() (*string, bool) {
	if o == nil || o.DocumentationNameOverride == nil {
		return nil, false
	}
	return o.DocumentationNameOverride, true
}

// HasDocumentationNameOverride returns a boolean if a field has been set.
func (o *BTCompanyUserParams) HasDocumentationNameOverride() bool {
	if o != nil && o.DocumentationNameOverride != nil {
		return true
	}

	return false
}

// SetDocumentationNameOverride gets a reference to the given string and assigns it to the DocumentationNameOverride field.
func (o *BTCompanyUserParams) SetDocumentationNameOverride(v string) {
	o.DocumentationNameOverride = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTCompanyUserParams) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserParams) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTCompanyUserParams) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTCompanyUserParams) SetEmail(v string) {
	o.Email = &v
}

// GetGlobalPermissions returns the GlobalPermissions field value if set, zero value otherwise.
func (o *BTCompanyUserParams) GetGlobalPermissions() []int32 {
	if o == nil || o.GlobalPermissions == nil {
		var ret []int32
		return ret
	}
	return o.GlobalPermissions
}

// GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserParams) GetGlobalPermissionsOk() ([]int32, bool) {
	if o == nil || o.GlobalPermissions == nil {
		return nil, false
	}
	return o.GlobalPermissions, true
}

// HasGlobalPermissions returns a boolean if a field has been set.
func (o *BTCompanyUserParams) HasGlobalPermissions() bool {
	if o != nil && o.GlobalPermissions != nil {
		return true
	}

	return false
}

// SetGlobalPermissions gets a reference to the given []int32 and assigns it to the GlobalPermissions field.
func (o *BTCompanyUserParams) SetGlobalPermissions(v []int32) {
	o.GlobalPermissions = v
}

// GetGuest returns the Guest field value if set, zero value otherwise.
func (o *BTCompanyUserParams) GetGuest() bool {
	if o == nil || o.Guest == nil {
		var ret bool
		return ret
	}
	return *o.Guest
}

// GetGuestOk returns a tuple with the Guest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserParams) GetGuestOk() (*bool, bool) {
	if o == nil || o.Guest == nil {
		return nil, false
	}
	return o.Guest, true
}

// HasGuest returns a boolean if a field has been set.
func (o *BTCompanyUserParams) HasGuest() bool {
	if o != nil && o.Guest != nil {
		return true
	}

	return false
}

// SetGuest gets a reference to the given bool and assigns it to the Guest field.
func (o *BTCompanyUserParams) SetGuest(v bool) {
	o.Guest = &v
}

// GetLight returns the Light field value if set, zero value otherwise.
func (o *BTCompanyUserParams) GetLight() bool {
	if o == nil || o.Light == nil {
		var ret bool
		return ret
	}
	return *o.Light
}

// GetLightOk returns a tuple with the Light field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCompanyUserParams) GetLightOk() (*bool, bool) {
	if o == nil || o.Light == nil {
		return nil, false
	}
	return o.Light, true
}

// HasLight returns a boolean if a field has been set.
func (o *BTCompanyUserParams) HasLight() bool {
	if o != nil && o.Light != nil {
		return true
	}

	return false
}

// SetLight gets a reference to the given bool and assigns it to the Light field.
func (o *BTCompanyUserParams) SetLight(v bool) {
	o.Light = &v
}

func (o BTCompanyUserParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.DocumentationNameOverride != nil {
		toSerialize["documentationNameOverride"] = o.DocumentationNameOverride
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.GlobalPermissions != nil {
		toSerialize["globalPermissions"] = o.GlobalPermissions
	}
	if o.Guest != nil {
		toSerialize["guest"] = o.Guest
	}
	if o.Light != nil {
		toSerialize["light"] = o.Light
	}
	return json.Marshal(toSerialize)
}

type NullableBTCompanyUserParams struct {
	value *BTCompanyUserParams
	isSet bool
}

func (v NullableBTCompanyUserParams) Get() *BTCompanyUserParams {
	return v.value
}

func (v *NullableBTCompanyUserParams) Set(val *BTCompanyUserParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCompanyUserParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCompanyUserParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCompanyUserParams(val *BTCompanyUserParams) *NullableBTCompanyUserParams {
	return &NullableBTCompanyUserParams{value: val, isSet: true}
}

func (v NullableBTCompanyUserParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCompanyUserParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
