@echo off
REM Start Docker Compose in the root folder (where the script is located)
start "Docker-Compose" cmd /k "docker-compose up"

REM Navigate to the client directory and start the React app
start "React-App" cmd /k "cd client && npm start"

REM Navigate to the server directory and start the Go server
start "Go-Server" cmd /k "cd server && go run main.go"
