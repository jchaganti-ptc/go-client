/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMInContextQuery2254 struct for BTMInContextQuery2254
type BTMInContextQuery2254 struct {
	BtType                     *string                    `json:"btType,omitempty"`
	DeterministicIdList        *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds           []string                   `json:"deterministicIds,omitempty"`
	GenerateSectionEntityQuery *bool                      `json:"generateSectionEntityQuery,omitempty"`
	GeneratedSectionQueryId    *string                    `json:"generatedSectionQueryId,omitempty"`
	ImportMicroversion         *string                    `json:"importMicroversion,omitempty"`
	NodeId                     *string                    `json:"nodeId,omitempty"`
	Query                      *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString                *string                    `json:"queryString,omitempty"`
	Path                       []string                   `json:"path,omitempty"`
}

// NewBTMInContextQuery2254 instantiates a new BTMInContextQuery2254 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMInContextQuery2254() *BTMInContextQuery2254 {
	this := BTMInContextQuery2254{}
	return &this
}

// NewBTMInContextQuery2254WithDefaults instantiates a new BTMInContextQuery2254 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMInContextQuery2254WithDefaults() *BTMInContextQuery2254 {
	this := BTMInContextQuery2254{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMInContextQuery2254) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMInContextQuery2254) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMInContextQuery2254) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetGenerateSectionEntityQuery returns the GenerateSectionEntityQuery field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetGenerateSectionEntityQuery() bool {
	if o == nil || o.GenerateSectionEntityQuery == nil {
		var ret bool
		return ret
	}
	return *o.GenerateSectionEntityQuery
}

// GetGenerateSectionEntityQueryOk returns a tuple with the GenerateSectionEntityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetGenerateSectionEntityQueryOk() (*bool, bool) {
	if o == nil || o.GenerateSectionEntityQuery == nil {
		return nil, false
	}
	return o.GenerateSectionEntityQuery, true
}

// HasGenerateSectionEntityQuery returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasGenerateSectionEntityQuery() bool {
	if o != nil && o.GenerateSectionEntityQuery != nil {
		return true
	}

	return false
}

// SetGenerateSectionEntityQuery gets a reference to the given bool and assigns it to the GenerateSectionEntityQuery field.
func (o *BTMInContextQuery2254) SetGenerateSectionEntityQuery(v bool) {
	o.GenerateSectionEntityQuery = &v
}

// GetGeneratedSectionQueryId returns the GeneratedSectionQueryId field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetGeneratedSectionQueryId() string {
	if o == nil || o.GeneratedSectionQueryId == nil {
		var ret string
		return ret
	}
	return *o.GeneratedSectionQueryId
}

// GetGeneratedSectionQueryIdOk returns a tuple with the GeneratedSectionQueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetGeneratedSectionQueryIdOk() (*string, bool) {
	if o == nil || o.GeneratedSectionQueryId == nil {
		return nil, false
	}
	return o.GeneratedSectionQueryId, true
}

// HasGeneratedSectionQueryId returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasGeneratedSectionQueryId() bool {
	if o != nil && o.GeneratedSectionQueryId != nil {
		return true
	}

	return false
}

// SetGeneratedSectionQueryId gets a reference to the given string and assigns it to the GeneratedSectionQueryId field.
func (o *BTMInContextQuery2254) SetGeneratedSectionQueryId(v string) {
	o.GeneratedSectionQueryId = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMInContextQuery2254) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMInContextQuery2254) SetNodeId(v string) {
	o.NodeId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMInContextQuery2254) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMInContextQuery2254) SetQueryString(v string) {
	o.QueryString = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTMInContextQuery2254) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMInContextQuery2254) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTMInContextQuery2254) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTMInContextQuery2254) SetPath(v []string) {
	o.Path = v
}

func (o BTMInContextQuery2254) MarshalJSON() ([]byte, error) {
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
	if o.GenerateSectionEntityQuery != nil {
		toSerialize["generateSectionEntityQuery"] = o.GenerateSectionEntityQuery
	}
	if o.GeneratedSectionQueryId != nil {
		toSerialize["generatedSectionQueryId"] = o.GeneratedSectionQueryId
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableBTMInContextQuery2254 struct {
	value *BTMInContextQuery2254
	isSet bool
}

func (v NullableBTMInContextQuery2254) Get() *BTMInContextQuery2254 {
	return v.value
}

func (v *NullableBTMInContextQuery2254) Set(val *BTMInContextQuery2254) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMInContextQuery2254) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMInContextQuery2254) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMInContextQuery2254(val *BTMInContextQuery2254) *NullableBTMInContextQuery2254 {
	return &NullableBTMInContextQuery2254{value: val, isSet: true}
}

func (v NullableBTMInContextQuery2254) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMInContextQuery2254) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
