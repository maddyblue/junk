#!/bin/sh

pg_dump -s -O -x -U $1 -h $2 $3 > structure.sql

#sed 's/^SET/-- SET/' structure.sql | \
#sed 's/^COMMENT/-- COMMENT/' | \
#sed 's/bigserial NOT NULL/BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY/' | \
#sed 's/::bigint//' | \
#sed 's/DEFAULT 0/DEFAULT "0"/' | \
#sed 's/DEFAULT (0)/DEFAULT "0"/' | \
#sed 's/CREATE TABLE "domain"/CREATE TABLE domain/' | \
#sed 's/CREATE TABLE "session"/CREATE TABLE session/' | \
#sed 's/bytea/blob/' | \
#sed 's/^ALTER TABLE ONLY/-- ALTER TABLE ONLY/' | \
#sed 's/ADD CONSTRAINT .* PRIMARY KEY .*/ -- PRIMARY KEY/' | \
#sed 's/ADD CONSTRAINT .* UNIQUE .*/ -- UNIQUE/' | \
#sed 's/CREATE INDEX .* ON/ALTER TABLE/' | \
#sed 's/USING btree/add index/' | \
#sed 's/add index (forum_word_word)/add index (forum_word_word (50))/' | \
#sed 's/ADD CONSTRAINT .* FOREIGN KEY/-- FOREIGN KEY/' | \
#sed 's/boolean/tinyint/' |\
#sed 's/^CREATE SEQUENCE/-- CREATE SEQUENCE/' |\
#sed 's/^    START WITH/--/' |\
#sed 's/^    INCREMENT BY 1$/--/' |\
#sed 's/^    NO MAXVALUE$/--/' |\
#sed 's/^    NO MINVALUE$/--/' |\
#sed 's/^    CACHE 1;$/--/' |\
#sed 's/^ALTER SEQUENCE/-- ALTER SEQUENCE/' |\
#sed 's/^ALTER TABLE .* ALTER COLUMN .* SET DEFAULT nextval.*regclass/--/' \
# > structure-mysql.sql

#sed 's/^ALTER TABLE .* add index .*/--/' structure-mysql.sql |\
#sed 's/^CREATE UNIQUE INDEX.*/--/' |\
#sed 's/^ALTER TABLE .* USING hash .*/--/' \
# > structure-sqlite.sql