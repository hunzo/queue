FROM golang:alpine as build

WORKDIR /app

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o loadtest


FROM scratch as production
# FROM alpine as production

COPY --from=build /app/loadtest .
COPY --from=build /app/config/ ./config

ENV TZ=Asia/Bangkok

CMD [ "./loadtest" ]