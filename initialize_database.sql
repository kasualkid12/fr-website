-- Create tables
CREATE TABLE
  persons (
    person_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
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

-- Add foreign key constraints
ALTER TABLE profiles ADD CONSTRAINT fk_profile_person FOREIGN KEY (person_id) REFERENCES persons (person_id) DEFERRABLE INITIALLY DEFERRED;

ALTER TABLE persons ADD CONSTRAINT fk_person_profile FOREIGN KEY (profile_id) REFERENCES profiles (profile_id);

ALTER TABLE relationships ADD CONSTRAINT fk_relationship_person1 FOREIGN KEY (person1_id) REFERENCES persons (person_id);

ALTER TABLE relationships ADD CONSTRAINT fk_relationship_person2 FOREIGN KEY (person2_id) REFERENCES persons (person_id);

-- Create views
CREATE OR REPLACE VIEW public.tree_initial_vw AS
WITH RECURSIVE family_tree AS (
    SELECT p.person_id,
           p.first_name,
           p.last_name,
           p.birth_date,
           p.death_date,
           p.gender,
           p.photo_url,
           p.profile_id,
           'Self'::text AS relationship,
           p.person_id AS parent_object
    FROM persons p
    UNION ALL
    SELECT p.person_id,
           p.first_name,
           p.last_name,
           p.birth_date,
           p.death_date,
           p.gender,
           p.photo_url,
           p.profile_id,
           CASE
               WHEN r.relationship_type::text = 'Parent'::text THEN (
                   (SELECT p2.first_name || ' ' || p2.last_name
                    FROM persons p2
                    WHERE p2.person_id = r.person1_id)::text
               ) || ' Child'::text
               WHEN r.relationship_type::text = 'Spouse'::text THEN (
                   (SELECT p2.first_name || ' ' || p2.last_name
                    FROM persons p2
                    WHERE p2.person_id = r.person1_id)::text
               ) || ' Spouse'::text
               ELSE NULL::text
           END AS relationship,
           f.parent_object
    FROM relationships r
             JOIN family_tree f ON r.person1_id = f.person_id
             JOIN persons p ON p.person_id = r.person2_id
    WHERE f.relationship = 'Self'::text
      AND (r.relationship_type::text = ANY (ARRAY['Parent'::character varying, 'Spouse'::character varying]::text[]))
)
SELECT family_tree.person_id,
       family_tree.first_name,
       family_tree.last_name,
       family_tree.birth_date,
       family_tree.death_date,
       family_tree.gender,
       family_tree.photo_url,
       family_tree.profile_id,
       family_tree.relationship,
       family_tree.parent_object
FROM family_tree;

ALTER TABLE public.tree_initial_vw
    OWNER TO root;

CREATE OR REPLACE VIEW public.tree_child_spouse_vw AS
WITH RECURSIVE family_tree AS (
    SELECT tree_initial_vw.person_id,
           tree_initial_vw.first_name,
           tree_initial_vw.last_name,
           tree_initial_vw.birth_date,
           tree_initial_vw.death_date,
           tree_initial_vw.gender,
           tree_initial_vw.photo_url,
           tree_initial_vw.profile_id,
           tree_initial_vw.relationship,
           tree_initial_vw.parent_object
    FROM tree_initial_vw
    UNION ALL
    SELECT p.person_id,
           p.first_name,
           p.last_name,
           p.birth_date,
           p.death_date,
           p.gender,
           p.photo_url,
           p.profile_id,
           (
               (SELECT p2.first_name || ' ' || p2.last_name
                FROM persons p2
                WHERE p2.person_id = r.person1_id)::text
           ) || ' Spouse'::text AS relationship,
           f.parent_object
    FROM relationships r
             JOIN family_tree f ON r.person1_id = f.person_id
             JOIN persons p ON p.person_id = r.person2_id
    WHERE f.relationship ~~ '%Child'::text
      AND r.relationship_type::text ~~ '%Spouse'::text
)
SELECT family_tree.person_id,
       family_tree.first_name,
       family_tree.last_name,
       family_tree.birth_date,
       family_tree.death_date,
       family_tree.gender,
       family_tree.photo_url,
       family_tree.profile_id,
       family_tree.relationship,
       family_tree.parent_object
FROM family_tree;

ALTER TABLE public.tree_child_spouse_vw
    OWNER TO root;
