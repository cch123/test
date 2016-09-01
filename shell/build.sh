#!/bin/sh
basepath=$(cd `dirname $0`; pwd)
#echo $basepath
sed -i'.bak' -e "s/define(['\"]ENVIRONMENT[\"'], [\"']development[\"'])/define(\"ENVIRONMENT\", \"production\")/g" $basepath/protected/yiic.php
sed -i'.bak' -e "s/define(['\"]ENVIRONMENT[\"'], [\"']development[\"'])/define(\"ENVIRONMENT\", \"production\")/g" $basepath/index.php
#sed -i'.bak' -e 's/define("ENVIRONMENT", "development")/define("ENVIRONMENT", "production")/g' $basepath/index.php
mkdir output
rm -rf .git/
ls|grep -v 'output\|build.sh'|xargs -i mv {} output/
