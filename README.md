# Bookshelf API with Golang Native
This is a RESTful API for storing books with almost zero configuration, as it doesn't integrate with any database. Instead, this backend app utilizes an array to store all data. However, it's important to note that when the web server is turned off, the data will be lost permanently.

## Prerequisites
- Golang
- gorilla/mux package

## Setting Up Project 
Please fork this repository first, then clone the repository that you have forked. You can clone your forked repository by using this following command:
```
git clone https://github.com/YOUR GITHUB USERNAME/go-bookshelf.git
```

This project written with native golang programming. No framework used in this project.

To support this application, I use [gorilla/mux](https://github.com/gorilla/mux). Gorilla Mux provides a powerful and flexible routing system, allowing you to define complex URL patterns, handle variables, and extract data from URLs easily. It supports both static and dynamic routing, making it suitable for a wide range of use cases.


Then, open a terminal like bash, zsh, command prompt or powershell and change the directory to where the cloned file is located.  

Finally, you can run the API with the following command:
```
go build && ./<the .exe file>
```
or

```
go run main.go
```

This project will run on `https://localhost:8080`


## Testing the API
You can try to hit the API using the POSTMAN application.

### API Endpoints
| METHOD   | ENDPOINT                 | FUNCTION |
| -------- | ------------------------ | ----------------------- |
| [GET]    | `/api/v1`                | to see if the API works |
| [POST]   | `/api/v1/book/insert`    | insert new book         |
| [GET]    | `/api/v1/book/show-all`  | show all books          |
| [GET]    | `/api/v1/book/find/{id}` | get book by id          |
| [PUT]    | `/api/v1/edit/{id}`      | edit book by id         |
| [DELETE] | `/api/v1/delete{id}`     | delete book by id       |

### API Payload
```json
{
    "title": string,
    "author": string,
    "publisher": string,
    "pageCount": int,
    "readPage": int,
}
```