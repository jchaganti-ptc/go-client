/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7010-841c6a8f62e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMaterialPropertyParams struct for BTMaterialPropertyParams
type BTMaterialPropertyParams struct {
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	Units       *string `json:"units,omitempty"`
	Value       *string `json:"value,omitempty"`
}

// NewBTMaterialPropertyParams instantiates a new BTMaterialPropertyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMaterialPropertyParams() *BTMaterialPropertyParams {
	this := BTMaterialPropertyParams{}
	return &this
}

// NewBTMaterialPropertyParamsWithDefaults instantiates a new BTMaterialPropertyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMaterialPropertyParamsWithDefaults() *BTMaterialPropertyParams {
	this := BTMaterialPropertyParams{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *BTMaterialPropertyParams) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialPropertyParams) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *BTMaterialPropertyParams) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *BTMaterialPropertyParams) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTMaterialPropertyParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialPropertyParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTMaterialPropertyParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTMaterialPropertyParams) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BTMaterialPropertyParams) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialPropertyParams) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BTMaterialPropertyParams) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BTMaterialPropertyParams) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMaterialPropertyParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialPropertyParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMaterialPropertyParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMaterialPropertyParams) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTMaterialPropertyParams) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialPropertyParams) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTMaterialPropertyParams) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTMaterialPropertyParams) SetType(v string) {
	o.Type = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTMaterialPropertyParams) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialPropertyParams) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BTMaterialPropertyParams) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *BTMaterialPropertyParams) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMaterialPropertyParams) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialPropertyParams) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMaterialPropertyParams) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTMaterialPropertyParams) SetValue(v string) {
	o.Value = &v
}

func (o BTMaterialPropertyParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTMaterialPropertyParams struct {
	value *BTMaterialPropertyParams
	isSet bool
}

func (v NullableBTMaterialPropertyParams) Get() *BTMaterialPropertyParams {
	return v.value
}

func (v *NullableBTMaterialPropertyParams) Set(val *BTMaterialPropertyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMaterialPropertyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMaterialPropertyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMaterialPropertyParams(val *BTMaterialPropertyParams) *NullableBTMaterialPropertyParams {
	return &NullableBTMaterialPropertyParams{value: val, isSet: true}
}

func (v NullableBTMaterialPropertyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMaterialPropertyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
