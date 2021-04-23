package shell

import (
	"fmt"
	"os"

	"github.com/reiver/go-telnet"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "shell",
		Short:   "Connect and get Liferay Gogo Shell",
		Aliases: []string{"sh"},
		Args:    cobra.NoArgs,
		Run:     run,
	}
	Host string
	Port int
)

func init() {
	Cmd.Flags().StringVar(&Host, "host", "localhost", "--host localhost")
	Cmd.Flags().IntVarP(&Port, "port", "p", 11311, "--port 11311")
}

func run(cmd *cobra.Command, args []string) {
	var caller telnet.Caller = GogoShellCaller
	destination := fmt.Sprintf("%s:%v", Host, Port)
	fmt.Printf("Connecting to %v...\n", destination)
	err := telnet.DialToAndCall(destination, caller)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
