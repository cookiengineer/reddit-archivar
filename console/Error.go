package console

import "os"
import "strings"

func Error(message string) {

	message = sanitize(message)
	offset := toOffset()

	Messages = append(Messages, NewMessage("Error", message))

	if strings.Contains(message, "\n") {

		var lines = strings.Split(message, "\n")

		if COLORS == true {

			for l := 0; l < len(lines); l++ {
				os.Stderr.WriteString("\u001b[41m" + offset + " " + lines[l] + "\u001b[K\n")
			}

			os.Stderr.WriteString("\u001b[0m")

		} else {

			for l := 0; l < len(lines); l++ {
				// os.Stdout.WriteString(offset + " " + lines[l] + "\n")
				os.Stderr.WriteString(offset + " " + lines[l] + "\n")
			}

		}

	} else {

		if COLORS == true {
			os.Stderr.WriteString("\u001b[41m" + offset + " " + message + "\u001b[K\u001b[0m\n")
		} else {
			os.Stderr.WriteString(offset + " " + message + "\n")
		}

	}

}
