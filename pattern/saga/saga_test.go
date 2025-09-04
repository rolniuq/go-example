package saga

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func createOrder() error {
	fmt.Println("create order")
	return nil
}

func cancelOrder() error {
	fmt.Println("cancel order")
	return nil
}

func createCustomer() error {
	fmt.Println("create customer")
	return fmt.Errorf("create customer error")
}

func cancelCustomer() error {
	fmt.Println("cancel customer")
	return nil
}

func TestSaga_Run(t *testing.T) {
	err := NewSaga().
		AddWorker(Worker{
			Action:       createOrder,
			Compensation: cancelOrder,
		}).
		AddWorker(Worker{
			Action:       createCustomer,
			Compensation: cancelCustomer,
		}).Run()

	require.True(t, strings.Contains(err.Error(), "Saga rolling back"))
}
