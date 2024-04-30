CREATE TABLE aircrafts (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    tail_number VARCHAR(100) UNIQUE NOT NULL,
    capacity int NOT NULL,
    -- type VARCHAR(255) NOT NULL CHECK (type IN ('passenger', 'cargo', 'helicopter')),
    year_of_manufacture INT,
    -- created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    airline_id UUID,
    FOREIGN KEY (airline_id) REFERENCES airlines(id) ON DELETE CASCADE
);
