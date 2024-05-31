package mapper

import (
	"context"
	"time"
)

type TimeStringConverter struct {
}

func (m *TimeStringConverter) StringToTime(ctx context.Context, source string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, source)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (m *TimeStringConverter) TimeToString(ctx context.Context, source *time.Time) (string, error) {
	return source.Format(time.RFC3339), nil
}

func AddTimeToStringConverter(mappers interface {
	Add(string, any)
}) {
	stringTime := &TimeStringConverter{}
	mappers.Add("TimeStringConverter", stringTime)
}
