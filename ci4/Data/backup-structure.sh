#!/bin/sh

pg_dump -s -O -x -U dolmant ci4 > structure.sql
