Version=0.0.1
CommitHash=$(shell git rev-parse --short HEAD)
BuildTime=$(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')

clean:
	rm -f pockats-build

default:
	go build -ldflags "\
		-X github.com/pockats/pockats-backend/version.Code=$(Version) \
		-X github.com/pockats/pockats-backend/version.CommitHash=$(CommitHash) \
		-X github.com/pockats/pockats-backend/version.BuildTime=$(BuildTime) \
		"

fmt:
	go fmt *.go
