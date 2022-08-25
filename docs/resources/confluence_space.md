---
layout: "confluence"
page_title: "Confluence: confluence_space"
sidebar_current: "docs-confluence-resource-space"
description: |-
  Provides space in Confluence
---

# confluence_space

Provides a space on your Confluence site.

## Example Usage

```hcl
resource confluence_space "default" {
  key  = "MYSPACE"
  name   = "My-Space"
}
```

## Argument Reference

The following arguments are supported:

- `key` - (Required) The key for the space

- `name` - (Required) the display name for the space.

## Import

space can be imported using the space key.

```
$ terraform import confluence_space.default {{key}}
```
