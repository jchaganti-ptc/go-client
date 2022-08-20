/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6126-5c3a878ad24b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAclInfo struct for BTAclInfo
type BTAclInfo struct {
	Admin             *bool                `json:"admin,omitempty"`
	Entries           []BTAclEntryInfo     `json:"entries,omitempty"`
	Href              *string              `json:"href,omitempty"`
	Id                *string              `json:"id,omitempty"`
	InheritedAcls     []BTInheritedAclInfo `json:"inheritedAcls,omitempty"`
	Name              *string              `json:"name,omitempty"`
	ObjectId          *string              `json:"objectId,omitempty"`
	ObjectType        *int64               `json:"objectType,omitempty"`
	Owner             *BTOwnerInfo         `json:"owner,omitempty"`
	Public            *bool                `json:"public,omitempty"`
	SharedWithSupport *bool                `json:"sharedWithSupport,omitempty"`
	ViewRef           *string              `json:"viewRef,omitempty"`
	Visibility        *string              `json:"visibility,omitempty"`
}

// NewBTAclInfo instantiates a new BTAclInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAclInfo() *BTAclInfo {
	this := BTAclInfo{}
	return &this
}

// NewBTAclInfoWithDefaults instantiates a new BTAclInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAclInfoWithDefaults() *BTAclInfo {
	this := BTAclInfo{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *BTAclInfo) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *BTAclInfo) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *BTAclInfo) SetAdmin(v bool) {
	o.Admin = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *BTAclInfo) GetEntries() []BTAclEntryInfo {
	if o == nil || o.Entries == nil {
		var ret []BTAclEntryInfo
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetEntriesOk() ([]BTAclEntryInfo, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *BTAclInfo) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []BTAclEntryInfo and assigns it to the Entries field.
func (o *BTAclInfo) SetEntries(v []BTAclEntryInfo) {
	o.Entries = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTAclInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTAclInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTAclInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTAclInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTAclInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTAclInfo) SetId(v string) {
	o.Id = &v
}

// GetInheritedAcls returns the InheritedAcls field value if set, zero value otherwise.
func (o *BTAclInfo) GetInheritedAcls() []BTInheritedAclInfo {
	if o == nil || o.InheritedAcls == nil {
		var ret []BTInheritedAclInfo
		return ret
	}
	return o.InheritedAcls
}

// GetInheritedAclsOk returns a tuple with the InheritedAcls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetInheritedAclsOk() ([]BTInheritedAclInfo, bool) {
	if o == nil || o.InheritedAcls == nil {
		return nil, false
	}
	return o.InheritedAcls, true
}

// HasInheritedAcls returns a boolean if a field has been set.
func (o *BTAclInfo) HasInheritedAcls() bool {
	if o != nil && o.InheritedAcls != nil {
		return true
	}

	return false
}

// SetInheritedAcls gets a reference to the given []BTInheritedAclInfo and assigns it to the InheritedAcls field.
func (o *BTAclInfo) SetInheritedAcls(v []BTInheritedAclInfo) {
	o.InheritedAcls = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAclInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAclInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAclInfo) SetName(v string) {
	o.Name = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *BTAclInfo) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *BTAclInfo) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *BTAclInfo) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTAclInfo) GetObjectType() int64 {
	if o == nil || o.ObjectType == nil {
		var ret int64
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetObjectTypeOk() (*int64, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTAclInfo) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given int64 and assigns it to the ObjectType field.
func (o *BTAclInfo) SetObjectType(v int64) {
	o.ObjectType = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BTAclInfo) GetOwner() BTOwnerInfo {
	if o == nil || o.Owner == nil {
		var ret BTOwnerInfo
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetOwnerOk() (*BTOwnerInfo, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BTAclInfo) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BTOwnerInfo and assigns it to the Owner field.
func (o *BTAclInfo) SetOwner(v BTOwnerInfo) {
	o.Owner = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *BTAclInfo) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *BTAclInfo) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *BTAclInfo) SetPublic(v bool) {
	o.Public = &v
}

// GetSharedWithSupport returns the SharedWithSupport field value if set, zero value otherwise.
func (o *BTAclInfo) GetSharedWithSupport() bool {
	if o == nil || o.SharedWithSupport == nil {
		var ret bool
		return ret
	}
	return *o.SharedWithSupport
}

// GetSharedWithSupportOk returns a tuple with the SharedWithSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetSharedWithSupportOk() (*bool, bool) {
	if o == nil || o.SharedWithSupport == nil {
		return nil, false
	}
	return o.SharedWithSupport, true
}

// HasSharedWithSupport returns a boolean if a field has been set.
func (o *BTAclInfo) HasSharedWithSupport() bool {
	if o != nil && o.SharedWithSupport != nil {
		return true
	}

	return false
}

// SetSharedWithSupport gets a reference to the given bool and assigns it to the SharedWithSupport field.
func (o *BTAclInfo) SetSharedWithSupport(v bool) {
	o.SharedWithSupport = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTAclInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTAclInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTAclInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BTAclInfo) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAclInfo) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BTAclInfo) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *BTAclInfo) SetVisibility(v string) {
	o.Visibility = &v
}

func (o BTAclInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
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
	if o.InheritedAcls != nil {
		toSerialize["inheritedAcls"] = o.InheritedAcls
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	if o.SharedWithSupport != nil {
		toSerialize["sharedWithSupport"] = o.SharedWithSupport
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableBTAclInfo struct {
	value *BTAclInfo
	isSet bool
}

func (v NullableBTAclInfo) Get() *BTAclInfo {
	return v.value
}

func (v *NullableBTAclInfo) Set(val *BTAclInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAclInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAclInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAclInfo(val *BTAclInfo) *NullableBTAclInfo {
	return &NullableBTAclInfo{value: val, isSet: true}
}

func (v NullableBTAclInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAclInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
