package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

type fnCreate struct {
}

func init() {
	function.Register(&fnCreate{})
}

func (s *fnCreate) Name() string {
	return "create"
}

func (s *fnCreate) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt, data.TypeInt, data.TypeInt, data.TypeInt, data.TypeInt, data.TypeInt, data.TypeInt, data.TypeString}, false
}

func (s *fnCreate) Eval(in ...interface{}) (interface{}, error) {
	years, err := coerce.ToInt(in[0])
	if err != nil {
		return nil, err
	}
	months, err := coerce.ToInt(in[1])
	if err != nil {
		return nil, err
	}
	days, err := coerce.ToInt(in[2])
	if err != nil {
		return nil, err
	}

	hh, err := coerce.ToInt(in[3])
	if err != nil {
		return nil, err
	}
	mm, err := coerce.ToInt(in[4])
	if err != nil {
		return nil, err
	}
	ss, err := coerce.ToInt(in[5])
	if err != nil {
		return nil, err
	}

	nsec, err := coerce.ToInt(in[6])
	if err != nil {
		return nil, err
	}

	loc, err := coerce.ToString(in[7])
	if err != nil {
		return nil, err
	}

	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	return time.Date(years, time.Month(months), days, hh, mm, ss, nsec, l), nil
}
