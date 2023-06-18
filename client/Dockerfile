# Base image
FROM node:16-alpine

# Set the working directory
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY . /app/
# COPY ["package.json", "package-lock.json*", "./"]


RUN npm install
# Install dependencies
RUN npm install -g vite

# Expose the port on which the app will run
EXPOSE 5173
# Set the command to run the app
CMD [ "npm", "run", "dev", "--", "--host", "0.0.0.0"]