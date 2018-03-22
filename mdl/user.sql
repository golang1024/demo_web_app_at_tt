DROP TABLE IF EXISTS tbl_users;
CREATE TABLE tbl_users (
  aid SERIAL,
  user_name text,
  PRIMARY KEY (aid)
);


DROP TABLE IF EXISTS tbl_relations;
CREATE TABLE tbl_relations (
  uida int,
  uidb int,
  statea int,
  stateb int,
  PRIMARY KEY (uida, uidb)
);