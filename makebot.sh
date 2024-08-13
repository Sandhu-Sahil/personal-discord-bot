#!/bin/bash

# Remove the old bot
rm -f make_bot

# Check if Go is installed
if ! command -v go &> /dev/null; then
    # If Go is not installed, install it
    sudo apt-get update
    sudo apt-get install golang-go
fi

# Install the required Go modules
# go mod download

# Build the Golang bot
go build -o make_bot main.go

# clear the screen
clear

# Make the bot executable
chmod +x make_bot

# Run the bot
./make_bot