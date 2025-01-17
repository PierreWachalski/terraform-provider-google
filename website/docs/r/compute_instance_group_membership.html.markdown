---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
description: |-
  Represents the Instance membership to the Instance Group.
---

# google\_compute\_instance\_group\_membership

Represents the Instance membership to the Instance Group.

**NOTE** You can use this resource instead of the `instances` field in the
`google_compute_instance_group`, however it's not recommended to use it alongside this field.
It might cause inconsistencies, as they can end up competing over control.

**NOTE** This resource has been added to avoid a situation, where after
Instance is recreated, it's removed from Instance Group and it's needed to
perform `apply` twice. To avoid situations like this, please use this resource
with the lifecycle `update_triggered_by` method, with the passed Instance's ID.


To get more information about InstanceGroupMembership, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroups)
* How-to Guides
    * [Add instances](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroups/addInstances)
    * [Remove instances](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroups/removeInstances)
    * [List instances](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroups/listInstances)

## Example Usage - Instance Group Membership


```hcl
resource "google_compute_network" "default-network" {
  name = "network"
}

resource "google_compute_instance" "default-instance" {
  name         = "instance"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = google_compute_network.default-network.name
  }
}

resource "google_compute_instance_group" "default-instance-group" {
  name      = "instance-group"
}

resource "google_compute_instance_group_membership" "default-ig-membership" {
  instance        = google_compute_instance.default-instance.self_link
  instance_group  = google_compute_instance_group.default-instance-group.name
}
```

## Argument Reference

The following arguments are supported:


* `instance` -
  (Required)
  An instance being added to the InstanceGroup

* `instance_group` -
  (Required)
  Represents an Instance Group resource name that the instance belongs to.


- - -


* `zone` -
  (Optional)
  A reference to the zone where the instance group resides.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{project}}/{{zone}}/{{instance_group}}/{{instance}}`


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


InstanceGroupMembership can be imported using any of these accepted formats:

* `projects/{{project}}/zones/{{zone}}/instanceGroups/{{instance_group}}/{{instance}}`
* `{{project}}/{{zone}}/{{instance_group}}/{{instance}}`
* `{{zone}}/{{instance_group}}/{{instance}}`
* `{{instance_group}}/{{instance}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import InstanceGroupMembership using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/zones/{{zone}}/instanceGroups/{{instance_group}}/{{instance}}"
  to = google_compute_instance_group_membership.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), InstanceGroupMembership can be imported using one of the formats above. For example:

```
$ terraform import google_compute_instance_group_membership.default projects/{{project}}/zones/{{zone}}/instanceGroups/{{instance_group}}/{{instance}}
$ terraform import google_compute_instance_group_membership.default {{project}}/{{zone}}/{{instance_group}}/{{instance}}
$ terraform import google_compute_instance_group_membership.default {{zone}}/{{instance_group}}/{{instance}}
$ terraform import google_compute_instance_group_membership.default {{instance_group}}/{{instance}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
