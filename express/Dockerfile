# Use the 'oven/bun' base image
FROM oven/bun

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to container
COPY package*.json ./

# Install app dependencies
RUN bun install

# Copy app source code to the container
COPY . .

# Expose the port that the app runs on
EXPOSE 3000

# Command to start the app
CMD ["bun", "index.js"]
