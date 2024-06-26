
# XYZ Books App Server

Server for XYZ Books CRUD APIs.

## Prerequisites
- Golang

## Tech Stack
- **Golang 1.22** - Language
- **Gin Gonic** - Web Framework
- **SQLite** - Embedded Database
- **GORM** - ORM Library
- **RabbitMQ** - Message Queue

## Installation

1. Clone repository
2. Navigate to `/server`
3. Run `go mod tidy`
4. Run `go run cmd/main.go`
5. Access the APIs via [localhost:8080](http://localhost:8080)

**Note:** 
- The server code already includes a `gorm.db` file with preloaded data. Delete this to have a fresh empty database.
- Rabbit MQ is configured to work when the server is ran via Docker.


## Author

- [Engr. Kirk Alyn Santos](https://github.com/kirkalyn13)

## API Reference

#### Get all books

```http
  GET /api/v1/books?q
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `q` | `string` | ISBN 13 of the books to filter |

#### Get book by ID

```http
  GET /api/v1/books/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the book to fetch |

#### Get book by ISBN13

```http
  GET /api/v1/books/isbn13/${isbn13}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `isbn13`      | `string` | **Required**. ISBN 13 of the book to fetch |

#### Get books with incomplete ISBNs

```http
  GET /api/v1/books/isbn/incomplete
```

#### Get all authors

```http
  GET /api/v1/authors?q
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `q`      | `string` | Name of the author to filter |

#### Get author by ID

```http
  GET /api/v1/authors/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the author to fetch |

#### Get all publishers

```http
  GET /api/v1/publishers?q
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `q`      | `string` | Name of the publisher to filter |

#### Get publisher by ID

```http
  GET /api/v1/publishers/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the publisher to fetch |

#### Add book

```http
  POST /api/v1/books
```

#### Add author

```http
  POST /api/v1/authors
```

#### Add publisher

```http
  POST /api/v1/publishers
```

#### Edit book

```http
  PUT /api/v1/books/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the book to edit |


#### Edit author

```http
  PUT /api/v1/authors/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the author to edit |

#### Edit publisher

```http
  PUT /api/v1/publishers/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the publisher to edit |

#### Delete book

```http
  DELETE /api/v1/books/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the book to delete |

#### Delete author

```http
  DELETE /api/v1/authors/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the author to delete |

#### Delete book

```http
  DELETE /api/v1/publishers/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of the publisher to delete |
