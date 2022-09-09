/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6392-2b80dda1e33c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMVariableStudioReference2764 struct for BTMVariableStudioReference2764
type BTMVariableStudioReference2764 struct {
	FeatureId               *string                                     `json:"featureId,omitempty"`
	FeatureType             *string                                     `json:"featureType,omitempty"`
	ImportMicroversion      *string                                     `json:"importMicroversion,omitempty"`
	Name                    *string                                     `json:"name,omitempty"`
	Namespace               *string                                     `json:"namespace,omitempty"`
	NodeId                  *string                                     `json:"nodeId,omitempty"`
	Parameters              []BTMParameter1                             `json:"parameters,omitempty"`
	ReturnAfterSubfeatures  *bool                                       `json:"returnAfterSubfeatures,omitempty"`
	SubFeatures             []BTMFeature134                             `json:"subFeatures,omitempty"`
	Suppressed              *bool                                       `json:"suppressed,omitempty"`
	SuppressionConfigured   *bool                                       `json:"suppressionConfigured,omitempty"`
	VariableStudioReference *bool                                       `json:"variableStudioReference,omitempty"`
	ApiConfiguration        *BTApiConfiguration                         `json:"apiConfiguration,omitempty"`
	BtType                  *string                                     `json:"btType,omitempty"`
	Configuration           []BTMParameter1                             `json:"configuration,omitempty"`
	DocumentId              *string                                     `json:"documentId,omitempty"`
	ElementId               *string                                     `json:"elementId,omitempty"`
	EntireVariableStudio    *bool                                       `json:"entireVariableStudio,omitempty"`
	IsAutomatic             *bool                                       `json:"isAutomatic,omitempty"`
	LockedState             *BTMParameter1                              `json:"lockedState,omitempty"`
	MicroversionId          *BTMicroversionId366                        `json:"microversionId,omitempty"`
	PartialReference        *bool                                       `json:"partialReference,omitempty"`
	ReferenceId             *string                                     `json:"referenceId,omitempty"`
	ReferenceNamespace      *string                                     `json:"referenceNamespace,omitempty"`
	ReferenceParameter      *BTMParameterReferenceWithConfiguration3028 `json:"referenceParameter,omitempty"`
	UnsetAutomaticEdit      *BTTreeEdit13                               `json:"unsetAutomaticEdit,omitempty"`
	ValidRevisionReference  *bool                                       `json:"validRevisionReference,omitempty"`
	VariableNames           []string                                    `json:"variableNames,omitempty"`
	VersionId               *string                                     `json:"versionId,omitempty"`
}

// NewBTMVariableStudioReference2764 instantiates a new BTMVariableStudioReference2764 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMVariableStudioReference2764() *BTMVariableStudioReference2764 {
	this := BTMVariableStudioReference2764{}
	return &this
}

// NewBTMVariableStudioReference2764WithDefaults instantiates a new BTMVariableStudioReference2764 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMVariableStudioReference2764WithDefaults() *BTMVariableStudioReference2764 {
	this := BTMVariableStudioReference2764{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMVariableStudioReference2764) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTMVariableStudioReference2764) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMVariableStudioReference2764) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMVariableStudioReference2764) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMVariableStudioReference2764) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMVariableStudioReference2764) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMVariableStudioReference2764) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetReturnAfterSubfeatures() bool {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		var ret bool
		return ret
	}
	return *o.ReturnAfterSubfeatures
}

// GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetReturnAfterSubfeaturesOk() (*bool, bool) {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		return nil, false
	}
	return o.ReturnAfterSubfeatures, true
}

// HasReturnAfterSubfeatures returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasReturnAfterSubfeatures() bool {
	if o != nil && o.ReturnAfterSubfeatures != nil {
		return true
	}

	return false
}

// SetReturnAfterSubfeatures gets a reference to the given bool and assigns it to the ReturnAfterSubfeatures field.
func (o *BTMVariableStudioReference2764) SetReturnAfterSubfeatures(v bool) {
	o.ReturnAfterSubfeatures = &v
}

// GetSubFeatures returns the SubFeatures field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetSubFeatures() []BTMFeature134 {
	if o == nil || o.SubFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.SubFeatures
}

// GetSubFeaturesOk returns a tuple with the SubFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetSubFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.SubFeatures == nil {
		return nil, false
	}
	return o.SubFeatures, true
}

// HasSubFeatures returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasSubFeatures() bool {
	if o != nil && o.SubFeatures != nil {
		return true
	}

	return false
}

