CREATE TABLE aircrafts (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    airline_id UUID REFERENCES airlines(id) NOT NULL ON DELETE CASCADE,
    wing_number VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL CHECK (type IN ('passenger', 'cargo', 'helicopter')),
    capacity INT NOT NULL,
    year_of_manufacture INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);