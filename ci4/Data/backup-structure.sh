#!/bin/sh

pg_dump -s -O -x -U $1 -h $2 $3 > structure.sql
