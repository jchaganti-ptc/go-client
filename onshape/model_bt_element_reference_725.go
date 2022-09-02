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

// BTElementReference725 - struct for BTElementReference725
type BTElementReference725 struct {
	implBTElementReference725 interface{}
}

// BTExternalReference1936AsBTElementReference725 is a convenience function that returns BTExternalReference1936 wrapped in BTElementReference725
func (o *BTExternalReference1936) AsBTElementReference725() *BTElementReference725 {
	return &BTElementReference725{o}
}

// NewBTElementReference725 instantiates a new BTElementReference725 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTElementReference725() *BTElementReference725 {
	this := BTElementReference725{Newbase_BTElementReference725()}
	return &this
}

// NewBTElementReference725WithDefaults instantiates a new BTElementReference725 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTElementReference725WithDefaults() *BTElementReference725 {
	this := BTElementReference725{Newbase_BTElementReference725WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTElementReference725) GetBtType() string {
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
func (o *BTElementReference725) GetBtTypeOk() (*string, bool) {
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
func (o *BTElementReference725) HasBtType() bool {
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
func (o *BTElementReference725) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *BTElementReference725) GetConfigured() bool {
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
func (o *BTElementReference725) GetConfiguredOk() (*bool, bool) {
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
func (o *BTElementReference725) HasConfigured() bool {
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
func (o *BTElementReference725) SetConfigured(v bool) {
	type getResult interface {
		SetConfigured(v bool)
	}

	o.GetActualInstance().(getResult).SetConfigured(v)
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTElementReference725) GetElementId() string {
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
func (o *BTElementReference725) GetElementIdOk() (*string, bool) {
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
func (o *BTElementReference725) HasElementId() bool {
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
func (o *BTElementReference725) SetElementId(v string) {
	type getResult interface {
		SetElementId(v string)
	}

	o.GetActualInstance().(getResult).SetElementId(v)
}

// GetExternalDocumentWithVersion returns the ExternalDocumentWithVersion field value if set, zero value otherwise.
func (o *BTElementReference725) GetExternalDocumentWithVersion() BTDocumentWithVersionId {
	type getResult interface {
		GetExternalDocumentWithVersion() BTDocumentWithVersionId
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExternalDocumentWithVersion()
	} else {
		var de BTDocumentWithVersionId
		return de
	}
}

// GetExternalDocumentWithVersionOk returns a tuple with the ExternalDocumentWithVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementReference725) GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool) {
	type getResult interface {
		GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExternalDocumentWithVersionOk()
	} else {
		return nil, false
	}
}

// HasExternalDocumentWithVersion returns a boolean if a field has been set.
func (o *BTElementReference725) HasExternalDocumentWithVersion() bool {
	type getResult interface {
		HasExternalDocumentWithVersion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasExternalDocumentWithVersion()
	} else {
		return false
	}
}

// SetExternalDocumentWithVersion gets a reference to the given BTDocumentWithVersionId and assigns it to the ExternalDocumentWithVersion field.
func (o *BTElementReference725) SetExternalDocumentWithVersion(v BTDocumentWithVersionId) {
	type getResult interface {
		SetExternalDocumentWithVersion(v BTDocumentWithVersionId)
	}

	o.GetActualInstance().(getResult).SetExternalDocumentWithVersion(v)
}

// GetExternalDocumentWithVersionAndElementId returns the ExternalDocumentWithVersionAndElementId field value if set, zero value otherwise.
func (o *BTElementReference725) GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId {
	type getResult interface {
		GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExternalDocumentWithVersionAndElementId()
	} else {
		var de BTDocumentWithVersionAndElementId
		return de
	}
}

// GetExternalDocumentWithVersionAndElementIdOk returns a tuple with the ExternalDocumentWithVersionAndElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementReference725) GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool) {
	type getResult interface {
		GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExternalDocumentWithVersionAndElementIdOk()
	} else {
		return nil, false
	}
}

// HasExternalDocumentWithVersionAndElementId returns a boolean if a field has been set.
func (o *BTElementReference725) HasExternalDocumentWithVersionAndElementId() bool {
	type getResult interface {
		HasExternalDocumentWithVersionAndElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasExternalDocumentWithVersionAndElementId()
	} else {
		return false
	}
}

// SetExternalDocumentWithVersionAndElementId gets a reference to the given BTDocumentWithVersionAndElementId and assigns it to the ExternalDocumentWithVersionAndElementId field.
func (o *BTElementReference725) SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId) {
	type getResult interface {
		SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId)
	}

	o.GetActualInstance().(getResult).SetExternalDocumentWithVersionAndElementId(v)
}

// GetExternalReference returns the ExternalReference field value if set, zero value otherwise.
func (o *BTElementReference725) GetExternalReference() bool {
	type getResult interface {
		GetExternalReference() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExternalReference()
	} else {
		var de bool
		return de
	}
}

// GetExternalReferenceOk returns a tuple with the ExternalReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementReference725) GetExternalReferenceOk() (*bool, bool) {
	type getResult interface {
		GetExternalReferenceOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExternalReferenceOk()
	} else {
		return nil, false
	}
}

// HasExternalReference returns a boolean if a field has been set.
func (o *BTElementReference725) HasExternalReference() bool {
	type getResult interface {
		HasExternalReference() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasExternalReference()
	} else {
		return false
	}
}

// SetExternalReference gets a reference to the given bool and assigns it to the ExternalReference field.
func (o *BTElementReference725) SetExternalReference(v bool) {
	type getResult interface {
		SetExternalReference(v bool)
	}

	o.GetActualInstance().(getResult).SetExternalReference(v)
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTElementReference725) GetFullElementId() BTFullElementId756 {
	type getResult interface {
		GetFullElementId() BTFullElementId756
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullElementId()
	} else {
		var de BTFullElementId756
		return de
	}
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementReference725) GetFullElementIdOk() (*BTFullElementId756, bool) {
	type getResult interface {
		GetFullElementIdOk() (*BTFullElementId756, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullElementIdOk()
	} else {
		return nil, false
	}
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTElementReference725) HasFullElementId() bool {
	type getResult interface {
		HasFullElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFullElementId()
	} else {
		return false
	}
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTElementReference725) SetFullElementId(v BTFullElementId756) {
	type getResult interface {
		SetFullElementId(v BTFullElementId756)
	}

	o.GetActualInstance().(getResult).SetFullElementId(v)
}

// GetMicroversionIdAndConfiguration returns the MicroversionIdAndConfiguration field value if set, zero value otherwise.
func (o *BTElementReference725) GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338 {
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
func (o *BTElementReference725) GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool) {
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
func (o *BTElementReference725) HasMicroversionIdAndConfiguration() bool {
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
func (o *BTElementReference725) SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338) {
	type getResult interface {
		SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338)
	}

	o.GetActualInstance().(getResult).SetMicroversionIdAndConfiguration(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTElementReference725) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementReference725) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTElementReference725) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTElementReference725) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTElementReference725) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTExternalReference-1936'
	if jsonDict["btType"] == "BTExternalReference-1936" {
		// try to unmarshal JSON data into BTExternalReference1936
		var qr *BTExternalReference1936
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTElementReference725 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTElementReference725 = nil
			return fmt.Errorf("Failed to unmarshal BTElementReference725 as BTExternalReference1936: %s", err.Error())
		}
	}

	var qtx *base_BTElementReference725
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTElementReference725 = qtx
		return nil // data stored in dst.base_BTElementReference725, return on the first match
	} else {
		dst.implBTElementReference725 = nil
		return fmt.Errorf("Failed to unmarshal BTElementReference725 as base_BTElementReference725: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTElementReference725) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTElementReference725) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTElementReference725
}

type NullableBTElementReference725 struct {
	value *BTElementReference725
	isSet bool
}

func (v NullableBTElementReference725) Get() *BTElementReference725 {
	return v.value
}

func (v *NullableBTElementReference725) Set(val *BTElementReference725) {
	v.value = val
	v.isSet = true
}

func (v NullableBTElementReference725) IsSet() bool {
	return v.isSet
}

func (v *NullableBTElementReference725) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTElementReference725(val *BTElementReference725) *NullableBTElementReference725 {
	return &NullableBTElementReference725{value: val, isSet: true}
}

func (v NullableBTElementReference725) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTElementReference725) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTElementReference725 struct {
	BtType                                  *string                               `json:"btType,omitempty"`
	Configured                              *bool                                 `json:"configured,omitempty"`
	ElementId                               *string                               `json:"elementId,omitempty"`
	ExternalDocumentWithVersion             *BTDocumentWithVersionId              `json:"externalDocumentWithVersion,omitempty"`
	ExternalDocumentWithVersionAndElementId *BTDocumentWithVersionAndElementId    `json:"externalDocumentWithVersionAndElementId,omitempty"`
	ExternalReference                       *bool                                 `json:"externalReference,omitempty"`
	FullElementId                           *BTFullElementId756                   `json:"fullElementId,omitempty"`
	MicroversionIdAndConfiguration          *BTMicroversionIdAndConfiguration2338 `json:"microversionIdAndConfiguration,omitempty"`
	NodeId                                  *string                               `json:"nodeId,omitempty"`
}

