FROM proto-gen

WORKDIR /app/client

COPY . .

RUN go mod tidy

RUN go build -o client client.go

CMD [ "./client" ]