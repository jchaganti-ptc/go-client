/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5805-1ec839b76add
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTIdentityInfo struct for BTIdentityInfo
type BTIdentityInfo struct {
	Href         *string            `json:"href,omitempty"`
	Id           *string            `json:"id,omitempty"`
	IdentityType *int32             `json:"identityType,omitempty"`
	Team         *BTTeamSummaryInfo `json:"team,omitempty"`
	User         *BTUserSummaryInfo `json:"user,omitempty"`
	ViewRef      *string            `json:"viewRef,omitempty"`
}

// NewBTIdentityInfo instantiates a new BTIdentityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTIdentityInfo() *BTIdentityInfo {
	this := BTIdentityInfo{}
	return &this
}

// NewBTIdentityInfoWithDefaults instantiates a new BTIdentityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTIdentityInfoWithDefaults() *BTIdentityInfo {
	this := BTIdentityInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTIdentityInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdentityInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTIdentityInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTIdentityInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTIdentityInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdentityInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTIdentityInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTIdentityInfo) SetId(v string) {
	o.Id = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *BTIdentityInfo) GetIdentityType() int32 {
	if o == nil || o.IdentityType == nil {
		var ret int32
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdentityInfo) GetIdentityTypeOk() (*int32, bool) {
	if o == nil || o.IdentityType == nil {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *BTIdentityInfo) HasIdentityType() bool {
	if o != nil && o.IdentityType != nil {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given int32 and assigns it to the IdentityType field.
func (o *BTIdentityInfo) SetIdentityType(v int32) {
	o.IdentityType = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *BTIdentityInfo) GetTeam() BTTeamSummaryInfo {
	if o == nil || o.Team == nil {
		var ret BTTeamSummaryInfo
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdentityInfo) GetTeamOk() (*BTTeamSummaryInfo, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *BTIdentityInfo) HasTeam() bool {
	if o != nil && o.Team != nil {
		return true
	}

	return false
}

// SetTeam gets a reference to the given BTTeamSummaryInfo and assigns it to the Team field.
func (o *BTIdentityInfo) SetTeam(v BTTeamSummaryInfo) {
	o.Team = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *BTIdentityInfo) GetUser() BTUserSummaryInfo {
	if o == nil || o.User == nil {
		var ret BTUserSummaryInfo
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdentityInfo) GetUserOk() (*BTUserSummaryInfo, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *BTIdentityInfo) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given BTUserSummaryInfo and assigns it to the User field.
func (o *BTIdentityInfo) SetUser(v BTUserSummaryInfo) {
	o.User = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTIdentityInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdentityInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTIdentityInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTIdentityInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTIdentityInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentityType != nil {
		toSerialize["identityType"] = o.IdentityType
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTIdentityInfo struct {
	value *BTIdentityInfo
	isSet bool
}

func (v NullableBTIdentityInfo) Get() *BTIdentityInfo {
	return v.value
}

func (v *NullableBTIdentityInfo) Set(val *BTIdentityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTIdentityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTIdentityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTIdentityInfo(val *BTIdentityInfo) *NullableBTIdentityInfo {
	return &NullableBTIdentityInfo{value: val, isSet: true}
}

func (v NullableBTIdentityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTIdentityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
