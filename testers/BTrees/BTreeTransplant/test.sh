cd ./main
go build
echo "Arguments:"
echo "root =
01
└── 07
    ├── 12
    │   └── 10
    └── 05
        └── 02
            └── 03
, node = 12 and
repla =
55
├── 60
└── 33
    └── 12
        └── 15

01
└── 07
    ├── 55
    │   ├── 60
    │   └── 33
    │       └── 12
    │           └── 15
    └── 05
        └── 02
            └── 03"
echo
echo "Result:"
./main 1
echo
echo
echo
echo
echo
echo
echo "Arguments:"
echo "root =
33
├── 5
│   └── 52
└── 20
    ├── 31
    └── 13
        └── 11
, node = 20 and
repla =
55
├── 60
└── 33
    └── 12
        └── 15

33
├── 5
│   └── 52
└── 55
    ├── 60
    └── 33
        └── 12
            └── 15"
echo
echo "Result:"
./main 2
echo
echo
echo
echo
echo
echo
echo "Arguments:"
echo "root =
03
└── 39
    ├── 99
    │   └── 44
    └── 11
        └── 14
            └── 11
, node = 11 and
repla =
55
├── 60
└── 33
    └── 12
        └── 15

03
└── 39
    ├── 99
    │   └── 44
    └── 55
        ├── 60
        └── 33
            └── 12
                └── 15"
echo
echo "Result:"
./main 3
rm main