# User Service - GraphQL API

A GraphQL-based user management service built with Go, GORM, and PostgreSQL.

## Features

- User CRUD operations (Create, Read, Update, Delete)
- JWT-based authentication
- Password hashing with bcrypt
- Role-based access (manager/member)
- GraphQL API with playground

## Getting Started

### Prerequisites

- Go 1.24+
- PostgreSQL
- Environment variables configured (see `.env.example`)

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up your environment variables
4. Run the service:
   ```bash
   go run main.go
   ```

The service will start at `http://localhost:8080`
- GraphQL Playground: `http://localhost:8080`
- GraphQL Endpoint: `http://localhost:8080/query`

## GraphQL Schema

### User Type
```graphql
type User {
  userId: ID!
  username: String!
  email: String!
  role: String!
  PasswordHash: String!
}
```

### Available Operations

#### Queries
- `fetchUsers`: Get all users
- `getUserByID(userId: ID!)`: Get user by ID

#### Mutations
- `createUser(username: String!, email: String!, role: String!, password: String!)`: Create new user
- `updateUser(userId: ID!, username: String, email: String, role: String, password: String)`: Update user
- `deleteUser(userId: ID!)`: Delete user
- `login(email: String!, password: String!)`: User login (returns JWT token)
- `logout`: User logout

## GraphQL Examples

### 1. Create User

```graphql
mutation CreateUser {
  createUser(
    username: "john_doe"
    email: "john.doe@example.com"
    role: "member"
    password: "securePassword123"
  ) {
    userId
    username
    email
    role
  }
}
```

### 2. Create Manager

```graphql
mutation CreateManager {
  createUser(
    username: "jane_manager"
    email: "jane.manager@example.com"
    role: "manager"
    password: "managerPassword456"
  ) {
    userId
    username
    email
    role
  }
}
```

### 3. Get All Users

```graphql
query GetAllUsers {
  fetchUsers {
    userId
    username
    email
    role
  }
}
```

### 4. Get User by ID

```graphql
query GetUserById {
  getUserByID(userId: "your-user-id-here") {
    userId
    username
    email
    role
  }
}
```

### 5. Update User (All Fields)

```graphql
mutation UpdateUser {
  updateUser(
    userId: "your-user-id-here"
    username: "john_updated"
    email: "john.updated@example.com"
    role: "manager"
    password: "newSecurePassword789"
  ) {
    userId
    username
    email
    role
  }
}
```

### 6. Update User (Password Only)

```graphql
mutation UpdatePassword {
  updateUser(
    userId: "your-user-id-here"
    password: "brandNewPassword123"
  ) {
    userId
    username
    email
    role
  }
}
```

### 7. Update User (Username and Email Only)

```graphql
mutation UpdateUserInfo {
  updateUser(
    userId: "your-user-id-here"
    username: "new_username"
    email: "new.email@example.com"
  ) {
    userId
    username
    email
    role
  }
}
```

### 8. Delete User

```graphql
mutation DeleteUser {
  deleteUser(userId: "your-user-id-here")
}
```

### 9. User Login

```graphql
mutation Login {
  login(
    email: "john.doe@example.com"
    password: "securePassword123"
  )
}
```

### 10. User Logout

```graphql
mutation Logout {
  logout
}
```

## Complete CRUD Workflow Example

Here's a complete workflow demonstrating all CRUD operations:

```graphql
# 1. Create a new user
mutation Step1_CreateUser {
  createUser(
    username: "test_user"
    email: "test@example.com"
    role: "member"
    password: "testPassword123"
  ) {
    userId
    username
    email
    role
  }
}

# 2. Read all users
query Step2_ReadAllUsers {
  fetchUsers {
    userId
    username
    email
    role
  }
}

# 3. Read specific user (use userId from step 1)
query Step3_ReadSpecificUser {
  getUserByID(userId: "USER_ID_FROM_STEP_1") {
    userId
    username
    email
    role
  }
}

# 4. Update the user (use userId from step 1)
mutation Step4_UpdateUser {
  updateUser(
    userId: "USER_ID_FROM_STEP_1"
    username: "updated_test_user"
    password: "newPassword456"
  ) {
    userId
    username
    email
    role
  }
}

# 5. Delete the user (use userId from step 1)
mutation Step5_DeleteUser {
  deleteUser(userId: "USER_ID_FROM_STEP_1")
}
```

## Authentication Flow Example

```graphql
# 1. Create a user
mutation CreateUserForAuth {
  createUser(
    username: "auth_user"
    email: "auth@example.com"
    role: "member"
    password: "authPassword123"
  ) {
    userId
    username
    email
  }
}

# 2. Login to get JWT token
mutation LoginUser {
  login(
    email: "auth@example.com"
    password: "authPassword123"
  )
}

# 3. Use the JWT token in your HTTP headers for authenticated requests
# Header: Authorization: Bearer YOUR_JWT_TOKEN_HERE

# 4. Logout when done
mutation LogoutUser {
  logout
}
```

## Error Handling

The API returns appropriate error messages for:
- Invalid credentials during login
- Duplicate email registration
- User not found
- Password hashing failures
- Database connection issues

## Security Features

- Passwords are hashed using bcrypt with default cost
- JWT tokens expire after 72 hours
- Email uniqueness is enforced
- Role validation for manager/member roles

## Database Schema

The service uses the following database table structure:

```sql
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    role VARCHAR(10) NOT NULL,
    password_hash VARCHAR NOT NULL
);
```

## Environment Variables

Create a `.env` file with:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=user_service
DB_PORT=5432
PORT=8080
```
