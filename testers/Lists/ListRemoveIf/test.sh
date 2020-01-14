cd ./main
go build
echo "Try with the arguments: l = <nil> and data_ref = 1

(<nil>)"
./main 1
echo
echo
echo
echo "Try with the arguments: l = <nil> and data_ref = 96

(<nil>)"
./main 2
echo
echo
echo
echo "Try with the arguments: l = 98-> 98-> 33-> 34-> 33-> 34-> 33-> 89-> 33-> <nil> and data_ref = 34

(98-> 98-> 33-> 33-> 33-> 89-> 33-> <nil>)"
./main 3
echo
echo
echo
echo "Try with the arguments: l = 79-> 74-> 99-> 79-> 7-> <nil> and data_ref = 99

(79-> 74-> 79-> 7-> <nil>)"
./main 4
echo
echo
echo
echo "Try with the arguments: l = 56-> 93-> 68-> 56-> 87-> 68-> 56-> 68-> <nil> and data_ref = 68

(56-> 93-> 56-> 87-> 56-> <nil>)"
./main 5
echo
echo
echo
echo "Try with the arguments: l = mvkUxbqhQve4l-> 4Zc4t hnf SQ-> q2If E8BPuX -> <nil> and data_ref = 4Zc4t hnf SQ

(mvkUxbqhQve4l-> q2If E8BPuX -> <nil>)"
./main 6
echo
echo
echo