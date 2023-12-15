/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27678-64d64396ca66
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTGeometryMateDisplayData1050 struct for BTGeometryMateDisplayData1050
type BTGeometryMateDisplayData1050 struct {
	BtType                *string                          `json:"btType,omitempty"`
	FirstDeterministicId  *string                          `json:"firstDeterministicId,omitempty"`
	FirstOccurrence       *BTOccurrence74                  `json:"firstOccurrence,omitempty"`
	Hidden                *bool                            `json:"hidden,omitempty"`
	IsDerivedFeature      *bool                            `json:"isDerivedFeature,omitempty"`
	Location              *BTCoordinateSystem387           `json:"location,omitempty"`
	NodeId                *string                          `json:"nodeId,omitempty"`
	OwnerOccurrence       *BTOccurrence74                  `json:"ownerOccurrence,omitempty"`
	SecondDeterministicId *string                          `json:"secondDeterministicId,omitempty"`
	SecondOccurrence      *BTOccurrence74                  `json:"secondOccurrence,omitempty"`
	Status                *GBTAssemblyFeatureDisplayStatus `json:"status,omitempty"`
}

// NewBTGeometryMateDisplayData1050 instantiates a new BTGeometryMateDisplayData1050 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGeometryMateDisplayData1050() *BTGeometryMateDisplayData1050 {
	this := BTGeometryMateDisplayData1050{}
	return &this
}

// NewBTGeometryMateDisplayData1050WithDefaults instantiates a new BTGeometryMateDisplayData1050 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGeometryMateDisplayData1050WithDefaults() *BTGeometryMateDisplayData1050 {
	this := BTGeometryMateDisplayData1050{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTGeometryMateDisplayData1050) SetBtType(v string) {
	o.BtType = &v
}

// GetFirstDeterministicId returns the FirstDeterministicId field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetFirstDeterministicId() string {
	if o == nil || o.FirstDeterministicId == nil {
		var ret string
		return ret
	}
	return *o.FirstDeterministicId
}

// GetFirstDeterministicIdOk returns a tuple with the FirstDeterministicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetFirstDeterministicIdOk() (*string, bool) {
	if o == nil || o.FirstDeterministicId == nil {
		return nil, false
	}
	return o.FirstDeterministicId, true
}

// HasFirstDeterministicId returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasFirstDeterministicId() bool {
	if o != nil && o.FirstDeterministicId != nil {
		return true
	}

	return false
}

// SetFirstDeterministicId gets a reference to the given string and assigns it to the FirstDeterministicId field.
func (o *BTGeometryMateDisplayData1050) SetFirstDeterministicId(v string) {
	o.FirstDeterministicId = &v
}

// GetFirstOccurrence returns the FirstOccurrence field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetFirstOccurrence() BTOccurrence74 {
	if o == nil || o.FirstOccurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.FirstOccurrence
}

// GetFirstOccurrenceOk returns a tuple with the FirstOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetFirstOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.FirstOccurrence == nil {
		return nil, false
	}
	return o.FirstOccurrence, true
}

// HasFirstOccurrence returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasFirstOccurrence() bool {
	if o != nil && o.FirstOccurrence != nil {
		return true
	}

	return false
}

// SetFirstOccurrence gets a reference to the given BTOccurrence74 and assigns it to the FirstOccurrence field.
func (o *BTGeometryMateDisplayData1050) SetFirstOccurrence(v BTOccurrence74) {
	o.FirstOccurrence = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTGeometryMateDisplayData1050) SetHidden(v bool) {
	o.Hidden = &v
}

// GetIsDerivedFeature returns the IsDerivedFeature field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetIsDerivedFeature() bool {
	if o == nil || o.IsDerivedFeature == nil {
		var ret bool
		return ret
	}
	return *o.IsDerivedFeature
}

// GetIsDerivedFeatureOk returns a tuple with the IsDerivedFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetIsDerivedFeatureOk() (*bool, bool) {
	if o == nil || o.IsDerivedFeature == nil {
		return nil, false
	}
	return o.IsDerivedFeature, true
}

// HasIsDerivedFeature returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasIsDerivedFeature() bool {
	if o != nil && o.IsDerivedFeature != nil {
		return true
	}

	return false
}

