/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6633-e6f3d20f07a4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPExpressionFunction1325 struct for BTPExpressionFunction1325
type BTPExpressionFunction1325 struct {
	Atomic              *bool                       `json:"atomic,omitempty"`
	BtType              *string                     `json:"btType,omitempty"`
	DocumentationType   *string                     `json:"documentationType,omitempty"`
	EndSourceLocation   *int32                      `json:"endSourceLocation,omitempty"`
	NodeId              *string                     `json:"nodeId,omitempty"`
	ShortDescriptor     *string                     `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10                 `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10                 `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool                       `json:"spaceDefault,omitempty"`
	StartSourceLocation *int32                      `json:"startSourceLocation,omitempty"`
	Arguments           []BTPArgumentDeclaration232 `json:"arguments,omitempty"`
	Body                *BTPStatementBlock271       `json:"body,omitempty"`
	Precondition        *BTPStatement269            `json:"precondition,omitempty"`
	ReturnType          *BTPTypeName290             `json:"returnType,omitempty"`
	SpaceAfterArglist   *BTPSpace10                 `json:"spaceAfterArglist,omitempty"`
	SpaceAfterFunction  *BTPSpace10                 `json:"spaceAfterFunction,omitempty"`
	SpaceInEmptyList    *BTPSpace10                 `json:"spaceInEmptyList,omitempty"`
}

// NewBTPExpressionFunction1325 instantiates a new BTPExpressionFunction1325 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPExpressionFunction1325() *BTPExpressionFunction1325 {
	this := BTPExpressionFunction1325{}
	return &this
}

// NewBTPExpressionFunction1325WithDefaults instantiates a new BTPExpressionFunction1325 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPExpressionFunction1325WithDefaults() *BTPExpressionFunction1325 {
	this := BTPExpressionFunction1325{}
	return &this
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPExpressionFunction1325) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPExpressionFunction1325) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetDocumentationType() string {
	if o == nil || o.DocumentationType == nil {
		var ret string
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetDocumentationTypeOk() (*string, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given string and assigns it to the DocumentationType field.
func (o *BTPExpressionFunction1325) SetDocumentationType(v string) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPExpressionFunction1325) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPExpressionFunction1325) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPExpressionFunction1325) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPExpressionFunction1325) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPExpressionFunction1325) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPExpressionFunction1325) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPExpressionFunction1325) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetArguments() []BTPArgumentDeclaration232 {
	if o == nil || o.Arguments == nil {
		var ret []BTPArgumentDeclaration232
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetArgumentsOk() ([]BTPArgumentDeclaration232, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []BTPArgumentDeclaration232 and assigns it to the Arguments field.
func (o *BTPExpressionFunction1325) SetArguments(v []BTPArgumentDeclaration232) {
	o.Arguments = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetBody() BTPStatementBlock271 {
	if o == nil || o.Body == nil {
		var ret BTPStatementBlock271
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetBodyOk() (*BTPStatementBlock271, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatementBlock271 and assigns it to the Body field.
func (o *BTPExpressionFunction1325) SetBody(v BTPStatementBlock271) {
	o.Body = &v
}

// GetPrecondition returns the Precondition field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetPrecondition() BTPStatement269 {
	if o == nil || o.Precondition == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Precondition
}

// GetPreconditionOk returns a tuple with the Precondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetPreconditionOk() (*BTPStatement269, bool) {
	if o == nil || o.Precondition == nil {
		return nil, false
	}
	return o.Precondition, true
}

// HasPrecondition returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasPrecondition() bool {
	if o != nil && o.Precondition != nil {
		return true
	}

	return false
}

// SetPrecondition gets a reference to the given BTPStatement269 and assigns it to the Precondition field.
func (o *BTPExpressionFunction1325) SetPrecondition(v BTPStatement269) {
	o.Precondition = &v
}

// GetReturnType returns the ReturnType field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetReturnType() BTPTypeName290 {
	if o == nil || o.ReturnType == nil {
		var ret BTPTypeName290
		return ret
	}
	return *o.ReturnType
}

// GetReturnTypeOk returns a tuple with the ReturnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetReturnTypeOk() (*BTPTypeName290, bool) {
	if o == nil || o.ReturnType == nil {
		return nil, false
	}
	return o.ReturnType, true
}

// HasReturnType returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasReturnType() bool {
	if o != nil && o.ReturnType != nil {
		return true
	}

	return false
}

// SetReturnType gets a reference to the given BTPTypeName290 and assigns it to the ReturnType field.
func (o *BTPExpressionFunction1325) SetReturnType(v BTPTypeName290) {
	o.ReturnType = &v
}

// GetSpaceAfterArglist returns the SpaceAfterArglist field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetSpaceAfterArglist() BTPSpace10 {
	if o == nil || o.SpaceAfterArglist == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterArglist
}

// GetSpaceAfterArglistOk returns a tuple with the SpaceAfterArglist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetSpaceAfterArglistOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterArglist == nil {
		return nil, false
	}
	return o.SpaceAfterArglist, true
}

// HasSpaceAfterArglist returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasSpaceAfterArglist() bool {
	if o != nil && o.SpaceAfterArglist != nil {
		return true
	}

	return false
}

// SetSpaceAfterArglist gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterArglist field.
func (o *BTPExpressionFunction1325) SetSpaceAfterArglist(v BTPSpace10) {
	o.SpaceAfterArglist = &v
}

// GetSpaceAfterFunction returns the SpaceAfterFunction field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetSpaceAfterFunction() BTPSpace10 {
	if o == nil || o.SpaceAfterFunction == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterFunction
}

// GetSpaceAfterFunctionOk returns a tuple with the SpaceAfterFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetSpaceAfterFunctionOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterFunction == nil {
		return nil, false
	}
	return o.SpaceAfterFunction, true
}

// HasSpaceAfterFunction returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasSpaceAfterFunction() bool {
	if o != nil && o.SpaceAfterFunction != nil {
		return true
	}

	return false
}

// SetSpaceAfterFunction gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterFunction field.
func (o *BTPExpressionFunction1325) SetSpaceAfterFunction(v BTPSpace10) {
	o.SpaceAfterFunction = &v
}

// GetSpaceInEmptyList returns the SpaceInEmptyList field value if set, zero value otherwise.
func (o *BTPExpressionFunction1325) GetSpaceInEmptyList() BTPSpace10 {
	if o == nil || o.SpaceInEmptyList == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInEmptyList
}

// GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPExpressionFunction1325) GetSpaceInEmptyListOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInEmptyList == nil {
		return nil, false
	}
	return o.SpaceInEmptyList, true
}

// HasSpaceInEmptyList returns a boolean if a field has been set.
func (o *BTPExpressionFunction1325) HasSpaceInEmptyList() bool {
	if o != nil && o.SpaceInEmptyList != nil {
		return true
	}

	return false
}

// SetSpaceInEmptyList gets a reference to the given BTPSpace10 and assigns it to the SpaceInEmptyList field.
func (o *BTPExpressionFunction1325) SetSpaceInEmptyList(v BTPSpace10) {
	o.SpaceInEmptyList = &v
}

func (o BTPExpressionFunction1325) MarshalJSON() ([]byte, error) {
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
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Precondition != nil {
		toSerialize["precondition"] = o.Precondition
	}
	if o.ReturnType != nil {
		toSerialize["returnType"] = o.ReturnType
	}
	if o.SpaceAfterArglist != nil {
		toSerialize["spaceAfterArglist"] = o.SpaceAfterArglist
	}
	if o.SpaceAfterFunction != nil {
		toSerialize["spaceAfterFunction"] = o.SpaceAfterFunction
	}
	if o.SpaceInEmptyList != nil {
		toSerialize["spaceInEmptyList"] = o.SpaceInEmptyList
	}
	return json.Marshal(toSerialize)
}

type NullableBTPExpressionFunction1325 struct {
	value *BTPExpressionFunction1325
	isSet bool
}

func (v NullableBTPExpressionFunction1325) Get() *BTPExpressionFunction1325 {
	return v.value
}

func (v *NullableBTPExpressionFunction1325) Set(val *BTPExpressionFunction1325) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPExpressionFunction1325) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPExpressionFunction1325) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPExpressionFunction1325(val *BTPExpressionFunction1325) *NullableBTPExpressionFunction1325 {
	return &NullableBTPExpressionFunction1325{value: val, isSet: true}
}

func (v NullableBTPExpressionFunction1325) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPExpressionFunction1325) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
