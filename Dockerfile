FROM golang:1.21

RUN groupadd -r appgroup && useradd -r -g appgroup appuser

WORKDIR /app

COPY main.go .

RUN go build -o myapp main.go

RUN chown -R appuser:appgroup /app
USER appuser

EXPOSE 4000

CMD ["./myapp"]
