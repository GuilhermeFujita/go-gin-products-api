## About the Project

This project is a backend API for managing products, following [this YouTube tutorial](https://www.youtube.com/watch?v=3p4mpId_ZU8).  

## Technologies

- **Golang** v1.24.1  
- **Postgres** v12 (using Docker)  

## Running the Project

Follow these steps to run the project:  

### 1. Install Dependencies  
Download the project dependencies using the following command:  

```bash
go mod tidy
```
### 2. Running in Development Mode
To run the project in development mode:
```bash
go run ./cmd/main.go  
```

### 3. Running in Release mode
To run the project in release mode using **make**, ensure it is installed, then execute:
```bash
make run
```
Once running, the API will be available on port 9000.

## Additional features beyond this tutorial
✅ Update products endpoint<br>
✅ Isolate logic to validate id query param<br>
✅ Added validation using [Go playground validator](https://github.com/go-playground/validator)<br>
✅ Usage of environment variables<br>