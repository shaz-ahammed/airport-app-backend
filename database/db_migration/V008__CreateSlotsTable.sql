CREATE TABLE slots (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    end_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    status VARCHAR(255) DEFAULT 'Available' NOT NULL,
    aircraft_id UUID  REFERENCES airlines (id),
    gate_id UUID REFERENCES gates (id)
);
