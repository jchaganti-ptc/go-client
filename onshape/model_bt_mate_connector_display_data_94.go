/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23307-4b97c8644a61
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMateConnectorDisplayData94 struct for BTMateConnectorDisplayData94
type BTMateConnectorDisplayData94 struct {
	BtType           *string                `json:"btType,omitempty"`
	ElementId        *string                `json:"elementId,omitempty"`
	EntityIds        []string               `json:"entityIds,omitempty"`
	Hidden           *bool                  `json:"hidden,omitempty"`
	Implicit         *bool                  `json:"implicit,omitempty"`
	IsDerivedFeature *bool                  `json:"isDerivedFeature,omitempty"`
	Location         *BTCoordinateSystem387 `json:"location,omitempty"`
	NodeId           *string                `json:"nodeId,omitempty"`
	Occurrence       *BTOccurrence74        `json:"occurrence,omitempty"`
	OwnerOccurrence  *BTOccurrence74        `json:"ownerOccurrence,omitempty"`
	PartId           *string                `json:"partId,omitempty"`
}

// NewBTMateConnectorDisplayData94 instantiates a new BTMateConnectorDisplayData94 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMateConnectorDisplayData94() *BTMateConnectorDisplayData94 {
	this := BTMateConnectorDisplayData94{}
	return &this
}

// NewBTMateConnectorDisplayData94WithDefaults instantiates a new BTMateConnectorDisplayData94 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMateConnectorDisplayData94WithDefaults() *BTMateConnectorDisplayData94 {
	this := BTMateConnectorDisplayData94{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMateConnectorDisplayData94) SetBtType(v string) {
	o.BtType = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTMateConnectorDisplayData94) SetElementId(v string) {
	o.ElementId = &v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetEntityIds() []string {
	if o == nil || o.EntityIds == nil {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetEntityIdsOk() ([]string, bool) {
	if o == nil || o.EntityIds == nil {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasEntityIds() bool {
	if o != nil && o.EntityIds != nil {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *BTMateConnectorDisplayData94) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTMateConnectorDisplayData94) SetHidden(v bool) {
	o.Hidden = &v
}

// GetImplicit returns the Implicit field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetImplicit() bool {
	if o == nil || o.Implicit == nil {
		var ret bool
		return ret
	}
	return *o.Implicit
}

// GetImplicitOk returns a tuple with the Implicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetImplicitOk() (*bool, bool) {
	if o == nil || o.Implicit == nil {
		return nil, false
	}
	return o.Implicit, true
}

// HasImplicit returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasImplicit() bool {
	if o != nil && o.Implicit != nil {
		return true
	}

	return false
}

// SetImplicit gets a reference to the given bool and assigns it to the Implicit field.
func (o *BTMateConnectorDisplayData94) SetImplicit(v bool) {
	o.Implicit = &v
}

// GetIsDerivedFeature returns the IsDerivedFeature field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetIsDerivedFeature() bool {
	if o == nil || o.IsDerivedFeature == nil {
		var ret bool
		return ret
	}
	return *o.IsDerivedFeature
}

// GetIsDerivedFeatureOk returns a tuple with the IsDerivedFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetIsDerivedFeatureOk() (*bool, bool) {
	if o == nil || o.IsDerivedFeature == nil {
		return nil, false
	}
	return o.IsDerivedFeature, true
}

// HasIsDerivedFeature returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasIsDerivedFeature() bool {
	if o != nil && o.IsDerivedFeature != nil {
		return true
	}

	return false
}

// SetIsDerivedFeature gets a reference to the given bool and assigns it to the IsDerivedFeature field.
func (o *BTMateConnectorDisplayData94) SetIsDerivedFeature(v bool) {
	o.IsDerivedFeature = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetLocation() BTCoordinateSystem387 {
	if o == nil || o.Location == nil {
		var ret BTCoordinateSystem387
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetLocationOk() (*BTCoordinateSystem387, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BTCoordinateSystem387 and assigns it to the Location field.
func (o *BTMateConnectorDisplayData94) SetLocation(v BTCoordinateSystem387) {
	o.Location = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMateConnectorDisplayData94) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTMateConnectorDisplayData94) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetOwnerOccurrence returns the OwnerOccurrence field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetOwnerOccurrence() BTOccurrence74 {
	if o == nil || o.OwnerOccurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.OwnerOccurrence
}

// GetOwnerOccurrenceOk returns a tuple with the OwnerOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetOwnerOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.OwnerOccurrence == nil {
		return nil, false
	}
	return o.OwnerOccurrence, true
}

// HasOwnerOccurrence returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasOwnerOccurrence() bool {
	if o != nil && o.OwnerOccurrence != nil {
		return true
	}

	return false
}

// SetOwnerOccurrence gets a reference to the given BTOccurrence74 and assigns it to the OwnerOccurrence field.
func (o *BTMateConnectorDisplayData94) SetOwnerOccurrence(v BTOccurrence74) {
	o.OwnerOccurrence = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTMateConnectorDisplayData94) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMateConnectorDisplayData94) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTMateConnectorDisplayData94) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTMateConnectorDisplayData94) SetPartId(v string) {
	o.PartId = &v
}

func (o BTMateConnectorDisplayData94) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.EntityIds != nil {
		toSerialize["entityIds"] = o.EntityIds
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.Implicit != nil {
		toSerialize["implicit"] = o.Implicit
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
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.OwnerOccurrence != nil {
		toSerialize["ownerOccurrence"] = o.OwnerOccurrence
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMateConnectorDisplayData94 struct {
	value *BTMateConnectorDisplayData94
	isSet bool
}

func (v NullableBTMateConnectorDisplayData94) Get() *BTMateConnectorDisplayData94 {
	return v.value
}

func (v *NullableBTMateConnectorDisplayData94) Set(val *BTMateConnectorDisplayData94) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMateConnectorDisplayData94) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMateConnectorDisplayData94) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMateConnectorDisplayData94(val *BTMateConnectorDisplayData94) *NullableBTMateConnectorDisplayData94 {
	return &NullableBTMateConnectorDisplayData94{value: val, isSet: true}
}

func (v NullableBTMateConnectorDisplayData94) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMateConnectorDisplayData94) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
