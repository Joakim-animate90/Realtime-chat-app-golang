FROM hub.c.163.com/public/centos:7.0
RUN [ "mkdir", "/usr/local/realtime-chat-go" ]
RUN [ "mkdir", "-p", "/usr/local/realtime-chat-go/web/static/file" ]
WORKDIR /usr/local/realtime-chat-go
COPY /bin/chat /usr/local/realtime-chat-go
COPY /config.toml /usr/local/realtime-chat-go
RUN chmod +x /usr/local/realtime-chat-go/chat
ENTRYPOINT ["/usr/local/realtime-chat-go/chat"]