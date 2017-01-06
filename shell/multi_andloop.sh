#! /bin/sh
function help() {
    echo "sh multi_grep.sh \$grepStr \$logfile \$machineType"
    exit 1
}
[ "$1" == "" ] && help
[ "$2" == "" ] && help

grepStr=$1
logFile=$2

[ -z "$3" ] && machineGroup="all" || machineGroup=$3

for machine in `cat machines.${machineGroup}`
do
ssh $machine "grep \"${grepStr}\" ${logFile}"
done
#! /bin/sh
function help() {
    echo "sh multi_grep.sh \$grepStr \$logfile \$machineType"
    exit 1
}
[ "$1" == "" ] && help
[ "$2" == "" ] && help

grepStr=$1
logFile=$2

[ -z "$3" ] && machineGroup="all" || machineGroup=$3

for machine in `cat machines.${machineGroup}`
do
ssh $machine "tail -f ${logFile} | grep \"${grepStr}\""
done

