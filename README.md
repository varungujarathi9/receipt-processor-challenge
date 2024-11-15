# Receipt Processor

## Running the Application

### Build and Run with Docker

```bash
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```

### Access Endpoints

- POST: <http://localhost:8080/receipts/process>

- GET: <http://localhost:8080/receipts/{id}/points>

### Postman Collection

[Receipt Processor.postman_collection.json](<Receipt Processor.postman_collection.json>)

## Code Overview

### Main Entry Point

The main entry point of the application is `main.go`. It sets up the HTTP server and routes.

### Handlers

Handlers for processing receipts and retrieving points are defined in `handlers/receipt_handler.go`.

### Models

Data models for the application are defined in the models package:

`item.go`
`points_response.go`
`receipt.go`

### Utilities

Utility functions for calculating points are defined in utils/points_calculator.go.
