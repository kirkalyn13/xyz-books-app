
# XYZ Books App

XYZ Books CRUD UI application.

## Prerequisites

- Docker

## Installation

1. Clone repository.
2. Navigate to the root folder of the project where the `docker-compose.yaml` file is located.
3. Run `docker compose up`. This should build and run the Docker images for both client and server, and download the Rabbit MQ image. Three docker containers are expected to run from this command, which may take a couple of minutes to complete.
4. Access the UI via [localhost:4173](http://localhost:4173)
5. Run the pipeline from [xyz-books-pipeline](https://github.com/kirkalyn13/xyz-books-pipeline). Note that the pipeline is dependent to the Rabbit MQ instance.

### Images

- `xyz-books-app-client`
- `xyz-books-app-server`
- `rabbitmq:3-management`

**Note:** The client and server could be ran separately. Instructions are specified on the readmes of each folder.

## Author

- [Engr. Kirk Alyn Santos](https://github.com/kirkalyn13)
