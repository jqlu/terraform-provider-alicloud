---
layout: "alicloud"
page_title: "Alicloud: alicloud_fc_function"
sidebar_current: "docs-alicloud-resource-fc"
description: |-
  Provides a Alicloud Function Compute Function resource. Function allows you to trigger execution of code in response to events in Alibaba Cloud. The Function itself includes source code and runtime configuration.
---

# alicloud\_fc\_function

Provides a Alicloud Function Compute Function resource. Function allows you to trigger execution of code in response to events in Alibaba Cloud. The Function itself includes source code and runtime configuration.
 For information about Service and how to use it, see [What is Function Compute](https://www.alibabacloud.com/help/doc-detail/52895.htm).

-> **NOTE:** The resource requires a provider field 'account_id'. [See account_id](https://www.terraform.io/docs/providers/alicloud/index.html#account_id).

## Example Usage

Basic Usage

```
variable "name" {
    default = "alicloudfcfunctionconfig"
}
resource "alicloud_log_project" "default" {
    name = "${var.name}"
    description = "tf unit test"
}

resource "alicloud_log_store" "default" {
    project = "${alicloud_log_project.default.name}"
    name = "${var.name}"
    retention_period = "3000"
    shard_count = 1
}
resource "alicloud_fc_service" "default" {
    name = "${var.name}"
    description = "tf unit test"
    log_config {
        project = "${alicloud_log_project.default.name}"
        logstore = "${alicloud_log_store.default.name}"
    }
    role = "${alicloud_ram_role.default.arn}"
    depends_on = ["alicloud_ram_role_policy_attachment.default"]
}
resource "alicloud_oss_bucket" "default" {
    bucket = "${var.name}"
}

resource "alicloud_oss_bucket_object" "default" {
    bucket = "${alicloud_oss_bucket.default.id}"
    key = "fc/hello.zip"
    content = <<EOF
        # -*- coding: utf-8 -*-
        def handler(event, context):
            print "hello world"
            return 'hello world'
    EOF
}

resource "alicloud_ram_role" "default" {
    name = "${var.name}"
    document = <<EOF
        {
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Effect": "Allow",
              "Principal": {
                "Service": [
                  "fc.aliyuncs.com"
                ]
              }
            }
          ],
          "Version": "1"
        }
    EOF
    description = "this is a test"
    force = true
}

resource "alicloud_ram_role_policy_attachment" "default" {
    role_name = "${alicloud_ram_role.default.name}"
    policy_name = "AliyunLogFullAccess"
    policy_type = "System"
}

resource "alicloud_fc_function" "foo" {
    service = "${alicloud_fc_service.default.name}"
    name = "${var.name}"
    description = "tf"
    oss_bucket = "${alicloud_oss_bucket.default.id}"
    oss_key = "${alicloud_oss_bucket_object.default.key}"
    memory_size = "512"
    runtime = "python2.7"
    handler = "hello.handler"
    environment_variables {
        prefix = "terraform"
    }
}
```
## Argument Reference

The following arguments are supported:

* `service` - (Required, ForceNew) The Function Compute service name.
* `name` - (Optional, ForceNew) The Function Compute function name. It is the only in one service and is conflict with "name_prefix".
* `name_prefix` - (Optional, ForceNew) Setting a prefix to get a only function name. It is conflict with "name".
* `description` - (Optional) The Function Compute function description.
* `filename` - (Optional) The path to the function's deployment package within the local filesystem. It is conflict with the `oss_`-prefixed options.
* `oss_bucket` - (Optional) The OSS bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same Alibaba Cloud region where you are creating the function.
* `oss_key` - (Optional) The OSS key of an object containing the function's deployment package. Conflicts with `filename`.
* `handler` - (Required) The function [entry point](https://www.alibabacloud.com/help/doc-detail/62213.htm) in your code.
* `memory_size` - (Optional) Amount of memory in MB your Function can use at runtime. Defaults to `128`. Limits to [128, 3072].
* `runtime` - (Required) See [Runtimes][https://www.alibabacloud.com/help/doc-detail/52077.htm] for valid values.
* `timeout` - (Optional) The amount of time your Function has to run in seconds.
* `environment_variables` - (Optional, Available in 1.36.0+) A map that defines environment variables for the function.
-> **NOTE:** For more information, see [Limits](https://www.alibabacloud.com/help/doc-detail/51907.htm).

## Attributes Reference

The following arguments are exported:

* `id` - The ID of the function. The value is formate as `<service>:<name>`.
* `last_modified` - The date this resource was last modified.

## Import

Function Compute function can be imported using the id, e.g.

```
$ terraform import alicloud_fc_service.foo my-fc-service:hello-world
```
