#!/bin/sh

group1='abilitytype area domain equipmentclass equipmenttype event group_def house item job monstertype site skin town'
group2='ability equipment monster'
group3='cor_area_monster cor_area_town cor_job_abilitytype cor_job_equipmenttype cor_job_joblv cor_monster_drop'

all="${group1} ${group2} ${group3}"
rev="${group3} ${group2} ${group1}"

o=data.sql
t=temp.sql
rm -f $o

for i in $rev
do
	echo "delete from ${i};" >> $o
done

for i in $all
do
	pg_dump -a -O -x -U dolmant -t $i ci4 | \
		grep -v "^-" | grep -v "^$" > $t
	grep -B 10 "COPY" $t > $t.top
	grep "^[0-9].*$" $t | sort -n > $t.mid
	grep -A 10 "\\\." $t > $t.bot
	cat $t.top $t.mid $t.bot >> $o
	rm $t $t.top $t.mid $t.bot
done
