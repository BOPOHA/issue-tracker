# docker-ce needs to be 17.05 or later.
FROM node:10-alpine as front_builder
RUN  apk add --no-cache bash git openssh
WORKDIR /static/
COPY static .
RUN npm install
RUN ./node_modules/bower/bin/bower --allow-root install

FROM nginx:alpine
COPY --from=front_builder /static/app /app
COPY provision/frontend/nginx.default.conf /etc/nginx/conf.d/default.conf
ENTRYPOINT ["nginx"]
