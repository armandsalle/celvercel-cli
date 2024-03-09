package prompt

import "testing"

func TestDrawPrompt_ver(t *testing.T) {
	got := DrawPrompt([]string{"ver"})
	want := `
   /\
  /__\`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestDrawPrompt_ver_ver(t *testing.T) {
	got := DrawPrompt([]string{"ver", "ver"})
	want := `
   /\
  /__\
   /\
  /__\`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestDrawPrompt_cel(t *testing.T) {
	got := DrawPrompt([]string{"cel"})
	want := `
 /¯¯¯¯\
/______\`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDrawPrompt_ver_cel(t *testing.T) {
	got := DrawPrompt([]string{"ver", "cel"})
	want := `
   /\
  /  \
 /    \
/______\`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestDrawPrompt_cel_ver(t *testing.T) {
	got := DrawPrompt([]string{"lec", "rev"})
	want := `
\¯¯¯¯¯¯/
 \    /
  \  /
   \/`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDrawPrompt_rev_rev_cel_cel_ver_rev(t *testing.T) {

	got := DrawPrompt([]string{"rev", "rev", "cel", "cel", "ver", "rev"})
	want := `
  \¯¯/
  _\/_
  \  /
  _\/_
 /    \
/_    _\
 /    \
/______\
   /\
  /  \
  \  /
   \/`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDrawPrompt_rev(t *testing.T) {

	got := DrawPrompt([]string{"rev"})
	want := `
  \¯¯/
   \/`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestDrawPrompt_lec(t *testing.T) {

	got := DrawPrompt([]string{"lec"})
	want := `
\¯¯¯¯¯¯/
 \____/`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestDrawPrompt_cel_cel_cel(t *testing.T) {

	got := DrawPrompt([]string{"cel", "cel", "cel"})
	want := `
 /¯¯¯¯\
/_    _\
 /    \
/_    _\
 /    \
/______\`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDrawPrompt_rev_lec_lec_rev_ver(t *testing.T) {

	got := DrawPrompt([]string{"rev", "lec", "lec", "rev", "ver"})
	want := `
  \¯¯/
 __\/__
\      /
_\    /_
\      /
 \    /
  \  /
   \/
   /\
  /__\`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestDrawPrompt_ver_ver_ver_lec_cel(t *testing.T) {

	got := DrawPrompt([]string{"ver", "ver", "ver", "lec", "cel"})
	want := `
   /\
  /__\
   /\
  /__\
   /\  
 _/  \_
\      /
 \    /
 /    \
/______\`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDrawPrompt_ver_cel_lec_ver(t *testing.T) {

	got := DrawPrompt([]string{"ver", "cel", "lec", "ver"})
	want := `
   /\
  /  \
 /    \
/      \
\      /
 \____/
   /\
  /__\`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
