---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mezmo_gcp_cloud_storage_destination Resource - terraform-provider-mezmo"
subcategory: ""
description: |-
  Publish log events to GCP Cloud Storage
---

# mezmo_gcp_cloud_storage_destination (Resource)

Publish log events to GCP Cloud Storage



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auth` (Attributes) Configure GCP Cloud Storage authentication (see [below for nested schema](#nestedatt--auth))
- `bucket` (String) The name of the bucket in GCP where the data will be stored.
- `pipeline_id` (String) The uuid of the pipeline

### Optional

- `ack_enabled` (Boolean) Acknowledge data from the source when it reaches the destination
- `batch_timeout_secs` (Number) The maximum amount of time, in seconds, events will be buffered before being flushed to the destination
- `compression` (String) The compression strategy used on the encoded data prior to sending.
- `description` (String) A user-defined value describing the destination
- `encoding` (String) Dictates how the data will be serialized before storing.
- `inputs` (List of String) The ids of the input components
- `prefix` (String) The prefix applied to the bucket name, giving the appearance of having directories.
- `title` (String) A user-defined title for the destination

### Read-Only

- `generation_id` (Number) An internal field used for component versioning
- `id` (String) The uuid of the destination

<a id="nestedatt--auth"></a>
### Nested Schema for `auth`

Required:

- `type` (String) The type of authentication to use.
- `value` (String, Sensitive) Authentication secret value.

