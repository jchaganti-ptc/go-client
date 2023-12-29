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

// BTRbacRoleInfo struct for BTRbacRoleInfo
type BTRbacRoleInfo struct {
	Active      *bool   `json:"active,omitempty"`
	Description *string `json:"description,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name           *string `json:"name,omitempty"`
	PredefinedRole *int32  `json:"predefinedRole,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTRbacRoleInfo instantiates a new BTRbacRoleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRbacRoleInfo() *BTRbacRoleInfo {
	this := BTRbacRoleInfo{}
	return &this
}

// NewBTRbacRoleInfoWithDefaults instantiates a new BTRbacRoleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRbacRoleInfoWithDefaults() *BTRbacRoleInfo {
	this := BTRbacRoleInfo{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *BTRbacRoleInfo) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacRoleInfo) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *BTRbacRoleInfo) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *BTRbacRoleInfo) SetActive(v bool) {
	o.Active = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTRbacRoleInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacRoleInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTRbacRoleInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTRbacRoleInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTRbacRoleInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacRoleInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTRbacRoleInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTRbacRoleInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTRbacRoleInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacRoleInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTRbacRoleInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTRbacRoleInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTRbacRoleInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacRoleInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTRbacRoleInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTRbacRoleInfo) SetName(v string) {
	o.Name = &v
}

// GetPredefinedRole returns the PredefinedRole field value if set, zero value otherwise.
func (o *BTRbacRoleInfo) GetPredefinedRole() int32 {
	if o == nil || o.PredefinedRole == nil {
		var ret int32
		return ret
	}
	return *o.PredefinedRole
}

// GetPredefinedRoleOk returns a tuple with the PredefinedRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacRoleInfo) GetPredefinedRoleOk() (*int32, bool) {
	if o == nil || o.PredefinedRole == nil {
		return nil, false
	}
	return o.PredefinedRole, true
}

// HasPredefinedRole returns a boolean if a field has been set.
func (o *BTRbacRoleInfo) HasPredefinedRole() bool {
	if o != nil && o.PredefinedRole != nil {
		return true
	}

	return false
}

// SetPredefinedRole gets a reference to the given int32 and assigns it to the PredefinedRole field.
func (o *BTRbacRoleInfo) SetPredefinedRole(v int32) {
	o.PredefinedRole = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTRbacRoleInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRbacRoleInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTRbacRoleInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTRbacRoleInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTRbacRoleInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
	if o.PredefinedRole != nil {
		toSerialize["predefinedRole"] = o.PredefinedRole
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTRbacRoleInfo struct {
	value *BTRbacRoleInfo
	isSet bool
}

func (v NullableBTRbacRoleInfo) Get() *BTRbacRoleInfo {
	return v.value
}

func (v *NullableBTRbacRoleInfo) Set(val *BTRbacRoleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRbacRoleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRbacRoleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRbacRoleInfo(val *BTRbacRoleInfo) *NullableBTRbacRoleInfo {
	return &NullableBTRbacRoleInfo{value: val, isSet: true}
}

func (v NullableBTRbacRoleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRbacRoleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
