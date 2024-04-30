INSERT INTO aircrafts (airline_id, tail_number, year_of_manufacture, capacity)
SELECT
    (SELECT id FROM airlines ORDER BY RANDOM() LIMIT 1),
    substr(md5(random()::text), 1, 10),
    FLOOR(random() * 50) + 1974,
    FLOOR(random() * 50) + 30
FROM generate_series(1, 24);


-- DO $$
-- BEGIN
--     FOR i IN 1..21 LOOP
--         INSERT INTO aircrafts (airline_id, wing_number, type, capacity, year_of_manufacture)
--         INSERT INTO aircrafts (airline_id, wing_number, year_of_manufacture)
--         VALUES (
--             (SELECT id FROM airlines ORDER BY RANDOM() LIMIT 1),
--             'Wing' || i,
--             CASE random() * 3
--                 WHEN 0 THEN 'passenger'
--                 WHEN 1 THEN 'cargo'
--                 ELSE 'helicopter'
--             END,
--             FLOOR(random() * 300) + 100,
--             FLOOR(random() * 50) + 1974
--         );
--     END LOOP;