// Newbase_BTElementReference725 instantiates a new base_BTElementReference725 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTElementReference725() *base_BTElementReference725 {
	this := base_BTElementReference725{}
	return &this
}

// Newbase_BTElementReference725WithDefaults instantiates a new base_BTElementReference725 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTElementReference725WithDefaults() *base_BTElementReference725 {
	this := base_BTElementReference725{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTElementReference725) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetConfigured() bool {
	if o == nil || o.Configured == nil {
		var ret bool
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetConfiguredOk() (*bool, bool) {
	if o == nil || o.Configured == nil {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasConfigured() bool {
	if o != nil && o.Configured != nil {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *base_BTElementReference725) SetConfigured(v bool) {
	o.Configured = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *base_BTElementReference725) SetElementId(v string) {
	o.ElementId = &v
}

// GetExternalDocumentWithVersion returns the ExternalDocumentWithVersion field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetExternalDocumentWithVersion() BTDocumentWithVersionId {
	if o == nil || o.ExternalDocumentWithVersion == nil {
		var ret BTDocumentWithVersionId
		return ret
	}
	return *o.ExternalDocumentWithVersion
}

// GetExternalDocumentWithVersionOk returns a tuple with the ExternalDocumentWithVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetExternalDocumentWithVersionOk() (*BTDocumentWithVersionId, bool) {
	if o == nil || o.ExternalDocumentWithVersion == nil {
		return nil, false
	}
	return o.ExternalDocumentWithVersion, true
}

// HasExternalDocumentWithVersion returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasExternalDocumentWithVersion() bool {
	if o != nil && o.ExternalDocumentWithVersion != nil {
		return true
	}

	return false
}

// SetExternalDocumentWithVersion gets a reference to the given BTDocumentWithVersionId and assigns it to the ExternalDocumentWithVersion field.
func (o *base_BTElementReference725) SetExternalDocumentWithVersion(v BTDocumentWithVersionId) {
	o.ExternalDocumentWithVersion = &v
}

// GetExternalDocumentWithVersionAndElementId returns the ExternalDocumentWithVersionAndElementId field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetExternalDocumentWithVersionAndElementId() BTDocumentWithVersionAndElementId {
	if o == nil || o.ExternalDocumentWithVersionAndElementId == nil {
		var ret BTDocumentWithVersionAndElementId
		return ret
	}
	return *o.ExternalDocumentWithVersionAndElementId
}

// GetExternalDocumentWithVersionAndElementIdOk returns a tuple with the ExternalDocumentWithVersionAndElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetExternalDocumentWithVersionAndElementIdOk() (*BTDocumentWithVersionAndElementId, bool) {
	if o == nil || o.ExternalDocumentWithVersionAndElementId == nil {
		return nil, false
	}
	return o.ExternalDocumentWithVersionAndElementId, true
}

