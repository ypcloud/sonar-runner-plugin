package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"text/template"
	"strings"
	"fmt"
)

type Plugin struct {
	Host     string
	Login    string
	Password string
	Key      string
	Name     string
	Version  string
	Sources  string
	Language string
	Profile  string
	Encoding string
	LcovPath string
	Debug    bool

	Path     string
	Repo     string
	Branch   string
	Default  string // default master branch
}

func (p *Plugin) Exec() error {

	err := p.buildRunnerProperties()
	if err != nil {
		return err
	}

	err = p.execSonarRunner()
	if err != nil {
		logrus.Println(err)
		return err
	}

	return nil
}

func (p Plugin) buildRunnerProperties() error {

	p.Key = strings.Replace(p.Key, "/", ":", -1)

	tmpl, err := template.ParseFiles("sonar-runner.properties.tmpl")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("/opt/sonar/conf/sonar-runner.properties")
	defer f.Close()
	if err != nil {
		fmt.Println("Error creating file: ", err)
		panic(err)
	}

	if p.Debug {
		err = tmpl.ExecuteTemplate(os.Stdout, "sonar-runner.properties.tmpl", p)
		if err != nil {
			panic(err)
		}
	}

	err = tmpl.ExecuteTemplate(f, "sonar-runner.properties.tmpl", p)
	if err != nil {
		panic(err)
	}

	return nil
}

func (p Plugin) execSonarRunner() error {
	// run archive command
	cmd := exec.Command("java", "-jar", "/opt/sonar/runner.jar", "-Drunner.home=/opt/sonar/")
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printOutput(output)

	if err != nil {
		return err
	}

	return nil
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}