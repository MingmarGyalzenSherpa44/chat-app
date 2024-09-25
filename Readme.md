# CLI Chat APP

This is a command-line interface (CLI) chat application built with Go. The app allows users to register, log in, and chat with other logged-in users in real time. User data and chat history are stored in a PostgreSQL database.

## Features

- **User Registration**: Register a new user with email, username, and password.
- **User Login**: Log in using your username and password.
- **Real-Time Chat**: Chat with all logged-in users in real time.
- **Persistent Storage**: Save registered users and chat history in a PostgreSQL database.


## Libraries Used

- [Cobra](https://github.com/spf13/cobra): For creating the CLI application.
- [go-socket.io](https://github.com/googollee/go-socket.io): For real-time communication.
- [go-socket.io-client](https://github.com/hesh915/go-socket.io-client): For connecting clients to the socket server.
- [pgx](https://github.com/jackc/pgx): For interacting with the PostgreSQL database.


## Technologies Used:

- **Go**
- **PostgreSQL**

## Getting Started

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/MingmarGyalzenSherpa44/chat-app
   cd chat-app

2. **Install Dependencies**:

    ```bash
    go get .

3. **Run the application**:

    ```bash
    go run cmd/server.go