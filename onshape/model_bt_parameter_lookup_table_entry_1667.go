/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6326-97b3616ccba2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTParameterLookupTableEntry1667 - struct for BTParameterLookupTableEntry1667
type BTParameterLookupTableEntry1667 struct {
	implBTParameterLookupTableEntry1667 interface{}
}

// BTParameterLookupTableListEntry1916AsBTParameterLookupTableEntry1667 is a convenience function that returns BTParameterLookupTableListEntry1916 wrapped in BTParameterLookupTableEntry1667
func (o *BTParameterLookupTableListEntry1916) AsBTParameterLookupTableEntry1667() *BTParameterLookupTableEntry1667 {
	return &BTParameterLookupTableEntry1667{o}
}

// NewBTParameterLookupTableEntry1667 instantiates a new BTParameterLookupTableEntry1667 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterLookupTableEntry1667() *BTParameterLookupTableEntry1667 {
	this := BTParameterLookupTableEntry1667{Newbase_BTParameterLookupTableEntry1667()}
	return &this
}

// NewBTParameterLookupTableEntry1667WithDefaults instantiates a new BTParameterLookupTableEntry1667 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterLookupTableEntry1667WithDefaults() *BTParameterLookupTableEntry1667 {
	this := BTParameterLookupTableEntry1667{Newbase_BTParameterLookupTableEntry1667WithDefaults()}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetAdditionalLocalizedStrings() int32 {
	type getResult interface {
		GetAdditionalLocalizedStrings() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAdditionalLocalizedStrings()
	} else {
		var de int32
		return de
	}
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	type getResult interface {
		GetAdditionalLocalizedStringsOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetAdditionalLocalizedStringsOk()
	} else {
		return nil, false
	}
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasAdditionalLocalizedStrings() bool {
	type getResult interface {
		HasAdditionalLocalizedStrings() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasAdditionalLocalizedStrings()
	} else {
		return false
	}
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTParameterLookupTableEntry1667) SetAdditionalLocalizedStrings(v int32) {
	type getResult interface {
		SetAdditionalLocalizedStrings(v int32)
	}

	o.GetActualInstance().(getResult).SetAdditionalLocalizedStrings(v)
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetBtType() string {
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
func (o *BTParameterLookupTableEntry1667) GetBtTypeOk() (*string, bool) {
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
func (o *BTParameterLookupTableEntry1667) HasBtType() bool {
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
func (o *BTParameterLookupTableEntry1667) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetLabel() string {
	type getResult interface {
		GetLabel() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLabel()
	} else {
		var de string
		return de
	}
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetLabelOk() (*string, bool) {
	type getResult interface {
		GetLabelOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLabelOk()
	} else {
		return nil, false
	}
}

// HasLabel returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasLabel() bool {
	type getResult interface {
		HasLabel() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasLabel()
	} else {
		return false
	}
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BTParameterLookupTableEntry1667) SetLabel(v string) {
	type getResult interface {
		SetLabel(v string)
	}

	o.GetActualInstance().(getResult).SetLabel(v)
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetLocalizableName() string {
	type getResult interface {
		GetLocalizableName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocalizableName()
	} else {
		var de string
		return de
	}
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetLocalizableNameOk() (*string, bool) {
	type getResult interface {
		GetLocalizableNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocalizableNameOk()
	} else {
		return nil, false
	}
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasLocalizableName() bool {
	type getResult interface {
		HasLocalizableName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasLocalizableName()
	} else {
		return false
	}
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTParameterLookupTableEntry1667) SetLocalizableName(v string) {
	type getResult interface {
		SetLocalizableName(v string)
	}

	o.GetActualInstance().(getResult).SetLocalizableName(v)
}

// GetLocalizedLabel returns the LocalizedLabel field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetLocalizedLabel() string {
	type getResult interface {
		GetLocalizedLabel() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocalizedLabel()
	} else {
		var de string
		return de
	}
}

// GetLocalizedLabelOk returns a tuple with the LocalizedLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetLocalizedLabelOk() (*string, bool) {
	type getResult interface {
		GetLocalizedLabelOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocalizedLabelOk()
	} else {
		return nil, false
	}
}

// HasLocalizedLabel returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasLocalizedLabel() bool {
	type getResult interface {
		HasLocalizedLabel() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasLocalizedLabel()
	} else {
		return false
	}
}

// SetLocalizedLabel gets a reference to the given string and assigns it to the LocalizedLabel field.
func (o *BTParameterLookupTableEntry1667) SetLocalizedLabel(v string) {
	type getResult interface {
		SetLocalizedLabel(v string)
	}

	o.GetActualInstance().(getResult).SetLocalizedLabel(v)
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetLocalizedName() string {
	type getResult interface {
		GetLocalizedName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocalizedName()
	} else {
		var de string
		return de
	}
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetLocalizedNameOk() (*string, bool) {
	type getResult interface {
		GetLocalizedNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetLocalizedNameOk()
	} else {
		return nil, false
	}
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasLocalizedName() bool {
	type getResult interface {
		HasLocalizedName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasLocalizedName()
	} else {
		return false
	}
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTParameterLookupTableEntry1667) SetLocalizedName(v string) {
	type getResult interface {
		SetLocalizedName(v string)
	}

	o.GetActualInstance().(getResult).SetLocalizedName(v)
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTParameterLookupTableEntry1667) GetStringsToLocalize() []string {
	type getResult interface {
		GetStringsToLocalize() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStringsToLocalize()
	} else {
		var de []string
		return de
	}
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterLookupTableEntry1667) GetStringsToLocalizeOk() ([]string, bool) {
	type getResult interface {
		GetStringsToLocalizeOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStringsToLocalizeOk()
	} else {
		return nil, false
	}
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTParameterLookupTableEntry1667) HasStringsToLocalize() bool {
	type getResult interface {
		HasStringsToLocalize() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasStringsToLocalize()
	} else {
		return false
	}
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTParameterLookupTableEntry1667) SetStringsToLocalize(v []string) {
	type getResult interface {
		SetStringsToLocalize(v []string)
	}

	o.GetActualInstance().(getResult).SetStringsToLocalize(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTParameterLookupTableEntry1667) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTParameterLookupTableListEntry-1916'
	if jsonDict["btType"] == "BTParameterLookupTableListEntry-1916" {
		// try to unmarshal JSON data into BTParameterLookupTableListEntry1916
		var qr *BTParameterLookupTableListEntry1916
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTParameterLookupTableEntry1667 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTParameterLookupTableEntry1667 = nil
			return fmt.Errorf("Failed to unmarshal BTParameterLookupTableEntry1667 as BTParameterLookupTableListEntry1916: %s", err.Error())
		}
	}

	var qtx *base_BTParameterLookupTableEntry1667
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTParameterLookupTableEntry1667 = qtx
		return nil // data stored in dst.base_BTParameterLookupTableEntry1667, return on the first match
	} else {
		dst.implBTParameterLookupTableEntry1667 = nil
		return fmt.Errorf("Failed to unmarshal BTParameterLookupTableEntry1667 as base_BTParameterLookupTableEntry1667: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTParameterLookupTableEntry1667) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTParameterLookupTableEntry1667) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTParameterLookupTableEntry1667
}

type NullableBTParameterLookupTableEntry1667 struct {
	value *BTParameterLookupTableEntry1667
	isSet bool
}

func (v NullableBTParameterLookupTableEntry1667) Get() *BTParameterLookupTableEntry1667 {
	return v.value
}

func (v *NullableBTParameterLookupTableEntry1667) Set(val *BTParameterLookupTableEntry1667) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterLookupTableEntry1667) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterLookupTableEntry1667) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterLookupTableEntry1667(val *BTParameterLookupTableEntry1667) *NullableBTParameterLookupTableEntry1667 {
	return &NullableBTParameterLookupTableEntry1667{value: val, isSet: true}
}

func (v NullableBTParameterLookupTableEntry1667) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterLookupTableEntry1667) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTParameterLookupTableEntry1667 struct {
	AdditionalLocalizedStrings *int32   `json:"additionalLocalizedStrings,omitempty"`
	BtType                     *string  `json:"btType,omitempty"`
	Label                      *string  `json:"label,omitempty"`
	LocalizableName            *string  `json:"localizableName,omitempty"`
	LocalizedLabel             *string  `json:"localizedLabel,omitempty"`
	LocalizedName              *string  `json:"localizedName,omitempty"`
	StringsToLocalize          []string `json:"stringsToLocalize,omitempty"`
}

// Newbase_BTParameterLookupTableEntry1667 instantiates a new base_BTParameterLookupTableEntry1667 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTParameterLookupTableEntry1667() *base_BTParameterLookupTableEntry1667 {
	this := base_BTParameterLookupTableEntry1667{}
	return &this
}

// Newbase_BTParameterLookupTableEntry1667WithDefaults instantiates a new base_BTParameterLookupTableEntry1667 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTParameterLookupTableEntry1667WithDefaults() *base_BTParameterLookupTableEntry1667 {
	this := base_BTParameterLookupTableEntry1667{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *base_BTParameterLookupTableEntry1667) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterLookupTableEntry1667) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *base_BTParameterLookupTableEntry1667) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *base_BTParameterLookupTableEntry1667) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTParameterLookupTableEntry1667) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterLookupTableEntry1667) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTParameterLookupTableEntry1667) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTParameterLookupTableEntry1667) SetBtType(v string) {
	o.BtType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *base_BTParameterLookupTableEntry1667) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterLookupTableEntry1667) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *base_BTParameterLookupTableEntry1667) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *base_BTParameterLookupTableEntry1667) SetLabel(v string) {
	o.Label = &v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *base_BTParameterLookupTableEntry1667) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterLookupTableEntry1667) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *base_BTParameterLookupTableEntry1667) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *base_BTParameterLookupTableEntry1667) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedLabel returns the LocalizedLabel field value if set, zero value otherwise.
