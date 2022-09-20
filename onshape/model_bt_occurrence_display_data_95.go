/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6524-ea93b01144ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOccurrenceDisplayData95 struct for BTOccurrenceDisplayData95
type BTOccurrenceDisplayData95 struct {
	BtType                          *string             `json:"btType,omitempty"`
	ElementId                       *string             `json:"elementId,omitempty"`
	ForceHighestQualityTessellation *bool               `json:"forceHighestQualityTessellation,omitempty"`
	FullElementId                   *BTFullElementId756 `json:"fullElementId,omitempty"`
	IsHidden                        *bool               `json:"isHidden,omitempty"`
	IsPatternDescendant             *bool               `json:"isPatternDescendant,omitempty"`
	OccurrenceData                  *BTOccurrenceData75 `json:"occurrenceData,omitempty"`
	PartIds                         []string            `json:"partIds,omitempty"`
}

// NewBTOccurrenceDisplayData95 instantiates a new BTOccurrenceDisplayData95 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOccurrenceDisplayData95() *BTOccurrenceDisplayData95 {
	this := BTOccurrenceDisplayData95{}
	return &this
}

// NewBTOccurrenceDisplayData95WithDefaults instantiates a new BTOccurrenceDisplayData95 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOccurrenceDisplayData95WithDefaults() *BTOccurrenceDisplayData95 {
	this := BTOccurrenceDisplayData95{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOccurrenceDisplayData95) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceDisplayData95) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOccurrenceDisplayData95) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOccurrenceDisplayData95) SetBtType(v string) {
	o.BtType = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTOccurrenceDisplayData95) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceDisplayData95) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTOccurrenceDisplayData95) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTOccurrenceDisplayData95) SetElementId(v string) {
	o.ElementId = &v
}

// GetForceHighestQualityTessellation returns the ForceHighestQualityTessellation field value if set, zero value otherwise.
func (o *BTOccurrenceDisplayData95) GetForceHighestQualityTessellation() bool {
	if o == nil || o.ForceHighestQualityTessellation == nil {
		var ret bool
		return ret
	}
	return *o.ForceHighestQualityTessellation
}

// GetForceHighestQualityTessellationOk returns a tuple with the ForceHighestQualityTessellation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceDisplayData95) GetForceHighestQualityTessellationOk() (*bool, bool) {
	if o == nil || o.ForceHighestQualityTessellation == nil {
		return nil, false
	}
	return o.ForceHighestQualityTessellation, true
}

// HasForceHighestQualityTessellation returns a boolean if a field has been set.
func (o *BTOccurrenceDisplayData95) HasForceHighestQualityTessellation() bool {
	if o != nil && o.ForceHighestQualityTessellation != nil {
		return true
	}

	return false
}

// SetForceHighestQualityTessellation gets a reference to the given bool and assigns it to the ForceHighestQualityTessellation field.
func (o *BTOccurrenceDisplayData95) SetForceHighestQualityTessellation(v bool) {
	o.ForceHighestQualityTessellation = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTOccurrenceDisplayData95) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceDisplayData95) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTOccurrenceDisplayData95) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTOccurrenceDisplayData95) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *BTOccurrenceDisplayData95) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceDisplayData95) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *BTOccurrenceDisplayData95) HasIsHidden() bool {
	if o != nil && o.IsHidden != nil {
		return true
	}

	return false
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *BTOccurrenceDisplayData95) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetIsPatternDescendant returns the IsPatternDescendant field value if set, zero value otherwise.
func (o *BTOccurrenceDisplayData95) GetIsPatternDescendant() bool {
	if o == nil || o.IsPatternDescendant == nil {
		var ret bool
		return ret
	}
	return *o.IsPatternDescendant
}

// GetIsPatternDescendantOk returns a tuple with the IsPatternDescendant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceDisplayData95) GetIsPatternDescendantOk() (*bool, bool) {
	if o == nil || o.IsPatternDescendant == nil {
		return nil, false
	}
	return o.IsPatternDescendant, true
}

