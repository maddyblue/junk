#!/bin/sh

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
