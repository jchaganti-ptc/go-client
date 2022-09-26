/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6621-03f431879e4b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOccurrenceFilter166 struct for BTOccurrenceFilter166
type BTOccurrenceFilter166 struct {
	BtType                     *string `json:"btType,omitempty"`
	ExcludeFlattenedParts      *bool   `json:"excludeFlattenedParts,omitempty"`
	ExcludePatternInstances    *bool   `json:"excludePatternInstances,omitempty"`
	ExcludeReplicatedInstances *bool   `json:"excludeReplicatedInstances,omitempty"`
	ExcludeSketch              *bool   `json:"excludeSketch,omitempty"`
	ExcludeStandardContent     *bool   `json:"excludeStandardContent,omitempty"`
	ExcludeStudioInserts       *bool   `json:"excludeStudioInserts,omitempty"`
	ExcludeSubAssemblies       *bool   `json:"excludeSubAssemblies,omitempty"`
	ExcludeSuppressed          *bool   `json:"excludeSuppressed,omitempty"`
	IncludeAssemblyRoot        *bool   `json:"includeAssemblyRoot,omitempty"`
	IncludeParametricInstance  *bool   `json:"includeParametricInstance,omitempty"`
	IncludePatternOccurrence   *bool   `json:"includePatternOccurrence,omitempty"`
	SolidOrCompositeBodyOnly   *bool   `json:"solidOrCompositeBodyOnly,omitempty"`
	TopLevelOnly               *bool   `json:"topLevelOnly,omitempty"`
}

// NewBTOccurrenceFilter166 instantiates a new BTOccurrenceFilter166 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOccurrenceFilter166() *BTOccurrenceFilter166 {
	this := BTOccurrenceFilter166{}
	return &this
}

// NewBTOccurrenceFilter166WithDefaults instantiates a new BTOccurrenceFilter166 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOccurrenceFilter166WithDefaults() *BTOccurrenceFilter166 {
	this := BTOccurrenceFilter166{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOccurrenceFilter166) SetBtType(v string) {
	o.BtType = &v
}

// GetExcludeFlattenedParts returns the ExcludeFlattenedParts field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetExcludeFlattenedParts() bool {
	if o == nil || o.ExcludeFlattenedParts == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeFlattenedParts
}

// GetExcludeFlattenedPartsOk returns a tuple with the ExcludeFlattenedParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetExcludeFlattenedPartsOk() (*bool, bool) {
	if o == nil || o.ExcludeFlattenedParts == nil {
		return nil, false
	}
	return o.ExcludeFlattenedParts, true
}

// HasExcludeFlattenedParts returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasExcludeFlattenedParts() bool {
	if o != nil && o.ExcludeFlattenedParts != nil {
		return true
	}

	return false
}

// SetExcludeFlattenedParts gets a reference to the given bool and assigns it to the ExcludeFlattenedParts field.
func (o *BTOccurrenceFilter166) SetExcludeFlattenedParts(v bool) {
	o.ExcludeFlattenedParts = &v
}

// GetExcludePatternInstances returns the ExcludePatternInstances field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetExcludePatternInstances() bool {
	if o == nil || o.ExcludePatternInstances == nil {
		var ret bool
		return ret
	}
	return *o.ExcludePatternInstances
}

// GetExcludePatternInstancesOk returns a tuple with the ExcludePatternInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetExcludePatternInstancesOk() (*bool, bool) {
	if o == nil || o.ExcludePatternInstances == nil {
		return nil, false
	}
	return o.ExcludePatternInstances, true
}

// HasExcludePatternInstances returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasExcludePatternInstances() bool {
	if o != nil && o.ExcludePatternInstances != nil {
		return true
	}

	return false
}

// SetExcludePatternInstances gets a reference to the given bool and assigns it to the ExcludePatternInstances field.
func (o *BTOccurrenceFilter166) SetExcludePatternInstances(v bool) {
	o.ExcludePatternInstances = &v
}

// GetExcludeReplicatedInstances returns the ExcludeReplicatedInstances field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetExcludeReplicatedInstances() bool {
	if o == nil || o.ExcludeReplicatedInstances == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeReplicatedInstances
}

