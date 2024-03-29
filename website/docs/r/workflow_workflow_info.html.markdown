---
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_info"
sidebar_current: "docs-intersight-resource-workflowWorkflowInfo"
description: |-
  Contains information for a workflow execution which is a runtime instance of workflow.

---

# Resource: intersight_workflow_workflow_info
Contains information for a workflow execution which is a runtime instance of workflow.

## Argument Reference
The following arguments are supported:
* `action`:(string)The action of the workflow such as start, cancel, retry, pause.
* `cleanup_time`:(string)(Computed)The time when the workflow info will be removed from database.
* `end_time`:(string)(Computed)The time when the workflow reached a final state.
* `failed_workflow_cleanup_duration`:(int)The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database.
* `input`:All the given inputs for the workflow.
* `inst_id`:(string)(Computed)A workflow instance Id which is the unique identified for the workflow execution.
* `internal`:(bool)Denotes if this workflow is internal and should be hidden from user view of running workflows.
* `meta_version`:(int)Version of the workflow metadata for which this workflow execution was started.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)A name of the workflow execution instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `organization`:(Array with Maximum of one item) -Relationship to the Organization that owns the Managed Object.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `output`:(Computed)All the generated outputs for the workflow.
* `parent_task_info`:(Array with Maximum of one item) -(Computed)Link to the task in the parent workflow which launched this workflow.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `pending_dynamic_workflow_info`:(Array with Maximum of one item) -(Computed)Reference to the PendingDynamicWorkflowInfo that was used to construct this workflow instance.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `progress`:(float)(Computed)This field indicates percentage of workflow task execution.
* `src`:(string)(Computed)The source microservice name which is the owner for this workflow.
* `start_time`:(string)(Computed)The time when the workflow was started for execution.
* `status`:(string)(Computed)A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED).
* `success_workflow_cleanup_duration`:(int)The duration in hours after which the workflow info for successful workflow will be removed from database.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
* `task_infos`:(Array)(Computed)List of task instances that ran as part of this workflow execution.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `type`:(string)(Computed)A type of the workflow (serverconfig, ansible_monitoring).
* `user_id`:(string)(Computed)The user identifier which indicates the user that started this workflow.
* `wait_reason`:(string)Denotes the reason workflow is in waiting status.
* `workflow_ctx`:The workflow context which contains initiator and target information.
* `workflow_definition`:(Array with Maximum of one item) -The workflow definition that was used to launch this workflow execution instance.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `workflow_meta_type`:(string)The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance.
* `workflow_task_count`:(int)(Computed)Total number of workflow tasks in this workflow.
