---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cudo_networks Data Source - terraform-provider-cudo"
subcategory: ""
description: |-
  Networks data source
---

# cudo_networks (Data Source)

Networks data source



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) Placeholder identifier attribute.
- `networks` (Attributes List) List of networks. (see [below for nested schema](#nestedatt--networks))

<a id="nestedatt--networks"></a>
### Nested Schema for `networks`

Read-Only:

- `cidr_prefix` (String) CIDR prefix i.e. 192.168.0.0/24
- `data_center_id` (String) The unique identifier of the datacenter where the network is located.
- `gateway` (String) gateway
- `id` (String) Network ID
- `vrouter_size` (String) Size of the vrouter 'small' 'medium' or 'large'

