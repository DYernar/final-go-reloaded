cd ./main
go build
echo "Arguments: nb = -7 and power = -2"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./main -7 -2
echo
echo "Arguments: nb = -8 and power = -7"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./main -8 -7
echo
echo "Arguments: nb = 4 and power = 8"
echo "Expected result:"
echo "65536"
echo "Your result:    "; ./main 4 8
echo
echo "Arguments: nb = 1 and power = 3"
echo "Expected result:"
echo "1"
echo "Your result:    "; ./main 1 3
echo
echo "Arguments: nb = -1 and power = 1"
echo "Expected result:"
echo "-1"
echo "Your result:    "; ./main -1 1
echo
echo "Arguments: nb = -6 and power = 5"
echo "Expected result:"
echo "-7776"
echo "Your result:    "; ./main -6 5
echo
rm main