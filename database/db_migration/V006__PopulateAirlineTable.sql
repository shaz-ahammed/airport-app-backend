DELETE FROM airlines;

INSERT INTO airlines (name, count)
SELECT
    substr(md5(random()::text), 0, 10),
    floor(random() * 10) + 1
FROM generate_series(1, 24);