// SetSubFeatures gets a reference to the given []BTMFeature134 and assigns it to the SubFeatures field.
func (o *BTMVariableStudioReference2764) SetSubFeatures(v []BTMFeature134) {
	o.SubFeatures = v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTMVariableStudioReference2764) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTMVariableStudioReference2764) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetVariableStudioReference returns the VariableStudioReference field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetVariableStudioReference() bool {
	if o == nil || o.VariableStudioReference == nil {
		var ret bool
		return ret
	}
	return *o.VariableStudioReference
}

// GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetVariableStudioReferenceOk() (*bool, bool) {
	if o == nil || o.VariableStudioReference == nil {
		return nil, false
	}
	return o.VariableStudioReference, true
}

// HasVariableStudioReference returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasVariableStudioReference() bool {
	if o != nil && o.VariableStudioReference != nil {
		return true
	}

	return false
}

// SetVariableStudioReference gets a reference to the given bool and assigns it to the VariableStudioReference field.
func (o *BTMVariableStudioReference2764) SetVariableStudioReference(v bool) {
	o.VariableStudioReference = &v
}

// GetApiConfiguration returns the ApiConfiguration field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetApiConfiguration() BTApiConfiguration {
	if o == nil || o.ApiConfiguration == nil {
		var ret BTApiConfiguration
		return ret
	}
	return *o.ApiConfiguration
}

// GetApiConfigurationOk returns a tuple with the ApiConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetApiConfigurationOk() (*BTApiConfiguration, bool) {
	if o == nil || o.ApiConfiguration == nil {
		return nil, false
	}
	return o.ApiConfiguration, true
}

// HasApiConfiguration returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasApiConfiguration() bool {
	if o != nil && o.ApiConfiguration != nil {
		return true
	}

	return false
}

// SetApiConfiguration gets a reference to the given BTApiConfiguration and assigns it to the ApiConfiguration field.
func (o *BTMVariableStudioReference2764) SetApiConfiguration(v BTApiConfiguration) {
	o.ApiConfiguration = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMVariableStudioReference2764) SetBtType(v string) {
	o.BtType = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetConfiguration() []BTMParameter1 {
	if o == nil || o.Configuration == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetConfigurationOk() ([]BTMParameter1, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given []BTMParameter1 and assigns it to the Configuration field.
func (o *BTMVariableStudioReference2764) SetConfiguration(v []BTMParameter1) {
	o.Configuration = v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTMVariableStudioReference2764) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTMVariableStudioReference2764) SetElementId(v string) {
	o.ElementId = &v
}

// GetEntireVariableStudio returns the EntireVariableStudio field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetEntireVariableStudio() bool {
	if o == nil || o.EntireVariableStudio == nil {
		var ret bool
		return ret
	}
	return *o.EntireVariableStudio
}

// GetEntireVariableStudioOk returns a tuple with the EntireVariableStudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetEntireVariableStudioOk() (*bool, bool) {
	if o == nil || o.EntireVariableStudio == nil {
		return nil, false
	}
	return o.EntireVariableStudio, true
}

// HasEntireVariableStudio returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasEntireVariableStudio() bool {
	if o != nil && o.EntireVariableStudio != nil {
		return true
	}

	return false
}

// SetEntireVariableStudio gets a reference to the given bool and assigns it to the EntireVariableStudio field.
func (o *BTMVariableStudioReference2764) SetEntireVariableStudio(v bool) {
	o.EntireVariableStudio = &v
}

// GetIsAutomatic returns the IsAutomatic field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetIsAutomatic() bool {
	if o == nil || o.IsAutomatic == nil {
		var ret bool
		return ret
	}
	return *o.IsAutomatic
}

// GetIsAutomaticOk returns a tuple with the IsAutomatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetIsAutomaticOk() (*bool, bool) {
	if o == nil || o.IsAutomatic == nil {
		return nil, false
	}
	return o.IsAutomatic, true
}

// HasIsAutomatic returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasIsAutomatic() bool {
	if o != nil && o.IsAutomatic != nil {
		return true
	}

	return false
}

// SetIsAutomatic gets a reference to the given bool and assigns it to the IsAutomatic field.
func (o *BTMVariableStudioReference2764) SetIsAutomatic(v bool) {
	o.IsAutomatic = &v
}

// GetLockedState returns the LockedState field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetLockedState() BTMParameter1 {
	if o == nil || o.LockedState == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.LockedState
}

