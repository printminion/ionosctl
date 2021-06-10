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
	ionoscloud "github.com/ionos-cloud/sdk-go/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	multierror "go.uber.org/multierr"
)

func firewallrule() *core.Command {
	ctx := context.TODO()
	firewallRuleCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "firewallrule",
			Aliases:          []string{"f", "fr", "firewall"},
			Short:            "Firewall Rule Operations",
			Long:             `The sub-commands of ` + "`" + `ionosctl firewallrule` + "`" + ` allow you to create, list, get, update, delete Firewall Rules.`,
			TraverseChildren: true,
		},
	}
	globalFlags := firewallRuleCmd.GlobalFlags()
	globalFlags.StringP(config.ArgDataCenterId, "", "", config.RequiredFlagDatacenterId)
	_ = viper.BindPFlag(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgDataCenterId), globalFlags.Lookup(config.ArgDataCenterId))
	_ = firewallRuleCmd.Command.RegisterFlagCompletionFunc(config.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getDataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	globalFlags.StringP(config.ArgServerId, "", "", config.RequiredFlagServerId)
	_ = viper.BindPFlag(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgServerId), globalFlags.Lookup(config.ArgServerId))
	_ = firewallRuleCmd.Command.RegisterFlagCompletionFunc(config.ArgServerId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getServersIds(os.Stderr, viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	globalFlags.StringP(config.ArgNicId, "", "", config.RequiredFlagNicId)
	_ = viper.BindPFlag(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgNicId), globalFlags.Lookup(config.ArgNicId))
	_ = firewallRuleCmd.Command.RegisterFlagCompletionFunc(config.ArgNicId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getNicsIds(os.Stderr,
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgDataCenterId)),
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgServerId)),
		), cobra.ShellCompDirectiveNoFileComp
	})
	globalFlags.StringSliceP(config.ArgCols, "", defaultFirewallRuleCols, utils.ColsMessage(allFirewallRuleCols))
	_ = viper.BindPFlag(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgCols), globalFlags.Lookup(config.ArgCols))
	_ = firewallRuleCmd.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allFirewallRuleCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	core.NewCommand(ctx, firewallRuleCmd, core.CommandBuilder{
		Namespace:  "firewallrule",
		Resource:   "firewallrule",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Firewall Rules",
		LongDesc:   "Use this command to get a list of Firewall Rules from a specified NIC from a Server.\n\nRequired values to run command:\n\n* Data Center Id\n* Server Id\n*Nic Id",
		Example:    listFirewallRuleExample,
		PreCmdRun:  PreRunGlobalDcServerNicIds,
		CmdRun:     RunFirewallRuleList,
		InitClient: true,
	})

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, firewallRuleCmd, core.CommandBuilder{
		Namespace:  "firewallrule",
		Resource:   "firewallrule",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get a Firewall Rule",
		LongDesc:   "Use this command to retrieve information of a specified Firewall Rule.\n\nRequired values to run command:\n\n* Data Center Id\n* Server Id\n*Nic Id\n* FirewallRule Id",
		Example:    getFirewallRuleExample,
		PreCmdRun:  PreRunGlobalDcServerNicIdsFRuleId,
		CmdRun:     RunFirewallRuleGet,
		InitClient: true,
	})
	get.AddStringFlag(config.ArgFirewallRuleId, config.ArgIdShort, "", config.RequiredFlagFirewallRuleId)
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgFirewallRuleId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getFirewallRulesIds(os.Stderr,
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgDataCenterId)),
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgServerId)),
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgNicId))), cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Create Command
	*/
	create := core.NewCommand(ctx, firewallRuleCmd, core.CommandBuilder{
		Namespace: "firewallrule",
		Resource:  "firewallrule",
		Verb:      "create",
		Aliases:   []string{"c"},
		ShortDesc: "Create a Firewall Rule",
		LongDesc: `Use this command to create/add a new Firewall Rule to the specified NIC. All Firewall Rules must be associated with a NIC.

NOTE: the Firewall Rule Protocol can only be set when creating a new Firewall Rule.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option.

Required values to run command:

* Data Center Id
* Server Id
* Nic Id 
* Protocol`,
		Example:    createFirewallRuleExample,
		PreCmdRun:  PreRunGlobalDcServerNicIdsFRuleProtocol,
		CmdRun:     RunFirewallRuleCreate,
		InitClient: true,
	})
	create.AddStringFlag(config.ArgName, config.ArgNameShort, "", "The name for the Firewall Rule")
	create.AddStringFlag(config.ArgProtocol, "", "", "The Protocol for Firewall Rule: TCP, UDP, ICMP, ANY "+config.RequiredFlag)
	create.AddStringFlag(config.ArgSourceMac, "", "", "Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Unset option allows all source MAC addresses.")
	create.AddStringFlag(config.ArgSourceIp, "", "", "Only traffic originating from the respective IPv4 address is allowed. Not setting option allows all source IPs.")
	create.AddStringFlag(config.ArgTargetIp, "", "", "In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Not setting option allows all target IPs.")
	create.AddIntFlag(config.ArgIcmpType, "", 0, "Define the allowed type (from 0 to 254) if the protocol ICMP is chosen. Not setting option allows all types.")
	create.AddIntFlag(config.ArgIcmpCode, "", 0, "Define the allowed code (from 0 to 254) if protocol ICMP is chosen. Not setting option allows all codes.")
	create.AddIntFlag(config.ArgPortRangeStart, "", 1, "Define the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Not setting portRangeStart and portRangeEnd allows all ports.")
	create.AddIntFlag(config.ArgPortRangeStop, "", 1, "Define the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Not setting portRangeStart and portRangeEnd allows all ports.")
	create.AddBoolFlag(config.ArgWaitForRequest, config.ArgWaitForRequestShort, config.DefaultWait, "Wait for Request for Firewall Rule creation to be executed")
	create.AddIntFlag(config.ArgTimeout, config.ArgTimeoutShort, config.DefaultTimeoutSeconds, "Timeout option for Request for Firewall Rule creation [seconds]")

	/*
		Update Command
	*/
	update := core.NewCommand(ctx, firewallRuleCmd, core.CommandBuilder{
		Namespace: "firewallrule",
		Resource:  "firewallrule",
		Verb:      "update",
		Aliases:   []string{"u", "up"},
		ShortDesc: "Update a FirewallRule",
		LongDesc: `Use this command to update a specified Firewall Rule.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option.

Required values to run command:

* Data Center Id
* Server Id
* Nic Id
* Firewall Rule Id`,
		Example:    updateFirewallRuleExample,
		PreCmdRun:  PreRunGlobalDcServerNicIdsFRuleId,
		CmdRun:     RunFirewallRuleUpdate,
		InitClient: true,
	})
	update.AddStringFlag(config.ArgName, config.ArgNameShort, "", "The name for the Firewall Rule")
	update.AddStringFlag(config.ArgSourceMac, "", "", "Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Not setting option allows all source MAC addresses.")
	update.AddStringFlag(config.ArgSourceIp, "", "", "Only traffic originating from the respective IPv4 address is allowed. Not setting option allows all source IPs.")
	update.AddStringFlag(config.ArgTargetIp, "", "", "In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Not setting option allows all target IPs.")
	update.AddIntFlag(config.ArgIcmpType, "", 0, "Redefine the allowed type (from 0 to 254) if the protocol ICMP is chosen. Not setting option allows all types.")
	update.AddIntFlag(config.ArgIcmpCode, "", 0, "Redefine the allowed code (from 0 to 254) if protocol ICMP is chosen. Not setting option allows all codes.")
	update.AddIntFlag(config.ArgPortRangeStart, "", 1, "Redefine the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Not setting portRangeStart and portRangeEnd allows all ports.")
	update.AddIntFlag(config.ArgPortRangeStop, "", 1, "Redefine the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Not setting portRangeStart and portRangeEnd allows all ports.")
	update.AddStringFlag(config.ArgFirewallRuleId, config.ArgIdShort, "", config.RequiredFlagFirewallRuleId)
	_ = update.Command.RegisterFlagCompletionFunc(config.ArgFirewallRuleId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getFirewallRulesIds(os.Stderr,
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgDataCenterId)),
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgServerId)),
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgNicId))), cobra.ShellCompDirectiveNoFileComp
	})
	update.AddBoolFlag(config.ArgWaitForRequest, config.ArgWaitForRequestShort, config.DefaultWait, "Wait for Request for Firewall Rule update to be executed")
	update.AddIntFlag(config.ArgTimeout, config.ArgTimeoutShort, config.DefaultTimeoutSeconds, "Timeout option for Request for Firewall Rule update [seconds]")

	/*
		Delete Command
	*/
	deleteCmd := core.NewCommand(ctx, firewallRuleCmd, core.CommandBuilder{
		Namespace: "firewallrule",
		Resource:  "firewallrule",
		Verb:      "delete",
		Aliases:   []string{"d"},
		ShortDesc: "Delete a FirewallRule",
		LongDesc: `Use this command to delete a specified Firewall Rule from a Virtual Data Center.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option. You can force the command to execute without user input using ` + "`" + `--force` + "`" + ` option.

Required values to run command:

* Data Center Id
* Server Id
* Nic Id
* Firewall Rule Id`,
		Example:    deleteFirewallRuleExample,
		PreCmdRun:  PreRunGlobalDcServerNicIdsFRuleId,
		CmdRun:     RunFirewallRuleDelete,
		InitClient: true,
	})
	deleteCmd.AddStringFlag(config.ArgFirewallRuleId, config.ArgIdShort, "", config.RequiredFlagFirewallRuleId)
	_ = deleteCmd.Command.RegisterFlagCompletionFunc(config.ArgFirewallRuleId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getFirewallRulesIds(os.Stderr,
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgDataCenterId)),
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgServerId)),
			viper.GetString(core.GetGlobalFlagName(firewallRuleCmd.Name(), config.ArgNicId))), cobra.ShellCompDirectiveNoFileComp
	})
	deleteCmd.AddBoolFlag(config.ArgWaitForRequest, config.ArgWaitForRequestShort, config.DefaultWait, "Wait for Request for Firewall Rule deletion to be executed")
	deleteCmd.AddIntFlag(config.ArgTimeout, config.ArgTimeoutShort, config.DefaultTimeoutSeconds, "Timeout option for Request for Firewall Rule deletion [seconds]")

	return firewallRuleCmd
}

func PreRunGlobalDcServerNicIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredGlobalFlags(c.Resource, config.ArgDataCenterId, config.ArgServerId, config.ArgNicId)
}

func PreRunGlobalDcServerNicIdsFRuleProtocol(c *core.PreCommandConfig) error {
	var result error
	if err := core.CheckRequiredGlobalFlags(c.Resource, config.ArgDataCenterId, config.ArgServerId, config.ArgNicId); err != nil {
		result = multierror.Append(result, err)
	}
	if err := core.CheckRequiredFlags(c.NS, config.ArgProtocol); err != nil {
		result = multierror.Append(result, err)
	}
	if result != nil {
		return result
	}
	return nil
}

func PreRunGlobalDcServerNicIdsFRuleId(c *core.PreCommandConfig) error {
	var result error
	if err := core.CheckRequiredGlobalFlags(c.Resource, config.ArgDataCenterId, config.ArgServerId, config.ArgNicId); err != nil {
		result = multierror.Append(result, err)
	}
	if err := core.CheckRequiredFlags(c.NS, config.ArgFirewallRuleId); err != nil {
		result = multierror.Append(result, err)
	}
	if result != nil {
		return result
	}
	return nil
}

func RunFirewallRuleList(c *core.CommandConfig) error {
	firewallRules, _, err := c.FirewallRules().List(
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgDataCenterId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgServerId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgNicId)),
	)
	if err != nil {
		return err
	}
	return c.Printer.Print(getFirewallRulePrint(nil, c, getFirewallRules(firewallRules)))
}

