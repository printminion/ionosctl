package commands

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/fatih/structs"
	"github.com/ionos-cloud/ionosctl/pkg/config"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	"github.com/ionos-cloud/ionosctl/pkg/resources"
	"github.com/ionos-cloud/ionosctl/pkg/utils"
	"github.com/ionos-cloud/ionosctl/pkg/utils/clierror"
	"github.com/ionos-cloud/ionosctl/pkg/utils/printer"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func k8sNodePoolLan() *core.Command {
	ctx := context.TODO()
	k8sCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "lan",
			Short:            "Kubernetes NodePool LAN Operations",
			Long:             "The sub-commands of `ionosctl k8s nodepool lan` allow you to list, add, remove Kubernetes Node Pool LANs.",
			TraverseChildren: true,
		},
	}

	/*
		List Command
	*/
	list := core.NewCommand(ctx, k8sCmd, core.CommandBuilder{
		Namespace:  "nodepool",
		Resource:   "lan",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Kubernetes NodePool LANs",
		LongDesc:   "Use this command to get a list of all contained NodePool LANs in a selected Kubernetes Cluster.\n\nRequired values to run command:\n\n* K8s Cluster Id\n* K8s NodePool Id",
		Example:    listK8sNodePoolLanExample,
		PreCmdRun:  PreRunK8sClusterNodePoolIds,
		CmdRun:     RunK8sNodePoolLanList,
		InitClient: true,
	})
	list.AddStringFlag(config.ArgK8sClusterId, "", "", config.RequiredFlagK8sClusterId)
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgK8sClusterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getK8sClustersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddStringFlag(config.ArgK8sNodePoolId, "", "", config.RequiredFlagK8sNodePoolId)
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgK8sNodePoolId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getK8sNodePoolsIds(os.Stderr, viper.GetString(core.GetFlagName(list.NS, config.ArgK8sClusterId))), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddStringSliceFlag(config.ArgCols, "", defaultK8sNodePoolLanCols, utils.ColsMessage(defaultK8sNodePoolLanCols))
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultK8sNodePoolLanCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Add Command
	*/
	add := core.NewCommand(ctx, k8sCmd, core.CommandBuilder{
		Namespace: "nodepool",
		Resource:  "lan",
		Verb:      "add",
		Aliases:   []string{"a"},
		ShortDesc: "Add a Kubernetes NodePool LAN",
		LongDesc: `Use this command to add a Node Pool LAN into an existing Node Pool.

You can wait for the Node Pool to be in "ACTIVE" state using ` + "`" + `--wait-for-state` + "`" + ` flag together with ` + "`" + `--timeout` + "`" + ` option.

Required values to run a command:

* K8s Cluster Id
* K8s NodePool Id
* Lan Id`,
		Example:    addK8sNodePoolLanExample,
		PreCmdRun:  PreRunK8sClusterNodePoolLanIds,
		CmdRun:     RunK8sNodePoolLanAdd,
		InitClient: true,
	})
	add.AddStringFlag(config.ArgK8sClusterId, "", "", config.RequiredFlagK8sClusterId)
	_ = add.Command.RegisterFlagCompletionFunc(config.ArgK8sClusterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getK8sClustersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	add.AddStringFlag(config.ArgK8sNodePoolId, "", "", config.RequiredFlagK8sNodePoolId)
	_ = add.Command.RegisterFlagCompletionFunc(config.ArgK8sNodePoolId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getK8sNodePoolsIds(os.Stderr, viper.GetString(core.GetFlagName(add.NS, config.ArgK8sClusterId))), cobra.ShellCompDirectiveNoFileComp
	})
	add.AddIntFlag(config.ArgLanId, config.ArgIdShort, 0, "The unique LAN Id of existing LANs to be attached to worker Nodes "+config.RequiredFlag)
	add.AddBoolFlag(config.ArgDhcp, "", true, "Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP")
	add.AddStringFlag(config.ArgNetwork, "", "", "IPv4 or IPv6 CIDR to be routed via the interface. Must be set with --gateway-ip flag")
	add.AddStringFlag(config.ArgGatewayIp, "", "", "IPv4 or IPv6 Gateway IP for the route. Must be set with --network flag")
	add.AddStringSliceFlag(config.ArgCols, "", defaultK8sNodePoolLanCols, utils.ColsMessage(defaultK8sNodePoolLanCols))
	_ = add.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultK8sNodePoolLanCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Remove Command
	*/
	removeCmd := core.NewCommand(ctx, k8sCmd, core.CommandBuilder{
		Namespace: "nodepool",
		Resource:  "lan",
		Verb:      "remove",
		Aliases:   []string{"r"},
		ShortDesc: "Remove a Kubernetes NodePool LAN",
		LongDesc: `This command removes a Kubernetes Node Pool LAN from a Node Pool.

Required values to run command:

* K8s Cluster Id
* K8s NodePool Id
* Lan Id`,
		Example:    removeK8sNodePoolLanExample,
		PreCmdRun:  PreRunK8sClusterNodePoolLanIds,
		CmdRun:     RunK8sNodePoolLanRemove,
		InitClient: true,
	})
	removeCmd.AddStringFlag(config.ArgK8sClusterId, "", "", config.RequiredFlagK8sClusterId)
	_ = removeCmd.Command.RegisterFlagCompletionFunc(config.ArgK8sClusterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getK8sClustersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	removeCmd.AddStringFlag(config.ArgK8sNodePoolId, "", "", config.RequiredFlagK8sNodePoolId)
	_ = removeCmd.Command.RegisterFlagCompletionFunc(config.ArgK8sNodePoolId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getK8sNodePoolsIds(os.Stderr, viper.GetString(core.GetFlagName(removeCmd.NS, config.ArgK8sClusterId))), cobra.ShellCompDirectiveNoFileComp
	})
	removeCmd.AddIntFlag(config.ArgLanId, config.ArgIdShort, 0, "The unique LAN Id of existing LANs to be detached from worker Nodes "+config.RequiredFlag)

	return k8sCmd
}

