/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6415-48a6b2252b8c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMetadataPropertyInfo struct for BTMetadataPropertyInfo
type BTMetadataPropertyInfo struct {
	ComputedProperty      *bool                            `json:"computedProperty,omitempty"`
	ComputedPropertyError *string                          `json:"computedPropertyError,omitempty"`
	DefaultValue          interface{}                      `json:"defaultValue,omitempty"`
	Dirty                 *bool                            `json:"dirty,omitempty"`
	Editable              *bool                            `json:"editable,omitempty"`
	EditableInUi          *bool                            `json:"editableInUi,omitempty"`
	EnumValues            []BTMetadataEnumValueInfo        `json:"enumValues,omitempty"`
	InitialValue          *map[string]interface{}          `json:"initialValue,omitempty"`
	Multivalued           *bool                            `json:"multivalued,omitempty"`
	Name                  *string                          `json:"name,omitempty"`
	PropertyId            *string                          `json:"propertyId,omitempty"`
	PropertySource        *int32                           `json:"propertySource,omitempty"`
	Required              *bool                            `json:"required,omitempty"`
	SchemaId              *string                          `json:"schemaId,omitempty"`
	UiHints               *BTMetadataPropertyUiHintsInfo   `json:"uiHints,omitempty"`
	Validator             *BTMetadataPropertyValidatorInfo `json:"validator,omitempty"`
	Value                 interface{}                      `json:"value,omitempty"`
	ValueType             *string                          `json:"valueType,omitempty"`
}

// NewBTMetadataPropertyInfo instantiates a new BTMetadataPropertyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataPropertyInfo() *BTMetadataPropertyInfo {
	this := BTMetadataPropertyInfo{}
	return &this
}

// NewBTMetadataPropertyInfoWithDefaults instantiates a new BTMetadataPropertyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataPropertyInfoWithDefaults() *BTMetadataPropertyInfo {
	this := BTMetadataPropertyInfo{}
	return &this
}

// GetComputedProperty returns the ComputedProperty field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetComputedProperty() bool {
	if o == nil || o.ComputedProperty == nil {
		var ret bool
		return ret
	}
	return *o.ComputedProperty
}

// GetComputedPropertyOk returns a tuple with the ComputedProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetComputedPropertyOk() (*bool, bool) {
	if o == nil || o.ComputedProperty == nil {
		return nil, false
	}
	return o.ComputedProperty, true
}

// HasComputedProperty returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasComputedProperty() bool {
	if o != nil && o.ComputedProperty != nil {
		return true
	}

	return false
}

// SetComputedProperty gets a reference to the given bool and assigns it to the ComputedProperty field.
func (o *BTMetadataPropertyInfo) SetComputedProperty(v bool) {
	o.ComputedProperty = &v
}

// GetComputedPropertyError returns the ComputedPropertyError field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetComputedPropertyError() string {
	if o == nil || o.ComputedPropertyError == nil {
		var ret string
		return ret
	}
	return *o.ComputedPropertyError
}

// GetComputedPropertyErrorOk returns a tuple with the ComputedPropertyError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetComputedPropertyErrorOk() (*string, bool) {
	if o == nil || o.ComputedPropertyError == nil {
		return nil, false
	}
	return o.ComputedPropertyError, true
}

// HasComputedPropertyError returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasComputedPropertyError() bool {
	if o != nil && o.ComputedPropertyError != nil {
		return true
	}

	return false
}

// SetComputedPropertyError gets a reference to the given string and assigns it to the ComputedPropertyError field.
func (o *BTMetadataPropertyInfo) SetComputedPropertyError(v string) {
	o.ComputedPropertyError = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BTMetadataPropertyInfo) GetDefaultValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BTMetadataPropertyInfo) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *BTMetadataPropertyInfo) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetDirty returns the Dirty field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetDirty() bool {
	if o == nil || o.Dirty == nil {
		var ret bool
		return ret
	}
	return *o.Dirty
}

