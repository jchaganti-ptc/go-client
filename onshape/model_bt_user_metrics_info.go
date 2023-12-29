/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28126-700645382199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUserMetricsInfo struct for BTUserMetricsInfo
type BTUserMetricsInfo struct {
	CreatedDocuments           *int64 `json:"createdDocuments,omitempty"`
	HasRecentlyOpenedDocuments *bool  `json:"hasRecentlyOpenedDocuments,omitempty"`
	HasSharedDocuments_        *bool  `json:"hasSharedDocuments,omitempty"`
	HasTrashedDocuments        *bool  `json:"hasTrashedDocuments,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name                     *string `json:"name,omitempty"`
	PrivateDocuments         *int64  `json:"privateDocuments,omitempty"`
	PublicDocuments          *int64  `json:"publicDocuments,omitempty"`
	SharedDocuments          *int64  `json:"sharedDocuments,omitempty"`
	UserAccountLimitsCrossed *bool   `json:"userAccountLimitsCrossed,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTUserMetricsInfo instantiates a new BTUserMetricsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserMetricsInfo() *BTUserMetricsInfo {
	this := BTUserMetricsInfo{}
	return &this
}

// NewBTUserMetricsInfoWithDefaults instantiates a new BTUserMetricsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserMetricsInfoWithDefaults() *BTUserMetricsInfo {
	this := BTUserMetricsInfo{}
	return &this
}

// GetCreatedDocuments returns the CreatedDocuments field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetCreatedDocuments() int64 {
	if o == nil || o.CreatedDocuments == nil {
		var ret int64
		return ret
	}
	return *o.CreatedDocuments
}

// GetCreatedDocumentsOk returns a tuple with the CreatedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetCreatedDocumentsOk() (*int64, bool) {
	if o == nil || o.CreatedDocuments == nil {
		return nil, false
	}
	return o.CreatedDocuments, true
}

// HasCreatedDocuments returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasCreatedDocuments() bool {
	if o != nil && o.CreatedDocuments != nil {
		return true
	}

	return false
}

// SetCreatedDocuments gets a reference to the given int64 and assigns it to the CreatedDocuments field.
func (o *BTUserMetricsInfo) SetCreatedDocuments(v int64) {
	o.CreatedDocuments = &v
}

// GetHasRecentlyOpenedDocuments returns the HasRecentlyOpenedDocuments field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetHasRecentlyOpenedDocuments() bool {
	if o == nil || o.HasRecentlyOpenedDocuments == nil {
		var ret bool
		return ret
	}
	return *o.HasRecentlyOpenedDocuments
}

// GetHasRecentlyOpenedDocumentsOk returns a tuple with the HasRecentlyOpenedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetHasRecentlyOpenedDocumentsOk() (*bool, bool) {
	if o == nil || o.HasRecentlyOpenedDocuments == nil {
		return nil, false
	}
	return o.HasRecentlyOpenedDocuments, true
}

// HasHasRecentlyOpenedDocuments returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasHasRecentlyOpenedDocuments() bool {
	if o != nil && o.HasRecentlyOpenedDocuments != nil {
		return true
	}

	return false
}

// SetHasRecentlyOpenedDocuments gets a reference to the given bool and assigns it to the HasRecentlyOpenedDocuments field.
func (o *BTUserMetricsInfo) SetHasRecentlyOpenedDocuments(v bool) {
	o.HasRecentlyOpenedDocuments = &v
}

// GetHasSharedDocuments_ returns the HasSharedDocuments_ field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetHasSharedDocuments_() bool {
	if o == nil || o.HasSharedDocuments_ == nil {
		var ret bool
		return ret
	}
	return *o.HasSharedDocuments_
}

// GetHasSharedDocuments_Ok returns a tuple with the HasSharedDocuments_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetHasSharedDocuments_Ok() (*bool, bool) {
	if o == nil || o.HasSharedDocuments_ == nil {
		return nil, false
	}
	return o.HasSharedDocuments_, true
}

// HasHasSharedDocuments_ returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasHasSharedDocuments_() bool {
	if o != nil && o.HasSharedDocuments_ != nil {
		return true
	}

	return false
}

// SetHasSharedDocuments_ gets a reference to the given bool and assigns it to the HasSharedDocuments_ field.
func (o *BTUserMetricsInfo) SetHasSharedDocuments_(v bool) {
	o.HasSharedDocuments_ = &v
}

// GetHasTrashedDocuments returns the HasTrashedDocuments field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetHasTrashedDocuments() bool {
	if o == nil || o.HasTrashedDocuments == nil {
		var ret bool
		return ret
	}
	return *o.HasTrashedDocuments
}

// GetHasTrashedDocumentsOk returns a tuple with the HasTrashedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetHasTrashedDocumentsOk() (*bool, bool) {
	if o == nil || o.HasTrashedDocuments == nil {
		return nil, false
	}
	return o.HasTrashedDocuments, true
}

// HasHasTrashedDocuments returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasHasTrashedDocuments() bool {
	if o != nil && o.HasTrashedDocuments != nil {
		return true
	}

	return false
}

// SetHasTrashedDocuments gets a reference to the given bool and assigns it to the HasTrashedDocuments field.
func (o *BTUserMetricsInfo) SetHasTrashedDocuments(v bool) {
	o.HasTrashedDocuments = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTUserMetricsInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUserMetricsInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTUserMetricsInfo) SetName(v string) {
	o.Name = &v
}

