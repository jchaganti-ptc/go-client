/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6108-60a91d537e42
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMModel141 struct for BTMModel141
type BTMModel141 struct {
	AllFeatures                   []BTMFeature134                  `json:"allFeatures,omitempty"`
	AllFeaturesAndOtherReferences []BTMFeature134                  `json:"allFeaturesAndOtherReferences,omitempty"`
	AllFeaturesAndSubFeatures     []BTMFeature134                  `json:"allFeaturesAndSubFeatures,omitempty"`
	BelScriptLibraryMajorVersion  *int32                           `json:"belScriptLibraryMajorVersion,omitempty"`
	BelScriptLibraryVersion       *string                          `json:"belScriptLibraryVersion,omitempty"`
	BelScriptLibraryVersionEnum   *string                          `json:"belScriptLibraryVersionEnum,omitempty"`
	BtType                        *string                          `json:"btType,omitempty"`
	Children                      []BTMNode19                      `json:"children,omitempty"`
	ConfigurationData             *BTMConfigurationData1560        `json:"configurationData,omitempty"`
	Configured                    *bool                            `json:"configured,omitempty"`
	DeepImports                   *map[string][]BTImport           `json:"deepImports,omitempty"`
	DefaultFeatures               *BTDefaultFeatures119            `json:"defaultFeatures,omitempty"`
	DefaultUnits                  *BTMUnitsDefault160              `json:"defaultUnits,omitempty"`
	FeatureImports                *map[string][]BTImport           `json:"featureImports,omitempty"`
	FirstRollbackIndex            *int32                           `json:"firstRollbackIndex,omitempty"`
	ImportMicroversion            *string                          `json:"importMicroversion,omitempty"`
	ImportSet                     []BTPModuleId235                 `json:"importSet,omitempty"`
	Imports                       []BTMImport136                   `json:"imports,omitempty"`
	IsVariableStudio              *bool                            `json:"isVariableStudio,omitempty"`
	LastFeatureBeforeRollBack     *BTMFeature134                   `json:"lastFeatureBeforeRollBack,omitempty"`
	Name                          *string                          `json:"name,omitempty"`
	NodeId                        *string                          `json:"nodeId,omitempty"`
	PartProperties                *BTPartProperties293             `json:"partProperties,omitempty"`
	PathToCache                   *BTCacheDataPath191              `json:"pathToCache,omitempty"`
	Properties                    *BTModelProperties1258           `json:"properties,omitempty"`
	RollbackBar                   *BTMRollback150                  `json:"rollbackBar,omitempty"`
	RolledBackToEnd               *bool                            `json:"rolledBackToEnd,omitempty"`
	VariableStudios               []BTMVariableStudioReference2764 `json:"variableStudios,omitempty"`
}

// NewBTMModel141 instantiates a new BTMModel141 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMModel141() *BTMModel141 {
	this := BTMModel141{}
	return &this
}

// NewBTMModel141WithDefaults instantiates a new BTMModel141 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMModel141WithDefaults() *BTMModel141 {
	this := BTMModel141{}
	return &this
}

// GetAllFeatures returns the AllFeatures field value if set, zero value otherwise.
func (o *BTMModel141) GetAllFeatures() []BTMFeature134 {
	if o == nil || o.AllFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.AllFeatures
}

// GetAllFeaturesOk returns a tuple with the AllFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetAllFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.AllFeatures == nil {
		return nil, false
	}
	return o.AllFeatures, true
}

// HasAllFeatures returns a boolean if a field has been set.
func (o *BTMModel141) HasAllFeatures() bool {
	if o != nil && o.AllFeatures != nil {
		return true
	}

	return false
}

// SetAllFeatures gets a reference to the given []BTMFeature134 and assigns it to the AllFeatures field.
func (o *BTMModel141) SetAllFeatures(v []BTMFeature134) {
	o.AllFeatures = v
}

// GetAllFeaturesAndOtherReferences returns the AllFeaturesAndOtherReferences field value if set, zero value otherwise.
func (o *BTMModel141) GetAllFeaturesAndOtherReferences() []BTMFeature134 {
	if o == nil || o.AllFeaturesAndOtherReferences == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.AllFeaturesAndOtherReferences
}

