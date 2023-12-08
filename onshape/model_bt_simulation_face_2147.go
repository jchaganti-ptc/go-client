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

// BTSimulationFace2147 struct for BTSimulationFace2147
type BTSimulationFace2147 struct {
	BtType                          *string                 `json:"btType,omitempty"`
	Compressed                      *bool                   `json:"compressed,omitempty"`
	Decompressed                    *BTEntityGeometry35     `json:"decompressed,omitempty"`
	ErrorCode                       *int32                  `json:"errorCode,omitempty"`
	EstimatedMemoryUsageInBytes     *int32                  `json:"estimatedMemoryUsageInBytes,omitempty"`
	Face                            *bool                   `json:"face,omitempty"`
	HasTessellationError            *bool                   `json:"hasTessellationError,omitempty"`
	SettingIndex                    *int32                  `json:"settingIndex,omitempty"`
	CompressedUvs                   *BTImmutableByteArray   `json:"compressedUvs,omitempty"`
	FlipComputedNormals             *bool                   `json:"flipComputedNormals,omitempty"`
	Indices                         *BTImmutableIntArray    `json:"indices,omitempty"`
	IndicesStoredAsDifferences      *bool                   `json:"indicesStoredAsDifferences,omitempty"`
	IsPlanar                        *bool                   `json:"isPlanar,omitempty"`
	MaxPrincipleCurvatureMagnitudes *BTImmutableFloatArray  `json:"maxPrincipleCurvatureMagnitudes,omitempty"`
	MinPrincipleCurvatureMagnitudes *BTImmutableFloatArray  `json:"minPrincipleCurvatureMagnitudes,omitempty"`
	Normals                         *BTImmutableFloatArray  `json:"normals,omitempty"`
	Points                          *BTImmutableFloatArray  `json:"points,omitempty"`
	SurfaceParameters               *BTImmutableDoubleArray `json:"surfaceParameters,omitempty"`
	SurfaceType                     *GBTSurfaceType         `json:"surfaceType,omitempty"`
	TextureCoordinates              *BTImmutableFloatArray  `json:"textureCoordinates,omitempty"`
	TriangleCount                   *int32                  `json:"triangleCount,omitempty"`
	Bounds                          *BTBoundingBox1052      `json:"bounds,omitempty"`
	SampleTrianglePointIndices      *BTImmutableIntArray    `json:"sampleTrianglePointIndices,omitempty"`
	TriangleNormalIndices           *BTImmutableIntArray    `json:"triangleNormalIndices,omitempty"`
	TrianglePointIndices            *BTImmutableIntArray    `json:"trianglePointIndices,omitempty"`
}

// NewBTSimulationFace2147 instantiates a new BTSimulationFace2147 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSimulationFace2147() *BTSimulationFace2147 {
	this := BTSimulationFace2147{}
	return &this
}

// NewBTSimulationFace2147WithDefaults instantiates a new BTSimulationFace2147 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSimulationFace2147WithDefaults() *BTSimulationFace2147 {
	this := BTSimulationFace2147{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSimulationFace2147) SetBtType(v string) {
	o.BtType = &v
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetCompressed() bool {
	if o == nil || o.Compressed == nil {
		var ret bool
		return ret
	}
	return *o.Compressed
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetCompressedOk() (*bool, bool) {
	if o == nil || o.Compressed == nil {
		return nil, false
	}
	return o.Compressed, true
}

// HasCompressed returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasCompressed() bool {
	if o != nil && o.Compressed != nil {
		return true
	}

	return false
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *BTSimulationFace2147) SetCompressed(v bool) {
	o.Compressed = &v
}

// GetDecompressed returns the Decompressed field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetDecompressed() BTEntityGeometry35 {
	if o == nil || o.Decompressed == nil {
		var ret BTEntityGeometry35
		return ret
	}
	return *o.Decompressed
}

// GetDecompressedOk returns a tuple with the Decompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetDecompressedOk() (*BTEntityGeometry35, bool) {
	if o == nil || o.Decompressed == nil {
		return nil, false
	}
	return o.Decompressed, true
}

// HasDecompressed returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasDecompressed() bool {
	if o != nil && o.Decompressed != nil {
		return true
	}

	return false
}

// SetDecompressed gets a reference to the given BTEntityGeometry35 and assigns it to the Decompressed field.
func (o *BTSimulationFace2147) SetDecompressed(v BTEntityGeometry35) {
	o.Decompressed = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTSimulationFace2147) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetEstimatedMemoryUsageInBytes() int32 {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedMemoryUsageInBytes
}

// GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetEstimatedMemoryUsageInBytesOk() (*int32, bool) {
	if o == nil || o.EstimatedMemoryUsageInBytes == nil {
		return nil, false
	}
	return o.EstimatedMemoryUsageInBytes, true
}

// HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasEstimatedMemoryUsageInBytes() bool {
	if o != nil && o.EstimatedMemoryUsageInBytes != nil {
		return true
	}

	return false
}

// SetEstimatedMemoryUsageInBytes gets a reference to the given int32 and assigns it to the EstimatedMemoryUsageInBytes field.
func (o *BTSimulationFace2147) SetEstimatedMemoryUsageInBytes(v int32) {
	o.EstimatedMemoryUsageInBytes = &v
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetFace() bool {
	if o == nil || o.Face == nil {
		var ret bool
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetFaceOk() (*bool, bool) {
	if o == nil || o.Face == nil {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasFace() bool {
	if o != nil && o.Face != nil {
		return true
	}

	return false
}

// SetFace gets a reference to the given bool and assigns it to the Face field.
func (o *BTSimulationFace2147) SetFace(v bool) {
	o.Face = &v
}

// GetHasTessellationError returns the HasTessellationError field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetHasTessellationError() bool {
	if o == nil || o.HasTessellationError == nil {
		var ret bool
		return ret
	}
	return *o.HasTessellationError
}

// GetHasTessellationErrorOk returns a tuple with the HasTessellationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetHasTessellationErrorOk() (*bool, bool) {
	if o == nil || o.HasTessellationError == nil {
		return nil, false
	}
	return o.HasTessellationError, true
}

// HasHasTessellationError returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasHasTessellationError() bool {
	if o != nil && o.HasTessellationError != nil {
		return true
	}

	return false
}

// SetHasTessellationError gets a reference to the given bool and assigns it to the HasTessellationError field.
func (o *BTSimulationFace2147) SetHasTessellationError(v bool) {
	o.HasTessellationError = &v
}

// GetSettingIndex returns the SettingIndex field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetSettingIndex() int32 {
	if o == nil || o.SettingIndex == nil {
		var ret int32
		return ret
	}
	return *o.SettingIndex
}

// GetSettingIndexOk returns a tuple with the SettingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetSettingIndexOk() (*int32, bool) {
	if o == nil || o.SettingIndex == nil {
		return nil, false
	}
	return o.SettingIndex, true
}

// HasSettingIndex returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasSettingIndex() bool {
	if o != nil && o.SettingIndex != nil {
		return true
	}

	return false
}

// SetSettingIndex gets a reference to the given int32 and assigns it to the SettingIndex field.
func (o *BTSimulationFace2147) SetSettingIndex(v int32) {
	o.SettingIndex = &v
}

// GetCompressedUvs returns the CompressedUvs field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetCompressedUvs() BTImmutableByteArray {
	if o == nil || o.CompressedUvs == nil {
		var ret BTImmutableByteArray
		return ret
	}
	return *o.CompressedUvs
}

// GetCompressedUvsOk returns a tuple with the CompressedUvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetCompressedUvsOk() (*BTImmutableByteArray, bool) {
	if o == nil || o.CompressedUvs == nil {
		return nil, false
	}
	return o.CompressedUvs, true
}

// HasCompressedUvs returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasCompressedUvs() bool {
	if o != nil && o.CompressedUvs != nil {
		return true
	}

	return false
}

// SetCompressedUvs gets a reference to the given BTImmutableByteArray and assigns it to the CompressedUvs field.
func (o *BTSimulationFace2147) SetCompressedUvs(v BTImmutableByteArray) {
	o.CompressedUvs = &v
}

// GetFlipComputedNormals returns the FlipComputedNormals field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetFlipComputedNormals() bool {
	if o == nil || o.FlipComputedNormals == nil {
		var ret bool
		return ret
	}
	return *o.FlipComputedNormals
}

// GetFlipComputedNormalsOk returns a tuple with the FlipComputedNormals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetFlipComputedNormalsOk() (*bool, bool) {
	if o == nil || o.FlipComputedNormals == nil {
		return nil, false
	}
	return o.FlipComputedNormals, true
}

// HasFlipComputedNormals returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasFlipComputedNormals() bool {
	if o != nil && o.FlipComputedNormals != nil {
		return true
	}

	return false
}

