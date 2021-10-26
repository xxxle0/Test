FROM node:16-alpine as builder
WORKDIR /usr/src/app
COPY package.json yarn.lock ./
RUN yarn
COPY . ./
RUN yarn build

FROM node:12-alpine
WORKDIR /app
RUN npm install -g webserver.local
COPY --from=builder  /usr/src/app/build ./build
EXPOSE 3000
CMD webserver.local -d ./build