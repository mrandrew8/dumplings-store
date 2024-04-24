FROM node:16 as builder
WORKDIR /usr/src/app
copy . .
RUN npm install
RUN npm run build

#release
FROM nginx:latest
COPY --from=builder /usr/src/app/dist /usr/share/nginx/html
COPY --from=builder /usr/src/app/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 80