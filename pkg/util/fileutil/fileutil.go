package fileutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/lgdd/deba/pkg/assets"
	"github.com/lgdd/deba/pkg/project"
	"github.com/lgdd/deba/pkg/util/printutil"
)

func CopyFromAssets(sourcePath, destPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	source, err := assets.Templates.Open(sourcePath)
	if err != nil {
		printutil.Error(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}

	defer source.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		printutil.Error(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}

	defer dest.Close()

	_, err = io.Copy(dest, source)
	if err != nil {
		printutil.Error(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}

	printutil.Success("create ")
	fmt.Printf("%s\n", destPath)
}

func UpdateWithData(pomPath string, metadata *project.Metadata) error {
	pomContent, err := ioutil.ReadFile(pomPath)
	if err != nil {
		return err
	}

	tpl, err := template.New(pomPath).Parse(string(pomContent))
	if err != nil {
		return err
	}

	var result bytes.Buffer
	err = tpl.Execute(&result, metadata)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(pomPath, result.Bytes(), 0664)
	if err != nil {
		return err
	}

	return nil
}

func VerifyCurrentDirAsWorkspace(build string) bool {
	files := make(map[string]void)
	dir, err := os.Getwd()

	if err != nil {
		printutil.Error(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files[strings.Split(path, dir)[1]] = void{}
		return nil
	})

	if err != nil {
		printutil.Error(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}

	switch {
	case build == project.Gradle && isGradleWorkspace(files):
		return true
	case build == project.Maven && isMavenWorkspace(files):
		return true
	case build == project.Gradle && isMavenWorkspace(files):
		printutil.Warning(fmt.Sprintln("Oops! It looks like you're trying to do Gradle stuff in a Maven workspace."))
		fmt.Print("Try again with the flag: ")
		printutil.Info("-b maven\n")
		os.Exit(1)
	case build == project.Maven && isGradleWorkspace(files):
		printutil.Warning(fmt.Sprintln("Oops! It looks like you're trying to do Maven stuff in a Gradle workspace."))
		fmt.Print("Try again with the flag: ")
		printutil.Info("-b gradle\n")
		fmt.Print("or without the flag: ")
		printutil.Info("-b maven\n")
		os.Exit(1)
	}
	return false
}

func isGradleWorkspace(files map[string]void) bool {
	sep := string(os.PathSeparator)
	expectedFiles := []string{
		sep + "configs",
		sep + "gradle.properties",
		sep + "settings.gradle",
		sep + "gradle" + sep + "wrapper",
		sep + "build.gradle",
		sep + "gradlew",
		sep + "platform.bndrun",
	}
	for _, expectedFile := range expectedFiles {
		if _, ok := files[expectedFile]; !ok {
			return false
		}
	}
	return true
}

func isMavenWorkspace(files map[string]void) bool {
	sep := string(os.PathSeparator)
	expectedFiles := []string{
		sep + "configs",
		sep + ".mvn" + sep + "wrapper",
		sep + "pom.xml",
		sep + "mvnw",
		sep + "platform.bndrun",
	}
	for _, expectedFile := range expectedFiles {
		if _, ok := files[expectedFile]; !ok {
			return false
		}
	}
	return true
}

func FindFileInParent(fileName string) (string, error) {

	dir, err := os.Getwd()

	if err != nil {
		printutil.Error(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}

	var filePath string
	pathSeparator := string(os.PathSeparator)

	slice := strings.Split(dir, pathSeparator)
	slice = slice[1:]

	for len(slice) > 0 {
		filePath =
			fmt.Sprintf("%s%s%s%s",
				pathSeparator, strings.Join(slice, string(os.PathSeparator)),
				pathSeparator, fileName)

		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			return filePath, nil
		}
		slice = slice[:len(slice)-1]
	}

	return "", fmt.Errorf("%s not found", fileName)
}

type void struct{}
