FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install --production

COPY . .

CMD ["npm", "run", "start:prod"]https://console.cloud.google.com/run/create?chat=true&inv=1&invt=AbtVHQ&project=cykt-399216https://console.cloud.google.com/run/create?chat=true&inv=1&invt=AbtVHQ&project=cykt-399216