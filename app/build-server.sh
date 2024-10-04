#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Compile Tailwind CSS
npx tailwindcss -i ./ui/main.css -o ./ui/tailwind.css

# Build the Go application
go build -buildvcs=false

# Run the backend
./backend
