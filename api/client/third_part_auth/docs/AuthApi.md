# \AuthApi

All URIs are relative to *https://api.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentApiAuthorizerTokenPost**](AuthApi.md#ComponentApiAuthorizerTokenPost) | **Post** /component/api_authorizer_token | 获取/刷新接口调用令牌
[**ComponentApiComponentTokenPost**](AuthApi.md#ComponentApiComponentTokenPost) | **Post** /component/api_component_token | 令牌
[**ComponentApiCreatePreauthcodePost**](AuthApi.md#ComponentApiCreatePreauthcodePost) | **Post** /component/api_create_preauthcode | 预授权码
[**ComponentApiGetAuthorizerInfoPost**](AuthApi.md#ComponentApiGetAuthorizerInfoPost) | **Post** /component/api_get_authorizer_info | 获取授权方的帐号基本信息
[**ComponentApiQueryAuthPost**](AuthApi.md#ComponentApiQueryAuthPost) | **Post** /component/api_query_auth | 使用授权码获取授权信息
[**ComponentApiStartPushTicketPost**](AuthApi.md#ComponentApiStartPushTicketPost) | **Post** /component/api_start_push_ticket | 启动ticket推送服务



## ComponentApiAuthorizerTokenPost

> AuthorizerTokenResponse ComponentApiAuthorizerTokenPost(ctx).ComponentAccessToken(componentAccessToken).AuthorizerTokenRequest(authorizerTokenRequest).Execute()

获取/刷新接口调用令牌

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
    componentAccessToken := "componentAccessToken_example" // string | 第三方平台component_access_token，不是authorizer_access_token
    authorizerTokenRequest := *openapiclient.NewAuthorizerTokenRequest("ComponentAppid_example", "AuthorizerAppid_example", "AuthorizerRefreshToken_example") // AuthorizerTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ComponentApiAuthorizerTokenPost(context.Background()).ComponentAccessToken(componentAccessToken).AuthorizerTokenRequest(authorizerTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ComponentApiAuthorizerTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentApiAuthorizerTokenPost`: AuthorizerTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ComponentApiAuthorizerTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentApiAuthorizerTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentAccessToken** | **string** | 第三方平台component_access_token，不是authorizer_access_token | 
 **authorizerTokenRequest** | [**AuthorizerTokenRequest**](AuthorizerTokenRequest.md) |  | 

### Return type

[**AuthorizerTokenResponse**](AuthorizerTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentApiComponentTokenPost

> ComponentTokenResponse ComponentApiComponentTokenPost(ctx).ComponentTokenRequest(componentTokenRequest).Execute()

令牌

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
    componentTokenRequest := *openapiclient.NewComponentTokenRequest("ComponentAppid_example", "ComponentAppsecret_example", "ComponentVerifyTicket_example") // ComponentTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ComponentApiComponentTokenPost(context.Background()).ComponentTokenRequest(componentTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ComponentApiComponentTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentApiComponentTokenPost`: ComponentTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ComponentApiComponentTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentApiComponentTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentTokenRequest** | [**ComponentTokenRequest**](ComponentTokenRequest.md) |  | 

### Return type

[**ComponentTokenResponse**](ComponentTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentApiCreatePreauthcodePost

> CreatePreauthcodeResponse ComponentApiCreatePreauthcodePost(ctx).ComponentAccessToken(componentAccessToken).CreatePreauthcodeRequest(createPreauthcodeRequest).Execute()

预授权码

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
    componentAccessToken := "componentAccessToken_example" // string | 第三方平台component_access_token，不是authorizer_access_token
    createPreauthcodeRequest := *openapiclient.NewCreatePreauthcodeRequest("ComponentAppid_example") // CreatePreauthcodeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ComponentApiCreatePreauthcodePost(context.Background()).ComponentAccessToken(componentAccessToken).CreatePreauthcodeRequest(createPreauthcodeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ComponentApiCreatePreauthcodePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentApiCreatePreauthcodePost`: CreatePreauthcodeResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ComponentApiCreatePreauthcodePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentApiCreatePreauthcodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentAccessToken** | **string** | 第三方平台component_access_token，不是authorizer_access_token | 
 **createPreauthcodeRequest** | [**CreatePreauthcodeRequest**](CreatePreauthcodeRequest.md) |  | 

### Return type

[**CreatePreauthcodeResponse**](CreatePreauthcodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentApiGetAuthorizerInfoPost

> GetAuthorizerInfoResponse ComponentApiGetAuthorizerInfoPost(ctx).ComponentAccessToken(componentAccessToken).GetAuthorizerInfoRequest(getAuthorizerInfoRequest).Execute()

获取授权方的帐号基本信息

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
    componentAccessToken := "componentAccessToken_example" // string | 第三方平台component_access_token，不是authorizer_access_token
    getAuthorizerInfoRequest := *openapiclient.NewGetAuthorizerInfoRequest("ComponentAppid_example", "AuthorizerAppid_example") // GetAuthorizerInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ComponentApiGetAuthorizerInfoPost(context.Background()).ComponentAccessToken(componentAccessToken).GetAuthorizerInfoRequest(getAuthorizerInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ComponentApiGetAuthorizerInfoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentApiGetAuthorizerInfoPost`: GetAuthorizerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ComponentApiGetAuthorizerInfoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentApiGetAuthorizerInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentAccessToken** | **string** | 第三方平台component_access_token，不是authorizer_access_token | 
 **getAuthorizerInfoRequest** | [**GetAuthorizerInfoRequest**](GetAuthorizerInfoRequest.md) |  | 

### Return type

[**GetAuthorizerInfoResponse**](GetAuthorizerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentApiQueryAuthPost

> QueryAuthResponse ComponentApiQueryAuthPost(ctx).ComponentAccessToken(componentAccessToken).QueryAuthRequest(queryAuthRequest).Execute()

使用授权码获取授权信息

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
    componentAccessToken := "componentAccessToken_example" // string | 第三方平台component_access_token，不是authorizer_access_token
    queryAuthRequest := *openapiclient.NewQueryAuthRequest("ComponentAppid_example", "AuthorizationCode_example") // QueryAuthRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ComponentApiQueryAuthPost(context.Background()).ComponentAccessToken(componentAccessToken).QueryAuthRequest(queryAuthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ComponentApiQueryAuthPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentApiQueryAuthPost`: QueryAuthResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ComponentApiQueryAuthPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentApiQueryAuthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentAccessToken** | **string** | 第三方平台component_access_token，不是authorizer_access_token | 
 **queryAuthRequest** | [**QueryAuthRequest**](QueryAuthRequest.md) |  | 

### Return type

[**QueryAuthResponse**](QueryAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentApiStartPushTicketPost

> StartPushTicketResponse ComponentApiStartPushTicketPost(ctx).StartPushTicketRequest(startPushTicketRequest).Execute()

启动ticket推送服务

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
    startPushTicketRequest := *openapiclient.NewStartPushTicketRequest("ComponentAppid_example", "ComponentSecret_example") // StartPushTicketRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ComponentApiStartPushTicketPost(context.Background()).StartPushTicketRequest(startPushTicketRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ComponentApiStartPushTicketPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentApiStartPushTicketPost`: StartPushTicketResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ComponentApiStartPushTicketPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComponentApiStartPushTicketPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startPushTicketRequest** | [**StartPushTicketRequest**](StartPushTicketRequest.md) |  | 

### Return type

[**StartPushTicketResponse**](StartPushTicketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