func PreRunK8sClusterNodePoolLanIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.NS, config.ArgK8sClusterId, config.ArgK8sNodePoolId, config.ArgLanId)
}

func RunK8sNodePoolLanList(c *core.CommandConfig) error {
	k8ss, _, err := c.K8s().GetNodePool(
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sClusterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sNodePoolId)),
	)
	if err != nil {
		return err
	}
	if properties, ok := k8ss.GetPropertiesOk(); ok && properties != nil {
		if lans, ok := properties.GetLansOk(); ok && lans != nil {
			return c.Printer.Print(getK8sNodePoolLanPrint(c, getK8sNodePoolLans(lans)))
		} else {
			return errors.New("error getting node pool lans")
		}
	} else {
		return errors.New("error getting node pool properties")
	}
}

func RunK8sNodePoolLanAdd(c *core.CommandConfig) error {
	ng, _, err := c.K8s().GetNodePool(
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sClusterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sNodePoolId)),
	)
	if err != nil {
		return err
	}
	input := getNewK8sNodePoolLanInfo(c, ng)
	ngNew, _, err := c.K8s().UpdateNodePool(
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sClusterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sNodePoolId)),
		input,
	)
	if err != nil {
		return err
	}
	return c.Printer.Print(getK8sNodePoolLanPrint(c, getK8sNodePoolLansForPut(ngNew)))
}

func RunK8sNodePoolLanRemove(c *core.CommandConfig) error {
	if err := utils.AskForConfirm(c.Stdin, c.Printer, "remove node pool lan"); err != nil {
		return err
	}
	ng, _, err := c.K8s().GetNodePool(
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sClusterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sNodePoolId)),
	)
	if err != nil {
		return err
	}
	input := removeK8sNodePoolLanInfo(c, ng)
	_, _, err = c.K8s().UpdateNodePool(
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sClusterId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgK8sNodePoolId)),
		input,
	)
	if err != nil {
		return err
	}
	return c.Printer.Print("Status: Command node pool lan remove has been successfully executed")
}

