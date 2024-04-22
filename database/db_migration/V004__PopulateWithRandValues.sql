INSERT INTO gates (gate_number, floor_number)
SELECT
    floor(random() * 100) + 1,
    floor(random() * 10) + 1
FROM generate_series(1, 23)
ON CONFLICT (gate_number) DO NOTHING;

INSERT INTO airlines (name)
SELECT
    substr(md5(random()::text), 0, 10)
FROM generate_series(1, 24);
