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

// BTAclEntryInfo struct for BTAclEntryInfo
type BTAclEntryInfo struct {
	AcceptOwnerTransfer  *bool        `json:"acceptOwnerTransfer,omitempty"`
	CompanyName          *string      `json:"companyName,omitempty"`
	Email                *string      `json:"email,omitempty"`
	EnterpriseMember     *bool        `json:"enterpriseMember,omitempty"`
	EntryId              *string      `json:"entryId,omitempty"`
	EntryState           *BTUserState `json:"entryState,omitempty"`
	EntryType            *int32       `json:"entryType,omitempty"`
	Image                *string      `json:"image,omitempty"`
	Name                 *string      `json:"name,omitempty"`
	ObjectId             *string      `json:"objectId,omitempty"`
	PendingOwnerTransfer *bool        `json:"pendingOwnerTransfer,omitempty"`
	Permission           *int64       `json:"permission,omitempty"`
	PermissionSet        []string     `json:"permissionSet,omitempty"`
	TeamName             *string      `json:"teamName,omitempty"`
}

// NewBTAclEntryInfo instantiates a new BTAclEntryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAclEntryInfo() *BTAclEntryInfo {
	this := BTAclEntryInfo{}
	return &this
}

// NewBTAclEntryInfoWithDefaults instantiates a new BTAclEntryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAclEntryInfoWithDefaults() *BTAclEntryInfo {
	this := BTAclEntryInfo{}
	return &this
}

// GetAcceptOwnerTransfer returns the AcceptOwnerTransfer field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetAcceptOwnerTransfer() bool {
	if o == nil || o.AcceptOwnerTransfer == nil {
		var ret bool
		return ret
	}
	return *o.AcceptOwnerTransfer
}

// GetAcceptOwnerTransferOk returns a tuple with the AcceptOwnerTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetAcceptOwnerTransferOk() (*bool, bool) {
	if o == nil || o.AcceptOwnerTransfer == nil {
		return nil, false
	}
	return o.AcceptOwnerTransfer, true
}

// HasAcceptOwnerTransfer returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasAcceptOwnerTransfer() bool {
	if o != nil && o.AcceptOwnerTransfer != nil {
		return true
	}

	return false
}

// SetAcceptOwnerTransfer gets a reference to the given bool and assigns it to the AcceptOwnerTransfer field.
func (o *BTAclEntryInfo) SetAcceptOwnerTransfer(v bool) {
	o.AcceptOwnerTransfer = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *BTAclEntryInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTAclEntryInfo) SetEmail(v string) {
	o.Email = &v
}

// GetEnterpriseMember returns the EnterpriseMember field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetEnterpriseMember() bool {
	if o == nil || o.EnterpriseMember == nil {
		var ret bool
		return ret
	}
	return *o.EnterpriseMember
}

// GetEnterpriseMemberOk returns a tuple with the EnterpriseMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetEnterpriseMemberOk() (*bool, bool) {
	if o == nil || o.EnterpriseMember == nil {
		return nil, false
	}
	return o.EnterpriseMember, true
}

// HasEnterpriseMember returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasEnterpriseMember() bool {
	if o != nil && o.EnterpriseMember != nil {
		return true
	}

	return false
}

// SetEnterpriseMember gets a reference to the given bool and assigns it to the EnterpriseMember field.
func (o *BTAclEntryInfo) SetEnterpriseMember(v bool) {
	o.EnterpriseMember = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *BTAclEntryInfo) SetEntryId(v string) {
	o.EntryId = &v
}

// GetEntryState returns the EntryState field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetEntryState() BTUserState {
	if o == nil || o.EntryState == nil {
		var ret BTUserState
		return ret
	}
	return *o.EntryState
}

// GetEntryStateOk returns a tuple with the EntryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetEntryStateOk() (*BTUserState, bool) {
	if o == nil || o.EntryState == nil {
		return nil, false
	}
	return o.EntryState, true
}

// HasEntryState returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasEntryState() bool {
	if o != nil && o.EntryState != nil {
		return true
	}

	return false
}

