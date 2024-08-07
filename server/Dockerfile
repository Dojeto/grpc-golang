FROM proto-gen

WORKDIR /app/server

COPY . .

RUN go mod tidy

RUN go build -o server *.go

CMD [ "./server" ]