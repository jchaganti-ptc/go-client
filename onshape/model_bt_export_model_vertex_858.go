/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28126-700645382199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportModelVertex858 struct for BTExportModelVertex858
type BTExportModelVertex858 struct {
	BtType *string        `json:"btType,omitempty"`
	Id     *string        `json:"id,omitempty"`
	Point  *BTVector3d389 `json:"point,omitempty"`
}

// NewBTExportModelVertex858 instantiates a new BTExportModelVertex858 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelVertex858() *BTExportModelVertex858 {
	this := BTExportModelVertex858{}
	return &this
}

// NewBTExportModelVertex858WithDefaults instantiates a new BTExportModelVertex858 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelVertex858WithDefaults() *BTExportModelVertex858 {
	this := BTExportModelVertex858{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportModelVertex858) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelVertex858) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportModelVertex858) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportModelVertex858) SetBtType(v string) {
	o.BtType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportModelVertex858) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelVertex858) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportModelVertex858) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportModelVertex858) SetId(v string) {
	o.Id = &v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *BTExportModelVertex858) GetPoint() BTVector3d389 {
	if o == nil || o.Point == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelVertex858) GetPointOk() (*BTVector3d389, bool) {
	if o == nil || o.Point == nil {
		return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *BTExportModelVertex858) HasPoint() bool {
	if o != nil && o.Point != nil {
		return true
	}

	return false
}

// SetPoint gets a reference to the given BTVector3d389 and assigns it to the Point field.
func (o *BTExportModelVertex858) SetPoint(v BTVector3d389) {
	o.Point = &v
}

func (o BTExportModelVertex858) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Point != nil {
		toSerialize["point"] = o.Point
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelVertex858 struct {
	value *BTExportModelVertex858
	isSet bool
}

func (v NullableBTExportModelVertex858) Get() *BTExportModelVertex858 {
	return v.value
}

func (v *NullableBTExportModelVertex858) Set(val *BTExportModelVertex858) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelVertex858) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelVertex858) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelVertex858(val *BTExportModelVertex858) *NullableBTExportModelVertex858 {
	return &NullableBTExportModelVertex858{value: val, isSet: true}
}

func (v NullableBTExportModelVertex858) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelVertex858) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
