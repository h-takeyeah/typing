FROM oven/bun:slim as base
WORKDIR /app
COPY . .
RUN bun i && bun run build

FROM oven/bun:slim as runner
WORKDIR /app
ENV NODE_ENV production
COPY --from=base /app/public ./public
COPY --from=base /app/.next/standalone ./
COPY --from=base /app/.next/static ./.next/static

EXPOSE 3000
ENV PORT 3000
ENV HOSTNAME "0.0.0.0"

CMD ["node", "server.js"]
