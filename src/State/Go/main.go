// State pattern is still useful in Go. A State interface with per-state implementations
// cleanly encapsulates behavior that varies by state (day vs night security vault system).
package main

import "fmt"

// State defines the behavior that changes based on the current state.
type State interface {
	DoClock(ctx *SafeFrame, hour int)
	DoUse(ctx *SafeFrame)
	DoAlarm(ctx *SafeFrame)
	DoPhone(ctx *SafeFrame)
	String() string
}

// DayState represents daytime behavior (9:00-16:59).
type DayState struct{}

func (d *DayState) DoClock(ctx *SafeFrame, hour int) {
	if hour < 9 || 17 <= hour {
		ctx.ChangeState(&NightState{})
	}
}

func (d *DayState) DoUse(ctx *SafeFrame) {
	ctx.RecordLog("金庫使用(昼間)")
}

func (d *DayState) DoAlarm(ctx *SafeFrame) {
	ctx.CallSecurityCenter("非常ベル(昼間)")
}

func (d *DayState) DoPhone(ctx *SafeFrame) {
	ctx.CallSecurityCenter("通常の通話(昼間)")
}

func (d *DayState) String() string {
	return "[昼間]"
}

// NightState represents nighttime behavior (17:00-8:59).
type NightState struct{}

func (n *NightState) DoClock(ctx *SafeFrame, hour int) {
	if 9 <= hour && hour < 17 {
		ctx.ChangeState(&DayState{})
	}
}

func (n *NightState) DoUse(ctx *SafeFrame) {
	ctx.CallSecurityCenter("非常：夜間の金庫使用！")
}

func (n *NightState) DoAlarm(ctx *SafeFrame) {
	ctx.CallSecurityCenter("非常ベル(夜間)")
}

func (n *NightState) DoPhone(ctx *SafeFrame) {
	ctx.RecordLog("夜間の通話録音")
}

func (n *NightState) String() string {
	return "[夜間]"
}

// SafeFrame is the context that holds the current state and handles actions.
type SafeFrame struct {
	state State
}

func NewSafeFrame() *SafeFrame {
	return &SafeFrame{state: &DayState{}}
}

func (sf *SafeFrame) SetClock(hour int) {
	fmt.Printf("現在時刻は%02d:00\n", hour)
	sf.state.DoClock(sf, hour)
}

func (sf *SafeFrame) ChangeState(state State) {
	fmt.Printf("%sから%sへ状態が変化しました。\n", sf.state, state)
	sf.state = state
}

func (sf *SafeFrame) CallSecurityCenter(msg string) {
	fmt.Println("call!", msg)
}

func (sf *SafeFrame) RecordLog(msg string) {
	fmt.Println("record ...", msg)
}

func (sf *SafeFrame) DoUse() {
	sf.state.DoUse(sf)
}

func (sf *SafeFrame) DoAlarm() {
	sf.state.DoAlarm(sf)
}

func (sf *SafeFrame) DoPhone() {
	sf.state.DoPhone(sf)
}

func main() {
	frame := NewSafeFrame()

	// Simulate one 24-hour cycle, triggering actions at key transition points.
	for hour := 0; hour < 24; hour++ {
		frame.SetClock(hour)

		// Simulate user actions at certain hours to demonstrate state-dependent behavior.
		switch hour {
		case 1: // Night: use vault triggers emergency
			fmt.Println("[Action] 金庫使用")
			frame.DoUse()
			fmt.Println("[Action] 非常ベル")
			frame.DoAlarm()
			fmt.Println("[Action] 通常通話")
			frame.DoPhone()
		case 10: // Day: use vault is normal
			fmt.Println("[Action] 金庫使用")
			frame.DoUse()
			fmt.Println("[Action] 非常ベル")
			frame.DoAlarm()
			fmt.Println("[Action] 通常通話")
			frame.DoPhone()
		case 20: // Night again
			fmt.Println("[Action] 金庫使用")
			frame.DoUse()
			fmt.Println("[Action] 通常通話")
			frame.DoPhone()
		}
	}
}