// GetAllFeaturesAndOtherReferencesOk returns a tuple with the AllFeaturesAndOtherReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetAllFeaturesAndOtherReferencesOk() ([]BTMFeature134, bool) {
	if o == nil || o.AllFeaturesAndOtherReferences == nil {
		return nil, false
	}
	return o.AllFeaturesAndOtherReferences, true
}

// HasAllFeaturesAndOtherReferences returns a boolean if a field has been set.
func (o *BTMModel141) HasAllFeaturesAndOtherReferences() bool {
	if o != nil && o.AllFeaturesAndOtherReferences != nil {
		return true
	}

	return false
}

// SetAllFeaturesAndOtherReferences gets a reference to the given []BTMFeature134 and assigns it to the AllFeaturesAndOtherReferences field.
func (o *BTMModel141) SetAllFeaturesAndOtherReferences(v []BTMFeature134) {
	o.AllFeaturesAndOtherReferences = v
}

// GetAllFeaturesAndSubFeatures returns the AllFeaturesAndSubFeatures field value if set, zero value otherwise.
func (o *BTMModel141) GetAllFeaturesAndSubFeatures() []BTMFeature134 {
	if o == nil || o.AllFeaturesAndSubFeatures == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.AllFeaturesAndSubFeatures
}

// GetAllFeaturesAndSubFeaturesOk returns a tuple with the AllFeaturesAndSubFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetAllFeaturesAndSubFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.AllFeaturesAndSubFeatures == nil {
		return nil, false
	}
	return o.AllFeaturesAndSubFeatures, true
}

// HasAllFeaturesAndSubFeatures returns a boolean if a field has been set.
func (o *BTMModel141) HasAllFeaturesAndSubFeatures() bool {
	if o != nil && o.AllFeaturesAndSubFeatures != nil {
		return true
	}

	return false
}

// SetAllFeaturesAndSubFeatures gets a reference to the given []BTMFeature134 and assigns it to the AllFeaturesAndSubFeatures field.
func (o *BTMModel141) SetAllFeaturesAndSubFeatures(v []BTMFeature134) {
	o.AllFeaturesAndSubFeatures = v
}

// GetBelScriptLibraryMajorVersion returns the BelScriptLibraryMajorVersion field value if set, zero value otherwise.
func (o *BTMModel141) GetBelScriptLibraryMajorVersion() int32 {
	if o == nil || o.BelScriptLibraryMajorVersion == nil {
		var ret int32
		return ret
	}
	return *o.BelScriptLibraryMajorVersion
}

// GetBelScriptLibraryMajorVersionOk returns a tuple with the BelScriptLibraryMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetBelScriptLibraryMajorVersionOk() (*int32, bool) {
	if o == nil || o.BelScriptLibraryMajorVersion == nil {
		return nil, false
	}
	return o.BelScriptLibraryMajorVersion, true
}

// HasBelScriptLibraryMajorVersion returns a boolean if a field has been set.
func (o *BTMModel141) HasBelScriptLibraryMajorVersion() bool {
	if o != nil && o.BelScriptLibraryMajorVersion != nil {
		return true
	}

	return false
}

// SetBelScriptLibraryMajorVersion gets a reference to the given int32 and assigns it to the BelScriptLibraryMajorVersion field.
func (o *BTMModel141) SetBelScriptLibraryMajorVersion(v int32) {
	o.BelScriptLibraryMajorVersion = &v
}

// GetBelScriptLibraryVersion returns the BelScriptLibraryVersion field value if set, zero value otherwise.
func (o *BTMModel141) GetBelScriptLibraryVersion() string {
	if o == nil || o.BelScriptLibraryVersion == nil {
		var ret string
		return ret
	}
	return *o.BelScriptLibraryVersion
}

// GetBelScriptLibraryVersionOk returns a tuple with the BelScriptLibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetBelScriptLibraryVersionOk() (*string, bool) {
	if o == nil || o.BelScriptLibraryVersion == nil {
		return nil, false
	}
	return o.BelScriptLibraryVersion, true
}

