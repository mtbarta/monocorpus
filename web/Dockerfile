FROM node:9.11.1-alpine as builder

# ARG NODE_ENV=production

COPY . /app

WORKDIR /app
RUN npm install && npm run build

# build production image
FROM nginx:1.13.12-alpine

RUN apk add --no-cache curl

COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist/ /usr/share/nginx/html

# CMD ['nginx', '-g', 'daemon off;']