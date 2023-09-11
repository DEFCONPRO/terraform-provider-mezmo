---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mezmo_elasticsearch_destination Resource - terraform-provider-mezmo"
subcategory: ""
description: |-
  Represents an ElasticSearch destination.
---

# mezmo_elasticsearch_destination (Resource)

Represents an ElasticSearch destination.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auth` (Attributes) Configures ES authentication (see [below for nested schema](#nestedatt--auth))
- `endpoints` (List of String) An array of ElasticSearch endpoints
- `pipeline_id` (String) The uuid of the pipeline

### Optional

- `ack_enabled` (Boolean) Acknowledge data from the source when it reaches the destination
- `compression` (String) The compression strategy used on the encoded data prior to sending
- `description` (String) A user-defined value describing the destination
- `index` (String) Index to use when writing ES events (default = mezmo-%Y.%m.%d)
- `inputs` (List of String) The ids of the input components
- `pipeline` (String) Name of an ES ingest pipeline to include
- `title` (String) A user-defined title for the destination

### Read-Only

- `generation_id` (Number) An internal field used for component versioning
- `id` (String) The uuid of the destination

<a id="nestedatt--auth"></a>
### Nested Schema for `auth`

Required:

- `strategy` (String) The ES authentication strategy to use

Optional:

- `access_key_id` (String, Sensitive) The AWS access key id
- `password` (String, Sensitive) The password to use for basic authentication
- `region` (String) The AWS Region
- `secret_access_key` (String, Sensitive) The AWS secret access key
- `user` (String) The username for basic authentication

