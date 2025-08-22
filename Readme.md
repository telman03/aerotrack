# ✈️ **AeroTrack: Real-Time Flight Tracker**

AeroTrack is a real-time geospatial tracking system built to monitor and visualize live flight data. It demonstrates a stateful, event-driven architecture by ingesting live data, managing the current state of each flight, and pushing updates in real-time to connected clients.

---

## ⚙️ **Technologies Used**

- **Go (Golang):** Core backend service  
- **WebSockets:** Low-latency, real-time communication  
- **PostgreSQL:** Source of truth for flight data  
- **Docker:** Containerization for easy setup  

---

## ✨ **Core Features**

- **Live Data Ingestion:**  
  Background process simulates or ingests flight coordinate data.

- **Real-Time State Management:**  
  Maintains current location and status for each flight in memory.

- **WebSocket Broadcast:**  
  Broadcasts location changes to all connected clients.

- **Data Persistence:**  
  Persists all flight data and route history in PostgreSQL.

---

## 🚀 **Getting Started**

This project is containerized with Docker for a straightforward setup.

### **Prerequisites**

- Docker & Docker Compose  
- Go (if running outside Docker)

### **Setup**

```sh
git clone https://github.com/telman03/aerotrack.git
cd aerotrack 
```

**Start the services:**

```sh
docker-compose up --build -d
```

**Check the logs:**

```sh
docker-compose logs -f aerotrack-service
```

---

## 🗺️ **Architecture Overview**

The system follows a simple client-server model with a focus on real-time data flow.

```
+-------------------+         WebSocket         +---------------------+
|   Go Backend      | <-----------------------> |   Client Dashboard  |
| (Data, DB, WS)    |                          | (Map Visualization) |
+-------------------+                          +---------------------+
         |
         | PostgreSQL
         v
+-------------------+
|   Flight Data DB  |
+-------------------+
```

- **Go Backend:**  
  - Data Simulation (Goroutine)  
  - PostgreSQL Interface  
  - WebSocket Server

- **PostgreSQL:**  
  Stores flight data (long-term memory)

- **Client Dashboard:**  
  (Not included) Connects via WebSockets to visualize live flight paths

---
