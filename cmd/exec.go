package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/ramiawar/superpet/config"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:     "exec",
	Aliases: []string{"x"},
	Short:   "Run the selected commands",
	Long:    `Run the selected commands directly`,
	RunE:    execute,
}

func execute(cmd *cobra.Command, args []string) (err error) {
	flag := config.Flag

	var options []string
	if len(args) != 0 {
		options = append(options, fmt.Sprintf("--query %s -1", strings.Join(args, " ")))
	}
	if config.Conf.General.SelectCmd == "fzf" {
		options = append(options, "--ansi")
		options = append(options, "--cycle")
		options = append(options, "-m")
	}

	commands, err := filter(options, flag.FilterTag)
	if err != nil {
		return err
	}
	command := strings.Join(commands, "; ")
	if config.Flag.Show {
		fmt.Printf("%s: %s\n", color.YellowString("Command"), command)
	}
	return run(command, os.Stdin, os.Stdout)
}

func init() {
	RootCmd.AddCommand(execCmd)
	execCmd.Flags().BoolVarP(&config.Flag.Color, "color", "", true,
		`Enable colorized output (only fzf)`)
	execCmd.Flags().StringVarP(&config.Flag.FilterTag, "tag", "t", "",
		`Filter tag`)
	execCmd.Flags().BoolVarP(&config.Flag.Show, "show", "s", false,
		`Show the command with the plain text before executing`)
}
