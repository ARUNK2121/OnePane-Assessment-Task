FROM golang:1.21.5-alpine3.19 AS build-stage
WORKDIR /home/project/
COPY ./ /home/project/
RUN mkdir -p /home/build
RUN go build -v -o /home/build/api ./


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /home/build/api /api
EXPOSE 8080
CMD ["/api"]
