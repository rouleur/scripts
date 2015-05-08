#!/bin/bash

# 
# Creates a csv file from a NAVTEQ POI layer. 
# Fields for csv are X,Y,POI_NAME,and ST_NAME (street)
# Usage: navpoi2csv <POI_TABLE_NAME> e.g navpoi2csv NAVTEQ_BANKS
#

layer_name=$1;
db_user=GISUSER;
db_pwd=GISUSER;
tns=TELE;
layer_csv=$layer_name.csv;

echo $db_user;

sqlplus $db_user/$db_pwd@$tns << EOF
   set embedded on
   set pagesize 0
   set colsep '|'
   set echo off
   set feedback off
   set linesize 1000
   set trimspool on
   set headsep off
   set heading off;
   spool $layer_csv;
  # select sde.st_x(SHAPE) || ',' || sde.st_y(SHAPE) ||',"'|| POI_NAME ||'","' || ST_NAME ||'"' from $layer_name;
    select sde.st_x($layer_name.SHAPE) || ',' || sde.st_y($layer_name.SHAPE) ||',"'||
 $layer_name.POI_NAME ||'","' || NAVTEQ_STREETS.ST_NAME ||'","' || NAVTEQ_MTDAREA.AREA_NAME ||'"'
 from $layer_name 
 join NAVTEQ_STREETS on $layer_name.LINK_ID=NAVTEQ_STREETS.LINK_ID 
  join NAVTEQ_MTDAREA on NAVTEQ_STREETS.R_AREA_ID = NAVTEQ_MTDAREA.AREA_ID;
EOF
sed -i -e /^SQL\>.*/d $layer_csv

