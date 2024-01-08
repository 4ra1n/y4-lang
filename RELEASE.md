## Release

### Release Steps

- cmd/main.go - cli.Version cli.BuildTime
- .github/y4-lang.yml - echo "VERSION=v0.0.1" >> $GITHUB_ENV
- CHANGELOG
- action.bat clean
- action.bat build

### Update Go Version

- go.mod - go *
- .github/y4-lang.yml - go-version: '*'
- Dockerfile - FROM golang:*
- CHANGELOG
