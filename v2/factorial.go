// Exercise 2. - Add a new file to the arithmetic package
// and paste the following code in the file
package arithmetic

func Factorial(n int) int {
	var f int = 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}
