/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6621-03f431879e4b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTEntityPoint29 struct for BTEntityPoint29
type BTEntityPoint29 struct {
	BtType                      *string             `json:"btType,omitempty"`
	Compressed                  *bool               `json:"compressed,omitempty"`
	Decompressed                *BTEntityGeometry35 `json:"decompressed,omitempty"`
	ErrorCode                   *int32              `json:"errorCode,omitempty"`
	EstimatedMemoryUsageInBytes *int32              `json:"estimatedMemoryUsageInBytes,omitempty"`
	Face                        *bool               `json:"face,omitempty"`
	HasTessellationError        *bool               `json:"hasTessellationError,omitempty"`
	SettingIndex                *int32              `json:"settingIndex,omitempty"`
	Point                       []float32           `json:"point,omitempty"`
}

// NewBTEntityPoint29 instantiates a new BTEntityPoint29 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEntityPoint29() *BTEntityPoint29 {
	this := BTEntityPoint29{}
	return &this
}

// NewBTEntityPoint29WithDefaults instantiates a new BTEntityPoint29 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEntityPoint29WithDefaults() *BTEntityPoint29 {
	this := BTEntityPoint29{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEntityPoint29) SetBtType(v string) {
	o.BtType = &v
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetCompressed() bool {
	if o == nil || o.Compressed == nil {
		var ret bool
		return ret
	}
	return *o.Compressed
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetCompressedOk() (*bool, bool) {
	if o == nil || o.Compressed == nil {
		return nil, false
	}
	return o.Compressed, true
}

// HasCompressed returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasCompressed() bool {
	if o != nil && o.Compressed != nil {
		return true
	}

	return false
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *BTEntityPoint29) SetCompressed(v bool) {
	o.Compressed = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetDecompressed() BTEntityGeometry35 {
	if o == nil || o.Decompressed == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetDecompressedOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTEntityGeometry35 and assigns it to the Decompressed field.
func (o *BTEntityPoint29) SetDecompressed(v BTEntityGeometry35) {
	o.Decompressed = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTEntityPoint29) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetEstimatedMemoryUsageInBytes() int32 {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedMemoryUsageInBytes
}

// GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetEstimatedMemoryUsageInBytesOk() (*int32, bool) {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		return nil, false
	}
	return o.EstimatedMemoryUsageInBytes, true
}

// HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasEstimatedMemoryUsageInBytes() bool {
	if o != nil && o.EstimatedMemoryUsageInBytes != nil {
		return true
	}

	return false
}

// SetEstimatedMemoryUsageInBytes gets a reference to the given int32 and assigns it to the EstimatedMemoryUsageInBytes field.
func (o *BTEntityPoint29) SetEstimatedMemoryUsageInBytes(v int32) {
	o.EstimatedMemoryUsageInBytes = &v
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetFace() bool {
	if o == nil || o.Face == nil {
		var ret bool
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetFaceOk() (*bool, bool) {
	if o == nil || o.Face == nil {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasFace() bool {
	if o != nil && o.Face != nil {
		return true
	}

	return false
}

// SetFace gets a reference to the given bool and assigns it to the Face field.
func (o *BTEntityPoint29) SetFace(v bool) {
	o.Face = &v
}

// GetHasTessellationError returns the HasTessellationError field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetHasTessellationError() bool {
	if o == nil || o.HasTessellationError == nil {
		var ret bool
		return ret
	}
	return *o.HasTessellationError
}

// GetHasTessellationErrorOk returns a tuple with the HasTessellationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetHasTessellationErrorOk() (*bool, bool) {
	if o == nil || o.HasTessellationError == nil {
		return nil, false
	}
	return o.HasTessellationError, true
}

// HasHasTessellationError returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasHasTessellationError() bool {
	if o != nil && o.HasTessellationError != nil {
		return true
	}

	return false
}

// SetHasTessellationError gets a reference to the given bool and assigns it to the HasTessellationError field.
func (o *BTEntityPoint29) SetHasTessellationError(v bool) {
	o.HasTessellationError = &v
}

// GetSettingIndex returns the SettingIndex field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetSettingIndex() int32 {
	if o == nil || o.SettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.SettingIndex
}

// GetSettingIndexOk returns a tuple with the SettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetSettingIndexOk() (*int32, bool) {
	if o == nil || o.SettingIndex == nil {
		return nil, false
	}
	return o.SettingIndex, true
}

// HasSettingIndex returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasSettingIndex() bool {
	if o != nil && o.SettingIndex != nil {
		return true
	}

	return false
}

// SetSettingIndex gets a reference to the given int32 and assigns it to the SettingIndex field.
func (o *BTEntityPoint29) SetSettingIndex(v int32) {
	o.SettingIndex = &v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *BTEntityPoint29) GetPoint() []float32 {
	if o == nil || o.Point == nil {
		var ret []float32
		return ret
	}
	return o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityPoint29) GetPointOk() ([]float32, bool) {
	if o == nil || o.Point == nil {
		return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *BTEntityPoint29) HasPoint() bool {
	if o != nil && o.Point != nil {
		return true
	}

	return false
}

// SetPoint gets a reference to the given []float32 and assigns it to the Point field.
func (o *BTEntityPoint29) SetPoint(v []float32) {
	o.Point = v
}

func (o BTEntityPoint29) MarshalJSON() ([]byte, error) {
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
	if o.Point != nil {
		toSerialize["point"] = o.Point
	}
	return json.Marshal(toSerialize)
}

type NullableBTEntityPoint29 struct {
	value *BTEntityPoint29
	isSet bool
}

func (v NullableBTEntityPoint29) Get() *BTEntityPoint29 {
	return v.value
}

func (v *NullableBTEntityPoint29) Set(val *BTEntityPoint29) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEntityPoint29) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEntityPoint29) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEntityPoint29(val *BTEntityPoint29) *NullableBTEntityPoint29 {
	return &NullableBTEntityPoint29{value: val, isSet: true}
}

func (v NullableBTEntityPoint29) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEntityPoint29) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
