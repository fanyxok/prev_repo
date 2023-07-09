

echo "run test for $1"
../../tools-bin/lexer $1 > ref.out
../../src/lexer $1 > out.out
diff   ref.out out.out
