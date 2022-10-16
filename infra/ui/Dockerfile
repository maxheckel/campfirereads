FROM node:lts-alpine

RUN npm install -g http-server

WORKDIR /app

COPY static/package*.json ./

RUN npm install

COPY ./static .

RUN npm run build

EXPOSE 8080
CMD ["http-server", "dist"]