# 1 шаг - сборки
FROM golang:1.19-alpine AS build_stage
COPY . /go/src/2023_1_BKS
WORKDIR /go/src/2023_1_BKS/cmd/post
RUN go env -w GO111MODULE=auto
RUN go install .

# 2 шаг
FROM alpine AS run_stage
WORKDIR /app_binary
COPY --from=build_stage /go/bin/post /app_binary/
RUN chmod +x ./post
ENTRYPOINT ./post