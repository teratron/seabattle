FROM ${ARG_FROM}

ADD bin/${ARG_OS}_${ARG_ARCH}/${ARG_BIN} /${ARG_BIN}

USER 65535:65535
ENV HOME /

ENTRYPOINT ["/${ARG_BIN}"]


################################################

#FROM golang:1.16

# Set the Current Working Directory inside the container
#WORKDIR $GOPATH/src/github.com/teratron/goseabattle

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
#COPY . .

# Download all the dependencies
#RUN go get -d -v ./...

# Install the package
#RUN go install -v ./...

# This container exposes port 8080 to the outside world
#EXPOSE 8080

# Run the executable
#CMD ["go-sample-app"]