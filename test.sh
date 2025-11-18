#!/bin/bash



# Test adding numbers and getting sorted list
echo "Testing adding numbers..."

# Add number 5
curl -X POST -d "number=5" http://localhost:8080/numbers
echo ""

# Add number 2
curl -X POST -d "number=2" http://localhost:8080/numbers
echo ""

# Add number 8
curl -X POST -d "number=8" http://localhost:8080/numbers
echo ""

# Add number 1
curl -X POST -d "number=1" http://localhost:8080/numbers
echo ""

# Add number 3
curl -X POST -d "number=3" http://localhost:8080/numbers
echo ""
