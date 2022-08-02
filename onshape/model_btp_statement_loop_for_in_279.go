/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5843-26f51e563fa4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPStatementLoopForIn279 struct for BTPStatementLoopForIn279
type BTPStatementLoopForIn279 struct {
	Annotation          *BTPAnnotation231 `json:"annotation,omitempty"`
	Atomic              *bool             `json:"atomic,omitempty"`
	BtType              *string           `json:"btType,omitempty"`
	DocumentationType   *string           `json:"documentationType,omitempty"`
	EndSourceLocation   *int32            `json:"endSourceLocation,omitempty"`
	NodeId              *string           `json:"nodeId,omitempty"`
	ShortDescriptor     *string           `json:"shortDescriptor,omitempty"`
	SpaceAfter          *BTPSpace10       `json:"spaceAfter,omitempty"`
	SpaceBefore         *BTPSpace10       `json:"spaceBefore,omitempty"`
	SpaceDefault        *bool             `json:"spaceDefault,omitempty"`
	StartSourceLocation *int32            `json:"startSourceLocation,omitempty"`
	Body                *BTPStatement269  `json:"body,omitempty"`
	SpaceAfterLoopType  *BTPSpace10       `json:"spaceAfterLoopType,omitempty"`
	Container           *BTPExpression9   `json:"container,omitempty"`
	Identifiers         []BTPIdentifier8  `json:"identifiers,omitempty"`
	IsVarDeclaredHere   *bool             `json:"isVarDeclaredHere,omitempty"`
	KeyVar              *BTPIdentifier8   `json:"keyVar,omitempty"`
	SpaceBeforeVar      *BTPSpace10       `json:"spaceBeforeVar,omitempty"`
	StandardTypes       []string          `json:"standardTypes,omitempty"`
	TypeNames           []string          `json:"typeNames,omitempty"`
	Var                 *BTPIdentifier8   `json:"var,omitempty"`
}

// NewBTPStatementLoopForIn279 instantiates a new BTPStatementLoopForIn279 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementLoopForIn279() *BTPStatementLoopForIn279 {
	this := BTPStatementLoopForIn279{}
	return &this
}

// NewBTPStatementLoopForIn279WithDefaults instantiates a new BTPStatementLoopForIn279 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementLoopForIn279WithDefaults() *BTPStatementLoopForIn279 {
	this := BTPStatementLoopForIn279{}
	return &this
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetAnnotation() BTPAnnotation231 {
	if o == nil || o.Annotation == nil {
		var ret BTPAnnotation231
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetAnnotationOk() (*BTPAnnotation231, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given BTPAnnotation231 and assigns it to the Annotation field.
func (o *BTPStatementLoopForIn279) SetAnnotation(v BTPAnnotation231) {
	o.Annotation = &v
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetAtomic() bool {
	if o == nil || o.Atomic == nil {
		var ret bool
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetAtomicOk() (*bool, bool) {
	if o == nil || o.Atomic == nil {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasAtomic() bool {
	if o != nil && o.Atomic != nil {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given bool and assigns it to the Atomic field.
func (o *BTPStatementLoopForIn279) SetAtomic(v bool) {
	o.Atomic = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementLoopForIn279) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentationType returns the DocumentationType field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetDocumentationType() string {
	if o == nil || o.DocumentationType == nil {
		var ret string
		return ret
	}
	return *o.DocumentationType
}

// GetDocumentationTypeOk returns a tuple with the DocumentationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetDocumentationTypeOk() (*string, bool) {
	if o == nil || o.DocumentationType == nil {
		return nil, false
	}
	return o.DocumentationType, true
}

// HasDocumentationType returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasDocumentationType() bool {
	if o != nil && o.DocumentationType != nil {
		return true
	}

	return false
}

// SetDocumentationType gets a reference to the given string and assigns it to the DocumentationType field.
func (o *BTPStatementLoopForIn279) SetDocumentationType(v string) {
	o.DocumentationType = &v
}

// GetEndSourceLocation returns the EndSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetEndSourceLocation() int32 {
	if o == nil || o.EndSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.EndSourceLocation
}

// GetEndSourceLocationOk returns a tuple with the EndSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetEndSourceLocationOk() (*int32, bool) {
	if o == nil || o.EndSourceLocation == nil {
		return nil, false
	}
	return o.EndSourceLocation, true
}

// HasEndSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasEndSourceLocation() bool {
	if o != nil && o.EndSourceLocation != nil {
		return true
	}

	return false
}

// SetEndSourceLocation gets a reference to the given int32 and assigns it to the EndSourceLocation field.
func (o *BTPStatementLoopForIn279) SetEndSourceLocation(v int32) {
	o.EndSourceLocation = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTPStatementLoopForIn279) SetNodeId(v string) {
	o.NodeId = &v
}

// GetShortDescriptor returns the ShortDescriptor field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetShortDescriptor() string {
	if o == nil || o.ShortDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ShortDescriptor
}

// GetShortDescriptorOk returns a tuple with the ShortDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetShortDescriptorOk() (*string, bool) {
	if o == nil || o.ShortDescriptor == nil {
		return nil, false
	}
	return o.ShortDescriptor, true
}

// HasShortDescriptor returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasShortDescriptor() bool {
	if o != nil && o.ShortDescriptor != nil {
		return true
	}

	return false
}

// SetShortDescriptor gets a reference to the given string and assigns it to the ShortDescriptor field.
func (o *BTPStatementLoopForIn279) SetShortDescriptor(v string) {
	o.ShortDescriptor = &v
}

// GetSpaceAfter returns the SpaceAfter field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetSpaceAfter() BTPSpace10 {
	if o == nil || o.SpaceAfter == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfter
}

// GetSpaceAfterOk returns a tuple with the SpaceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetSpaceAfterOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfter == nil {
		return nil, false
	}
	return o.SpaceAfter, true
}

