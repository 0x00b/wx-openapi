# \SubscribemessageApi

All URIs are relative to *https://api.weixin.qq.com/cgi-bin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MessageSubscribeSendPost**](SubscribemessageApi.md#MessageSubscribeSendPost) | **Post** /message/subscribe/send | 发送订阅消息



## MessageSubscribeSendPost

> SubscribeMessageSendResponse MessageSubscribeSendPost(ctx).AccessToken(accessToken).SubscribeMessageSendRequest(subscribeMessageSendRequest).Execute()

发送订阅消息

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
    subscribeMessageSendRequest := *openapiclient.NewSubscribeMessageSendRequest() // SubscribeMessageSendRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscribemessageApi.MessageSubscribeSendPost(context.Background()).AccessToken(accessToken).SubscribeMessageSendRequest(subscribeMessageSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribemessageApi.MessageSubscribeSendPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessageSubscribeSendPost`: SubscribeMessageSendResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscribemessageApi.MessageSubscribeSendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessageSubscribeSendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string** | access_token | 
 **subscribeMessageSendRequest** | [**SubscribeMessageSendRequest**](SubscribeMessageSendRequest.md) |  | 

### Return type

[**SubscribeMessageSendResponse**](SubscribeMessageSendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

