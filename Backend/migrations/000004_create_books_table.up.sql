CREATE TABLE IF NOT EXISTS books (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    author_id BIGINT UNSIGNED NOT NULL,
    author VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    isbn VARCHAR(255) NOT NULL UNIQUE KEY,
    genre VARCHAR(255) NOT NULL,
    description TEXT,
    cover_image VARCHAR(500),
    published_year YEAR,
    total_copies INT NOT NULL DEFAULT 1,
    available_copies INT NOT NULL DEFAULT 1,

    file_path VARCHAR(500),
    file_size_bytes BIGINT,
    file_format ENUM('pdf', 'epub'),
    is_digital BOOLEAN NOT NULL DEFAULT FALSE,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE,

    INDEX idx_books_author (author),
    INDEX idx_books_genre (genre),
    FULLTEXT idx_books_search (title, author, description)  -- for full-text search
);