// SetEntryState gets a reference to the given BTUserState and assigns it to the EntryState field.
func (o *BTAclEntryInfo) SetEntryState(v BTUserState) {
	o.EntryState = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetEntryType() int32 {
	if o == nil || o.EntryType == nil {
		var ret int32
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetEntryTypeOk() (*int32, bool) {
	if o == nil || o.EntryType == nil {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasEntryType() bool {
	if o != nil && o.EntryType != nil {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given int32 and assigns it to the EntryType field.
func (o *BTAclEntryInfo) SetEntryType(v int32) {
	o.EntryType = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTAclEntryInfo) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAclEntryInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *BTAclEntryInfo) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetPendingOwnerTransfer returns the PendingOwnerTransfer field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetPendingOwnerTransfer() bool {
	if o == nil || o.PendingOwnerTransfer == nil {
		var ret bool
		return ret
	}
	return *o.PendingOwnerTransfer
}

// GetPendingOwnerTransferOk returns a tuple with the PendingOwnerTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetPendingOwnerTransferOk() (*bool, bool) {
	if o == nil || o.PendingOwnerTransfer == nil {
		return nil, false
	}
	return o.PendingOwnerTransfer, true
}

// HasPendingOwnerTransfer returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasPendingOwnerTransfer() bool {
	if o != nil && o.PendingOwnerTransfer != nil {
		return true
	}

	return false
}

// SetPendingOwnerTransfer gets a reference to the given bool and assigns it to the PendingOwnerTransfer field.
func (o *BTAclEntryInfo) SetPendingOwnerTransfer(v bool) {
	o.PendingOwnerTransfer = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetPermission() int64 {
	if o == nil || o.Permission == nil {
		var ret int64
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetPermissionOk() (*int64, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given int64 and assigns it to the Permission field.
func (o *BTAclEntryInfo) SetPermission(v int64) {
	o.Permission = &v
}

// GetPermissionSet returns the PermissionSet field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetPermissionSet() []string {
	if o == nil || o.PermissionSet == nil {
		var ret []string
		return ret
	}
	return o.PermissionSet
}

// GetPermissionSetOk returns a tuple with the PermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetPermissionSetOk() ([]string, bool) {
	if o == nil || o.PermissionSet == nil {
		return nil, false
	}
	return o.PermissionSet, true
}

// HasPermissionSet returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasPermissionSet() bool {
	if o != nil && o.PermissionSet != nil {
		return true
	}

	return false
}

// SetPermissionSet gets a reference to the given []string and assigns it to the PermissionSet field.
func (o *BTAclEntryInfo) SetPermissionSet(v []string) {
	o.PermissionSet = v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *BTAclEntryInfo) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclEntryInfo) GetTeamNameOk() (*string, bool) {
	if o == nil || o.TeamName == nil {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *BTAclEntryInfo) HasTeamName() bool {
	if o != nil && o.TeamName != nil {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *BTAclEntryInfo) SetTeamName(v string) {
	o.TeamName = &v
}

func (o BTAclEntryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptOwnerTransfer != nil {
		toSerialize["acceptOwnerTransfer"] = o.AcceptOwnerTransfer
	}
	if o.CompanyName != nil {
		toSerialize["companyName"] = o.CompanyName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EnterpriseMember != nil {
		toSerialize["enterpriseMember"] = o.EnterpriseMember
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.EntryState != nil {
		toSerialize["entryState"] = o.EntryState
	}
	if o.EntryType != nil {
		toSerialize["entryType"] = o.EntryType
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.PendingOwnerTransfer != nil {
		toSerialize["pendingOwnerTransfer"] = o.PendingOwnerTransfer
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.PermissionSet != nil {
		toSerialize["permissionSet"] = o.PermissionSet
	}
	if o.TeamName != nil {
		toSerialize["teamName"] = o.TeamName
	}
	return json.Marshal(toSerialize)
}

type NullableBTAclEntryInfo struct {
	value *BTAclEntryInfo
	isSet bool
}

func (v NullableBTAclEntryInfo) Get() *BTAclEntryInfo {
	return v.value
}

func (v *NullableBTAclEntryInfo) Set(val *BTAclEntryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAclEntryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAclEntryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAclEntryInfo(val *BTAclEntryInfo) *NullableBTAclEntryInfo {
	return &NullableBTAclEntryInfo{value: val, isSet: true}
}

func (v NullableBTAclEntryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAclEntryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
