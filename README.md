GO LANG
Creating a Go Lang backend app.

CRUD

CREATE

```bash
curl --location --request POST 'http://localhost:8000/books' \
--header 'Content-Type: application/json' \
--data-raw '{
    "publication": "12342",
    "name": "New Book 22",
    "author": "John Doe"
}'
```

READ

1.  GET ALL

    ```bash
    curl --location --request GET 'http://localhost:8000/books'
    ```

2.  GET ONE
    ```bash
    curl --location --request GET 'http://localhost:8000/books/12'
    ```

UPDATE

```bash
curl --location --request PUT 'http://localhost:8000/books/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "publication": "9990111",
    "name": "New Book 1090",
    "author": "John Doe 99"
}'
```

DELETE

```bash
curl --location --request DELETE 'http://localhost:8000/books/12'
```

Run app
`go run cmd/main/main.go`

Build app
`go build`
