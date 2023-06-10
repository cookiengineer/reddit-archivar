package console

import "os"

func Group(message string) {

	offset := toOffset()
	OFFSET++

	message = sanitize(message)
	Messages = append(Messages, NewMessage("Group", message))

	if COLORS == true {
		os.Stdout.WriteString("\u001b[40m" + offset + "\u001b[K\u001b[0m\n")
		os.Stdout.WriteString("\u001b[40m" + offset + "/ " + message + "\u001b[K\u001b[0m\n")
	} else {
		os.Stdout.WriteString(offset + "\n")
		os.Stdout.WriteString(offset + "/ " + message + "\n")
	}

}
