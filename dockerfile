# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:latest as builder

RUN mkdir -p /build
WORKDIR /build

# Copy proval code to the container image.d
# Build the binary.
COPY . ./

# RUN apt-get update && apt-get install -y libvirt-dev
# RUN git config --global url."https://bwleee:ATBBJFDzSJBUX65zh5EfuVwqXHy50E7DBF9A@bitbucket.org/okestrolab".insteadOf "https://bitbucket.org/okestrolab"
# RUN go env -w GOPRIVATE=bitbucket.org/okestrolab
# RUN go get -u bitbucket.org/okestrolab/baton-om-sdk
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a


# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:latest
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
RUN mkdir -p /app
COPY --from=builder /build/myapp /app/

# Run the web service_vw on container startup.
WORKDIR /app
CMD ["/app/myapp"]