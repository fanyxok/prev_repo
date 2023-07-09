echo $1  >> result_record
echo "one-bit"  >> result_record
./run-sniper -c ./config-one-bit.cfg -- ./mytest/"$1"
sleep 1s
sed -n 11,11p sim.out >> result_record

echo "pentium_m"  >> result_record
./run-sniper -c ./config-pentium_m.cfg -- ./mytest/"$1"
sed -n 11,11p sim.out >> result_record

echo "perceptron"  >> result_record
./run-sniper -c ./config-perceptron.cfg -- ./mytest/"$1"
sed -n 11,11p sim.out >> result_record

echo "perceptron-local"  >> result_record
./run-sniper -c ./config-perceptron-local.cfg -- ./mytest/"$1"
sed -n 11,11p sim.out >> result_record
echo "------" >> result_record