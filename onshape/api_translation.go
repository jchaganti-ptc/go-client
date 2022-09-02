/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6326-97b3616ccba2
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
	"strings"
)

// TranslationApiService TranslationApi service
type TranslationApiService service

type ApiCreateTranslationRequest struct {
	ctx                                  context.Context
	ApiService                           *TranslationApiService
	did                                  string
	wid                                  string
	file                                 *map[string]interface{}
	allowFaultyParts                     *bool
	createComposite                      *bool
	createDrawingIfPossible              *bool
	encodedFilename                      *string
	extractAssemblyHierarchy             *bool
	flattenAssemblies                    *bool
	formatName                           *string
	joinAdjacentSurfaces                 *bool
	locationElementId                    *string
	locationGroupId                      *string
	locationPosition                     *int32
	notifyUser                           *bool
	ownerId                              *string
	parentId                             *string
	projectId                            *string
	public                               *bool
	onePartPerDoc                        *bool
	splitAssembliesIntoMultipleDocuments *bool
	storeInDocument                      *bool
	translate                            *bool
	unit                                 *string
	uploadId                             *string
	versionString                        *string
	yAxisIsUp                            *bool
	importWithinDocument                 *bool
}

// The file to upload.
func (r ApiCreateTranslationRequest) File(file map[string]interface{}) ApiCreateTranslationRequest {
	r.file = &file
	return r
}

func (r ApiCreateTranslationRequest) AllowFaultyParts(allowFaultyParts bool) ApiCreateTranslationRequest {
	r.allowFaultyParts = &allowFaultyParts
	return r
}

func (r ApiCreateTranslationRequest) CreateComposite(createComposite bool) ApiCreateTranslationRequest {
	r.createComposite = &createComposite
	return r
}

func (r ApiCreateTranslationRequest) CreateDrawingIfPossible(createDrawingIfPossible bool) ApiCreateTranslationRequest {
	r.createDrawingIfPossible = &createDrawingIfPossible
	return r
}

func (r ApiCreateTranslationRequest) EncodedFilename(encodedFilename string) ApiCreateTranslationRequest {
	r.encodedFilename = &encodedFilename
	return r
}

func (r ApiCreateTranslationRequest) ExtractAssemblyHierarchy(extractAssemblyHierarchy bool) ApiCreateTranslationRequest {
	r.extractAssemblyHierarchy = &extractAssemblyHierarchy
	return r
}

func (r ApiCreateTranslationRequest) FlattenAssemblies(flattenAssemblies bool) ApiCreateTranslationRequest {
	r.flattenAssemblies = &flattenAssemblies
	return r
}

func (r ApiCreateTranslationRequest) FormatName(formatName string) ApiCreateTranslationRequest {
	r.formatName = &formatName
	return r
}

func (r ApiCreateTranslationRequest) JoinAdjacentSurfaces(joinAdjacentSurfaces bool) ApiCreateTranslationRequest {
	r.joinAdjacentSurfaces = &joinAdjacentSurfaces
	return r
}

func (r ApiCreateTranslationRequest) LocationElementId(locationElementId string) ApiCreateTranslationRequest {
	r.locationElementId = &locationElementId
	return r
}

func (r ApiCreateTranslationRequest) LocationGroupId(locationGroupId string) ApiCreateTranslationRequest {
	r.locationGroupId = &locationGroupId
	return r
}

func (r ApiCreateTranslationRequest) LocationPosition(locationPosition int32) ApiCreateTranslationRequest {
	r.locationPosition = &locationPosition
	return r
}

func (r ApiCreateTranslationRequest) NotifyUser(notifyUser bool) ApiCreateTranslationRequest {
	r.notifyUser = &notifyUser
	return r
}

func (r ApiCreateTranslationRequest) OwnerId(ownerId string) ApiCreateTranslationRequest {
	r.ownerId = &ownerId
	return r
}

func (r ApiCreateTranslationRequest) ParentId(parentId string) ApiCreateTranslationRequest {
	r.parentId = &parentId
	return r
}

func (r ApiCreateTranslationRequest) ProjectId(projectId string) ApiCreateTranslationRequest {
	r.projectId = &projectId
	return r
}

func (r ApiCreateTranslationRequest) Public(public bool) ApiCreateTranslationRequest {
	r.public = &public
	return r
}

func (r ApiCreateTranslationRequest) OnePartPerDoc(onePartPerDoc bool) ApiCreateTranslationRequest {
	r.onePartPerDoc = &onePartPerDoc
	return r
}

func (r ApiCreateTranslationRequest) SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments bool) ApiCreateTranslationRequest {
	r.splitAssembliesIntoMultipleDocuments = &splitAssembliesIntoMultipleDocuments
	return r
}

func (r ApiCreateTranslationRequest) StoreInDocument(storeInDocument bool) ApiCreateTranslationRequest {
	r.storeInDocument = &storeInDocument
	return r
}