// SetFlipComputedNormals gets a reference to the given bool and assigns it to the FlipComputedNormals field.
func (o *BTSimulationFace2147) SetFlipComputedNormals(v bool) {
	o.FlipComputedNormals = &v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetIndices() BTImmutableIntArray {
	if o == nil || o.Indices == nil {
		var ret BTImmutableIntArray
		return ret
	}
	return *o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetIndicesOk() (*BTImmutableIntArray, bool) {
	if o == nil || o.Indices == nil {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasIndices() bool {
	if o != nil && o.Indices != nil {
		return true
	}

	return false
}

// SetIndices gets a reference to the given BTImmutableIntArray and assigns it to the Indices field.
func (o *BTSimulationFace2147) SetIndices(v BTImmutableIntArray) {
	o.Indices = &v
}

// GetIndicesStoredAsDifferences returns the IndicesStoredAsDifferences field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetIndicesStoredAsDifferences() bool {
	if o == nil || o.IndicesStoredAsDifferences == nil {
		var ret bool
		return ret
	}
	return *o.IndicesStoredAsDifferences
}

// GetIndicesStoredAsDifferencesOk returns a tuple with the IndicesStoredAsDifferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetIndicesStoredAsDifferencesOk() (*bool, bool) {
	if o == nil || o.IndicesStoredAsDifferences == nil {
		return nil, false
	}
	return o.IndicesStoredAsDifferences, true
}

// HasIndicesStoredAsDifferences returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasIndicesStoredAsDifferences() bool {
	if o != nil && o.IndicesStoredAsDifferences != nil {
		return true
	}

	return false
}

// SetIndicesStoredAsDifferences gets a reference to the given bool and assigns it to the IndicesStoredAsDifferences field.
func (o *BTSimulationFace2147) SetIndicesStoredAsDifferences(v bool) {
	o.IndicesStoredAsDifferences = &v
}

// GetIsPlanar returns the IsPlanar field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetIsPlanar() bool {
	if o == nil || o.IsPlanar == nil {
		var ret bool
		return ret
	}
	return *o.IsPlanar
}

// GetIsPlanarOk returns a tuple with the IsPlanar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetIsPlanarOk() (*bool, bool) {
	if o == nil || o.IsPlanar == nil {
		return nil, false
	}
	return o.IsPlanar, true
}

// HasIsPlanar returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasIsPlanar() bool {
	if o != nil && o.IsPlanar != nil {
		return true
	}

	return false
}

// SetIsPlanar gets a reference to the given bool and assigns it to the IsPlanar field.
func (o *BTSimulationFace2147) SetIsPlanar(v bool) {
	o.IsPlanar = &v
}

// GetMaxPrincipleCurvatureMagnitudes returns the MaxPrincipleCurvatureMagnitudes field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetMaxPrincipleCurvatureMagnitudes() BTImmutableFloatArray {
	if o == nil || o.MaxPrincipleCurvatureMagnitudes == nil {
		var ret BTImmutableFloatArray
		return ret
	}
	return *o.MaxPrincipleCurvatureMagnitudes
}

// GetMaxPrincipleCurvatureMagnitudesOk returns a tuple with the MaxPrincipleCurvatureMagnitudes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetMaxPrincipleCurvatureMagnitudesOk() (*BTImmutableFloatArray, bool) {
	if o == nil || o.MaxPrincipleCurvatureMagnitudes == nil {
		return nil, false
	}
	return o.MaxPrincipleCurvatureMagnitudes, true
}

// HasMaxPrincipleCurvatureMagnitudes returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasMaxPrincipleCurvatureMagnitudes() bool {
	if o != nil && o.MaxPrincipleCurvatureMagnitudes != nil {
		return true
	}

	return false
}

// SetMaxPrincipleCurvatureMagnitudes gets a reference to the given BTImmutableFloatArray and assigns it to the MaxPrincipleCurvatureMagnitudes field.
func (o *BTSimulationFace2147) SetMaxPrincipleCurvatureMagnitudes(v BTImmutableFloatArray) {
	o.MaxPrincipleCurvatureMagnitudes = &v
}

// GetMinPrincipleCurvatureMagnitudes returns the MinPrincipleCurvatureMagnitudes field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetMinPrincipleCurvatureMagnitudes() BTImmutableFloatArray {
	if o == nil || o.MinPrincipleCurvatureMagnitudes == nil {
		var ret BTImmutableFloatArray
		return ret
	}
	return *o.MinPrincipleCurvatureMagnitudes
}

