package stop

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/lgdd/liferay-cli/lfr/pkg/util/fileutil"
	"github.com/lgdd/liferay-cli/lfr/pkg/util/printutil"
	"github.com/lgdd/liferay-cli/lfr/pkg/util/procutil"
	"github.com/spf13/cobra"
)

var (
	// Cmd is the command 'stop' which shutdowns the Liferay bundle
	Cmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop a Liferay Tomcat bundle",
		Args:  cobra.NoArgs,
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	if runtime.GOOS == "windows" {
		printutil.Info("not available for Windows since the Tomcat process run in another command window")
		os.Exit(0)
	}
	shutdownScript, err := fileutil.GetTomcatScriptPath("shutdown")

	if err != nil {
		printutil.Danger(err.Error())
		os.Exit(1)
	}

	_, err = procutil.GetCatalinaPid()

	if err != nil {
		printutil.Danger(err.Error())
		os.Exit(1)
	}

	shutdownCmd := exec.Command(shutdownScript)
	shutdownCmd.Stdout = os.Stdout

	err = shutdownCmd.Run()

	if err != nil {
		printutil.Danger(err.Error())
		os.Exit(1)
	}

}
