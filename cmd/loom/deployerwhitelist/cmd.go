package deployer_whitelist

import (
	"encoding/json"
	"fmt"
	"strings"

	dwtypes "github.com/loomnetwork/go-loom/builtin/types/deployer_whitelist"
	"github.com/loomnetwork/go-loom/cli"
	dw "github.com/loomnetwork/loomchain/builtin/plugins/deployer_whitelist"
	"github.com/spf13/cobra"
)

var (
	dwContractName = "deployerwhitelist"
)

type deployerInfo struct {
	Address string
	Flags   string
}

func NewDeployCommand() *cobra.Command {
	cmd := cli.ContractCallCommand("deployerwhitelist")
	cmd.Use = "deployer"
	cmd.Short = "Deployer Whitelist CLI"
	cmd.AddCommand(
		addDeployerCmd(),
		getDeployerCmd(),
		listDeployersCmd(),
		removeDeployerCmd(),
		setDefaultDeployerCmd(),
		getDefaultDeployerCmd(),
	)
	return cmd
}

const addDeployerCmdExample = `
loom deployer add 0x7262d4c97c7B93937E4810D289b7320e9dA82857 all
`

func addDeployerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "add <deployer address> <permission (go|evm|migration|all)>",
		Short:   "Add deployer with permision to deployer list",
		Example: addDeployerCmdExample,
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := parseAddress(args[0])
			if err != nil {
				return err
			}

			var flags uint32
			if strings.EqualFold(args[1], "evm") {
				flags = uint32(dw.AllowEVMDeployFlag)
			} else if strings.EqualFold(args[1], "go") {
				flags = uint32(dw.AllowGoDeployFlag)
			} else if strings.EqualFold(args[1], "migration") {
				flags = uint32(dw.AllowMigrationFlag)
			} else if strings.EqualFold(args[1], "all") {
				flags = dw.PackFlags(
					uint32(dw.AllowEVMDeployFlag), uint32(dw.AllowGoDeployFlag),
					uint32(dw.AllowMigrationFlag),
				)
			} else {
				return fmt.Errorf("Please specify deploy permission (go|evm|migration|all)")
			}

			cmd.SilenceUsage = true

			req := &dwtypes.AddDeployerRequest{
				DeployerAddr: addr.MarshalPB(),
				Flags:        flags,
			}

			return cli.CallContract(dwContractName, "AddDeployer", req, nil)
		},
	}
}

const removeDeployerCmdExample = `
loom deployer remove 0x7262d4c97c7B93937E4810D289b7320e9dA82857
`

func removeDeployerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "remove <deployer address>",
		Short:   "Remove deployer from whitelist",
		Example: removeDeployerCmdExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := parseAddress(args[0])
			if err != nil {
				return err
			}

			cmd.SilenceUsage = true

			req := &dwtypes.RemoveDeployerRequest{
				DeployerAddr: addr.MarshalPB(),
			}

			return cli.CallContract(dwContractName, "RemoveDeployer", req, nil)
		},
	}
}

const getDeployerCmdExample = `
loom deployer get 0x7262d4c97c7B93937E4810D289b7320e9dA82857
`

func getDeployerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "get <deployer address>",
		Short:   "Show current permissions of a deployer",
		Example: getDeployerCmdExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := parseAddress(args[0])
			if err != nil {
				return err
			}

			cmd.SilenceUsage = true

			req := &dwtypes.GetDeployerRequest{
				DeployerAddr: addr.MarshalPB(),
			}

			var resp dwtypes.GetDeployerResponse
			if err := cli.StaticCallContract(dwContractName, "GetDeployer", req, &resp); err != nil {
				return err
			}

			deployer := getDeployerInfo(resp.Deployer)

			output, err := json.MarshalIndent(deployer, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))
			return nil
		},
	}
}

const listDeployersCmdExample = `
loom deployer list
`

func listDeployersCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "Display all deployers in whitelist",
		Example: listDeployersCmdExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			req := &dwtypes.ListDeployersRequest{}
			var resp dwtypes.ListDeployersResponse
			if err := cli.StaticCallContract(dwContractName, "ListDeployers", req, &resp); err != nil {
				return err
			}

			deployersInfo := []*deployerInfo{}
			for _, deployer := range resp.Deployers {
				deployerInfo := getDeployerInfo(deployer)
				deployersInfo = append(deployersInfo, &deployerInfo)
			}

			output, err := json.MarshalIndent(deployersInfo, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))
			return nil
		},
	}
}

const setDefaultDeployerCmdExample = `
loom deployer set-default go evm migration
`

func setDefaultDeployerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "set-default <flag1> .. <flagN>",
		Short:   "Set default deployer permision",
		Example: setDefaultDeployerCmdExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags := uint32(0)
			for _, f := range args {
				if strings.EqualFold(f, "evm") {
					flags = dw.PackFlags(flags, uint32(dw.AllowEVMDeployFlag))
				} else if strings.EqualFold(f, "go") {
					flags = dw.PackFlags(flags, uint32(dw.AllowGoDeployFlag))
				} else if strings.EqualFold(f, "migration") {
					flags = dw.PackFlags(flags, uint32(dw.AllowMigrationFlag))
				} else if strings.EqualFold(f, "all") {
					flags = dw.PackFlags(
						flags,
						uint32(dw.AllowEVMDeployFlag),
						uint32(dw.AllowGoDeployFlag),
						uint32(dw.AllowMigrationFlag),
					)
				}
			}

			cmd.SilenceUsage = true

			req := &dwtypes.SetDefaultDeployerRequest{
				Flags: flags,
			}

			return cli.CallContract(dwContractName, "SetDefaultDeployer", req, nil)
		},
	}
}

const getDefaultDeployerCmdExample = `
loom deployer get-default
`

func getDefaultDeployerCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "get-default",
		Short:   "Show current permissions of default deployer",
		Example: getDeployerCmdExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			var resp dwtypes.GetDefaultDeployerResponse
			if err := cli.StaticCallContract(
				dwContractName,
				"GetDefaultDeployer",
				&dwtypes.GetDefaultDeployerRequest{},
				&resp); err != nil {
				return err
			}

			deployer := getDeployerInfo(resp.Deployer)

			output, err := json.MarshalIndent(deployer, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(output))
			return nil
		},
	}
}

func getDeployerInfo(deployer *dwtypes.Deployer) deployerInfo {
	flagsInt := dw.UnpackFlags(deployer.Flags)
	flags := []string{}
	for _, flag := range flagsInt {
		flags = append(flags, dwtypes.Flags_name[int32(flag)])
	}
	f := strings.Join(flags, "|")
	deployerInfo := deployerInfo{
		Address: deployer.Address.ChainId + ":" + deployer.Address.Local.String(),
		Flags:   f,
	}
	return deployerInfo
}
