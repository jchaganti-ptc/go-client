/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27313-c3b730633f3c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMAssemblyFeature12218 struct for BTMAssemblyFeature12218
type BTMAssemblyFeature12218 struct {
	BtType                                 *string                                   `json:"btType,omitempty"`
	FeatureId                              *string                                   `json:"featureId,omitempty"`
	FeatureType                            *string                                   `json:"featureType,omitempty"`
	ImportMicroversion                     *string                                   `json:"importMicroversion,omitempty"`
	Name                                   *string                                   `json:"name,omitempty"`
	Namespace                              *string                                   `json:"namespace,omitempty"`
	NodeId                                 *string                                   `json:"nodeId,omitempty"`
	Parameters                             []BTMParameter1                           `json:"parameters,omitempty"`
	ReturnAfterSubfeatures                 *bool                                     `json:"returnAfterSubfeatures,omitempty"`
	SubFeatures                            []BTMFeature134                           `json:"subFeatures,omitempty"`
	Suppressed                             *bool                                     `json:"suppressed,omitempty"`
	SuppressionConfigured                  *bool                                     `json:"suppressionConfigured,omitempty"`
	VariableStudioReference                *bool                                     `json:"variableStudioReference,omitempty"`
	AuxiliaryTreeFeature                   *bool                                     `json:"auxiliaryTreeFeature,omitempty"`
	FeatureFolder                          *bool                                     `json:"featureFolder,omitempty"`
	FeatureListFieldIndex                  *int32                                    `json:"featureListFieldIndex,omitempty"`
	FieldIndexForOwnedMateConnectors       *int32                                    `json:"fieldIndexForOwnedMateConnectors,omitempty"`
	OccurrenceQueriesFromAllConfigurations []BTMIndividualQueryWithOccurrenceBase904 `json:"occurrenceQueriesFromAllConfigurations,omitempty"`
	ParametricInstanceFeature              *bool                                     `json:"parametricInstanceFeature,omitempty"`
	Version                                *int32                                    `json:"version,omitempty"`
}

// NewBTMAssemblyFeature12218 instantiates a new BTMAssemblyFeature12218 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMAssemblyFeature12218() *BTMAssemblyFeature12218 {
	this := BTMAssemblyFeature12218{}
	return &this
}

// NewBTMAssemblyFeature12218WithDefaults instantiates a new BTMAssemblyFeature12218 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMAssemblyFeature12218WithDefaults() *BTMAssemblyFeature12218 {
	this := BTMAssemblyFeature12218{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMAssemblyFeature12218) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMAssemblyFeature12218) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTMAssemblyFeature12218) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMAssemblyFeature12218) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMAssemblyFeature12218) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMAssemblyFeature12218) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMAssemblyFeature12218) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMAssemblyFeature12218) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetReturnAfterSubfeatures returns the ReturnAfterSubfeatures field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetReturnAfterSubfeatures() bool {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		var ret bool
		return ret
	}
	return *o.ReturnAfterSubfeatures
}

// GetReturnAfterSubfeaturesOk returns a tuple with the ReturnAfterSubfeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetReturnAfterSubfeaturesOk() (*bool, bool) {
	if o == nil || o.ReturnAfterSubfeatures == nil {
		return nil, false
	}
	return o.ReturnAfterSubfeatures, true
}

// HasReturnAfterSubfeatures returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasReturnAfterSubfeatures() bool {
	if o != nil && o.ReturnAfterSubfeatures != nil {
		return true
	}

	return false
}

// SetReturnAfterSubfeatures gets a reference to the given bool and assigns it to the ReturnAfterSubfeatures field.
func (o *BTMAssemblyFeature12218) SetReturnAfterSubfeatures(v bool) {
	o.ReturnAfterSubfeatures = &v
}

// GetSubFeatures returns the SubFeatures field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetSubFeatures() []BTMFeature134 {
	if o == nil || o.SubFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.SubFeatures
}

// GetSubFeaturesOk returns a tuple with the SubFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetSubFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.SubFeatures == nil {
		return nil, false
	}
	return o.SubFeatures, true
}

// HasSubFeatures returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasSubFeatures() bool {
	if o != nil && o.SubFeatures != nil {
		return true
	}

	return false
}

// SetSubFeatures gets a reference to the given []BTMFeature134 and assigns it to the SubFeatures field.
func (o *BTMAssemblyFeature12218) SetSubFeatures(v []BTMFeature134) {
	o.SubFeatures = v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetSuppressed() bool {
	if o == nil || o.Suppressed == nil {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetSuppressedOk() (*bool, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *BTMAssemblyFeature12218) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetSuppressionConfigured returns the SuppressionConfigured field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetSuppressionConfigured() bool {
	if o == nil || o.SuppressionConfigured == nil {
		var ret bool
		return ret
	}
	return *o.SuppressionConfigured
}

// GetSuppressionConfiguredOk returns a tuple with the SuppressionConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetSuppressionConfiguredOk() (*bool, bool) {
	if o == nil || o.SuppressionConfigured == nil {
		return nil, false
	}
	return o.SuppressionConfigured, true
}

// HasSuppressionConfigured returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasSuppressionConfigured() bool {
	if o != nil && o.SuppressionConfigured != nil {
		return true
	}

	return false
}

// SetSuppressionConfigured gets a reference to the given bool and assigns it to the SuppressionConfigured field.
func (o *BTMAssemblyFeature12218) SetSuppressionConfigured(v bool) {
	o.SuppressionConfigured = &v
}

// GetVariableStudioReference returns the VariableStudioReference field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetVariableStudioReference() bool {
	if o == nil || o.VariableStudioReference == nil {
		var ret bool
		return ret
	}
	return *o.VariableStudioReference
}

