/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6358-16771368e3d1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsTable1073 struct for BTBillOfMaterialsTable1073
type BTBillOfMaterialsTable1073 struct {
	AllRowValues                            [][]string              `json:"allRowValues,omitempty"`
	ColumnCount                             *int32                  `json:"columnCount,omitempty"`
	FrozenColumns                           *int32                  `json:"frozenColumns,omitempty"`
	IsFailed                                *bool                   `json:"isFailed,omitempty"`
	NodeId                                  *string                 `json:"nodeId,omitempty"`
	ReadOnly                                *bool                   `json:"readOnly,omitempty"`
	RowCount                                *int32                  `json:"rowCount,omitempty"`
	SortOrder                               *BTTableSortOrder4371   `json:"sortOrder,omitempty"`
	TableColumns                            []BTTableColumnInfo1222 `json:"tableColumns,omitempty"`
	TableId                                 *string                 `json:"tableId,omitempty"`
	TableRows                               []BTTableRow1054        `json:"tableRows,omitempty"`
	Title                                   *string                 `json:"title,omitempty"`
	BtType                                  *string                 `json:"btType,omitempty"`
	FailedMetadataRepresentativeOccurrences []string                `json:"failedMetadataRepresentativeOccurrences,omitempty"`
	Indented                                *bool                   `json:"indented,omitempty"`
	PartNumber                              *string                 `json:"partNumber,omitempty"`
	Revision                                *string                 `json:"revision,omitempty"`
	ShowingExcluded                         *bool                   `json:"showingExcluded,omitempty"`
}

// NewBTBillOfMaterialsTable1073 instantiates a new BTBillOfMaterialsTable1073 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsTable1073() *BTBillOfMaterialsTable1073 {
	this := BTBillOfMaterialsTable1073{}
	return &this
}

// NewBTBillOfMaterialsTable1073WithDefaults instantiates a new BTBillOfMaterialsTable1073 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsTable1073WithDefaults() *BTBillOfMaterialsTable1073 {
	this := BTBillOfMaterialsTable1073{}
	return &this
}

// GetAllRowValues returns the AllRowValues field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetAllRowValues() [][]string {
	if o == nil || o.AllRowValues == nil {
		var ret [][]string
		return ret
	}
	return o.AllRowValues
}

// GetAllRowValuesOk returns a tuple with the AllRowValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetAllRowValuesOk() ([][]string, bool) {
	if o == nil || o.AllRowValues == nil {
		return nil, false
	}
	return o.AllRowValues, true
}

// HasAllRowValues returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasAllRowValues() bool {
	if o != nil && o.AllRowValues != nil {
		return true
	}

	return false
}

// SetAllRowValues gets a reference to the given [][]string and assigns it to the AllRowValues field.
func (o *BTBillOfMaterialsTable1073) SetAllRowValues(v [][]string) {
	o.AllRowValues = v
}

// GetColumnCount returns the ColumnCount field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetColumnCount() int32 {
	if o == nil || o.ColumnCount == nil {
		var ret int32
		return ret
	}
	return *o.ColumnCount
}

// GetColumnCountOk returns a tuple with the ColumnCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetColumnCountOk() (*int32, bool) {
	if o == nil || o.ColumnCount == nil {
		return nil, false
	}
	return o.ColumnCount, true
}

// HasColumnCount returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasColumnCount() bool {
	if o != nil && o.ColumnCount != nil {
		return true
	}

	return false
}

// SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.
func (o *BTBillOfMaterialsTable1073) SetColumnCount(v int32) {
	o.ColumnCount = &v
}

// GetFrozenColumns returns the FrozenColumns field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetFrozenColumns() int32 {
	if o == nil || o.FrozenColumns == nil {
		var ret int32
		return ret
	}
	return *o.FrozenColumns
}

// GetFrozenColumnsOk returns a tuple with the FrozenColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetFrozenColumnsOk() (*int32, bool) {
	if o == nil || o.FrozenColumns == nil {
		return nil, false
	}
	return o.FrozenColumns, true
}

// HasFrozenColumns returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasFrozenColumns() bool {
	if o != nil && o.FrozenColumns != nil {
		return true
	}

	return false
}

// SetFrozenColumns gets a reference to the given int32 and assigns it to the FrozenColumns field.
func (o *BTBillOfMaterialsTable1073) SetFrozenColumns(v int32) {
	o.FrozenColumns = &v
}

