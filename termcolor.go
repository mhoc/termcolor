package termcolor

import "fmt"

const (
	CTRL_RESET            = "\033[0m"
	CTRL_BOLD             = "\033[1m"
	CTRL_DIM              = "\033[2m"
	CTRL_ITALIC           = "\033[3m"
	CTRL_UNDERLINE        = "\033[4m"
	CTRL_BLINK            = "\033[5m"
	CTRL_REVERSE          = "\033[7m"
	CTRL_HIDDEN           = "\033[8m"
	CTRL_STRIKE           = "\033[9m"
	CTRL_FG_BLACK         = "\033[30m"
	CTRL_FG_RED           = "\033[31m"
	CTRL_FG_GREEN         = "\033[32m"
	CTRL_FG_YELLOW        = "\033[33m"
	CTRL_FG_BLUE          = "\033[34m"
	CTRL_FG_MAGENTA       = "\033[35m"
	CTRL_FG_CYAN          = "\033[36m"
	CTRL_FG_LIGHT_GRAY    = "\033[37m"
	CTRL_BG_BLACK         = "\033[40m"
	CTRL_BG_RED           = "\033[41m"
	CTRL_BG_GREEN         = "\033[42m"
	CTRL_BG_YELLOW        = "\033[43m"
	CTRL_BG_BLUE          = "\033[44m"
	CTRL_BG_MAGENTA       = "\033[45m"
	CTRL_BG_CYAN          = "\033[46m"
	CTRL_BG_LIGHT_GRAY    = "\033[47m"
	CTRL_FG_DARK_GRAY     = "\033[90m"
	CTRL_FG_LIGHT_RED     = "\033[91m"
	CTRL_FG_LIGHT_GREEN   = "\033[92m"
	CTRL_FG_LIGHT_YELLOW  = "\033[93m"
	CTRL_FG_LIGHT_BLUE    = "\033[94m"
	CTRL_FG_LIGHT_MAGENTA = "\033[95m"
	CTRL_FG_LIGHT_CYAN    = "\033[96m"
	CTRL_FG_WHITE         = "\033[97m"
	CTRL_BG_DARK_GRAY     = "\033[100m"
	CTRL_BG_LIGHT_RED     = "\033[101m"
	CTRL_BG_LIGHT_GREEN   = "\033[102m"
	CTRL_BG_LIGHT_YELLOW  = "\033[103m"
	CTRL_BG_LIGHT_BLUE    = "\033[104m"
	CTRL_BG_LIGHT_MAGENTA = "\033[105m"
	CTRL_BG_LIGHT_CYAN    = "\033[106m"
	CTRL_BG_WHITE         = "\033[107m"
)

// Black colorizes the provided text as black
func Black(s string) string {
	return CTRL_FG_BLACK + s + CTRL_RESET
}

// Blackf is a fmt.Printf version of Black()
func Blackf(s string, args ...any) string {
	return fmt.Sprintf(Black(s), args...)
}

// Blue colorizes the provided text as blue
func Blue(s string) string {
	return CTRL_FG_BLUE + s + CTRL_RESET
}

// Bluef is a fmt.Printf version of Blue()
func Bluef(s string, args ...any) string {
	return fmt.Sprintf(Blue(s), args...)
}

// Cyan colorizes the provided text as cyan
func Cyan(s string) string {
	return CTRL_FG_CYAN + s + CTRL_RESET
}

// Cyanf is a fmt.Printf version of Cyan()
func Cyanf(s string, args ...any) string {
	return fmt.Sprintf(Cyan(s), args...)
}

// DarkGray colorizes the provided text as dark gray
func DarkGray(s string) string {
	return CTRL_FG_DARK_GRAY + s + CTRL_RESET
}

// DarkGrayf is a fmt.Printf version of DarkGray()
func DarkGrayf(s string, args ...any) string {
	return fmt.Sprintf(DarkGray(s), args...)
}

// Green coloorizes the provided text as green
func Green(s string) string {
	return CTRL_FG_GREEN + s + CTRL_RESET
}

// Greenf is a fmt.Printf version of Green()
func Greenf(s string, args ...any) string {
	return fmt.Sprintf(Green(s), args...)
}

// LightCyan colorizes the provided text as light cyan
func LightCyan(s string) string {
	return CTRL_FG_LIGHT_CYAN + s + CTRL_RESET
}

// LightCyanf is a fmt.Printf version of DarkGray()
func LightCyanf(s string, args ...any) string {
	return fmt.Sprintf(LightCyan(s), args...)
}

// LightGray colorizes the provided text as light gray
func LightGray(s string) string {
	return CTRL_FG_LIGHT_GRAY + s + CTRL_RESET
}

// LightGrayf is a fmt.Printf version of DarkGray()
func LightGrayf(s string, args ...any) string {
	return fmt.Sprintf(LightGray(s), args...)
}

// Red colorizes the provided text as red
func LightGreen(s string) string {
	return CTRL_FG_LIGHT_GREEN + s + CTRL_RESET
}

// LightGreenf is a fmt.Printf version of DarkGray()
func LightGreenf(s string, args ...any) string {
	return fmt.Sprintf(LightGreen(s), args...)
}

// LightMagenta colorizes the provided text as light magenta
func LightMagenta(s string) string {
	return CTRL_FG_LIGHT_MAGENTA + s + CTRL_RESET
}

// LightMagentaf is a fmt.Printf version of LightMagenta()
func LightMagentaf(s string, args ...any) string {
	return fmt.Sprintf(LightMagenta(s), args...)
}

// LightRed colorizes the provided text as light red
func LightRed(s string) string {
	return CTRL_FG_LIGHT_RED + s + CTRL_RESET
}

// LightRedf is a fmt.Printf version of LightRed()
func LightRedf(s string, args ...any) string {
	return fmt.Sprintf(LightRed(s), args...)
}

// Magenta colorizes the provided text as magenta
func Magenta(s string) string {
	return CTRL_FG_MAGENTA + s + CTRL_RESET
}

// Magentaf is a fmt.Printf version of Magenta()
func Magentaf(s string, args ...any) string {
	return fmt.Sprintf(Magenta(s), args...)
}

// Red colorizes the provided text as red
func Red(s string) string {
	return CTRL_FG_RED + s + CTRL_RESET
}

// Redf is a fmt.Printf version of Red()
func Redf(s string, args ...any) string {
	return fmt.Sprintf(Red(s), args...)
}

// White colorizes the provided text as white
func White(s string) string {
	return CTRL_FG_WHITE + s + CTRL_RESET
}

// Whitef is a fmt.Printf version of White()
func Whitef(s string, args ...any) string {
	return fmt.Sprintf(White(s), args...)
}

// Yellow colorizes the provided text as yellow
func Yellow(s string) string {
	return CTRL_FG_YELLOW + s + CTRL_RESET
}

// LightMagentaf is a fmt.Printf version of LightMagenta()
func Yellowf(s string, args ...any) string {
	return fmt.Sprintf(Yellow(s), args...)
}
