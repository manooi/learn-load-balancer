# Use Node.js 14
FROM node:14

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to container
COPY package*.json ./

# Install app dependencies
RUN npm install

# Copy app source code to the container
COPY . .

# Expose the port that the app runs on
EXPOSE 3000

# Command to start the app
CMD ["node", "index.js"]

