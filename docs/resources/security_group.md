---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cudo_security_group Resource - terraform-provider-cudo"
subcategory: ""
description: |-
  Security group resource
---

# cudo_security_group (Resource)

Security group resource

## Example Usage

```terraform
resource "cudo_security_group" "my-sg" {
  id             = "my-sg"
  data_center_id = "gb-london-1"
  description    = "security group for a purpose"
  rules = [
    {
      ports     = "22,80,443"
      rule_type = "outbound"
      protocol  = "tcp"
    },
    {
      ports     = "22,80,443"
      rule_type = "inbound"
      protocol  = "tcp"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `data_center_id` (String) The unique identifier of the datacenter where the network is located.
- `id` (String) Security Group ID
- `rules` (Attributes List) List of security group rules (see [below for nested schema](#nestedatt--rules))

### Optional

- `description` (String) Description of the security group

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Required:

- `protocol` (String) Protocol for rule, use one of: tcp or udp
- `rule_type` (String) Type for rule either 'inbound' or 'outbound'

Optional:

- `icmp_type` (String) Specific ICMP type of the rule. If a type has multiple codes, it includes all the codes within. This can only be used with ICMP
- `ip_range_cidr` (String) A single IP address or CIDR format range to apply rule to
- `ports` (String) A comma separated list of ports (80,443,8080) or a single port range (1024:2048)

Read-Only:

- `id` (String) The unique identifier of the rule

