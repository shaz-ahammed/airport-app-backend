CREATE TABLE aircrafts (
    id VARCHAR(255) PRIMARY KEY AUTO_INCREMENT,
    tail_number VARCHAR(100) UNIQUE NOT NULL,
    -- type VARCHAR(255) NOT NULL CHECK (type IN ('passenger', 'cargo', 'helicopter')),
    year_of_manufacture INT,
    -- created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    airline_id VARCHAR(255),
    FOREIGN KEY (airline_id) REFERENCES airlines(id) ON DELETE CASCADE,
);
