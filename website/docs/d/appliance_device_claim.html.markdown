---
layout: "intersight"
page_title: "Intersight: intersight_appliance_device_claim"
sidebar_current: "docs-intersight-data-source-applianceDeviceClaim"
description: |-
DeviceClaim managed object represents a user initiated claim request for claiming
an endpoint device. There can be many DeviceClaim managed object for a given endpoint
device when users claim and unclaim devices repeatedly.

Claiming an endpoint device is a multi-step operation. The Intersight Appliance
starts a workflow with multiple tasks to process the device claim request. The status
of the device claim operation can be obtained from the claim workflow.

---

# Data Source: intersight_appliance_device_claim
DeviceClaim managed object represents a user initiated claim request for claiming
an endpoint device. There can be many DeviceClaim managed object for a given endpoint
device when users claim and unclaim devices repeatedly.

Claiming an endpoint device is a multi-step operation. The Intersight Appliance
starts a workflow with multiple tasks to process the device claim request. The status
of the device claim operation can be obtained from the claim workflow.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_id`:(string)Device identifier of the endpoint device.
* `hostname`:(string)Hostname or IP address of the endpoint device the user wants to claim.
* `is_password_set`:(bool)
* `message`:(string)Message set by the device claim process.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `password`:(string)Password to be used to login to the endpoint device.
* `platform_type`:(string)Platform type of the endpoint device.
* `request_id`:(string)User defined claim request identifier set by the UI. The RequestId field is not a mandatory. The Intersight Appliance will assign a unique value automatically if the field is not set.
* `security_token`:(string)Device security token of the endpoint device.
* `status`:(string)Status of the device claim process.
* `username`:(string)Username to log in to the endpoint device.
