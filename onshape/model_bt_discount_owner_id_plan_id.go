/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27960-e195de6ac56c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDiscountOwnerIdPlanId struct for BTDiscountOwnerIdPlanId
type BTDiscountOwnerIdPlanId struct {
	OwnerId *string `json:"ownerId,omitempty"`
	PlanId  *string `json:"planId,omitempty"`
}

// NewBTDiscountOwnerIdPlanId instantiates a new BTDiscountOwnerIdPlanId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDiscountOwnerIdPlanId() *BTDiscountOwnerIdPlanId {
	this := BTDiscountOwnerIdPlanId{}
	return &this
}

// NewBTDiscountOwnerIdPlanIdWithDefaults instantiates a new BTDiscountOwnerIdPlanId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDiscountOwnerIdPlanIdWithDefaults() *BTDiscountOwnerIdPlanId {
	this := BTDiscountOwnerIdPlanId{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTDiscountOwnerIdPlanId) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountOwnerIdPlanId) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTDiscountOwnerIdPlanId) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTDiscountOwnerIdPlanId) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *BTDiscountOwnerIdPlanId) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiscountOwnerIdPlanId) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *BTDiscountOwnerIdPlanId) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *BTDiscountOwnerIdPlanId) SetPlanId(v string) {
	o.PlanId = &v
}

func (o BTDiscountOwnerIdPlanId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.PlanId != nil {
		toSerialize["planId"] = o.PlanId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDiscountOwnerIdPlanId struct {
	value *BTDiscountOwnerIdPlanId
	isSet bool
}

func (v NullableBTDiscountOwnerIdPlanId) Get() *BTDiscountOwnerIdPlanId {
	return v.value
}

func (v *NullableBTDiscountOwnerIdPlanId) Set(val *BTDiscountOwnerIdPlanId) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDiscountOwnerIdPlanId) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDiscountOwnerIdPlanId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDiscountOwnerIdPlanId(val *BTDiscountOwnerIdPlanId) *NullableBTDiscountOwnerIdPlanId {
	return &NullableBTDiscountOwnerIdPlanId{value: val, isSet: true}
}

func (v NullableBTDiscountOwnerIdPlanId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDiscountOwnerIdPlanId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
