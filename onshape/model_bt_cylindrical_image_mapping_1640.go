/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.26754-ceeaad064d4a
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCylindricalImageMapping1640 struct for BTCylindricalImageMapping1640
type BTCylindricalImageMapping1640 struct {
	BtType           *string                `json:"btType,omitempty"`
	DeterministicIds []string               `json:"deterministicIds,omitempty"`
	UvTransform      *BTMatrix3x3340        `json:"uvTransform,omitempty"`
	CylinderSystem   *BTCoordinateSystem387 `json:"cylinderSystem,omitempty"`
	Radius           *float32               `json:"radius,omitempty"`
}

// NewBTCylindricalImageMapping1640 instantiates a new BTCylindricalImageMapping1640 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCylindricalImageMapping1640() *BTCylindricalImageMapping1640 {
	this := BTCylindricalImageMapping1640{}
	return &this
}

// NewBTCylindricalImageMapping1640WithDefaults instantiates a new BTCylindricalImageMapping1640 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCylindricalImageMapping1640WithDefaults() *BTCylindricalImageMapping1640 {
	this := BTCylindricalImageMapping1640{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCylindricalImageMapping1640) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCylindricalImageMapping1640) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCylindricalImageMapping1640) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCylindricalImageMapping1640) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTCylindricalImageMapping1640) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCylindricalImageMapping1640) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTCylindricalImageMapping1640) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTCylindricalImageMapping1640) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetUvTransform returns the UvTransform field value if set, zero value otherwise.
func (o *BTCylindricalImageMapping1640) GetUvTransform() BTMatrix3x3340 {
	if o == nil || o.UvTransform == nil {
		var ret BTMatrix3x3340
		return ret
	}
	return *o.UvTransform
}

// GetUvTransformOk returns a tuple with the UvTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCylindricalImageMapping1640) GetUvTransformOk() (*BTMatrix3x3340, bool) {
	if o == nil || o.UvTransform == nil {
		return nil, false
	}
	return o.UvTransform, true
}

// HasUvTransform returns a boolean if a field has been set.
func (o *BTCylindricalImageMapping1640) HasUvTransform() bool {
	if o != nil && o.UvTransform != nil {
		return true
	}

	return false
}

// SetUvTransform gets a reference to the given BTMatrix3x3340 and assigns it to the UvTransform field.
func (o *BTCylindricalImageMapping1640) SetUvTransform(v BTMatrix3x3340) {
	o.UvTransform = &v
}

// GetCylinderSystem returns the CylinderSystem field value if set, zero value otherwise.
func (o *BTCylindricalImageMapping1640) GetCylinderSystem() BTCoordinateSystem387 {
	if o == nil || o.CylinderSystem == nil {
		var ret BTCoordinateSystem387
		return ret
	}
	return *o.CylinderSystem
}

// GetCylinderSystemOk returns a tuple with the CylinderSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCylindricalImageMapping1640) GetCylinderSystemOk() (*BTCoordinateSystem387, bool) {
	if o == nil || o.CylinderSystem == nil {
		return nil, false
	}
	return o.CylinderSystem, true
}

// HasCylinderSystem returns a boolean if a field has been set.
func (o *BTCylindricalImageMapping1640) HasCylinderSystem() bool {
	if o != nil && o.CylinderSystem != nil {
		return true
	}

	return false
}

// SetCylinderSystem gets a reference to the given BTCoordinateSystem387 and assigns it to the CylinderSystem field.
func (o *BTCylindricalImageMapping1640) SetCylinderSystem(v BTCoordinateSystem387) {
	o.CylinderSystem = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BTCylindricalImageMapping1640) GetRadius() float32 {
	if o == nil || o.Radius == nil {
		var ret float32
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCylindricalImageMapping1640) GetRadiusOk() (*float32, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BTCylindricalImageMapping1640) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float32 and assigns it to the Radius field.
func (o *BTCylindricalImageMapping1640) SetRadius(v float32) {
	o.Radius = &v
}

func (o BTCylindricalImageMapping1640) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.UvTransform != nil {
		toSerialize["uvTransform"] = o.UvTransform
	}
	if o.CylinderSystem != nil {
		toSerialize["cylinderSystem"] = o.CylinderSystem
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	return json.Marshal(toSerialize)
}

type NullableBTCylindricalImageMapping1640 struct {
	value *BTCylindricalImageMapping1640
	isSet bool
}

func (v NullableBTCylindricalImageMapping1640) Get() *BTCylindricalImageMapping1640 {
	return v.value
}

func (v *NullableBTCylindricalImageMapping1640) Set(val *BTCylindricalImageMapping1640) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCylindricalImageMapping1640) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCylindricalImageMapping1640) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCylindricalImageMapping1640(val *BTCylindricalImageMapping1640) *NullableBTCylindricalImageMapping1640 {
	return &NullableBTCylindricalImageMapping1640{value: val, isSet: true}
}

func (v NullableBTCylindricalImageMapping1640) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCylindricalImageMapping1640) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
