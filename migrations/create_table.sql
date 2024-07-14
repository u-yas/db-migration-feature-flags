CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INT NOT NULL
);

CREATE TABLE feature_flags (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  enabled BOOLEAN NOT NULL
);

insert INTO users (name, age) VALUES
('Alice', 20),
('Bob', 25),
('Charlie', 30),
('David', 35),
('Eve', 40),
('Frank', 45),
('Grace', 50),
('Heidi', 55),
('Ivan', 60),
('Judy', 65)
;

insert INTO feature_flags (name, enabled) VALUES
('users_name_to_display_name', false);

COMMIT;