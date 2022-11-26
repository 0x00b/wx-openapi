# \MaterialApi

All URIs are relative to *https://api.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaterialBatchgetMaterialPost**](MaterialApi.md#MaterialBatchgetMaterialPost) | **Post** /material/batchget_material | 素材管理 /获取素材列表
[**MaterialGetMaterialPost**](MaterialApi.md#MaterialGetMaterialPost) | **Post** /material/get_material | 素材管理 /获取永久素材



## MaterialBatchgetMaterialPost

> BatchgetMaterialResponse MaterialBatchgetMaterialPost(ctx).AccessToken(accessToken).BatchgetMaterialRequest(batchgetMaterialRequest).Execute()

素材管理 /获取素材列表

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accessToken := "accessToken_example" // string | access_token
    batchgetMaterialRequest := *openapiclient.NewBatchgetMaterialRequest() // BatchgetMaterialRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaterialApi.MaterialBatchgetMaterialPost(context.Background()).AccessToken(accessToken).BatchgetMaterialRequest(batchgetMaterialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterialApi.MaterialBatchgetMaterialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaterialBatchgetMaterialPost`: BatchgetMaterialResponse
    fmt.Fprintf(os.Stdout, "Response from `MaterialApi.MaterialBatchgetMaterialPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaterialBatchgetMaterialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string** | access_token | 
 **batchgetMaterialRequest** | [**BatchgetMaterialRequest**](BatchgetMaterialRequest.md) |  | 

### Return type

[**BatchgetMaterialResponse**](BatchgetMaterialResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaterialGetMaterialPost

> *os.File MaterialGetMaterialPost(ctx).AccessToken(accessToken).GetMaterialRequest(getMaterialRequest).Execute()

素材管理 /获取永久素材

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accessToken := "accessToken_example" // string | access_token
    getMaterialRequest := *openapiclient.NewGetMaterialRequest() // GetMaterialRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaterialApi.MaterialGetMaterialPost(context.Background()).AccessToken(accessToken).GetMaterialRequest(getMaterialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterialApi.MaterialGetMaterialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MaterialGetMaterialPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MaterialApi.MaterialGetMaterialPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaterialGetMaterialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string** | access_token | 
 **getMaterialRequest** | [**GetMaterialRequest**](GetMaterialRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

