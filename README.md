# Blog Platform

A full-stack blog platform built with **Go (Gin)** and **PostgreSQL**, featuring authentication, posts, comments, and likes.  
This project was developed as a pet project to practice real-world backend architecture and frontend integration.

---

## Features

### Authentication
- User registration and login
- JWT-based authentication
- Protected routes using middleware
- Secure password hashing

### Posts
- Create, read, update, delete posts
- Only the author can edit or delete their post
- Posts ordered by creation date

### Likes
- Like / Unlike posts
- One like per user per post
- Like count displayed
- “Liked by me” state supported

### Comments
- Add comments to posts
- Delete only your own comments
- Comments loaded dynamically

### Frontend
- Vanilla HTML, CSS, JavaScript
- Bootstrap for layout and styling
- Served directly from Gin
- Pages:
  - Home (posts list)
  - Single post page
  - Login / Register
  - Create post

---

## Tech Stack

**Backend**
- Go
- Gin
- PostgreSQL
- JWT

**Frontend**
- HTML
- CSS
- JavaScript
- Bootstrap

---

## Project Structure

```text
blog-platform/
├── main.go
├── routes/
│   └── routes.go
├── handlers/
│   ├── auth.go
│   ├── post.go
│   ├── comment.go
│   └── like.go
├── repository/
│   ├── user_repository.go
│   ├── post_repository.go
│   ├── comment_repository.go
│   └── like_repository.go
├── middleware/
│   └── auth.go
├── models/
│   ├── user.go
│   ├── post.go
│   └── comment.go
├── utils/
│   ├── jwt.go
│   └── password.go
├── database/
│   └── db.go
├── public/
│   ├── index.html
│   ├── login.html
│   ├── register.html
│   ├── post.html
│   ├── create-post.html
│   ├── css/
│   └── script/
└── README.md
```

---

## Authentication Flow

1. User registers or logs in
2. Server returns a JWT token
3. Token is stored in `localStorage`
4. Token is sent in `Authorization: Bearer <token>` header
5. Middleware validates token and extracts `user_id`
6. Handlers use `user_id` for authorization

---

## Authorization Rules

- Only authenticated users can:
  - Create posts
  - Like / unlike posts
  - Add comments
- Only the author can:
  - Delete their post
  - Delete their comment
- Backend always enforces ownership checks

---

## API Overview

### Auth
```text
POST /auth/register
POST /auth/login
```

### Posts
```text
GET    /posts
GET    /posts/:id
POST   /posts
PUT    /posts/:id
PATCH  /posts/:id
DELETE /posts/:id
```

### Likes
```text
POST   /posts/:id/like
DELETE /posts/:id/like
GET    /posts/:id/likes
```

### Comments
```text
GET    /posts/:id/comments
POST   /posts/:id/comments
DELETE /comments/:id
```

---

## Database

PostgreSQL is used as the main database.

Tables:
- users
- posts
- comments
- likes

Relationships are enforced using foreign keys and ownership checks in queries.

---

## Running the Project

### 1. Clone the repository
```bash
git clone https://github.com/Kair97/GO_RESTAPI_Mini_Blog-Platform.git
cd GO_RESTAPI_Mini_Blog-Platform
```

### 2. Create `.env`
```
DB_URL=postgres://user:password@localhost:5432/blog
JWT_SECRET=your_secret_key
```

### 3. Run the server

```bash
go run main.go
```

### 4. Open in browser
```
http://localhost:8080
```

---

## Project Goals

* Practice clean backend architecture
* Learn JWT authentication
* Implement proper authorization
* Connect backend with frontend
* Build a realistic full-stack application

---

## Future Improvements

* Edit post UI
* Pagination
* User profiles
* Improved error handling
* Modal confirmations
* Deployment

---

## Author

Built by **Orynbek Kairzhan**
Pet project · Learning · Growth
