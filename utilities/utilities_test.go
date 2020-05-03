package utilities

import "testing"

func TestContains(t *testing.T) {
	imageNames := []string{"alpine", "busybox", "ubuntu", "windows"}
	if !Contains(imageNames, "alpine") || !Contains(imageNames, "busybox") ||
		!Contains(imageNames, "ubuntu") || !Contains(imageNames, "windows") {
		t.Error("Error on Contains function")
	}
	if Contains(imageNames, "kali") || Contains(imageNames, "") {
		t.Error("Error on Contains function")
	}
}