func getNewK8sNodePoolLanInfo(c *core.CommandConfig, oldNg *resources.K8sNodePool) resources.K8sNodePool {
	propertiesUpdated := resources.K8sNodePoolProperties{}
	if properties, ok := oldNg.GetPropertiesOk(); ok && properties != nil {
		if n, ok := properties.GetNodeCountOk(); ok && n != nil {
			propertiesUpdated.SetNodeCount(*n)
		}
		if n, ok := properties.GetAutoScalingOk(); ok && n != nil {
			propertiesUpdated.SetAutoScaling(*n)
		}
		if n, ok := properties.GetMaintenanceWindowOk(); ok && n != nil {
			propertiesUpdated.SetMaintenanceWindow(*n)
		}
		if n, ok := properties.GetK8sVersionOk(); ok && n != nil {
			propertiesUpdated.SetK8sVersion(*n)
		}
		if viper.IsSet(core.GetFlagName(c.NS, config.ArgLanId)) {
			newLans := make([]ionoscloud.KubernetesNodePoolLan, 0)
			// Append existing LANs
			if existingLans, ok := properties.GetLansOk(); ok && existingLans != nil {
				for _, existingLan := range *existingLans {
					newLans = append(newLans, existingLan)
				}
			}
			// Add new LANs
			lanId := viper.GetInt32(core.GetFlagName(c.NS, config.ArgLanId))
			dhcp := viper.GetBool(core.GetFlagName(c.NS, config.ArgDhcp))
			newLan := ionoscloud.KubernetesNodePoolLan{
				Id:   &lanId,
				Dhcp: &dhcp,
			}
			if viper.IsSet(core.GetFlagName(c.NS, config.ArgNetwork)) &&
				viper.IsSet(core.GetFlagName(c.NS, config.ArgGatewayIp)) {
				newRoute := ionoscloud.KubernetesNodePoolLanRoutes{}
				if viper.IsSet(core.GetFlagName(c.NS, config.ArgNetwork)) {
					newRoute.SetNetwork(viper.GetString(core.GetFlagName(c.NS, config.ArgNetwork)))
				}
				if viper.IsSet(core.GetFlagName(c.NS, config.ArgGatewayIp)) {
					newRoute.SetGatewayIp(viper.GetString(core.GetFlagName(c.NS, config.ArgGatewayIp)))
				}
				newLan.SetRoutes([]ionoscloud.KubernetesNodePoolLanRoutes{newRoute})
			}
			newLans = append(newLans, newLan)
			propertiesUpdated.SetLans(newLans)
		}
	}
	return resources.K8sNodePool{
		KubernetesNodePool: ionoscloud.KubernetesNodePool{
			Properties: &propertiesUpdated.KubernetesNodePoolProperties,
		},
	}
}

func removeK8sNodePoolLanInfo(c *core.CommandConfig, oldNg *resources.K8sNodePool) resources.K8sNodePool {
	propertiesUpdated := resources.K8sNodePoolProperties{}
	if properties, ok := oldNg.GetPropertiesOk(); ok && properties != nil {
		if n, ok := properties.GetNodeCountOk(); ok && n != nil {
			propertiesUpdated.SetNodeCount(*n)
		}
		if n, ok := properties.GetAutoScalingOk(); ok && n != nil {
			propertiesUpdated.SetAutoScaling(*n)
		}
		if n, ok := properties.GetMaintenanceWindowOk(); ok && n != nil {
			propertiesUpdated.SetMaintenanceWindow(*n)
		}
		if n, ok := properties.GetK8sVersionOk(); ok && n != nil {
			propertiesUpdated.SetK8sVersion(*n)
		}
		if viper.IsSet(core.GetFlagName(c.NS, config.ArgLanId)) {
			lanId := viper.GetInt32(core.GetFlagName(c.NS, config.ArgLanId))
			newLans := make([]ionoscloud.KubernetesNodePoolLan, 0)
			if existingLans, ok := properties.GetLansOk(); ok && existingLans != nil {
				for _, existingLan := range *existingLans {
					if id, ok := existingLan.GetIdOk(); ok && id != nil {
						if *id != lanId {
							newLans = append(newLans, existingLan)
						}
					}
				}
			}
			propertiesUpdated.SetLans(newLans)
		}
	}
	return resources.K8sNodePool{
		KubernetesNodePool: ionoscloud.KubernetesNodePool{
			Properties: &propertiesUpdated.KubernetesNodePoolProperties,
		},
	}
}

