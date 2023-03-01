FROM node:latest

WORKDIR /app/medusa

COPY . .

RUN yarn global add @medusajs/medusa-cli@latest

RUN npm i

EXPOSE 8080

ENTRYPOINT ["npm", "run", "start"]
