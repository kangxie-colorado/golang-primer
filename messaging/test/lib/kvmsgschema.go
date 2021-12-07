package lib

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// scheme for msg: SET|GET|DEL + LEN-of-KEY(5 bytes) + Key + [LEN-of-VAL(5 bytes) + Val]

func parseKeyValue(msg string) (string, string) {
	keyLenStr := msg[3:8]
	keyLen, _ := strconv.Atoi(keyLenStr)
	key := msg[8 : 8+keyLen]

	valLenStr := msg[8+keyLen : 13+keyLen]
	valLen, _ := strconv.Atoi(valLenStr)
	value := msg[13+valLen:]

	log.Infoln("Key:", key, "Value:", value)

	return key, value
}

func parseKey(msg string) string {
	keyLenStr := msg[3:8]
	keyLen, _ := strconv.Atoi(keyLenStr)
	key := msg[8 : 8+keyLen]
	log.Infoln("Key:", key)

	return key
}
func encodeKeyValue(key, value string) string {
	return fmt.Sprintf("%05d", len(key)) + key + fmt.Sprintf("%05d", len(value)) + value
}

func encodeKey(key string) string {
	return fmt.Sprintf("%05d", len(key)) + key
}
