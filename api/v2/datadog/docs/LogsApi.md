# \LogsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogs**](LogsApi.md#ListLogs) | **Post** /api/v2/logs/events/search | Get a list of logs
[**ListLogsGet**](LogsApi.md#ListLogsGet) | **Get** /api/v2/logs/events | Get a quick list of logs



## ListLogs

> LogsListResponse ListLogs(ctx).Body(body).Execute()

Get a list of logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    body := datadog.LogsListRequest{Filter: datadog.LogsListRequest_filter{From: "TODO", Indexes: []string{"Indexes_example"), Query: "Query_example", To: "TODO"}, Page: datadog.LogsListRequest_page{Cursor: "Cursor_example", Limit: 123}, Sort: datadog.LogsSort{}} // LogsListRequest |  (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.ListLogs(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogs`: LogsListResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsListRequest**](LogsListRequest.md) |  | 

### Return type

[**LogsListResponse**](LogsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogsGet

> LogsListResponse ListLogsGet(ctx).FilterQuery(filterQuery).FilterIndex(filterIndex).FilterFrom(filterFrom).FilterTo(filterTo).Sort(sort).PageCursor(pageCursor).PageLimit(pageLimit).Execute()

Get a quick list of logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    filterQuery := "filterQuery_example" // string | Search query following logs syntax. (optional)
    filterIndex := "filterIndex_example" // string | For customers with multiple indexes, the indexes to search Defaults to '*' which means all indexes (optional)
    filterFrom := Get-Date // time.Time | Minimum timestamp for requested logs. (optional)
    filterTo := Get-Date // time.Time | Maximum timestamp for requested logs. (optional)
    sort := datadog.LogsSort{} // LogsSort | Order of logs in results. (optional)
    pageCursor := "pageCursor_example" // string | List following results with a cursor provided in the previous query. (optional)
    pageLimit := 987 // int32 | Maximum number of logs in the response. (optional) (default to 10)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsApi.ListLogsGet(ctx).FilterQuery(filterQuery).FilterIndex(filterIndex).FilterFrom(filterFrom).FilterTo(filterTo).Sort(sort).PageCursor(pageCursor).PageLimit(pageLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogsGet`: LogsListResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.ListLogsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterQuery** | **string** | Search query following logs syntax. | 
 **filterIndex** | **string** | For customers with multiple indexes, the indexes to search Defaults to &#39;*&#39; which means all indexes | 
 **filterFrom** | **time.Time** | Minimum timestamp for requested logs. | 
 **filterTo** | **time.Time** | Maximum timestamp for requested logs. | 
 **sort** | [**LogsSort**](.md) | Order of logs in results. | 
 **pageCursor** | **string** | List following results with a cursor provided in the previous query. | 
 **pageLimit** | **int32** | Maximum number of logs in the response. | [default to 10]

### Return type

[**LogsListResponse**](LogsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