// HasBelScriptLibraryVersion returns a boolean if a field has been set.
func (o *BTMModel141) HasBelScriptLibraryVersion() bool {
	if o != nil && o.BelScriptLibraryVersion != nil {
		return true
	}

	return false
}

// SetBelScriptLibraryVersion gets a reference to the given string and assigns it to the BelScriptLibraryVersion field.
func (o *BTMModel141) SetBelScriptLibraryVersion(v string) {
	o.BelScriptLibraryVersion = &v
}

// GetBelScriptLibraryVersionEnum returns the BelScriptLibraryVersionEnum field value if set, zero value otherwise.
func (o *BTMModel141) GetBelScriptLibraryVersionEnum() string {
	if o == nil || o.BelScriptLibraryVersionEnum == nil {
		var ret string
		return ret
	}
	return *o.BelScriptLibraryVersionEnum
}

// GetBelScriptLibraryVersionEnumOk returns a tuple with the BelScriptLibraryVersionEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetBelScriptLibraryVersionEnumOk() (*string, bool) {
	if o == nil || o.BelScriptLibraryVersionEnum == nil {
		return nil, false
	}
	return o.BelScriptLibraryVersionEnum, true
}

// HasBelScriptLibraryVersionEnum returns a boolean if a field has been set.
func (o *BTMModel141) HasBelScriptLibraryVersionEnum() bool {
	if o != nil && o.BelScriptLibraryVersionEnum != nil {
		return true
	}

	return false
}

