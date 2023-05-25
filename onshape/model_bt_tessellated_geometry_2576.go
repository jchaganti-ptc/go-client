/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16419-6916b772b99f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTTessellatedGeometry2576 - struct for BTTessellatedGeometry2576
type BTTessellatedGeometry2576 struct {
	implBTTessellatedGeometry2576 interface{}
}

// BTEntityPoint29AsBTTessellatedGeometry2576 is a convenience function that returns BTEntityPoint29 wrapped in BTTessellatedGeometry2576
func (o *BTEntityPoint29) AsBTTessellatedGeometry2576() *BTTessellatedGeometry2576 {
	return &BTTessellatedGeometry2576{o}
}

// BTEntityDegenerateEdge1129AsBTTessellatedGeometry2576 is a convenience function that returns BTEntityDegenerateEdge1129 wrapped in BTTessellatedGeometry2576
func (o *BTEntityDegenerateEdge1129) AsBTTessellatedGeometry2576() *BTTessellatedGeometry2576 {
	return &BTTessellatedGeometry2576{o}
}

// BTEntityEdge30AsBTTessellatedGeometry2576 is a convenience function that returns BTEntityEdge30 wrapped in BTTessellatedGeometry2576
func (o *BTEntityEdge30) AsBTTessellatedGeometry2576() *BTTessellatedGeometry2576 {
	return &BTTessellatedGeometry2576{o}
}

// BTSimulationFace2147AsBTTessellatedGeometry2576 is a convenience function that returns BTSimulationFace2147 wrapped in BTTessellatedGeometry2576
func (o *BTSimulationFace2147) AsBTTessellatedGeometry2576() *BTTessellatedGeometry2576 {
	return &BTTessellatedGeometry2576{o}
}

// BTEntityFace31AsBTTessellatedGeometry2576 is a convenience function that returns BTEntityFace31 wrapped in BTTessellatedGeometry2576
func (o *BTEntityFace31) AsBTTessellatedGeometry2576() *BTTessellatedGeometry2576 {
	return &BTTessellatedGeometry2576{o}
}

// NewBTTessellatedGeometry2576 instantiates a new BTTessellatedGeometry2576 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTessellatedGeometry2576() *BTTessellatedGeometry2576 {
	this := BTTessellatedGeometry2576{Newbase_BTTessellatedGeometry2576()}
	return &this
}

