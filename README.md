# Hotstar Mini API

A simple streaming platform backend built with Go, GORM, and PostgreSQL.

## Features
- **Media Management**: Add, upload, and retrieve media content.
- **Categorization**: Organize media by categories and genres.
- **Ratings**: Users can rate media content.
- **File Uploads**: Supports uploading thumbnails and banners.
- **Static Serving**: Serves uploaded images directly.

## Prerequisites
- Go 1.25+
- PostgreSQL

## Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd hotstar_mini
   ```

2. **Configure Environment Variables**:
   Create a `.env` file in the root directory (or use the existing one):
   ```env
   PORT=8080
   DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
   ```

3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the Application**:
   ```bash
   go run main.go
   ```

## API Endpoints

### Media
- `GET /media` - Get all media
- `GET /media_detail?id={id}` - Get media by ID
- `POST /media` - Upload media (Multipart Form)
  - Fields: `title`, `desc`, `thumbnail` (file), `banner` (file), `age_rating`, `release_year`, `category_id`, `genre` (comma-separated).

### Category & Genre
- `POST /category` - Create a category
- `GET /category` - Get all categories
- `POST /genre` - Create a genre

### Ratings
- `POST /rating` - Submit a rating

## Project Structure
- `handlers/`: HTTP request handlers and logic.
- `models/`: GORM database models.
- `db/`: Database initialization.
- `uploads/`: Directory where uploaded files are stored.
