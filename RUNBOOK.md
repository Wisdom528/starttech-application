\# Runbook



\## Frontend



Install dependencies:



npm install



Run locally:



npm start



Build:



npm run build



\## Backend



Run:



go run main.go



\## CI/CD



Push changes:



git add .

git commit -m "update"

git push origin main



GitHub Actions automatically executes the workflows.



\## Troubleshooting



Frontend build:



npm install

npm run build



Backend build:



go mod tidy

go test ./...