// Output Printing

var defaultK8sNodePoolLanCols = []string{"LanId", "Dhcp", "RoutesNetwork", "RoutesGatewayIp"}

type K8sNodePoolLanPrint struct {
	LanId           int32    `json:"LanId,omitempty"`
	Dhcp            bool     `json:"Dhcp,omitempty"`
	RoutesNetwork   []string `json:"RoutesNetwork,omitempty"`
	RoutesGatewayIp []string `json:"RoutesGatewayIp,omitempty"`
}

func getK8sNodePoolLanPrint(c *core.CommandConfig, k8ss []resources.K8sNodePoolLan) printer.Result {
	r := printer.Result{}
	if c != nil {
		if k8ss != nil {
			r.OutputJSON = k8ss
			r.KeyValue = getK8sNodePoolLansKVMaps(k8ss)
			r.Columns = getK8sNodePoolLanCols(core.GetGlobalFlagName(c.Resource, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getK8sNodePoolLansForPut(ng *resources.K8sNodePoolForPut) []resources.K8sNodePoolLan {
	ss := make([]resources.K8sNodePoolLan, 0)
	if ng != nil {
		if properties, ok := ng.GetPropertiesOk(); ok && properties != nil {
			if lans, ok := properties.GetLansOk(); ok && lans != nil {
				for _, lanItem := range *lans {
					ss = append(ss, resources.K8sNodePoolLan{
						KubernetesNodePoolLan: lanItem,
					})
				}
			}
		}
	}
	return ss
}

func getK8sNodePoolLanCols(flagName string, outErr io.Writer) []string {
	if viper.IsSet(flagName) {
		var k8sCols []string
		columnsMap := map[string]string{
			"LanId":           "LanId",
			"Dhcp":            "Dhcp",
			"RoutesNetwork":   "RoutesNetwork",
			"RoutesGatewayIp": "RoutesGatewayIp",
		}
		for _, k := range viper.GetStringSlice(flagName) {
			col := columnsMap[k]
			if col != "" {
				k8sCols = append(k8sCols, col)
			} else {
				clierror.CheckError(errors.New("unknown column "+k), outErr)
			}
		}
		return k8sCols
	} else {
		return defaultK8sNodePoolLanCols
	}
}

func getK8sNodePoolLans(k8ss *[]ionoscloud.KubernetesNodePoolLan) []resources.K8sNodePoolLan {
	u := make([]resources.K8sNodePoolLan, 0)
	if k8ss != nil {
		for _, item := range *k8ss {
			u = append(u, resources.K8sNodePoolLan{KubernetesNodePoolLan: item})
		}
	}
	return u
}

func getK8sNodePoolLansKVMaps(us []resources.K8sNodePoolLan) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(us))
	for _, u := range us {
		var uPrint K8sNodePoolLanPrint
		if id, ok := u.GetIdOk(); ok && id != nil {
			uPrint.LanId = *id
		}
		if dhcp, ok := u.GetDhcpOk(); ok && dhcp != nil {
			uPrint.Dhcp = *dhcp
		}
		if routes, ok := u.GetRoutesOk(); ok && routes != nil {
			newRoutesNetwork := make([]string, 0)
			newRoutesGatewayIp := make([]string, 0)
			for _, route := range *routes {
				if net, ok := route.GetNetworkOk(); ok && net != nil {
					newRoutesNetwork = append(newRoutesNetwork, *net)
				}
				if ip, ok := route.GetGatewayIpOk(); ok && ip != nil {
					newRoutesGatewayIp = append(newRoutesGatewayIp, *ip)
				}
			}
			uPrint.RoutesNetwork = newRoutesNetwork
			uPrint.RoutesGatewayIp = newRoutesGatewayIp
		}
		o := structs.Map(uPrint)
		out = append(out, o)
	}
	return out
}