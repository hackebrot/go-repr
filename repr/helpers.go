package repr

import "time"

// boolPtr returns a *bool for the given value
func boolPtr(b bool) *bool {
	return &b
}

// intPtr returns a *int for the given value
func intPtr(i int) *int {
	return &i
}

// stringPtr returns a *string for the given value
func stringPtr(s string) *string {
	return &s
}

// timePtr returns a *string for the given value
func timePtr(t time.Time) *time.Time {
	return &t
}
