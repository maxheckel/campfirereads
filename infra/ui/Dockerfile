FROM node:lts-alpine

RUN mkdir -p /app
WORKDIR /app
COPY ./static /app

EXPOSE 4200

CMD ["npm", "run", "dev", "--",  "--port", "4200", "--host"]