/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6246-994505a8b5bf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTFullElementId756 - struct for BTFullElementId756
type BTFullElementId756 struct {
	implBTFullElementId756 interface{}
}

// BTFullElementIdWithDocument1729AsBTFullElementId756 is a convenience function that returns BTFullElementIdWithDocument1729 wrapped in BTFullElementId756
func (o *BTFullElementIdWithDocument1729) AsBTFullElementId756() *BTFullElementId756 {
	return &BTFullElementId756{o}
}

// NewBTFullElementId756 instantiates a new BTFullElementId756 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFullElementId756() *BTFullElementId756 {
	this := BTFullElementId756{Newbase_BTFullElementId756()}
	return &this
}

// NewBTFullElementId756WithDefaults instantiates a new BTFullElementId756 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFullElementId756WithDefaults() *BTFullElementId756 {
	this := BTFullElementId756{Newbase_BTFullElementId756WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFullElementId756) GetBtType() string {
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
func (o *BTFullElementId756) GetBtTypeOk() (*string, bool) {
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
func (o *BTFullElementId756) HasBtType() bool {
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
func (o *BTFullElementId756) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *BTFullElementId756) GetConfigured() bool {
	type getResult interface {
		GetConfigured() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConfigured()
	} else {
		var de bool
		return de
	}
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementId756) GetConfiguredOk() (*bool, bool) {
	type getResult interface {
		GetConfiguredOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetConfiguredOk()
	} else {
		return nil, false
	}
}

// HasConfigured returns a boolean if a field has been set.
func (o *BTFullElementId756) HasConfigured() bool {
	type getResult interface {
		HasConfigured() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasConfigured()
	} else {
		return false
	}
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *BTFullElementId756) SetConfigured(v bool) {
	type getResult interface {
		SetConfigured(v bool)
	}

	o.GetActualInstance().(getResult).SetConfigured(v)
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTFullElementId756) GetElementId() string {
	type getResult interface {
		GetElementId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementId()
	} else {
		var de string
		return de
	}
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementId756) GetElementIdOk() (*string, bool) {
	type getResult interface {
		GetElementIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementIdOk()
	} else {
		return nil, false
	}
}

// HasElementId returns a boolean if a field has been set.
func (o *BTFullElementId756) HasElementId() bool {
	type getResult interface {
		HasElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasElementId()
	} else {
		return false
	}
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTFullElementId756) SetElementId(v string) {
	type getResult interface {
		SetElementId(v string)
	}

	o.GetActualInstance().(getResult).SetElementId(v)
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTFullElementId756) GetMicroversionId() BTMicroversionId366 {
	type getResult interface {
		GetMicroversionId() BTMicroversionId366
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionId()
	} else {
		var de BTMicroversionId366
		return de
	}
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementId756) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	type getResult interface {
		GetMicroversionIdOk() (*BTMicroversionId366, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdOk()
	} else {
		return nil, false
	}
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTFullElementId756) HasMicroversionId() bool {
	type getResult interface {
		HasMicroversionId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversionId()
	} else {
		return false
	}
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTFullElementId756) SetMicroversionId(v BTMicroversionId366) {
	type getResult interface {
		SetMicroversionId(v BTMicroversionId366)
	}

	o.GetActualInstance().(getResult).SetMicroversionId(v)
}

// GetMicroversionIdAndConfiguration returns the MicroversionIdAndConfiguration field value if set, zero value otherwise.
func (o *BTFullElementId756) GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338 {
	type getResult interface {
		GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdAndConfiguration()
	} else {
		var de BTMicroversionIdAndConfiguration2338
		return de
	}
}

// GetMicroversionIdAndConfigurationOk returns a tuple with the MicroversionIdAndConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementId756) GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	type getResult interface {
		GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdAndConfigurationOk()
	} else {
		return nil, false
	}
}

// HasMicroversionIdAndConfiguration returns a boolean if a field has been set.
func (o *BTFullElementId756) HasMicroversionIdAndConfiguration() bool {
	type getResult interface {
		HasMicroversionIdAndConfiguration() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversionIdAndConfiguration()
	} else {
		return false
	}
}

// SetMicroversionIdAndConfiguration gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the MicroversionIdAndConfiguration field.
func (o *BTFullElementId756) SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338) {
	type getResult interface {
		SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338)
	}

	o.GetActualInstance().(getResult).SetMicroversionIdAndConfiguration(v)
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *BTFullElementId756) GetTarget() BTMicroversionIdAndConfiguration2338 {
	type getResult interface {
		GetTarget() BTMicroversionIdAndConfiguration2338
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTarget()
	} else {
		var de BTMicroversionIdAndConfiguration2338
		return de
	}
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementId756) GetTargetOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	type getResult interface {
		GetTargetOk() (*BTMicroversionIdAndConfiguration2338, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetTargetOk()
	} else {
		return nil, false
	}
}

