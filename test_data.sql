-- Insert entries into the 'persons' table
INSERT INTO
  persons (
    first_name,
    last_name,
    birth_date,
    death_date,
    gender,
    profile_id,
    photo_url
  )
VALUES
  (
    'Ric', -- #1
    'Boyd',
    '1946-10-06',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Dana', -- #2
    'Boyd',
    '1949-07-16',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Dava', -- #3
    'Walker',
    '1968-09-10',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Jami', -- #4
    'Garzella',
    '1970-02-24',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Richard', -- #5
    'Boyd',
    '1971-09-08',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Matthew', -- #6
    'Boyd',
    '1977-06-05',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Paul', -- #7
    'Walker',
    '1965-11-29',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Tom', -- #8
    'Garzella',
    '1968-11-18',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Christy', -- #9
    'Boyd',
    '1974-05-16',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Kelly', -- #10
    'Boyd',
    '1977-05-24',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Haleigh', -- #11
    'Forrest',
    '1990-04-20',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Hana', -- #12
    'Guess',
    '1991-06-23',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Jon', -- #13
    'Walker',
    '1994-06-29',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Trey', -- #14
    'Walker',
    '1999-06-21',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Calli', -- #15
    'Pawlicki',
    '1996-04-18',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Mikinna', -- #16
    'White',
    '1999-05-04',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Zackary', -- #17
    'Garzella',
    '2004-03-24',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Collin', -- #18
    'Boyd',
    '2000-09-12',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Kamrin', -- #19
    'Boyd',
    '2006-02-20',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Ozzie', -- #20
    'Boyd',
    '2010-02-16',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Easton', -- #21
    'Boyd',
    '2013-06-25',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Alexa', -- #22
    'Boyd',
    '2003-06-07',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Olivia', -- #23
    'Boyd',
    '2006-05-24',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'John', -- #24
    'Forrest',
    '1987-07-23',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Ian', -- #25
    'Guess',
    '1988-03-28',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Jacob', -- #26
    'Pawlicki',
    '1996-02-05',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Stephen', -- #27
    'White',
    '1998-03-15',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Desmond', -- #28
    'Guess',
    '2013-09-29',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Nona', -- #29
    'Guess',
    '2014-10-07',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Everett', -- #30
    'Guess',
    '2016-10-26',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Remington', -- #31
    'Guess',
    '2018-12-28',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Sage', -- #32
    'Guess',
    '2020-11-15',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Lewis', -- #33
    'Guess',
    '2022-08-23',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Madison', -- #34
    'Pawlicki',
    '2017-08-25',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Paisley', -- #35
    'Pawlicki',
    '2020-07-30',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Levi', -- #36
    'Pawlicki',
    '2021-09-30',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Kinsley', -- #37
    'Pawlicki',
    '2024-05-21',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'James', -- #38
    'Kruger',
    '2021-03-21',
    NULL,
    'Male',
    NULL,
    NULL
  );

-- Insert relationships among persons
INSERT INTO
  relationships (person1_id, person2_id, relationship_type)
VALUES
  (1, 2, 'Spouse'),
  (1, 3, 'Parent'),
  (1, 4, 'Parent'),
  (1, 5, 'Parent'),
  (1, 6, 'Parent'),
  (2, 3, 'Parent'),
  (2, 4, 'Parent'),
  (2, 5, 'Parent'),
  (2, 6, 'Parent'),
  (3, 7, 'Spouse'),
  (3, 11, 'Parent'),
  (3, 12, 'Parent'),
  (3, 13, 'Parent'),
  (3, 14, 'Parent'),
  (7, 11, 'Parent'),
  (7, 12, 'Parent'),
  (7, 13, 'Parent'),
  (7, 14, 'Parent'),
  (4, 8, 'Spouse'),
  (4, 15, 'Parent'),
  (4, 16, 'Parent'),
  (4, 17, 'Parent'),
  (8, 15, 'Parent'),
  (8, 16, 'Parent'),
  (8, 17, 'Parent'),
  (5, 9, 'Spouse'),
  (5, 18, 'Parent'),
  (5, 19, 'Parent'),
  (5, 20, 'Parent'),
  (5, 21, 'Parent'),
  (9, 18, 'Parent'),
  (9, 19, 'Parent'),
  (9, 20, 'Parent'),
  (9, 21, 'Parent'),
  (6, 10, 'Spouse'),
  (6, 22, 'Parent'),
  (6, 23, 'Parent'),
  (10, 22, 'Parent'),
  (10, 23, 'Parent'),
  (11, 24, 'Spouse'),
  (12, 25, 'Spouse'),
  (12, 28, 'Parent'),
  (12, 29, 'Parent'),
  (12, 30, 'Parent'),
  (12, 31, 'Parent'),
  (12, 32, 'Parent'),
  (12, 33, 'Parent'),
  (25, 28, 'Parent'),
  (25, 29, 'Parent'),
  (25, 30, 'Parent'),
  (25, 31, 'Parent'),
  (25, 32, 'Parent'),
  (25, 33, 'Parent'),
  (15, 26, 'Spouse'),
  (15, 34, 'Parent'),
  (15, 35, 'Parent'),
  (15, 36, 'Parent'),
  (15, 37, 'Parent'),
  (26, 34, 'Parent'),
  (26, 35, 'Parent'),
  (26, 36, 'Parent'),
  (26, 37, 'Parent'),
  (16, 27, 'Spouse'),
  (18, 38, 'Parent');