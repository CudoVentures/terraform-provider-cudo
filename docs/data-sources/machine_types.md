---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cudo_machine_types Data Source - terraform-provider-cudo"
subcategory: ""
description: |-
  MachineType data source
---

# cudo_machine_types (Data Source)

MachineType data source

## Example Usage

```terraform
data "cudo_machine_types" "types" {
  search_params = {
    memory_gib = 4
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `search_params` (Attributes) Search parameters for limiting machine types (see [below for nested schema](#nestedatt--search_params))

### Read-Only

- `id` (String) Placeholder identifier attribute.
- `machine_types` (Attributes List) List of available machine types (see [below for nested schema](#nestedatt--machine_types))

<a id="nestedatt--search_params"></a>
### Nested Schema for `search_params`

Optional:

- `cpu_model` (String) CPU model name
- `data_center_id` (String) ID of the data center where the VM is located
- `gpu_count` (Number) Number of GPUs
- `gpu_model` (String) GPU model name
- `memory_gib` (Number) Amount of memory in GiB
- `order_by` (String) Field to order results by
- `page_number` (Number) Page number of the results to return
- `page_size` (Number) Number of results per page
- `region_id` (String) ID of the region to search in
- `storage_gib` (Number) Amount of storage in GiB
- `vcpu` (Number) Number of vCPUs


<a id="nestedatt--machine_types"></a>
### Nested Schema for `machine_types`

Read-Only:

- `count_vm_available` (Number) Number of available VMs of this configuration
- `cpu_model` (String) CPU model name
- `data_center_id` (String) ID of the data center where the VM is located
- `gpu_model` (String) GPU model name
- `gpu_price_hr` (String) Price per GPU per hour
- `id` (String) Machine type identifier
- `machine_type` (String) Machine type identifier
- `memory_gib_price_hr` (String) Price per GiB of memory per hour
- `storage_gib_price_hr` (String) Price per GiB of storage per hour
- `total_gpu_price_hr` (String) Total price for all GPUs per hour
- `total_memory_price_hr` (String) Total price for all memory per hour
- `total_price_hr` (String) Total price for the VM per hour
- `total_storage_price_hr` (String) Total price for all storage per hour
- `total_vcpu_price_hr` (String) Total price for all vCPUs per hour
- `vcpu_price_hr` (String) Price per vCPU per hour

