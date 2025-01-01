FROM golang:1.22-bookworm

WORKDIR /app

COPY go.mod .
RUN go mod download 

COPY . ./

RUN go build -o cdp ./cmd/cdp/  

EXPOSE 8080

CMD [ "./cdp" ]