// SetBelScriptLibraryVersionEnum gets a reference to the given string and assigns it to the BelScriptLibraryVersionEnum field.
func (o *BTMModel141) SetBelScriptLibraryVersionEnum(v string) {
	o.BelScriptLibraryVersionEnum = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMModel141) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMModel141) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMModel141) SetBtType(v string) {
	o.BtType = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *BTMModel141) GetChildren() []BTMNode19 {
	if o == nil || o.Children == nil {
		var ret []BTMNode19
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetChildrenOk() ([]BTMNode19, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *BTMModel141) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []BTMNode19 and assigns it to the Children field.
func (o *BTMModel141) SetChildren(v []BTMNode19) {
	o.Children = v
}

// GetConfigurationData returns the ConfigurationData field value if set, zero value otherwise.
func (o *BTMModel141) GetConfigurationData() BTMConfigurationData1560 {
	if o == nil || o.ConfigurationData == nil {
		var ret BTMConfigurationData1560
		return ret
	}
	return *o.ConfigurationData
}

// GetConfigurationDataOk returns a tuple with the ConfigurationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetConfigurationDataOk() (*BTMConfigurationData1560, bool) {
	if o == nil || o.ConfigurationData == nil {
		return nil, false
	}
	return o.ConfigurationData, true
}

// HasConfigurationData returns a boolean if a field has been set.
func (o *BTMModel141) HasConfigurationData() bool {
	if o != nil && o.ConfigurationData != nil {
		return true
	}

	return false
}

// SetConfigurationData gets a reference to the given BTMConfigurationData1560 and assigns it to the ConfigurationData field.
func (o *BTMModel141) SetConfigurationData(v BTMConfigurationData1560) {
	o.ConfigurationData = &v
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *BTMModel141) GetConfigured() bool {
	if o == nil || o.Configured == nil {
		var ret bool
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetConfiguredOk() (*bool, bool) {
	if o == nil || o.Configured == nil {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *BTMModel141) HasConfigured() bool {
	if o != nil && o.Configured != nil {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *BTMModel141) SetConfigured(v bool) {
	o.Configured = &v
}

// GetDeepImports returns the DeepImports field value if set, zero value otherwise.
func (o *BTMModel141) GetDeepImports() map[string][]BTImport {
	if o == nil || o.DeepImports == nil {
		var ret map[string][]BTImport
		return ret
	}
	return *o.DeepImports
}

// GetDeepImportsOk returns a tuple with the DeepImports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetDeepImportsOk() (*map[string][]BTImport, bool) {
	if o == nil || o.DeepImports == nil {
		return nil, false
	}
	return o.DeepImports, true
}

// HasDeepImports returns a boolean if a field has been set.
func (o *BTMModel141) HasDeepImports() bool {
	if o != nil && o.DeepImports != nil {
		return true
	}

	return false
}

// SetDeepImports gets a reference to the given map[string][]BTImport and assigns it to the DeepImports field.
func (o *BTMModel141) SetDeepImports(v map[string][]BTImport) {
	o.DeepImports = &v
}

// GetDefaultFeatures returns the DefaultFeatures field value if set, zero value otherwise.
func (o *BTMModel141) GetDefaultFeatures() BTDefaultFeatures119 {
	if o == nil || o.DefaultFeatures == nil {
		var ret BTDefaultFeatures119
		return ret
	}
	return *o.DefaultFeatures
}

// GetDefaultFeaturesOk returns a tuple with the DefaultFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetDefaultFeaturesOk() (*BTDefaultFeatures119, bool) {
	if o == nil || o.DefaultFeatures == nil {
		return nil, false
	}
	return o.DefaultFeatures, true
}

// HasDefaultFeatures returns a boolean if a field has been set.
func (o *BTMModel141) HasDefaultFeatures() bool {
	if o != nil && o.DefaultFeatures != nil {
		return true
	}

	return false
}

// SetDefaultFeatures gets a reference to the given BTDefaultFeatures119 and assigns it to the DefaultFeatures field.
func (o *BTMModel141) SetDefaultFeatures(v BTDefaultFeatures119) {
	o.DefaultFeatures = &v
}

// GetDefaultUnits returns the DefaultUnits field value if set, zero value otherwise.
func (o *BTMModel141) GetDefaultUnits() BTMUnitsDefault160 {
	if o == nil || o.DefaultUnits == nil {
		var ret BTMUnitsDefault160
		return ret
	}
	return *o.DefaultUnits
}

// GetDefaultUnitsOk returns a tuple with the DefaultUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetDefaultUnitsOk() (*BTMUnitsDefault160, bool) {
	if o == nil || o.DefaultUnits == nil {
		return nil, false
	}
	return o.DefaultUnits, true
}

// HasDefaultUnits returns a boolean if a field has been set.
func (o *BTMModel141) HasDefaultUnits() bool {
	if o != nil && o.DefaultUnits != nil {
		return true
	}

	return false
}

// SetDefaultUnits gets a reference to the given BTMUnitsDefault160 and assigns it to the DefaultUnits field.
func (o *BTMModel141) SetDefaultUnits(v BTMUnitsDefault160) {
	o.DefaultUnits = &v
}

// GetFeatureImports returns the FeatureImports field value if set, zero value otherwise.
func (o *BTMModel141) GetFeatureImports() map[string][]BTImport {
	if o == nil || o.FeatureImports == nil {
		var ret map[string][]BTImport
		return ret
	}
	return *o.FeatureImports
}

// GetFeatureImportsOk returns a tuple with the FeatureImports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetFeatureImportsOk() (*map[string][]BTImport, bool) {
	if o == nil || o.FeatureImports == nil {
		return nil, false
	}
	return o.FeatureImports, true
}

// HasFeatureImports returns a boolean if a field has been set.
func (o *BTMModel141) HasFeatureImports() bool {
	if o != nil && o.FeatureImports != nil {
		return true
	}

	return false
}

// SetFeatureImports gets a reference to the given map[string][]BTImport and assigns it to the FeatureImports field.
func (o *BTMModel141) SetFeatureImports(v map[string][]BTImport) {
	o.FeatureImports = &v
}

// GetFirstRollbackIndex returns the FirstRollbackIndex field value if set, zero value otherwise.
func (o *BTMModel141) GetFirstRollbackIndex() int32 {
	if o == nil || o.FirstRollbackIndex == nil {
		var ret int32
		return ret
	}
	return *o.FirstRollbackIndex
}

// GetFirstRollbackIndexOk returns a tuple with the FirstRollbackIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetFirstRollbackIndexOk() (*int32, bool) {
	if o == nil || o.FirstRollbackIndex == nil {
		return nil, false
	}
	return o.FirstRollbackIndex, true
}

// HasFirstRollbackIndex returns a boolean if a field has been set.
func (o *BTMModel141) HasFirstRollbackIndex() bool {
	if o != nil && o.FirstRollbackIndex != nil {
		return true
	}

	return false
}