// GetDirtyOk returns a tuple with the Dirty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetDirtyOk() (*bool, bool) {
	if o == nil || o.Dirty == nil {
		return nil, false
	}
	return o.Dirty, true
}

// HasDirty returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasDirty() bool {
	if o != nil && o.Dirty != nil {
		return true
	}

	return false
}

// SetDirty gets a reference to the given bool and assigns it to the Dirty field.
func (o *BTMetadataPropertyInfo) SetDirty(v bool) {
	o.Dirty = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *BTMetadataPropertyInfo) SetEditable(v bool) {
	o.Editable = &v
}

// GetEditableInUi returns the EditableInUi field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetEditableInUi() bool {
	if o == nil || o.EditableInUi == nil {
		var ret bool
		return ret
	}
	return *o.EditableInUi
}

// GetEditableInUiOk returns a tuple with the EditableInUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetEditableInUiOk() (*bool, bool) {
	if o == nil || o.EditableInUi == nil {
		return nil, false
	}
	return o.EditableInUi, true
}

// HasEditableInUi returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasEditableInUi() bool {
	if o != nil && o.EditableInUi != nil {
		return true
	}

	return false
}

// SetEditableInUi gets a reference to the given bool and assigns it to the EditableInUi field.
func (o *BTMetadataPropertyInfo) SetEditableInUi(v bool) {
	o.EditableInUi = &v
}

// GetEnumValues returns the EnumValues field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetEnumValues() []BTMetadataEnumValueInfo {
	if o == nil || o.EnumValues == nil {
		var ret []BTMetadataEnumValueInfo
		return ret
	}
	return o.EnumValues
}

// GetEnumValuesOk returns a tuple with the EnumValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetEnumValuesOk() ([]BTMetadataEnumValueInfo, bool) {
	if o == nil || o.EnumValues == nil {
		return nil, false
	}
	return o.EnumValues, true
}

// HasEnumValues returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasEnumValues() bool {
	if o != nil && o.EnumValues != nil {
		return true
	}

	return false
}

// SetEnumValues gets a reference to the given []BTMetadataEnumValueInfo and assigns it to the EnumValues field.
func (o *BTMetadataPropertyInfo) SetEnumValues(v []BTMetadataEnumValueInfo) {
	o.EnumValues = v
}

// GetInitialValue returns the InitialValue field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetInitialValue() map[string]interface{} {
	if o == nil || o.InitialValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.InitialValue
}

// GetInitialValueOk returns a tuple with the InitialValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetInitialValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.InitialValue == nil {
		return nil, false
	}
	return o.InitialValue, true
}

// HasInitialValue returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasInitialValue() bool {
	if o != nil && o.InitialValue != nil {
		return true
	}

	return false
}

// SetInitialValue gets a reference to the given map[string]interface{} and assigns it to the InitialValue field.
func (o *BTMetadataPropertyInfo) SetInitialValue(v map[string]interface{}) {
	o.InitialValue = &v
}

// GetMultivalued returns the Multivalued field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetMultivalued() bool {
	if o == nil || o.Multivalued == nil {
		var ret bool
		return ret
	}
	return *o.Multivalued
}

// GetMultivaluedOk returns a tuple with the Multivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetMultivaluedOk() (*bool, bool) {
	if o == nil || o.Multivalued == nil {
		return nil, false
	}
	return o.Multivalued, true
}

// HasMultivalued returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasMultivalued() bool {
	if o != nil && o.Multivalued != nil {
		return true
	}

	return false
}

