module github.com/XBeg9/pulumi-databricks/provider

go 1.14

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/databrickslabs/databricks-terraform v0.2.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.13.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.3.3
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
)
