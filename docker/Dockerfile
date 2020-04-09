#Use the minimal golang alpine base image
FROM golang:1.8-alpine

LABEL maintainer="ian@constructd.com"

WORKDIR /go/src/app

COPY ./app/helloworld.go .

#Expose is not strictly required, but should be included as best practice
EXPOSE 80

#Note that this container runs as root, and should not be used in production!
CMD ["go", "run", "helloworld.go"]