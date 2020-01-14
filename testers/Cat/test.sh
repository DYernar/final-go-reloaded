go build
echo "file: quest8.txt"
echo "Expected result:"
cat quest8.txt ; echo
echo "Your result:"
./Cat quest8.txt ; echo
echo
echo "file: quest8T.txt"
echo "Expected result:"
cat quest8T.txt ; echo
echo "Your result:"
./Cat quest8T.txt ; echo
echo
echo "file: quest8.txt quest8T.txt"
echo "Expected result:"
cat quest8.txt quest8T.txt ; echo
echo "Your result:"
./Cat quest8.txt quest8T.txt ; echo
echo
echo "file: abc"
echo "Expected result:"
cat abc 
echo "Your result:"
./Cat abc ; echo
echo
echo "Running without arguments:"
./Cat