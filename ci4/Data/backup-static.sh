#!/bin/sh

# List here only the tables that contain static data.
# Do not list things like player and chocobo tables - those are dynamic and
# should only the structure should be stored, not the data.

mysqldump -u user --password=usersql -h faye ci4 \
	ability \
	abilitytype \
	cor_job_ability \
	cor_job_itemtype \
	cor_job_joblv \
	cor_monster_item \
	item \
	itemtype \
	job \
	monster \
	monstertype \
	site \
	> backup-static.sql
