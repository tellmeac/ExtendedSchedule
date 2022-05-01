FROM node:16.14.0 as builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci --silent

COPY . .

RUN npm run build

# Runner 
FROM nginx:1.15.2-alpine

# Replace default nginx config with custom
RUN rm -rf /etc/nginx/conf.d
COPY conf.d /etc/nginx

# Static build
COPY --from=builder /app/build /usr/share/nginx/html/

# Copy .env file and shell script to container
WORKDIR /usr/share/nginx/html

# Start Nginx server
CMD ["/bin/ash", "-c", "nginx", "-g", "daemon off;"]