#!/bin/bash
# 
# Creates a Jboss Module XML from the jars in the current directory.
#

if [ $# -eq 0 ]
  then
    echo "Usage: modulizer <module name>"
    exit 1
fi
echo "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
echo "<module xmlns=\"urn:jboss:module:1.1\" name=\"$1\">"
echo "  <resources>"
for f in *.jar; do echo "    <resource-root path=\""$f"\"/>"; done
echo "  </resources>"
echo "  <dependencies>"
echo "    <module name=\"sun.jdk\" export=\"true\" >"
echo "      <imports>"
echo "        <include path=\"sun/misc/Unsafe\" />"
echo "      </imports>"
echo "    </module>"
echo "    <module name=\"org.apache.log4j\"/>"
echo "    <module name=\"org.apache.commons.logging\"/>"
echo "    <module name=\"javax.api\"/>"
echo "  </dependencies>"
echo "</module>"
