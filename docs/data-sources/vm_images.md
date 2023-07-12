---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cudo_vm_images Data Source - terraform-provider-cudo"
subcategory: ""
description: |-
  Images data source
---

# cudo_vm_images (Data Source)

Images data source



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) Placeholder identifier attribute.
- `images` (Attributes List) List of images. (see [below for nested schema](#nestedatt--images))

<a id="nestedatt--images"></a>
### Nested Schema for `images`

Read-Only:

- `description` (String) Image description
- `id` (String) Image identifier
- `name` (String) Image name
- `size_gib` (String) Image size in GiB

