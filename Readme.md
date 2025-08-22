AeroTrack: Real-Time Flight Tracker
AeroTrack is a real-time geospatial tracking system built to monitor and visualize live flight data. It demonstrates a stateful, event-driven architecture by ingesting live data, managing the current state of each flight, and pushing updates in real-time to connected clients.

⚙️ Technologies Used
Go (Golang): The core backend service.

WebSockets: For low-latency, real-time communication between the server and a client dashboard.

PostgreSQL: The source of truth for storing all flight data and its history.

Docker: For containerizing the application and its database for easy setup.

✨ Core Features
Live Data Ingestion: A background process simulates or ingests flight coordinate data into the main tracking service.

Real-Time State Management: The service maintains the current location and status for each flight in memory, allowing for fast lookups.

WebSocket Broadcast: When a flight's location changes, the server broadcasts the new coordinates to all connected clients.

Data Persistence: All flight data, including route history, is persisted in a PostgreSQL database.

🚀 Getting Started
This project is containerized with Docker for a straightforward setup.

Prerequisites
Docker and Docker Compose

Go (if running outside of Docker)

Setup
Clone the repository:

git clone [your-repository-url]
cd aerotrack

Start the services:
This command will build the Go application, start the PostgreSQL database, and run all services in the background.

docker-compose up --build -d

Check the logs:
You can view the logs of your application to see the data ingestion and WebSocket activity.

docker-compose logs -f aerotrack-service

🗺️ Architecture Overview
The system follows a simple client-server model with a focus on real-time data flow.

A simplified client-server architecture diagram with a WebSocket connection would go here.

Go Backend: The central Go application handles three key functions:

Data Simulation: A Goroutine within the main service simulates live flight data.

PostgreSQL Interface: Reads and writes flight data to the database.

WebSocket Server: Manages all client connections, receiving requests and broadcasting real-time updates.

PostgreSQL: Stores the flight data, serving as the system's long-term memory.

Client Dashboard (not included in this repo): A hypothetical frontend (e.g., built with React or Vue) that connects to the Go backend via WebSockets to visualize the live flight paths on a map.

🎓 What I Learned
Designing and implementing a real-time backend using Go's net/http package and WebSockets.

Managing in-memory application state for low-latency lookups.

The importance of concurrency in Go for handling multiple client connections and background data ingestion simultaneously.

Working with Docker Compose for local development and dependency management.

Persisting complex data structures, such as geospatial coordinates, in a PostgreSQL database.