
if [ $# == 0 ]
then 
    ./run-sniper -c ./config-raw.cfg -- ./mytest/lab0.exe
elif [ $# == 1 ]
then 
    ./run-sniper -c ./config-"$1".cfg -- ./mytest/lab0.exe 
elif [ $# == 2 ]
then
    ./run-sniper -c ./config-"$1".cfg -- ./mytest/"$2"
fi 
    