// HasSpaceAfter returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasSpaceAfter() bool {
	if o != nil && o.SpaceAfter != nil {
		return true
	}

	return false
}

// SetSpaceAfter gets a reference to the given BTPSpace10 and assigns it to the SpaceAfter field.
func (o *BTPStatementLoopForIn279) SetSpaceAfter(v BTPSpace10) {
	o.SpaceAfter = &v
}

// GetSpaceBefore returns the SpaceBefore field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetSpaceBefore() BTPSpace10 {
	if o == nil || o.SpaceBefore == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBefore
}

// GetSpaceBeforeOk returns a tuple with the SpaceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetSpaceBeforeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBefore == nil {
		return nil, false
	}
	return o.SpaceBefore, true
}

// HasSpaceBefore returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasSpaceBefore() bool {
	if o != nil && o.SpaceBefore != nil {
		return true
	}

	return false
}

// SetSpaceBefore gets a reference to the given BTPSpace10 and assigns it to the SpaceBefore field.
func (o *BTPStatementLoopForIn279) SetSpaceBefore(v BTPSpace10) {
	o.SpaceBefore = &v
}

// GetSpaceDefault returns the SpaceDefault field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetSpaceDefault() bool {
	if o == nil || o.SpaceDefault == nil {
		var ret bool
		return ret
	}
	return *o.SpaceDefault
}

// GetSpaceDefaultOk returns a tuple with the SpaceDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetSpaceDefaultOk() (*bool, bool) {
	if o == nil || o.SpaceDefault == nil {
		return nil, false
	}
	return o.SpaceDefault, true
}

// HasSpaceDefault returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasSpaceDefault() bool {
	if o != nil && o.SpaceDefault != nil {
		return true
	}

	return false
}

// SetSpaceDefault gets a reference to the given bool and assigns it to the SpaceDefault field.
func (o *BTPStatementLoopForIn279) SetSpaceDefault(v bool) {
	o.SpaceDefault = &v
}

// GetStartSourceLocation returns the StartSourceLocation field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetStartSourceLocation() int32 {
	if o == nil || o.StartSourceLocation == nil {
		var ret int32
		return ret
	}
	return *o.StartSourceLocation
}

// GetStartSourceLocationOk returns a tuple with the StartSourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetStartSourceLocationOk() (*int32, bool) {
	if o == nil || o.StartSourceLocation == nil {
		return nil, false
	}
	return o.StartSourceLocation, true
}

// HasStartSourceLocation returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasStartSourceLocation() bool {
	if o != nil && o.StartSourceLocation != nil {
		return true
	}

	return false
}

