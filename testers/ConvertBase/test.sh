cd ./main
go build
echo "Arguments: nbr = "4506C", baseFrom = "0123456789ABCDEF" and baseTo = "choumi""
echo "Expected result:"
echo "hccocimc"
echo "Your result:    "; ./main 4506C 0123456789ABCDEF choumi
echo
echo "Arguments: nbr = "babcbaaaaabcb", baseFrom = "abc" and baseTo = "0123456789ABCDEF""
echo "Expected result:"
echo "9B611"
echo "Your result:    "; ./main babcbaaaaabcb abc 0123456789ABCDEF
echo
echo "Arguments: nbr = "256850", baseFrom = "0123456789" and baseTo = "01""
echo "Expected result:"
echo "111110101101010010"
echo "Your result:    "; ./main 256850 0123456789 01
echo
echo "Arguments: nbr = "uuhoumo", baseFrom = "choumi" and baseTo = "Zone01""
echo "Expected result:"
echo "eeone0n"
echo "Your result:    "; ./main uuhoumo choumi Zone01
echo
echo "Arguments: nbr = "683241", baseFrom = "0123456789" and baseTo = "0123456789""
echo "Expected result:"
echo "683241"
echo "Your result:    "; ./main 683241 0123456789 0123456789
echo
rm main