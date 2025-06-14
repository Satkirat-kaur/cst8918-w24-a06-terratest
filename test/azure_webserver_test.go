package test
import (
    "testing"

    "github.com/gruntwork-io/terratest/modules/azure"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "eec42dd1-da1e-4954-9b40-f1b92133a742"

func TestAzureLinuxVMCreation(t *testing.T) {
    terraformOptions := &terraform.Options{
        // The path to where our Terraform code is located
        TerraformDir: "../",
        // Override the default terraform variables
        Vars: map[string]interface{}{
            "labelPrefix": "kaur1852",
        },
    }

    defer terraform.Destroy(t, terraformOptions)

    // Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
    terraform.InitAndApply(t, terraformOptions)

    // Run `terraform output` to get the value of output variable
    vmName := terraform.Output(t, terraformOptions, "vm_name")
    resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

    // Confirm VM exists
    assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))

    //Test that NIC exists
    nicName := terraform.Output(t, terraformOptions, "nic_name")
    assert.True(t, azure.NetworkInterfaceExists(t, nicName, resourceGroupName, subscriptionID))

    // Test VM is running correct image
    // vm := azure.GetVirtualMachineImage(t, vmName, resourceGroupName, subscriptionID) 
    // assert.Contains(t, vm.Offer, "0001-com-ubuntu-server-jammy")
    // assert.Contains(t, vm.Publisher,"Canonical")clear
    // assert.Contains(t, vm.SKU,"22_04-lts-gen2")
}
