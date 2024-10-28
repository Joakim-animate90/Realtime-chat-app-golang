# Realtime Chat App

This is a realtime chat application built using Go and Gin framework, designed to allow users to chat with each other and manage user information, friends, and groups. The application supports both individual and group messaging, and includes features like user registration, login, and avatar management.

## Features

- User registration, login, and information modification
- Avatar upload and management
- Add and manage friends
- Create and join groups
- Send and receive messages in real-time
- Support for different message types (text, file, image, etc.)

## Getting Started

### Prerequisites

- Go 1.16 or later
- Docker (optional, for containerized deployment)
- MySQL database

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/realtime-chat-app.git
   cd realtime-chat-app
   ```

2. Set up the database:

   - Import the `chat.sql` file into your MySQL database to create the necessary tables.

3. Configure the application:

   - Copy the `config.example.toml` to `config.toml` and update the database connection settings.

4. Build the project:

   ```bash
   make build
   ```

5. Run the application:

   ```bash
   ./bin/chat
   ```

### Using Docker

1. Build the Docker image:

   ```bash
   docker build -t realtime-chat-app .
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 realtime-chat-app
   ```

## API Endpoints

- **User Management**
  - `POST /user/register` - Register a new user
  - `POST /user/login` - Login with existing credentials
  - `PUT /user` - Modify user information
  - `GET /user/:uuid` - Get user details by UUID
  - `GET /user/name` - Get user or group by name

- **Friend Management**
  - `POST /friend` - Add a friend

- **Group Management**
  - `POST /group/:uuid` - Save group details
  - `POST /group/join/:userUuid/:groupUuid` - Join a group

- **File Management**
  - `POST /file` - Upload a file (e.g., avatar)

- **Messaging**
  - `GET /message` - Get messages

- **WebSocket**
  - `/socket.io` - WebSocket connection for real-time communication

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries, please contact [joakimbwire23@gmail.com].
