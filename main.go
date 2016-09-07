package main

import (
	"fmt"
	"os"
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {

	cli.VersionFlag = cli.BoolFlag{
		Name: "plugin-version, V",
		Usage: "print plugin version",
	}

	app := cli.NewApp()
	app.Name = "sonar runner plugin"
	app.Usage = "sonar runner plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository full name",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "repo.branch",
			Usage:  "repository default branch",
			EnvVar: "DRONE_REPO_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "repository branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},


		cli.StringFlag{
			Name:   "host",
			Usage:  "Sonar host URL",
			EnvVar: "SONAR_HOST,PLUGIN_HOST",
			Value: "localhost",
		},
		cli.StringFlag{
			Name:   "login",
			Usage:  "sonar login",
			EnvVar: "SONAR_LOGIN,PLUGIN_LOGIN",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "sonar password",
			EnvVar: "SONAR_PASSWORD,PLUGIN_PASSWORD",
		},
		cli.StringFlag{
			Name:   "key",
			Usage:  "Project Key",
			EnvVar: "PLUGIN_KEY,DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "name",
			Usage:  "Project name",
			EnvVar: "PLUGIN_NAME,DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "version",
			Usage:  "Project version",
			EnvVar: "PLUGIN_VERSION",
		},
		cli.StringFlag{
			Name:   "sources",
			Usage:  "Project sources to scan",
			EnvVar: "PLUGIN_SOURCES",
		},
		cli.StringFlag{
			Name:   "language",
			Usage:  "Project language",
			EnvVar: "PLUGIN_LANGUAGE",
			Value:  "js",
		},
		cli.StringFlag{
			Name:   "profile",
			Usage:  "Project profile",
			EnvVar: "PLUGIN_PROFILE",
			Value:  "node",
		},
		cli.StringFlag{
			Name:   "encoding",
			Usage:  "Project source encondig",
			EnvVar: "PLUGIN_ENCODING",
			Value:  "UTF-8",
		},
		cli.StringFlag{
			Name:   "lcovpath",
			Usage:  "Project lcov report path",
			EnvVar: "PLUGIN_LCOVPATH",
			Value:  "test/coverage/reports/lcov.info",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Print generated config - debug purposes",
			EnvVar: "PLUGIN_DEBUG",
		},

	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {

	plugin := Plugin {

		Host:     c.String("host"),
		Login:    c.String("login"),
		Password: c.String("password"),
		Key:      c.String("key"),
		Name:     c.String("name"),
		Version:  c.String("version"),
		Sources:  c.String("sources"),
		Language: c.String("language"),
		Profile:  c.String("profile"),
		Encoding: c.String("encoding"),
		LcovPath: c.String("lcovpath"),
		Debug:    c.Bool("debug"),

		Path:     c.String("path"),
		Repo:     c.String("repo.name"),
		Default:  c.String("repo.branch"),
		Branch:   c.String("commit.branch"),
	}

	return plugin.Exec()
}
