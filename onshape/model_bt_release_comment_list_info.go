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

// BTReleaseCommentListInfo struct for BTReleaseCommentListInfo
type BTReleaseCommentListInfo struct {
	Comments []BTCommentInfo `json:"comments,omitempty"`
	RpId     *string         `json:"rpId,omitempty"`
	RpName   *string         `json:"rpName,omitempty"`
}

// NewBTReleaseCommentListInfo instantiates a new BTReleaseCommentListInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTReleaseCommentListInfo() *BTReleaseCommentListInfo {
	this := BTReleaseCommentListInfo{}
	return &this
}

// NewBTReleaseCommentListInfoWithDefaults instantiates a new BTReleaseCommentListInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTReleaseCommentListInfoWithDefaults() *BTReleaseCommentListInfo {
	this := BTReleaseCommentListInfo{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *BTReleaseCommentListInfo) GetComments() []BTCommentInfo {
	if o == nil || o.Comments == nil {
		var ret []BTCommentInfo
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseCommentListInfo) GetCommentsOk() ([]BTCommentInfo, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *BTReleaseCommentListInfo) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []BTCommentInfo and assigns it to the Comments field.
func (o *BTReleaseCommentListInfo) SetComments(v []BTCommentInfo) {
	o.Comments = v
}

// GetRpId returns the RpId field value if set, zero value otherwise.
func (o *BTReleaseCommentListInfo) GetRpId() string {
	if o == nil || o.RpId == nil {
		var ret string
		return ret
	}
	return *o.RpId
}

// GetRpIdOk returns a tuple with the RpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseCommentListInfo) GetRpIdOk() (*string, bool) {
	if o == nil || o.RpId == nil {
		return nil, false
	}
	return o.RpId, true
}

// HasRpId returns a boolean if a field has been set.
func (o *BTReleaseCommentListInfo) HasRpId() bool {
	if o != nil && o.RpId != nil {
		return true
	}

	return false
}

// SetRpId gets a reference to the given string and assigns it to the RpId field.
func (o *BTReleaseCommentListInfo) SetRpId(v string) {
	o.RpId = &v
}

// GetRpName returns the RpName field value if set, zero value otherwise.
func (o *BTReleaseCommentListInfo) GetRpName() string {
	if o == nil || o.RpName == nil {
		var ret string
		return ret
	}
	return *o.RpName
}

// GetRpNameOk returns a tuple with the RpName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTReleaseCommentListInfo) GetRpNameOk() (*string, bool) {
	if o == nil || o.RpName == nil {
		return nil, false
	}
	return o.RpName, true
}

// HasRpName returns a boolean if a field has been set.
func (o *BTReleaseCommentListInfo) HasRpName() bool {
	if o != nil && o.RpName != nil {
		return true
	}

	return false
}

// SetRpName gets a reference to the given string and assigns it to the RpName field.
func (o *BTReleaseCommentListInfo) SetRpName(v string) {
	o.RpName = &v
}

func (o BTReleaseCommentListInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.RpId != nil {
		toSerialize["rpId"] = o.RpId
	}
	if o.RpName != nil {
		toSerialize["rpName"] = o.RpName
	}
	return json.Marshal(toSerialize)
}

type NullableBTReleaseCommentListInfo struct {
	value *BTReleaseCommentListInfo
	isSet bool
}

func (v NullableBTReleaseCommentListInfo) Get() *BTReleaseCommentListInfo {
	return v.value
}

func (v *NullableBTReleaseCommentListInfo) Set(val *BTReleaseCommentListInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTReleaseCommentListInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTReleaseCommentListInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTReleaseCommentListInfo(val *BTReleaseCommentListInfo) *NullableBTReleaseCommentListInfo {
	return &NullableBTReleaseCommentListInfo{value: val, isSet: true}
}

func (v NullableBTReleaseCommentListInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTReleaseCommentListInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
