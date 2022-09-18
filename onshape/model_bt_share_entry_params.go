/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6510-4a0bae47e25c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTShareEntryParams struct for BTShareEntryParams
type BTShareEntryParams struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	CompanyId     *string `json:"companyId,omitempty"`
	Email         *string `json:"email,omitempty"`
	EntryType     *int32  `json:"entryType,omitempty"`
	TeamId        *string `json:"teamId,omitempty"`
	UserId        *string `json:"userId,omitempty"`
}

// NewBTShareEntryParams instantiates a new BTShareEntryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTShareEntryParams() *BTShareEntryParams {
	this := BTShareEntryParams{}
	return &this
}

// NewBTShareEntryParamsWithDefaults instantiates a new BTShareEntryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTShareEntryParamsWithDefaults() *BTShareEntryParams {
	this := BTShareEntryParams{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *BTShareEntryParams) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareEntryParams) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *BTShareEntryParams) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *BTShareEntryParams) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTShareEntryParams) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareEntryParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTShareEntryParams) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTShareEntryParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTShareEntryParams) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareEntryParams) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTShareEntryParams) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTShareEntryParams) SetEmail(v string) {
	o.Email = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *BTShareEntryParams) GetEntryType() int32 {
	if o == nil || o.EntryType == nil {
		var ret int32
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareEntryParams) GetEntryTypeOk() (*int32, bool) {
	if o == nil || o.EntryType == nil {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *BTShareEntryParams) HasEntryType() bool {
	if o != nil && o.EntryType != nil {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given int32 and assigns it to the EntryType field.
func (o *BTShareEntryParams) SetEntryType(v int32) {
	o.EntryType = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *BTShareEntryParams) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareEntryParams) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *BTShareEntryParams) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *BTShareEntryParams) SetTeamId(v string) {
	o.TeamId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTShareEntryParams) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShareEntryParams) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTShareEntryParams) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTShareEntryParams) SetUserId(v string) {
	o.UserId = &v
}

func (o BTShareEntryParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationId != nil {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EntryType != nil {
		toSerialize["entryType"] = o.EntryType
	}
	if o.TeamId != nil {
		toSerialize["teamId"] = o.TeamId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableBTShareEntryParams struct {
	value *BTShareEntryParams
	isSet bool
}

func (v NullableBTShareEntryParams) Get() *BTShareEntryParams {
	return v.value
}

func (v *NullableBTShareEntryParams) Set(val *BTShareEntryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTShareEntryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTShareEntryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTShareEntryParams(val *BTShareEntryParams) *NullableBTShareEntryParams {
	return &NullableBTShareEntryParams{value: val, isSet: true}
}

func (v NullableBTShareEntryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTShareEntryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