// GetMinPrincipleCurvatureMagnitudesOk returns a tuple with the MinPrincipleCurvatureMagnitudes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetMinPrincipleCurvatureMagnitudesOk() (*BTImmutableFloatArray, bool) {
	if o == nil || o.MinPrincipleCurvatureMagnitudes == nil {
		return nil, false
	}
	return o.MinPrincipleCurvatureMagnitudes, true
}

// HasMinPrincipleCurvatureMagnitudes returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasMinPrincipleCurvatureMagnitudes() bool {
	if o != nil && o.MinPrincipleCurvatureMagnitudes != nil {
		return true
	}

	return false
}

// SetMinPrincipleCurvatureMagnitudes gets a reference to the given BTImmutableFloatArray and assigns it to the MinPrincipleCurvatureMagnitudes field.
func (o *BTSimulationFace2147) SetMinPrincipleCurvatureMagnitudes(v BTImmutableFloatArray) {
	o.MinPrincipleCurvatureMagnitudes = &v
}

// GetNormals returns the Normals field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetNormals() BTImmutableFloatArray {
	if o == nil || o.Normals == nil {
		var ret BTImmutableFloatArray
		return ret
	}
	return *o.Normals
}

// GetNormalsOk returns a tuple with the Normals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetNormalsOk() (*BTImmutableFloatArray, bool) {
	if o == nil || o.Normals == nil {
		return nil, false
	}
	return o.Normals, true
}

// HasNormals returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasNormals() bool {
	if o != nil && o.Normals != nil {
		return true
	}

	return false
}

// SetNormals gets a reference to the given BTImmutableFloatArray and assigns it to the Normals field.
func (o *BTSimulationFace2147) SetNormals(v BTImmutableFloatArray) {
	o.Normals = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetPoints() BTImmutableFloatArray {
	if o == nil || o.Points == nil {
		var ret BTImmutableFloatArray
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetPointsOk() (*BTImmutableFloatArray, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given BTImmutableFloatArray and assigns it to the Points field.
func (o *BTSimulationFace2147) SetPoints(v BTImmutableFloatArray) {
	o.Points = &v
}

// GetSurfaceParameters returns the SurfaceParameters field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetSurfaceParameters() BTImmutableDoubleArray {
	if o == nil || o.SurfaceParameters == nil {
		var ret BTImmutableDoubleArray
		return ret
	}
	return *o.SurfaceParameters
}

// GetSurfaceParametersOk returns a tuple with the SurfaceParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetSurfaceParametersOk() (*BTImmutableDoubleArray, bool) {
	if o == nil || o.SurfaceParameters == nil {
		return nil, false
	}
	return o.SurfaceParameters, true
}

// HasSurfaceParameters returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasSurfaceParameters() bool {
	if o != nil && o.SurfaceParameters != nil {
		return true
	}

	return false
}

// SetSurfaceParameters gets a reference to the given BTImmutableDoubleArray and assigns it to the SurfaceParameters field.
func (o *BTSimulationFace2147) SetSurfaceParameters(v BTImmutableDoubleArray) {
	o.SurfaceParameters = &v
}

// GetSurfaceType returns the SurfaceType field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetSurfaceType() GBTSurfaceType {
	if o == nil || o.SurfaceType == nil {
		var ret GBTSurfaceType
		return ret
	}
	return *o.SurfaceType
}

// GetSurfaceTypeOk returns a tuple with the SurfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetSurfaceTypeOk() (*GBTSurfaceType, bool) {
	if o == nil || o.SurfaceType == nil {
		return nil, false
	}
	return o.SurfaceType, true
}

// HasSurfaceType returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasSurfaceType() bool {
	if o != nil && o.SurfaceType != nil {
		return true
	}

	return false
}

// SetSurfaceType gets a reference to the given GBTSurfaceType and assigns it to the SurfaceType field.
func (o *BTSimulationFace2147) SetSurfaceType(v GBTSurfaceType) {
	o.SurfaceType = &v
}

// GetTextureCoordinates returns the TextureCoordinates field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetTextureCoordinates() BTImmutableFloatArray {
	if o == nil || o.TextureCoordinates == nil {
		var ret BTImmutableFloatArray
		return ret
	}
	return *o.TextureCoordinates
}

