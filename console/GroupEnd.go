package console

import "os"

func GroupEnd(message string) {

	if OFFSET > 0 {
		OFFSET--
	}

	message = sanitize(message)
	offset := toOffset()
	Messages = append(Messages, NewMessage("GroupEnd", message))

	if COLORS == true {
		os.Stdout.WriteString("\u001b[40m" + offset + "\\ " + message + "\u001b[K\u001b[0m\n")
		os.Stdout.WriteString("\u001b[40m" + offset + "\u001b[K\u001b[0m\n")
	} else {
		os.Stdout.WriteString(offset + "\\ " + message + "\n")
		os.Stdout.WriteString(offset + "\n")
	}

}