// GetVariableStudioReferenceOk returns a tuple with the VariableStudioReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetVariableStudioReferenceOk() (*bool, bool) {
	if o == nil || o.VariableStudioReference == nil {
		return nil, false
	}
	return o.VariableStudioReference, true
}

// HasVariableStudioReference returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasVariableStudioReference() bool {
	if o != nil && o.VariableStudioReference != nil {
		return true
	}

	return false
}

// SetVariableStudioReference gets a reference to the given bool and assigns it to the VariableStudioReference field.
func (o *BTMAssemblyFeature12218) SetVariableStudioReference(v bool) {
	o.VariableStudioReference = &v
}

// GetAuxiliaryTreeFeature returns the AuxiliaryTreeFeature field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetAuxiliaryTreeFeature() bool {
	if o == nil || o.AuxiliaryTreeFeature == nil {
		var ret bool
		return ret
	}
	return *o.AuxiliaryTreeFeature
}

// GetAuxiliaryTreeFeatureOk returns a tuple with the AuxiliaryTreeFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetAuxiliaryTreeFeatureOk() (*bool, bool) {
	if o == nil || o.AuxiliaryTreeFeature == nil {
		return nil, false
	}
	return o.AuxiliaryTreeFeature, true
}

// HasAuxiliaryTreeFeature returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasAuxiliaryTreeFeature() bool {
	if o != nil && o.AuxiliaryTreeFeature != nil {
		return true
	}

	return false
}

// SetAuxiliaryTreeFeature gets a reference to the given bool and assigns it to the AuxiliaryTreeFeature field.
func (o *BTMAssemblyFeature12218) SetAuxiliaryTreeFeature(v bool) {
	o.AuxiliaryTreeFeature = &v
}

// GetFeatureFolder returns the FeatureFolder field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetFeatureFolder() bool {
	if o == nil || o.FeatureFolder == nil {
		var ret bool
		return ret
	}
	return *o.FeatureFolder
}

// GetFeatureFolderOk returns a tuple with the FeatureFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetFeatureFolderOk() (*bool, bool) {
	if o == nil || o.FeatureFolder == nil {
		return nil, false
	}
	return o.FeatureFolder, true
}

// HasFeatureFolder returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasFeatureFolder() bool {
	if o != nil && o.FeatureFolder != nil {
		return true
	}

	return false
}

// SetFeatureFolder gets a reference to the given bool and assigns it to the FeatureFolder field.
func (o *BTMAssemblyFeature12218) SetFeatureFolder(v bool) {
	o.FeatureFolder = &v
}

// GetFeatureListFieldIndex returns the FeatureListFieldIndex field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetFeatureListFieldIndex() int32 {
	if o == nil || o.FeatureListFieldIndex == nil {
		var ret int32
		return ret
	}
	return *o.FeatureListFieldIndex
}

// GetFeatureListFieldIndexOk returns a tuple with the FeatureListFieldIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetFeatureListFieldIndexOk() (*int32, bool) {
	if o == nil || o.FeatureListFieldIndex == nil {
		return nil, false
	}
	return o.FeatureListFieldIndex, true
}

// HasFeatureListFieldIndex returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasFeatureListFieldIndex() bool {
	if o != nil && o.FeatureListFieldIndex != nil {
		return true
	}

	return false
}

// SetFeatureListFieldIndex gets a reference to the given int32 and assigns it to the FeatureListFieldIndex field.
func (o *BTMAssemblyFeature12218) SetFeatureListFieldIndex(v int32) {
	o.FeatureListFieldIndex = &v
}

// GetFieldIndexForOwnedMateConnectors returns the FieldIndexForOwnedMateConnectors field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetFieldIndexForOwnedMateConnectors() int32 {
	if o == nil || o.FieldIndexForOwnedMateConnectors == nil {
		var ret int32
		return ret
	}
	return *o.FieldIndexForOwnedMateConnectors
}

// GetFieldIndexForOwnedMateConnectorsOk returns a tuple with the FieldIndexForOwnedMateConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetFieldIndexForOwnedMateConnectorsOk() (*int32, bool) {
	if o == nil || o.FieldIndexForOwnedMateConnectors == nil {
		return nil, false
	}
	return o.FieldIndexForOwnedMateConnectors, true
}

