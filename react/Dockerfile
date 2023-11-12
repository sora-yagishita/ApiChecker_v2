FROM node:18.17.1
WORKDIR /app/react
COPY package*.json ./
RUN yarn install
COPY . .
ENTRYPOINT [ "npm", "start" ]