// NewBTTessellatedGeometry2576WithDefaults instantiates a new BTTessellatedGeometry2576 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTessellatedGeometry2576WithDefaults() *BTTessellatedGeometry2576 {
	this := BTTessellatedGeometry2576{Newbase_BTTessellatedGeometry2576WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTessellatedGeometry2576) GetBtType() string {
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
func (o *BTTessellatedGeometry2576) GetBtTypeOk() (*string, bool) {
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
func (o *BTTessellatedGeometry2576) HasBtType() bool {
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
func (o *BTTessellatedGeometry2576) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *BTTessellatedGeometry2576) GetCompressed() bool {
	type getResult interface {
		GetCompressed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCompressed()
	} else {
		var de bool
		return de
	}
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellatedGeometry2576) GetCompressedOk() (*bool, bool) {
	type getResult interface {
		GetCompressedOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetCompressedOk()
	} else {
		return nil, false
	}
}

// HasCompressed returns a boolean if a field has been set.
func (o *BTTessellatedGeometry2576) HasCompressed() bool {
	type getResult interface {
		HasCompressed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasCompressed()
	} else {
		return false
	}
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *BTTessellatedGeometry2576) SetCompressed(v bool) {
	type getResult interface {
		SetCompressed(v bool)
	}

	o.GetActualInstance().(getResult).SetCompressed(v)
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTTessellatedGeometry2576) GetDecompressed() BTEntityGeometry35 {
	type getResult interface {
		GetDecompressed() BTEntityGeometry35
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDecompressed()
	} else {
		var de BTEntityGeometry35
		return de
	}
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellatedGeometry2576) GetDecompressedOk() (*BTEntityGeometry35, bool) {
	type getResult interface {
		GetDecompressedOk() (*BTEntityGeometry35, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDecompressedOk()
	} else {
		return nil, false
	}
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTTessellatedGeometry2576) HasDecompressed() bool {
	type getResult interface {
		HasDecompressed() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDecompressed()
	} else {
		return false
	}
}

// SetDecompressed gets a reference to the given BTEntityGeometry35 and assigns it to the Decompressed field.
func (o *BTTessellatedGeometry2576) SetDecompressed(v BTEntityGeometry35) {
	type getResult interface {
		SetDecompressed(v BTEntityGeometry35)
	}

	o.GetActualInstance().(getResult).SetDecompressed(v)
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTTessellatedGeometry2576) GetErrorCode() int32 {
	type getResult interface {
		GetErrorCode() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorCode()
	} else {
		var de int32
		return de
	}
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellatedGeometry2576) GetErrorCodeOk() (*int32, bool) {
	type getResult interface {
		GetErrorCodeOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetErrorCodeOk()
	} else {
		return nil, false
	}
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTTessellatedGeometry2576) HasErrorCode() bool {
	type getResult interface {
		HasErrorCode() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasErrorCode()
	} else {
		return false
	}
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTTessellatedGeometry2576) SetErrorCode(v int32) {
	type getResult interface {
		SetErrorCode(v int32)
	}

	o.GetActualInstance().(getResult).SetErrorCode(v)
}

// GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field value if set, zero value otherwise.
func (o *BTTessellatedGeometry2576) GetEstimatedMemoryUsageInBytes() int32 {
	type getResult interface {
		GetEstimatedMemoryUsageInBytes() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEstimatedMemoryUsageInBytes()
	} else {
		var de int32
		return de
	}
}

// GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellatedGeometry2576) GetEstimatedMemoryUsageInBytesOk() (*int32, bool) {
	type getResult interface {
		GetEstimatedMemoryUsageInBytesOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEstimatedMemoryUsageInBytesOk()
	} else {
		return nil, false
	}
}

// HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.
func (o *BTTessellatedGeometry2576) HasEstimatedMemoryUsageInBytes() bool {
	type getResult interface {
		HasEstimatedMemoryUsageInBytes() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEstimatedMemoryUsageInBytes()
	} else {
		return false
	}
}

// SetEstimatedMemoryUsageInBytes gets a reference to the given int32 and assigns it to the EstimatedMemoryUsageInBytes field.
func (o *BTTessellatedGeometry2576) SetEstimatedMemoryUsageInBytes(v int32) {
	type getResult interface {
		SetEstimatedMemoryUsageInBytes(v int32)
	}

	o.GetActualInstance().(getResult).SetEstimatedMemoryUsageInBytes(v)
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *BTTessellatedGeometry2576) GetFace() bool {
	type getResult interface {
		GetFace() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFace()
	} else {
		var de bool
		return de
	}
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellatedGeometry2576) GetFaceOk() (*bool, bool) {
	type getResult interface {
		GetFaceOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFaceOk()
	} else {
		return nil, false
	}
}

// HasFace returns a boolean if a field has been set.
func (o *BTTessellatedGeometry2576) HasFace() bool {
	type getResult interface {
		HasFace() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFace()
	} else {
		return false
	}
}

// SetFace gets a reference to the given bool and assigns it to the Face field.
func (o *BTTessellatedGeometry2576) SetFace(v bool) {
	type getResult interface {
		SetFace(v bool)
	}

	o.GetActualInstance().(getResult).SetFace(v)
}

// GetHasTessellationError returns the HasTessellationError field value if set, zero value otherwise.
func (o *BTTessellatedGeometry2576) GetHasTessellationError() bool {
	type getResult interface {
		GetHasTessellationError() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHasTessellationError()
	} else {
		var de bool
		return de
	}
}

// GetHasTessellationErrorOk returns a tuple with the HasTessellationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellatedGeometry2576) GetHasTessellationErrorOk() (*bool, bool) {
	type getResult interface {
		GetHasTessellationErrorOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetHasTessellationErrorOk()
	} else {
		return nil, false
	}
}

// HasHasTessellationError returns a boolean if a field has been set.
func (o *BTTessellatedGeometry2576) HasHasTessellationError() bool {
	type getResult interface {
		HasHasTessellationError() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasHasTessellationError()
	} else {
		return false
	}
}

// SetHasTessellationError gets a reference to the given bool and assigns it to the HasTessellationError field.
func (o *BTTessellatedGeometry2576) SetHasTessellationError(v bool) {
	type getResult interface {
		SetHasTessellationError(v bool)
	}

	o.GetActualInstance().(getResult).SetHasTessellationError(v)
}

// GetSettingIndex returns the SettingIndex field value if set, zero value otherwise.
func (o *BTTessellatedGeometry2576) GetSettingIndex() int32 {
	type getResult interface {
		GetSettingIndex() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSettingIndex()
	} else {
		var de int32
		return de
	}
}

// GetSettingIndexOk returns a tuple with the SettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTessellatedGeometry2576) GetSettingIndexOk() (*int32, bool) {
	type getResult interface {
		GetSettingIndexOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSettingIndexOk()
	} else {
		return nil, false
	}
}

