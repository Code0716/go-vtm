package util

import "time"

const layout = "2006-01-02 15:04:05"

// TimeFromStr is format mysql style
func TimeFromStr(s string) time.Time {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// タイムゾーンの取得に失敗した場合の処理
		return time.Time{}
	}

	t, err := time.ParseInLocation(layout, s, loc)
	if err != nil {
		// 時間のパースに失敗した場合の処理
		return time.Time{}
	}
	return t
}

// TimeToString time to string
func TimeToString(t time.Time) string {
	s := t.Format(layout)
	return s
}