// SetFirstRollbackIndex gets a reference to the given int32 and assigns it to the FirstRollbackIndex field.
func (o *BTMModel141) SetFirstRollbackIndex(v int32) {
	o.FirstRollbackIndex = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMModel141) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMModel141) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMModel141) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetImportSet returns the ImportSet field value if set, zero value otherwise.
func (o *BTMModel141) GetImportSet() []BTPModuleId235 {
	if o == nil || o.ImportSet == nil {
		var ret []BTPModuleId235
		return ret
	}
	return o.ImportSet
}

// GetImportSetOk returns a tuple with the ImportSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetImportSetOk() ([]BTPModuleId235, bool) {
	if o == nil || o.ImportSet == nil {
		return nil, false
	}
	return o.ImportSet, true
}

// HasImportSet returns a boolean if a field has been set.
func (o *BTMModel141) HasImportSet() bool {
	if o != nil && o.ImportSet != nil {
		return true
	}

	return false
}

// SetImportSet gets a reference to the given []BTPModuleId235 and assigns it to the ImportSet field.
func (o *BTMModel141) SetImportSet(v []BTPModuleId235) {
	o.ImportSet = v
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *BTMModel141) GetImports() []BTMImport136 {
	if o == nil || o.Imports == nil {
		var ret []BTMImport136
		return ret
	}
	return o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetImportsOk() ([]BTMImport136, bool) {
	if o == nil || o.Imports == nil {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *BTMModel141) HasImports() bool {
	if o != nil && o.Imports != nil {
		return true
	}

	return false
}

// SetImports gets a reference to the given []BTMImport136 and assigns it to the Imports field.
func (o *BTMModel141) SetImports(v []BTMImport136) {
	o.Imports = v
}

// GetIsVariableStudio returns the IsVariableStudio field value if set, zero value otherwise.
func (o *BTMModel141) GetIsVariableStudio() bool {
	if o == nil || o.IsVariableStudio == nil {
		var ret bool
		return ret
	}
	return *o.IsVariableStudio
}

// GetIsVariableStudioOk returns a tuple with the IsVariableStudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetIsVariableStudioOk() (*bool, bool) {
	if o == nil || o.IsVariableStudio == nil {
		return nil, false
	}
	return o.IsVariableStudio, true
}

// HasIsVariableStudio returns a boolean if a field has been set.
func (o *BTMModel141) HasIsVariableStudio() bool {
	if o != nil && o.IsVariableStudio != nil {
		return true
	}

	return false
}

// SetIsVariableStudio gets a reference to the given bool and assigns it to the IsVariableStudio field.
func (o *BTMModel141) SetIsVariableStudio(v bool) {
	o.IsVariableStudio = &v
}

// GetLastFeatureBeforeRollBack returns the LastFeatureBeforeRollBack field value if set, zero value otherwise.
func (o *BTMModel141) GetLastFeatureBeforeRollBack() BTMFeature134 {
	if o == nil || o.LastFeatureBeforeRollBack == nil {
		var ret BTMFeature134
		return ret
	}
	return *o.LastFeatureBeforeRollBack
}

// GetLastFeatureBeforeRollBackOk returns a tuple with the LastFeatureBeforeRollBack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetLastFeatureBeforeRollBackOk() (*BTMFeature134, bool) {
	if o == nil || o.LastFeatureBeforeRollBack == nil {
		return nil, false
	}
	return o.LastFeatureBeforeRollBack, true
}

// HasLastFeatureBeforeRollBack returns a boolean if a field has been set.
func (o *BTMModel141) HasLastFeatureBeforeRollBack() bool {
	if o != nil && o.LastFeatureBeforeRollBack != nil {
		return true
	}

	return false
}

// SetLastFeatureBeforeRollBack gets a reference to the given BTMFeature134 and assigns it to the LastFeatureBeforeRollBack field.
func (o *BTMModel141) SetLastFeatureBeforeRollBack(v BTMFeature134) {
	o.LastFeatureBeforeRollBack = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMModel141) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMModel141) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMModel141) SetName(v string) {
	o.Name = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMModel141) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMModel141) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMModel141) SetNodeId(v string) {
	o.NodeId = &v
}