// HasSettingIndex returns a boolean if a field has been set.
func (o *BTTessellatedGeometry2576) HasSettingIndex() bool {
	type getResult interface {
		HasSettingIndex() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSettingIndex()
	} else {
		return false
	}
}

// SetSettingIndex gets a reference to the given int32 and assigns it to the SettingIndex field.
func (o *BTTessellatedGeometry2576) SetSettingIndex(v int32) {
	type getResult interface {
		SetSettingIndex(v int32)
	}

	o.GetActualInstance().(getResult).SetSettingIndex(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTTessellatedGeometry2576) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTEntityDegenerateEdge-1129'
	if jsonDict["btType"] == "BTEntityDegenerateEdge-1129" {
		// try to unmarshal JSON data into BTEntityDegenerateEdge1129
		var qr *BTEntityDegenerateEdge1129
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTessellatedGeometry2576 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTessellatedGeometry2576 = nil
			return fmt.Errorf("Failed to unmarshal BTTessellatedGeometry2576 as BTEntityDegenerateEdge1129: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEntityEdge-30'
	if jsonDict["btType"] == "BTEntityEdge-30" {
		// try to unmarshal JSON data into BTEntityEdge30
		var qr *BTEntityEdge30
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTessellatedGeometry2576 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTessellatedGeometry2576 = nil
			return fmt.Errorf("Failed to unmarshal BTTessellatedGeometry2576 as BTEntityEdge30: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEntityFace-31'
	if jsonDict["btType"] == "BTEntityFace-31" {
		// try to unmarshal JSON data into BTEntityFace31
		var qr *BTEntityFace31
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTessellatedGeometry2576 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTessellatedGeometry2576 = nil
			return fmt.Errorf("Failed to unmarshal BTTessellatedGeometry2576 as BTEntityFace31: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEntityPoint-29'
	if jsonDict["btType"] == "BTEntityPoint-29" {
		// try to unmarshal JSON data into BTEntityPoint29
		var qr *BTEntityPoint29
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTessellatedGeometry2576 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTessellatedGeometry2576 = nil
			return fmt.Errorf("Failed to unmarshal BTTessellatedGeometry2576 as BTEntityPoint29: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSimulationFace-2147'
	if jsonDict["btType"] == "BTSimulationFace-2147" {
		// try to unmarshal JSON data into BTSimulationFace2147
		var qr *BTSimulationFace2147
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTessellatedGeometry2576 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTessellatedGeometry2576 = nil
			return fmt.Errorf("Failed to unmarshal BTTessellatedGeometry2576 as BTSimulationFace2147: %s", err.Error())
		}
	}

	var qtx *base_BTTessellatedGeometry2576
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTTessellatedGeometry2576 = qtx
		return nil // data stored in dst.base_BTTessellatedGeometry2576, return on the first match
	} else {
		dst.implBTTessellatedGeometry2576 = nil
		return fmt.Errorf("Failed to unmarshal BTTessellatedGeometry2576 as base_BTTessellatedGeometry2576: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTTessellatedGeometry2576) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTTessellatedGeometry2576) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTTessellatedGeometry2576
}

type NullableBTTessellatedGeometry2576 struct {
	value *BTTessellatedGeometry2576
	isSet bool
}

func (v NullableBTTessellatedGeometry2576) Get() *BTTessellatedGeometry2576 {
	return v.value
}

func (v *NullableBTTessellatedGeometry2576) Set(val *BTTessellatedGeometry2576) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTessellatedGeometry2576) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTessellatedGeometry2576) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTessellatedGeometry2576(val *BTTessellatedGeometry2576) *NullableBTTessellatedGeometry2576 {
	return &NullableBTTessellatedGeometry2576{value: val, isSet: true}
}

func (v NullableBTTessellatedGeometry2576) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTessellatedGeometry2576) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTTessellatedGeometry2576 struct {
	BtType                      *string             `json:"btType,omitempty"`
	Compressed                  *bool               `json:"compressed,omitempty"`
	Decompressed                *BTEntityGeometry35 `json:"decompressed,omitempty"`
	ErrorCode                   *int32              `json:"errorCode,omitempty"`
	EstimatedMemoryUsageInBytes *int32              `json:"estimatedMemoryUsageInBytes,omitempty"`
	Face                        *bool               `json:"face,omitempty"`
	HasTessellationError        *bool               `json:"hasTessellationError,omitempty"`
	SettingIndex                *int32              `json:"settingIndex,omitempty"`
}

// Newbase_BTTessellatedGeometry2576 instantiates a new base_BTTessellatedGeometry2576 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTTessellatedGeometry2576() *base_BTTessellatedGeometry2576 {
	this := base_BTTessellatedGeometry2576{}
	return &this
}

// Newbase_BTTessellatedGeometry2576WithDefaults instantiates a new base_BTTessellatedGeometry2576 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTTessellatedGeometry2576WithDefaults() *base_BTTessellatedGeometry2576 {
	this := base_BTTessellatedGeometry2576{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTTessellatedGeometry2576) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTessellatedGeometry2576) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTTessellatedGeometry2576) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTTessellatedGeometry2576) SetBtType(v string) {
	o.BtType = &v
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *base_BTTessellatedGeometry2576) GetCompressed() bool {
	if o == nil || o.Compressed == nil {
		var ret bool
		return ret
	}
	return *o.Compressed
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTessellatedGeometry2576) GetCompressedOk() (*bool, bool) {
	if o == nil || o.Compressed == nil {
		return nil, false
	}
	return o.Compressed, true
}

// HasCompressed returns a boolean if a field has been set.
func (o *base_BTTessellatedGeometry2576) HasCompressed() bool {
	if o != nil && o.Compressed != nil {
		return true
	}

	return false
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *base_BTTessellatedGeometry2576) SetCompressed(v bool) {
	o.Compressed = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *base_BTTessellatedGeometry2576) GetDecompressed() BTEntityGeometry35 {
	if o == nil || o.Decompressed == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTessellatedGeometry2576) GetDecompressedOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *base_BTTessellatedGeometry2576) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTEntityGeometry35 and assigns it to the Decompressed field.
