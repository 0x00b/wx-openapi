# \MiniProgramAuthApi

All URIs are relative to *https://api.weixin.qq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CgiBinTokenGet**](MiniProgramAuthApi.md#CgiBinTokenGet) | **Get** /cgi-bin/token | 获取小程序全局唯一后台接口调用凭据（access_token）。
[**SnsJscode2sessionGet**](MiniProgramAuthApi.md#SnsJscode2sessionGet) | **Get** /sns/jscode2session | 登录凭证校验。
[**WxaGetwxacodeunlimitGet**](MiniProgramAuthApi.md#WxaGetwxacodeunlimitGet) | **Get** /wxa/getwxacodeunlimit | 获取小程序全局唯一后台接口调用凭据（access_token）。



## CgiBinTokenGet

> TokenRsp CgiBinTokenGet(ctx).Appid(appid).Secret(secret).GrantType(grantType).Execute()

获取小程序全局唯一后台接口调用凭据（access_token）。

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
    appid := "appid_example" // string | 小程序 appId
    secret := "secret_example" // string | 小程序 secret
    grantType := "grantType_example" // string | 授权类型，此处只需填写 client_credential

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiniProgramAuthApi.CgiBinTokenGet(context.Background()).Appid(appid).Secret(secret).GrantType(grantType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiniProgramAuthApi.CgiBinTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CgiBinTokenGet`: TokenRsp
    fmt.Fprintf(os.Stdout, "Response from `MiniProgramAuthApi.CgiBinTokenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCgiBinTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appid** | **string** | 小程序 appId | 
 **secret** | **string** | 小程序 secret | 
 **grantType** | **string** | 授权类型，此处只需填写 client_credential | 

### Return type

[**TokenRsp**](TokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnsJscode2sessionGet

> Jscode2sessionRsp SnsJscode2sessionGet(ctx).Appid(appid).Secret(secret).JsCode(jsCode).GrantType(grantType).Execute()

登录凭证校验。

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
    appid := "appid_example" // string | 小程序 appId
    secret := "secret_example" // string | 小程序 secret
    jsCode := "jsCode_example" // string | 登录时获取的 code
    grantType := "grantType_example" // string | 授权类型，此处只需填写 authorization_code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiniProgramAuthApi.SnsJscode2sessionGet(context.Background()).Appid(appid).Secret(secret).JsCode(jsCode).GrantType(grantType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiniProgramAuthApi.SnsJscode2sessionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnsJscode2sessionGet`: Jscode2sessionRsp
    fmt.Fprintf(os.Stdout, "Response from `MiniProgramAuthApi.SnsJscode2sessionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnsJscode2sessionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appid** | **string** | 小程序 appId | 
 **secret** | **string** | 小程序 secret | 
 **jsCode** | **string** | 登录时获取的 code | 
 **grantType** | **string** | 授权类型，此处只需填写 authorization_code | 

### Return type

[**Jscode2sessionRsp**](Jscode2sessionRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WxaGetwxacodeunlimitGet

> *os.File WxaGetwxacodeunlimitGet(ctx).AccessToken(accessToken).GetwxacodeunlimitReq(getwxacodeunlimitReq).Execute()

获取小程序全局唯一后台接口调用凭据（access_token）。

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
    getwxacodeunlimitReq := *openapiclient.NewGetwxacodeunlimitReq() // GetwxacodeunlimitReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MiniProgramAuthApi.WxaGetwxacodeunlimitGet(context.Background()).AccessToken(accessToken).GetwxacodeunlimitReq(getwxacodeunlimitReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiniProgramAuthApi.WxaGetwxacodeunlimitGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WxaGetwxacodeunlimitGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MiniProgramAuthApi.WxaGetwxacodeunlimitGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWxaGetwxacodeunlimitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string** | access_token | 
 **getwxacodeunlimitReq** | [**GetwxacodeunlimitReq**](GetwxacodeunlimitReq.md) |  | 

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

