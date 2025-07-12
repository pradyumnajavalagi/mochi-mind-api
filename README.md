# Kanji Learning API (Mochi Mind API)

A RESTful API backend for a Kanji learning application built with Go. This application provides endpoints for managing Japanese Kanji flashcards with image upload capabilities.

## 🎯 Features

- **Flashcard Management**: Create, read, update, and delete Kanji flashcards
- **Image Upload**: Upload and serve Kanji character images
- **Random Flashcards**: Get random flashcards for study sessions
- **PostgreSQL Database**: Persistent storage for flashcard data
- **RESTful API**: Clean HTTP endpoints for frontend integration

## 🏗️ Project Structure

```
kanji-app-backend/
├── main.go              # Application entry point
├── go.mod               # Go module dependencies
├── go.sum               # Dependency checksums
├── models/
│   └── models.go        # Database models and operations
├── router/
│   └── router.go        # HTTP route definitions
├── middleware/
│   └── handlers.go      # HTTP request handlers
├── uploads/             # Image storage directory
└── mochi-mind-api.exe   # Compiled executable
```

## 🚀 Getting Started

### Prerequisites

- Go 1.22.4 or higher
- PostgreSQL database
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd kanji-app-backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgresql://username:password@localhost:5432/database_name
   ```

4. **Set up the database**
   Create a PostgreSQL database and run the following SQL to create the flashcards table:
   ```sql
   CREATE TABLE flashcards (
       id SERIAL PRIMARY KEY,
       kanji_image_url TEXT NOT NULL,
       onyomi VARCHAR(255),
       kunyomi VARCHAR(255),
       example_usage TEXT
   );
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 📚 API Endpoints

### Flashcards

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/flashcards` | Get all flashcards |
| `POST` | `/flashcards` | Create a new flashcard |
| `PUT` | `/flashcards/{id}` | Update a flashcard |
| `DELETE` | `/flashcards/{id}` | Delete a flashcard |
| `GET` | `/flashcards/random` | Get random flashcards |

### File Upload

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/upload` | Upload a Kanji image file |

### Static Files

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/uploads/{filename}` | Serve uploaded images |

## 📝 API Usage Examples

### Create a Flashcard
```bash
curl -X POST http://localhost:8080/flashcards \
  -H "Content-Type: application/json" \
  -d '{
    "kanji_image_url": "http://localhost:8080/uploads/kanji_image.jpg",
    "onyomi": "カン",
    "kunyomi": "かん",
    "example_usage": "漢字 (かんじ) - Chinese characters"
  }'
```

### Get All Flashcards
```bash
curl http://localhost:8080/flashcards
```

### Upload an Image
```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@path/to/kanji_image.jpg"
```

### Get Random Flashcards
```bash
curl http://localhost:8080/flashcards/random
```

## 🗄️ Database Schema

### Flashcards Table
```sql
CREATE TABLE flashcards (
    id SERIAL PRIMARY KEY,
    kanji_image_url TEXT NOT NULL,
    onyomi VARCHAR(255),
    kunyomi VARCHAR(255),
    example_usage TEXT
);
```

## 🛠️ Dependencies

- **gorilla/mux**: HTTP router and URL matcher
- **jackc/pgx/v5**: PostgreSQL driver
- **joho/godotenv**: Environment variable management

## 🔧 Configuration

The application uses environment variables for configuration:

- `DATABASE_URL`: PostgreSQL connection string (required)

## 📁 File Storage

Uploaded images are stored in the `uploads/` directory with timestamped filenames to prevent conflicts. The application automatically handles file cleanup when flashcards are updated or deleted.

## 🚀 Deployment

### Build the Application
```bash
go build -o mochi-mind-api main.go
```

### Run the Executable
```bash
./mochi-mind-api
```

## 🔒 Security Considerations

- Ensure proper database credentials and connection security
- Consider implementing authentication for production use
- Validate file uploads to prevent malicious files
- Use HTTPS in production environments

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions, please open an issue in the repository or contact the development team.

---

**Note**: This API is designed to work with a frontend application for a complete Kanji learning experience. The image URLs returned are configured for Android emulator compatibility (`10.0.2.2:8080`). 