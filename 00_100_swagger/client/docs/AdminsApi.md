# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/YWSROCK_1/hello/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInventory**](AdminsApi.md#AddInventory) | **Post** /inventory | adds an inventory item

# **AddInventory**
> AddInventory(ctx, optional)
adds an inventory item

Adds an item to the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdminsApiAddInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminsApiAddInventoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InventoryItem**](InventoryItem.md)| Inventory item to add | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

