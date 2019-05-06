# go-maria
Maria was made to simplify releases for the most modern integrations.

## Features
Maria generate your application version to: 
- Git Repository (Github, Bitbucket, GitLab, etc.); 
- DockerHub
- Sonar Cloud (WIP);
- Generic Webhooks (WIP);

## Usage
Install the binary into your $GOPATH running. Make sure is your $GOPATH/bin into $PATH variable on your OS.
```bash
#Install go-maria CLI into $GOPATH
$ go install github.com/solivaf/go-maria/

#Navigate into your project root
$ cd <project-root>
 
#Init a your project with maria configuration
$ <project-root> go-maria init <module-name>
```

After running **init** command will be created a file named **.goversion.toml**. 
In this file will be configured your project's integrations. Here is an example of a file created

```toml

[module]
  name = "module-name"
  version = "v0.0.1-SNAPSHOT"

```

The version should respect [Semantic Versioning 2.0.0](https://semver.org/) with prefix **v** conforms 
go module versioning. For more information see [Go Modules #NewConcepts](https://github.com/golang/go/wiki/Modules#new-concepts).

## Commands
All commands should be run at root directory of your project.

- init: Initialize a your go module creating a new versioning file named **.goversion.toml**.
- release: Releases a new version into git. More information in [#Release a new version](#release)

## Releasing a new version
Maria will use **docker** and **git** local command line to release new versions, so you will need these clients configured.

With all settled now you can run release command to create a tag and commit a new release into your git repository.

```bash
#Release a new major version
$ go-maria release major

#Release a new minor version
$ go-maria release minor

#Release a new patch version
$ go-maria release patch
```

Running one of these commands above, maria will create a new tag increment a version according SemVer and publish this 
into your remote. The release command will work to all tags you have into your .goversion.toml (**docker**, **module**, **sonar**, etc.) 

If you want avoid the publication into remote you can set the arg **skip-publish**.

```bash
#Release a new version, but don't push to remote
$ go-maria release patch skip-push
```

For more information use go-maria --help or -h.


#### Docker Template
To enable docker on your release you should add [docker] inside **.goversion.toml**. 
Maria will release a new image with these configurations.

```toml

[module]
  name = "go-maria"
  version = "v0.2.1-SNAPSHOT"

[docker]
  organization = "private-organization"
  imageName = "some-image"
  buildDirectory = "./relative/path"
  dockerCompose = "true"
  releaseLatest = "true"
```
- Configurations

| Variable | Required | Default | Description |
|:---------:|:---------:|:--------:|:------------:|
| organization | true | N/A | Docker private repository organization |
| imageName | true | N/A | Name of target docker image |
| buildDirectory | true | N/A | Directory where are placed your Dockerfile or docker-compose.yml |
| dockerCompose | false | false | Indicates if build shoud be made by docker or docker-compose command |
| releaseLatests | false | false | Indicates if should release a tag named **latest** too. |

 