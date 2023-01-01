# Is prime

Write a function, is_prime, that takes in a number as an argument. The function should return a boolean indicating whether or not the given number is prime.

A prime number is a number that is only divisible by two distinct numbers: 1 and itself.

For example, 7 is a prime because it is only divisible by 1 and 7. For example, 6 is not a prime because it is divisible by 1, 2, 3, and 6.

You can assume that the input number is a positive integer.


test_00:
```js
is_prime(2) # -> True
```

test_01:
```js
is_prime(3) # -> True
```

test_02:
```js
is_prime(4) # -> False
```

test_03:
```js
is_prime(5) # -> True
```

test_04:
```js
is_prime(6) # -> False
```

test_05:
```js
is_prime(2017) # -> True
```

test_06:
```js
is_prime(713) # -> False
```