// GetPartProperties returns the PartProperties field value if set, zero value otherwise.
func (o *BTMModel141) GetPartProperties() BTPartProperties293 {
	if o == nil || o.PartProperties == nil {
		var ret BTPartProperties293
		return ret
	}
	return *o.PartProperties
}

// GetPartPropertiesOk returns a tuple with the PartProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetPartPropertiesOk() (*BTPartProperties293, bool) {
	if o == nil || o.PartProperties == nil {
		return nil, false
	}
	return o.PartProperties, true
}

// HasPartProperties returns a boolean if a field has been set.
func (o *BTMModel141) HasPartProperties() bool {
	if o != nil && o.PartProperties != nil {
		return true
	}

	return false
}

// SetPartProperties gets a reference to the given BTPartProperties293 and assigns it to the PartProperties field.
func (o *BTMModel141) SetPartProperties(v BTPartProperties293) {
	o.PartProperties = &v
}

// GetPathToCache returns the PathToCache field value if set, zero value otherwise.
func (o *BTMModel141) GetPathToCache() BTCacheDataPath191 {
	if o == nil || o.PathToCache == nil {
		var ret BTCacheDataPath191
		return ret
	}
	return *o.PathToCache
}

// GetPathToCacheOk returns a tuple with the PathToCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetPathToCacheOk() (*BTCacheDataPath191, bool) {
	if o == nil || o.PathToCache == nil {
		return nil, false
	}
	return o.PathToCache, true
}

// HasPathToCache returns a boolean if a field has been set.
func (o *BTMModel141) HasPathToCache() bool {
	if o != nil && o.PathToCache != nil {
		return true
	}

	return false
}

// SetPathToCache gets a reference to the given BTCacheDataPath191 and assigns it to the PathToCache field.
func (o *BTMModel141) SetPathToCache(v BTCacheDataPath191) {
	o.PathToCache = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTMModel141) GetProperties() BTModelProperties1258 {
	if o == nil || o.Properties == nil {
		var ret BTModelProperties1258
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetPropertiesOk() (*BTModelProperties1258, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTMModel141) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given BTModelProperties1258 and assigns it to the Properties field.
func (o *BTMModel141) SetProperties(v BTModelProperties1258) {
	o.Properties = &v
}

// GetRollbackBar returns the RollbackBar field value if set, zero value otherwise.
func (o *BTMModel141) GetRollbackBar() BTMRollback150 {
	if o == nil || o.RollbackBar == nil {
		var ret BTMRollback150
		return ret
	}
	return *o.RollbackBar
}

// GetRollbackBarOk returns a tuple with the RollbackBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetRollbackBarOk() (*BTMRollback150, bool) {
	if o == nil || o.RollbackBar == nil {
		return nil, false
	}
	return o.RollbackBar, true
}

// HasRollbackBar returns a boolean if a field has been set.
func (o *BTMModel141) HasRollbackBar() bool {
	if o != nil && o.RollbackBar != nil {
		return true
	}

	return false
}

// SetRollbackBar gets a reference to the given BTMRollback150 and assigns it to the RollbackBar field.
func (o *BTMModel141) SetRollbackBar(v BTMRollback150) {
	o.RollbackBar = &v
}

// GetRolledBackToEnd returns the RolledBackToEnd field value if set, zero value otherwise.
func (o *BTMModel141) GetRolledBackToEnd() bool {
	if o == nil || o.RolledBackToEnd == nil {
		var ret bool
		return ret
	}
	return *o.RolledBackToEnd
}

// GetRolledBackToEndOk returns a tuple with the RolledBackToEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetRolledBackToEndOk() (*bool, bool) {
	if o == nil || o.RolledBackToEnd == nil {
		return nil, false
	}
	return o.RolledBackToEnd, true
}

// HasRolledBackToEnd returns a boolean if a field has been set.
func (o *BTMModel141) HasRolledBackToEnd() bool {
	if o != nil && o.RolledBackToEnd != nil {
		return true
	}

	return false
}

// SetRolledBackToEnd gets a reference to the given bool and assigns it to the RolledBackToEnd field.
func (o *BTMModel141) SetRolledBackToEnd(v bool) {
	o.RolledBackToEnd = &v
}

