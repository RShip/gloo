
---
title: "azure.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `azure.plugins.gloo.solo.io` 
#### Types:


- [UpstreamSpec](#upstreamspec)
- [FunctionSpec](#functionspec)
- [AuthLevel](#authlevel)
- [DestinationSpec](#destinationspec)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/plugins/azure/azure.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/azure/azure.proto)





---
### UpstreamSpec

 
Upstream Spec for Azure Functions Upstreams
Azure Upstreams represent a collection of Azure Functions for a particular Azure Account
within a particular Function App

```yaml
"functionAppName": string
"secretRef": .core.solo.io.ResourceRef
"functions": []azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `functionAppName` | `string` | The Name of the Azure Function App where the functions are grouped. |  |
| `secretRef` | [.core.solo.io.ResourceRef](../../../../../../../../solo-kit/api/v1/ref.proto.sk/#resourceref) | A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an [Azure Publish Profile JSON file](https://azure.microsoft.com/en-us/downloads/publishing-profile-overview/). {{ hide_not_implemented "Azure Secrets can be created with `glooctl secret create azure ...`" }} Note that this secret is not required unless Function Discovery is enabled. |  |
| `functions` | [[]azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec](../azure.proto.sk/#functionspec) |  |  |




---
### FunctionSpec

 
Function Spec for Functions on Azure Functions Upstreams
The Function Spec contains data necessary for Gloo to invoke Azure functions

```yaml
"functionName": string
"authLevel": .azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `functionName` | `string` | The Name of the Azure Function as it appears in the Azure Functions Portal. |  |
| `authLevel` | [.azure.plugins.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel](../azure.proto.sk/#authlevel) | Auth Level can bve either "anonymous" "function" or "admin" See https://vincentlauzon.com/2017/12/04/azure-functions-http-authorization-levels/ for more details. |  |




---
### AuthLevel



| Name | Description |
| ----- | ----------- | 
| `Anonymous` |  |
| `Function` |  |
| `Admin` |  |




---
### DestinationSpec



```yaml
"functionName": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `functionName` | `string` | The Function Name of the FunctionSpec to be invoked. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->