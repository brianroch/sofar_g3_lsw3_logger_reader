package otlp

import (
	"testing"
	"time"
)

func Test_getTariff(t *testing.T) {
	time01Val := "3:00AM"
	time01, _ := time.Parse(time.Kitchen, time01Val)
	time02Val := "8:00AM"
	time02, _ := time.Parse(time.Kitchen, time02Val)
	time03Val := "8:05AM"
	time03, _ := time.Parse(time.Kitchen, time03Val)
	time04Val := "2:05PM"
	time04, _ := time.Parse(time.Kitchen, time04Val)
	time05Val := "5:00PM"
	time05, _ := time.Parse(time.Kitchen, time05Val)
	time06Val := "5:05PM"
	time06, _ := time.Parse(time.Kitchen, time06Val)
	time07Val := "7:00PM"
	time07, _ := time.Parse(time.Kitchen, time07Val)
	time08Val := "7:05PM"
	time08, _ := time.Parse(time.Kitchen, time08Val)
	time09Val := "9:00PM"
	time09, _ := time.Parse(time.Kitchen, time09Val)
	time10Val := "9:05PM"
	time10, _ := time.Parse(time.Kitchen, time10Val)
	time11Val := "11:00PM"
	time11, _ := time.Parse(time.Kitchen, time11Val)
	time12Val := "11:05PM"
	time12, _ := time.Parse(time.Kitchen, time12Val)

	type args struct {
		t time.Time
	}
	tests := []struct {
		name      string
		args      args
		wantName  string
		wantIndex string
	}{
		{name: time01Val, args: args{time01}, wantName: tariffNight, wantIndex: "0"},
		{name: time02Val, args: args{time02}, wantName: tariffDay, wantIndex: "0"},
		{name: time03Val, args: args{time03}, wantName: tariffDay, wantIndex: "0"},
		{name: time04Val, args: args{time04}, wantName: tariffDay, wantIndex: "0"},
		{name: time05Val, args: args{time05}, wantName: tariffPeak, wantIndex: "0"},
		{name: time06Val, args: args{time06}, wantName: tariffPeak, wantIndex: "0"},
		{name: time07Val, args: args{time07}, wantName: tariffDay, wantIndex: "1"},
		{name: time08Val, args: args{time08}, wantName: tariffDay, wantIndex: "1"},
		{name: time09Val, args: args{time09}, wantName: tariffDay, wantIndex: "1"},
		{name: time10Val, args: args{time10}, wantName: tariffDay, wantIndex: "1"},
		{name: time11Val, args: args{time11}, wantName: tariffNight, wantIndex: "1"},
		{name: time12Val, args: args{time12}, wantName: tariffNight, wantIndex: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, i := getTariffNameAndIndex(tt.args.t); got != tt.wantName || i != tt.wantIndex {
				t.Errorf("getTariffNameAndIndex() = %v, %v, wantName %v, wantIndex %v", got, i, tt.wantName, tt.wantIndex)
			}
		})
	}
}
