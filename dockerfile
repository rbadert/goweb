#The ADD, RUN, and ENTRYPOINT steps are common tasks for any Go project.
#To simplify this, there is an onbuild variant of the golang image that automatically copies the package source,
#fetches the application dependencies, builds the program, and configures it to run on startup.
#With the onbuild variant, the Dockerfile is much simpler:

#FROM golang:onbuild

FROM irisgo/cloud-native-go:latest
EXPOSE 8081

ENV APPSOURCES /go/src/github.com/iris-contrib/cloud-native-go

RUN ${APPSOURCES}/cloud-native-go