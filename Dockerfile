FROM golang:latest

RUN mkdir -p /usr/src/app/
RUN mkdir -p $GOPATH/src/github.com/Ubidy
WORKDIR /usr/src/app

COPY . /usr/src/app/

RUN ln -s /usr/src/app $GOPATH/src/github.com/Ubidy/Ubidy_EmployerNotificationAPI

RUN go get github.com/auth0-community/auth0
RUN go get github.com/denisenkom/go-mssqldb
RUN go get github.com/gin-gonic/gin
RUN go get github.com/google/uuid
RUN go get github.com/gin-contrib/cors
RUN go get gopkg.in/GetStream/stream-go2.v1
RUN go get github.com/Microsoft/ApplicationInsights-Go/appinsights
RUN go build -o employeractivitystreamapi .

EXPOSE 5020 80
ENTRYPOINT [ "./employeractivitystreamapi" ]
