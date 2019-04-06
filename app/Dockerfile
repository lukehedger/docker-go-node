FROM node:11 as builder
COPY . /app
WORKDIR /app
RUN npm install
RUN npm run build

CMD ["npm", "start"]
