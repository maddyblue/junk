#!/bin/sh

o=data.sql
rm -f $o

for i in ability abilitytype area cor_area_monster cor_area_town cor_job_abilitytype cor_job_equipmenttype cor_job_joblv cor_monster_drop domain equipment equipmentclass equipmenttype event group_def house item job monster monstertype site skin town
do
	echo "truncate table ${i};" >> $o
	pg_dump -a -O -x -U dolmant -t $i ci4 >> $o
	echo "select setval('${i}_${i}_id_seq', max(${i}_id)) from $i;" >> $o
done