func (r ApiCreateTranslationRequest) Translate(translate bool) ApiCreateTranslationRequest {
	r.translate = &translate
	return r
}

func (r ApiCreateTranslationRequest) Unit(unit string) ApiCreateTranslationRequest {
	r.unit = &unit
	return r
}

func (r ApiCreateTranslationRequest) UploadId(uploadId string) ApiCreateTranslationRequest {
	r.uploadId = &uploadId
	return r
}

func (r ApiCreateTranslationRequest) VersionString(versionString string) ApiCreateTranslationRequest {
	r.versionString = &versionString
	return r
}

func (r ApiCreateTranslationRequest) YAxisIsUp(yAxisIsUp bool) ApiCreateTranslationRequest {
	r.yAxisIsUp = &yAxisIsUp
	return r
}

func (r ApiCreateTranslationRequest) ImportWithinDocument(importWithinDocument bool) ApiCreateTranslationRequest {
	r.importWithinDocument = &importWithinDocument
	return r
}

func (r ApiCreateTranslationRequest) Execute() (*BTTranslationRequestInfo, *http.Response, error) {
	return r.ApiService.CreateTranslationExecute(r)
}

/*
CreateTranslation Upload foreign data (for example, an X_T file) into Onshape, and then translate the data to generate a part, Part Studio, Assembly, or subassembly.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wid
 @return ApiCreateTranslationRequest
*/
func (a *TranslationApiService) CreateTranslation(ctx context.Context, did string, wid string) ApiCreateTranslationRequest {
	return ApiCreateTranslationRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
	}
}

