package main

func main() {
	Go("test_command", func() error {
		return nil
	}, func(err error) error {
		return nil
	})
}