// GetPrivateDocuments returns the PrivateDocuments field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetPrivateDocuments() int64 {
	if o == nil || o.PrivateDocuments == nil {
		var ret int64
		return ret
	}
	return *o.PrivateDocuments
}

// GetPrivateDocumentsOk returns a tuple with the PrivateDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetPrivateDocumentsOk() (*int64, bool) {
	if o == nil || o.PrivateDocuments == nil {
		return nil, false
	}
	return o.PrivateDocuments, true
}

// HasPrivateDocuments returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasPrivateDocuments() bool {
	if o != nil && o.PrivateDocuments != nil {
		return true
	}

	return false
}

// SetPrivateDocuments gets a reference to the given int64 and assigns it to the PrivateDocuments field.
func (o *BTUserMetricsInfo) SetPrivateDocuments(v int64) {
	o.PrivateDocuments = &v
}

// GetPublicDocuments returns the PublicDocuments field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetPublicDocuments() int64 {
	if o == nil || o.PublicDocuments == nil {
		var ret int64
		return ret
	}
	return *o.PublicDocuments
}

// GetPublicDocumentsOk returns a tuple with the PublicDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetPublicDocumentsOk() (*int64, bool) {
	if o == nil || o.PublicDocuments == nil {
		return nil, false
	}
	return o.PublicDocuments, true
}

// HasPublicDocuments returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasPublicDocuments() bool {
	if o != nil && o.PublicDocuments != nil {
		return true
	}

	return false
}

// SetPublicDocuments gets a reference to the given int64 and assigns it to the PublicDocuments field.
func (o *BTUserMetricsInfo) SetPublicDocuments(v int64) {
	o.PublicDocuments = &v
}

// GetSharedDocuments returns the SharedDocuments field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetSharedDocuments() int64 {
	if o == nil || o.SharedDocuments == nil {
		var ret int64
		return ret
	}
	return *o.SharedDocuments
}

// GetSharedDocumentsOk returns a tuple with the SharedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetSharedDocumentsOk() (*int64, bool) {
	if o == nil || o.SharedDocuments == nil {
		return nil, false
	}
	return o.SharedDocuments, true
}

// HasSharedDocuments returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasSharedDocuments() bool {
	if o != nil && o.SharedDocuments != nil {
		return true
	}

	return false
}

// SetSharedDocuments gets a reference to the given int64 and assigns it to the SharedDocuments field.
func (o *BTUserMetricsInfo) SetSharedDocuments(v int64) {
	o.SharedDocuments = &v
}

// GetUserAccountLimitsCrossed returns the UserAccountLimitsCrossed field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetUserAccountLimitsCrossed() bool {
	if o == nil || o.UserAccountLimitsCrossed == nil {
		var ret bool
		return ret
	}
	return *o.UserAccountLimitsCrossed
}

// GetUserAccountLimitsCrossedOk returns a tuple with the UserAccountLimitsCrossed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetUserAccountLimitsCrossedOk() (*bool, bool) {
	if o == nil || o.UserAccountLimitsCrossed == nil {
		return nil, false
	}
	return o.UserAccountLimitsCrossed, true
}

// HasUserAccountLimitsCrossed returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasUserAccountLimitsCrossed() bool {
	if o != nil && o.UserAccountLimitsCrossed != nil {
		return true
	}

	return false
}

// SetUserAccountLimitsCrossed gets a reference to the given bool and assigns it to the UserAccountLimitsCrossed field.
func (o *BTUserMetricsInfo) SetUserAccountLimitsCrossed(v bool) {
	o.UserAccountLimitsCrossed = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTUserMetricsInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserMetricsInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTUserMetricsInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTUserMetricsInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTUserMetricsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDocuments != nil {
		toSerialize["createdDocuments"] = o.CreatedDocuments
	}
	if o.HasRecentlyOpenedDocuments != nil {
		toSerialize["hasRecentlyOpenedDocuments"] = o.HasRecentlyOpenedDocuments
	}
	if o.HasSharedDocuments_ != nil {
		toSerialize["hasSharedDocuments"] = o.HasSharedDocuments_
	}
	if o.HasTrashedDocuments != nil {
		toSerialize["hasTrashedDocuments"] = o.HasTrashedDocuments
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PrivateDocuments != nil {
		toSerialize["privateDocuments"] = o.PrivateDocuments
	}
	if o.PublicDocuments != nil {
		toSerialize["publicDocuments"] = o.PublicDocuments
	}
	if o.SharedDocuments != nil {
		toSerialize["sharedDocuments"] = o.SharedDocuments
	}
	if o.UserAccountLimitsCrossed != nil {
		toSerialize["userAccountLimitsCrossed"] = o.UserAccountLimitsCrossed
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserMetricsInfo struct {
	value *BTUserMetricsInfo
	isSet bool
}

func (v NullableBTUserMetricsInfo) Get() *BTUserMetricsInfo {
	return v.value
}

func (v *NullableBTUserMetricsInfo) Set(val *BTUserMetricsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserMetricsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserMetricsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserMetricsInfo(val *BTUserMetricsInfo) *NullableBTUserMetricsInfo {
	return &NullableBTUserMetricsInfo{value: val, isSet: true}
}

func (v NullableBTUserMetricsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserMetricsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