// GetVariableStudios returns the VariableStudios field value if set, zero value otherwise.
func (o *BTMModel141) GetVariableStudios() []BTMVariableStudioReference2764 {
	if o == nil || o.VariableStudios == nil {
		var ret []BTMVariableStudioReference2764
		return ret
	}
	return o.VariableStudios
}

// GetVariableStudiosOk returns a tuple with the VariableStudios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMModel141) GetVariableStudiosOk() ([]BTMVariableStudioReference2764, bool) {
	if o == nil || o.VariableStudios == nil {
		return nil, false
	}
	return o.VariableStudios, true
}

// HasVariableStudios returns a boolean if a field has been set.
func (o *BTMModel141) HasVariableStudios() bool {
	if o != nil && o.VariableStudios != nil {
		return true
	}

	return false
}

// SetVariableStudios gets a reference to the given []BTMVariableStudioReference2764 and assigns it to the VariableStudios field.
func (o *BTMModel141) SetVariableStudios(v []BTMVariableStudioReference2764) {
	o.VariableStudios = v
}

func (o BTMModel141) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllFeatures != nil {
		toSerialize["allFeatures"] = o.AllFeatures
	}
	if o.AllFeaturesAndOtherReferences != nil {
		toSerialize["allFeaturesAndOtherReferences"] = o.AllFeaturesAndOtherReferences
	}
	if o.AllFeaturesAndSubFeatures != nil {
		toSerialize["allFeaturesAndSubFeatures"] = o.AllFeaturesAndSubFeatures
	}
	if o.BelScriptLibraryMajorVersion != nil {
		toSerialize["belScriptLibraryMajorVersion"] = o.BelScriptLibraryMajorVersion
	}
	if o.BelScriptLibraryVersion != nil {
		toSerialize["belScriptLibraryVersion"] = o.BelScriptLibraryVersion
	}
	if o.BelScriptLibraryVersionEnum != nil {
		toSerialize["belScriptLibraryVersionEnum"] = o.BelScriptLibraryVersionEnum
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.ConfigurationData != nil {
		toSerialize["configurationData"] = o.ConfigurationData
	}
	if o.Configured != nil {
		toSerialize["configured"] = o.Configured
	}
	if o.DeepImports != nil {
		toSerialize["deepImports"] = o.DeepImports
	}
	if o.DefaultFeatures != nil {
		toSerialize["defaultFeatures"] = o.DefaultFeatures
	}
	if o.DefaultUnits != nil {
		toSerialize["defaultUnits"] = o.DefaultUnits
	}
	if o.FeatureImports != nil {
		toSerialize["featureImports"] = o.FeatureImports
	}
	if o.FirstRollbackIndex != nil {
		toSerialize["firstRollbackIndex"] = o.FirstRollbackIndex
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.ImportSet != nil {
		toSerialize["importSet"] = o.ImportSet
	}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	if o.IsVariableStudio != nil {
		toSerialize["isVariableStudio"] = o.IsVariableStudio
	}
	if o.LastFeatureBeforeRollBack != nil {
		toSerialize["lastFeatureBeforeRollBack"] = o.LastFeatureBeforeRollBack
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.PartProperties != nil {
		toSerialize["partProperties"] = o.PartProperties
	}
	if o.PathToCache != nil {
		toSerialize["pathToCache"] = o.PathToCache
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.RollbackBar != nil {
		toSerialize["rollbackBar"] = o.RollbackBar
	}
	if o.RolledBackToEnd != nil {
		toSerialize["rolledBackToEnd"] = o.RolledBackToEnd
	}
	if o.VariableStudios != nil {
		toSerialize["variableStudios"] = o.VariableStudios
	}
	return json.Marshal(toSerialize)
}

type NullableBTMModel141 struct {
	value *BTMModel141
	isSet bool
}

func (v NullableBTMModel141) Get() *BTMModel141 {
	return v.value
}

func (v *NullableBTMModel141) Set(val *BTMModel141) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMModel141) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMModel141) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMModel141(val *BTMModel141) *NullableBTMModel141 {
	return &NullableBTMModel141{value: val, isSet: true}
}

func (v NullableBTMModel141) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMModel141) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
