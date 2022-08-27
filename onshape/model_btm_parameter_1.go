/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6246-994505a8b5bf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTMParameter1 - struct for BTMParameter1
type BTMParameter1 struct {
	implBTMParameter1 interface{}
}

// BTMParameterQueryList148AsBTMParameter1 is a convenience function that returns BTMParameterQueryList148 wrapped in BTMParameter1
func (o *BTMParameterQueryList148) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterFeatureList1749AsBTMParameter1 is a convenience function that returns BTMParameterFeatureList1749 wrapped in BTMParameter1
func (o *BTMParameterFeatureList1749) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterReferenceImage2014AsBTMParameter1 is a convenience function that returns BTMParameterReferenceImage2014 wrapped in BTMParameter1
func (o *BTMParameterReferenceImage2014) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterBlobReference1679AsBTMParameter1 is a convenience function that returns BTMParameterBlobReference1679 wrapped in BTMParameter1
func (o *BTMParameterBlobReference1679) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterNullableQuantity807AsBTMParameter1 is a convenience function that returns BTMParameterNullableQuantity807 wrapped in BTMParameter1
func (o *BTMParameterNullableQuantity807) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterAppearance627AsBTMParameter1 is a convenience function that returns BTMParameterAppearance627 wrapped in BTMParameter1
func (o *BTMParameterAppearance627) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterBoolean144AsBTMParameter1 is a convenience function that returns BTMParameterBoolean144 wrapped in BTMParameter1
func (o *BTMParameterBoolean144) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterEnum145AsBTMParameter1 is a convenience function that returns BTMParameterEnum145 wrapped in BTMParameter1
func (o *BTMParameterEnum145) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterDerived864AsBTMParameter1 is a convenience function that returns BTMParameterDerived864 wrapped in BTMParameter1
func (o *BTMParameterDerived864) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMDatabaseParameter2229AsBTMParameter1 is a convenience function that returns BTMDatabaseParameter2229 wrapped in BTMParameter1
func (o *BTMDatabaseParameter2229) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterArray2025AsBTMParameter1 is a convenience function that returns BTMParameterArray2025 wrapped in BTMParameter1
func (o *BTMParameterArray2025) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterStringWithTolerances4286AsBTMParameter1 is a convenience function that returns BTMParameterStringWithTolerances4286 wrapped in BTMParameter1
func (o *BTMParameterStringWithTolerances4286) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterReferenceTable917AsBTMParameter1 is a convenience function that returns BTMParameterReferenceTable917 wrapped in BTMParameter1
func (o *BTMParameterReferenceTable917) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterMaterial1388AsBTMParameter1 is a convenience function that returns BTMParameterMaterial1388 wrapped in BTMParameter1
func (o *BTMParameterMaterial1388) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterInvalid1664AsBTMParameter1 is a convenience function that returns BTMParameterInvalid1664 wrapped in BTMParameter1
func (o *BTMParameterInvalid1664) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMReadOnlyParameter3800AsBTMParameter1 is a convenience function that returns BTMReadOnlyParameter3800 wrapped in BTMParameter1
func (o *BTMReadOnlyParameter3800) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterQuantity147AsBTMParameter1 is a convenience function that returns BTMParameterQuantity147 wrapped in BTMParameter1
func (o *BTMParameterQuantity147) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterProgress3232AsBTMParameter1 is a convenience function that returns BTMParameterProgress3232 wrapped in BTMParameter1
func (o *BTMParameterProgress3232) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterReference2434AsBTMParameter1 is a convenience function that returns BTMParameterReference2434 wrapped in BTMParameter1
func (o *BTMParameterReference2434) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterConfigured2222AsBTMParameter1 is a convenience function that returns BTMParameterConfigured2222 wrapped in BTMParameter1
func (o *BTMParameterConfigured2222) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterReferenceJSON790AsBTMParameter1 is a convenience function that returns BTMParameterReferenceJSON790 wrapped in BTMParameter1
func (o *BTMParameterReferenceJSON790) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterLookupTablePath1419AsBTMParameter1 is a convenience function that returns BTMParameterLookupTablePath1419 wrapped in BTMParameter1
func (o *BTMParameterLookupTablePath1419) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterReferenceCADImport2016AsBTMParameter1 is a convenience function that returns BTMParameterReferenceCADImport2016 wrapped in BTMParameter1
func (o *BTMParameterReferenceCADImport2016) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterForeignId146AsBTMParameter1 is a convenience function that returns BTMParameterForeignId146 wrapped in BTMParameter1
func (o *BTMParameterForeignId146) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterReferenceBlob3281AsBTMParameter1 is a convenience function that returns BTMParameterReferenceBlob3281 wrapped in BTMParameter1
func (o *BTMParameterReferenceBlob3281) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterString149AsBTMParameter1 is a convenience function that returns BTMParameterString149 wrapped in BTMParameter1
func (o *BTMParameterString149) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// BTMParameterQueryWithOccurrenceList67AsBTMParameter1 is a convenience function that returns BTMParameterQueryWithOccurrenceList67 wrapped in BTMParameter1
func (o *BTMParameterQueryWithOccurrenceList67) AsBTMParameter1() *BTMParameter1 {
	return &BTMParameter1{o}
}

