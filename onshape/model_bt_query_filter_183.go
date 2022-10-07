/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6797-6824f8df9946
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTQueryFilter183 - struct for BTQueryFilter183
type BTQueryFilter183 struct {
	implBTQueryFilter183 interface{}
}

// BTOccurrenceFilter166AsBTQueryFilter183 is a convenience function that returns BTOccurrenceFilter166 wrapped in BTQueryFilter183
func (o *BTOccurrenceFilter166) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTAllowMeshGeometryFilter1026AsBTQueryFilter183 is a convenience function that returns BTAllowMeshGeometryFilter1026 wrapped in BTQueryFilter183
func (o *BTAllowMeshGeometryFilter1026) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTFeatureFilter127AsBTQueryFilter183 is a convenience function that returns BTFeatureFilter127 wrapped in BTQueryFilter183
func (o *BTFeatureFilter127) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTAllowFlattenedGeometryFilter2140AsBTQueryFilter183 is a convenience function that returns BTAllowFlattenedGeometryFilter2140 wrapped in BTQueryFilter183
func (o *BTAllowFlattenedGeometryFilter2140) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTSketchSolveStatusFilter3657AsBTQueryFilter183 is a convenience function that returns BTSketchSolveStatusFilter3657 wrapped in BTQueryFilter183
func (o *BTSketchSolveStatusFilter3657) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTActiveSheetMetalFilter2944AsBTQueryFilter183 is a convenience function that returns BTActiveSheetMetalFilter2944 wrapped in BTQueryFilter183
func (o *BTActiveSheetMetalFilter2944) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTBodyTypeFilter112AsBTQueryFilter183 is a convenience function that returns BTBodyTypeFilter112 wrapped in BTQueryFilter183
func (o *BTBodyTypeFilter112) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTEntityTypeFilter124AsBTQueryFilter183 is a convenience function that returns BTEntityTypeFilter124 wrapped in BTQueryFilter183
func (o *BTEntityTypeFilter124) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTSMDefinitionEntityTypeFilter1651AsBTQueryFilter183 is a convenience function that returns BTSMDefinitionEntityTypeFilter1651 wrapped in BTQueryFilter183
func (o *BTSMDefinitionEntityTypeFilter1651) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTSplineHandleFilter2971AsBTQueryFilter183 is a convenience function that returns BTSplineHandleFilter2971 wrapped in BTQueryFilter183
func (o *BTSplineHandleFilter2971) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTEdgeTopologyFilter122AsBTQueryFilter183 is a convenience function that returns BTEdgeTopologyFilter122 wrapped in BTQueryFilter183
func (o *BTEdgeTopologyFilter122) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTFeatureTypeFilter962AsBTQueryFilter183 is a convenience function that returns BTFeatureTypeFilter962 wrapped in BTQueryFilter183
func (o *BTFeatureTypeFilter962) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTAndFilter110AsBTQueryFilter183 is a convenience function that returns BTAndFilter110 wrapped in BTQueryFilter183
func (o *BTAndFilter110) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTConstructionObjectFilter113AsBTQueryFilter183 is a convenience function that returns BTConstructionObjectFilter113 wrapped in BTQueryFilter183
func (o *BTConstructionObjectFilter113) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTAllowedMateTypeFilter1511AsBTQueryFilter183 is a convenience function that returns BTAllowedMateTypeFilter1511 wrapped in BTQueryFilter183
func (o *BTAllowedMateTypeFilter1511) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTPlaneOrientationFilter1700AsBTQueryFilter183 is a convenience function that returns BTPlaneOrientationFilter1700 wrapped in BTQueryFilter183
func (o *BTPlaneOrientationFilter1700) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTClosedCurveFilter1206AsBTQueryFilter183 is a convenience function that returns BTClosedCurveFilter1206 wrapped in BTQueryFilter183
func (o *BTClosedCurveFilter1206) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTFlatSheetMetalFilter3018AsBTQueryFilter183 is a convenience function that returns BTFlatSheetMetalFilter3018 wrapped in BTQueryFilter183
func (o *BTFlatSheetMetalFilter3018) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTMateConnectorFilter163AsBTQueryFilter183 is a convenience function that returns BTMateConnectorFilter163 wrapped in BTQueryFilter183
func (o *BTMateConnectorFilter163) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTOrFilter167AsBTQueryFilter183 is a convenience function that returns BTOrFilter167 wrapped in BTQueryFilter183
func (o *BTOrFilter167) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTSplineControlPolygonFilter1130AsBTQueryFilter183 is a convenience function that returns BTSplineControlPolygonFilter1130 wrapped in BTQueryFilter183
func (o *BTSplineControlPolygonFilter1130) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTImageFilter853AsBTQueryFilter183 is a convenience function that returns BTImageFilter853 wrapped in BTQueryFilter183
func (o *BTImageFilter853) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTMateFilter162AsBTQueryFilter183 is a convenience function that returns BTMateFilter162 wrapped in BTQueryFilter183
func (o *BTMateFilter162) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTModifiableEntityOnlyFilter1593AsBTQueryFilter183 is a convenience function that returns BTModifiableEntityOnlyFilter1593 wrapped in BTQueryFilter183
func (o *BTModifiableEntityOnlyFilter1593) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTNotFilter165AsBTQueryFilter183 is a convenience function that returns BTNotFilter165 wrapped in BTQueryFilter183
func (o *BTNotFilter165) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTTextStrokeFilter461AsBTQueryFilter183 is a convenience function that returns BTTextStrokeFilter461 wrapped in BTQueryFilter183
func (o *BTTextStrokeFilter461) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTAllowEdgePointFilter2371AsBTQueryFilter183 is a convenience function that returns BTAllowEdgePointFilter2371 wrapped in BTQueryFilter183
func (o *BTAllowEdgePointFilter2371) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTSketchObjectFilter184AsBTQueryFilter183 is a convenience function that returns BTSketchObjectFilter184 wrapped in BTQueryFilter183
func (o *BTSketchObjectFilter184) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTTextObjectFilter1515AsBTQueryFilter183 is a convenience function that returns BTTextObjectFilter1515 wrapped in BTQueryFilter183
func (o *BTTextObjectFilter1515) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// BTGeometryFilter130AsBTQueryFilter183 is a convenience function that returns BTGeometryFilter130 wrapped in BTQueryFilter183
func (o *BTGeometryFilter130) AsBTQueryFilter183() *BTQueryFilter183 {
	return &BTQueryFilter183{o}
}

