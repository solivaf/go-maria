package config

import (
	"flag"
	_init "github.com/solivaf/go-maria/internal/app/command/init"
	_release "github.com/solivaf/go-maria/internal/app/command/release"
)

type Config struct {
	InitConfig    *_init.Config
	ReleaseConfig *_release.Config
}

func InitConfig() *Config {
	skipPublish := flag.Bool("skip-publish", false, "use skip-release to not modify app version")
	patch := flag.Bool("patch", false, "use to define if should increment patch version")
	minor := flag.Bool("minor", false, "use to increment minor version")
	major := flag.Bool("major", false, "use to increment major version")
	release := flag.Bool("release", false, "")
	initCmd := flag.Bool("init", false, "")
	appName := flag.String("app.name", "app-name", "use to specify the app name")
	flag.Parse()

	if *release {
		return &Config{ReleaseConfig: &_release.Config{Patch: patch, Minor: minor, Major: major, SkipPublish: skipPublish}}
	}
	if *initCmd {
		return &Config{InitConfig: &_init.Config{AppName: appName}}
	}
	return nil
}
