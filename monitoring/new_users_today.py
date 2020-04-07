#!/usr/bin/python3

import os
import mysql.connector as mariadb

mariadb_connection = mariadb.connect(user=os.environ['DBUSER'], password=os.environ['DBPASSWORD'], database=os.environ['DBNAMESPACE'])

cursor = mariadb_connection.cursor()
cursor.execute('SELECT * FROM LOGS WHERE occurred_on > UNIX_TIMESTAMP(CURRENT_DATE() -1) AND occurrence_type = \'USERCREATED\';')

for row in cursor:
    print(row[2]," ",row[4])


