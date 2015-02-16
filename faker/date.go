package faker

import "time"

func randomDate() time.Time {
	return time.Now().AddDate(-randomInRange(0, 3), -randomInRange(0, 12), -randomInRange(0, 360))
}
