g++ ./main.cpp -o main 

cat ./test/$1.in | ./main > a.out
cat ./test/$1.out > b.out
diff -Z a.out b.out
