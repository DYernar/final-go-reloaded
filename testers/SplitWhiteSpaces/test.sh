cd ./main
go build
echo "Expected result:"
echo "[The earliest foundations of what would become computer science predate the invention of the modern digital computer]"
echo "Your result:    "; ./main  "The earliest foundations of what would become computer science predate the invention of the modern digital computer"
echo
echo "Expected result:"
echo "[Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,]"
echo "Your result:    "; ./main  "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
echo
echo "Expected result:"
echo "[aiding in computations such as multiplication and division .]"
echo "Your result:    "; ./main  "aiding in computations such as multiplication and division ."
echo
echo "Expected result:"
echo "[Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment.]"
echo "Your result:    "; ./main  "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
echo
echo "Expected result:"
echo "[Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]]"
echo "Your result:    "; ./main  "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
echo
echo "Expected result:"
echo "[In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,]"
echo "Your result:    "; ./main  " In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
echo
rm main