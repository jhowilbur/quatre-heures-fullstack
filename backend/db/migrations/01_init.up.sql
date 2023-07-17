CREATE TABLE IF NOT EXISTS Events (
                        ID SERIAL PRIMARY KEY,
                        Title VARCHAR(32),
                        Description TEXT,
                        InitialDate TIMESTAMP WITH TIME ZONE,
                        FinalDate TIMESTAMP WITH TIME ZONE
);

INSERT INTO Events (Title, Description, InitialDate, FinalDate)
VALUES ('Event 1', 'Description 1', '2023-07-16 10:00:00+00:00', '2023-07-16 12:00:00+00:00'),
       ('Event 2', 'Description 2', '2023-07-17 14:00:00+00:00', '2023-07-17 16:00:00+00:00'),
       ('Event 3', 'Description 3', '2023-07-18 09:00:00+00:00', '2023-07-18 11:00:00+00:00');