func RunFirewallRuleGet(c *core.CommandConfig) error {
	firewallRule, _, err := c.FirewallRules().Get(
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgDataCenterId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgServerId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgNicId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgFirewallRuleId)),
	)
	if err != nil {
		return err
	}
	return c.Printer.Print(getFirewallRulePrint(nil, c, getFirewallRule(firewallRule)))
}

func RunFirewallRuleCreate(c *core.CommandConfig) error {
	properties := getFirewallRulePropertiesSet(c)
	input := resources.FirewallRule{
		FirewallRule: ionoscloud.FirewallRule{
			Properties: &properties.FirewallruleProperties,
		},
	}
	firewallRule, resp, err := c.FirewallRules().Create(
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgDataCenterId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgServerId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgNicId)),
		input,
	)
	if err != nil {
		return err
	}

	if err = utils.WaitForRequest(c, printer.GetRequestPath(resp)); err != nil {
		return err
	}
	return c.Printer.Print(getFirewallRulePrint(resp, c, getFirewallRule(firewallRule)))
}

func RunFirewallRuleUpdate(c *core.CommandConfig) error {
	firewallRule, resp, err := c.FirewallRules().Update(
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgDataCenterId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgServerId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgNicId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgFirewallRuleId)),
		getFirewallRulePropertiesSet(c),
	)
	if err != nil {
		return err
	}

	if err = utils.WaitForRequest(c, printer.GetRequestPath(resp)); err != nil {
		return err
	}
	return c.Printer.Print(getFirewallRulePrint(resp, c, getFirewallRule(firewallRule)))
}