// GetLockedStateOk returns a tuple with the LockedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetLockedStateOk() (*BTMParameter1, bool) {
	if o == nil || o.LockedState == nil {
		return nil, false
	}
	return o.LockedState, true
}

// HasLockedState returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasLockedState() bool {
	if o != nil && o.LockedState != nil {
		return true
	}

	return false
}

// SetLockedState gets a reference to the given BTMParameter1 and assigns it to the LockedState field.
func (o *BTMVariableStudioReference2764) SetLockedState(v BTMParameter1) {
	o.LockedState = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTMVariableStudioReference2764) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetPartialReference returns the PartialReference field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetPartialReference() bool {
	if o == nil || o.PartialReference == nil {
		var ret bool
		return ret
	}
	return *o.PartialReference
}

// GetPartialReferenceOk returns a tuple with the PartialReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetPartialReferenceOk() (*bool, bool) {
	if o == nil || o.PartialReference == nil {
		return nil, false
	}
	return o.PartialReference, true
}

// HasPartialReference returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasPartialReference() bool {
	if o != nil && o.PartialReference != nil {
		return true
	}

	return false
}

// SetPartialReference gets a reference to the given bool and assigns it to the PartialReference field.
func (o *BTMVariableStudioReference2764) SetPartialReference(v bool) {
	o.PartialReference = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetReferenceId() string {
	if o == nil || o.ReferenceId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetReferenceIdOk() (*string, bool) {
	if o == nil || o.ReferenceId == nil {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *BTMVariableStudioReference2764) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetReferenceNamespace returns the ReferenceNamespace field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetReferenceNamespace() string {
	if o == nil || o.ReferenceNamespace == nil {
		var ret string
		return ret
	}
	return *o.ReferenceNamespace
}

// GetReferenceNamespaceOk returns a tuple with the ReferenceNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetReferenceNamespaceOk() (*string, bool) {
	if o == nil || o.ReferenceNamespace == nil {
		return nil, false
	}
	return o.ReferenceNamespace, true
}

// HasReferenceNamespace returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasReferenceNamespace() bool {
	if o != nil && o.ReferenceNamespace != nil {
		return true
	}

	return false
}

// SetReferenceNamespace gets a reference to the given string and assigns it to the ReferenceNamespace field.
func (o *BTMVariableStudioReference2764) SetReferenceNamespace(v string) {
	o.ReferenceNamespace = &v
}

// GetReferenceParameter returns the ReferenceParameter field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetReferenceParameter() BTMParameterReferenceWithConfiguration3028 {
	if o == nil || o.ReferenceParameter == nil {
		var ret BTMParameterReferenceWithConfiguration3028
		return ret
	}
	return *o.ReferenceParameter
}

// GetReferenceParameterOk returns a tuple with the ReferenceParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetReferenceParameterOk() (*BTMParameterReferenceWithConfiguration3028, bool) {
	if o == nil || o.ReferenceParameter == nil {
		return nil, false
	}
	return o.ReferenceParameter, true
}

// HasReferenceParameter returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasReferenceParameter() bool {
	if o != nil && o.ReferenceParameter != nil {
		return true
	}

	return false
}

// SetReferenceParameter gets a reference to the given BTMParameterReferenceWithConfiguration3028 and assigns it to the ReferenceParameter field.
func (o *BTMVariableStudioReference2764) SetReferenceParameter(v BTMParameterReferenceWithConfiguration3028) {
	o.ReferenceParameter = &v
}

// GetUnsetAutomaticEdit returns the UnsetAutomaticEdit field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetUnsetAutomaticEdit() BTTreeEdit13 {
	if o == nil || o.UnsetAutomaticEdit == nil {
		var ret BTTreeEdit13
		return ret
	}
	return *o.UnsetAutomaticEdit
}

// GetUnsetAutomaticEditOk returns a tuple with the UnsetAutomaticEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetUnsetAutomaticEditOk() (*BTTreeEdit13, bool) {
	if o == nil || o.UnsetAutomaticEdit == nil {
		return nil, false
	}
	return o.UnsetAutomaticEdit, true
}

// HasUnsetAutomaticEdit returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasUnsetAutomaticEdit() bool {
	if o != nil && o.UnsetAutomaticEdit != nil {
		return true
	}

	return false
}

