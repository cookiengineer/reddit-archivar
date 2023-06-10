package console

import "os"

func GroupEndResult(result bool, message string) {

	if OFFSET > 0 {
		OFFSET--
	}

	message = sanitize(message)
	offset := toOffset()

	if result == true {

		message += " succeeded"
		Messages = append(Messages, NewMessage("GroupEnd", message))

		if COLORS == true {
			os.Stdout.WriteString("\u001b[42m" + offset + "\\ " + message + "\u001b[K\u001b[0m\n")
			os.Stdout.WriteString("\u001b[42m" + offset + "\u001b[K\u001b[0m\n")
		} else {
			os.Stdout.WriteString(offset + "\\ " + message + "\n")
			os.Stdout.WriteString(offset + "\n")
		}

	} else {

		message += " failed"
		Messages = append(Messages, NewMessage("GroupEnd", message))

		if COLORS == true {
			os.Stdout.WriteString("\u001b[41m" + offset + "\\ " + message + "\u001b[K\u001b[0m\n")
			os.Stdout.WriteString("\u001b[41m" + offset + "\u001b[K\u001b[0m\n")
		} else {
			os.Stdout.WriteString(offset + "\\ " + message + "\n")
			os.Stdout.WriteString(offset + "\n")
		}

	}

}
