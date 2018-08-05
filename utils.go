package tricks

import "testing"

func shouldPanic(t *testing.T, f func(), msg string) {
	defer func() {
		rec := recover()
		if rec == nil {
			t.Error("func did not panic")
		}

		err, ok := rec.(error)
		if !ok || err.Error() != msg {
			t.Errorf("unexpected error ok = %t, err = %v", ok, err)
		}
	}()

	f()
}