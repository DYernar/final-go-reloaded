go build
echo "Arguments: 861 + 870"
echo "Expected result:"
echo "1731"
echo "Your result:    "; ./Doop 861 + 870
echo
echo "Arguments: 861 - 870"
echo "Expected result:"
echo "-9"
echo "Your result:    "; ./Doop 861 - 870
echo
echo "Arguments: 861 * 870"
echo "Expected result:"
echo "749070"
echo "Your result:    "; ./Doop 861 "*" 870
echo
echo "Arguments: 861 % 870"
echo "Expected result:"
echo "861"
echo "Your result:    "; ./Doop 861 % 870
echo
echo "Arguments: hello + 1"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./Doop hello + 1
echo
echo "Arguments: 1 p 1"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./Doop 1 p 1
echo
echo "Arguments: 1 / 0"
echo "Expected result:"
echo "No division by 0"
echo "Your result:    "; ./Doop 1 / 0
echo
echo "Arguments: 1 % 0"
echo "Expected result:"
echo "No remainder of division by 0"
echo "Your result:    "; ./Doop 1 % 0
echo
echo "Arguments: 1 * 1"
echo "Expected result:"
echo "1"
echo "Your result:    "; ./Doop 1 "*" 1
echo
echo "Arguments: 9223372036854775807 + 1"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./Doop 9223372036854775807 + 1
echo
echo "Arguments: 9223372036854775809 - 3"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./Doop 9223372036854775809 - 3
echo
echo "Arguments: 9223372036854775807 * 3"
echo "Expected result:"
echo "0"
echo "Your result:    "; ./Doop 9223372036854775807 "*" 3
echo
rm Doop