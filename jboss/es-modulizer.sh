#!/bin/bash

#
# Elasticsearch JBoss Module Generator
# Prerequisite: maven (version > 3.0)
# How it works: Creates a dummy maven project having elasticsearch as dependency
# Using mvn dependency plugin gets the jars and generates a module.xml
#

pomxml="<project xmlns=\"http://maven.apache.org/POM/4.0.0\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd\"><modelVersion>4.0.0</modelVersion><groupId>es</groupId><artifactId>es</artifactId><version>0.0.1-SNAPSHOT</version><dependencies><dependency><groupId>org.elasticsearch</groupId><artifactId>elasticsearch</artifactId><version>2.2.0</version><scope>provided</scope></dependency></dependencies></project>"
workdir=$PWD
projdir=es-$RANDOM
if [ -d $projdir ]; then
   projdir=es-$RANDOM
fi
echo "Creating temporary mvn project: $projdir"
mkdir $projdir
echo $pomxml > $projdir/pom.xml
cd $projdir
outputdir=es-jboss-module
moduledir=$workdir/$outputdir/org/elasticsearch/main
mvn dependency:copy-dependencies -DoutputDirectory="$moduledir"
cd $moduledir
echo "Generating module.xml..."
cat << HEAD >> module.xml
<module xmlns="urn:jboss:module:1.1" name="org.elasticsearch">
  <resources>
HEAD
for f in *.jar
  do  echo "   <resource-root path=\""$f"\"/>" >> module.xml
done
cat << TAIL >> module.xml
  </resources>
  <dependencies>
    <module name="sun.jdk" export="true" >
      <imports>
        <include path="sun/misc/Unsafe" />
      </imports>
    </module>
    <module name="org.apache.log4j"/>
    <module name="org.apache.commons.logging"/>
    <module name="javax.api"/>
  </dependencies>
</module>
TAIL
echo "Cleaning temporary mvn project $projdir..."
rm -rf $workdir/$projdir
echo "Elasticsearch module ready under $outputdir"





