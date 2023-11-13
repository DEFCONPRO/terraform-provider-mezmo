---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mezmo_filter_processor Resource - terraform-provider-mezmo"
subcategory: ""
description: |-
  Define condition(s) to include or exclude events from the pipeline
---

# mezmo_filter_processor (Resource)

Define condition(s) to include or exclude events from the pipeline



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `action` (String) How to handle events matching this criteria
- `conditional` (Attributes) A group of expressions (optionally nested) joined by a logical operator (see [below for nested schema](#nestedatt--conditional))
- `pipeline_id` (String) The uuid of the pipeline

### Optional

- `description` (String) A user-defined value describing the processor
- `inputs` (List of String) The ids of the input components
- `title` (String) A user-defined title for the processor

### Read-Only

- `generation_id` (Number) An internal field used for component versioning
- `id` (String) The uuid of the processor

<a id="nestedatt--conditional"></a>
### Nested Schema for `conditional`

Optional:

- `expressions` (Attributes List) Defines a list of expressions for field comparisons (see [below for nested schema](#nestedatt--conditional--expressions))
- `expressions_group` (Attributes List) A group of expressions joined by a logical operator (see [below for nested schema](#nestedatt--conditional--expressions_group))
- `logical_operation` (String) The logical operation (AND/OR) to be applied to the list of conditionals

<a id="nestedatt--conditional--expressions"></a>
### Nested Schema for `conditional.expressions`

Required:

- `field` (String) The field path whose value will be used in the comparison
- `operator` (String) The comparison operator

Optional:

- `value_number` (Number) The operand to compare the field value with, when the value is a number
- `value_string` (String) The operand to compare the field value with, when the value is a string


<a id="nestedatt--conditional--expressions_group"></a>
### Nested Schema for `conditional.expressions_group`

Optional:

- `expressions` (Attributes List) Defines a list of expressions for field comparisons (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions))
- `expressions_group` (Attributes List) A group of expressions joined by a logical operator (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group))
- `logical_operation` (String) The logical operation (AND/OR) to be applied to the list of conditionals

<a id="nestedatt--conditional--expressions_group--expressions"></a>
### Nested Schema for `conditional.expressions_group.expressions`

Required:

- `field` (String) The field path whose value will be used in the comparison
- `operator` (String) The comparison operator

Optional:

- `value_number` (Number) The operand to compare the field value with, when the value is a number
- `value_string` (String) The operand to compare the field value with, when the value is a string


<a id="nestedatt--conditional--expressions_group--expressions_group"></a>
### Nested Schema for `conditional.expressions_group.expressions_group`

Optional:

- `expressions` (Attributes List) Defines a list of expressions for field comparisons (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--expressions))
- `expressions_group` (Attributes List) A group of expressions joined by a logical operator (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--expressions_group))
- `logical_operation` (String) The logical operation (AND/OR) to be applied to the list of conditionals

<a id="nestedatt--conditional--expressions_group--expressions_group--expressions"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation`

Required:

- `field` (String) The field path whose value will be used in the comparison
- `operator` (String) The comparison operator

Optional:

- `value_number` (Number) The operand to compare the field value with, when the value is a number
- `value_string` (String) The operand to compare the field value with, when the value is a string


<a id="nestedatt--conditional--expressions_group--expressions_group--expressions_group"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation`

Optional:

- `expressions` (Attributes List) Defines a list of expressions for field comparisons (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions))
- `expressions_group` (Attributes List) A group of expressions joined by a logical operator (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group))
- `logical_operation` (String) The logical operation (AND/OR) to be applied to the list of conditionals

<a id="nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation.expressions`

Required:

- `field` (String) The field path whose value will be used in the comparison
- `operator` (String) The comparison operator

Optional:

- `value_number` (Number) The operand to compare the field value with, when the value is a number
- `value_string` (String) The operand to compare the field value with, when the value is a string


<a id="nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation.expressions_group`

Optional:

- `expressions` (Attributes List) Defines a list of expressions for field comparisons (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--expressions))
- `expressions_group` (Attributes List) A group of expressions joined by a logical operator (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--expressions_group))
- `logical_operation` (String) The logical operation (AND/OR) to be applied to the list of conditionals

<a id="nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--expressions"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation.expressions_group.logical_operation`

Required:

- `field` (String) The field path whose value will be used in the comparison
- `operator` (String) The comparison operator

Optional:

- `value_number` (Number) The operand to compare the field value with, when the value is a number
- `value_string` (String) The operand to compare the field value with, when the value is a string


<a id="nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--expressions_group"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation.expressions_group.logical_operation`

Optional:

- `expressions` (Attributes List) Defines a list of expressions for field comparisons (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--logical_operation--expressions))
- `expressions_group` (Attributes List) A group of expressions joined by a logical operator (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--logical_operation--expressions_group))
- `logical_operation` (String) The logical operation (AND/OR) to be applied to the list of conditionals

<a id="nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--logical_operation--expressions"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation.expressions_group.logical_operation.logical_operation`

Required:

- `field` (String) The field path whose value will be used in the comparison
- `operator` (String) The comparison operator

Optional:

- `value_number` (Number) The operand to compare the field value with, when the value is a number
- `value_string` (String) The operand to compare the field value with, when the value is a string


<a id="nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--logical_operation--expressions_group"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation.expressions_group.logical_operation.logical_operation`

Optional:

- `expressions` (Attributes List) Defines a list of expressions for field comparisons (see [below for nested schema](#nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--logical_operation--logical_operation--expressions))
- `logical_operation` (String) The logical operation (AND/OR) to be applied to the list of conditionals

<a id="nestedatt--conditional--expressions_group--expressions_group--logical_operation--expressions_group--logical_operation--logical_operation--expressions"></a>
### Nested Schema for `conditional.expressions_group.expressions_group.logical_operation.expressions_group.logical_operation.logical_operation.logical_operation`

Required:

- `field` (String) The field path whose value will be used in the comparison
- `operator` (String) The comparison operator

Optional:

- `value_number` (Number) The operand to compare the field value with, when the value is a number
- `value_string` (String) The operand to compare the field value with, when the value is a string