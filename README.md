# 🌟 **Student Profile:**

🎓 **Student ID:** 220103167  
📚 **Subject:** Go Programming Language  
📝 **Assignment:** Graded Project 6 – **Simple REST API Server**

---

---

# 📚 **Project 6: Simple REST API Server**

## 🎯 **Objective:**
Create a RESTful API for a mock Bookstore system using the **Go programming language**.

---

## 📝 **Guidelines:**
- 🖥️ Use the `net/http` package to set up the server.
- 🛠️ **Endpoints to implement**:
  - `GET /books`: Retrieve all books.
  - `POST /books`: Add a new book.
  - `DELETE /books/{id}`: Delete a book by its ID.
- 📦 Store the data using slices or a map.
- 🔁 Return **JSON** responses for each endpoint.

---

## ⚡ **Requirements:**
- Make sure you have the **latest version of Go** installed. You can check and update it by following the instructions [here](https://golang.org/doc/install).

---

## 🧭 **Setting up Swagger Documentation:**

1. To initialize the Swagger documentation, run the following command:
    ```bash
    swag init --parseDependency --parseInternal -g ./cmd/main.go
    ```

2. To start the server, execute:
    ```bash
    go run cmd/main.go
    ```

---

## 🌐 **Accessing Swagger Documentation:**
Once the server is running, access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```
