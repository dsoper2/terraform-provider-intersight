---
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_meta"
sidebar_current: "docs-intersight-data-source-workflowTaskMeta"
description: |-
This MO contains a task definition.

---

# Data Source: intersight_workflow_task_meta
This MO contains a task definition.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action_type`:(string)A task execution type to indicate if it is a system task.
* `description`:(string)A description of the task.
* `internal`:(bool)Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)A task name that should be unique in Conductor DB.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `response_timeout_sec`:(int)The worker respnose timeout value.
* `retry_count`:(int)A number of reties for this task.
* `retry_delay_sec`:(int)The time on which the retry will be delayed.
* `retry_logic`:(string)A logic which defines the way to handle retry (FIXED, EXPONENTIAL_BACKOFF).
* `src`:(string)A service owns the task metadata.
* `timeout_policy`:(string)A policy which defines the way to handle timeout (RETRY, TIME_OUT_WF, ALERT_ONLY).
* `timeout_sec`:(int)A timeout value for the task in seconds.
