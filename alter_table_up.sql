BEGIN;
update feature_flags set enabled = true where name = 'users_name_to_display_name';

ALTER TABLE users RENAME COLUMN name TO display_name;

COMMIT;