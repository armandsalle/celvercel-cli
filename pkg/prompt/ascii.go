package prompt

type AsciiArt struct {
	Only, Open, Close, Default                                 string
	BeforeVer, BeforeCel, BeforeRev, BeforeLec                 string
	OpenBeforeVer, OpenBeforeCel, OpenBeforeRev, OpenBeforeLec string
}

var VerAscii = AsciiArt{
	Default: "\n   /\\\n  /  \\",
	Only:    "\n   /\\\n  /__\\",
	Open:    "\n   /\\\n  /  \\",
	Close:   "\n   /\\\n  /__\\",

	BeforeVer: "\n   /\\\n  /__\\",
	BeforeLec: "\n   /\\  \n _/  \\_",

	OpenBeforeVer: "\n   /\\\n  /__\\",
	OpenBeforeLec: "\n   /\\\n__/  \\__",
}

var RevAscii = AsciiArt{
	Default: "\n  \\  /\n   \\/",
	Only:    "\n  \\¯¯/\n   \\/",
	Open:    "\n  \\¯¯/\n   \\/",
	Close:   "\n  \\  /\n   \\/",

	BeforeRev: "\n  \\  /\n  _\\/_",
	BeforeLec: "\n  \\  /\n __\\/__",
	BeforeCel: "\n  \\  /\n  _\\/_",

	OpenBeforeRev: "\n  \\¯¯/\n  _\\/_",
	OpenBeforeLec: "\n  \\¯¯/\n __\\/__",
	OpenBeforeCel: "\n  \\¯¯/\n  _\\/_",
}

var CelAscii = AsciiArt{
	Default: "\n /    \\\n/      \\",
	Only:    "\n /¯¯¯¯\\\n/______\\",
	Open:    "\n /¯¯¯¯\\\n/      \\",
	Close:   "\n /    \\\n/______\\",

	BeforeCel: "\n /    \\\n/_    _\\",
	BeforeVer: "\n /    \\\n/______\\",
	BeforeRev: "\n /    \\\n/_    _\\",

	OpenBeforeCel: "\n /¯¯¯¯\\\n/_    _\\",
	OpenBeforeVer: "\n /¯¯¯¯\\\n/______\\",
}

var LecAscii = AsciiArt{
	Default: "\n\\      /\n \\    /",
	Only:    "\n\\¯¯¯¯¯¯/\n \\____/",
	Open:    "\n\\¯¯¯¯¯¯/\n \\    /",
	Close:   "\n\\      /\n \\____/",

	BeforeLec: "\n\\      /\n_\\    /_",
	BeforeVer: "\n\\      /\n \\____/",

	OpenBeforeLec: "\n\\¯¯¯¯¯¯/\n_\\    /_",
	OpenBeforeVer: "\n\\¯¯¯¯¯¯/\n \\____/",
}

func (a *AsciiArt) HandleBefore(word string) string {
	if word == "ver" && a.BeforeVer != "" {
		return a.BeforeVer
	}

	if word == "cel" && a.BeforeCel != "" {
		return a.BeforeCel
	}

	if word == "rev" && a.BeforeRev != "" {
		return a.BeforeRev
	}

	if word == "lec" && a.BeforeLec != "" {
		return a.BeforeLec
	}

	return ""
}

func (a *AsciiArt) HandleOpenBefore(word string) string {
	if word == "ver" && a.OpenBeforeVer != "" {
		return a.OpenBeforeVer
	}

	if word == "cel" && a.OpenBeforeCel != "" {
		return a.OpenBeforeCel
	}

	if word == "rev" && a.OpenBeforeRev != "" {
		return a.OpenBeforeRev
	}

	if word == "lec" && a.OpenBeforeLec != "" {
		return a.OpenBeforeLec
	}

	return ""
}