func (o *base_BTTessellatedGeometry2576) SetDecompressed(v BTEntityGeometry35) {
	o.Decompressed = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *base_BTTessellatedGeometry2576) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTessellatedGeometry2576) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *base_BTTessellatedGeometry2576) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *base_BTTessellatedGeometry2576) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field value if set, zero value otherwise.
func (o *base_BTTessellatedGeometry2576) GetEstimatedMemoryUsageInBytes() int32 {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedMemoryUsageInBytes
}

// GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTessellatedGeometry2576) GetEstimatedMemoryUsageInBytesOk() (*int32, bool) {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		return nil, false
	}
	return o.EstimatedMemoryUsageInBytes, true
}

// HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.
func (o *base_BTTessellatedGeometry2576) HasEstimatedMemoryUsageInBytes() bool {
	if o != nil && o.EstimatedMemoryUsageInBytes != nil {
		return true
	}

	return false
}

// SetEstimatedMemoryUsageInBytes gets a reference to the given int32 and assigns it to the EstimatedMemoryUsageInBytes field.
func (o *base_BTTessellatedGeometry2576) SetEstimatedMemoryUsageInBytes(v int32) {
	o.EstimatedMemoryUsageInBytes = &v
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *base_BTTessellatedGeometry2576) GetFace() bool {
	if o == nil || o.Face == nil {
		var ret bool
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTessellatedGeometry2576) GetFaceOk() (*bool, bool) {
	if o == nil || o.Face == nil {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *base_BTTessellatedGeometry2576) HasFace() bool {
	if o != nil && o.Face != nil {
		return true
	}

	return false
}

// SetFace gets a reference to the given bool and assigns it to the Face field.
func (o *base_BTTessellatedGeometry2576) SetFace(v bool) {
	o.Face = &v
}

// GetHasTessellationError returns the HasTessellationError field value if set, zero value otherwise.
func (o *base_BTTessellatedGeometry2576) GetHasTessellationError() bool {
	if o == nil || o.HasTessellationError == nil {
		var ret bool
		return ret
	}
	return *o.HasTessellationError
}

// GetHasTessellationErrorOk returns a tuple with the HasTessellationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTessellatedGeometry2576) GetHasTessellationErrorOk() (*bool, bool) {
	if o == nil || o.HasTessellationError == nil {
		return nil, false
	}
	return o.HasTessellationError, true
}

// HasHasTessellationError returns a boolean if a field has been set.
func (o *base_BTTessellatedGeometry2576) HasHasTessellationError() bool {
	if o != nil && o.HasTessellationError != nil {
		return true
	}

	return false
}

// SetHasTessellationError gets a reference to the given bool and assigns it to the HasTessellationError field.
func (o *base_BTTessellatedGeometry2576) SetHasTessellationError(v bool) {
	o.HasTessellationError = &v
}

// GetSettingIndex returns the SettingIndex field value if set, zero value otherwise.
func (o *base_BTTessellatedGeometry2576) GetSettingIndex() int32 {
	if o == nil || o.SettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.SettingIndex
}

// GetSettingIndexOk returns a tuple with the SettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTessellatedGeometry2576) GetSettingIndexOk() (*int32, bool) {
	if o == nil || o.SettingIndex == nil {
		return nil, false
	}
	return o.SettingIndex, true
}

// HasSettingIndex returns a boolean if a field has been set.
func (o *base_BTTessellatedGeometry2576) HasSettingIndex() bool {
	if o != nil && o.SettingIndex != nil {
		return true
	}

	return false
}

// SetSettingIndex gets a reference to the given int32 and assigns it to the SettingIndex field.
func (o *base_BTTessellatedGeometry2576) SetSettingIndex(v int32) {
	o.SettingIndex = &v
}

func (o base_BTTessellatedGeometry2576) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Compressed != nil {
		toSerialize["compressed"] = o.Compressed
	}
	if o.Decompressed != nil {
		toSerialize["decompressed"] = o.Decompressed
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.EstimatedMemoryUsageInBytes != nil {
		toSerialize["estimatedMemoryUsageInBytes"] = o.EstimatedMemoryUsageInBytes
	}
	if o.Face != nil {
		toSerialize["face"] = o.Face
	}
	if o.HasTessellationError != nil {
		toSerialize["hasTessellationError"] = o.HasTessellationError
	}
	if o.SettingIndex != nil {
		toSerialize["settingIndex"] = o.SettingIndex
	}
	return json.Marshal(toSerialize)
}
