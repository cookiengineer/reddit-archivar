package schemas

import "strconv"
import "strings"
import "time"

type Timestamp time.Time

func (timestamp *Timestamp) UnmarshalJSON(buffer []byte) error {

	tmp := string(buffer)

	if strings.Contains(tmp, ".") {
		tmp = tmp[0:strings.Index(tmp, ".")]
	}

	check, err := strconv.ParseInt(tmp, 10, 64)

	if err != nil {
		return err
	}

	*timestamp = Timestamp(time.Unix(check, 0))

	return nil

}

func (timestamp *Timestamp) String() string {
	return time.Time(*timestamp).Format(time.RFC3339)
}
