#!/bin/sh

# dump structure only (-d)

mysqldump -u user --password=usersql -h faye -d ci4 > structure.sql
