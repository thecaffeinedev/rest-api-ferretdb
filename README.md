# REST-API With Go and FerretDB

This repository contains the code of CRUD operations on [FerretDB](https://www.ferretdb.io/) written in Golang. Using HTTP requests, you can create, read, update, and delete users from the [FerretDB](https://www.ferretdb.io/) instance.

> Note: This is not a production-ready application. I was just trying out FerretDB.
> 

## Get Started

Development environment needs:

- [Docker](https://www.docker.com/get-started)
- [Task](https://taskfile.dev/installation/)

Make sure docker is installed and running. And the Task tool is installed.

## **How to run?**

First, Clone the repository 

```bash
git clone https://github.com/thecaffeinedev/rest-api-ferretdb.git
```

Next, change the current directory to the repository:

```bash
cd rest-api-ferretdb
```

Create file `.env` with content copied from `.env.sample`.  There won't be any need of changing the values. ****

Next, run the application:

```bash
task run
```

After the build is successful, You can access the application on port `8080` 

## Endpoints

```markdown
| Name        | HTTP Method | Route            |
|-------------|-------------|----------------- |
| Health      | GET         | /alive           |
|             |             |                  |
| Get User    | GET         | /api/user/{email}|
| Create User | POST        | /api/user        |
| Update User | PUT         | /api/user/{email}|
| Delete User | DELETE      | /api/user/{email}|
```

### Create User

This endpoint inserts a document in the `users` collection of the `users` database.Send a `POST` request to `/api/user`:

```bash
curl -X POST \
  'http://127.0.0.1:8080/api/user' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "Anakin",
  "email": "anakin@gmail.com",
  "password": "skywalker"
}'
```

Response:

```json
{
  "id": "ObjectID(\"6420090d2b2859af4fc17b4d\")",
  "name": "Anakin",
  "email": "anakin@gmail.com",
  "password": "skywalker"
}
```

### Get User

This endpoint retrieves a user given the email. Send a `GET` request to `/api/user/{email}`:

```bash
curl -X GET \
  'http://127.0.0.1:8080/api/user/anakin@gmail.com' \
  --header 'Accept: */*'
```

Response:

```json
{
  "id": "ObjectID(\"6420090d2b2859af4fc17b4d\")",
  "name": "Anakin",
  "email": "anakin@gmail.com",
  "password": "skywalker"
}
```

### Update User

This endpoint updates the provided fields within the specified document filtered by email. Send a `PUT` request to `/api/user/{email`}:

```bash
curl -X PUT \
  'http://127.0.0.1:8080/api/user/anakin@gmail.com' \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
  --data-raw '{"password": "padme"}'
```

Response:

```json
{
  "id": "",
  "name": "",
  "email": "anakin@gmail.com",
  "password": "padme"
}
```

### Delete User

This endpoint deletes the user from the database given the email.Send a `DELETE` request to `/api/user/{email}`:

```bash
curl -X DELETE \
  'http://127.0.0.1:8080/api/user/anakin@gmail.com' \
  --header 'Accept: */*'
```

Response:

```json
{
  "message": "Successfully Deleted"
}
```

## Technologies Used:

- [Docker + Docker-Compose](https://docs.docker.com/compose/)
- [FerretDB](https://docs.ferretdb.io/)
- [Go](https://go.dev/) (Of course)
- [Task](https://taskfile.dev/)

## Frameworks + Libraries Used

- [sirupsen/logrus](https://github.com/sirupsen/logrus)
- [Go-Chi](https://go-chi.io/#/)
- [Godotenv](https://github.com/joho/godotenv)

## EndNotes

I would recommend checking the FerretDB Github Repo. They are doing some awesome work.