Distributed Arithmetic Expression Calculator


This project is a distributed arithmetic expression calculator that can be used to calculate the result of an arithmetic expression asynchronously. The calculator is designed to be horizontally scalable and can handle a large number of concurrent requests.

Architecture
The calculator is composed of two main parts: the orchestrator and the agent.

The orchestrator is responsible for accepting arithmetic expressions from clients and storing them in a PostgreSQL database. The orchestrator also provides an HTTP API for querying the status of a task and retrieving the result of a completed task. The orchestrator uses a Redis database to store task status information and a RabbitMQ queue to distribute tasks to agents.

The agent is responsible for consuming tasks from the RabbitMQ queue and processing them. The agent uses the math/big package to perform the arithmetic calculations and stores the result of the calculation in the PostgreSQL database.
To run the distributed calculator project, you will need to have the following prerequisites installed:

Running the application

To run the application, you need to have Go installed on your machine.

1. Clone the repository:

```bash
git clone https://github.com/dm4brl/distributed-calculator.git
