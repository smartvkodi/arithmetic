// Exercise 1. Create a package called arithmetic
// in $GOPATH/packages/mypackages
// and paste the following code in the file:
package arithmetic

// Checks if a number is prime or not
func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func IsEven(n int) bool {
	return n%2 == 0
}

func IsOdd(n int) bool {
	return n%2 != 0
}
