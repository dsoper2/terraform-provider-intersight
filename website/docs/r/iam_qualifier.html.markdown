---
layout: "intersight"
page_title: "Intersight: intersight_iam_qualifier"
sidebar_current: "docs-intersight-resource-iamQualifier"
description: |-
  The qualifier defines how a user qualifies to be part of a user group.

---

# Resource: intersight_iam_qualifier
The qualifier defines how a user qualifies to be part of a user group.

## Argument Reference
The following arguments are supported:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)(Computed)The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
* `usergroup`:(Array with Maximum of one item) -A collection of references to the [iam.UserGroup](mo://iam.UserGroup) Managed Object.When this managed object is deleted, the referenced [iam.UserGroup](mo://iam.UserGroup) MO unsets its reference to this deleted MO.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `value`:
                (Array of schema.TypeString) -The value of the SAML attribute.
