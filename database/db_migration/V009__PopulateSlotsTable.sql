INSERT INTO slots (start_time, end_time, status, aircraft_id, gate_id)
SELECT
    NOW() - (random() * INTERVAL '30 days') AS start_time,
    NOW() + ((random() * 2 + 2) * INTERVAL '1 hour') AS end_time,
    CASE WHEN random() < 0.5 THEN 'Reserved' ELSE 'Booked' END AS status,
    (SELECT id FROM airlines ORDER BY random() LIMIT 1) AS aircraft_id,
    (SELECT id FROM gates ORDER BY random() LIMIT 1) AS gate_id
FROM generate_series(1, 16);

INSERT INTO slots (start_time, end_time, status)
SELECT
    NOW() - (random() * INTERVAL '30 days') AS start_time,
    NOW() + ((random() * 2 + 2) * INTERVAL '1 hour') AS end_time,
    'Available' AS is_available
FROM generate_series(1, 12);
