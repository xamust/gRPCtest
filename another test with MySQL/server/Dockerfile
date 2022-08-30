FROM golang:1.18
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /server/
COPY . ./
RUN make build
CMD [ "build/server" ]
EXPOSE 8080