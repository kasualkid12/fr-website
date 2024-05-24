-- Insert entries into the 'persons' table
INSERT INTO
  persons (
    name,
    birth_date,
    death_date,
    gender,
    profile_id,
    photo_url
  )
VALUES
  (
    'Ric Boyd', -- #1
    '1946-10-06',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Dana Boyd', -- #2
    '1949-07-16',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Dava Walker', -- #3
    '1968-09-10',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Jami Garzella', -- #4
    '1970-02-24',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Richard Boyd', -- #5
    '1971-09-08',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Matthew Boyd', -- #6
    '1977-06-05',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Paul Walker II', -- #7
    '1965-11-29',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Tom Garzella', -- #8
    '1968-11-18',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Christy Boyd', -- #9
    '1974-05-16',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Kelly Boyd', -- #10
    '1977-05-24',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Haleigh Forrest', -- #11
    '1990-04-20',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Hana Guess', -- #12
    '1991-06-23',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Jon Walker', -- #13
    '1994-06-29',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Trey Walker', -- #14
    '1999-06-21',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Calli Pawlicki', -- #15
    '1996-04-18',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Mikinna White', -- #16
    '1999-05-04',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Zackary Garzella', -- #17
    '2004-03-24',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Collin Boyd', -- #18
    '2000-09-12',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Kamrin Boyd', -- #19
    '2006-02-20',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Ozzie Boyd', -- #20
    '2010-02-16',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Easton Boyd', -- #21
    '2013-06-25',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Alexa Boyd', -- #22
    '2003-06-07',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Olivia Boyd', -- #23
    '2006-05-24',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'John Forrest', -- #24
    '1987-07-23',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Ian Guess', -- #25
    '1988-03-28',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Jacob Pawlicki', -- #26
    '1996-02-05',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Stephen White', -- #27
    '1998-03-15',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Desmond Guess', -- #28
    '2013-09-29',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Nona Guess', -- #29
    '2014-10-07',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Everett Guess', -- #30
    '2016-10-26',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Remington Guess', -- #31
    '2018-12-28',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Sage Guess', -- #32
    '2020-11-15',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Lewis Guess', -- #33
    '2022-08-23',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Madison Pawlicki', -- #34
    '2017-08-25',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Paisley Pawlicki', -- #35
    '2020-07-30',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Levi Pawlicki', -- #36
    '2021-09-30',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Kinsley Pawlicki', -- #37
    '2024-05-21',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'James Kruger', -- #38
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