// GetIsFailed returns the IsFailed field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetIsFailed() bool {
	if o == nil || o.IsFailed == nil {
		var ret bool
		return ret
	}
	return *o.IsFailed
}

// GetIsFailedOk returns a tuple with the IsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetIsFailedOk() (*bool, bool) {
	if o == nil || o.IsFailed == nil {
		return nil, false
	}
	return o.IsFailed, true
}

// HasIsFailed returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasIsFailed() bool {
	if o != nil && o.IsFailed != nil {
		return true
	}

	return false
}

// SetIsFailed gets a reference to the given bool and assigns it to the IsFailed field.
func (o *BTBillOfMaterialsTable1073) SetIsFailed(v bool) {
	o.IsFailed = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTBillOfMaterialsTable1073) SetNodeId(v string) {
	o.NodeId = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BTBillOfMaterialsTable1073) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetRowCount() int32 {
	if o == nil || o.RowCount == nil {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetRowCountOk() (*int32, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasRowCount() bool {
	if o != nil && o.RowCount != nil {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *BTBillOfMaterialsTable1073) SetRowCount(v int32) {
	o.RowCount = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetSortOrder() BTTableSortOrder4371 {
	if o == nil || o.SortOrder == nil {
		var ret BTTableSortOrder4371
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetSortOrderOk() (*BTTableSortOrder4371, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given BTTableSortOrder4371 and assigns it to the SortOrder field.
func (o *BTBillOfMaterialsTable1073) SetSortOrder(v BTTableSortOrder4371) {
	o.SortOrder = &v
}

// GetTableColumns returns the TableColumns field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetTableColumns() []BTTableColumnInfo1222 {
	if o == nil || o.TableColumns == nil {
		var ret []BTTableColumnInfo1222
		return ret
	}
	return o.TableColumns
}

// GetTableColumnsOk returns a tuple with the TableColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetTableColumnsOk() ([]BTTableColumnInfo1222, bool) {
	if o == nil || o.TableColumns == nil {
		return nil, false
	}
	return o.TableColumns, true
}

// HasTableColumns returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasTableColumns() bool {
	if o != nil && o.TableColumns != nil {
		return true
	}

	return false
}

// SetTableColumns gets a reference to the given []BTTableColumnInfo1222 and assigns it to the TableColumns field.
func (o *BTBillOfMaterialsTable1073) SetTableColumns(v []BTTableColumnInfo1222) {
	o.TableColumns = v
}

// GetTableId returns the TableId field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetTableId() string {
	if o == nil || o.TableId == nil {
		var ret string
		return ret
	}
	return *o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetTableIdOk() (*string, bool) {
	if o == nil || o.TableId == nil {
		return nil, false
	}
	return o.TableId, true
}

// HasTableId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasTableId() bool {
	if o != nil && o.TableId != nil {
		return true
	}

	return false
}

// SetTableId gets a reference to the given string and assigns it to the TableId field.
func (o *BTBillOfMaterialsTable1073) SetTableId(v string) {
	o.TableId = &v
}

// GetTableRows returns the TableRows field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetTableRows() []BTTableRow1054 {
	if o == nil || o.TableRows == nil {
		var ret []BTTableRow1054
		return ret
	}
	return o.TableRows
}

// GetTableRowsOk returns a tuple with the TableRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetTableRowsOk() ([]BTTableRow1054, bool) {
	if o == nil || o.TableRows == nil {
		return nil, false
	}
	return o.TableRows, true
}

// HasTableRows returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasTableRows() bool {
	if o != nil && o.TableRows != nil {
		return true
	}

	return false
}

// SetTableRows gets a reference to the given []BTTableRow1054 and assigns it to the TableRows field.
func (o *BTBillOfMaterialsTable1073) SetTableRows(v []BTTableRow1054) {
	o.TableRows = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BTBillOfMaterialsTable1073) SetTitle(v string) {
	o.Title = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBillOfMaterialsTable1073) SetBtType(v string) {
	o.BtType = &v
}

// GetFailedMetadataRepresentativeOccurrences returns the FailedMetadataRepresentativeOccurrences field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetFailedMetadataRepresentativeOccurrences() []string {
	if o == nil || o.FailedMetadataRepresentativeOccurrences == nil {
		var ret []string
		return ret
	}
	return o.FailedMetadataRepresentativeOccurrences
}

