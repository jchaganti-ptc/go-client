/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27881-5dbbe8053cdf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSetFeatureRollbackResponse1042 struct for BTSetFeatureRollbackResponse1042
type BTSetFeatureRollbackResponse1042 struct {
	BtType                 *string              `json:"btType,omitempty"`
	LibraryVersion         *int32               `json:"libraryVersion,omitempty"`
	MicroversionId         *BTMicroversionId366 `json:"microversionId,omitempty"`
	MicroversionSkew       *bool                `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew *bool                `json:"rejectMicroversionSkew,omitempty"`
	RollbackIndex          *int32               `json:"rollbackIndex,omitempty"`
	SerializationVersion   *string              `json:"serializationVersion,omitempty"`
	SourceMicroversion     *string              `json:"sourceMicroversion,omitempty"`
}

// NewBTSetFeatureRollbackResponse1042 instantiates a new BTSetFeatureRollbackResponse1042 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSetFeatureRollbackResponse1042() *BTSetFeatureRollbackResponse1042 {
	this := BTSetFeatureRollbackResponse1042{}
	return &this
}

// NewBTSetFeatureRollbackResponse1042WithDefaults instantiates a new BTSetFeatureRollbackResponse1042 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSetFeatureRollbackResponse1042WithDefaults() *BTSetFeatureRollbackResponse1042 {
	this := BTSetFeatureRollbackResponse1042{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackResponse1042) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackResponse1042) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackResponse1042) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSetFeatureRollbackResponse1042) SetBtType(v string) {
	o.BtType = &v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackResponse1042) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackResponse1042) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackResponse1042) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTSetFeatureRollbackResponse1042) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackResponse1042) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackResponse1042) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackResponse1042) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTSetFeatureRollbackResponse1042) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackResponse1042) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackResponse1042) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackResponse1042) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTSetFeatureRollbackResponse1042) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackResponse1042) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackResponse1042) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackResponse1042) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTSetFeatureRollbackResponse1042) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetRollbackIndex returns the RollbackIndex field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackResponse1042) GetRollbackIndex() int32 {
	if o == nil || o.RollbackIndex == nil {
		var ret int32
		return ret
	}
	return *o.RollbackIndex
}

// GetRollbackIndexOk returns a tuple with the RollbackIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackResponse1042) GetRollbackIndexOk() (*int32, bool) {
	if o == nil || o.RollbackIndex == nil {
		return nil, false
	}
	return o.RollbackIndex, true
}

// HasRollbackIndex returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackResponse1042) HasRollbackIndex() bool {
	if o != nil && o.RollbackIndex != nil {
		return true
	}

	return false
}

// SetRollbackIndex gets a reference to the given int32 and assigns it to the RollbackIndex field.
func (o *BTSetFeatureRollbackResponse1042) SetRollbackIndex(v int32) {
	o.RollbackIndex = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackResponse1042) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackResponse1042) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackResponse1042) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTSetFeatureRollbackResponse1042) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackResponse1042) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackResponse1042) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackResponse1042) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTSetFeatureRollbackResponse1042) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o BTSetFeatureRollbackResponse1042) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.RollbackIndex != nil {
		toSerialize["rollbackIndex"] = o.RollbackIndex
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTSetFeatureRollbackResponse1042 struct {
	value *BTSetFeatureRollbackResponse1042
	isSet bool
}

func (v NullableBTSetFeatureRollbackResponse1042) Get() *BTSetFeatureRollbackResponse1042 {
	return v.value
}

func (v *NullableBTSetFeatureRollbackResponse1042) Set(val *BTSetFeatureRollbackResponse1042) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSetFeatureRollbackResponse1042) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSetFeatureRollbackResponse1042) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSetFeatureRollbackResponse1042(val *BTSetFeatureRollbackResponse1042) *NullableBTSetFeatureRollbackResponse1042 {
	return &NullableBTSetFeatureRollbackResponse1042{value: val, isSet: true}
}

func (v NullableBTSetFeatureRollbackResponse1042) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSetFeatureRollbackResponse1042) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