// SetStartSourceLocation gets a reference to the given int32 and assigns it to the StartSourceLocation field.
func (o *BTPStatementLoopForIn279) SetStartSourceLocation(v int32) {
	o.StartSourceLocation = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetBody() BTPStatement269 {
	if o == nil || o.Body == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetBodyOk() (*BTPStatement269, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatement269 and assigns it to the Body field.
func (o *BTPStatementLoopForIn279) SetBody(v BTPStatement269) {
	o.Body = &v
}

// GetSpaceAfterLoopType returns the SpaceAfterLoopType field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetSpaceAfterLoopType() BTPSpace10 {
	if o == nil || o.SpaceAfterLoopType == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterLoopType
}

// GetSpaceAfterLoopTypeOk returns a tuple with the SpaceAfterLoopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetSpaceAfterLoopTypeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterLoopType == nil {
		return nil, false
	}
	return o.SpaceAfterLoopType, true
}

// HasSpaceAfterLoopType returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasSpaceAfterLoopType() bool {
	if o != nil && o.SpaceAfterLoopType != nil {
		return true
	}

	return false
}

// SetSpaceAfterLoopType gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterLoopType field.
func (o *BTPStatementLoopForIn279) SetSpaceAfterLoopType(v BTPSpace10) {
	o.SpaceAfterLoopType = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetContainer() BTPExpression9 {
	if o == nil || o.Container == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetContainerOk() (*BTPExpression9, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given BTPExpression9 and assigns it to the Container field.
func (o *BTPStatementLoopForIn279) SetContainer(v BTPExpression9) {
	o.Container = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetIdentifiers() []BTPIdentifier8 {
	if o == nil || o.Identifiers == nil {
		var ret []BTPIdentifier8
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetIdentifiersOk() ([]BTPIdentifier8, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []BTPIdentifier8 and assigns it to the Identifiers field.
func (o *BTPStatementLoopForIn279) SetIdentifiers(v []BTPIdentifier8) {
	o.Identifiers = v
}

// GetIsVarDeclaredHere returns the IsVarDeclaredHere field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetIsVarDeclaredHere() bool {
	if o == nil || o.IsVarDeclaredHere == nil {
		var ret bool
		return ret
	}
	return *o.IsVarDeclaredHere
}

// GetIsVarDeclaredHereOk returns a tuple with the IsVarDeclaredHere field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetIsVarDeclaredHereOk() (*bool, bool) {
	if o == nil || o.IsVarDeclaredHere == nil {
		return nil, false
	}
	return o.IsVarDeclaredHere, true
}

// HasIsVarDeclaredHere returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasIsVarDeclaredHere() bool {
	if o != nil && o.IsVarDeclaredHere != nil {
		return true
	}

	return false
}

// SetIsVarDeclaredHere gets a reference to the given bool and assigns it to the IsVarDeclaredHere field.
func (o *BTPStatementLoopForIn279) SetIsVarDeclaredHere(v bool) {
	o.IsVarDeclaredHere = &v
}

// GetKeyVar returns the KeyVar field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetKeyVar() BTPIdentifier8 {
	if o == nil || o.KeyVar == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.KeyVar
}

// GetKeyVarOk returns a tuple with the KeyVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetKeyVarOk() (*BTPIdentifier8, bool) {
	if o == nil || o.KeyVar == nil {
		return nil, false
	}
	return o.KeyVar, true
}

// HasKeyVar returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasKeyVar() bool {
	if o != nil && o.KeyVar != nil {
		return true
	}

	return false
}

// SetKeyVar gets a reference to the given BTPIdentifier8 and assigns it to the KeyVar field.
func (o *BTPStatementLoopForIn279) SetKeyVar(v BTPIdentifier8) {
	o.KeyVar = &v
}

// GetSpaceBeforeVar returns the SpaceBeforeVar field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetSpaceBeforeVar() BTPSpace10 {
	if o == nil || o.SpaceBeforeVar == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeVar
}

// GetSpaceBeforeVarOk returns a tuple with the SpaceBeforeVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetSpaceBeforeVarOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeVar == nil {
		return nil, false
	}
	return o.SpaceBeforeVar, true
}

// HasSpaceBeforeVar returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasSpaceBeforeVar() bool {
	if o != nil && o.SpaceBeforeVar != nil {
		return true
	}

	return false
}

