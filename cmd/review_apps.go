package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/Scalingo/cli/reviewapps"

	"github.com/Scalingo/cli/cmd/autocomplete"
	"github.com/Scalingo/cli/detect"
)

var (
	reviewAppsShowCommand = cli.Command{
		Name:     "review-apps",
		Category: "Review Apps",
		Flags:    []cli.Flag{&appFlag},
		Usage:    "Show review apps of the parent application",
		Description: CommandDescription{
			Description: "Show review apps of the parent application",
			Examples:    []string{"scalingo --app my-app review-apps"},
		}.Render(),
		Action: func(c *cli.Context) error {
			if c.Args().Len() != 0 {
				cli.ShowCommandHelp(c, "review-apps")
				return nil
			}

			currentApp := detect.CurrentApp(c)
			err := reviewapps.Show(c.Context, currentApp)
			if err != nil {
				errorQuit(c.Context, err)
			}
			return nil
		},
		BashComplete: func(c *cli.Context) {
			_ = autocomplete.CmdFlagsAutoComplete(c, "review-apps")
		},
	}
)
