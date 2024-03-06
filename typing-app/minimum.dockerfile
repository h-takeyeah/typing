FROM node:20.11-slim as builder
WORKDIR /app
COPY package*.json .
RUN npm ci && npm cache clean --force

COPY . .
RUN npm run build

FROM node:20.11-slim
WORKDIR /app
COPY --from=builder --chown=node:node /app/public ./public
COPY --from=builder --chown=node:node /app/.next/standalone ./
COPY --from=builder --chown=node:node /app/.next/static ./.next/static

CMD ["node", "server.js"]
