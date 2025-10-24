package emergency

import (
	"testing"
	"time"
)

func TestTimelock(t *testing.T) {
	unlockTime := time.Now().Add(1 * time.Second)
	timelock := NewTimelock(unlockTime)

	// Test Unlock before timelock has expired
	err := timelock.Unlock()
	if err == nil {
		t.Errorf("Unlock() expected error for timelock not expired, but got nil")
	}

	// Wait for timelock to expire
	time.Sleep(1 * time.Second)

	// Test Unlock after timelock has expired
	err = timelock.Unlock()
	if err != nil {
		t.Errorf("Unlock() error = %v, want nil", err)
	}
}