// GetTextureCoordinatesOk returns a tuple with the TextureCoordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetTextureCoordinatesOk() (*BTImmutableFloatArray, bool) {
	if o == nil || o.TextureCoordinates == nil {
		return nil, false
	}
	return o.TextureCoordinates, true
}

// HasTextureCoordinates returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasTextureCoordinates() bool {
	if o != nil && o.TextureCoordinates != nil {
		return true
	}

	return false
}

// SetTextureCoordinates gets a reference to the given BTImmutableFloatArray and assigns it to the TextureCoordinates field.
func (o *BTSimulationFace2147) SetTextureCoordinates(v BTImmutableFloatArray) {
	o.TextureCoordinates = &v
}

// GetTriangleCount returns the TriangleCount field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetTriangleCount() int32 {
	if o == nil || o.TriangleCount == nil {
		var ret int32
		return ret
	}
	return *o.TriangleCount
}

// GetTriangleCountOk returns a tuple with the TriangleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetTriangleCountOk() (*int32, bool) {
	if o == nil || o.TriangleCount == nil {
		return nil, false
	}
	return o.TriangleCount, true
}

// HasTriangleCount returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasTriangleCount() bool {
	if o != nil && o.TriangleCount != nil {
		return true
	}

	return false
}

// SetTriangleCount gets a reference to the given int32 and assigns it to the TriangleCount field.
func (o *BTSimulationFace2147) SetTriangleCount(v int32) {
	o.TriangleCount = &v
}

// GetBounds returns the Bounds field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetBounds() BTBoundingBox1052 {
	if o == nil || o.Bounds == nil {
		var ret BTBoundingBox1052
		return ret
	}
	return *o.Bounds
}

// GetBoundsOk returns a tuple with the Bounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetBoundsOk() (*BTBoundingBox1052, bool) {
	if o == nil || o.Bounds == nil {
		return nil, false
	}
	return o.Bounds, true
}

// HasBounds returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasBounds() bool {
	if o != nil && o.Bounds != nil {
		return true
	}

	return false
}

// SetBounds gets a reference to the given BTBoundingBox1052 and assigns it to the Bounds field.
func (o *BTSimulationFace2147) SetBounds(v BTBoundingBox1052) {
	o.Bounds = &v
}

// GetSampleTrianglePointIndices returns the SampleTrianglePointIndices field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetSampleTrianglePointIndices() BTImmutableIntArray {
	if o == nil || o.SampleTrianglePointIndices == nil {
		var ret BTImmutableIntArray
		return ret
	}
	return *o.SampleTrianglePointIndices
}

// GetSampleTrianglePointIndicesOk returns a tuple with the SampleTrianglePointIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetSampleTrianglePointIndicesOk() (*BTImmutableIntArray, bool) {
	if o == nil || o.SampleTrianglePointIndices == nil {
		return nil, false
	}
	return o.SampleTrianglePointIndices, true
}

// HasSampleTrianglePointIndices returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasSampleTrianglePointIndices() bool {
	if o != nil && o.SampleTrianglePointIndices != nil {
		return true
	}

	return false
}

// SetSampleTrianglePointIndices gets a reference to the given BTImmutableIntArray and assigns it to the SampleTrianglePointIndices field.
func (o *BTSimulationFace2147) SetSampleTrianglePointIndices(v BTImmutableIntArray) {
	o.SampleTrianglePointIndices = &v
}

// GetTriangleNormalIndices returns the TriangleNormalIndices field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetTriangleNormalIndices() BTImmutableIntArray {
	if o == nil || o.TriangleNormalIndices == nil {
		var ret BTImmutableIntArray
		return ret
	}
	return *o.TriangleNormalIndices
}

// GetTriangleNormalIndicesOk returns a tuple with the TriangleNormalIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetTriangleNormalIndicesOk() (*BTImmutableIntArray, bool) {
	if o == nil || o.TriangleNormalIndices == nil {
		return nil, false
	}
	return o.TriangleNormalIndices, true
}

// HasTriangleNormalIndices returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasTriangleNormalIndices() bool {
	if o != nil && o.TriangleNormalIndices != nil {
		return true
	}

	return false
}

// SetTriangleNormalIndices gets a reference to the given BTImmutableIntArray and assigns it to the TriangleNormalIndices field.
func (o *BTSimulationFace2147) SetTriangleNormalIndices(v BTImmutableIntArray) {
	o.TriangleNormalIndices = &v
}

