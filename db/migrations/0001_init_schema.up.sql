DROP TABLE IF EXISTS client CASCADE;

DROP SEQUENCE IF EXISTS client_id_seq;

CREATE SEQUENCE client_id_seq START 1000;

create table client (
  Id bigint NOT NULL PRIMARY KEY,
  Name varchar(40) NOT NULL,
  LogoUrl text
);
