FROM alpine
WORKDIR /app

RUN mkdir .config
RUN mkdir target
COPY .config .config/

COPY ./target/karaca-assignment ./target/
ENV APP_ENV=local_env

CMD ["./target/karaca-assignment"]