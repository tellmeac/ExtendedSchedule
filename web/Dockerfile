FROM node:16.14.0 as builder

WORKDIR /app

COPY package*.json /app/

RUN npm install

COPY . .

# docker build \
# -t essearch/ess-elasticsearch:1.7.6 \
# --build-arg REACT_APP_API_BASE_URL='http://195.133.196.111/api' \
# --build-arg REACT_APP_GOOGLE_CLIENT_ID='client_id' \
# --no-cache .
ARG REACT_APP_API_BASE_URL
ARG REACT_APP_GOOGLE_CLIENT_ID

ENV REACT_APP_API_BASE_URL $REACT_APP_API_BASE_URL
ENV REACT_APP_GOOGLE_CLIENT_ID $REACT_APP_GOOGLE_CLIENT_ID

RUN npm run build

FROM nginx:1.15.2-alpine

COPY --from=builder /app/build /var/www/html

COPY --from=builder /app/nginx.conf /etc/nginx/conf.d/default.conf