func RunFirewallRuleDelete(c *core.CommandConfig) error {
	if err := utils.AskForConfirm(c.Stdin, c.Printer, "delete firewall rule"); err != nil {
		return err
	}
	resp, err := c.FirewallRules().Delete(
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgDataCenterId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgServerId)),
		viper.GetString(core.GetGlobalFlagName(c.Resource, config.ArgNicId)),
		viper.GetString(core.GetFlagName(c.NS, config.ArgFirewallRuleId)),
	)
	if err != nil {
		return err
	}

	if err = utils.WaitForRequest(c, printer.GetRequestPath(resp)); err != nil {
		return err
	}
	return c.Printer.Print(getFirewallRulePrint(resp, c, nil))
}

// Get Firewall Rule Properties set used for create and update commands
func getFirewallRulePropertiesSet(c *core.CommandConfig) resources.FirewallRuleProperties {
	properties := resources.FirewallRuleProperties{}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgName)) {
		properties.SetName(viper.GetString(core.GetFlagName(c.NS, config.ArgName)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgProtocol)) {
		properties.SetProtocol(viper.GetString(core.GetFlagName(c.NS, config.ArgProtocol)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgSourceIp)) {
		properties.SetSourceIp(viper.GetString(core.GetFlagName(c.NS, config.ArgSourceIp)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgSourceMac)) {
		properties.SetSourceMac(viper.GetString(core.GetFlagName(c.NS, config.ArgSourceMac)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgTargetIp)) {
		properties.SetTargetIp(viper.GetString(core.GetFlagName(c.NS, config.ArgTargetIp)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgIcmpCode)) {
		properties.SetIcmpCode(viper.GetInt32(core.GetFlagName(c.NS, config.ArgIcmpCode)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgIcmpType)) {
		properties.SetIcmpType(viper.GetInt32(core.GetFlagName(c.NS, config.ArgIcmpType)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgPortRangeStart)) {
		properties.SetPortRangeStart(viper.GetInt32(core.GetFlagName(c.NS, config.ArgPortRangeStart)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgPortRangeStop)) {
		properties.SetPortRangeEnd(viper.GetInt32(core.GetFlagName(c.NS, config.ArgPortRangeStop)))
	}
	return properties
}

// Output Printing

var (
	defaultFirewallRuleCols = []string{"FirewallRuleId", "Name", "Protocol", "PortRangeStart", "PortRangeEnd", "State"}
	allFirewallRuleCols     = []string{"FirewallRuleId", "Name", "Protocol", "SourceMac", "SourceIP", "TargetIP", "PortRangeStart", "PortRangeEnd", "State"}
)

type FirewallRulePrint struct {
	FirewallRuleId string `json:"FirewallRuleId,omitempty"`
	Name           string `json:"Name,omitempty"`
	Protocol       string `json:"Protocol,omitempty"`
	SourceMac      string `json:"SourceMac,omitempty"`
	SourceIP       string `json:"SourceIP,omitempty"`
	TargetIP       string `json:"TargetIP,omitempty"`
	PortRangeStart int32  `json:"PortRangeStart,omitempty"`
	PortRangeEnd   int32  `json:"PortRangeEnd,omitempty"`
	State          string `json:"State,omitempty"`
}

func getFirewallRulePrint(resp *resources.Response, c *core.CommandConfig, rule []resources.FirewallRule) printer.Result {
	var r printer.Result
	if c != nil {
		if resp != nil {
			r.ApiResponse = resp
			r.Resource = c.Resource
			r.Verb = c.Verb
			r.WaitForRequest = viper.GetBool(core.GetFlagName(c.NS, config.ArgWaitForRequest))
		}
		if rule != nil {
			r.OutputJSON = rule
			r.KeyValue = getFirewallRulesKVMaps(rule)
			r.Columns = getFirewallRulesCols(core.GetGlobalFlagName(c.Resource, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getFirewallRulesCols(flagName string, outErr io.Writer) []string {
	if viper.IsSet(flagName) && len(viper.GetStringSlice(flagName)) > 0 {
		var firewallRuleCols []string
		columnsMap := map[string]string{
			"FirewallRuleId": "FirewallRuleId",
			"Name":           "Name",
			"Protocol":       "Protocol",
			"SourceMac":      "SourceMac",
			"SourceIP":       "SourceIP",
			"TargetIP":       "TargetIP",
			"PortRangeStart": "PortRangeStart",
			"PortRangeEnd":   "PortRangeEnd",
			"State":          "State",
		}
		for _, k := range viper.GetStringSlice(flagName) {
			col := columnsMap[k]
			if col != "" {
				firewallRuleCols = append(firewallRuleCols, col)
			} else {
				clierror.CheckError(errors.New("unknown column "+k), outErr)
			}
		}
		return firewallRuleCols
	} else {
		return defaultFirewallRuleCols
	}
}

func getFirewallRules(firewallRules resources.FirewallRules) []resources.FirewallRule {
	ls := make([]resources.FirewallRule, 0)
	if items, ok := firewallRules.GetItemsOk(); ok && items != nil {
		for _, s := range *items {
			ls = append(ls, resources.FirewallRule{FirewallRule: s})
		}
	}
	return ls
}

func getFirewallRule(firewallRule *resources.FirewallRule) []resources.FirewallRule {
	ss := make([]resources.FirewallRule, 0)
	if firewallRule != nil {
		ss = append(ss, resources.FirewallRule{FirewallRule: firewallRule.FirewallRule})
	}
	return ss
}

func getFirewallRulesKVMaps(ls []resources.FirewallRule) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(ls))
	if len(ls) > 0 {
		for _, l := range ls {
			o := getFirewallRuleKVMap(l)
			out = append(out, o)
		}
	}
	return out
}

func getFirewallRuleKVMap(l resources.FirewallRule) map[string]interface{} {
	var firewallRulePrint FirewallRulePrint
	if id, ok := l.GetIdOk(); ok && id != nil {
		firewallRulePrint.FirewallRuleId = *id
	}
	if properties, ok := l.GetPropertiesOk(); ok && properties != nil {
		if name, ok := properties.GetNameOk(); ok && name != nil {
			firewallRulePrint.Name = *name
		}
		if protocol, ok := properties.GetProtocolOk(); ok && protocol != nil {
			firewallRulePrint.Protocol = *protocol
		}
		if portRangeStart, ok := properties.GetPortRangeStartOk(); ok && portRangeStart != nil {
			firewallRulePrint.PortRangeStart = *portRangeStart
		}
		if portRangeEnd, ok := properties.GetPortRangeEndOk(); ok && portRangeEnd != nil {
			firewallRulePrint.PortRangeEnd = *portRangeEnd
		}
		if sourceMac, ok := properties.GetSourceMacOk(); ok && sourceMac != nil {
			firewallRulePrint.SourceMac = *sourceMac
		}
		if sourceIp, ok := properties.GetSourceIpOk(); ok && sourceIp != nil {
			firewallRulePrint.SourceIP = *sourceIp
		}
		if targetIp, ok := properties.GetTargetIpOk(); ok && targetIp != nil {
			firewallRulePrint.TargetIP = *targetIp
		}
	}
	if metadata, ok := l.GetMetadataOk(); ok && metadata != nil {
		if state, ok := metadata.GetStateOk(); ok && state != nil {
			firewallRulePrint.State = *state
		}
	}
	return structs.Map(firewallRulePrint)
}

func getFirewallRulesIds(outErr io.Writer, datacenterId, serverId, nicId string) []string {
	err := config.Load()
	clierror.CheckError(err, outErr)
	clientSvc, err := resources.NewClientService(
		viper.GetString(config.Username),
		viper.GetString(config.Password),
		viper.GetString(config.Token),
		viper.GetString(config.ArgServerUrl),
	)
	clierror.CheckError(err, outErr)
	firewallRuleSvc := resources.NewFirewallRuleService(clientSvc.Get(), context.TODO())
	firewallRules, _, err := firewallRuleSvc.List(datacenterId, serverId, nicId)
	clierror.CheckError(err, outErr)
	firewallRulesIds := make([]string, 0)
	if items, ok := firewallRules.FirewallRules.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			if itemId, ok := item.GetIdOk(); ok && itemId != nil {
				firewallRulesIds = append(firewallRulesIds, *itemId)
			}
		}
	} else {
		return nil
	}
	return firewallRulesIds
}
