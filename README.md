
# To-do Application

Task management application written in Go. This app helps to manage tasks. User can create, update, delete and sort by status tasks.


## Getting Started

### Swagger
You can use doccumentation for understanding endpoints of project. After successfull downloading, you can use:
`http://localhost:8080/swagger/index.html`


### Requests

1. **Create task:** ``POST /api/todo-list/tasks``

```json
Request:
{
  "title": "Do homework"
}

Response:
{
    "id": 1,
    "title": "Do homework",
    "activeAt": "16 Jul 24 15:16 +0545",
    "completed": false
}
```
Fields are required for filling, else user can face with 404 error


2. **Update existing tasks:** `PUT /api/todo-list/tasks/{id}`

```json
Request:
{
  "title": "Do homework"
  "activeAt": ""15 Jul 24 17:30 +0545""
}

Response:
{
    "id": 1,
    "title": "Deploy project",
    "activeAt": "16 Jul 24 15:17 +0545",
    "completed": false
}
```
Field {id} is required

3. **Delete task:** ```DELETE /api/todo-list/tasks/{id}```

```json
Response:
{
    "message": "Task deleted"
}
```
Field {id} is required

4. **Mark as done task:** ```PUT /api/todo-list/tasks/{ID}done```

```json
Response:
{
    "id": 2,
    "title": "Deploy project",
    "activeAt": "16 Jul 24 15:22 +0545",
    "completed": true
}
```
The field "compleated" change value to true  
{Id} is required


5. **Get tasks by status:** ```GET /api/todo-list/tasks?status=active``` or ```GET /api/todo-list/tasks?status=done```

```json
[
    {
        "id": 3,
        "title": "Study golang",
        "activeAt": "16 Jul 24 15:29 +0545",
        "completed": false
    },
    {
        "id": 4,
        "title": "Run outside",
        "activeAt": "16 Jul 24 15:29 +0545",
        "completed": false
    },
    {
        "id": 5,
        "title": "Run outside",
        "activeAt": "16 Jul 24 15:30 +0545",
        "completed": false
    },
    {
        "id": 6,
        "title": "Run outside",
        "activeAt": "16 Jul 24 15:30 +0545",
        "completed": false
    }
]
```