// SetSpaceBeforeVar gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeVar field.
func (o *BTPStatementLoopForIn279) SetSpaceBeforeVar(v BTPSpace10) {
	o.SpaceBeforeVar = &v
}

// GetStandardTypes returns the StandardTypes field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetStandardTypes() []string {
	if o == nil || o.StandardTypes == nil {
		var ret []string
		return ret
	}
	return o.StandardTypes
}

// GetStandardTypesOk returns a tuple with the StandardTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetStandardTypesOk() ([]string, bool) {
	if o == nil || o.StandardTypes == nil {
		return nil, false
	}
	return o.StandardTypes, true
}

// HasStandardTypes returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasStandardTypes() bool {
	if o != nil && o.StandardTypes != nil {
		return true
	}

	return false
}

// SetStandardTypes gets a reference to the given []string and assigns it to the StandardTypes field.
func (o *BTPStatementLoopForIn279) SetStandardTypes(v []string) {
	o.StandardTypes = v
}

// GetTypeNames returns the TypeNames field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetTypeNames() []string {
	if o == nil || o.TypeNames == nil {
		var ret []string
		return ret
	}
	return o.TypeNames
}

// GetTypeNamesOk returns a tuple with the TypeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetTypeNamesOk() ([]string, bool) {
	if o == nil || o.TypeNames == nil {
		return nil, false
	}
	return o.TypeNames, true
}

// HasTypeNames returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasTypeNames() bool {
	if o != nil && o.TypeNames != nil {
		return true
	}

	return false
}

// SetTypeNames gets a reference to the given []string and assigns it to the TypeNames field.
func (o *BTPStatementLoopForIn279) SetTypeNames(v []string) {
	o.TypeNames = v
}

// GetVar returns the Var field value if set, zero value otherwise.
func (o *BTPStatementLoopForIn279) GetVar() BTPIdentifier8 {
	if o == nil || o.Var == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.Var
}

// GetVarOk returns a tuple with the Var field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopForIn279) GetVarOk() (*BTPIdentifier8, bool) {
	if o == nil || o.Var == nil {
		return nil, false
	}
	return o.Var, true
}

// HasVar returns a boolean if a field has been set.
func (o *BTPStatementLoopForIn279) HasVar() bool {
	if o != nil && o.Var != nil {
		return true
	}

	return false
}

// SetVar gets a reference to the given BTPIdentifier8 and assigns it to the Var field.
func (o *BTPStatementLoopForIn279) SetVar(v BTPIdentifier8) {
	o.Var = &v
}

func (o BTPStatementLoopForIn279) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Annotation != nil {
		toSerialize["annotation"] = o.Annotation
	}
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
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.SpaceAfterLoopType != nil {
		toSerialize["spaceAfterLoopType"] = o.SpaceAfterLoopType
	}
	if o.Container != nil {
		toSerialize["container"] = o.Container
	}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if o.IsVarDeclaredHere != nil {
		toSerialize["isVarDeclaredHere"] = o.IsVarDeclaredHere
	}
	if o.KeyVar != nil {
		toSerialize["keyVar"] = o.KeyVar
	}
	if o.SpaceBeforeVar != nil {
		toSerialize["spaceBeforeVar"] = o.SpaceBeforeVar
	}
	if o.StandardTypes != nil {
		toSerialize["standardTypes"] = o.StandardTypes
	}
	if o.TypeNames != nil {
		toSerialize["typeNames"] = o.TypeNames
	}
	if o.Var != nil {
		toSerialize["var"] = o.Var
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementLoopForIn279 struct {
	value *BTPStatementLoopForIn279
	isSet bool
}

func (v NullableBTPStatementLoopForIn279) Get() *BTPStatementLoopForIn279 {
	return v.value
}

func (v *NullableBTPStatementLoopForIn279) Set(val *BTPStatementLoopForIn279) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementLoopForIn279) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementLoopForIn279) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementLoopForIn279(val *BTPStatementLoopForIn279) *NullableBTPStatementLoopForIn279 {
	return &NullableBTPStatementLoopForIn279{value: val, isSet: true}
}

func (v NullableBTPStatementLoopForIn279) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementLoopForIn279) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
