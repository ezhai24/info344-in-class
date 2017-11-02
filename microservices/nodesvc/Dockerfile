FROM node:alpine

# creates directory and cd into it
WORKDIR /app

COPY package.json .
RUN npm install
COPY . .
EXPOSE 80
ENTRYPOINT ["node", "index.js"]
