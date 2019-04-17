//go:generate mockgen -source=interfaces.go -destination=client_mock.go -package=commands -self_package=mather/internal/commands MatherClient

package commands

// Defines the interface for performing math operations the commands require
type MatherClient interface {
	Add(a, b int) (int, error)
	Subtract(a, b int) (int, error)
	Multiply(a, b int) (int, error)
	Divide(a, b int) (float64, error)
}
