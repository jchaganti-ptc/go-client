/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19458-7ff87863110f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTExportModelProperties3216 - struct for BTExportModelProperties3216
type BTExportModelProperties3216 struct {
	implBTExportModelProperties3216 interface{}
}

// BTExportBodyProperties3559AsBTExportModelProperties3216 is a convenience function that returns BTExportBodyProperties3559 wrapped in BTExportModelProperties3216
func (o *BTExportBodyProperties3559) AsBTExportModelProperties3216() *BTExportModelProperties3216 {
	return &BTExportModelProperties3216{o}
}

// NewBTExportModelProperties3216 instantiates a new BTExportModelProperties3216 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelProperties3216() *BTExportModelProperties3216 {
	this := BTExportModelProperties3216{Newbase_BTExportModelProperties3216()}
	return &this
}

// NewBTExportModelProperties3216WithDefaults instantiates a new BTExportModelProperties3216 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelProperties3216WithDefaults() *BTExportModelProperties3216 {
	this := BTExportModelProperties3216{Newbase_BTExportModelProperties3216WithDefaults()}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTExportModelProperties3216) GetAppearance() BTGraphicsAppearance1152 {
	type getResult interface {
		GetAppearance() BTGraphicsAppearance1152
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAppearance()
	} else {
		var de BTGraphicsAppearance1152
		return de
	}
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelProperties3216) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	type getResult interface {
		GetAppearanceOk() (*BTGraphicsAppearance1152, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAppearanceOk()
	} else {
		return nil, false
	}
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTExportModelProperties3216) HasAppearance() bool {
	type getResult interface {
		HasAppearance() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasAppearance()
	} else {
		return false
	}
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTExportModelProperties3216) SetAppearance(v BTGraphicsAppearance1152) {
	type getResult interface {
		SetAppearance(v BTGraphicsAppearance1152)
	}

	o.GetActualInstance().(getResult).SetAppearance(v)
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportModelProperties3216) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelProperties3216) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportModelProperties3216) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportModelProperties3216) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTExportModelProperties3216) GetName() string {
	type getResult interface {
		GetName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetName()
	} else {
		var de string
		return de
	}
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelProperties3216) GetNameOk() (*string, bool) {
	type getResult interface {
		GetNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNameOk()
	} else {
		return nil, false
	}
}

// HasName returns a boolean if a field has been set.
func (o *BTExportModelProperties3216) HasName() bool {
	type getResult interface {
		HasName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasName()
	} else {
		return false
	}
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTExportModelProperties3216) SetName(v string) {
	type getResult interface {
		SetName(v string)
	}

	o.GetActualInstance().(getResult).SetName(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTExportModelProperties3216) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTExportBodyProperties-3559'
	if jsonDict["btType"] == "BTExportBodyProperties-3559" {
		// try to unmarshal JSON data into BTExportBodyProperties3559
		var qr *BTExportBodyProperties3559
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTExportModelProperties3216 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTExportModelProperties3216 = nil
			return fmt.Errorf("Failed to unmarshal BTExportModelProperties3216 as BTExportBodyProperties3559: %s", err.Error())
		}
	}

	var qtx *base_BTExportModelProperties3216
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTExportModelProperties3216 = qtx
		return nil // data stored in dst.base_BTExportModelProperties3216, return on the first match
	} else {
		dst.implBTExportModelProperties3216 = nil
		return fmt.Errorf("Failed to unmarshal BTExportModelProperties3216 as base_BTExportModelProperties3216: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTExportModelProperties3216) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTExportModelProperties3216) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTExportModelProperties3216
}

type NullableBTExportModelProperties3216 struct {
	value *BTExportModelProperties3216
	isSet bool
}

func (v NullableBTExportModelProperties3216) Get() *BTExportModelProperties3216 {
	return v.value
}

func (v *NullableBTExportModelProperties3216) Set(val *BTExportModelProperties3216) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelProperties3216) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelProperties3216) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelProperties3216(val *BTExportModelProperties3216) *NullableBTExportModelProperties3216 {
	return &NullableBTExportModelProperties3216{value: val, isSet: true}
}

func (v NullableBTExportModelProperties3216) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelProperties3216) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTExportModelProperties3216 struct {
	Appearance *BTGraphicsAppearance1152 `json:"appearance,omitempty"`
	BtType     *string                   `json:"btType,omitempty"`
	Name       *string                   `json:"name,omitempty"`
}

// Newbase_BTExportModelProperties3216 instantiates a new base_BTExportModelProperties3216 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTExportModelProperties3216() *base_BTExportModelProperties3216 {
	this := base_BTExportModelProperties3216{}
	return &this
}

// Newbase_BTExportModelProperties3216WithDefaults instantiates a new base_BTExportModelProperties3216 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTExportModelProperties3216WithDefaults() *base_BTExportModelProperties3216 {
	this := base_BTExportModelProperties3216{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *base_BTExportModelProperties3216) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelProperties3216) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *base_BTExportModelProperties3216) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *base_BTExportModelProperties3216) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTExportModelProperties3216) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelProperties3216) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTExportModelProperties3216) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTExportModelProperties3216) SetBtType(v string) {
	o.BtType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *base_BTExportModelProperties3216) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTExportModelProperties3216) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *base_BTExportModelProperties3216) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *base_BTExportModelProperties3216) SetName(v string) {
	o.Name = &v
}

func (o base_BTExportModelProperties3216) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}
