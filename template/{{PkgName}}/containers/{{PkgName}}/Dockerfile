{{- $gopath := env "GOPATH" -}}
{{- $pwd := env "PWD" -}}
{{- $relPath := trimPrefix $pwd $gopath -}}
{{- $path := trimPrefix $relPath "/src/" -}}
{{- $pkgpath := printf "%s/%s" $path PkgName -}}
FROM golang:alpine AS build-machine

MAINTAINER {{Author}} <{{Email}}>

RUN apk --update add git openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

RUN mkdir -p /go/src/{{$pkgpath}}
ADD . /go/src/{{$pkgpath}}

WORKDIR /go/src/{{$pkgpath}}

RUN go-wrapper download
RUN go-wrapper install

{{- if UseMySQL -}}
RUN go get -u -d github.com/go-sql-driver/mysql
RUN go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
RUN go build -tags 'mysql' -o /usr/local/bin/migrate github.com/mattes/migrate/cli
{{- end -}}

FROM alpine:latest

{{- if UseMySQL -}}
COPY --from=build-machine /usr/local/bin/migrate /usr/local/bin/migrate
{{- end -}}
COPY --from=build-machine /go/bin/{{PkgName}} /usr/local/bin/{{PkgName}}

CMD ["/usr/local/bin/{{PkgName}}"]
