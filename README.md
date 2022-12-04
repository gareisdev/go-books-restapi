
# GO Books REST API

[![Go Report](https://goreportcard.com/badge/github.com/gareisdev/go-books-restapi)](https://goreportcard.com/report/github.com/gareisdev/go-books-restapi)

This is a little REST API built with GO and uses MariaDB to store data.

Packages:

+ GORM
+ Mux

## Endpoints

### Get all books

```http
  GET /api/books
```

#### Schema response

```json
[
  {
    "ID": null,
    "title": "",
    "author": "",
    "publication": "",
    "UpdatedAt": "2022-12-04T00:00:00.000Z",
    "CreatedAt": "2022-12-04T00:00:00.000Z"
  }
]
```

### Get book

```http
  GET /api/books/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Id of book to fetch |

#### Schema response

```json
{
    "ID": null,
    "title": "",
    "author": "",
    "publication": "",
    "UpdatedAt": "2022-12-04T00:00:00.000Z",
    "CreatedAt": "2022-12-04T00:00:00.000Z"
}
```

### Create book

```http
  POST /api/books
```

#### Body

```json
{
    "title": "",
    "author": "",
    "publication": "",
}
```

#### Schema response

```json
{
    "ID": null,
    "title": "",
    "author": "",
    "publication": "",
    "UpdatedAt": "2022-12-04T00:00:00.000Z",
    "CreatedAt": "2022-12-04T00:00:00.000Z"
}
```

### Update book

```http
  PUT /api/books/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Id of book to fetch |

#### Body

```json
{
    "title": "",
    "author": "",
    "publication": "",
}
```

#### Schema response

```json
{
    "ID": null,
    "title": "",
    "author": "",
    "publication": "",
    "UpdatedAt": "2022-12-04T00:00:00.000Z",
    "CreatedAt": "2022-12-04T00:00:00.000Z"
}
```

### Delete book

```http
  DELETE /api/books/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Id of book to fetch |

#### Schema response

```json
{
    "ID": null,
    "title": "",
    "author": "",
    "publication": "",
    "UpdatedAt": "2022-12-04T00:00:00.000Z",
    "CreatedAt": "2022-12-04T00:00:00.000Z"
}
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/gareisdev/go-books-restapi
```

Go to the project directory

```bash
  cd go-books-restapi
```

Up containers

```bash
  docker-compose up
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