// HasFieldIndexForOwnedMateConnectors returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasFieldIndexForOwnedMateConnectors() bool {
	if o != nil && o.FieldIndexForOwnedMateConnectors != nil {
		return true
	}

	return false
}

// SetFieldIndexForOwnedMateConnectors gets a reference to the given int32 and assigns it to the FieldIndexForOwnedMateConnectors field.
func (o *BTMAssemblyFeature12218) SetFieldIndexForOwnedMateConnectors(v int32) {
	o.FieldIndexForOwnedMateConnectors = &v
}

// GetOccurrenceQueriesFromAllConfigurations returns the OccurrenceQueriesFromAllConfigurations field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetOccurrenceQueriesFromAllConfigurations() []BTMIndividualQueryWithOccurrenceBase904 {
	if o == nil || o.OccurrenceQueriesFromAllConfigurations == nil {
		var ret []BTMIndividualQueryWithOccurrenceBase904
		return ret
	}
	return o.OccurrenceQueriesFromAllConfigurations
}

// GetOccurrenceQueriesFromAllConfigurationsOk returns a tuple with the OccurrenceQueriesFromAllConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetOccurrenceQueriesFromAllConfigurationsOk() ([]BTMIndividualQueryWithOccurrenceBase904, bool) {
	if o == nil || o.OccurrenceQueriesFromAllConfigurations == nil {
		return nil, false
	}
	return o.OccurrenceQueriesFromAllConfigurations, true
}

// HasOccurrenceQueriesFromAllConfigurations returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasOccurrenceQueriesFromAllConfigurations() bool {
	if o != nil && o.OccurrenceQueriesFromAllConfigurations != nil {
		return true
	}

	return false
}

// SetOccurrenceQueriesFromAllConfigurations gets a reference to the given []BTMIndividualQueryWithOccurrenceBase904 and assigns it to the OccurrenceQueriesFromAllConfigurations field.
func (o *BTMAssemblyFeature12218) SetOccurrenceQueriesFromAllConfigurations(v []BTMIndividualQueryWithOccurrenceBase904) {
	o.OccurrenceQueriesFromAllConfigurations = v
}

// GetParametricInstanceFeature returns the ParametricInstanceFeature field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetParametricInstanceFeature() bool {
	if o == nil || o.ParametricInstanceFeature == nil {
		var ret bool
		return ret
	}
	return *o.ParametricInstanceFeature
}

// GetParametricInstanceFeatureOk returns a tuple with the ParametricInstanceFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetParametricInstanceFeatureOk() (*bool, bool) {
	if o == nil || o.ParametricInstanceFeature == nil {
		return nil, false
	}
	return o.ParametricInstanceFeature, true
}

// HasParametricInstanceFeature returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasParametricInstanceFeature() bool {
	if o != nil && o.ParametricInstanceFeature != nil {
		return true
	}

	return false
}

// SetParametricInstanceFeature gets a reference to the given bool and assigns it to the ParametricInstanceFeature field.
func (o *BTMAssemblyFeature12218) SetParametricInstanceFeature(v bool) {
	o.ParametricInstanceFeature = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BTMAssemblyFeature12218) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMAssemblyFeature12218) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BTMAssemblyFeature12218) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BTMAssemblyFeature12218) SetVersion(v int32) {
	o.Version = &v
}

func (o BTMAssemblyFeature12218) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
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
	if o.AuxiliaryTreeFeature != nil {
		toSerialize["auxiliaryTreeFeature"] = o.AuxiliaryTreeFeature
	}
	if o.FeatureFolder != nil {
		toSerialize["featureFolder"] = o.FeatureFolder
	}
	if o.FeatureListFieldIndex != nil {
		toSerialize["featureListFieldIndex"] = o.FeatureListFieldIndex
	}
	if o.FieldIndexForOwnedMateConnectors != nil {
		toSerialize["fieldIndexForOwnedMateConnectors"] = o.FieldIndexForOwnedMateConnectors
	}
	if o.OccurrenceQueriesFromAllConfigurations != nil {
		toSerialize["occurrenceQueriesFromAllConfigurations"] = o.OccurrenceQueriesFromAllConfigurations
	}
	if o.ParametricInstanceFeature != nil {
		toSerialize["parametricInstanceFeature"] = o.ParametricInstanceFeature
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBTMAssemblyFeature12218 struct {
	value *BTMAssemblyFeature12218
	isSet bool
}

func (v NullableBTMAssemblyFeature12218) Get() *BTMAssemblyFeature12218 {
	return v.value
}

func (v *NullableBTMAssemblyFeature12218) Set(val *BTMAssemblyFeature12218) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMAssemblyFeature12218) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMAssemblyFeature12218) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMAssemblyFeature12218(val *BTMAssemblyFeature12218) *NullableBTMAssemblyFeature12218 {
	return &NullableBTMAssemblyFeature12218{value: val, isSet: true}
}

func (v NullableBTMAssemblyFeature12218) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMAssemblyFeature12218) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
