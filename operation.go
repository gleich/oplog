package oplog

type Operation string

func (o Operation) Extend(opt string) Operation {
	return Operation(string(o) + "," + opt)
}
