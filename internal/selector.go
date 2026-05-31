package internal

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrInvalidRange       = errors.New("range start cannot be greater than range end")
	ErrFieldBelowOne      = errors.New("field must be greater than 0")
	ErrInvalidRangeFormat = errors.New("range must have exactly two parts separated by '-'")
	ErrFieldOutOfBounds   = errors.New("field index exceeds number of fields in line")
)

type FieldSelector struct {
	maxField int
	sorted   []int
}

func NewFieldSelector(fieldStr string) (*FieldSelector, error) {
	maxField := 0
	var fieldList []int

	if fieldStr == "" {
		return nil, ErrFieldStrEmpty
	}

	for part := range strings.SplitSeq(fieldStr, ",") {
		if strings.HasPrefix(part, "-") {
			return nil, fmt.Errorf("field %q: %w", part, ErrFieldBelowOne)
		}
		if strings.Contains(part, "-") {
			st, end, err := extractRange(part)
			if err != nil {
				return nil, err
			}
			err = validateRange(st, end)
			if err != nil {
				return nil, err
			}
			for i := st; i <= end; i++ {
				fieldList = append(fieldList, i)
			}
			if end > maxField {
				maxField = end
			}
		} else {
			field, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid field %q: %w", part, err)
			}
			if field < 1 {
				return nil, fmt.Errorf("field %d: %w", field, ErrFieldBelowOne)
			}
			fieldList = append(fieldList, field)
			if field > maxField {
				maxField = field
			}
		}
	}

	sort.Ints(fieldList)
	return &FieldSelector{
		maxField: maxField,
		sorted:   fieldList,
	}, nil
}

func (fs *FieldSelector) Extract(line string, delimiter string) (string, error) {
	words := strings.Split(line, delimiter)
	if fs.maxField > len(words) {
		return "", fmt.Errorf("field %d: %w", fs.maxField, ErrFieldOutOfBounds)
	}
	var result []string
	seen := -1
	for _, item := range fs.sorted {
		if item == seen {
			continue
		} else {
			seen = item
		}

		result = append(result, words[item-1])
	}
	return strings.Join(result, delimiter), nil
}

func validateRange(firstRange, secondRange int) error {
	if firstRange < 1 {
		return fmt.Errorf("start %d: %w", firstRange, ErrFieldBelowOne)
	}
	if secondRange < 1 {
		return fmt.Errorf("end %d: %w", secondRange, ErrFieldBelowOne)
	}
	if firstRange > secondRange {
		return fmt.Errorf("start %d, end %d: %w", firstRange, secondRange, ErrInvalidRange)
	}
	return nil
}

func extractRange(rangeStr string) (st, end int, err error) {
	rangeSplit := strings.Split(rangeStr, "-")
	if len(rangeSplit) != 2 {
		return 0, 0, fmt.Errorf("range %q: %w", rangeStr, ErrInvalidRangeFormat)
	}
	st, err = strconv.Atoi(rangeSplit[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid range start %q: %w", rangeSplit[0], err)
	}
	end, err = strconv.Atoi(rangeSplit[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid range end %q: %w", rangeSplit[1], err)
	}
	return st, end, nil
}
