INSERT INTO slots (start_time, end_time, is_available, aircraft_id, gate_id)
SELECT
    NOW() - (random() * INTERVAL '30 days') AS start_time,
    NOW() + ((random() * 2 + 2) * INTERVAL '1 hour') AS end_time,
    false AS is_available,
    (SELECT id FROM airlines ORDER BY random() LIMIT 1) AS aircraft_id,
    (SELECT id FROM gates ORDER BY random() LIMIT 1) AS gate_id
FROM generate_series(1, 16);

INSERT INTO slots (start_time, end_time, is_available)
SELECT
    NOW() - (random() * INTERVAL '30 days') AS start_time,
    NOW() + ((random() * 2 + 2) * INTERVAL '1 hour') AS end_time,
    true AS is_available
FROM generate_series(1, 12);