func (o *base_BTParameterLookupTableEntry1667) GetLocalizedLabel() string {
	if o == nil || o.LocalizedLabel == nil {
		var ret string
		return ret
	}
	return *o.LocalizedLabel
}

// GetLocalizedLabelOk returns a tuple with the LocalizedLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterLookupTableEntry1667) GetLocalizedLabelOk() (*string, bool) {
	if o == nil || o.LocalizedLabel == nil {
		return nil, false
	}
	return o.LocalizedLabel, true
}

// HasLocalizedLabel returns a boolean if a field has been set.
func (o *base_BTParameterLookupTableEntry1667) HasLocalizedLabel() bool {
	if o != nil && o.LocalizedLabel != nil {
		return true
	}

	return false
}

// SetLocalizedLabel gets a reference to the given string and assigns it to the LocalizedLabel field.
func (o *base_BTParameterLookupTableEntry1667) SetLocalizedLabel(v string) {
	o.LocalizedLabel = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *base_BTParameterLookupTableEntry1667) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterLookupTableEntry1667) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *base_BTParameterLookupTableEntry1667) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *base_BTParameterLookupTableEntry1667) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *base_BTParameterLookupTableEntry1667) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTParameterLookupTableEntry1667) GetStringsToLocalizeOk() ([]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *base_BTParameterLookupTableEntry1667) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *base_BTParameterLookupTableEntry1667) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = v
}

func (o base_BTParameterLookupTableEntry1667) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalLocalizedStrings != nil {
		toSerialize["additionalLocalizedStrings"] = o.AdditionalLocalizedStrings
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LocalizableName != nil {
		toSerialize["localizableName"] = o.LocalizableName
	}
	if o.LocalizedLabel != nil {
		toSerialize["localizedLabel"] = o.LocalizedLabel
	}
	if o.LocalizedName != nil {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if o.StringsToLocalize != nil {
		toSerialize["stringsToLocalize"] = o.StringsToLocalize
	}
	return json.Marshal(toSerialize)
}
