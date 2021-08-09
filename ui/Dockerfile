FROM node:14.4.0

WORKDIR source

COPY package.json .

RUN yarn

COPY . .

env REACT_APP_SERVICE_URL http://some.internal.service.msb.com

RUN yarn build

CMD ["yarn", "start"]
