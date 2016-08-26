#!/bin/sh

tables='ability abilitytype area cor_area_monster cor_area_town cor_job_abilitytype cor_job_equipmenttype cor_job_joblv cor_monster_drop domain equipment equipmentclass equipmenttype event group_def house item job monster monstertype site skin town'

for i in $tables
do
	t=$t" -t "$i
done

pg_dump -a -Fc -O -U $1 -h $2$t $3 > data.dump
pg_dump -a -Fp -O -U $1 -h $2$t $3 > data.sql

