FROM node:18-alpine

WORKDIR /app

COPY frontend/package.json frontend/package-lock.json ./

RUN npm install --frozen-lockfile

RUN npm i -g serve

COPY frontend .

RUN npm run build

EXPOSE 3000

CMD [ "serve", "-s", "dist" ]