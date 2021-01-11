FROM node:12-alpine as builder
ENV NODE_ENV=production
WORKDIR /service
COPY package.json yarn.lock ./
COPY . .
RUN yarn
RUN yarn build
FROM node:12-alpine
WORKDIR /service
COPY --from=builder /service/node_modules ./node_modules/
COPY --from=builder /service/dist ./dist/
CMD ["node", "dist/server.js"]