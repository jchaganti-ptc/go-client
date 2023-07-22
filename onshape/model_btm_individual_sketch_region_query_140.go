/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19303-3cbf47a47fe4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMIndividualSketchRegionQuery140 struct for BTMIndividualSketchRegionQuery140
type BTMIndividualSketchRegionQuery140 struct {
	BtType              *string                    `json:"btType,omitempty"`
	DeterministicIdList *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds    []string                   `json:"deterministicIds,omitempty"`
	ImportMicroversion  *string                    `json:"importMicroversion,omitempty"`
	NodeId              *string                    `json:"nodeId,omitempty"`
	Query               *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString         *string                    `json:"queryString,omitempty"`
	PersistentQuery     *BTPStatement269           `json:"persistentQuery,omitempty"`
	QueryStatement      *BTPStatement269           `json:"queryStatement,omitempty"`
	VariableName        *BTMIndividualQuery138     `json:"variableName,omitempty"`
	FeatureId           *string                    `json:"featureId,omitempty"`
	FilterInnerLoops    *bool                      `json:"filterInnerLoops,omitempty"`
}

// NewBTMIndividualSketchRegionQuery140 instantiates a new BTMIndividualSketchRegionQuery140 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualSketchRegionQuery140() *BTMIndividualSketchRegionQuery140 {
	this := BTMIndividualSketchRegionQuery140{}
	return &this
}

// NewBTMIndividualSketchRegionQuery140WithDefaults instantiates a new BTMIndividualSketchRegionQuery140 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualSketchRegionQuery140WithDefaults() *BTMIndividualSketchRegionQuery140 {
	this := BTMIndividualSketchRegionQuery140{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualSketchRegionQuery140) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMIndividualSketchRegionQuery140) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMIndividualSketchRegionQuery140) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMIndividualSketchRegionQuery140) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMIndividualSketchRegionQuery140) SetNodeId(v string) {
	o.NodeId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMIndividualSketchRegionQuery140) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMIndividualSketchRegionQuery140) SetQueryString(v string) {
	o.QueryString = &v
}

// GetPersistentQuery returns the PersistentQuery field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetPersistentQuery() BTPStatement269 {
	if o == nil || o.PersistentQuery == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.PersistentQuery
}

// GetPersistentQueryOk returns a tuple with the PersistentQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetPersistentQueryOk() (*BTPStatement269, bool) {
	if o == nil || o.PersistentQuery == nil {
		return nil, false
	}
	return o.PersistentQuery, true
}

// HasPersistentQuery returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasPersistentQuery() bool {
	if o != nil && o.PersistentQuery != nil {
		return true
	}

	return false
}

// SetPersistentQuery gets a reference to the given BTPStatement269 and assigns it to the PersistentQuery field.
func (o *BTMIndividualSketchRegionQuery140) SetPersistentQuery(v BTPStatement269) {
	o.PersistentQuery = &v
}

// GetQueryStatement returns the QueryStatement field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetQueryStatement() BTPStatement269 {
	if o == nil || o.QueryStatement == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.QueryStatement
}

// GetQueryStatementOk returns a tuple with the QueryStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetQueryStatementOk() (*BTPStatement269, bool) {
	if o == nil || o.QueryStatement == nil {
		return nil, false
	}
	return o.QueryStatement, true
}

// HasQueryStatement returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasQueryStatement() bool {
	if o != nil && o.QueryStatement != nil {
		return true
	}

	return false
}

// SetQueryStatement gets a reference to the given BTPStatement269 and assigns it to the QueryStatement field.
func (o *BTMIndividualSketchRegionQuery140) SetQueryStatement(v BTPStatement269) {
	o.QueryStatement = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetVariableName() BTMIndividualQuery138 {
	if o == nil || o.VariableName == nil {
		var ret BTMIndividualQuery138
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetVariableNameOk() (*BTMIndividualQuery138, bool) {
	if o == nil || o.VariableName == nil {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasVariableName() bool {
	if o != nil && o.VariableName != nil {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given BTMIndividualQuery138 and assigns it to the VariableName field.
func (o *BTMIndividualSketchRegionQuery140) SetVariableName(v BTMIndividualQuery138) {
	o.VariableName = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTMIndividualSketchRegionQuery140) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFilterInnerLoops returns the FilterInnerLoops field value if set, zero value otherwise.
func (o *BTMIndividualSketchRegionQuery140) GetFilterInnerLoops() bool {
	if o == nil || o.FilterInnerLoops == nil {
		var ret bool
		return ret
	}
	return *o.FilterInnerLoops
}

// GetFilterInnerLoopsOk returns a tuple with the FilterInnerLoops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualSketchRegionQuery140) GetFilterInnerLoopsOk() (*bool, bool) {
	if o == nil || o.FilterInnerLoops == nil {
		return nil, false
	}
	return o.FilterInnerLoops, true
}

// HasFilterInnerLoops returns a boolean if a field has been set.
func (o *BTMIndividualSketchRegionQuery140) HasFilterInnerLoops() bool {
	if o != nil && o.FilterInnerLoops != nil {
		return true
	}

	return false
}

// SetFilterInnerLoops gets a reference to the given bool and assigns it to the FilterInnerLoops field.
func (o *BTMIndividualSketchRegionQuery140) SetFilterInnerLoops(v bool) {
	o.FilterInnerLoops = &v
}

func (o BTMIndividualSketchRegionQuery140) MarshalJSON() ([]byte, error) {
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
	if o.PersistentQuery != nil {
		toSerialize["persistentQuery"] = o.PersistentQuery
	}
	if o.QueryStatement != nil {
		toSerialize["queryStatement"] = o.QueryStatement
	}
	if o.VariableName != nil {
		toSerialize["variableName"] = o.VariableName
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FilterInnerLoops != nil {
		toSerialize["filterInnerLoops"] = o.FilterInnerLoops
	}
	return json.Marshal(toSerialize)
}

type NullableBTMIndividualSketchRegionQuery140 struct {
	value *BTMIndividualSketchRegionQuery140
	isSet bool
}

func (v NullableBTMIndividualSketchRegionQuery140) Get() *BTMIndividualSketchRegionQuery140 {
	return v.value
}

func (v *NullableBTMIndividualSketchRegionQuery140) Set(val *BTMIndividualSketchRegionQuery140) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualSketchRegionQuery140) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualSketchRegionQuery140) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualSketchRegionQuery140(val *BTMIndividualSketchRegionQuery140) *NullableBTMIndividualSketchRegionQuery140 {
	return &NullableBTMIndividualSketchRegionQuery140{value: val, isSet: true}
}

func (v NullableBTMIndividualSketchRegionQuery140) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualSketchRegionQuery140) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