// GetExcludeReplicatedInstancesOk returns a tuple with the ExcludeReplicatedInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetExcludeReplicatedInstancesOk() (*bool, bool) {
	if o == nil || o.ExcludeReplicatedInstances == nil {
		return nil, false
	}
	return o.ExcludeReplicatedInstances, true
}

// HasExcludeReplicatedInstances returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasExcludeReplicatedInstances() bool {
	if o != nil && o.ExcludeReplicatedInstances != nil {
		return true
	}

	return false
}

// SetExcludeReplicatedInstances gets a reference to the given bool and assigns it to the ExcludeReplicatedInstances field.
func (o *BTOccurrenceFilter166) SetExcludeReplicatedInstances(v bool) {
	o.ExcludeReplicatedInstances = &v
}

// GetExcludeSketch returns the ExcludeSketch field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetExcludeSketch() bool {
	if o == nil || o.ExcludeSketch == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeSketch
}

// GetExcludeSketchOk returns a tuple with the ExcludeSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetExcludeSketchOk() (*bool, bool) {
	if o == nil || o.ExcludeSketch == nil {
		return nil, false
	}
	return o.ExcludeSketch, true
}

// HasExcludeSketch returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasExcludeSketch() bool {
	if o != nil && o.ExcludeSketch != nil {
		return true
	}

	return false
}

// SetExcludeSketch gets a reference to the given bool and assigns it to the ExcludeSketch field.
func (o *BTOccurrenceFilter166) SetExcludeSketch(v bool) {
	o.ExcludeSketch = &v
}

// GetExcludeStandardContent returns the ExcludeStandardContent field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetExcludeStandardContent() bool {
	if o == nil || o.ExcludeStandardContent == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeStandardContent
}

// GetExcludeStandardContentOk returns a tuple with the ExcludeStandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetExcludeStandardContentOk() (*bool, bool) {
	if o == nil || o.ExcludeStandardContent == nil {
		return nil, false
	}
	return o.ExcludeStandardContent, true
}

// HasExcludeStandardContent returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasExcludeStandardContent() bool {
	if o != nil && o.ExcludeStandardContent != nil {
		return true
	}

	return false
}

// SetExcludeStandardContent gets a reference to the given bool and assigns it to the ExcludeStandardContent field.
func (o *BTOccurrenceFilter166) SetExcludeStandardContent(v bool) {
	o.ExcludeStandardContent = &v
}

// GetExcludeStudioInserts returns the ExcludeStudioInserts field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetExcludeStudioInserts() bool {
	if o == nil || o.ExcludeStudioInserts == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeStudioInserts
}

// GetExcludeStudioInsertsOk returns a tuple with the ExcludeStudioInserts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetExcludeStudioInsertsOk() (*bool, bool) {
	if o == nil || o.ExcludeStudioInserts == nil {
		return nil, false
	}
	return o.ExcludeStudioInserts, true
}

// HasExcludeStudioInserts returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasExcludeStudioInserts() bool {
	if o != nil && o.ExcludeStudioInserts != nil {
		return true
	}

	return false
}

// SetExcludeStudioInserts gets a reference to the given bool and assigns it to the ExcludeStudioInserts field.
func (o *BTOccurrenceFilter166) SetExcludeStudioInserts(v bool) {
	o.ExcludeStudioInserts = &v
}

// GetExcludeSubAssemblies returns the ExcludeSubAssemblies field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetExcludeSubAssemblies() bool {
	if o == nil || o.ExcludeSubAssemblies == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeSubAssemblies
}

// GetExcludeSubAssembliesOk returns a tuple with the ExcludeSubAssemblies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetExcludeSubAssembliesOk() (*bool, bool) {
	if o == nil || o.ExcludeSubAssemblies == nil {
		return nil, false
	}
	return o.ExcludeSubAssemblies, true
}

// HasExcludeSubAssemblies returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasExcludeSubAssemblies() bool {
	if o != nil && o.ExcludeSubAssemblies != nil {
		return true
	}

	return false
}

// SetExcludeSubAssemblies gets a reference to the given bool and assigns it to the ExcludeSubAssemblies field.
func (o *BTOccurrenceFilter166) SetExcludeSubAssemblies(v bool) {
	o.ExcludeSubAssemblies = &v
}

