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

// BTBaseEntityAppearanceEntry3607 struct for BTBaseEntityAppearanceEntry3607
type BTBaseEntityAppearanceEntry3607 struct {
	AffectedDeterministicIds []string                  `json:"affectedDeterministicIds,omitempty"`
	Appearance               *BTGraphicsAppearance1152 `json:"appearance,omitempty"`
	BtType                   *string                   `json:"btType,omitempty"`
	Source                   *BTPartMetadataSource2895 `json:"source,omitempty"`
}

// NewBTBaseEntityAppearanceEntry3607 instantiates a new BTBaseEntityAppearanceEntry3607 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBaseEntityAppearanceEntry3607() *BTBaseEntityAppearanceEntry3607 {
	this := BTBaseEntityAppearanceEntry3607{}
	return &this
}

// NewBTBaseEntityAppearanceEntry3607WithDefaults instantiates a new BTBaseEntityAppearanceEntry3607 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBaseEntityAppearanceEntry3607WithDefaults() *BTBaseEntityAppearanceEntry3607 {
	this := BTBaseEntityAppearanceEntry3607{}
	return &this
}

// GetAffectedDeterministicIds returns the AffectedDeterministicIds field value if set, zero value otherwise.
func (o *BTBaseEntityAppearanceEntry3607) GetAffectedDeterministicIds() []string {
	if o == nil || o.AffectedDeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.AffectedDeterministicIds
}

// GetAffectedDeterministicIdsOk returns a tuple with the AffectedDeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseEntityAppearanceEntry3607) GetAffectedDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.AffectedDeterministicIds == nil {
		return nil, false
	}
	return o.AffectedDeterministicIds, true
}

// HasAffectedDeterministicIds returns a boolean if a field has been set.
func (o *BTBaseEntityAppearanceEntry3607) HasAffectedDeterministicIds() bool {
	if o != nil && o.AffectedDeterministicIds != nil {
		return true
	}

	return false
}

// SetAffectedDeterministicIds gets a reference to the given []string and assigns it to the AffectedDeterministicIds field.
func (o *BTBaseEntityAppearanceEntry3607) SetAffectedDeterministicIds(v []string) {
	o.AffectedDeterministicIds = v
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTBaseEntityAppearanceEntry3607) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseEntityAppearanceEntry3607) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTBaseEntityAppearanceEntry3607) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTBaseEntityAppearanceEntry3607) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBaseEntityAppearanceEntry3607) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseEntityAppearanceEntry3607) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBaseEntityAppearanceEntry3607) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBaseEntityAppearanceEntry3607) SetBtType(v string) {
	o.BtType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BTBaseEntityAppearanceEntry3607) GetSource() BTPartMetadataSource2895 {
	if o == nil || o.Source == nil {
		var ret BTPartMetadataSource2895
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseEntityAppearanceEntry3607) GetSourceOk() (*BTPartMetadataSource2895, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BTBaseEntityAppearanceEntry3607) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given BTPartMetadataSource2895 and assigns it to the Source field.
func (o *BTBaseEntityAppearanceEntry3607) SetSource(v BTPartMetadataSource2895) {
	o.Source = &v
}

func (o BTBaseEntityAppearanceEntry3607) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AffectedDeterministicIds != nil {
		toSerialize["affectedDeterministicIds"] = o.AffectedDeterministicIds
	}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableBTBaseEntityAppearanceEntry3607 struct {
	value *BTBaseEntityAppearanceEntry3607
	isSet bool
}

func (v NullableBTBaseEntityAppearanceEntry3607) Get() *BTBaseEntityAppearanceEntry3607 {
	return v.value
}

func (v *NullableBTBaseEntityAppearanceEntry3607) Set(val *BTBaseEntityAppearanceEntry3607) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBaseEntityAppearanceEntry3607) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBaseEntityAppearanceEntry3607) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBaseEntityAppearanceEntry3607(val *BTBaseEntityAppearanceEntry3607) *NullableBTBaseEntityAppearanceEntry3607 {
	return &NullableBTBaseEntityAppearanceEntry3607{value: val, isSet: true}
}

func (v NullableBTBaseEntityAppearanceEntry3607) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBaseEntityAppearanceEntry3607) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
