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
    'John Doe',
    '1980-01-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Jane Smith',
    '1985-02-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Alice Johnson',
    '1990-03-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Bob Brown',
    '1995-04-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Charlie Davis',
    '2000-05-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Diana Adams',
    '1975-06-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Evan Firth',
    '1987-07-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Grace Halt',
    '1992-08-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Ian Grubb',
    '1998-09-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Judy Klein',
    '2001-10-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 11',
    '1970-01-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 12',
    '1980-02-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 13',
    '1990-03-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 14',
    '2000-04-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 15',
    '2010-05-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 16',
    '1985-06-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 17',
    '1995-07-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 18',
    '2005-08-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 19',
    '1975-09-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 20',
    '1985-10-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 21',
    '1995-11-01',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 22',
    '2005-12-01',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 23',
    '1985-01-15',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 24',
    '1995-02-15',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 25',
    '2005-03-15',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 26',
    '1975-04-15',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 27',
    '1985-05-15',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 28',
    '1995-06-15',
    NULL,
    'Female',
    NULL,
    NULL
  ),
  (
    'Person 29',
    '2005-07-15',
    NULL,
    'Male',
    NULL,
    NULL
  ),
  (
    'Person 30',
    '1975-08-15',
    NULL,
    'Female',
    NULL,
    NULL
  );

-- Insert relationships among persons
INSERT INTO
  relationships (person1_id, person2_id, relationship_type)
VALUES
  (1, 2, 'Spouse'),
  (1, 3, 'Parent'),
  (1, 5, 'Parent'),
  (1, 7, 'Parent'),
  (2, 3, 'Parent'),
  (2, 5, 'Parent'),
  (2, 7, 'Parent'),
  (3, 4, 'Spouse'),
  (3, 5, 'Sibling'),
  (5, 7, 'Sibling'),
  (5, 6, 'Spouse'),
  (7, 8, 'Spouse');