// Execute executes the request
//  @return BTTranslationRequestInfo
func (a *TranslationApiService) CreateTranslationExecute(r ApiCreateTranslationRequest) (*BTTranslationRequestInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTTranslationRequestInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TranslationApiService.CreateTranslation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/translations/d/{did}/w/{wid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.file != nil {
		paramJson, err := parameterToJson(*r.file)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("file", paramJson)
	}
	if r.allowFaultyParts != nil {
		localVarFormParams.Add("allowFaultyParts", parameterToString(*r.allowFaultyParts, ""))
	}
	if r.createComposite != nil {
		localVarFormParams.Add("createComposite", parameterToString(*r.createComposite, ""))
	}
	if r.createDrawingIfPossible != nil {
		localVarFormParams.Add("createDrawingIfPossible", parameterToString(*r.createDrawingIfPossible, ""))
	}
	if r.encodedFilename != nil {
		localVarFormParams.Add("encodedFilename", parameterToString(*r.encodedFilename, ""))
	}
	if r.extractAssemblyHierarchy != nil {
		localVarFormParams.Add("extractAssemblyHierarchy", parameterToString(*r.extractAssemblyHierarchy, ""))
	}
	if r.flattenAssemblies != nil {
		localVarFormParams.Add("flattenAssemblies", parameterToString(*r.flattenAssemblies, ""))
	}
	if r.formatName != nil {
		localVarFormParams.Add("formatName", parameterToString(*r.formatName, ""))
	}
	if r.joinAdjacentSurfaces != nil {
		localVarFormParams.Add("joinAdjacentSurfaces", parameterToString(*r.joinAdjacentSurfaces, ""))
	}
	if r.locationElementId != nil {
		localVarFormParams.Add("locationElementId", parameterToString(*r.locationElementId, ""))
	}
	if r.locationGroupId != nil {
		localVarFormParams.Add("locationGroupId", parameterToString(*r.locationGroupId, ""))
	}
	if r.locationPosition != nil {
		localVarFormParams.Add("locationPosition", parameterToString(*r.locationPosition, ""))
	}
	if r.notifyUser != nil {
		localVarFormParams.Add("notifyUser", parameterToString(*r.notifyUser, ""))
	}
	if r.ownerId != nil {
		localVarFormParams.Add("ownerId", parameterToString(*r.ownerId, ""))
	}
	if r.parentId != nil {
		localVarFormParams.Add("parentId", parameterToString(*r.parentId, ""))
	}
	if r.projectId != nil {
		localVarFormParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.public != nil {
		localVarFormParams.Add("public", parameterToString(*r.public, ""))
	}
	if r.onePartPerDoc != nil {
		localVarFormParams.Add("onePartPerDoc", parameterToString(*r.onePartPerDoc, ""))
	}
	if r.splitAssembliesIntoMultipleDocuments != nil {
		localVarFormParams.Add("splitAssembliesIntoMultipleDocuments", parameterToString(*r.splitAssembliesIntoMultipleDocuments, ""))
	}
	if r.storeInDocument != nil {
		localVarFormParams.Add("storeInDocument", parameterToString(*r.storeInDocument, ""))
	}
	if r.translate != nil {
		localVarFormParams.Add("translate", parameterToString(*r.translate, ""))
	}
	if r.unit != nil {
		localVarFormParams.Add("unit", parameterToString(*r.unit, ""))
	}
	if r.uploadId != nil {
		localVarFormParams.Add("uploadId", parameterToString(*r.uploadId, ""))
	}
	if r.versionString != nil {
		localVarFormParams.Add("versionString", parameterToString(*r.versionString, ""))
	}
	if r.yAxisIsUp != nil {
		localVarFormParams.Add("yAxisIsUp", parameterToString(*r.yAxisIsUp, ""))
	}
	if r.importWithinDocument != nil {
		localVarFormParams.Add("importWithinDocument", parameterToString(*r.importWithinDocument, ""))
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
		var v BTTranslationRequestInfo
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

type ApiDeleteTranslationRequest struct {
	ctx        context.Context
	ApiService *TranslationApiService
	tid        string
}

func (r ApiDeleteTranslationRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteTranslationExecute(r)
}

/*
DeleteTranslation Delete translation status entry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tid
 @return ApiDeleteTranslationRequest
*/
func (a *TranslationApiService) DeleteTranslation(ctx context.Context, tid string) ApiDeleteTranslationRequest {
	return ApiDeleteTranslationRequest{
		ApiService: a,
		ctx:        ctx,
		tid:        tid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *TranslationApiService) DeleteTranslationExecute(r ApiDeleteTranslationRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TranslationApiService.DeleteTranslation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/translations/{tid}"
	localVarPath = strings.Replace(localVarPath, "{"+"tid"+"}", url.PathEscape(parameterToString(r.tid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetAllTranslatorFormatsRequest struct {
	ctx        context.Context
	ApiService *TranslationApiService
}

func (r ApiGetAllTranslatorFormatsRequest) Execute() ([]BTModelFormatFullInfo, *http.Response, error) {
	return r.ApiService.GetAllTranslatorFormatsExecute(r)
}

/*
GetAllTranslatorFormats Retrieve a list of translation formats that can work for this translation. Some are valid only as an input format and cannot be used as an output format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllTranslatorFormatsRequest
*/
func (a *TranslationApiService) GetAllTranslatorFormats(ctx context.Context) ApiGetAllTranslatorFormatsRequest {
	return ApiGetAllTranslatorFormatsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []BTModelFormatFullInfo
func (a *TranslationApiService) GetAllTranslatorFormatsExecute(r ApiGetAllTranslatorFormatsRequest) ([]BTModelFormatFullInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTModelFormatFullInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TranslationApiService.GetAllTranslatorFormats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/translations/translationformats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		var v []BTModelFormatFullInfo
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

type ApiGetDocumentTranslationsRequest struct {
	ctx        context.Context
	ApiService *TranslationApiService
	did        string
	offset     *int32
	limit      *int32
}

func (r ApiGetDocumentTranslationsRequest) Offset(offset int32) ApiGetDocumentTranslationsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetDocumentTranslationsRequest) Limit(limit int32) ApiGetDocumentTranslationsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetDocumentTranslationsRequest) Execute() (*BTListResponseBTTranslationRequestInfo, *http.Response, error) {
	return r.ApiService.GetDocumentTranslationsExecute(r)
}

/*
GetDocumentTranslations Request an array of translations that were made against this document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @return ApiGetDocumentTranslationsRequest
*/
func (a *TranslationApiService) GetDocumentTranslations(ctx context.Context, did string) ApiGetDocumentTranslationsRequest {
	return ApiGetDocumentTranslationsRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
	}
}

// Execute executes the request
//  @return BTListResponseBTTranslationRequestInfo
func (a *TranslationApiService) GetDocumentTranslationsExecute(r ApiGetDocumentTranslationsRequest) (*BTListResponseBTTranslationRequestInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTTranslationRequestInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TranslationApiService.GetDocumentTranslations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/translations/d/{did}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
		var v BTListResponseBTTranslationRequestInfo
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

type ApiGetTranslationRequest struct {
	ctx        context.Context
	ApiService *TranslationApiService
	tid        string
}

func (r ApiGetTranslationRequest) Execute() (*BTTranslationRequestInfo, *http.Response, error) {
	return r.ApiService.GetTranslationExecute(r)
}

/*
GetTranslation Request information on an in-progress or completed translation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tid
 @return ApiGetTranslationRequest
*/
func (a *TranslationApiService) GetTranslation(ctx context.Context, tid string) ApiGetTranslationRequest {
	return ApiGetTranslationRequest{
		ApiService: a,
		ctx:        ctx,
		tid:        tid,
	}
}

// Execute executes the request
//  @return BTTranslationRequestInfo
func (a *TranslationApiService) GetTranslationExecute(r ApiGetTranslationRequest) (*BTTranslationRequestInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTTranslationRequestInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TranslationApiService.GetTranslation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/translations/{tid}"
	localVarPath = strings.Replace(localVarPath, "{"+"tid"+"}", url.PathEscape(parameterToString(r.tid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		var v BTTranslationRequestInfo
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
