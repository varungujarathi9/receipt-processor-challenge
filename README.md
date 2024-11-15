# Receipt Processor

## Running the Application

- ### Build and Run with Docker:

  ```bash
  docker build -t receipt-processor .
  docker run -p 8080:8080 receipt-processor
  ```

- ### Access Endpoints:

  - POST: http://localhost:8080/receipts/process

  - GET: http://localhost:8080/receipts/{id}/points
