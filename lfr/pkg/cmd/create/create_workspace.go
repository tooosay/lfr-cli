package create

import (
	"fmt"
	"os"
	"runtime"

	"github.com/lgdd/liferay-cli/lfr/pkg/cmd/exec"
	"github.com/lgdd/liferay-cli/lfr/pkg/generate/workspace"
	"github.com/lgdd/liferay-cli/lfr/pkg/project"
	"github.com/lgdd/liferay-cli/lfr/pkg/util/fileutil"
	"github.com/lgdd/liferay-cli/lfr/pkg/util/printutil"
	"github.com/spf13/cobra"
)

var (
	createWorkspace = &cobra.Command{
		Use:     "workspace NAME",
		Aliases: []string{"ws"},
		Args:    cobra.ExactArgs(1),
		Run:     generateWorkspace,
	}
	Version string
	Build   string
	Init    bool
)

func init() {
	createWorkspace.Flags().StringVarP(&Version, "version", "v", "7.3", "--version 7.3")
	createWorkspace.Flags().StringVarP(&Build, "build", "b", "gradle", "--build gradle")
	createWorkspace.Flags().BoolVarP(&Init, "init", "i", false, "--init")
}

func generateWorkspace(cmd *cobra.Command, args []string) {
	if fileutil.IsInWorkspaceDir() {
		printutil.Danger("You're already in a Liferay Workspace and I can't create a new one in it.\n")
		os.Exit(1)
	}
	name := args[0]
	err := workspace.Generate(name, Build, Version)
	if err != nil {
		printutil.Danger(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}
	printutil.Success(fmt.Sprintf("\nSuccessfully created a Liferay Workspace '%s'\n", name))

	if Init {
		runInit(name, Build)
	} else {
		printInitCmd(name, Build)
	}
}

func runInit(name, build string) {
	if err := os.Chdir(name); err != nil {
		printutil.Danger(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}

	fmt.Print("\nInitializing Liferay Bundle using:\n\n")

	switch build {
	case project.Gradle:
		printutil.Info("lfr exec initBundle\n\n")
		exec.RunWrapperCmd([]string{"initBundle"})
	case project.Maven:
		printutil.Info("lfr exec bundle-support:init\n\n")
		exec.RunWrapperCmd([]string{"bundle-support:init"})
	}
}

func printInitCmd(name, build string) {
	fmt.Println("\nInitialize your Liferay bundle:")
	if runtime.GOOS == "windows" {
		switch build {
		case project.Gradle:
			printutil.Info(fmt.Sprintf("cd %s && lfr exec initBundle\n", name))
		case project.Maven:
			printutil.Info(fmt.Sprintf("cd %s && lfr exec bundle-support:init\n", name))
		}
	} else {
		switch build {
		case project.Gradle:
			printutil.Info(fmt.Sprintf("cd %s && lfr exec initBundle\n", name))
		case project.Maven:
			printutil.Info(fmt.Sprintf("cd %s && lfr exec bundle-support:init\n", name))
		}
	}
	fmt.Print("\n")
}