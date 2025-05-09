# Go Auction Concurrency

A concurrent auction system built with Go and MongoDB.

## 🚀 Running the Project with Docker and Docker Compose

This project uses Docker and Docker Compose to simplify setup and running.

---

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) installed
- [Docker Compose](https://docs.docker.com/compose/install/) installed

---

## 🏗️ How to Run

1. **Clone the repository:**

   ```sh
   git clone https://github.com/marcofilho/go-auction-concurrency.git
   cd go-auction-concurrency
   ```

2. **Build and start the containers:**

   ```sh
   docker-compose up --build
   ```

   This will:

   - Build the Go application image
   - Start the Go app and MongoDB containers

3. **Access the application:**

   - The Go API will be available at: `http://localhost:8080`
   - MongoDB will be running at: `mongodb://localhost:27017`

---

## 🛑 Stopping the Project

To stop and remove the containers, run:

```sh
docker-compose down
```

---

## 🐳 Useful Docker Commands

- View running containers: `docker ps`
- View logs: `docker-compose logs -f`
- Rebuild after code changes: `docker-compose up --build`

---

## 📬 Need Help?

Open an issue or contact [@marcofilho](https://github.com/marcofilho).

---
