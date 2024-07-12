# BallastLane TechTest Project

## Project Structure

The project is organized as follows:

- **`./cmd/clock/main.go`**: Is the entry point for this solution. [View `./cmd/clock/main.go`](./cmd/clock/main.go)
- **`./pkg/`**: Directory with the software components of this solution. [View `./pkg/`](./pkg/)
- **`./scripts/`**: Auxiliar scripts to build docker images.

## Getting Started

To start working on this project, follow the steps below:

1. **Clone the Repository**: Clone this repository to your local machine.
2. **Install Dependencies**: `make deps`
3. **Run the solution**: `make run`
4. **To change the print tick, tock and bong values**:
```shell
curl -X POST http://localhost:8080/update-config \
-H "Content-Type: application/json" \
-d '{"run_span":10800,"tick":"quack","tock":"quock","bong":"quong"}'
```