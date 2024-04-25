DO $$
BEGIN
    FOR i IN 1..21 LOOP
        INSERT INTO aircrafts (airline_id, wing_number, type, capacity, year_of_manufacture)
        VALUES (
            (SELECT id FROM airlines ORDER BY RANDOM() LIMIT 1),
            'Wing' || i,
            CASE random() * 3
                WHEN 0 THEN 'passenger'
                WHEN 1 THEN 'cargo'
                ELSE 'helicopter'
            END,
            FLOOR(random() * 300) + 100,  
            FLOOR(random() * 50) + 1974
        );
    END LOOP;
END $$;