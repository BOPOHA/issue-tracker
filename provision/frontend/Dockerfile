# docker-ce needs to be 17.05 or later.
FROM node:10-alpine as front_builder
RUN  apk add --no-cache bash git openssh
WORKDIR /root/
COPY static/package.json .
COPY static/package-lock.json .
RUN npm config set prefix /root/
RUN npm install
RUN find / -name bower
ENV PATH=$PATH:/root/node_modules/.bin
RUN echo $PATH
WORKDIR /static/
COPY static/ .
RUN bower --allow-root install

FROM nginx:alpine
COPY --from=front_builder /static/app /app
COPY provision/frontend/nginx.default.conf /etc/nginx/conf.d/default.conf
ENTRYPOINT ["nginx"]
