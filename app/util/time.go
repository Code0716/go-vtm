package util

import "time"

const layout = "2006-01-02 15:04:05"

// TimeFromStr is format mysql style
func TimeFromStr(s string) time.Time {
	t, _ := time.Parse(layout, s)
	return t
	// utcLoc, err := time.LoadLocation("UTC")
	// if err != nil {
	// 	panic(err)
	// }
	// utc, err := time.ParseInLocation("2006-01-02 15:04:05", s, utcLoc)
	// if err != nil {
	// 	panic(err)
	// }

	// jstLoc, err := time.LoadLocation("Asia/Tokyo")
	// if err != nil {
	// 	panic(err)
	// }
	// jst := utc.In(jstLoc)
	// return jst
}

// TimeToString time to string
func TimeToString(t time.Time) string {
	s := t.Format(layout)
	return s
}
