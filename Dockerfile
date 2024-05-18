FROM golang:1.22-alpine AS stage1
WORKDIR /project/ciao/

COPY go.* .
RUN  go mod download

COPY . .
RUN go build -o ./cmd/ciaPostNRelExec ./cmd/main.go

FROM alpine:latest
WORKDIR /project/ciao/


COPY --from=stage1 /project/ciao/cmd/ciaPostNRelExec ./cmd/
COPY --from=stage1 /project/ciao/dev.env ./

EXPOSE 50052
ENTRYPOINT [ "/project/ciao/cmd/ciaPostNRelExec" ]