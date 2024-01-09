/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28658-06d4d4923fc7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPExpressionCall240 struct for BTPExpressionCall240
type BTPExpressionCall240 struct {
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
	FunctionExpression  *BTPExpression9     `json:"functionExpression,omitempty"`
	FunctionName        *BTPName261         `json:"functionName,omitempty"`
	FunctionNameString  *string             `json:"functionNameString,omitempty"`
	IsArrowCall         *bool               `json:"isArrowCall,omitempty"`
	SpaceInEmptyList    *BTPSpace10         `json:"spaceInEmptyList,omitempty"`
}

// NewBTPExpressionCall240 instantiates a new BTPExpressionCall240 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPExpressionCall240() *BTPExpressionCall240 {
	this := BTPExpressionCall240{}
	return &this
}

// NewBTPExpressionCall240WithDefaults instantiates a new BTPExpressionCall240 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPExpressionCall240WithDefaults() *BTPExpressionCall240 {
	this := BTPExpressionCall240{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPExpressionCall240) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionCall240) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetDocumentationType() GBTPDefinitionType {
	if o == nil || o.DocumentationType == nil {
		var ret GBTPDefinitionType
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetDocumentationTypeOk() (*GBTPDefinitionType, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given GBTPDefinitionType and assigns it to the DocumentationType field.
func (o *BTPExpressionCall240) SetDocumentationType(v GBTPDefinitionType) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPExpressionCall240) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPExpressionCall240) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPExpressionCall240) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPExpressionCall240) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPExpressionCall240) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPExpressionCall240) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPExpressionCall240) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetFunctionExpression returns the FunctionExpression field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetFunctionExpression() BTPExpression9 {
	if o == nil || o.FunctionExpression == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.FunctionExpression
}

// GetFunctionExpressionOk returns a tuple with the FunctionExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetFunctionExpressionOk() (*BTPExpression9, bool) {
	if o == nil || o.FunctionExpression == nil {
		return nil, false
	}
	return o.FunctionExpression, true
}

// HasFunctionExpression returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasFunctionExpression() bool {
	if o != nil && o.FunctionExpression != nil {
		return true
	}

	return false
}

// SetFunctionExpression gets a reference to the given BTPExpression9 and assigns it to the FunctionExpression field.
func (o *BTPExpressionCall240) SetFunctionExpression(v BTPExpression9) {
	o.FunctionExpression = &v
}

// GetFunctionName returns the FunctionName field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetFunctionName() BTPName261 {
	if o == nil || o.FunctionName == nil {
		var ret BTPName261
		return ret
	}
	return *o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetFunctionNameOk() (*BTPName261, bool) {
	if o == nil || o.FunctionName == nil {
		return nil, false
	}
	return o.FunctionName, true
}

// HasFunctionName returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasFunctionName() bool {
	if o != nil && o.FunctionName != nil {
		return true
	}

	return false
}

// SetFunctionName gets a reference to the given BTPName261 and assigns it to the FunctionName field.
func (o *BTPExpressionCall240) SetFunctionName(v BTPName261) {
	o.FunctionName = &v
}

// GetFunctionNameString returns the FunctionNameString field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetFunctionNameString() string {
	if o == nil || o.FunctionNameString == nil {
		var ret string
		return ret
	}
	return *o.FunctionNameString
}

// GetFunctionNameStringOk returns a tuple with the FunctionNameString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetFunctionNameStringOk() (*string, bool) {
	if o == nil || o.FunctionNameString == nil {
		return nil, false
	}
	return o.FunctionNameString, true
}

// HasFunctionNameString returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasFunctionNameString() bool {
	if o != nil && o.FunctionNameString != nil {
		return true
	}

	return false
}

// SetFunctionNameString gets a reference to the given string and assigns it to the FunctionNameString field.
func (o *BTPExpressionCall240) SetFunctionNameString(v string) {
	o.FunctionNameString = &v
}

// GetIsArrowCall returns the IsArrowCall field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetIsArrowCall() bool {
	if o == nil || o.IsArrowCall == nil {
		var ret bool
		return ret
	}
	return *o.IsArrowCall
}

// GetIsArrowCallOk returns a tuple with the IsArrowCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetIsArrowCallOk() (*bool, bool) {
	if o == nil || o.IsArrowCall == nil {
		return nil, false
	}
	return o.IsArrowCall, true
}

// HasIsArrowCall returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasIsArrowCall() bool {
	if o != nil && o.IsArrowCall != nil {
		return true
	}

	return false
}

// SetIsArrowCall gets a reference to the given bool and assigns it to the IsArrowCall field.
func (o *BTPExpressionCall240) SetIsArrowCall(v bool) {
	o.IsArrowCall = &v
}

// GetSpaceInEmptyList returns the SpaceInEmptyList field value if set, zero value otherwise.
func (o *BTPExpressionCall240) GetSpaceInEmptyList() BTPSpace10 {
	if o == nil || o.SpaceInEmptyList == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInEmptyList
}

// GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionCall240) GetSpaceInEmptyListOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInEmptyList == nil {
		return nil, false
	}
	return o.SpaceInEmptyList, true
}

// HasSpaceInEmptyList returns a boolean if a field has been set.
func (o *BTPExpressionCall240) HasSpaceInEmptyList() bool {
	if o != nil && o.SpaceInEmptyList != nil {
		return true
	}

	return false
}

// SetSpaceInEmptyList gets a reference to the given BTPSpace10 and assigns it to the SpaceInEmptyList field.
func (o *BTPExpressionCall240) SetSpaceInEmptyList(v BTPSpace10) {
	o.SpaceInEmptyList = &v
}

func (o BTPExpressionCall240) MarshalJSON() ([]byte, error) {
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
	if o.FunctionExpression != nil {
		toSerialize["functionExpression"] = o.FunctionExpression
	}
	if o.FunctionName != nil {
		toSerialize["functionName"] = o.FunctionName
	}
	if o.FunctionNameString != nil {
		toSerialize["functionNameString"] = o.FunctionNameString
	}
	if o.IsArrowCall != nil {
		toSerialize["isArrowCall"] = o.IsArrowCall
	}
	if o.SpaceInEmptyList != nil {
		toSerialize["spaceInEmptyList"] = o.SpaceInEmptyList
	}
	return json.Marshal(toSerialize)
}

type NullableBTPExpressionCall240 struct {
	value *BTPExpressionCall240
	isSet bool
}

func (v NullableBTPExpressionCall240) Get() *BTPExpressionCall240 {
	return v.value
}

func (v *NullableBTPExpressionCall240) Set(val *BTPExpressionCall240) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPExpressionCall240) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPExpressionCall240) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPExpressionCall240(val *BTPExpressionCall240) *NullableBTPExpressionCall240 {
	return &NullableBTPExpressionCall240{value: val, isSet: true}
}

func (v NullableBTPExpressionCall240) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPExpressionCall240) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
