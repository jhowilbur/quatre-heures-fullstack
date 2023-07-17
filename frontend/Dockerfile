# develop stage
FROM node:16-alpine as develop-stage
WORKDIR /app

COPY package*.json ./
RUN npm install --legacy-peer-deps
COPY . .

EXPOSE 8080
CMD ["npm", "run", "serve"]

# TODO
## build stage
#FROM develop-stage as build-stage
#RUN npm run build
#
## production stage
#FROM nginx:1.15.7-alpine as production-stage
#COPY --from=build-stage /app/dist /usr/share/nginx/html
#EXPOSE 8080
#CMD ["nginx", "-g", "daemon off;"]