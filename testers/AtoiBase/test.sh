cd ./main
go build
echo "Arguments: s = bcbbbbaab and base = abc"
echo "Expected result:"
echo "12016"
echo "Your result:    "; ./main bcbbbbaab abc
echo
echo "Arguments: s = 00 and base = 01"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./main 00 01
echo
echo "Arguments: s = saDt!I!sI and base = CHOUMIisDAcat!"
echo "Expected result:"
echo "11557277891"
echo "Your result:    "; ./main "saDt!I!sI" "CHOUMIisDAcat!"
echo
echo "Arguments: s = AAho?Ao and base = WhoAmI?"
echo "Expected result:"
echo "406772"
echo "Your result:    "; ./main "AAho?Ao" "WhoAmI?"
echo
echo "Arguments: s = thisinputshouldnotmatter and base = abca"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./main "thisinputshouldnotmatter" "abca"
echo
echo "Arguments: s = 125 and base = 0123456789"
echo "Expected result:"
echo "125"
echo "Your result:    "; ./main "125" "0123456789"
echo
echo "Arguments: s = bbbbbab and base = -ab"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./main "bbbbbab" "-ab"
echo
rm main