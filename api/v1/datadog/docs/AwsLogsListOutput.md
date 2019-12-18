# AwsLogsListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Your AWS Account ID without dashes. | [optional] 
**Lambdas** | Pointer to [**[]AwsLogsListOutputLambdas**](AWSLogsListOutput_lambdas.md) | List of ARNs configured in your Datadog account. | [optional] 
**Services** | Pointer to **[]string** | Array of services IDs. | [optional] 

## Methods

### GetAccountId

`func (o *AwsLogsListOutput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsLogsListOutput) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *AwsLogsListOutput) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *AwsLogsListOutput) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetLambdas

`func (o *AwsLogsListOutput) GetLambdas() []AwsLogsListOutputLambdas`

GetLambdas returns the Lambdas field if non-nil, zero value otherwise.

### GetLambdasOk

`func (o *AwsLogsListOutput) GetLambdasOk() ([]AwsLogsListOutputLambdas, bool)`

GetLambdasOk returns a tuple with the Lambdas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLambdas

`func (o *AwsLogsListOutput) HasLambdas() bool`

HasLambdas returns a boolean if a field has been set.

### SetLambdas

`func (o *AwsLogsListOutput) SetLambdas(v []AwsLogsListOutputLambdas)`

SetLambdas gets a reference to the given []AwsLogsListOutputLambdas and assigns it to the Lambdas field.

### GetServices

`func (o *AwsLogsListOutput) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AwsLogsListOutput) GetServicesOk() ([]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServices

`func (o *AwsLogsListOutput) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServices

`func (o *AwsLogsListOutput) SetServices(v []string)`

SetServices gets a reference to the given []string and assigns it to the Services field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