// GetFailedMetadataRepresentativeOccurrencesOk returns a tuple with the FailedMetadataRepresentativeOccurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetFailedMetadataRepresentativeOccurrencesOk() ([]string, bool) {
	if o == nil || o.FailedMetadataRepresentativeOccurrences == nil {
		return nil, false
	}
	return o.FailedMetadataRepresentativeOccurrences, true
}

// HasFailedMetadataRepresentativeOccurrences returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasFailedMetadataRepresentativeOccurrences() bool {
	if o != nil && o.FailedMetadataRepresentativeOccurrences != nil {
		return true
	}

	return false
}

// SetFailedMetadataRepresentativeOccurrences gets a reference to the given []string and assigns it to the FailedMetadataRepresentativeOccurrences field.
func (o *BTBillOfMaterialsTable1073) SetFailedMetadataRepresentativeOccurrences(v []string) {
	o.FailedMetadataRepresentativeOccurrences = v
}

// GetIndented returns the Indented field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetIndented() bool {
	if o == nil || o.Indented == nil {
		var ret bool
		return ret
	}
	return *o.Indented
}

// GetIndentedOk returns a tuple with the Indented field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetIndentedOk() (*bool, bool) {
	if o == nil || o.Indented == nil {
		return nil, false
	}
	return o.Indented, true
}

// HasIndented returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasIndented() bool {
	if o != nil && o.Indented != nil {
		return true
	}

	return false
}

// SetIndented gets a reference to the given bool and assigns it to the Indented field.
func (o *BTBillOfMaterialsTable1073) SetIndented(v bool) {
	o.Indented = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTBillOfMaterialsTable1073) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTBillOfMaterialsTable1073) SetRevision(v string) {
	o.Revision = &v
}

// GetShowingExcluded returns the ShowingExcluded field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTable1073) GetShowingExcluded() bool {
	if o == nil || o.ShowingExcluded == nil {
		var ret bool
		return ret
	}
	return *o.ShowingExcluded
}

// GetShowingExcludedOk returns a tuple with the ShowingExcluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTable1073) GetShowingExcludedOk() (*bool, bool) {
	if o == nil || o.ShowingExcluded == nil {
		return nil, false
	}
	return o.ShowingExcluded, true
}

// HasShowingExcluded returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTable1073) HasShowingExcluded() bool {
	if o != nil && o.ShowingExcluded != nil {
		return true
	}

	return false
}

// SetShowingExcluded gets a reference to the given bool and assigns it to the ShowingExcluded field.
func (o *BTBillOfMaterialsTable1073) SetShowingExcluded(v bool) {
	o.ShowingExcluded = &v
}

func (o BTBillOfMaterialsTable1073) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllRowValues != nil {
		toSerialize["allRowValues"] = o.AllRowValues
	}
	if o.ColumnCount != nil {
		toSerialize["columnCount"] = o.ColumnCount
	}
	if o.FrozenColumns != nil {
		toSerialize["frozenColumns"] = o.FrozenColumns
	}
	if o.IsFailed != nil {
		toSerialize["isFailed"] = o.IsFailed
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.RowCount != nil {
		toSerialize["rowCount"] = o.RowCount
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if o.TableColumns != nil {
		toSerialize["tableColumns"] = o.TableColumns
	}
	if o.TableId != nil {
		toSerialize["tableId"] = o.TableId
	}
	if o.TableRows != nil {
		toSerialize["tableRows"] = o.TableRows
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FailedMetadataRepresentativeOccurrences != nil {
		toSerialize["failedMetadataRepresentativeOccurrences"] = o.FailedMetadataRepresentativeOccurrences
	}
	if o.Indented != nil {
		toSerialize["indented"] = o.Indented
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.ShowingExcluded != nil {
		toSerialize["showingExcluded"] = o.ShowingExcluded
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsTable1073 struct {
	value *BTBillOfMaterialsTable1073
	isSet bool
}

func (v NullableBTBillOfMaterialsTable1073) Get() *BTBillOfMaterialsTable1073 {
	return v.value
}

func (v *NullableBTBillOfMaterialsTable1073) Set(val *BTBillOfMaterialsTable1073) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsTable1073) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsTable1073) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsTable1073(val *BTBillOfMaterialsTable1073) *NullableBTBillOfMaterialsTable1073 {
	return &NullableBTBillOfMaterialsTable1073{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsTable1073) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsTable1073) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
