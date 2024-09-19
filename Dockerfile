FROM golang
WORKDIR /udise-api  
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags='-s' -o=./bin/api ./cmd/api
CMD ./bin/api -port=8080 -env=production
EXPOSE 8080
