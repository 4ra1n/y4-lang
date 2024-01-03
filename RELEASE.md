## Release

### Release Steps

- cmd/main.go - cli.Version cli.BuildTime
- .github/y4-lang.yml - echo "VERSION=v0.0.1" >> $GITHUB_ENV
- Dockerfile - ARG VERSION=v0.0.1
- CHANGELOG
- github/delete-caches/main.go
- github/delete-runs/main.go
- github/build/main.go (windows)
- sudo ./docker-build.sh (linux)

### Update Go Version

- go.mod - go *
- .github/y4-lang.yml - go-version: '*'
- Dockerfile - FROM golang:*
- CHANGELOG
