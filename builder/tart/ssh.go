package tart

import (
	"context"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TartMachineIP(ctx context.Context, vmName string, ipResolver string) func(multistep.StateBag) (string, error) {
	return func(state multistep.StateBag) (string, error) {
		cmdArgs := []string{"ip", "--wait", "120"}

		if ipResolver != "" {
			cmdArgs = append(cmdArgs, "--resolver", ipResolver)
		}

		cmdArgs = append(cmdArgs, vmName)
		return TartExec(ctx, cmdArgs...)
	}
}