// NewBTMParameter1 instantiates a new BTMParameter1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameter1() *BTMParameter1 {
	this := BTMParameter1{Newbase_BTMParameter1()}
	return &this
}

// NewBTMParameter1WithDefaults instantiates a new BTMParameter1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameter1WithDefaults() *BTMParameter1 {
	this := BTMParameter1{Newbase_BTMParameter1WithDefaults()}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameter1) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameter1) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameter1) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameter1) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameter1) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameter1) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameter1) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameter1) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameter1) GetParameterId() string {
	type getResult interface {
		GetParameterId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterId()
	} else {
		var de string
		return de
	}
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameter1) GetParameterIdOk() (*string, bool) {
	type getResult interface {
		GetParameterIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetParameterIdOk()
	} else {
		return nil, false
	}
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameter1) HasParameterId() bool {
	type getResult interface {
		HasParameterId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasParameterId()
	} else {
		return false
	}
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameter1) SetParameterId(v string) {
	type getResult interface {
		SetParameterId(v string)
	}

	o.GetActualInstance().(getResult).SetParameterId(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMParameter1) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMDatabaseParameter-2229'
	if jsonDict["btType"] == "BTMDatabaseParameter-2229" {
		// try to unmarshal JSON data into BTMDatabaseParameter2229
		var qr *BTMDatabaseParameter2229
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMDatabaseParameter2229: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterAppearance-627'
	if jsonDict["btType"] == "BTMParameterAppearance-627" {
		// try to unmarshal JSON data into BTMParameterAppearance627
		var qr *BTMParameterAppearance627
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterAppearance627: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterArray-2025'
	if jsonDict["btType"] == "BTMParameterArray-2025" {
		// try to unmarshal JSON data into BTMParameterArray2025
		var qr *BTMParameterArray2025
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterArray2025: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterBlobReference-1679'
	if jsonDict["btType"] == "BTMParameterBlobReference-1679" {
		// try to unmarshal JSON data into BTMParameterBlobReference1679
		var qr *BTMParameterBlobReference1679
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterBlobReference1679: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterBoolean-144'
	if jsonDict["btType"] == "BTMParameterBoolean-144" {
		// try to unmarshal JSON data into BTMParameterBoolean144
		var qr *BTMParameterBoolean144
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterBoolean144: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterConfigured-2222'
	if jsonDict["btType"] == "BTMParameterConfigured-2222" {
		// try to unmarshal JSON data into BTMParameterConfigured2222
		var qr *BTMParameterConfigured2222
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterConfigured2222: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterDerived-864'
	if jsonDict["btType"] == "BTMParameterDerived-864" {
		// try to unmarshal JSON data into BTMParameterDerived864
		var qr *BTMParameterDerived864
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterDerived864: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterEnum-145'
	if jsonDict["btType"] == "BTMParameterEnum-145" {
		// try to unmarshal JSON data into BTMParameterEnum145
		var qr *BTMParameterEnum145
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterEnum145: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterFeatureList-1749'
	if jsonDict["btType"] == "BTMParameterFeatureList-1749" {
		// try to unmarshal JSON data into BTMParameterFeatureList1749
		var qr *BTMParameterFeatureList1749
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterFeatureList1749: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterForeignId-146'
	if jsonDict["btType"] == "BTMParameterForeignId-146" {
		// try to unmarshal JSON data into BTMParameterForeignId146
		var qr *BTMParameterForeignId146
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterForeignId146: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterInvalid-1664'
	if jsonDict["btType"] == "BTMParameterInvalid-1664" {
		// try to unmarshal JSON data into BTMParameterInvalid1664
		var qr *BTMParameterInvalid1664
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterInvalid1664: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterLookupTablePath-1419'
	if jsonDict["btType"] == "BTMParameterLookupTablePath-1419" {
		// try to unmarshal JSON data into BTMParameterLookupTablePath1419
		var qr *BTMParameterLookupTablePath1419
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterLookupTablePath1419: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterMaterial-1388'
	if jsonDict["btType"] == "BTMParameterMaterial-1388" {
		// try to unmarshal JSON data into BTMParameterMaterial1388
		var qr *BTMParameterMaterial1388
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterMaterial1388: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterNullableQuantity-807'
	if jsonDict["btType"] == "BTMParameterNullableQuantity-807" {
		// try to unmarshal JSON data into BTMParameterNullableQuantity807
		var qr *BTMParameterNullableQuantity807
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterNullableQuantity807: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterProgress-3232'
	if jsonDict["btType"] == "BTMParameterProgress-3232" {
		// try to unmarshal JSON data into BTMParameterProgress3232
		var qr *BTMParameterProgress3232
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterProgress3232: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterQuantity-147'
	if jsonDict["btType"] == "BTMParameterQuantity-147" {
		// try to unmarshal JSON data into BTMParameterQuantity147
		var qr *BTMParameterQuantity147
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterQuantity147: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterQueryList-148'
	if jsonDict["btType"] == "BTMParameterQueryList-148" {
		// try to unmarshal JSON data into BTMParameterQueryList148
		var qr *BTMParameterQueryList148
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterQueryList148: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterQueryWithOccurrenceList-67'
	if jsonDict["btType"] == "BTMParameterQueryWithOccurrenceList-67" {
		// try to unmarshal JSON data into BTMParameterQueryWithOccurrenceList67
		var qr *BTMParameterQueryWithOccurrenceList67
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterQueryWithOccurrenceList67: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReference-2434'
	if jsonDict["btType"] == "BTMParameterReference-2434" {
		// try to unmarshal JSON data into BTMParameterReference2434
		var qr *BTMParameterReference2434
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterReference2434: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceBlob-3281'
	if jsonDict["btType"] == "BTMParameterReferenceBlob-3281" {
		// try to unmarshal JSON data into BTMParameterReferenceBlob3281
		var qr *BTMParameterReferenceBlob3281
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterReferenceBlob3281: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceCADImport-2016'
	if jsonDict["btType"] == "BTMParameterReferenceCADImport-2016" {
		// try to unmarshal JSON data into BTMParameterReferenceCADImport2016
		var qr *BTMParameterReferenceCADImport2016
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterReferenceCADImport2016: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceImage-2014'
	if jsonDict["btType"] == "BTMParameterReferenceImage-2014" {
		// try to unmarshal JSON data into BTMParameterReferenceImage2014
		var qr *BTMParameterReferenceImage2014
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterReferenceImage2014: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceJSON-790'
	if jsonDict["btType"] == "BTMParameterReferenceJSON-790" {
		// try to unmarshal JSON data into BTMParameterReferenceJSON790
		var qr *BTMParameterReferenceJSON790
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterReferenceJSON790: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterReferenceTable-917'
	if jsonDict["btType"] == "BTMParameterReferenceTable-917" {
		// try to unmarshal JSON data into BTMParameterReferenceTable917
		var qr *BTMParameterReferenceTable917
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterReferenceTable917: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterString-149'
	if jsonDict["btType"] == "BTMParameterString-149" {
		// try to unmarshal JSON data into BTMParameterString149
		var qr *BTMParameterString149
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterString149: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMParameterStringWithTolerances-4286'
	if jsonDict["btType"] == "BTMParameterStringWithTolerances-4286" {
		// try to unmarshal JSON data into BTMParameterStringWithTolerances4286
		var qr *BTMParameterStringWithTolerances4286
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMParameterStringWithTolerances4286: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMReadOnlyParameter-3800'
	if jsonDict["btType"] == "BTMReadOnlyParameter-3800" {
		// try to unmarshal JSON data into BTMReadOnlyParameter3800
		var qr *BTMReadOnlyParameter3800
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMParameter1 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMParameter1 = nil
			return fmt.Errorf("Failed to unmarshal BTMParameter1 as BTMReadOnlyParameter3800: %s", err.Error())
		}
	}

	var qtx *base_BTMParameter1
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMParameter1 = qtx
		return nil // data stored in dst.base_BTMParameter1, return on the first match
	} else {
		dst.implBTMParameter1 = nil
		return fmt.Errorf("Failed to unmarshal BTMParameter1 as base_BTMParameter1: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMParameter1) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMParameter1) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMParameter1
}

type NullableBTMParameter1 struct {
	value *BTMParameter1
	isSet bool
}

func (v NullableBTMParameter1) Get() *BTMParameter1 {
	return v.value
}

func (v *NullableBTMParameter1) Set(val *BTMParameter1) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameter1) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameter1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameter1(val *BTMParameter1) *NullableBTMParameter1 {
	return &NullableBTMParameter1{value: val, isSet: true}
}

func (v NullableBTMParameter1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameter1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMParameter1 struct {
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
	ParameterId        *string `json:"parameterId,omitempty"`
}

// Newbase_BTMParameter1 instantiates a new base_BTMParameter1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMParameter1() *base_BTMParameter1 {
	this := base_BTMParameter1{}
	return &this
}

// Newbase_BTMParameter1WithDefaults instantiates a new base_BTMParameter1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMParameter1WithDefaults() *base_BTMParameter1 {
	this := base_BTMParameter1{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMParameter1) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameter1) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMParameter1) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMParameter1) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMParameter1) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameter1) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMParameter1) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMParameter1) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *base_BTMParameter1) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMParameter1) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *base_BTMParameter1) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *base_BTMParameter1) SetParameterId(v string) {
	o.ParameterId = &v
}

func (o base_BTMParameter1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	return json.Marshal(toSerialize)
}
