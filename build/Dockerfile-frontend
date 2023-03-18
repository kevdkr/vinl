FROM node:lts AS dependencies
WORKDIR /app
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json
RUN npm ci

FROM dependencies AS build
COPY . .
RUN npm run build

FROM nginx:alpine
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/build /usr/share/nginx/html/
