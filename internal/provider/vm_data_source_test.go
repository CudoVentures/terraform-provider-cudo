package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/virtual_machines"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAcc_VMInstanceDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "tf-ds-test"

	resourcesConfig := fmt.Sprintf(`
resource "cudo_vm" "my-vm" {
   machine_type       = "standard"
   data_center_id     = "black-mesa"
   vcpus              = 1
   boot_disk = {
     image_id = "alpine-linux-317"
     size_gib = 1
   }
   memory_gib         = 2
   id                 = "%s"
 }`, name)

	testAccVMInstanceDataSourceConfig := fmt.Sprintf(`
data "cudo_vm" "test" {
	id = "%s"
}`, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			cl := getClient()

			getParams := virtual_machines.NewGetVMParamsWithContext(ctx)
			getParams.ID = name
			getParams.ProjectID = projectID
			ins, err := cl.VirtualMachines.GetVM(getParams)
			if err == nil && ins.Payload.VM.ShortState != "epil" {
				terminateParams := virtual_machines.NewTerminateVMParamsWithContext(ctx)
				terminateParams.ID = name
				terminateParams.ProjectID = projectID
				res, err := cl.VirtualMachines.TerminateVM(terminateParams)
				t.Log(res, err)

				return fmt.Errorf("vm resource not destroyed %s, %s, %s", ins.Payload.VM.ID, ins.Payload.VM.LcmState, ins.Payload.VM.OneState)
			}
			return nil
		},

		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + resourcesConfig,
			},
			{
				Config: getProviderConfig() + resourcesConfig + testAccVMInstanceDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_vm.test", "id", name)),
			},
		},
	})
}
