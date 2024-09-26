package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type StepTemplateUpdate struct {
	root  *Otto
	Quiet bool `usage:"Only print IDs of updated step template" short:"q"`
}

func (l *StepTemplateUpdate) Customize(cmd *cobra.Command) {
	cmd.Use = "update [flags] NAME REFERENCE"
	cmd.Args = cobra.ExactArgs(2)
}

func (l *StepTemplateUpdate) Run(cmd *cobra.Command, args []string) error {
	tr, err := l.root.Client.UpdateToolReference(cmd.Context(), args[0], args[1])
	if err != nil {
		return err
	}
	if l.Quiet {
		fmt.Println(tr.ID)
	} else {
		fmt.Println("Step template updated:", tr.ID)
	}
	return nil
}