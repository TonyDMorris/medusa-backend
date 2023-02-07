FROM node:latest

WORKDIR /app/medusa

COPY . .

RUN yarn global add @medusajs/medusa-cli@latest

RUN yarn

EXPOSE 9000

ENTRYPOINT ["./develop.sh"]
