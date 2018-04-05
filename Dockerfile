FROM golang

RUN rm -rf /go/src/github.com/x64puzzle/go-playground
COPY . /go/src/github.com/x64puzzle/go-playground
WORKDIR /go/src/github.com/x64puzzle/go-playground

# RUN go-wrapper download
# RUN go-wrapper install

CMD go-playground

EXPOSE 8080