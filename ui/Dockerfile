FROM node:14.4.0

WORKDIR source

COPY package.json .

RUN yarn

COPY . .

RUN yarn build

CMD ["yarn", "start"]
