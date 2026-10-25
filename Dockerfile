# custom: 运行时基础镜像固定小版本，避免浮动标签
FROM alpine:3.23

ARG BUILDARCH
WORKDIR /app
RUN apk add --no-cache tzdata
COPY ./${BUILDARCH}/release /app/
VOLUME /app/data

EXPOSE 21114
CMD ["./apimain"]