// HasIsPatternDescendant returns a boolean if a field has been set.
func (o *BTOccurrenceDisplayData95) HasIsPatternDescendant() bool {
	if o != nil && o.IsPatternDescendant != nil {
		return true
	}

	return false
}

// SetIsPatternDescendant gets a reference to the given bool and assigns it to the IsPatternDescendant field.
func (o *BTOccurrenceDisplayData95) SetIsPatternDescendant(v bool) {
	o.IsPatternDescendant = &v
}

// GetOccurrenceData returns the OccurrenceData field value if set, zero value otherwise.
func (o *BTOccurrenceDisplayData95) GetOccurrenceData() BTOccurrenceData75 {
	if o == nil || o.OccurrenceData == nil {
		var ret BTOccurrenceData75
		return ret
	}
	return *o.OccurrenceData
}

// GetOccurrenceDataOk returns a tuple with the OccurrenceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceDisplayData95) GetOccurrenceDataOk() (*BTOccurrenceData75, bool) {
	if o == nil || o.OccurrenceData == nil {
		return nil, false
	}
	return o.OccurrenceData, true
}

// HasOccurrenceData returns a boolean if a field has been set.
func (o *BTOccurrenceDisplayData95) HasOccurrenceData() bool {
	if o != nil && o.OccurrenceData != nil {
		return true
	}

	return false
}

// SetOccurrenceData gets a reference to the given BTOccurrenceData75 and assigns it to the OccurrenceData field.
func (o *BTOccurrenceDisplayData95) SetOccurrenceData(v BTOccurrenceData75) {
	o.OccurrenceData = &v
}

// GetPartIds returns the PartIds field value if set, zero value otherwise.
func (o *BTOccurrenceDisplayData95) GetPartIds() []string {
	if o == nil || o.PartIds == nil {
		var ret []string
		return ret
	}
	return o.PartIds
}

// GetPartIdsOk returns a tuple with the PartIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceDisplayData95) GetPartIdsOk() ([]string, bool) {
	if o == nil || o.PartIds == nil {
		return nil, false
	}
	return o.PartIds, true
}

// HasPartIds returns a boolean if a field has been set.
func (o *BTOccurrenceDisplayData95) HasPartIds() bool {
	if o != nil && o.PartIds != nil {
		return true
	}

	return false
}

// SetPartIds gets a reference to the given []string and assigns it to the PartIds field.
func (o *BTOccurrenceDisplayData95) SetPartIds(v []string) {
	o.PartIds = v
}

func (o BTOccurrenceDisplayData95) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ForceHighestQualityTessellation != nil {
		toSerialize["forceHighestQualityTessellation"] = o.ForceHighestQualityTessellation
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.IsHidden != nil {
		toSerialize["isHidden"] = o.IsHidden
	}
	if o.IsPatternDescendant != nil {
		toSerialize["isPatternDescendant"] = o.IsPatternDescendant
	}
	if o.OccurrenceData != nil {
		toSerialize["occurrenceData"] = o.OccurrenceData
	}
	if o.PartIds != nil {
		toSerialize["partIds"] = o.PartIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTOccurrenceDisplayData95 struct {
	value *BTOccurrenceDisplayData95
	isSet bool
}

func (v NullableBTOccurrenceDisplayData95) Get() *BTOccurrenceDisplayData95 {
	return v.value
}

func (v *NullableBTOccurrenceDisplayData95) Set(val *BTOccurrenceDisplayData95) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOccurrenceDisplayData95) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOccurrenceDisplayData95) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOccurrenceDisplayData95(val *BTOccurrenceDisplayData95) *NullableBTOccurrenceDisplayData95 {
	return &NullableBTOccurrenceDisplayData95{value: val, isSet: true}
}

func (v NullableBTOccurrenceDisplayData95) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOccurrenceDisplayData95) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