// GetExcludeSuppressed returns the ExcludeSuppressed field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetExcludeSuppressed() bool {
	if o == nil || o.ExcludeSuppressed == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeSuppressed
}

// GetExcludeSuppressedOk returns a tuple with the ExcludeSuppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetExcludeSuppressedOk() (*bool, bool) {
	if o == nil || o.ExcludeSuppressed == nil {
		return nil, false
	}
	return o.ExcludeSuppressed, true
}

// HasExcludeSuppressed returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasExcludeSuppressed() bool {
	if o != nil && o.ExcludeSuppressed != nil {
		return true
	}

	return false
}

// SetExcludeSuppressed gets a reference to the given bool and assigns it to the ExcludeSuppressed field.
func (o *BTOccurrenceFilter166) SetExcludeSuppressed(v bool) {
	o.ExcludeSuppressed = &v
}

// GetIncludeAssemblyRoot returns the IncludeAssemblyRoot field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetIncludeAssemblyRoot() bool {
	if o == nil || o.IncludeAssemblyRoot == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAssemblyRoot
}

// GetIncludeAssemblyRootOk returns a tuple with the IncludeAssemblyRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetIncludeAssemblyRootOk() (*bool, bool) {
	if o == nil || o.IncludeAssemblyRoot == nil {
		return nil, false
	}
	return o.IncludeAssemblyRoot, true
}

// HasIncludeAssemblyRoot returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasIncludeAssemblyRoot() bool {
	if o != nil && o.IncludeAssemblyRoot != nil {
		return true
	}

	return false
}

// SetIncludeAssemblyRoot gets a reference to the given bool and assigns it to the IncludeAssemblyRoot field.
func (o *BTOccurrenceFilter166) SetIncludeAssemblyRoot(v bool) {
	o.IncludeAssemblyRoot = &v
}

// GetIncludeParametricInstance returns the IncludeParametricInstance field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetIncludeParametricInstance() bool {
	if o == nil || o.IncludeParametricInstance == nil {
		var ret bool
		return ret
	}
	return *o.IncludeParametricInstance
}

// GetIncludeParametricInstanceOk returns a tuple with the IncludeParametricInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetIncludeParametricInstanceOk() (*bool, bool) {
	if o == nil || o.IncludeParametricInstance == nil {
		return nil, false
	}
	return o.IncludeParametricInstance, true
}

// HasIncludeParametricInstance returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasIncludeParametricInstance() bool {
	if o != nil && o.IncludeParametricInstance != nil {
		return true
	}

	return false
}

// SetIncludeParametricInstance gets a reference to the given bool and assigns it to the IncludeParametricInstance field.
func (o *BTOccurrenceFilter166) SetIncludeParametricInstance(v bool) {
	o.IncludeParametricInstance = &v
}

// GetIncludePatternOccurrence returns the IncludePatternOccurrence field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetIncludePatternOccurrence() bool {
	if o == nil || o.IncludePatternOccurrence == nil {
		var ret bool
		return ret
	}
	return *o.IncludePatternOccurrence
}

// GetIncludePatternOccurrenceOk returns a tuple with the IncludePatternOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetIncludePatternOccurrenceOk() (*bool, bool) {
	if o == nil || o.IncludePatternOccurrence == nil {
		return nil, false
	}
	return o.IncludePatternOccurrence, true
}

// HasIncludePatternOccurrence returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasIncludePatternOccurrence() bool {
	if o != nil && o.IncludePatternOccurrence != nil {
		return true
	}

	return false
}

// SetIncludePatternOccurrence gets a reference to the given bool and assigns it to the IncludePatternOccurrence field.
func (o *BTOccurrenceFilter166) SetIncludePatternOccurrence(v bool) {
	o.IncludePatternOccurrence = &v
}

// GetSolidOrCompositeBodyOnly returns the SolidOrCompositeBodyOnly field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetSolidOrCompositeBodyOnly() bool {
	if o == nil || o.SolidOrCompositeBodyOnly == nil {
		var ret bool
		return ret
	}
	return *o.SolidOrCompositeBodyOnly
}

// GetSolidOrCompositeBodyOnlyOk returns a tuple with the SolidOrCompositeBodyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetSolidOrCompositeBodyOnlyOk() (*bool, bool) {
	if o == nil || o.SolidOrCompositeBodyOnly == nil {
		return nil, false
	}
	return o.SolidOrCompositeBodyOnly, true
}

