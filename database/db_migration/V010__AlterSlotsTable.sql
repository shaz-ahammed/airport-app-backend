ALTER TABLE slots
    DROP CONSTRAINT IF EXISTS aircraft_id_foreign_key_name,
    DROP CONSTRAINT IF EXISTS gate_id_foreign_key_name,
    ADD CONSTRAINT aircraft_id_foreign_key_name FOREIGN KEY (aircraft_id) REFERENCES airlines (id) ON DELETE CASCADE,
    ADD CONSTRAINT gate_id_foreign_key_name FOREIGN KEY (gate_id) REFERENCES gates (id) ON DELETE CASCADE;
