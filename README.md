# DinoDO - Go-based Todo List API

<p align="center">
<img src="https://i.imgur.com/FHyvrxG.png" alt="DinoDO" style="height: 300px;"/>
</p>

## Description

DinoDO is a simple Todo List API developed in Go, utilizing the following tools:

- [gin-gonic](https://github.com/gin-gonic/gin) - Web framework for Go
- [gorm](https://gorm.io/) - Go ORM
- [uuid](https://github.com/google/uuid) - UUID generation in Go
- [godotenv](https://github.com/joho/godotenv) - Loading environment variables in Go

The application uses a MySQL database to store tasks.

## Features

- **Task Creation**: Add new tasks to the list.
- **Task Editing**: Update the content or status of an existing task.
- **Task Listing**: View all available tasks.
- **Task Listing By ID**: View task by provided ID.
- **Delete Tasks**: Remove tasks from the list.
- **JWT Authentication**: Secure routes with JSON Web Token authentication.

## Configuration

1. Install dependencies:

    ```bash
        go mod tidy
    ```

2. Set up environment variables:

   - Create a `.env` file in the project root and configure the necessary variables. An example is provided in the `.env.example` file.

3. Run the application:

    ```bash
    go run main.go
    ```

   The application will be available at `http://localhost:4000`.

## Usage Examples

### Task Creation

```bash
curl -X POST http://localhost:4000/api/v1/tasks -d '{"task": "New Task", "finished":false}'
```

### Task Editing

```bash
curl -X PUT http://localhost:4000/api/v1/tasks/{id} -d '{"title": "Updated Task", "finished": true}'
```

### Task Listing

```bash
curl http://localhost:4000/api/v1/tasks
```


### Task Listing By ID

```bash
curl http://localhost:4000/api/v1/tasks/{id}
```

### Delete Task

```bash
curl -X DELETE http://localhost:4000/api/v1/tasks/{id}
```

### User Register

```bash
curl -X POST http://localhost:4000/auth/register -d '{"name": "User Name","email": "user@email.com", "username": "example", "password": "password"}'
```

### Get User Details

```bash
curl http://localhost:4000/users/{id} -H "Authorization: Bearer <your_token>"
```

## Contributions

Contributions are welcome! Feel free to open issues and pull requests.

## License

This project is for educational purposes only and is licensed under the MIT License.
[MIT License](LICENSE).