// HasSolidOrCompositeBodyOnly returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasSolidOrCompositeBodyOnly() bool {
	if o != nil && o.SolidOrCompositeBodyOnly != nil {
		return true
	}

	return false
}

// SetSolidOrCompositeBodyOnly gets a reference to the given bool and assigns it to the SolidOrCompositeBodyOnly field.
func (o *BTOccurrenceFilter166) SetSolidOrCompositeBodyOnly(v bool) {
	o.SolidOrCompositeBodyOnly = &v
}

// GetTopLevelOnly returns the TopLevelOnly field value if set, zero value otherwise.
func (o *BTOccurrenceFilter166) GetTopLevelOnly() bool {
	if o == nil || o.TopLevelOnly == nil {
		var ret bool
		return ret
	}
	return *o.TopLevelOnly
}

// GetTopLevelOnlyOk returns a tuple with the TopLevelOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOccurrenceFilter166) GetTopLevelOnlyOk() (*bool, bool) {
	if o == nil || o.TopLevelOnly == nil {
		return nil, false
	}
	return o.TopLevelOnly, true
}

// HasTopLevelOnly returns a boolean if a field has been set.
func (o *BTOccurrenceFilter166) HasTopLevelOnly() bool {
	if o != nil && o.TopLevelOnly != nil {
		return true
	}

	return false
}

// SetTopLevelOnly gets a reference to the given bool and assigns it to the TopLevelOnly field.
func (o *BTOccurrenceFilter166) SetTopLevelOnly(v bool) {
	o.TopLevelOnly = &v
}

func (o BTOccurrenceFilter166) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ExcludeFlattenedParts != nil {
		toSerialize["excludeFlattenedParts"] = o.ExcludeFlattenedParts
	}
	if o.ExcludePatternInstances != nil {
		toSerialize["excludePatternInstances"] = o.ExcludePatternInstances
	}
	if o.ExcludeReplicatedInstances != nil {
		toSerialize["excludeReplicatedInstances"] = o.ExcludeReplicatedInstances
	}
	if o.ExcludeSketch != nil {
		toSerialize["excludeSketch"] = o.ExcludeSketch
	}
	if o.ExcludeStandardContent != nil {
		toSerialize["excludeStandardContent"] = o.ExcludeStandardContent
	}
	if o.ExcludeStudioInserts != nil {
		toSerialize["excludeStudioInserts"] = o.ExcludeStudioInserts
	}
	if o.ExcludeSubAssemblies != nil {
		toSerialize["excludeSubAssemblies"] = o.ExcludeSubAssemblies
	}
	if o.ExcludeSuppressed != nil {
		toSerialize["excludeSuppressed"] = o.ExcludeSuppressed
	}
	if o.IncludeAssemblyRoot != nil {
		toSerialize["includeAssemblyRoot"] = o.IncludeAssemblyRoot
	}
	if o.IncludeParametricInstance != nil {
		toSerialize["includeParametricInstance"] = o.IncludeParametricInstance
	}
	if o.IncludePatternOccurrence != nil {
		toSerialize["includePatternOccurrence"] = o.IncludePatternOccurrence
	}
	if o.SolidOrCompositeBodyOnly != nil {
		toSerialize["solidOrCompositeBodyOnly"] = o.SolidOrCompositeBodyOnly
	}
	if o.TopLevelOnly != nil {
		toSerialize["topLevelOnly"] = o.TopLevelOnly
	}
	return json.Marshal(toSerialize)
}

type NullableBTOccurrenceFilter166 struct {
	value *BTOccurrenceFilter166
	isSet bool
}

func (v NullableBTOccurrenceFilter166) Get() *BTOccurrenceFilter166 {
	return v.value
}

func (v *NullableBTOccurrenceFilter166) Set(val *BTOccurrenceFilter166) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOccurrenceFilter166) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOccurrenceFilter166) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOccurrenceFilter166(val *BTOccurrenceFilter166) *NullableBTOccurrenceFilter166 {
	return &NullableBTOccurrenceFilter166{value: val, isSet: true}
}

func (v NullableBTOccurrenceFilter166) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOccurrenceFilter166) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
