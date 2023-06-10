package console

import "os"
import "strings"

func Info(message string) {

	message = sanitize(message)
	offset := toOffset()

	Messages = append(Messages, NewMessage("Info", message))

	if strings.Contains(message, "\n") {

		var lines = strings.Split(message, "\n")

		if COLORS == true {

			for l := 0; l < len(lines); l++ {
				os.Stdout.WriteString("\u001b[42m" + offset + " " + lines[l] + "\u001b[K\n")
			}

			os.Stdout.WriteString("\u001b[0m")

		} else {

			for l := 0; l < len(lines); l++ {
				os.Stdout.WriteString(offset + " " + lines[l] + "\n")
			}

		}

	} else {

		if COLORS == true {
			os.Stdout.WriteString("\u001b[42m" + offset + " " + message + "\u001b[K\u001b[0m\n")
		} else {
			os.Stdout.WriteString(offset + " " + message + "\n")
		}

	}

}
