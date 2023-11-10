/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25652-944cf1af37c9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTCountDimensionDisplayData1778 struct for BTCountDimensionDisplayData1778
type BTCountDimensionDisplayData1778 struct {
	BtType               *string         `json:"btType,omitempty"`
	CoordinateSystem     *BTMatrix3x3340 `json:"coordinateSystem,omitempty"`
	FeatureId            *string         `json:"featureId,omitempty"`
	HasMaximumLimit_     *bool           `json:"hasMaximumLimit,omitempty"`
	HasMinimumLimit_     *bool           `json:"hasMinimumLimit,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	IsAssociatedWithFlat *bool           `json:"isAssociatedWithFlat,omitempty"`
	IsDriven             *bool           `json:"isDriven,omitempty"`
	IsOverDefined        *bool           `json:"isOverDefined,omitempty"`
	MaximumLimit         *float64        `json:"maximumLimit,omitempty"`
	MinimumLimit         *float64        `json:"minimumLimit,omitempty"`
	ParameterId          *string         `json:"parameterId,omitempty"`
	PlaneMatrix          *BTBSMatrix386  `json:"planeMatrix,omitempty"`
	Value                *float64        `json:"value,omitempty"`
	PositionX            *float64        `json:"positionX,omitempty"`
	PositionY            *float64        `json:"positionY,omitempty"`
}

// NewBTCountDimensionDisplayData1778 instantiates a new BTCountDimensionDisplayData1778 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCountDimensionDisplayData1778() *BTCountDimensionDisplayData1778 {
	this := BTCountDimensionDisplayData1778{}
	return &this
}

// NewBTCountDimensionDisplayData1778WithDefaults instantiates a new BTCountDimensionDisplayData1778 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCountDimensionDisplayData1778WithDefaults() *BTCountDimensionDisplayData1778 {
	this := BTCountDimensionDisplayData1778{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCountDimensionDisplayData1778) SetBtType(v string) {
	o.BtType = &v
}

// GetCoordinateSystem returns the CoordinateSystem field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetCoordinateSystem() BTMatrix3x3340 {
	if o == nil || o.CoordinateSystem == nil {
		var ret BTMatrix3x3340
		return ret
	}
	return *o.CoordinateSystem
}

// GetCoordinateSystemOk returns a tuple with the CoordinateSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetCoordinateSystemOk() (*BTMatrix3x3340, bool) {
	if o == nil || o.CoordinateSystem == nil {
		return nil, false
	}
	return o.CoordinateSystem, true
}

// HasCoordinateSystem returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasCoordinateSystem() bool {
	if o != nil && o.CoordinateSystem != nil {
		return true
	}

	return false
}

// SetCoordinateSystem gets a reference to the given BTMatrix3x3340 and assigns it to the CoordinateSystem field.
func (o *BTCountDimensionDisplayData1778) SetCoordinateSystem(v BTMatrix3x3340) {
	o.CoordinateSystem = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTCountDimensionDisplayData1778) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetHasMaximumLimit_ returns the HasMaximumLimit_ field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetHasMaximumLimit_() bool {
	if o == nil || o.HasMaximumLimit_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMaximumLimit_
}

// GetHasMaximumLimit_Ok returns a tuple with the HasMaximumLimit_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetHasMaximumLimit_Ok() (*bool, bool) {
	if o == nil || o.HasMaximumLimit_ == nil {
		return nil, false
	}
	return o.HasMaximumLimit_, true
}

// HasHasMaximumLimit_ returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasHasMaximumLimit_() bool {
	if o != nil && o.HasMaximumLimit_ != nil {
		return true
	}

	return false
}

// SetHasMaximumLimit_ gets a reference to the given bool and assigns it to the HasMaximumLimit_ field.
func (o *BTCountDimensionDisplayData1778) SetHasMaximumLimit_(v bool) {
	o.HasMaximumLimit_ = &v
}

// GetHasMinimumLimit_ returns the HasMinimumLimit_ field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetHasMinimumLimit_() bool {
	if o == nil || o.HasMinimumLimit_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMinimumLimit_
}

// GetHasMinimumLimit_Ok returns a tuple with the HasMinimumLimit_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetHasMinimumLimit_Ok() (*bool, bool) {
	if o == nil || o.HasMinimumLimit_ == nil {
		return nil, false
	}
	return o.HasMinimumLimit_, true
}

// HasHasMinimumLimit_ returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasHasMinimumLimit_() bool {
	if o != nil && o.HasMinimumLimit_ != nil {
		return true
	}

	return false
}

// SetHasMinimumLimit_ gets a reference to the given bool and assigns it to the HasMinimumLimit_ field.
func (o *BTCountDimensionDisplayData1778) SetHasMinimumLimit_(v bool) {
	o.HasMinimumLimit_ = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCountDimensionDisplayData1778) SetId(v string) {
	o.Id = &v
}

// GetIsAssociatedWithFlat returns the IsAssociatedWithFlat field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetIsAssociatedWithFlat() bool {
	if o == nil || o.IsAssociatedWithFlat == nil {
		var ret bool
		return ret
	}
	return *o.IsAssociatedWithFlat
}

// GetIsAssociatedWithFlatOk returns a tuple with the IsAssociatedWithFlat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetIsAssociatedWithFlatOk() (*bool, bool) {
	if o == nil || o.IsAssociatedWithFlat == nil {
		return nil, false
	}
	return o.IsAssociatedWithFlat, true
}

// HasIsAssociatedWithFlat returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasIsAssociatedWithFlat() bool {
	if o != nil && o.IsAssociatedWithFlat != nil {
		return true
	}

	return false
}

// SetIsAssociatedWithFlat gets a reference to the given bool and assigns it to the IsAssociatedWithFlat field.
func (o *BTCountDimensionDisplayData1778) SetIsAssociatedWithFlat(v bool) {
	o.IsAssociatedWithFlat = &v
}

// GetIsDriven returns the IsDriven field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetIsDriven() bool {
	if o == nil || o.IsDriven == nil {
		var ret bool
		return ret
	}
	return *o.IsDriven
}

// GetIsDrivenOk returns a tuple with the IsDriven field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetIsDrivenOk() (*bool, bool) {
	if o == nil || o.IsDriven == nil {
		return nil, false
	}
	return o.IsDriven, true
}

// HasIsDriven returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasIsDriven() bool {
	if o != nil && o.IsDriven != nil {
		return true
	}

	return false
}

// SetIsDriven gets a reference to the given bool and assigns it to the IsDriven field.
func (o *BTCountDimensionDisplayData1778) SetIsDriven(v bool) {
	o.IsDriven = &v
}

// GetIsOverDefined returns the IsOverDefined field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetIsOverDefined() bool {
	if o == nil || o.IsOverDefined == nil {
		var ret bool
		return ret
	}
	return *o.IsOverDefined
}

// GetIsOverDefinedOk returns a tuple with the IsOverDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetIsOverDefinedOk() (*bool, bool) {
	if o == nil || o.IsOverDefined == nil {
		return nil, false
	}
	return o.IsOverDefined, true
}

// HasIsOverDefined returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasIsOverDefined() bool {
	if o != nil && o.IsOverDefined != nil {
		return true
	}

	return false
}

// SetIsOverDefined gets a reference to the given bool and assigns it to the IsOverDefined field.
func (o *BTCountDimensionDisplayData1778) SetIsOverDefined(v bool) {
	o.IsOverDefined = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetMaximumLimit() float64 {
	if o == nil || o.MaximumLimit == nil {
		var ret float64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetMaximumLimitOk() (*float64, bool) {
	if o == nil || o.MaximumLimit == nil {
		return nil, false
	}
	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasMaximumLimit() bool {
	if o != nil && o.MaximumLimit != nil {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given float64 and assigns it to the MaximumLimit field.
func (o *BTCountDimensionDisplayData1778) SetMaximumLimit(v float64) {
	o.MaximumLimit = &v
}

// GetMinimumLimit returns the MinimumLimit field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetMinimumLimit() float64 {
	if o == nil || o.MinimumLimit == nil {
		var ret float64
		return ret
	}
	return *o.MinimumLimit
}

// GetMinimumLimitOk returns a tuple with the MinimumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetMinimumLimitOk() (*float64, bool) {
	if o == nil || o.MinimumLimit == nil {
		return nil, false
	}
	return o.MinimumLimit, true
}

// HasMinimumLimit returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasMinimumLimit() bool {
	if o != nil && o.MinimumLimit != nil {
		return true
	}

	return false
}

// SetMinimumLimit gets a reference to the given float64 and assigns it to the MinimumLimit field.
func (o *BTCountDimensionDisplayData1778) SetMinimumLimit(v float64) {
	o.MinimumLimit = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTCountDimensionDisplayData1778) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetPlaneMatrix returns the PlaneMatrix field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetPlaneMatrix() BTBSMatrix386 {
	if o == nil || o.PlaneMatrix == nil {
		var ret BTBSMatrix386
		return ret
	}
	return *o.PlaneMatrix
}

// GetPlaneMatrixOk returns a tuple with the PlaneMatrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetPlaneMatrixOk() (*BTBSMatrix386, bool) {
	if o == nil || o.PlaneMatrix == nil {
		return nil, false
	}
	return o.PlaneMatrix, true
}

// HasPlaneMatrix returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasPlaneMatrix() bool {
	if o != nil && o.PlaneMatrix != nil {
		return true
	}

	return false
}

// SetPlaneMatrix gets a reference to the given BTBSMatrix386 and assigns it to the PlaneMatrix field.
func (o *BTCountDimensionDisplayData1778) SetPlaneMatrix(v BTBSMatrix386) {
	o.PlaneMatrix = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BTCountDimensionDisplayData1778) SetValue(v float64) {
	o.Value = &v
}

// GetPositionX returns the PositionX field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetPositionX() float64 {
	if o == nil || o.PositionX == nil {
		var ret float64
		return ret
	}
	return *o.PositionX
}

// GetPositionXOk returns a tuple with the PositionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetPositionXOk() (*float64, bool) {
	if o == nil || o.PositionX == nil {
		return nil, false
	}
	return o.PositionX, true
}

// HasPositionX returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasPositionX() bool {
	if o != nil && o.PositionX != nil {
		return true
	}

	return false
}

// SetPositionX gets a reference to the given float64 and assigns it to the PositionX field.
func (o *BTCountDimensionDisplayData1778) SetPositionX(v float64) {
	o.PositionX = &v
}

// GetPositionY returns the PositionY field value if set, zero value otherwise.
func (o *BTCountDimensionDisplayData1778) GetPositionY() float64 {
	if o == nil || o.PositionY == nil {
		var ret float64
		return ret
	}
	return *o.PositionY
}

// GetPositionYOk returns a tuple with the PositionY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCountDimensionDisplayData1778) GetPositionYOk() (*float64, bool) {
	if o == nil || o.PositionY == nil {
		return nil, false
	}
	return o.PositionY, true
}

// HasPositionY returns a boolean if a field has been set.
func (o *BTCountDimensionDisplayData1778) HasPositionY() bool {
	if o != nil && o.PositionY != nil {
		return true
	}

	return false
}

// SetPositionY gets a reference to the given float64 and assigns it to the PositionY field.
func (o *BTCountDimensionDisplayData1778) SetPositionY(v float64) {
	o.PositionY = &v
}

func (o BTCountDimensionDisplayData1778) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CoordinateSystem != nil {
		toSerialize["coordinateSystem"] = o.CoordinateSystem
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.HasMaximumLimit_ != nil {
		toSerialize["hasMaximumLimit"] = o.HasMaximumLimit_
	}
	if o.HasMinimumLimit_ != nil {
		toSerialize["hasMinimumLimit"] = o.HasMinimumLimit_
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsAssociatedWithFlat != nil {
		toSerialize["isAssociatedWithFlat"] = o.IsAssociatedWithFlat
	}
	if o.IsDriven != nil {
		toSerialize["isDriven"] = o.IsDriven
	}
	if o.IsOverDefined != nil {
		toSerialize["isOverDefined"] = o.IsOverDefined
	}
	if o.MaximumLimit != nil {
		toSerialize["maximumLimit"] = o.MaximumLimit
	}
	if o.MinimumLimit != nil {
		toSerialize["minimumLimit"] = o.MinimumLimit
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.PlaneMatrix != nil {
		toSerialize["planeMatrix"] = o.PlaneMatrix
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.PositionX != nil {
		toSerialize["positionX"] = o.PositionX
	}
	if o.PositionY != nil {
		toSerialize["positionY"] = o.PositionY
	}
	return json.Marshal(toSerialize)
}

type NullableBTCountDimensionDisplayData1778 struct {
	value *BTCountDimensionDisplayData1778
	isSet bool
}

func (v NullableBTCountDimensionDisplayData1778) Get() *BTCountDimensionDisplayData1778 {
	return v.value
}

func (v *NullableBTCountDimensionDisplayData1778) Set(val *BTCountDimensionDisplayData1778) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCountDimensionDisplayData1778) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCountDimensionDisplayData1778) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCountDimensionDisplayData1778(val *BTCountDimensionDisplayData1778) *NullableBTCountDimensionDisplayData1778 {
	return &NullableBTCountDimensionDisplayData1778{value: val, isSet: true}
}

func (v NullableBTCountDimensionDisplayData1778) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCountDimensionDisplayData1778) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
