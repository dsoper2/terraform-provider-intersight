---
layout: "intersight"
page_title: "Intersight: intersight_appliance_setup_info"
sidebar_current: "docs-intersight-data-source-applianceSetupInfo"
description: |-
SetupInfo will have only one managed object. SetupInfo managed object is to keep
track of the Intersight Appliance's setup information and guide the UI through
the initial configuration of the Intersight Appliance.

The SetupInfo managed object is created during the Intersight Appliance setup.
The Intersight UI uses this object to store the initial configuration states
that the user has completed. If the user closes the Intersight UI without
finishing all the initial configuration, then the Intersight UI will use this
managed object to display the next configuration that the user needs to complete
when the user uses the Intersight Appliance next time.

---

# Data Source: intersight_appliance_setup_info
SetupInfo will have only one managed object. SetupInfo managed object is to keep
track of the Intersight Appliance's setup information and guide the UI through
the initial configuration of the Intersight Appliance.

The SetupInfo managed object is created during the Intersight Appliance setup.
The Intersight UI uses this object to store the initial configuration states
that the user has completed. If the user closes the Intersight UI without
finishing all the initial configuration, then the Intersight UI will use this
managed object to display the next configuration that the user needs to complete
when the user uses the Intersight Appliance next time.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `build_type`:(string)Build type of the Intersight Appliance setup (e.g. release or debug).
* `cloud_url`:(string)URL of the Intersight to which this Intersight Appliance is connected to.
* `end_time`:(string)End date of the Intersight Appliance's initial setup.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `start_time`:(string)Start date of the Intersight Appliance's initial setup.
