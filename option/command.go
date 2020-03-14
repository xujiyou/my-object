package option

import (
	"github.com/spf13/cobra"
	"my-object/server"
)

func BuildCommand() {

	var fileCmd = &cobra.Command{
		Use:   "file [path to store]",
		Short: "Object store path.",
		Long:  `Object store path. Object will store in this path.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			server.Run(server.FILE)
		},
	}

	var mongoCmd = &cobra.Command{
		Use:   "mongodb [mongo uri to store]",
		Short: "Mongo uri to store.",
		Long:  `Mongo uri to store. Object will store in mongodb.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			server.Run(server.MONGODB)
		},
	}

	var rootCmd = &cobra.Command{Use: "my-object"}
	rootCmd.AddCommand(fileCmd, mongoCmd)
	_ = rootCmd.Execute()
}
