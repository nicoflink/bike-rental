# Stage1: Build #
FROM node:17.8.0-alpine3.15 AS build
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage2: Run Server #
FROM nginx:1.23.1-alpine
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/server.conf /etc/nginx/conf.d/http_server.conf
COPY --from=build /app/dist/bike-map-app /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf