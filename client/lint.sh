echo "Move to client folder"
cd client 

echo "Checking tsc file, type checking..."
yarn type-check

echo "Lint file..."
yarn lint