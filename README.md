# Post-it-Pad

Post-it Pad contains generic html email layouts. It's integrated and intended to be used with [Post-it by apio](https://telescope.apiobuild.com/app/post-it). However, the generated html can also be used anywhere.

It comes with cmd interface and can be imported to other go projects. The API supports the following basic functionality:

1. Create new blank layout for customization
2. Generate HTML with example json for specific or all layouts
3. Generate HTML with custom json

We welcome contributions for more layouts as the community sees fit.

## Develop

Use the command to build and test out functionality:

```bash
go run ./cmd/pad/main.go generate -h
```

## Usage

### Command Cli

Alternatively, use the pre-built docker container to use the cmd:

```bash
docker run -v $(pwd)/output:/app/output ghcr.io/apiobuild/post-it-pad generate -l receipt
```

The output will be at `output/generated.html`

### Go Package

```bash
go get github.com/apiobuild/post-it-pad
```

### Build

Use the `go build` command to build the cli:

```bash
go build -o pad cmd/pad/main.go
```

Alternatively, use `docker build -t post-it-pad .` to build the docker image.

### Tests

Use the command to run test suite.

```bash
go test ./... -coverprofile cp.out
```

## Email Template CSS

The email template css is based off [https://github.com/leemunroe/responsive-html-email-template](https://github.com/leemunroe/responsive-html-email-template).
