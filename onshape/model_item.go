/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5928-bd774e9c9f97
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// Item - struct for Item
type Item struct {
	implItem interface{}
}

// BlobItemAsItem is a convenience function that returns BlobItem wrapped in Item
func (o *BlobItem) AsItem() *Item {
	return &Item{o}
}

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem(jsonType string) *Item {
	this := Item{Newbase_Item(jsonType)}
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{Newbase_ItemWithDefaults()}
	return &this
}

// GetBaseHref returns the BaseHref field value if set, zero value otherwise.
func (o *Item) GetBaseHref() string {
	type getResult interface {
		GetBaseHref() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBaseHref()
	} else {
		var de string
		return de
	}
}

// GetBaseHrefOk returns a tuple with the BaseHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetBaseHrefOk() (*string, bool) {
	type getResult interface {
		GetBaseHrefOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBaseHrefOk()
	} else {
		return nil, false
	}
}

// HasBaseHref returns a boolean if a field has been set.
func (o *Item) HasBaseHref() bool {
	type getResult interface {
		HasBaseHref() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBaseHref()
	} else {
		return false
	}
}

// SetBaseHref gets a reference to the given string and assigns it to the BaseHref field.
func (o *Item) SetBaseHref(v string) {
	type getResult interface {
		SetBaseHref(v string)
	}

	o.GetActualInstance().(getResult).SetBaseHref(v)
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *Item) GetDataType() string {
	type getResult interface {
		GetDataType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDataType()
	} else {
		var de string
		return de
	}
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDataTypeOk() (*string, bool) {
	type getResult interface {
		GetDataTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDataTypeOk()
	} else {
		return nil, false
	}
}

// HasDataType returns a boolean if a field has been set.
func (o *Item) HasDataType() bool {
	type getResult interface {
		HasDataType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDataType()
	} else {
		return false
	}
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *Item) SetDataType(v string) {
	type getResult interface {
		SetDataType(v string)
	}

	o.GetActualInstance().(getResult).SetDataType(v)
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *Item) GetDocumentId() string {
	type getResult interface {
		GetDocumentId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentId()
	} else {
		var de string
		return de
	}
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetDocumentIdOk() (*string, bool) {
	type getResult interface {
		GetDocumentIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDocumentIdOk()
	} else {
		return nil, false
	}
}

// HasDocumentId returns a boolean if a field has been set.
func (o *Item) HasDocumentId() bool {
	type getResult interface {
		HasDocumentId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDocumentId()
	} else {
		return false
	}
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *Item) SetDocumentId(v string) {
	type getResult interface {
		SetDocumentId(v string)
	}

	o.GetActualInstance().(getResult).SetDocumentId(v)
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *Item) GetElementId() string {
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
func (o *Item) GetElementIdOk() (*string, bool) {
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
func (o *Item) HasElementId() bool {
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
func (o *Item) SetElementId(v string) {
	type getResult interface {
		SetElementId(v string)
	}

	o.GetActualInstance().(getResult).SetElementId(v)
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *Item) GetElementType() string {
	type getResult interface {
		GetElementType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementType()
	} else {
		var de string
		return de
	}
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetElementTypeOk() (*string, bool) {
	type getResult interface {
		GetElementTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementTypeOk()
	} else {
		return nil, false
	}
}

// HasElementType returns a boolean if a field has been set.
func (o *Item) HasElementType() bool {
	type getResult interface {
		HasElementType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasElementType()
	} else {
		return false
	}
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *Item) SetElementType(v string) {
	type getResult interface {
		SetElementType(v string)
	}

	o.GetActualInstance().(getResult).SetElementType(v)
}

// GetEncodedConfiguration returns the EncodedConfiguration field value if set, zero value otherwise.
func (o *Item) GetEncodedConfiguration() string {
	type getResult interface {
		GetEncodedConfiguration() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEncodedConfiguration()
	} else {
		var de string
		return de
	}
}

// GetEncodedConfigurationOk returns a tuple with the EncodedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetEncodedConfigurationOk() (*string, bool) {
	type getResult interface {
		GetEncodedConfigurationOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEncodedConfigurationOk()
	} else {
		return nil, false
	}
}

// HasEncodedConfiguration returns a boolean if a field has been set.
func (o *Item) HasEncodedConfiguration() bool {
	type getResult interface {
		HasEncodedConfiguration() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEncodedConfiguration()
	} else {
		return false
	}
}

// SetEncodedConfiguration gets a reference to the given string and assigns it to the EncodedConfiguration field.
func (o *Item) SetEncodedConfiguration(v string) {
	type getResult interface {
		SetEncodedConfiguration(v string)
	}

	o.GetActualInstance().(getResult).SetEncodedConfiguration(v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Item) GetId() string {
	type getResult interface {
		GetId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetId()
	} else {
		var de string
		return de
	}
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetIdOk() (*string, bool) {
	type getResult interface {
		GetIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIdOk()
	} else {
		return nil, false
	}
}

// HasId returns a boolean if a field has been set.
func (o *Item) HasId() bool {
	type getResult interface {
		HasId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasId()
	} else {
		return false
	}
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Item) SetId(v string) {
	type getResult interface {
		SetId(v string)
	}

	o.GetActualInstance().(getResult).SetId(v)
}

// GetJsonType returns the JsonType field value
func (o *Item) GetJsonType() string {
	type getResult interface {
		GetJsonType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetJsonType()
	} else {
		var de string
		return de
	}
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *Item) GetJsonTypeOk() (*string, bool) {
	type getResult interface {
		GetJsonTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetJsonTypeOk()
	} else {
		return nil, false
	}
}

// SetJsonType sets field value
func (o *Item) SetJsonType(v string) {
	type getResult interface {
		SetJsonType(v string)
	}

	o.GetActualInstance().(getResult).SetJsonType(v)
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *Item) GetPartId() string {
	type getResult interface {
		GetPartId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPartId()
	} else {
		var de string
		return de
	}
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetPartIdOk() (*string, bool) {
	type getResult interface {
		GetPartIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPartIdOk()
	} else {
		return nil, false
	}
}

// HasPartId returns a boolean if a field has been set.
func (o *Item) HasPartId() bool {
	type getResult interface {
		HasPartId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPartId()
	} else {
		return false
	}
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *Item) SetPartId(v string) {
	type getResult interface {
		SetPartId(v string)
	}

	o.GetActualInstance().(getResult).SetPartId(v)
}

// GetPartName returns the PartName field value if set, zero value otherwise.
func (o *Item) GetPartName() string {
	type getResult interface {
		GetPartName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPartName()
	} else {
		var de string
		return de
	}
}

// GetPartNameOk returns a tuple with the PartName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetPartNameOk() (*string, bool) {
	type getResult interface {
		GetPartNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPartNameOk()
	} else {
		return nil, false
	}
}

// HasPartName returns a boolean if a field has been set.
func (o *Item) HasPartName() bool {
	type getResult interface {
		HasPartName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPartName()
	} else {
		return false
	}
}

// SetPartName gets a reference to the given string and assigns it to the PartName field.
func (o *Item) SetPartName(v string) {
	type getResult interface {
		SetPartName(v string)
	}

	o.GetActualInstance().(getResult).SetPartName(v)
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *Item) GetPartNumber() string {
	type getResult interface {
		GetPartNumber() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPartNumber()
	} else {
		var de string
		return de
	}
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetPartNumberOk() (*string, bool) {
	type getResult interface {
		GetPartNumberOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPartNumberOk()
	} else {
		return nil, false
	}
}

// HasPartNumber returns a boolean if a field has been set.
func (o *Item) HasPartNumber() bool {
	type getResult interface {
		HasPartNumber() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPartNumber()
	} else {
		return false
	}
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *Item) SetPartNumber(v string) {
	type getResult interface {
		SetPartNumber(v string)
	}

	o.GetActualInstance().(getResult).SetPartNumber(v)
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *Item) GetRevision() string {
	type getResult interface {
		GetRevision() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRevision()
	} else {
		var de string
		return de
	}
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetRevisionOk() (*string, bool) {
	type getResult interface {
		GetRevisionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRevisionOk()
	} else {
		return nil, false
	}
}

// HasRevision returns a boolean if a field has been set.
func (o *Item) HasRevision() bool {
	type getResult interface {
		HasRevision() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRevision()
	} else {
		return false
	}
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *Item) SetRevision(v string) {
	type getResult interface {
		SetRevision(v string)
	}

	o.GetActualInstance().(getResult).SetRevision(v)
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *Item) GetRevisionId() string {
	type getResult interface {
		GetRevisionId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRevisionId()
	} else {
		var de string
		return de
	}
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetRevisionIdOk() (*string, bool) {
	type getResult interface {
		GetRevisionIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRevisionIdOk()
	} else {
		return nil, false
	}
}

// HasRevisionId returns a boolean if a field has been set.
func (o *Item) HasRevisionId() bool {
	type getResult interface {
		HasRevisionId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRevisionId()
	} else {
		return false
	}
}

// SetRevisionId gets a reference to the given string and assigns it to the RevisionId field.
func (o *Item) SetRevisionId(v string) {
	type getResult interface {
		SetRevisionId(v string)
	}

	o.GetActualInstance().(getResult).SetRevisionId(v)
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Item) GetState() int32 {
	type getResult interface {
		GetState() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetState()
	} else {
		var de int32
		return de
	}
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetStateOk() (*int32, bool) {
	type getResult interface {
		GetStateOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetStateOk()
	} else {
		return nil, false
	}
}

// HasState returns a boolean if a field has been set.
func (o *Item) HasState() bool {
	type getResult interface {
		HasState() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasState()
	} else {
		return false
	}
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *Item) SetState(v int32) {
	type getResult interface {
		SetState(v int32)
	}

	o.GetActualInstance().(getResult).SetState(v)
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *Item) GetVersionId() string {
	type getResult interface {
		GetVersionId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVersionId()
	} else {
		var de string
		return de
	}
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetVersionIdOk() (*string, bool) {
	type getResult interface {
		GetVersionIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVersionIdOk()
	} else {
		return nil, false
	}
}

// HasVersionId returns a boolean if a field has been set.
func (o *Item) HasVersionId() bool {
	type getResult interface {
		HasVersionId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasVersionId()
	} else {
		return false
	}
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *Item) SetVersionId(v string) {
	type getResult interface {
		SetVersionId(v string)
	}

	o.GetActualInstance().(getResult).SetVersionId(v)
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *Item) GetVersionName() string {
	type getResult interface {
		GetVersionName() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVersionName()
	} else {
		var de string
		return de
	}
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetVersionNameOk() (*string, bool) {
	type getResult interface {
		GetVersionNameOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVersionNameOk()
	} else {
		return nil, false
	}
}

// HasVersionName returns a boolean if a field has been set.
func (o *Item) HasVersionName() bool {
	type getResult interface {
		HasVersionName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasVersionName()
	} else {
		return false
	}
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *Item) SetVersionName(v string) {
	type getResult interface {
		SetVersionName(v string)
	}

	o.GetActualInstance().(getResult).SetVersionName(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Item) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BlobItem'
	if jsonDict["jsonType"] == "BlobItem" {
		// try to unmarshal JSON data into BlobItem
		var qr *BlobItem
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implItem = qr
			return nil // data stored, return on the first match
		} else {
			dst.implItem = nil
			return fmt.Errorf("Failed to unmarshal Item as BlobItem: %s", err.Error())
		}
	}

	// check if the discriminator value is 'publication-blob-item'
	if jsonDict["jsonType"] == "publication-blob-item" {
		// try to unmarshal JSON data into BlobItem
		var qr *BlobItem
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implItem = qr
			return nil // data stored, return on the first match
		} else {
			dst.implItem = nil
			return fmt.Errorf("Failed to unmarshal Item as BlobItem: %s", err.Error())
		}
	}

	var qtx *base_Item
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implItem = qtx
		return nil // data stored in dst.base_Item, return on the first match
	} else {
		dst.implItem = nil
		return fmt.Errorf("Failed to unmarshal Item as base_Item: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Item) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *Item) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implItem
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_Item struct {
	BaseHref             *string `json:"baseHref,omitempty"`
	DataType             *string `json:"dataType,omitempty"`
	DocumentId           *string `json:"documentId,omitempty"`
	ElementId            *string `json:"elementId,omitempty"`
	ElementType          *string `json:"elementType,omitempty"`
	EncodedConfiguration *string `json:"encodedConfiguration,omitempty"`
	Id                   *string `json:"id,omitempty"`
	JsonType             string  `json:"jsonType"`
	PartId               *string `json:"partId,omitempty"`
	PartName             *string `json:"partName,omitempty"`
	PartNumber           *string `json:"partNumber,omitempty"`
	Revision             *string `json:"revision,omitempty"`
	RevisionId           *string `json:"revisionId,omitempty"`
	State                *int32  `json:"state,omitempty"`
	VersionId            *string `json:"versionId,omitempty"`
	VersionName          *string `json:"versionName,omitempty"`
}

// Newbase_Item instantiates a new base_Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_Item(jsonType string) *base_Item {
	this := base_Item{}
	this.JsonType = jsonType
	return &this
}

// Newbase_ItemWithDefaults instantiates a new base_Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_ItemWithDefaults() *base_Item {
	this := base_Item{}
	return &this
}

// GetBaseHref returns the BaseHref field value if set, zero value otherwise.
func (o *base_Item) GetBaseHref() string {
	if o == nil || o.BaseHref == nil {
		var ret string
		return ret
	}
	return *o.BaseHref
}

// GetBaseHrefOk returns a tuple with the BaseHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetBaseHrefOk() (*string, bool) {
	if o == nil || o.BaseHref == nil {
		return nil, false
	}
	return o.BaseHref, true
}

// HasBaseHref returns a boolean if a field has been set.
func (o *base_Item) HasBaseHref() bool {
	if o != nil && o.BaseHref != nil {
		return true
	}

	return false
}

// SetBaseHref gets a reference to the given string and assigns it to the BaseHref field.
func (o *base_Item) SetBaseHref(v string) {
	o.BaseHref = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *base_Item) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *base_Item) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *base_Item) SetDataType(v string) {
	o.DataType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *base_Item) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *base_Item) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *base_Item) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *base_Item) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *base_Item) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *base_Item) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *base_Item) GetElementType() string {
	if o == nil || o.ElementType == nil {
		var ret string
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetElementTypeOk() (*string, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *base_Item) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *base_Item) SetElementType(v string) {
	o.ElementType = &v
}

// GetEncodedConfiguration returns the EncodedConfiguration field value if set, zero value otherwise.
func (o *base_Item) GetEncodedConfiguration() string {
	if o == nil || o.EncodedConfiguration == nil {
		var ret string
		return ret
	}
	return *o.EncodedConfiguration
}

// GetEncodedConfigurationOk returns a tuple with the EncodedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetEncodedConfigurationOk() (*string, bool) {
	if o == nil || o.EncodedConfiguration == nil {
		return nil, false
	}
	return o.EncodedConfiguration, true
}

// HasEncodedConfiguration returns a boolean if a field has been set.
func (o *base_Item) HasEncodedConfiguration() bool {
	if o != nil && o.EncodedConfiguration != nil {
		return true
	}

	return false
}

// SetEncodedConfiguration gets a reference to the given string and assigns it to the EncodedConfiguration field.
func (o *base_Item) SetEncodedConfiguration(v string) {
	o.EncodedConfiguration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *base_Item) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *base_Item) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *base_Item) SetId(v string) {
	o.Id = &v
}

