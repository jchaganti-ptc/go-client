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

// BTGetJsonResponse2137 struct for BTGetJsonResponse2137
type BTGetJsonResponse2137 struct {
	BtType   *string                    `json:"btType,omitempty"`
	ChangeId *string                    `json:"changeId,omitempty"`
	Tree     *BTGetJsonResponse2137Tree `json:"tree,omitempty"`
}

// NewBTGetJsonResponse2137 instantiates a new BTGetJsonResponse2137 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGetJsonResponse2137() *BTGetJsonResponse2137 {
	this := BTGetJsonResponse2137{}
	return &this
}

// NewBTGetJsonResponse2137WithDefaults instantiates a new BTGetJsonResponse2137 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGetJsonResponse2137WithDefaults() *BTGetJsonResponse2137 {
	this := BTGetJsonResponse2137{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTGetJsonResponse2137) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGetJsonResponse2137) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTGetJsonResponse2137) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTGetJsonResponse2137) SetBtType(v string) {
	o.BtType = &v
}

// GetChangeId returns the ChangeId field value if set, zero value otherwise.
func (o *BTGetJsonResponse2137) GetChangeId() string {
	if o == nil || o.ChangeId == nil {
		var ret string
		return ret
	}
	return *o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGetJsonResponse2137) GetChangeIdOk() (*string, bool) {
	if o == nil || o.ChangeId == nil {
		return nil, false
	}
	return o.ChangeId, true
}

// HasChangeId returns a boolean if a field has been set.
func (o *BTGetJsonResponse2137) HasChangeId() bool {
	if o != nil && o.ChangeId != nil {
		return true
	}

	return false
}

// SetChangeId gets a reference to the given string and assigns it to the ChangeId field.
func (o *BTGetJsonResponse2137) SetChangeId(v string) {
	o.ChangeId = &v
}

// GetTree returns the Tree field value if set, zero value otherwise.
func (o *BTGetJsonResponse2137) GetTree() BTGetJsonResponse2137Tree {
	if o == nil || o.Tree == nil {
		var ret BTGetJsonResponse2137Tree
		return ret
	}
	return *o.Tree
}

// GetTreeOk returns a tuple with the Tree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGetJsonResponse2137) GetTreeOk() (*BTGetJsonResponse2137Tree, bool) {
	if o == nil || o.Tree == nil {
		return nil, false
	}
	return o.Tree, true
}

// HasTree returns a boolean if a field has been set.
func (o *BTGetJsonResponse2137) HasTree() bool {
	if o != nil && o.Tree != nil {
		return true
	}

	return false
}

// SetTree gets a reference to the given BTGetJsonResponse2137Tree and assigns it to the Tree field.
func (o *BTGetJsonResponse2137) SetTree(v BTGetJsonResponse2137Tree) {
	o.Tree = &v
}

func (o BTGetJsonResponse2137) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ChangeId != nil {
		toSerialize["changeId"] = o.ChangeId
	}
	if o.Tree != nil {
		toSerialize["tree"] = o.Tree
	}
	return json.Marshal(toSerialize)
}

type NullableBTGetJsonResponse2137 struct {
	value *BTGetJsonResponse2137
	isSet bool
}

func (v NullableBTGetJsonResponse2137) Get() *BTGetJsonResponse2137 {
	return v.value
}

func (v *NullableBTGetJsonResponse2137) Set(val *BTGetJsonResponse2137) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGetJsonResponse2137) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGetJsonResponse2137) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGetJsonResponse2137(val *BTGetJsonResponse2137) *NullableBTGetJsonResponse2137 {
	return &NullableBTGetJsonResponse2137{value: val, isSet: true}
}

func (v NullableBTGetJsonResponse2137) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGetJsonResponse2137) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
