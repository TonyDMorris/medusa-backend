FROM node:latest

WORKDIR /app/medusa

COPY . .


RUN npm install

EXPOSE 8080

ENTRYPOINT ["npm", "run", "start"]
