---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cudo_security_groups Data Source - terraform-provider-cudo"
subcategory: ""
description: |-
  Security groups data source
---

# cudo_security_groups (Data Source)

Security groups data source



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `data_center_id` (String) Datacenter ID to request security groups from

### Read-Only

- `id` (String) Placeholder identifier attribute.
- `security_groups` (Attributes List) List of security groups (see [below for nested schema](#nestedatt--security_groups))

<a id="nestedatt--security_groups"></a>
### Nested Schema for `security_groups`

Read-Only:

- `data_center_id` (String) Datacenter ID
- `description` (String) Security group description
- `id` (String) Image identifier
- `rules` (Attributes List) List of rules in security group (see [below for nested schema](#nestedatt--security_groups--rules))

<a id="nestedatt--security_groups--rules"></a>
### Nested Schema for `security_groups.rules`

Read-Only:

- `icmp_type` (String) ICMP type
- `id` (String) Rule ID
- `ip_range_cidr` (String) IP range
- `ports` (String) Image size in GiB
- `protocol` (String) Image size in GiB
- `rule_type` (String) Image size in GiB