// NewBTQueryFilter183 instantiates a new BTQueryFilter183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTQueryFilter183() *BTQueryFilter183 {
	this := BTQueryFilter183{Newbase_BTQueryFilter183()}
	return &this
}

// NewBTQueryFilter183WithDefaults instantiates a new BTQueryFilter183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTQueryFilter183WithDefaults() *BTQueryFilter183 {
	this := BTQueryFilter183{Newbase_BTQueryFilter183WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTQueryFilter183) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTQueryFilter183) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTQueryFilter183) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTQueryFilter183) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTQueryFilter183) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTActiveSheetMetalFilter-2944'
	if jsonDict["btType"] == "BTActiveSheetMetalFilter-2944" {
		// try to unmarshal JSON data into BTActiveSheetMetalFilter2944
		var qr *BTActiveSheetMetalFilter2944
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTActiveSheetMetalFilter2944: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAllowEdgePointFilter-2371'
	if jsonDict["btType"] == "BTAllowEdgePointFilter-2371" {
		// try to unmarshal JSON data into BTAllowEdgePointFilter2371
		var qr *BTAllowEdgePointFilter2371
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTAllowEdgePointFilter2371: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAllowFlattenedGeometryFilter-2140'
	if jsonDict["btType"] == "BTAllowFlattenedGeometryFilter-2140" {
		// try to unmarshal JSON data into BTAllowFlattenedGeometryFilter2140
		var qr *BTAllowFlattenedGeometryFilter2140
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTAllowFlattenedGeometryFilter2140: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAllowMeshGeometryFilter-1026'
	if jsonDict["btType"] == "BTAllowMeshGeometryFilter-1026" {
		// try to unmarshal JSON data into BTAllowMeshGeometryFilter1026
		var qr *BTAllowMeshGeometryFilter1026
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTAllowMeshGeometryFilter1026: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAllowedMateTypeFilter-1511'
	if jsonDict["btType"] == "BTAllowedMateTypeFilter-1511" {
		// try to unmarshal JSON data into BTAllowedMateTypeFilter1511
		var qr *BTAllowedMateTypeFilter1511
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTAllowedMateTypeFilter1511: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTAndFilter-110'
	if jsonDict["btType"] == "BTAndFilter-110" {
		// try to unmarshal JSON data into BTAndFilter110
		var qr *BTAndFilter110
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTAndFilter110: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTBodyTypeFilter-112'
	if jsonDict["btType"] == "BTBodyTypeFilter-112" {
		// try to unmarshal JSON data into BTBodyTypeFilter112
		var qr *BTBodyTypeFilter112
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTBodyTypeFilter112: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTClosedCurveFilter-1206'
	if jsonDict["btType"] == "BTClosedCurveFilter-1206" {
		// try to unmarshal JSON data into BTClosedCurveFilter1206
		var qr *BTClosedCurveFilter1206
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTClosedCurveFilter1206: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConstructionObjectFilter-113'
	if jsonDict["btType"] == "BTConstructionObjectFilter-113" {
		// try to unmarshal JSON data into BTConstructionObjectFilter113
		var qr *BTConstructionObjectFilter113
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTConstructionObjectFilter113: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEdgeTopologyFilter-122'
	if jsonDict["btType"] == "BTEdgeTopologyFilter-122" {
		// try to unmarshal JSON data into BTEdgeTopologyFilter122
		var qr *BTEdgeTopologyFilter122
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTEdgeTopologyFilter122: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTEntityTypeFilter-124'
	if jsonDict["btType"] == "BTEntityTypeFilter-124" {
		// try to unmarshal JSON data into BTEntityTypeFilter124
		var qr *BTEntityTypeFilter124
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTEntityTypeFilter124: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureFilter-127'
	if jsonDict["btType"] == "BTFeatureFilter-127" {
		// try to unmarshal JSON data into BTFeatureFilter127
		var qr *BTFeatureFilter127
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTFeatureFilter127: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFeatureTypeFilter-962'
	if jsonDict["btType"] == "BTFeatureTypeFilter-962" {
		// try to unmarshal JSON data into BTFeatureTypeFilter962
		var qr *BTFeatureTypeFilter962
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTFeatureTypeFilter962: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFlatSheetMetalFilter-3018'
	if jsonDict["btType"] == "BTFlatSheetMetalFilter-3018" {
		// try to unmarshal JSON data into BTFlatSheetMetalFilter3018
		var qr *BTFlatSheetMetalFilter3018
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTFlatSheetMetalFilter3018: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTGeometryFilter-130'
	if jsonDict["btType"] == "BTGeometryFilter-130" {
		// try to unmarshal JSON data into BTGeometryFilter130
		var qr *BTGeometryFilter130
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTGeometryFilter130: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTImageFilter-853'
	if jsonDict["btType"] == "BTImageFilter-853" {
		// try to unmarshal JSON data into BTImageFilter853
		var qr *BTImageFilter853
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTImageFilter853: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMateConnectorFilter-163'
	if jsonDict["btType"] == "BTMateConnectorFilter-163" {
		// try to unmarshal JSON data into BTMateConnectorFilter163
		var qr *BTMateConnectorFilter163
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTMateConnectorFilter163: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMateFilter-162'
	if jsonDict["btType"] == "BTMateFilter-162" {
		// try to unmarshal JSON data into BTMateFilter162
		var qr *BTMateFilter162
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTMateFilter162: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTModifiableEntityOnlyFilter-1593'
	if jsonDict["btType"] == "BTModifiableEntityOnlyFilter-1593" {
		// try to unmarshal JSON data into BTModifiableEntityOnlyFilter1593
		var qr *BTModifiableEntityOnlyFilter1593
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTModifiableEntityOnlyFilter1593: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTNotFilter-165'
	if jsonDict["btType"] == "BTNotFilter-165" {
		// try to unmarshal JSON data into BTNotFilter165
		var qr *BTNotFilter165
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTNotFilter165: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTOccurrenceFilter-166'
	if jsonDict["btType"] == "BTOccurrenceFilter-166" {
		// try to unmarshal JSON data into BTOccurrenceFilter166
		var qr *BTOccurrenceFilter166
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTOccurrenceFilter166: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTOrFilter-167'
	if jsonDict["btType"] == "BTOrFilter-167" {
		// try to unmarshal JSON data into BTOrFilter167
		var qr *BTOrFilter167
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTOrFilter167: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPlaneOrientationFilter-1700'
	if jsonDict["btType"] == "BTPlaneOrientationFilter-1700" {
		// try to unmarshal JSON data into BTPlaneOrientationFilter1700
		var qr *BTPlaneOrientationFilter1700
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTPlaneOrientationFilter1700: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSMDefinitionEntityTypeFilter-1651'
	if jsonDict["btType"] == "BTSMDefinitionEntityTypeFilter-1651" {
		// try to unmarshal JSON data into BTSMDefinitionEntityTypeFilter1651
		var qr *BTSMDefinitionEntityTypeFilter1651
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTSMDefinitionEntityTypeFilter1651: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchObjectFilter-184'
	if jsonDict["btType"] == "BTSketchObjectFilter-184" {
		// try to unmarshal JSON data into BTSketchObjectFilter184
		var qr *BTSketchObjectFilter184
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTSketchObjectFilter184: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSketchSolveStatusFilter-3657'
	if jsonDict["btType"] == "BTSketchSolveStatusFilter-3657" {
		// try to unmarshal JSON data into BTSketchSolveStatusFilter3657
		var qr *BTSketchSolveStatusFilter3657
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTSketchSolveStatusFilter3657: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSplineControlPolygonFilter-1130'
	if jsonDict["btType"] == "BTSplineControlPolygonFilter-1130" {
		// try to unmarshal JSON data into BTSplineControlPolygonFilter1130
		var qr *BTSplineControlPolygonFilter1130
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTSplineControlPolygonFilter1130: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSplineHandleFilter-2971'
	if jsonDict["btType"] == "BTSplineHandleFilter-2971" {
		// try to unmarshal JSON data into BTSplineHandleFilter2971
		var qr *BTSplineHandleFilter2971
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTSplineHandleFilter2971: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTextObjectFilter-1515'
	if jsonDict["btType"] == "BTTextObjectFilter-1515" {
		// try to unmarshal JSON data into BTTextObjectFilter1515
		var qr *BTTextObjectFilter1515
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTTextObjectFilter1515: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTTextStrokeFilter-461'
	if jsonDict["btType"] == "BTTextStrokeFilter-461" {
		// try to unmarshal JSON data into BTTextStrokeFilter461
		var qr *BTTextStrokeFilter461
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTQueryFilter183 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTQueryFilter183 = nil
			return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as BTTextStrokeFilter461: %s", err.Error())
		}
	}

	var qtx *base_BTQueryFilter183
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTQueryFilter183 = qtx
		return nil // data stored in dst.base_BTQueryFilter183, return on the first match
	} else {
		dst.implBTQueryFilter183 = nil
		return fmt.Errorf("Failed to unmarshal BTQueryFilter183 as base_BTQueryFilter183: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTQueryFilter183) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTQueryFilter183) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTQueryFilter183
}

type NullableBTQueryFilter183 struct {
	value *BTQueryFilter183
	isSet bool
}

func (v NullableBTQueryFilter183) Get() *BTQueryFilter183 {
	return v.value
}

func (v *NullableBTQueryFilter183) Set(val *BTQueryFilter183) {
	v.value = val
	v.isSet = true
}

func (v NullableBTQueryFilter183) IsSet() bool {
	return v.isSet
}

func (v *NullableBTQueryFilter183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTQueryFilter183(val *BTQueryFilter183) *NullableBTQueryFilter183 {
	return &NullableBTQueryFilter183{value: val, isSet: true}
}

func (v NullableBTQueryFilter183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTQueryFilter183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTQueryFilter183 struct {
	BtType *string `json:"btType,omitempty"`
}

// Newbase_BTQueryFilter183 instantiates a new base_BTQueryFilter183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTQueryFilter183() *base_BTQueryFilter183 {
	this := base_BTQueryFilter183{}
	return &this
}

// Newbase_BTQueryFilter183WithDefaults instantiates a new base_BTQueryFilter183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTQueryFilter183WithDefaults() *base_BTQueryFilter183 {
	this := base_BTQueryFilter183{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTQueryFilter183) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTQueryFilter183) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTQueryFilter183) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTQueryFilter183) SetBtType(v string) {
	o.BtType = &v
}

func (o base_BTQueryFilter183) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}
