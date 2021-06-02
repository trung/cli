package admin

import (
	"github.com/smallstep/cli/errs"
	"github.com/smallstep/cli/flags"
	"github.com/smallstep/cli/utils/cautils"
	"github.com/urfave/cli"
)

func removeCommand() cli.Command {
	return cli.Command{
		Name:   "remove",
		Action: cli.ActionFunc(removeAction),
		Usage:  "remove an admin from the CA configuration",
		UsageText: `**step ca admin remove** <subject> [**--provisioner**=<id>] [**--ca-url**=<uri>]
[**--root**=<file>]`,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "provisioner",
				Usage: `Filter admins by provisioner name.`,
			},
			flags.X5cCert,
			flags.X5cKey,
			flags.PasswordFile,
			flags.CaURL,
			flags.Root,
		},
		Description: `**step ca admin remove** removes an admin from the CA configuration.

## POSITIONAL ARGUMENTS

<name>
: The name of the admin to be removed.

## EXAMPLES

Remove an admin:
'''
$ step ca admin remove max@smallstep.com
'''

Remove an admin with additional filtering by provisioner:
'''
$ step ca admin remove max@smallstep.com --provisioner admin-jwk
'''
`,
	}
}

func removeAction(ctx *cli.Context) error {
	if err := errs.NumberOfArguments(ctx, 1); err != nil {
		return err
	}

	client, err := cautils.NewAdminClient(ctx)
	if err != nil {
		return err
	}

	admins, err := client.GetAdmins()
	if err != nil {
		return err
	}
	adm, err := adminPrompt(ctx, client, admins)
	if err != nil {
		return err
	}

	if err := client.RemoveAdmin(adm.Id); err != nil {
		return err
	}

	return nil
}