---
layout: "intersight"
page_title: "Intersight: intersight_networkconfig_policy"
sidebar_current: "docs-intersight-resource-networkconfigPolicy"
description: |-
  Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.

---

# Resource: intersight_networkconfig_policy
Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.

## Argument Reference
The following arguments are supported:
* `alternate_ipv4dns_server`:(string)IP address of the secondary DNS server.
* `alternate_ipv6dns_server`:(string)IP address of the secondary DNS server.
* `description`:(string)Description of the policy.
* `dynamic_dns_domain`:(string)The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request.
* `enable_dynamic_dns`:(bool)If enabled, updates the resource records to the DNS from Cisco IMC.
* `enable_ipv6`:(bool)If enabled, allows to configure IPv6 properties.
* `enable_ipv4dns_from_dhcp`:(bool)If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it.
* `enable_ipv6dns_from_dhcp`:(bool)If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `organization`:(Array with Maximum of one item) -Relationship to the Organization that owns the Managed Object.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `preferred_ipv4dns_server`:(string)IP address of the primary DNS server.
* `preferred_ipv6dns_server`:(string)IP address of the primary DNS server.
* `profiles`:(Array)Relationship to the profile object.
This complex property has following sub-properties:
  + `moid`:(string)(Computed)The Moid of the referenced REST resource.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `selector`:(string)(Computed)An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'is set and 'moid' is empty/absent from the request, Intersight will determine the Moid of theresource matching the filter expression and populate it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request. An error is returned if the filtermatches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'.
* `tags`:(Array)The array of tags, which allow to add key, value meta-data to managed objects.
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string)The string representation of a tag key.
  + `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
  + `value`:(string)The string representation of a tag value.