// GetTrianglePointIndices returns the TrianglePointIndices field value if set, zero value otherwise.
func (o *BTSimulationFace2147) GetTrianglePointIndices() BTImmutableIntArray {
	if o == nil || o.TrianglePointIndices == nil {
		var ret BTImmutableIntArray
		return ret
	}
	return *o.TrianglePointIndices
}

// GetTrianglePointIndicesOk returns a tuple with the TrianglePointIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSimulationFace2147) GetTrianglePointIndicesOk() (*BTImmutableIntArray, bool) {
	if o == nil || o.TrianglePointIndices == nil {
		return nil, false
	}
	return o.TrianglePointIndices, true
}

// HasTrianglePointIndices returns a boolean if a field has been set.
func (o *BTSimulationFace2147) HasTrianglePointIndices() bool {
	if o != nil && o.TrianglePointIndices != nil {
		return true
	}

	return false
}

// SetTrianglePointIndices gets a reference to the given BTImmutableIntArray and assigns it to the TrianglePointIndices field.
func (o *BTSimulationFace2147) SetTrianglePointIndices(v BTImmutableIntArray) {
	o.TrianglePointIndices = &v
}

func (o BTSimulationFace2147) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Compressed != nil {
		toSerialize["compressed"] = o.Compressed
	}
	if o.Decompressed != nil {
		toSerialize["decompressed"] = o.Decompressed
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.EstimatedMemoryUsageInBytes != nil {
		toSerialize["estimatedMemoryUsageInBytes"] = o.EstimatedMemoryUsageInBytes
	}
	if o.Face != nil {
		toSerialize["face"] = o.Face
	}
	if o.HasTessellationError != nil {
		toSerialize["hasTessellationError"] = o.HasTessellationError
	}
	if o.SettingIndex != nil {
		toSerialize["settingIndex"] = o.SettingIndex
	}
	if o.CompressedUvs != nil {
		toSerialize["compressedUvs"] = o.CompressedUvs
	}
	if o.FlipComputedNormals != nil {
		toSerialize["flipComputedNormals"] = o.FlipComputedNormals
	}
	if o.Indices != nil {
		toSerialize["indices"] = o.Indices
	}
	if o.IndicesStoredAsDifferences != nil {
		toSerialize["indicesStoredAsDifferences"] = o.IndicesStoredAsDifferences
	}
	if o.IsPlanar != nil {
		toSerialize["isPlanar"] = o.IsPlanar
	}
	if o.MaxPrincipleCurvatureMagnitudes != nil {
		toSerialize["maxPrincipleCurvatureMagnitudes"] = o.MaxPrincipleCurvatureMagnitudes
	}
	if o.MinPrincipleCurvatureMagnitudes != nil {
		toSerialize["minPrincipleCurvatureMagnitudes"] = o.MinPrincipleCurvatureMagnitudes
	}
	if o.Normals != nil {
		toSerialize["normals"] = o.Normals
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.SurfaceParameters != nil {
		toSerialize["surfaceParameters"] = o.SurfaceParameters
	}
	if o.SurfaceType != nil {
		toSerialize["surfaceType"] = o.SurfaceType
	}
	if o.TextureCoordinates != nil {
		toSerialize["textureCoordinates"] = o.TextureCoordinates
	}
	if o.TriangleCount != nil {
		toSerialize["triangleCount"] = o.TriangleCount
	}
	if o.Bounds != nil {
		toSerialize["bounds"] = o.Bounds
	}
	if o.SampleTrianglePointIndices != nil {
		toSerialize["sampleTrianglePointIndices"] = o.SampleTrianglePointIndices
	}
	if o.TriangleNormalIndices != nil {
		toSerialize["triangleNormalIndices"] = o.TriangleNormalIndices
	}
	if o.TrianglePointIndices != nil {
		toSerialize["trianglePointIndices"] = o.TrianglePointIndices
	}
	return json.Marshal(toSerialize)
}

type NullableBTSimulationFace2147 struct {
	value *BTSimulationFace2147
	isSet bool
}

func (v NullableBTSimulationFace2147) Get() *BTSimulationFace2147 {
	return v.value
}

func (v *NullableBTSimulationFace2147) Set(val *BTSimulationFace2147) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSimulationFace2147) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSimulationFace2147) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSimulationFace2147(val *BTSimulationFace2147) *NullableBTSimulationFace2147 {
	return &NullableBTSimulationFace2147{value: val, isSet: true}
}

func (v NullableBTSimulationFace2147) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSimulationFace2147) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
