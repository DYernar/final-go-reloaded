cd ./main
go build
echo "Argument: \"\""
echo "Expected result:"
echo "0"
echo "Your result:    "; ./main ""
echo
echo "Argument: "-""
echo "Expected result:"
echo "0"
echo "Your result:    "; ./main "-"
echo
echo "Argument: "--123""
echo "Expected result:"
echo "0"
echo "Your result:    "; ./main "--123"
echo
echo "Argument: "1""
echo "Expected result:"
echo "1"
echo "Your result:    "; ./main "1"
echo
echo "Argument: "-3""
echo "Expected result:"
echo "-3"
echo "Your result:    "; ./main "-3"
echo
echo "Argument: "8292""
echo "Expected result:"
echo "8292"
echo "Your result:    "; ./main "8292"
echo
echo "Argument: "9223372036854775807""
echo "Expected result:"
echo "9223372036854775807"
echo "Your result:    "; ./main "9223372036854775807"
echo
echo "Argument: "-9223372036854775808""
echo "Expected result:"
echo "-9223372036854775808"
echo "Your result:    "; ./main "-9223372036854775808"
echo
rm main