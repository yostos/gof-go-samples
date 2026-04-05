// Mediator pattern: a LoginDialog mediator coordinates colleague widgets.
// CLI-based simulation demonstrates the same enable/disable logic as the Java GUI version.
package main

import "fmt"

// Colleague is anything the mediator can enable/disable.
type Colleague interface {
	SetEnabled(enabled bool)
	IsEnabled() bool
	Name() string
}

// RadioButton represents a selectable option (Guest / Login).
type RadioButton struct {
	name     string
	selected bool
	enabled  bool
}

func (r *RadioButton) Name() string         { return r.name }
func (r *RadioButton) SetEnabled(e bool)     { r.enabled = e }
func (r *RadioButton) IsEnabled() bool       { return r.enabled }
func (r *RadioButton) SetSelected(s bool)    { r.selected = s }
func (r *RadioButton) IsSelected() bool      { return r.selected }

// TextField represents a text input field.
type TextField struct {
	name    string
	text    string
	enabled bool
}

func (t *TextField) Name() string         { return t.name }
func (t *TextField) SetEnabled(e bool)     { t.enabled = e }
func (t *TextField) IsEnabled() bool       { return t.enabled }
func (t *TextField) SetText(s string)      { t.text = s }
func (t *TextField) Text() string          { return t.text }

// Button represents a clickable button.
type Button struct {
	name    string
	enabled bool
}

func (b *Button) Name() string         { return b.name }
func (b *Button) SetEnabled(e bool)     { b.enabled = e }
func (b *Button) IsEnabled() bool       { return b.enabled }

// LoginDialog is the mediator that coordinates all colleagues.
type LoginDialog struct {
	checkGuest  *RadioButton
	checkLogin  *RadioButton
	textUser    *TextField
	textPass    *TextField
	buttonOk    *Button
	buttonCancel *Button
}

// NewLoginDialog creates and initializes the dialog.
func NewLoginDialog() *LoginDialog {
	d := &LoginDialog{
		checkGuest:   &RadioButton{name: "Guest", selected: true, enabled: true},
		checkLogin:   &RadioButton{name: "Login", selected: false, enabled: true},
		textUser:     &TextField{name: "Username"},
		textPass:     &TextField{name: "Password"},
		buttonOk:     &Button{name: "OK"},
		buttonCancel: &Button{name: "Cancel", enabled: true},
	}
	d.colleagueChanged()
	return d
}

// colleagueChanged recalculates the enabled state of all colleagues.
func (d *LoginDialog) colleagueChanged() {
	if d.checkGuest.IsSelected() {
		// Guest mode: disable username/password, enable OK
		d.textUser.SetEnabled(false)
		d.textPass.SetEnabled(false)
		d.buttonOk.SetEnabled(true)
	} else {
		// Login mode: enable username, conditionally enable password and OK
		d.textUser.SetEnabled(true)
		if len(d.textUser.Text()) > 0 {
			d.textPass.SetEnabled(true)
			if len(d.textPass.Text()) > 0 {
				d.buttonOk.SetEnabled(true)
			} else {
				d.buttonOk.SetEnabled(false)
			}
		} else {
			d.textPass.SetEnabled(false)
			d.buttonOk.SetEnabled(false)
		}
	}
}

// selectGuest simulates selecting the Guest radio button.
func (d *LoginDialog) selectGuest() {
	d.checkGuest.SetSelected(true)
	d.checkLogin.SetSelected(false)
	d.colleagueChanged()
}

// selectLogin simulates selecting the Login radio button.
func (d *LoginDialog) selectLogin() {
	d.checkGuest.SetSelected(false)
	d.checkLogin.SetSelected(true)
	d.colleagueChanged()
}

// setUsername simulates typing in the username field.
func (d *LoginDialog) setUsername(text string) {
	d.textUser.SetText(text)
	d.colleagueChanged()
}

// setPassword simulates typing in the password field.
func (d *LoginDialog) setPassword(text string) {
	d.textPass.SetText(text)
	d.colleagueChanged()
}

// printState displays the current state of all colleagues.
func (d *LoginDialog) printState() {
	status := func(c Colleague) string {
		if c.IsEnabled() {
			return "enabled"
		}
		return "disabled"
	}

	fmt.Printf("  Guest=[%s selected=%v] Login=[%s selected=%v]\n",
		status(d.checkGuest), d.checkGuest.IsSelected(),
		status(d.checkLogin), d.checkLogin.IsSelected())
	fmt.Printf("  Username=[%s text=%q] Password=[%s text=%q]\n",
		status(d.textUser), d.textUser.Text(),
		status(d.textPass), d.textPass.Text())
	fmt.Printf("  OK=[%s] Cancel=[%s]\n",
		status(d.buttonOk), status(d.buttonCancel))
	fmt.Println()
}

func main() {
	dialog := NewLoginDialog()

	fmt.Println("=== Initial state (Guest selected) ===")
	dialog.printState()

	fmt.Println("=== Select Login ===")
	dialog.selectLogin()
	dialog.printState()

	fmt.Println("=== Type username: alice ===")
	dialog.setUsername("alice")
	dialog.printState()

	fmt.Println("=== Type password: secret ===")
	dialog.setPassword("secret")
	dialog.printState()

	fmt.Println("=== Clear password ===")
	dialog.setPassword("")
	dialog.printState()

	fmt.Println("=== Clear username ===")
	dialog.setUsername("")
	dialog.printState()

	fmt.Println("=== Select Guest again ===")
	dialog.selectGuest()
	dialog.printState()
}
