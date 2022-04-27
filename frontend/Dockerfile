FROM nginx:stable-alpine

ARG BUILD_DIR_ARG=dist
ENV BUILD_DIR=${BUILD_DIR_ARG}

COPY ./${BUILD_DIR} /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]