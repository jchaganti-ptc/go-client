/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23149-3610d8cd750e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPLValueBoxDereference251 struct for BTPLValueBoxDereference251
type BTPLValueBoxDereference251 struct {
	Atomic              *bool               `json:"atomic,omitempty"`
	BtType              *string             `json:"btType,omitempty"`
	DocumentationType   *GBTPDefinitionType `json:"documentationType,omitempty"`
	EndSourceLocation   *int32              `json:"endSourceLocation,omitempty"`
	NodeId              *string             `json:"nodeId,omitempty"`
	ShortDescriptor     *string             `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10         `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10         `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool               `json:"spaceDefault,omitempty"`
	StartSourceLocation *int32              `json:"startSourceLocation,omitempty"`
	Box                 *BTPExpression9     `json:"box,omitempty"`
	SpaceInside         *BTPSpace10         `json:"spaceInside,omitempty"`
}

// NewBTPLValueBoxDereference251 instantiates a new BTPLValueBoxDereference251 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPLValueBoxDereference251() *BTPLValueBoxDereference251 {
	this := BTPLValueBoxDereference251{}
	return &this
}

// NewBTPLValueBoxDereference251WithDefaults instantiates a new BTPLValueBoxDereference251 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPLValueBoxDereference251WithDefaults() *BTPLValueBoxDereference251 {
	this := BTPLValueBoxDereference251{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPLValueBoxDereference251) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPLValueBoxDereference251) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPLValueBoxDereference251) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPLValueBoxDereference251) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPLValueBoxDereference251) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPLValueBoxDereference251) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPLValueBoxDereference251) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPLValueBoxDereference251) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPLValueBoxDereference251) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPLValueBoxDereference251) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetBox returns the Box field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetBox() BTPExpression9 {
	if o == nil || o.Box == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Box
}

// GetBoxOk returns a tuple with the Box field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetBoxOk() (*BTPExpression9, bool) {
	if o == nil || o.Box == nil {
		return nil, false
	}
	return o.Box, true
}

// HasBox returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasBox() bool {
	if o != nil && o.Box != nil {
		return true
	}

	return false
}

// SetBox gets a reference to the given BTPExpression9 and assigns it to the Box field.
func (o *BTPLValueBoxDereference251) SetBox(v BTPExpression9) {
	o.Box = &v
}

// GetSpaceInside returns the SpaceInside field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetSpaceInside() BTPSpace10 {
	if o == nil || o.SpaceInside == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInside
}

// GetSpaceInsideOk returns a tuple with the SpaceInside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetSpaceInsideOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInside == nil {
		return nil, false
	}
	return o.SpaceInside, true
}

// HasSpaceInside returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasSpaceInside() bool {
	if o != nil && o.SpaceInside != nil {
		return true
	}

	return false
}

// SetSpaceInside gets a reference to the given BTPSpace10 and assigns it to the SpaceInside field.
func (o *BTPLValueBoxDereference251) SetSpaceInside(v BTPSpace10) {
	o.SpaceInside = &v
}

func (o BTPLValueBoxDereference251) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Atomic != nil {
		toSerialize["atomic"] = o.Atomic
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentationType != nil {
		toSerialize["documentationType"] = o.DocumentationType
	}
	if o.EndSourceLocation != nil {
		toSerialize["endSourceLocation"] = o.EndSourceLocation
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ShortDescriptor != nil {
		toSerialize["shortDescriptor"] = o.ShortDescriptor
	}
	if o.SpaceAfter != nil {
		toSerialize["spaceAfter"] = o.SpaceAfter
	}
	if o.SpaceBefore != nil {
		toSerialize["spaceBefore"] = o.SpaceBefore
	}
	if o.SpaceDefault != nil {
		toSerialize["spaceDefault"] = o.SpaceDefault
	}
	if o.StartSourceLocation != nil {
		toSerialize["startSourceLocation"] = o.StartSourceLocation
	}
	if o.Box != nil {
		toSerialize["box"] = o.Box
	}
	if o.SpaceInside != nil {
		toSerialize["spaceInside"] = o.SpaceInside
	}
	return json.Marshal(toSerialize)
}

type NullableBTPLValueBoxDereference251 struct {
	value *BTPLValueBoxDereference251
	isSet bool
}

func (v NullableBTPLValueBoxDereference251) Get() *BTPLValueBoxDereference251 {
	return v.value
}

func (v *NullableBTPLValueBoxDereference251) Set(val *BTPLValueBoxDereference251) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPLValueBoxDereference251) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPLValueBoxDereference251) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPLValueBoxDereference251(val *BTPLValueBoxDereference251) *NullableBTPLValueBoxDereference251 {
	return &NullableBTPLValueBoxDereference251{value: val, isSet: true}
}

func (v NullableBTPLValueBoxDereference251) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPLValueBoxDereference251) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
