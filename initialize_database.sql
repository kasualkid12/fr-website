CREATE TABLE
  persons (
    person_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    birth_date DATE,
    death_date DATE,
    gender VARCHAR(50),
    photo_url VARCHAR(255),
    profile_id INTEGER
  );

CREATE TABLE
  relationships (
    relationship_id SERIAL PRIMARY KEY,
    person1_id INTEGER,
    person2_id INTEGER,
    relationship_type VARCHAR(50)
  );

CREATE TABLE
  profiles (
    profile_id SERIAL PRIMARY KEY,
    person_id INTEGER UNIQUE,
    biography TEXT,
    photo_url VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
  );

ALTER TABLE profiles ADD CONSTRAINT fk_profile_person FOREIGN KEY (person_id) REFERENCES persons (person_id) DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE persons ADD CONSTRAINT fk_person_profile FOREIGN KEY (profile_id) REFERENCES profiles (profile_id);

ALTER TABLE relationships ADD CONSTRAINT fk_relationship_person1 FOREIGN KEY (person1_id) REFERENCES persons (person_id);

ALTER TABLE relationships ADD CONSTRAINT fk_relationship_person2 FOREIGN KEY (person2_id) REFERENCES persons (person_id);