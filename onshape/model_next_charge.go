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

// NextCharge struct for NextCharge
type NextCharge struct {
	Amount           *int64    `json:"amount,omitempty"`
	CurrentPeriodEnd *JSONTime `json:"currentPeriodEnd,omitempty"`
	Interval         *string   `json:"interval,omitempty"`
	Total            *int64    `json:"total,omitempty"`
}

// NewNextCharge instantiates a new NextCharge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextCharge() *NextCharge {
	this := NextCharge{}
	return &this
}

// NewNextChargeWithDefaults instantiates a new NextCharge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextChargeWithDefaults() *NextCharge {
	this := NextCharge{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *NextCharge) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextCharge) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *NextCharge) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *NextCharge) SetAmount(v int64) {
	o.Amount = &v
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value if set, zero value otherwise.
func (o *NextCharge) GetCurrentPeriodEnd() JSONTime {
	if o == nil || o.CurrentPeriodEnd == nil {
		var ret JSONTime
		return ret
	}
	return *o.CurrentPeriodEnd
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextCharge) GetCurrentPeriodEndOk() (*JSONTime, bool) {
	if o == nil || o.CurrentPeriodEnd == nil {
		return nil, false
	}
	return o.CurrentPeriodEnd, true
}

// HasCurrentPeriodEnd returns a boolean if a field has been set.
func (o *NextCharge) HasCurrentPeriodEnd() bool {
	if o != nil && o.CurrentPeriodEnd != nil {
		return true
	}

	return false
}

// SetCurrentPeriodEnd gets a reference to the given JSONTime and assigns it to the CurrentPeriodEnd field.
func (o *NextCharge) SetCurrentPeriodEnd(v JSONTime) {
	o.CurrentPeriodEnd = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *NextCharge) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextCharge) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *NextCharge) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *NextCharge) SetInterval(v string) {
	o.Interval = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *NextCharge) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextCharge) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *NextCharge) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *NextCharge) SetTotal(v int64) {
	o.Total = &v
}

func (o NextCharge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.CurrentPeriodEnd != nil {
		toSerialize["currentPeriodEnd"] = o.CurrentPeriodEnd
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableNextCharge struct {
	value *NextCharge
	isSet bool
}

func (v NullableNextCharge) Get() *NextCharge {
	return v.value
}

func (v *NullableNextCharge) Set(val *NextCharge) {
	v.value = val
	v.isSet = true
}

func (v NullableNextCharge) IsSet() bool {
	return v.isSet
}

func (v *NullableNextCharge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextCharge(val *NextCharge) *NullableNextCharge {
	return &NullableNextCharge{value: val, isSet: true}
}

func (v NullableNextCharge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextCharge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
