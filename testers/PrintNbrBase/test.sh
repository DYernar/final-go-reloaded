cd ./main
go build
echo "Arguments: nbr = 919617 and base = 01"
echo "Expected result:"
echo "11100000100001000001"
echo "Your result:    "; ./main 919617 "01"
echo
echo "Arguments: nbr = 753639 and base = CHOUMIisDAcat!"
echo "Expected result:"
echo "HIDAHI"
echo "Your result:    "; ./main 753639 "CHOUMIisDAcat!"
echo
echo "Arguments: nbr = -174336 and base = CHOUMIisDAcat!"
echo "Expected result:"
echo "-MssiD"
echo "Your result:    "; ./main -174336 "CHOUMIisDAcat!"
echo
echo "Arguments: nbr = -661165 and base = 1"
echo "Expected result:"
echo "NV"
echo "Your result:    "; ./main -661165 "1"
echo
echo "Arguments: nbr = -861737 and base = Zone01Zone01"
echo "Expected result:"
echo "NV"
echo "Your result:    "; ./main -861737 "Zone01Zone01"
echo
echo "Arguments: nbr = 125 and base = 0123456789ABCDEF"
echo "Expected result:"
echo "7D"
echo "Your result:    "; ./main 125 "0123456789ABCDEF"
echo
echo "Arguments: nbr = -125 and base = choumi"
echo "Expected result:"
echo "-uoi"
echo "Your result:    "; ./main -125 "choumi"
echo
echo "Arguments: nbr = 125 and base = -ab"
echo "Expected result:"
echo "NV"
echo "Your result:    "; ./main 125 "-ab"
echo
echo "Arguments: nbr = -9223372036854775808 and base = 0123456789"
echo "Expected result:"
echo "-9223372036854775808"
echo "Your result:    "; ./main -9223372036854775808 "0123456789"
echo
rm main