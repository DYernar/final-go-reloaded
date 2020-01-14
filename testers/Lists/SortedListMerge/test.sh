cd ./main
go build
echo "Try with the arguments: n1 = <nil> and n2 = <nil>

<nil>"
echo "Your result:"
./main 1
echo
echo
echo
echo "Try with the arguments: n1 = <nil> and n2 = 2-> 2-> 4-> 9-> 12-> 12-> 19-> 20-> <nil>

2 -> 2 -> 4 -> 9 -> 12 -> 12 -> 19 -> 20 -> <nil>"
./main 2
echo
echo
echo
echo "Try with the arguments: n1 = 4-> 4-> 6-> 9-> 13-> 18-> 20-> 20-> <nil> and n2 = <nil>

4 -> 4 -> 6 -> 9 -> 13 -> 18 -> 20 -> 20 -> <nil>"
./main 3
echo
echo
echo

echo "Try with the arguments: n1 = 0-> 7-> 39-> 92-> 97-> 93-> 91-> 28-> 64-> <nil> and n2 = 80-> 23-> 27-> 30-> 85-> 81-> 75-> 70-> <nil>

0 -> 7 -> 23 -> 27 -> 28 -> 30 -> 39 -> 64 -> 70 -> 75 -> 80 -> 81 -> 85 -> 91 -> 92 -> 93 -> 97 -> <nil>"
./main 4
echo
echo
echo
echo "Try with the arguments: n1 =0-> 2-> 11-> 30-> 54-> 56-> 70-> 79-> 99-> <nil> and n2 = 23-> 28-> 38-> 67-> 67-> 79-> 95-> 97-> <nil>

0 -> 2 -> 11 -> 23 -> 28 -> 30 -> 38 -> 54 -> 56 -> 67 -> 67 -> 70 -> 79 -> 79 -> 95 -> 97 -> 99 -> <nil>"
./main 5
echo
echo
echo
echo "Try with the arguments: n1 = 0-> 3-> 8-> 8-> 13-> 19-> 34-> 38-> 46-> <nil> and n2 = 7-> 39-> 45-> 53-> 59-> 70-> 76-> 79-> <nil>

0 -> 3 -> 7 -> 8 -> 8 -> 13 -> 19 -> 34 -> 38 -> 39 -> 45 -> 46 -> 53 -> 59 -> 70 -> 76 -> 79 -> <nil>"
./main 6
echo
echo
echo
rm main