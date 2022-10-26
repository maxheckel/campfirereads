FROM node:lts-alpine

RUN mkdir -p /app
WORKDIR /app
COPY ./static /app
RUN npm install -g @vue/cli

EXPOSE 4200

CMD ["npm", "run", "dev", "--",  "--port", "4200", "--host"]