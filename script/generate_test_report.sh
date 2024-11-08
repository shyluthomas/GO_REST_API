#!/bin/bash

echo "Running tests with coverage..."
go test ./... -coverprofile=coverage.out

echo "Generating HTML report..."
go tool cover -html=coverage.out -o coverage.html

echo "Coverage report generated: coverage.html"