// SetIsDerivedFeature gets a reference to the given bool and assigns it to the IsDerivedFeature field.
func (o *BTGeometryMateDisplayData1050) SetIsDerivedFeature(v bool) {
	o.IsDerivedFeature = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetLocation() BTCoordinateSystem387 {
	if o == nil || o.Location == nil {
		var ret BTCoordinateSystem387
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetLocationOk() (*BTCoordinateSystem387, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BTCoordinateSystem387 and assigns it to the Location field.
func (o *BTGeometryMateDisplayData1050) SetLocation(v BTCoordinateSystem387) {
	o.Location = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTGeometryMateDisplayData1050) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOwnerOccurrence returns the OwnerOccurrence field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetOwnerOccurrence() BTOccurrence74 {
	if o == nil || o.OwnerOccurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OwnerOccurrence
}

// GetOwnerOccurrenceOk returns a tuple with the OwnerOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetOwnerOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.OwnerOccurrence == nil {
		return nil, false
	}
	return o.OwnerOccurrence, true
}

// HasOwnerOccurrence returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasOwnerOccurrence() bool {
	if o != nil && o.OwnerOccurrence != nil {
		return true
	}

	return false
}

// SetOwnerOccurrence gets a reference to the given BTOccurrence74 and assigns it to the OwnerOccurrence field.
func (o *BTGeometryMateDisplayData1050) SetOwnerOccurrence(v BTOccurrence74) {
	o.OwnerOccurrence = &v
}

// GetSecondDeterministicId returns the SecondDeterministicId field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetSecondDeterministicId() string {
	if o == nil || o.SecondDeterministicId == nil {
		var ret string
		return ret
	}
	return *o.SecondDeterministicId
}

// GetSecondDeterministicIdOk returns a tuple with the SecondDeterministicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetSecondDeterministicIdOk() (*string, bool) {
	if o == nil || o.SecondDeterministicId == nil {
		return nil, false
	}
	return o.SecondDeterministicId, true
}

// HasSecondDeterministicId returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasSecondDeterministicId() bool {
	if o != nil && o.SecondDeterministicId != nil {
		return true
	}

	return false
}

// SetSecondDeterministicId gets a reference to the given string and assigns it to the SecondDeterministicId field.
func (o *BTGeometryMateDisplayData1050) SetSecondDeterministicId(v string) {
	o.SecondDeterministicId = &v
}

// GetSecondOccurrence returns the SecondOccurrence field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetSecondOccurrence() BTOccurrence74 {
	if o == nil || o.SecondOccurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.SecondOccurrence
}

// GetSecondOccurrenceOk returns a tuple with the SecondOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetSecondOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.SecondOccurrence == nil {
		return nil, false
	}
	return o.SecondOccurrence, true
}

// HasSecondOccurrence returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasSecondOccurrence() bool {
	if o != nil && o.SecondOccurrence != nil {
		return true
	}

	return false
}

// SetSecondOccurrence gets a reference to the given BTOccurrence74 and assigns it to the SecondOccurrence field.
func (o *BTGeometryMateDisplayData1050) SetSecondOccurrence(v BTOccurrence74) {
	o.SecondOccurrence = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTGeometryMateDisplayData1050) GetStatus() GBTAssemblyFeatureDisplayStatus {
	if o == nil || o.Status == nil {
		var ret GBTAssemblyFeatureDisplayStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryMateDisplayData1050) GetStatusOk() (*GBTAssemblyFeatureDisplayStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTGeometryMateDisplayData1050) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GBTAssemblyFeatureDisplayStatus and assigns it to the Status field.
func (o *BTGeometryMateDisplayData1050) SetStatus(v GBTAssemblyFeatureDisplayStatus) {
	o.Status = &v
}

func (o BTGeometryMateDisplayData1050) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FirstDeterministicId != nil {
		toSerialize["firstDeterministicId"] = o.FirstDeterministicId
	}
	if o.FirstOccurrence != nil {
		toSerialize["firstOccurrence"] = o.FirstOccurrence
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.IsDerivedFeature != nil {
		toSerialize["isDerivedFeature"] = o.IsDerivedFeature
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.OwnerOccurrence != nil {
		toSerialize["ownerOccurrence"] = o.OwnerOccurrence
	}
	if o.SecondDeterministicId != nil {
		toSerialize["secondDeterministicId"] = o.SecondDeterministicId
	}
	if o.SecondOccurrence != nil {
		toSerialize["secondOccurrence"] = o.SecondOccurrence
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBTGeometryMateDisplayData1050 struct {
	value *BTGeometryMateDisplayData1050
	isSet bool
}

func (v NullableBTGeometryMateDisplayData1050) Get() *BTGeometryMateDisplayData1050 {
	return v.value
}

func (v *NullableBTGeometryMateDisplayData1050) Set(val *BTGeometryMateDisplayData1050) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGeometryMateDisplayData1050) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGeometryMateDisplayData1050) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGeometryMateDisplayData1050(val *BTGeometryMateDisplayData1050) *NullableBTGeometryMateDisplayData1050 {
	return &NullableBTGeometryMateDisplayData1050{value: val, isSet: true}
}

func (v NullableBTGeometryMateDisplayData1050) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGeometryMateDisplayData1050) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