// SetMultivalued gets a reference to the given bool and assigns it to the Multivalued field.
func (o *BTMetadataPropertyInfo) SetMultivalued(v bool) {
	o.Multivalued = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMetadataPropertyInfo) SetName(v string) {
	o.Name = &v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetPropertyId() string {
	if o == nil || o.PropertyId == nil {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetPropertyIdOk() (*string, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *BTMetadataPropertyInfo) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetPropertySource returns the PropertySource field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetPropertySource() int32 {
	if o == nil || o.PropertySource == nil {
		var ret int32
		return ret
	}
	return *o.PropertySource
}

// GetPropertySourceOk returns a tuple with the PropertySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetPropertySourceOk() (*int32, bool) {
	if o == nil || o.PropertySource == nil {
		return nil, false
	}
	return o.PropertySource, true
}

// HasPropertySource returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasPropertySource() bool {
	if o != nil && o.PropertySource != nil {
		return true
	}

	return false
}

// SetPropertySource gets a reference to the given int32 and assigns it to the PropertySource field.
func (o *BTMetadataPropertyInfo) SetPropertySource(v int32) {
	o.PropertySource = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *BTMetadataPropertyInfo) SetRequired(v bool) {
	o.Required = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *BTMetadataPropertyInfo) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetUiHints() BTMetadataPropertyUiHintsInfo {
	if o == nil || o.UiHints == nil {
		var ret BTMetadataPropertyUiHintsInfo
		return ret
	}
	return *o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetUiHintsOk() (*BTMetadataPropertyUiHintsInfo, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given BTMetadataPropertyUiHintsInfo and assigns it to the UiHints field.
func (o *BTMetadataPropertyInfo) SetUiHints(v BTMetadataPropertyUiHintsInfo) {
	o.UiHints = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetValidator() BTMetadataPropertyValidatorInfo {
	if o == nil || o.Validator == nil {
		var ret BTMetadataPropertyValidatorInfo
		return ret
	}
	return *o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetValidatorOk() (*BTMetadataPropertyValidatorInfo, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasValidator() bool {
	if o != nil && o.Validator != nil {
		return true
	}

	return false
}

// SetValidator gets a reference to the given BTMetadataPropertyValidatorInfo and assigns it to the Validator field.
func (o *BTMetadataPropertyInfo) SetValidator(v BTMetadataPropertyValidatorInfo) {
	o.Validator = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BTMetadataPropertyInfo) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BTMetadataPropertyInfo) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *BTMetadataPropertyInfo) SetValue(v interface{}) {
	o.Value = v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *BTMetadataPropertyInfo) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyInfo) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *BTMetadataPropertyInfo) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *BTMetadataPropertyInfo) SetValueType(v string) {
	o.ValueType = &v
}

func (o BTMetadataPropertyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComputedProperty != nil {
		toSerialize["computedProperty"] = o.ComputedProperty
	}
	if o.ComputedPropertyError != nil {
		toSerialize["computedPropertyError"] = o.ComputedPropertyError
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Dirty != nil {
		toSerialize["dirty"] = o.Dirty
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	if o.EditableInUi != nil {
		toSerialize["editableInUi"] = o.EditableInUi
	}
	if o.EnumValues != nil {
		toSerialize["enumValues"] = o.EnumValues
	}
	if o.InitialValue != nil {
		toSerialize["initialValue"] = o.InitialValue
	}
	if o.Multivalued != nil {
		toSerialize["multivalued"] = o.Multivalued
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PropertyId != nil {
		toSerialize["propertyId"] = o.PropertyId
	}
	if o.PropertySource != nil {
		toSerialize["propertySource"] = o.PropertySource
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.SchemaId != nil {
		toSerialize["schemaId"] = o.SchemaId
	}
	if o.UiHints != nil {
		toSerialize["uiHints"] = o.UiHints
	}
	if o.Validator != nil {
		toSerialize["validator"] = o.Validator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataPropertyInfo struct {
	value *BTMetadataPropertyInfo
	isSet bool
}

func (v NullableBTMetadataPropertyInfo) Get() *BTMetadataPropertyInfo {
	return v.value
}

func (v *NullableBTMetadataPropertyInfo) Set(val *BTMetadataPropertyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataPropertyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataPropertyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataPropertyInfo(val *BTMetadataPropertyInfo) *NullableBTMetadataPropertyInfo {
	return &NullableBTMetadataPropertyInfo{value: val, isSet: true}
}

func (v NullableBTMetadataPropertyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataPropertyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
