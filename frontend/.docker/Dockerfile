# Use a Node 16 base image
FROM node:18.15.0-alpine3.17 as builder

# Set the working directory to /app inside the container
WORKDIR /app

# Copy app files
COPY package.json .

RUN yarn install
COPY . .
RUN yarn build

# Bundle static assets with nginx
FROM nginx:1.23.3-alpine as production
# Copy built assets from `builder` image
COPY --from=builder /app/build /usr/share/nginx/html
# Add your nginx.conf
COPY .docker/nginx.conf /etc/nginx/conf.d/default.conf
# Expose port
EXPOSE 80
# Start nginx
CMD ["nginx", "-g", "daemon off;"]