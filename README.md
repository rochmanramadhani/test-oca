# Test OCA Parking System

This project implements a parking management system using Event Driven Architecture and Domain Driven Design principles. It consists of multiple services orchestrated via Docker-compose.

[![Report Service Build](https://github.com/rochmanramadhani/test-oca-bak/actions/workflows/report-service.yaml/badge.svg)](https://github.com/rochmanramadhani/test-oca-bak/actions/workflows/report-service.yaml)
[![Parking Service Build](https://github.com/rochmanramadhani/test-oca-bak/actions/workflows/parking-service.yaml/badge.svg)](https://github.com/rochmanramadhani/test-oca-bak/actions/workflows/parking-service.yaml)

## Overview

This system aims to manage parking operations efficiently, leveraging modern architectural patterns for scalability and maintainability. Key features include:

- **Domain Driven Design**: Each module encapsulates domain logic, enhancing clarity and modularity within the system.

## Architecture Diagram
![Architecture Diagram](https://github.com/rochmanramadhani/bank-techschool-service/assets/115830242/1fed9b9d-ba55-43a6-8f65-dd6851980f0b)

This architecture diagram provides a high-level overview of the system components and their interactions within a microservices setup. The diagram includes a client, Nginx, and Docker containers hosting different services. Below is a detailed explanation of each component and the flow:

1. **Client**
   - The client represents the user or an application that sends requests to access various services provided by the server.

2. **Nginx (port 1500)**
   - Nginx is configured as a reverse proxy and load balancer, listening on port 1500.
   - It receives incoming requests from the client and forwards them to the appropriate service running in Docker containers.

3. **Docker**
   - Docker is used to containerize and run different services independently.
   - The Docker environment hosts the following services:
      - **Parking Service (port 8081)**
         - This service handles requests related to parking functionalities.
      - **Report Service (port 8082)**
         - This service handles requests related to reporting functionalities.

### Flow of Requests

- The client sends a request to Nginx on port 1500.
- Nginx processes the incoming request and determines which service should handle it.
- Depending on the type of request, Nginx forwards the request to either the Parking Service on port 8081 or the Report Service on port 8082.
- The respective service processes the request and sends a response back to Nginx.
- Nginx then forwards the response back to the client.

### Key Points

- **Isolation and Scalability**: By using Docker, each service is isolated in its container, making it easier to scale and manage.
- **Reverse Proxy**: Nginx acts as an intermediary, improving the security, performance, and reliability of the communication between the client and services.
- **Port Configuration**: Specific ports are assigned to each service to handle incoming requests, ensuring clear and organized traffic routing.

This architecture ensures that the system is modular, scalable, and easy to manage, providing a robust solution for handling different types of service requests.

## Sequence Diagram
![Sequence Diagram](https://github.com/rochmanramadhani/bank-techschool-service/assets/115830242/b608d4fa-a2ac-4e25-90ec-9f7c345a996a)

This sequence diagram illustrates the interaction between various components in a microservices architecture setup involving a client, server, Docker, Nginx, ParkingService, and ReportService. Below is a step-by-step breakdown of the process:

1. **Client Requests to Access Services**
   - The client sends a request to access the services hosted on the server.

2. **Server Initializes Docker Environment**
   - The server initializes the Docker environment to set up the required services.

3. **Start Nginx**
   - Nginx is started on port 1500 to act as a reverse proxy and load balancer.

4. **Check Services Availability**
   - Nginx checks the availability of the services by attempting to connect to:
      - ParkingService on port 8081
      - ReportService on port 8082
   - This check is performed in a loop until all services are available.

5. **Services are Available**
   - Once both services are confirmed to be available, Nginx responds to the server indicating that services are available.

6. **Client Requests Parking Service**
   - The client sends a request specifically for the ParkingService.

7. **Forward Request to Parking Service**
   - Nginx forwards this request to the ParkingService running on port 8081.

8. **Parking Service Responds with Data**
   - The ParkingService processes the request and sends back the required data to Nginx.

9. **Forward Parking Service Data**
   - Nginx forwards the response from the ParkingService back to the server, which then relays it to the client.

This sequence ensures that all services are up and running before handling client requests, providing a reliable and scalable architecture for service interaction and data flow.

## Installation and Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/rochmanramadhani/test-oca.git
   cd test-oca
   ```

2. Run Docker-compose to start all services:
   ```bash
   docker-compose up
   ```

3. Wait until all services are running. Once ready, you can interact with the system using the provided Postman collection.

## Usage

For detailed usage instructions and API endpoints, refer to the [Postman Collection](https://documenter.getpostman.com/view/12136790/2sA3XMhiF5#29777650-a0f6-4f09-83ca-54a0f3c797fb). This collection includes predefined test cases and scenarios to demonstrate the system's capabilities.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to submit a pull request or open an issue in the repository.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- Thanks to the open-source community for their valuable libraries and tools that make this project possible.
