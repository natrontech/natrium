# stage build
FROM node:18.11-alpine as build

WORKDIR /app

# copy everything to the container
COPY . .

# clean install all dependencies
RUN npm ci

# remove potential security issues
RUN npm audit fix

# build SvelteKit app
RUN npm run build

# stage run
FROM node:18.11-alpine

WORKDIR /app

# create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# copy dependency list
COPY --from=build /app/package*.json ./

# clean install dependencies, no devDependencies, no prepare script
RUN npm ci --production --ignore-scripts

# remove potential security issues
RUN npm audit fix

# copy built SvelteKit app to /app
COPY --from=build /app/build ./

RUN chown -R appuser:appgroup /app
USER appuser

EXPOSE 3000
CMD ["npm", "run", "start"]