// GetJsonType returns the JsonType field value
func (o *base_Item) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *base_Item) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *base_Item) SetJsonType(v string) {
	o.JsonType = v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *base_Item) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *base_Item) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *base_Item) SetPartId(v string) {
	o.PartId = &v
}

// GetPartName returns the PartName field value if set, zero value otherwise.
func (o *base_Item) GetPartName() string {
	if o == nil || o.PartName == nil {
		var ret string
		return ret
	}
	return *o.PartName
}

// GetPartNameOk returns a tuple with the PartName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetPartNameOk() (*string, bool) {
	if o == nil || o.PartName == nil {
		return nil, false
	}
	return o.PartName, true
}

// HasPartName returns a boolean if a field has been set.
func (o *base_Item) HasPartName() bool {
	if o != nil && o.PartName != nil {
		return true
	}

	return false
}

// SetPartName gets a reference to the given string and assigns it to the PartName field.
func (o *base_Item) SetPartName(v string) {
	o.PartName = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *base_Item) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *base_Item) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *base_Item) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *base_Item) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *base_Item) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *base_Item) SetRevision(v string) {
	o.Revision = &v
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *base_Item) GetRevisionId() string {
	if o == nil || o.RevisionId == nil {
		var ret string
		return ret
	}
	return *o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetRevisionIdOk() (*string, bool) {
	if o == nil || o.RevisionId == nil {
		return nil, false
	}
	return o.RevisionId, true
}

// HasRevisionId returns a boolean if a field has been set.
func (o *base_Item) HasRevisionId() bool {
	if o != nil && o.RevisionId != nil {
		return true
	}

	return false
}

// SetRevisionId gets a reference to the given string and assigns it to the RevisionId field.
func (o *base_Item) SetRevisionId(v string) {
	o.RevisionId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *base_Item) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *base_Item) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *base_Item) SetState(v int32) {
	o.State = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *base_Item) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *base_Item) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *base_Item) SetVersionId(v string) {
	o.VersionId = &v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *base_Item) GetVersionName() string {
	if o == nil || o.VersionName == nil {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_Item) GetVersionNameOk() (*string, bool) {
	if o == nil || o.VersionName == nil {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *base_Item) HasVersionName() bool {
	if o != nil && o.VersionName != nil {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *base_Item) SetVersionName(v string) {
	o.VersionName = &v
}

func (o base_Item) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseHref != nil {
		toSerialize["baseHref"] = o.BaseHref
	}
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.EncodedConfiguration != nil {
		toSerialize["encodedConfiguration"] = o.EncodedConfiguration
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartName != nil {
		toSerialize["partName"] = o.PartName
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.RevisionId != nil {
		toSerialize["revisionId"] = o.RevisionId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.VersionName != nil {
		toSerialize["versionName"] = o.VersionName
	}
	return json.Marshal(toSerialize)
}
