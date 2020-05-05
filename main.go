package main

import (
	"io/ioutil"
	"log"
	"os"

	cmd "thy/commands"
	cst "thy/constants"
	"thy/format"
	"thy/utils"
	"thy/version"

	"github.com/thycotic-rd/cli"
)

func main() {
	exitStatus, err := runCLI(os.Args)
	if err != nil {
		format.NewDefaultOutClient().Fail(err)
		os.Exit(utils.GetExecStatus(err))
	}
	os.Exit(exitStatus)
}

func runCLI(args []string) (exitStatus int, err error) {
	c := cli.NewCLI(cst.CmdRoot, version.Version)
	c.Args = args[1:]
	c.Name = cst.CmdRoot
	c.Commands = map[string]cli.CommandFactory{
		"secret":                        cmd.GetSecretCmd,
		"secret read":                   cmd.GetReadCmd,
		"secret describe":               cmd.GetDescribeCmd,
		"secret search":                 cmd.GetSecretSearchCmd,
		"secret delete":                 cmd.GetDeleteCmd,
		"secret restore":                cmd.GetSecretRestoreCmd,
		"secret create":                 cmd.GetCreateCmd,
		"secret update":                 cmd.GetUpdateCmd,
		"secret rollback":               cmd.GetRollbackCmd,
		"secret edit":                   cmd.GetEditCmd,
		"secret bustcache":              cmd.GetBustCacheCmd,
		"policy":                        cmd.GetPolicyCmd,
		"policy read":                   cmd.GetPolicyReadCmd,
		"policy search":                 cmd.GetPolicySearchCommand,
		"policy delete":                 cmd.GetPolicyDeleteCmd,
		"policy restore":                cmd.GetPolicyRestoreCmd,
		"policy create":                 cmd.GetPolicyCreateCmd,
		"policy edit":                   cmd.GetPolicyEditCmd,
		"policy update":                 cmd.GetPolicyUpdateCmd,
		"policy rollback":               cmd.GetPolicyRollbackCmd,
		"auth":                          cmd.GetAuthCmd,
		"auth clear":                    cmd.GetAuthClearCmd,
		"auth list":                     cmd.GetAuthListCmd,
		"auth change-password":          cmd.GetAuthChangePasswordCmd,
		"user":                          cmd.GetUserCmd,
		"user read":                     cmd.GetUserReadCmd,
		"user search":                   cmd.GetUserSearchCmd,
		"user delete":                   cmd.GetUserDeleteCmd,
		"user restore":                  cmd.GetUserRestoreCmd,
		"user create":                   cmd.GetUserCreateCmd,
		"whoami":                        cmd.GetWhoAmICmd,
		"eval":                          cmd.GetEvaluateFlagCmd,
		"cli-config":                    cmd.GetCliConfigCmd,
		"init":                          cmd.GetCliConfigInitCmd,
		"cli-config init":               cmd.GetCliConfigInitCmd,
		"cli-config update":             cmd.GetCliConfigUpdateCmd,
		"cli-config clear":              cmd.GetCliConfigClearCmd,
		"cli-config read":               cmd.GetCliConfigReadCmd,
		"cli-config edit":               cmd.GetCliConfigEditCmd,
		"config":                        cmd.GetConfigCmd,
		"config read":                   cmd.GetConfigReadCmd,
		"config update":                 cmd.GetConfigUpdateCmd,
		"config edit":                   cmd.GetConfigEditCmd,
		"config auth-provider":          cmd.GetAuthProviderCmd,
		"config auth-provider read":     cmd.GetAuthProviderReadCmd,
		"config auth-provider search":   cmd.GetAuthProviderSearchCommand,
		"config auth-provider delete":   cmd.GetAuthProviderDeleteCmd,
		"config auth-provider restore":  cmd.GetAuthProviderRestoreCmd,
		"config auth-provider create":   cmd.GetAuthProviderCreateCmd,
		"config auth-provider update":   cmd.GetAuthProviderUpdateCmd,
		"config auth-provider rollback": cmd.GetAuthProviderRollbackCmd,
		"role":                          cmd.GetRoleCmd,
		"role read":                     cmd.GetRoleReadCmd,
		"role search":                   cmd.GetRoleSearchCmd,
		"role delete":                   cmd.GetRoleDeleteCmd,
		"role restore":                  cmd.GetRoleRestoreCmd,
		"role create":                   cmd.GetRoleCreateCmd,
		"role update":                   cmd.GetRoleUpdateCmd,
		"client":                        cmd.GetClientCmd,
		"client read":                   cmd.GetClientReadCmd,
		"client delete":                 cmd.GetClientDeleteCmd,
		"client restore":                cmd.GetClientRestoreCmd,
		"client create":                 cmd.GetClientCreateCmd,
		"client search":                 cmd.GetClientSearchCmd,
		"usage":                         cmd.GetUsageCmd,
		"logs":                          cmd.GetLogsSearchCmd,
		"audit":                         cmd.GetAuditSearchCmd,
		"group":                         cmd.GetGroupCmd,
		"group read":                    cmd.GetGroupReadCmd,
		"group create":                  cmd.GetGroupCreateCmd,
		"group delete":                  cmd.GetGroupDeleteCmd,
		"group restore":                 cmd.GetGroupRestoreCmd,
		"group search":                  cmd.GetGroupSearchCmd,
		"group add-members":             cmd.GetAddMembersCmd,
		"user groups":                   cmd.GetMemberGroupsCmd,
		"group delete-members":          cmd.GetDeleteMembersCmd,
		"pki":                           cmd.GetPkiCmd,
		"pki register":                  cmd.GetPkiRegisterCmd,
		"pki sign":                      cmd.GetPkiSignCmd,
		"pki leaf":                      cmd.GetPkiLeafCmd,
		"pki generate-root":             cmd.GetPkiGenerateRootCmd,
	}

	c.Autocomplete = true
	c.AutocompleteInstall = "install"
	c.AutocompleteUninstall = "uninstall"
	log.SetOutput(ioutil.Discard)
	return c.Run()
}