FROM golang

WORKDIR usr/src/app

COPY . ./

RUN go mod download

RUN go build -v -o logarithm_api ./cmd
RUN go build -v -o logarithm_api_migrate ./db/migration

CMD ["sh", "-c", "./logarithm_api_migrate && ./logarithm_api"]