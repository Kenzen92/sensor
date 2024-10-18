# Use an official Python runtime as a base image
FROM python:3.10-slim

# Set the working directory inside the container
WORKDIR /app

# Install dependencies for building native extensions (gcc, musl-dev, and sqlite)
RUN apt-get update && apt-get install -y gcc libsqlite3-dev build-essential


# Copy the requirements.txt file and install dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt


# Copy the Python source code to the working directory
COPY . .

# Run the Python script
CMD ["python", "main.py"]
