CREATE TABLE gates (
    id UUID DEFAULT uuid_generate_v4(),
    gate_number INT UNIQUE NOT NULL,
    floor_number INT NOT NULL,
    PRIMARY KEY (id)
);