// HasTarget returns a boolean if a field has been set.
func (o *BTFullElementId756) HasTarget() bool {
	type getResult interface {
		HasTarget() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasTarget()
	} else {
		return false
	}
}

// SetTarget gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the Target field.
func (o *BTFullElementId756) SetTarget(v BTMicroversionIdAndConfiguration2338) {
	type getResult interface {
		SetTarget(v BTMicroversionIdAndConfiguration2338)
	}

	o.GetActualInstance().(getResult).SetTarget(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTFullElementId756) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTFullElementIdWithDocument-1729'
	if jsonDict["btType"] == "BTFullElementIdWithDocument-1729" {
		// try to unmarshal JSON data into BTFullElementIdWithDocument1729
		var qr *BTFullElementIdWithDocument1729
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTFullElementId756 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTFullElementId756 = nil
			return fmt.Errorf("Failed to unmarshal BTFullElementId756 as BTFullElementIdWithDocument1729: %s", err.Error())
		}
	}

	var qtx *base_BTFullElementId756
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTFullElementId756 = qtx
		return nil // data stored in dst.base_BTFullElementId756, return on the first match
	} else {
		dst.implBTFullElementId756 = nil
		return fmt.Errorf("Failed to unmarshal BTFullElementId756 as base_BTFullElementId756: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTFullElementId756) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTFullElementId756) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTFullElementId756
}

type NullableBTFullElementId756 struct {
	value *BTFullElementId756
	isSet bool
}

func (v NullableBTFullElementId756) Get() *BTFullElementId756 {
	return v.value
}

func (v *NullableBTFullElementId756) Set(val *BTFullElementId756) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFullElementId756) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFullElementId756) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFullElementId756(val *BTFullElementId756) *NullableBTFullElementId756 {
	return &NullableBTFullElementId756{value: val, isSet: true}
}

func (v NullableBTFullElementId756) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFullElementId756) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTFullElementId756 struct {
	BtType                         *string                               `json:"btType,omitempty"`
	Configured                     *bool                                 `json:"configured,omitempty"`
	ElementId                      *string                               `json:"elementId,omitempty"`
	MicroversionId                 *BTMicroversionId366                  `json:"microversionId,omitempty"`
	MicroversionIdAndConfiguration *BTMicroversionIdAndConfiguration2338 `json:"microversionIdAndConfiguration,omitempty"`
	Target                         *BTMicroversionIdAndConfiguration2338 `json:"target,omitempty"`
}

// Newbase_BTFullElementId756 instantiates a new base_BTFullElementId756 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTFullElementId756() *base_BTFullElementId756 {
	this := base_BTFullElementId756{}
	return &this
}

// Newbase_BTFullElementId756WithDefaults instantiates a new base_BTFullElementId756 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTFullElementId756WithDefaults() *base_BTFullElementId756 {
	this := base_BTFullElementId756{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTFullElementId756) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementId756) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTFullElementId756) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTFullElementId756) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *base_BTFullElementId756) GetConfigured() bool {
	if o == nil || o.Configured == nil {
		var ret bool
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementId756) GetConfiguredOk() (*bool, bool) {
	if o == nil || o.Configured == nil {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *base_BTFullElementId756) HasConfigured() bool {
	if o != nil && o.Configured != nil {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *base_BTFullElementId756) SetConfigured(v bool) {
	o.Configured = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *base_BTFullElementId756) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementId756) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *base_BTFullElementId756) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *base_BTFullElementId756) SetElementId(v string) {
	o.ElementId = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *base_BTFullElementId756) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementId756) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *base_BTFullElementId756) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *base_BTFullElementId756) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetMicroversionIdAndConfiguration returns the MicroversionIdAndConfiguration field value if set, zero value otherwise.
func (o *base_BTFullElementId756) GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.MicroversionIdAndConfiguration
}

// GetMicroversionIdAndConfigurationOk returns a tuple with the MicroversionIdAndConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementId756) GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfiguration, true
}

// HasMicroversionIdAndConfiguration returns a boolean if a field has been set.
func (o *base_BTFullElementId756) HasMicroversionIdAndConfiguration() bool {
	if o != nil && o.MicroversionIdAndConfiguration != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfiguration gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the MicroversionIdAndConfiguration field.
func (o *base_BTFullElementId756) SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338) {
	o.MicroversionIdAndConfiguration = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *base_BTFullElementId756) GetTarget() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.Target == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTFullElementId756) GetTargetOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *base_BTFullElementId756) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the Target field.
func (o *base_BTFullElementId756) SetTarget(v BTMicroversionIdAndConfiguration2338) {
	o.Target = &v
}

func (o base_BTFullElementId756) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Configured != nil {
		toSerialize["configured"] = o.Configured
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MicroversionIdAndConfiguration != nil {
		toSerialize["microversionIdAndConfiguration"] = o.MicroversionIdAndConfiguration
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}