// SetUnsetAutomaticEdit gets a reference to the given BTTreeEdit13 and assigns it to the UnsetAutomaticEdit field.
func (o *BTMVariableStudioReference2764) SetUnsetAutomaticEdit(v BTTreeEdit13) {
	o.UnsetAutomaticEdit = &v
}

// GetValidRevisionReference returns the ValidRevisionReference field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetValidRevisionReference() bool {
	if o == nil || o.ValidRevisionReference == nil {
		var ret bool
		return ret
	}
	return *o.ValidRevisionReference
}

// GetValidRevisionReferenceOk returns a tuple with the ValidRevisionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetValidRevisionReferenceOk() (*bool, bool) {
	if o == nil || o.ValidRevisionReference == nil {
		return nil, false
	}
	return o.ValidRevisionReference, true
}

// HasValidRevisionReference returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasValidRevisionReference() bool {
	if o != nil && o.ValidRevisionReference != nil {
		return true
	}

	return false
}

// SetValidRevisionReference gets a reference to the given bool and assigns it to the ValidRevisionReference field.
func (o *BTMVariableStudioReference2764) SetValidRevisionReference(v bool) {
	o.ValidRevisionReference = &v
}

// GetVariableNames returns the VariableNames field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetVariableNames() []string {
	if o == nil || o.VariableNames == nil {
		var ret []string
		return ret
	}
	return o.VariableNames
}

// GetVariableNamesOk returns a tuple with the VariableNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetVariableNamesOk() ([]string, bool) {
	if o == nil || o.VariableNames == nil {
		return nil, false
	}
	return o.VariableNames, true
}

// HasVariableNames returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasVariableNames() bool {
	if o != nil && o.VariableNames != nil {
		return true
	}

	return false
}

// SetVariableNames gets a reference to the given []string and assigns it to the VariableNames field.
func (o *BTMVariableStudioReference2764) SetVariableNames(v []string) {
	o.VariableNames = v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTMVariableStudioReference2764) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMVariableStudioReference2764) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTMVariableStudioReference2764) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTMVariableStudioReference2764) SetVersionId(v string) {
	o.VersionId = &v
}

func (o BTMVariableStudioReference2764) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.ReturnAfterSubfeatures != nil {
		toSerialize["returnAfterSubfeatures"] = o.ReturnAfterSubfeatures
	}
	if o.SubFeatures != nil {
		toSerialize["subFeatures"] = o.SubFeatures
	}
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
	if o.SuppressionConfigured != nil {
		toSerialize["suppressionConfigured"] = o.SuppressionConfigured
	}
	if o.VariableStudioReference != nil {
		toSerialize["variableStudioReference"] = o.VariableStudioReference
	}
	if o.ApiConfiguration != nil {
		toSerialize["apiConfiguration"] = o.ApiConfiguration
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.EntireVariableStudio != nil {
		toSerialize["entireVariableStudio"] = o.EntireVariableStudio
	}
	if o.IsAutomatic != nil {
		toSerialize["isAutomatic"] = o.IsAutomatic
	}
	if o.LockedState != nil {
		toSerialize["lockedState"] = o.LockedState
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.PartialReference != nil {
		toSerialize["partialReference"] = o.PartialReference
	}
	if o.ReferenceId != nil {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if o.ReferenceNamespace != nil {
		toSerialize["referenceNamespace"] = o.ReferenceNamespace
	}
	if o.ReferenceParameter != nil {
		toSerialize["referenceParameter"] = o.ReferenceParameter
	}
	if o.UnsetAutomaticEdit != nil {
		toSerialize["unsetAutomaticEdit"] = o.UnsetAutomaticEdit
	}
	if o.ValidRevisionReference != nil {
		toSerialize["validRevisionReference"] = o.ValidRevisionReference
	}
	if o.VariableNames != nil {
		toSerialize["variableNames"] = o.VariableNames
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMVariableStudioReference2764 struct {
	value *BTMVariableStudioReference2764
	isSet bool
}

func (v NullableBTMVariableStudioReference2764) Get() *BTMVariableStudioReference2764 {
	return v.value
}

func (v *NullableBTMVariableStudioReference2764) Set(val *BTMVariableStudioReference2764) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMVariableStudioReference2764) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMVariableStudioReference2764) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMVariableStudioReference2764(val *BTMVariableStudioReference2764) *NullableBTMVariableStudioReference2764 {
	return &NullableBTMVariableStudioReference2764{value: val, isSet: true}
}

func (v NullableBTMVariableStudioReference2764) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMVariableStudioReference2764) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
