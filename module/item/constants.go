package item

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

//enum area

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var itemStatusStringList = [3]string{"doing", "done", "deleted"}

func (item ItemStatus) String() string {
	return fmt.Sprintf(itemStatusStringList[item])
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range itemStatusStringList {
		if itemStatusStringList[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New(fmt.Sprintf("Invalid status string"))
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintf("There is an error at %s", value))
	}

	strValue := string(bytes)
	v, error := parseStr2ItemStatus(strValue)

	if error != nil {
		return errors.New(fmt.Sprintf("There is an error at %s", value))
	}
	*item = v

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnMarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStr2ItemStatus(str)

	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}

//end enum area
