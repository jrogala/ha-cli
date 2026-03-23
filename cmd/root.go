// Package cmd implements the ha-cli commands.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jrogala/ha-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "ha",
	Short: "CLI for controlling Home Assistant",
	Long: `CLI for controlling Home Assistant via REST API.

Quick examples:
  ha setup --token TOKEN                       Save access token
  ha config                                    Show HA server info
  ha entities                                  List all entities
  ha entities --domain light                   List only lights
  ha state light.kitchen                       Get entity state
  ha on light.kitchen                          Turn on entity
  ha off light.kitchen                         Turn off entity
  ha toggle light.kitchen                      Toggle entity
  ha call light turn_on --data '{"entity_id":"light.kitchen","brightness":255}'
  ha services                                  List available services
  ha services --domain light                   List light services`,
}

func Execute() {
	config.Init()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("json", false, "Output raw JSON responses")
	rootCmd.SetHelpFunc(customHelp)
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

func customHelp(cmd *cobra.Command, _ []string) {
	if cmd == rootCmd {
		printTree()
		return
	}
	if !cmd.HasSubCommands() {
		printLeafHelp(cmd)
		return
	}
	printSubtree(cmd)
}

func printTree() {
	fmt.Println("ha-cli - Home Assistant CLI client")
	fmt.Println("")
	fmt.Println("Global: --json (raw JSON output)")
	fmt.Println("")
	fmt.Println("Commands:")

	for _, cmd := range rootCmd.Commands() {
		if cmd.Hidden || cmd.Name() == "help" || cmd.Name() == "completion" {
			continue
		}
		if cmd.HasSubCommands() {
			fmt.Printf("  %s\n", cmd.Name())
			for _, sub := range cmd.Commands() {
				if sub.Hidden {
					continue
				}
				aliases := ""
				if len(sub.Aliases) > 0 {
					aliases = " (" + strings.Join(sub.Aliases, ", ") + ")"
				}
				fmt.Printf("    %-10s %s%s\n", sub.Name(), sub.Short, aliases)
			}
		} else {
			fmt.Printf("  %-12s %s\n", cmd.Name(), cmd.Short)
		}
	}

	fmt.Println("")
	fmt.Println("Run 'ha <command> --help' for full details.")
}

func printSubtree(cmd *cobra.Command) {
	fmt.Printf("%s\n\n", cmd.Short)

	for _, sub := range cmd.Commands() {
		if sub.Hidden {
			continue
		}
		aliases := ""
		if len(sub.Aliases) > 0 {
			aliases = " (" + strings.Join(sub.Aliases, ", ") + ")"
		}
		fmt.Printf("  %-10s %s%s\n", sub.Name(), sub.Short, aliases)
	}

	fmt.Println("")
	fmt.Printf("Run 'ha %s <subcommand> --help' for full details.\n", cmd.Name())
}

func printLeafHelp(cmd *cobra.Command) {
	fmt.Printf("%s %s\n", cmd.UseLine(), "")
	fmt.Println(cmd.Short)

	if cmd.HasLocalFlags() {
		fmt.Println("")
		fmt.Println("Flags:")
		cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
			shorthand := ""
			if f.Shorthand != "" {
				shorthand = "-" + f.Shorthand + ", "
			}
			def := ""
			if f.DefValue != "" && f.DefValue != "false" && f.DefValue != "0" {
				def = " (default: " + f.DefValue + ")"
			}
			fmt.Printf("  %s--%s %s%s\n", shorthand, f.Name, f.Usage, def)
		})
	}
}
