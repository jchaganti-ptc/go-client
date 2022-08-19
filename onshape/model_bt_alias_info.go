/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6108-60a91d537e42
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAliasInfo struct for BTAliasInfo
type BTAliasInfo struct {
	CompanyId   *string            `json:"companyId,omitempty"`
	CreatedAt   *JSONTime          `json:"createdAt,omitempty"`
	Description *string            `json:"description,omitempty"`
	Entries     []BTAliasEntryInfo `json:"entries,omitempty"`
	Href        *string            `json:"href,omitempty"`
	Id          *string            `json:"id,omitempty"`
	Identities  []BTIdentityInfo   `json:"identities,omitempty"`
	Name        *string            `json:"name,omitempty"`
	ViewRef     *string            `json:"viewRef,omitempty"`
}

// NewBTAliasInfo instantiates a new BTAliasInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAliasInfo() *BTAliasInfo {
	this := BTAliasInfo{}
	return &this
}

// NewBTAliasInfoWithDefaults instantiates a new BTAliasInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAliasInfoWithDefaults() *BTAliasInfo {
	this := BTAliasInfo{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTAliasInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTAliasInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTAliasInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTAliasInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTAliasInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTAliasInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAliasInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAliasInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAliasInfo) SetDescription(v string) {
	o.Description = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BTAliasInfo) GetEntries() []BTAliasEntryInfo {
	if o == nil || o.Entries == nil {
		var ret []BTAliasEntryInfo
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetEntriesOk() ([]BTAliasEntryInfo, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BTAliasInfo) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []BTAliasEntryInfo and assigns it to the Entries field.
func (o *BTAliasInfo) SetEntries(v []BTAliasEntryInfo) {
	o.Entries = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTAliasInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTAliasInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTAliasInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTAliasInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTAliasInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTAliasInfo) SetId(v string) {
	o.Id = &v
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *BTAliasInfo) GetIdentities() []BTIdentityInfo {
	if o == nil || o.Identities == nil {
		var ret []BTIdentityInfo
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetIdentitiesOk() ([]BTIdentityInfo, bool) {
	if o == nil || o.Identities == nil {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *BTAliasInfo) HasIdentities() bool {
	if o != nil && o.Identities != nil {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []BTIdentityInfo and assigns it to the Identities field.
func (o *BTAliasInfo) SetIdentities(v []BTIdentityInfo) {
	o.Identities = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAliasInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAliasInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAliasInfo) SetName(v string) {
	o.Name = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTAliasInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAliasInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTAliasInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTAliasInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTAliasInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Identities != nil {
		toSerialize["identities"] = o.Identities
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTAliasInfo struct {
	value *BTAliasInfo
	isSet bool
}

func (v NullableBTAliasInfo) Get() *BTAliasInfo {
	return v.value
}

func (v *NullableBTAliasInfo) Set(val *BTAliasInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAliasInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAliasInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAliasInfo(val *BTAliasInfo) *NullableBTAliasInfo {
	return &NullableBTAliasInfo{value: val, isSet: true}
}

func (v NullableBTAliasInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAliasInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
