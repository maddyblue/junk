#!/bin/sh

# dump data of static tables

mysqldump -u user --password=usersql -h faye -t ci4 \
	ability \
	abilitytype \
	cor_job_ability \
	cor_job_equipmenttype \
	cor_job_joblv \
	cor_monster_drop \
	domain \
	equipment \
	equipmenttype \
	item \
	job \
	monster \
	monstertype \
	site \
	> data.sql
