# multiple-knapsack-problem
An implementation of the 0/1 Multiple Knapsack problem.

The 0-1 Multiple Knapsack Problem (MKP) is: give a set of n items and a set of m knapsacks (m <= n)

### build for windowns xp 32bit

```
env GOOS=windows GOARCH=386 go build
```

### reference

```
./doc/Chapter6.pdf
```

### how to use

1. change "input" file
    ```
    ; lines starting with ; or # are comments
    ; knapsack list
    ; num weight/size
    2 21000
    ; following lines specify list of objects with properties
    ; name   price/value  weight/size
    ZM170740765 10180 10180
    ZM170840636 6435 6435
    ZM170840637 6340 6340
    ZM170840804 10100 10100
    ZM170840805 10005 10005
    ZM170840806 10055 10055
    ZM170840807 10940 10940
    ZM161240072 1600 1600
    ZM170240664 800 800
    ZM170840710 9888 9888
    ```
1. go run main.go
1. check "output" file
    ```
    capacity: 42000, max: 41.900, num: 5, time: 0 seconds
    knap: name(knap-1-21000), capacity(21.000), used(20.995), num(2)
        item: name(ZM170840807), weight(10.940), value(10.940)
        item: name(ZM170840806), weight(10.055), value(10.055)
    knap: name(knap-2-21000), capacity(21.000), used(20.905), num(3)
        item: name(ZM170840804), weight(10.100), value(10.100)
        item: name(ZM170840805), weight(10.005), value(10.005)
        item: name(ZM170240664), weight(0.800), value(0.800)
    ```

