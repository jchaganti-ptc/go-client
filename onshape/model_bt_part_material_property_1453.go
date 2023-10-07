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

// BTPartMaterialProperty1453 struct for BTPartMaterialProperty1453
type BTPartMaterialProperty1453 struct {
	BtType      *string `json:"btType,omitempty"`
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	Units       *string `json:"units,omitempty"`
	Value       *string `json:"value,omitempty"`
}

// NewBTPartMaterialProperty1453 instantiates a new BTPartMaterialProperty1453 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartMaterialProperty1453() *BTPartMaterialProperty1453 {
	this := BTPartMaterialProperty1453{}
	return &this
}

// NewBTPartMaterialProperty1453WithDefaults instantiates a new BTPartMaterialProperty1453 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartMaterialProperty1453WithDefaults() *BTPartMaterialProperty1453 {
	this := BTPartMaterialProperty1453{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartMaterialProperty1453) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialProperty1453) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartMaterialProperty1453) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartMaterialProperty1453) SetBtType(v string) {
	o.BtType = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *BTPartMaterialProperty1453) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialProperty1453) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *BTPartMaterialProperty1453) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *BTPartMaterialProperty1453) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTPartMaterialProperty1453) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialProperty1453) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTPartMaterialProperty1453) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTPartMaterialProperty1453) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BTPartMaterialProperty1453) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialProperty1453) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BTPartMaterialProperty1453) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BTPartMaterialProperty1453) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPartMaterialProperty1453) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialProperty1453) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPartMaterialProperty1453) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPartMaterialProperty1453) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTPartMaterialProperty1453) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialProperty1453) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTPartMaterialProperty1453) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTPartMaterialProperty1453) SetType(v string) {
	o.Type = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTPartMaterialProperty1453) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialProperty1453) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BTPartMaterialProperty1453) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *BTPartMaterialProperty1453) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTPartMaterialProperty1453) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMaterialProperty1453) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTPartMaterialProperty1453) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTPartMaterialProperty1453) SetValue(v string) {
	o.Value = &v
}

func (o BTPartMaterialProperty1453) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
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

type NullableBTPartMaterialProperty1453 struct {
	value *BTPartMaterialProperty1453
	isSet bool
}

func (v NullableBTPartMaterialProperty1453) Get() *BTPartMaterialProperty1453 {
	return v.value
}

func (v *NullableBTPartMaterialProperty1453) Set(val *BTPartMaterialProperty1453) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartMaterialProperty1453) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartMaterialProperty1453) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartMaterialProperty1453(val *BTPartMaterialProperty1453) *NullableBTPartMaterialProperty1453 {
	return &NullableBTPartMaterialProperty1453{value: val, isSet: true}
}

func (v NullableBTPartMaterialProperty1453) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartMaterialProperty1453) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
