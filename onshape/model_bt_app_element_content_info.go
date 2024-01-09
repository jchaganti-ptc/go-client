/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28658-06d4d4923fc7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementContentInfo struct for BTAppElementContentInfo
type BTAppElementContentInfo struct {
	ChangeId *string                        `json:"changeId,omitempty"`
	Data     []BTAppElementContentEntryInfo `json:"data,omitempty"`
	// The numeric code identifying the error that occurred, if one occurred.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// A human-readable value for the error that occurred, if one occurred.
	ErrorDescription *string                `json:"errorDescription,omitempty"`
	ErrorValue       *BTAppElementErrorCode `json:"errorValue,omitempty"`
}

// NewBTAppElementContentInfo instantiates a new BTAppElementContentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementContentInfo() *BTAppElementContentInfo {
	this := BTAppElementContentInfo{}
	return &this
}

// NewBTAppElementContentInfoWithDefaults instantiates a new BTAppElementContentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementContentInfoWithDefaults() *BTAppElementContentInfo {
	this := BTAppElementContentInfo{}
	return &this
}

// GetChangeId returns the ChangeId field value if set, zero value otherwise.
func (o *BTAppElementContentInfo) GetChangeId() string {
	if o == nil || o.ChangeId == nil {
		var ret string
		return ret
	}
	return *o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentInfo) GetChangeIdOk() (*string, bool) {
	if o == nil || o.ChangeId == nil {
		return nil, false
	}
	return o.ChangeId, true
}

// HasChangeId returns a boolean if a field has been set.
func (o *BTAppElementContentInfo) HasChangeId() bool {
	if o != nil && o.ChangeId != nil {
		return true
	}

	return false
}

// SetChangeId gets a reference to the given string and assigns it to the ChangeId field.
func (o *BTAppElementContentInfo) SetChangeId(v string) {
	o.ChangeId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTAppElementContentInfo) GetData() []BTAppElementContentEntryInfo {
	if o == nil || o.Data == nil {
		var ret []BTAppElementContentEntryInfo
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentInfo) GetDataOk() ([]BTAppElementContentEntryInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTAppElementContentInfo) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BTAppElementContentEntryInfo and assigns it to the Data field.
func (o *BTAppElementContentInfo) SetData(v []BTAppElementContentEntryInfo) {
	o.Data = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTAppElementContentInfo) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentInfo) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTAppElementContentInfo) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTAppElementContentInfo) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *BTAppElementContentInfo) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentInfo) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *BTAppElementContentInfo) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *BTAppElementContentInfo) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorValue returns the ErrorValue field value if set, zero value otherwise.
func (o *BTAppElementContentInfo) GetErrorValue() BTAppElementErrorCode {
	if o == nil || o.ErrorValue == nil {
		var ret BTAppElementErrorCode
		return ret
	}
	return *o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementContentInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool) {
	if o == nil || o.ErrorValue == nil {
		return nil, false
	}
	return o.ErrorValue, true
}

// HasErrorValue returns a boolean if a field has been set.
func (o *BTAppElementContentInfo) HasErrorValue() bool {
	if o != nil && o.ErrorValue != nil {
		return true
	}

	return false
}

// SetErrorValue gets a reference to the given BTAppElementErrorCode and assigns it to the ErrorValue field.
func (o *BTAppElementContentInfo) SetErrorValue(v BTAppElementErrorCode) {
	o.ErrorValue = &v
}

func (o BTAppElementContentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeId != nil {
		toSerialize["changeId"] = o.ChangeId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorValue != nil {
		toSerialize["errorValue"] = o.ErrorValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementContentInfo struct {
	value *BTAppElementContentInfo
	isSet bool
}

func (v NullableBTAppElementContentInfo) Get() *BTAppElementContentInfo {
	return v.value
}

func (v *NullableBTAppElementContentInfo) Set(val *BTAppElementContentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementContentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementContentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementContentInfo(val *BTAppElementContentInfo) *NullableBTAppElementContentInfo {
	return &NullableBTAppElementContentInfo{value: val, isSet: true}
}

func (v NullableBTAppElementContentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementContentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
