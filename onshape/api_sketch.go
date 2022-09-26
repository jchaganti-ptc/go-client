/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6621-03f431879e4b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// SketchApiService SketchApi service
type SketchApiService service

type ApiGetSketchBoundingBoxesRequest struct {
	ctx            context.Context
	ApiService     *SketchApiService
	did            string
	wvm            string
	wvmid          string
	eid            string
	sid            string
	configuration  *string
	linkDocumentId *string
}

func (r ApiGetSketchBoundingBoxesRequest) Configuration(configuration string) ApiGetSketchBoundingBoxesRequest {
	r.configuration = &configuration
	return r
}

func (r ApiGetSketchBoundingBoxesRequest) LinkDocumentId(linkDocumentId string) ApiGetSketchBoundingBoxesRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetSketchBoundingBoxesRequest) Execute() (*BTBoundingBoxInfo, *http.Response, error) {
	return r.ApiService.GetSketchBoundingBoxesExecute(r)
}

/*
GetSketchBoundingBoxes Retrieve sketch bounding boxes by document ID, workspace or version or microversion ID, tab ID, and sketch ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @param sid
 @return ApiGetSketchBoundingBoxesRequest
*/
func (a *SketchApiService) GetSketchBoundingBoxes(ctx context.Context, did string, wvm string, wvmid string, eid string, sid string) ApiGetSketchBoundingBoxesRequest {
	return ApiGetSketchBoundingBoxesRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
		sid:        sid,
	}
}

// Execute executes the request
//  @return BTBoundingBoxInfo
func (a *SketchApiService) GetSketchBoundingBoxesExecute(r ApiGetSketchBoundingBoxesRequest) (*BTBoundingBoxInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTBoundingBoxInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SketchApiService.GetSketchBoundingBoxes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches/{sid}/boundingboxes"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sid"+"}", url.PathEscape(parameterToString(r.sid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.configuration != nil {
		localVarQueryParams.Add("configuration", parameterToString(*r.configuration, ""))
	}
	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTBoundingBoxInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSketchInfoRequest struct {
	ctx             context.Context
	ApiService      *SketchApiService
	did             string
	wvm             string
	wvmid           string
	eid             string
	configuration   *string
	sketchId        *[]string
	output3D        *bool
	curvePoints     *bool
	includeGeometry *bool
	linkDocumentId  *string
}

func (r ApiGetSketchInfoRequest) Configuration(configuration string) ApiGetSketchInfoRequest {
	r.configuration = &configuration
	return r
}

func (r ApiGetSketchInfoRequest) SketchId(sketchId []string) ApiGetSketchInfoRequest {
	r.sketchId = &sketchId
	return r
}

func (r ApiGetSketchInfoRequest) Output3D(output3D bool) ApiGetSketchInfoRequest {
	r.output3D = &output3D
	return r
}

func (r ApiGetSketchInfoRequest) CurvePoints(curvePoints bool) ApiGetSketchInfoRequest {
	r.curvePoints = &curvePoints
	return r
}

func (r ApiGetSketchInfoRequest) IncludeGeometry(includeGeometry bool) ApiGetSketchInfoRequest {
	r.includeGeometry = &includeGeometry
	return r
}

func (r ApiGetSketchInfoRequest) LinkDocumentId(linkDocumentId string) ApiGetSketchInfoRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetSketchInfoRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetSketchInfoExecute(r)
}

/*
GetSketchInfo Retrieve sketches by document ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @return ApiGetSketchInfoRequest
*/
func (a *SketchApiService) GetSketchInfo(ctx context.Context, did string, wvm string, wvmid string, eid string) ApiGetSketchInfoRequest {
	return ApiGetSketchInfoRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SketchApiService) GetSketchInfoExecute(r ApiGetSketchInfoRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SketchApiService.GetSketchInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.configuration != nil {
		localVarQueryParams.Add("configuration", parameterToString(*r.configuration, ""))
	}
	if r.sketchId != nil {
		t := *r.sketchId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sketchId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sketchId", parameterToString(t, "multi"))
		}
	}
	if r.output3D != nil {
		localVarQueryParams.Add("output3D", parameterToString(*r.output3D, ""))
	}
	if r.curvePoints != nil {
		localVarQueryParams.Add("curvePoints", parameterToString(*r.curvePoints, ""))
	}
	if r.includeGeometry != nil {
		localVarQueryParams.Add("includeGeometry", parameterToString(*r.includeGeometry, ""))
	}
	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v map[string]interface{}
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTessellatedEntitiesRequest struct {
	ctx            context.Context
	ApiService     *SketchApiService
	did            string
	wvm            string
	wvmid          string
	eid            string
	sid            string
	configuration  *string
	entityId       *[]string
	angleTolerance *float64
	chordTolerance *float64
	linkDocumentId *string
}

func (r ApiGetTessellatedEntitiesRequest) Configuration(configuration string) ApiGetTessellatedEntitiesRequest {
	r.configuration = &configuration
	return r
}

func (r ApiGetTessellatedEntitiesRequest) EntityId(entityId []string) ApiGetTessellatedEntitiesRequest {
	r.entityId = &entityId
	return r
}

func (r ApiGetTessellatedEntitiesRequest) AngleTolerance(angleTolerance float64) ApiGetTessellatedEntitiesRequest {
	r.angleTolerance = &angleTolerance
	return r
}

func (r ApiGetTessellatedEntitiesRequest) ChordTolerance(chordTolerance float64) ApiGetTessellatedEntitiesRequest {
	r.chordTolerance = &chordTolerance
	return r
}

func (r ApiGetTessellatedEntitiesRequest) LinkDocumentId(linkDocumentId string) ApiGetTessellatedEntitiesRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetTessellatedEntitiesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetTessellatedEntitiesExecute(r)
}

/*
GetTessellatedEntities Retrieve tessellated entities of sketches by document ID, workspace or version or microversion ID, tab ID, and sketch ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wvm
 @param wvmid
 @param eid
 @param sid
 @return ApiGetTessellatedEntitiesRequest
*/
func (a *SketchApiService) GetTessellatedEntities(ctx context.Context, did string, wvm string, wvmid string, eid string, sid string) ApiGetTessellatedEntitiesRequest {
	return ApiGetTessellatedEntitiesRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
		sid:        sid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SketchApiService) GetTessellatedEntitiesExecute(r ApiGetTessellatedEntitiesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SketchApiService.GetTessellatedEntities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches/{sid}/tessellatedentities"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sid"+"}", url.PathEscape(parameterToString(r.sid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.configuration != nil {
		localVarQueryParams.Add("configuration", parameterToString(*r.configuration, ""))
	}
	if r.entityId != nil {
		t := *r.entityId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("entityId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("entityId", parameterToString(t, "multi"))
		}
	}
	if r.angleTolerance != nil {
		localVarQueryParams.Add("angleTolerance", parameterToString(*r.angleTolerance, ""))
	}
	if r.chordTolerance != nil {
		localVarQueryParams.Add("chordTolerance", parameterToString(*r.chordTolerance, ""))
	}
	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v map[string]interface{}
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
