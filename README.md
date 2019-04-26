# go-maria
Go Maria was made to simplify releases for the most modern integrations.

## Features
Maria generate your application version to: 
- Git Repository (Github, Bitbucket, GitLab, etc.); 
- DockerHub (WIP);
- Sonar Cloud (WIP);
- Generic Webhooks (WIP);

If you need any other platform integration, please open an issue so we can attendee a new request.

## Usage
Install the binary into your $GOPATH running. Make sure is your $GOPATH/bin into $PATH variable on your OS.
```bash
#Install gomaria CLI into $GOPATH
$ go install github.com/solivaf/go-maria/

#Navigate into your project root
$ cd <project-root>
 
#Init a your project with maria configuration
$ <project-root> gomaria init
```

After running **init** command will be created a file named **.goversion.toml**. 
In this file will be configured your project's integrations. Here is an example of a file created

```toml

[module]
  name = "app-name"
  version = "v0.0.1-SNAPSHOT"

```

The version should respect [Semantic Versioning 2.0.0](https://semver.org/) with prefix **v** conforms 
go module versioning. For more information see [Go Modules #NewConcepts](https://github.com/golang/go/wiki/Modules#new-concepts).

## Commands
All commands should be run at root directory of your project.

- init: Initialize a your go module creating a new versioning file named **.goversion.toml**.
- release: Releases a new version into git. More information in [#Release a new version](#release)

## Releasing a new version
With all settled now you can run release command to create a tag and commit a new release into your git repository.

```bash
#Release a new major version
$ gomaria --release --major

#Release a new minor version
$ gomaria --release --minor

#Release a new patch version
$ gomaria --release --patch
```

Running one of these commands above, maria will create a new tag increment a version according SemVer and publish this 
into your remote. If you want avoid the publication into remote you can set the arg **skip-publish**.

```bash
#Release a new version, but don't push to remote
$ gomaria --release --patch --skip-publish
```
