FROM centos

MAINTAINER xxxx@qq.com

RUN yum install -y epel-release
RUN yum install -y gcc
RUN yum install -y go

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
#ADD conf /go/src/conf
ADD github.com /go/src/github.com
ADD golang.org /go/src/golang.org
ADD gopkg.in /go/src/gopkg.in
#ADD ./service/protoc /go/src/service/protoc
#ADD ./service/common /go/src/service/common
ADD build.sh /build.sh
RUN chmod +x /build.sh
RUN /build.sh