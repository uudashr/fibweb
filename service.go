package fibweb

// FibonacciService provide all fibonacci operation
type FibonacciService interface {
	Seq(int) ([]int, error)
}