// HasExternalDocumentWithVersionAndElementId returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasExternalDocumentWithVersionAndElementId() bool {
	if o != nil && o.ExternalDocumentWithVersionAndElementId != nil {
		return true
	}

	return false
}

// SetExternalDocumentWithVersionAndElementId gets a reference to the given BTDocumentWithVersionAndElementId and assigns it to the ExternalDocumentWithVersionAndElementId field.
func (o *base_BTElementReference725) SetExternalDocumentWithVersionAndElementId(v BTDocumentWithVersionAndElementId) {
	o.ExternalDocumentWithVersionAndElementId = &v
}

// GetExternalReference returns the ExternalReference field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetExternalReference() bool {
	if o == nil || o.ExternalReference == nil {
		var ret bool
		return ret
	}
	return *o.ExternalReference
}

// GetExternalReferenceOk returns a tuple with the ExternalReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetExternalReferenceOk() (*bool, bool) {
	if o == nil || o.ExternalReference == nil {
		return nil, false
	}
	return o.ExternalReference, true
}

// HasExternalReference returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasExternalReference() bool {
	if o != nil && o.ExternalReference != nil {
		return true
	}

	return false
}

// SetExternalReference gets a reference to the given bool and assigns it to the ExternalReference field.
func (o *base_BTElementReference725) SetExternalReference(v bool) {
	o.ExternalReference = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *base_BTElementReference725) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetMicroversionIdAndConfiguration returns the MicroversionIdAndConfiguration field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.MicroversionIdAndConfiguration
}

// GetMicroversionIdAndConfigurationOk returns a tuple with the MicroversionIdAndConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfiguration, true
}

// HasMicroversionIdAndConfiguration returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasMicroversionIdAndConfiguration() bool {
	if o != nil && o.MicroversionIdAndConfiguration != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfiguration gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the MicroversionIdAndConfiguration field.
func (o *base_BTElementReference725) SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338) {
	o.MicroversionIdAndConfiguration = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTElementReference725) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementReference725) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTElementReference725) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTElementReference725) SetNodeId(v string) {
	o.NodeId = &v
}

func (o base_BTElementReference725) MarshalJSON() ([]byte, error) {
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
	if o.ExternalDocumentWithVersion != nil {
		toSerialize["externalDocumentWithVersion"] = o.ExternalDocumentWithVersion
	}
	if o.ExternalDocumentWithVersionAndElementId != nil {
		toSerialize["externalDocumentWithVersionAndElementId"] = o.ExternalDocumentWithVersionAndElementId
	}
	if o.ExternalReference != nil {
		toSerialize["externalReference"] = o.ExternalReference
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.MicroversionIdAndConfiguration != nil {
		toSerialize["microversionIdAndConfiguration"] = o.MicroversionIdAndConfiguration
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	return json.Marshal(toSerialize)
}
