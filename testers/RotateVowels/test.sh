go build
echo "Arguments: """
echo "Expected result:"
echo 
echo "Your result:    "; ./RotateVowels ""
echo
echo "Arguments: "Hello World""
echo "Expected result:"
echo "Hollo Werld"
echo "Your result:    "; ./RotateVowels "Hello World"
echo
echo "Arguments: "HEllO World" "problem solved""
echo "Expected result:"
echo "Hello Werld problom sOlvEd"
echo "Your result:    "; ./RotateVowels "HEllO World" "problem solved"
echo
echo "Arguments: "str" "shh" "psst""
echo "Expected result:"
echo "str shh psst"
echo "Your result:    "; ./RotateVowels "str" "shh" "psst"
echo
echo "Arguments: "happy thoughts" "good luck""
echo "Expected result:"
echo "huppy thooghts guod lack"
echo "Your result:    "; ./RotateVowels "happy thoughts" "good luck"
echo
echo "Arguments:  al's elEphAnt is overly underweight!"
echo "Expected result:"
echo "il's elephunt es ovirly AndErweaght!"
echo "Your result:    "; ./RotateVowels "al's elEphAnt is overly underweight!"
echo
rm RotateVowels