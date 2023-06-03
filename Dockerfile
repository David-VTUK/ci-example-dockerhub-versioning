FROM ubuntu:22.10
RUN mkdir /app
COPY ./bin/example-app /app/
WORKDIR /app
RUN chmod +x /app/example-app
CMD ["/app/example-app"]
EXPOSE 8081
