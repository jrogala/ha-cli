package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/jrogala/ha-cli/config"
	"github.com/jrogala/ha-cli/internal/cmdutil"
	"github.com/jrogala/ha-cli/pkg/ops"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(setupCmd, configCmd, entitiesCmd, stateCmd,
		onCmd, offCmd, toggleCmd, callCmd, servicesCmd)

	setupCmd.Flags().String("token", "", "Long-lived access token")
	setupCmd.Flags().String("url", "", "Home Assistant URL")
	setupCmd.MarkFlagRequired("token")

	entitiesCmd.Flags().String("domain", "", "Filter by domain (light, switch, sensor, etc.)")
	entitiesCmd.Flags().String("search", "", "Search by name or entity_id")

	servicesCmd.Flags().String("domain", "", "Filter by domain")

	callCmd.Flags().String("data", "{}", "JSON data to send with service call")
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure Home Assistant connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, _ := cmd.Flags().GetString("token")
		url, _ := cmd.Flags().GetString("url")

		viper.Set("token", token)
		if url != "" {
			viper.Set("url", url)
		}
		if err := config.Save(); err != nil {
			return err
		}
		fmt.Println("Configuration saved")
		return nil
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show Home Assistant server info",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmdutil.NewClient()
		if err != nil {
			return err
		}
		info, err := ops.GetConfig(c)
		if err != nil {
			return err
		}
		cmdutil.Render(cmd, info, func() {
			fmt.Printf("Name:     %s\n", info.LocationName)
			fmt.Printf("Version:  %s\n", info.Version)
			fmt.Printf("Timezone: %s\n", info.TimeZone)
		})
		return nil
	},
}

var entitiesCmd = &cobra.Command{
	Use:     "entities",
	Aliases: []string{"ls", "list"},
	Short:   "List entities",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmdutil.NewClient()
		if err != nil {
			return err
		}
		domain, _ := cmd.Flags().GetString("domain")
		search, _ := cmd.Flags().GetString("search")

		entries, err := ops.ListEntities(c, ops.ListOptions{Domain: domain, Search: search})
		if err != nil {
			return err
		}
		cmdutil.Render(cmd, entries, func() {
			w := cmdutil.NewTabWriter()
			fmt.Fprintln(w, "ENTITY_ID\tSTATE\tNAME")
			for _, e := range entries {
				fmt.Fprintf(w, "%s\t%s\t%s\n", e.EntityID, e.State, e.Name)
			}
			w.Flush()
		})
		return nil
	},
}

var stateCmd = &cobra.Command{
	Use:   "state <entity_id>",
	Short: "Get entity state",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmdutil.NewClient()
		if err != nil {
			return err
		}
		detail, err := ops.GetState(c, args[0])
		if err != nil {
			return err
		}
		cmdutil.Render(cmd, detail, func() {
			fmt.Printf("Entity:  %s\n", detail.EntityID)
			fmt.Printf("Name:    %s\n", detail.Name)
			fmt.Printf("State:   %s\n", detail.State)
			fmt.Printf("Updated: %s\n", detail.LastUpdated)
			for k, v := range detail.Attributes {
				if k == "friendly_name" {
					continue
				}
				fmt.Printf("  %s: %v\n", k, v)
			}
		})
		return nil
	},
}

var onCmd = &cobra.Command{
	Use:   "on <entity_id>",
	Short: "Turn on an entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmdutil.NewClient()
		if err != nil {
			return err
		}
		result, err := ops.TurnOn(c, args[0])
		if err != nil {
			return err
		}
		cmdutil.Render(cmd, result, func() {
			fmt.Printf("Turned on %s\n", args[0])
		})
		return nil
	},
}

var offCmd = &cobra.Command{
	Use:   "off <entity_id>",
	Short: "Turn off an entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmdutil.NewClient()
		if err != nil {
			return err
		}
		result, err := ops.TurnOff(c, args[0])
		if err != nil {
			return err
		}
		cmdutil.Render(cmd, result, func() {
			fmt.Printf("Turned off %s\n", args[0])
		})
		return nil
	},
}

var toggleCmd = &cobra.Command{
	Use:   "toggle <entity_id>",
	Short: "Toggle an entity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmdutil.NewClient()
		if err != nil {
			return err
		}
		result, err := ops.Toggle(c, args[0])
		if err != nil {
			return err
		}
		cmdutil.Render(cmd, result, func() {
			fmt.Printf("Toggled %s\n", args[0])
		})
		return nil
	},
}

var callCmd = &cobra.Command{
	Use:   "call <domain> <service>",
	Short: "Call a Home Assistant service",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmdutil.NewClient()
		if err != nil {
			return err
		}
		dataStr, _ := cmd.Flags().GetString("data")
		var data map[string]any
		if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
			return fmt.Errorf("invalid JSON data: %w", err)
		}
		result, err := ops.CallService(c, args[0], args[1], data)
		if err != nil {
			return err
		}
		cmdutil.Render(cmd, result, func() {
			fmt.Printf("Called %s.%s\n", args[0], args[1])
		})
		return nil
	},
}

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List available services",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cmdutil.NewClient()
		if err != nil {
			return err
		}
		domain, _ := cmd.Flags().GetString("domain")

		entries, err := ops.ListServices(c, domain)
		if err != nil {
			return err
		}
		cmdutil.Render(cmd, entries, func() {
			w := cmdutil.NewTabWriter()
			fmt.Fprintln(w, "DOMAIN\tSERVICE")
			for _, e := range entries {
				fmt.Fprintf(w, "%s\t%s\n", e.Domain, e.Service)
			}
			w.Flush()
		})
		return nil
	},
}
