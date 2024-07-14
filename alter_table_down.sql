BEGIN;
update feature_flags set enabled = false where name = 'users_name_to_display_name';

ALTER TABLE users RENAME COLUMN display_name TO name;

COMMIT;