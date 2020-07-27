FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY configs ./configs/

COPY antdate-admin .

ENV ENV_NAME dev

CMD ["sh","-c","/root/antdate-admin --config=/root/configs/${ENV_NAME}.yaml"]
