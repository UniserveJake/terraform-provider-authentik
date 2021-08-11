---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "authentik_property_mapping_ldap Data Source - terraform-provider-authentik"
subcategory: ""
description: |-
  Get LDAP Property mappings
---

# authentik_property_mapping_ldap (Data Source)

Get LDAP Property mappings

## Example Usage

```terraform
# To get the ID of a LDAP Property mapping

data "authentik_property_mapping_ldap" "test" {
  managed = "goauthentik.io/sources/ldap/default-name"
}

# Then use `data.authentik_property_mapping_ldap.test.id`

# Or, to get the IDs of multiple mappings

data "authentik_property_mapping_ldap" "test" {
  managed_list = [
    "goauthentik.io/sources/ldap/default-name",
    "goauthentik.io/sources/ldap/default-mail"
  ]
}

# Then use data.authentik_property_mapping_ldap.test.ids
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **ids** (List of String) List of ids when `managed_list` is set.
- **managed** (String)
- **managed_list** (List of String) Retrive multiple property mappings
- **name** (String)
- **object_field** (String)

### Read-Only

- **expression** (String)

