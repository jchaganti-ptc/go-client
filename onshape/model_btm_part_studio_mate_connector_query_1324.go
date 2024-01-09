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

// BTMPartStudioMateConnectorQuery1324 struct for BTMPartStudioMateConnectorQuery1324
type BTMPartStudioMateConnectorQuery1324 struct {
	BtType                       *string                    `json:"btType,omitempty"`
	DeterministicIdList          *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds             []string                   `json:"deterministicIds,omitempty"`
	FullPathAsString             *string                    `json:"fullPathAsString,omitempty"`
	ImportMicroversion           *string                    `json:"importMicroversion,omitempty"`
	NodeId                       *string                    `json:"nodeId,omitempty"`
	Occurrence                   *BTOccurrence74            `json:"occurrence,omitempty"`
	Path                         []string                   `json:"path,omitempty"`
	Query                        *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString                  *string                    `json:"queryString,omitempty"`
	FeatureId                    *string                    `json:"featureId,omitempty"`
	FeatureIdWithOccurrence      *string                    `json:"featureIdWithOccurrence,omitempty"`
	PartStudioMateConnectorQuery *bool                      `json:"partStudioMateConnectorQuery,omitempty"`
	QueryData                    *string                    `json:"queryData,omitempty"`
}

// NewBTMPartStudioMateConnectorQuery1324 instantiates a new BTMPartStudioMateConnectorQuery1324 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMPartStudioMateConnectorQuery1324() *BTMPartStudioMateConnectorQuery1324 {
	this := BTMPartStudioMateConnectorQuery1324{}
	return &this
}

// NewBTMPartStudioMateConnectorQuery1324WithDefaults instantiates a new BTMPartStudioMateConnectorQuery1324 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMPartStudioMateConnectorQuery1324WithDefaults() *BTMPartStudioMateConnectorQuery1324 {
	this := BTMPartStudioMateConnectorQuery1324{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMPartStudioMateConnectorQuery1324) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMPartStudioMateConnectorQuery1324) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMPartStudioMateConnectorQuery1324) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetFullPathAsString() string {
	if o == nil || o.FullPathAsString == nil {
		var ret string
		return ret
	}
	return *o.FullPathAsString
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetFullPathAsStringOk() (*string, bool) {
	if o == nil || o.FullPathAsString == nil {
		return nil, false
	}
	return o.FullPathAsString, true
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasFullPathAsString() bool {
	if o != nil && o.FullPathAsString != nil {
		return true
	}

	return false
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *BTMPartStudioMateConnectorQuery1324) SetFullPathAsString(v string) {
	o.FullPathAsString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMPartStudioMateConnectorQuery1324) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMPartStudioMateConnectorQuery1324) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTMPartStudioMateConnectorQuery1324) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTMPartStudioMateConnectorQuery1324) SetPath(v []string) {
	o.Path = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMPartStudioMateConnectorQuery1324) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMPartStudioMateConnectorQuery1324) SetQueryString(v string) {
	o.QueryString = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMPartStudioMateConnectorQuery1324) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureIdWithOccurrence returns the FeatureIdWithOccurrence field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetFeatureIdWithOccurrence() string {
	if o == nil || o.FeatureIdWithOccurrence == nil {
		var ret string
		return ret
	}
	return *o.FeatureIdWithOccurrence
}

// GetFeatureIdWithOccurrenceOk returns a tuple with the FeatureIdWithOccurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetFeatureIdWithOccurrenceOk() (*string, bool) {
	if o == nil || o.FeatureIdWithOccurrence == nil {
		return nil, false
	}
	return o.FeatureIdWithOccurrence, true
}

// HasFeatureIdWithOccurrence returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasFeatureIdWithOccurrence() bool {
	if o != nil && o.FeatureIdWithOccurrence != nil {
		return true
	}

	return false
}

// SetFeatureIdWithOccurrence gets a reference to the given string and assigns it to the FeatureIdWithOccurrence field.
func (o *BTMPartStudioMateConnectorQuery1324) SetFeatureIdWithOccurrence(v string) {
	o.FeatureIdWithOccurrence = &v
}

// GetPartStudioMateConnectorQuery returns the PartStudioMateConnectorQuery field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetPartStudioMateConnectorQuery() bool {
	if o == nil || o.PartStudioMateConnectorQuery == nil {
		var ret bool
		return ret
	}
	return *o.PartStudioMateConnectorQuery
}

// GetPartStudioMateConnectorQueryOk returns a tuple with the PartStudioMateConnectorQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetPartStudioMateConnectorQueryOk() (*bool, bool) {
	if o == nil || o.PartStudioMateConnectorQuery == nil {
		return nil, false
	}
	return o.PartStudioMateConnectorQuery, true
}

// HasPartStudioMateConnectorQuery returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasPartStudioMateConnectorQuery() bool {
	if o != nil && o.PartStudioMateConnectorQuery != nil {
		return true
	}

	return false
}

// SetPartStudioMateConnectorQuery gets a reference to the given bool and assigns it to the PartStudioMateConnectorQuery field.
func (o *BTMPartStudioMateConnectorQuery1324) SetPartStudioMateConnectorQuery(v bool) {
	o.PartStudioMateConnectorQuery = &v
}

// GetQueryData returns the QueryData field value if set, zero value otherwise.
func (o *BTMPartStudioMateConnectorQuery1324) GetQueryData() string {
	if o == nil || o.QueryData == nil {
		var ret string
		return ret
	}
	return *o.QueryData
}

// GetQueryDataOk returns a tuple with the QueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMPartStudioMateConnectorQuery1324) GetQueryDataOk() (*string, bool) {
	if o == nil || o.QueryData == nil {
		return nil, false
	}
	return o.QueryData, true
}

// HasQueryData returns a boolean if a field has been set.
func (o *BTMPartStudioMateConnectorQuery1324) HasQueryData() bool {
	if o != nil && o.QueryData != nil {
		return true
	}

	return false
}

// SetQueryData gets a reference to the given string and assigns it to the QueryData field.
func (o *BTMPartStudioMateConnectorQuery1324) SetQueryData(v string) {
	o.QueryData = &v
}

func (o BTMPartStudioMateConnectorQuery1324) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.FullPathAsString != nil {
		toSerialize["fullPathAsString"] = o.FullPathAsString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FeatureIdWithOccurrence != nil {
		toSerialize["featureIdWithOccurrence"] = o.FeatureIdWithOccurrence
	}
	if o.PartStudioMateConnectorQuery != nil {
		toSerialize["partStudioMateConnectorQuery"] = o.PartStudioMateConnectorQuery
	}
	if o.QueryData != nil {
		toSerialize["queryData"] = o.QueryData
	}
	return json.Marshal(toSerialize)
}

type NullableBTMPartStudioMateConnectorQuery1324 struct {
	value *BTMPartStudioMateConnectorQuery1324
	isSet bool
}

func (v NullableBTMPartStudioMateConnectorQuery1324) Get() *BTMPartStudioMateConnectorQuery1324 {
	return v.value
}

func (v *NullableBTMPartStudioMateConnectorQuery1324) Set(val *BTMPartStudioMateConnectorQuery1324) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMPartStudioMateConnectorQuery1324) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMPartStudioMateConnectorQuery1324) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMPartStudioMateConnectorQuery1324(val *BTMPartStudioMateConnectorQuery1324) *NullableBTMPartStudioMateConnectorQuery1324 {
	return &NullableBTMPartStudioMateConnectorQuery1324{value: val, isSet: true}
}

func (v NullableBTMPartStudioMateConnectorQuery1324) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMPartStudioMateConnectorQuery1324) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
