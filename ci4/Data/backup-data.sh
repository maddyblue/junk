#!/bin/sh

group1='abilitytype area domain equipmentclass equipmenttype event group_def house item job monstertype site skin town'
group2='ability equipment monster'
group3='cor_area_monster cor_area_town cor_job_abilitytype cor_job_equipmenttype cor_job_joblv cor_monster_drop'

all="${group1} ${group2} ${group3}"
rev="${group3} ${group2} ${group1}"

o=data.sql
t=temp.sql
om=data-mysql.sql
os=data-sqlite.sql
rm -f $o $om

for i in $rev
do
	echo "delete from ${i};" >> $o
	echo "truncate ${i};" >> $om
done

for i in $all
do
	pg_dump -a -O -x -U $1 -h $2 -t $i $3 | \
		grep -v "^-" | grep -v "^$" > $t
	head -n 5 $t > $t.top
	sed 1,5d $t | sed '$d' | sort -n > $t.mid
	tail -n 1 $t > $t.bot
	cat $t.top $t.mid $t.bot >> $o
	rm $t $t.top $t.mid $t.bot

	pg_dump -a -O -x -d -U $1 -h $2 -t $i $3 | \
	sed 's/^SET/--SET/' | \
	sed "s/SELECT pg_catalog.setval('\(.*\)', \(.*\), true);/ALTER TABLE \1 AUTO_INCREMENT =\2;/" | \
	sed 's/"domain"/domain/' | \
	sed "s/', E'/', '/" \
		>> $om
done

sed 's/^.*AUTO_INCREMENT/--/' $om |\
sed 's/^truncate/delete from/' \
 > $os