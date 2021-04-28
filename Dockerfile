FROM ubuntu:20.10
RUN mkdir /app
COPY echolife-exporter /app/
WORKDIR /app
CMD ["/app/echolife-exporter"]
EXPOSE 8080