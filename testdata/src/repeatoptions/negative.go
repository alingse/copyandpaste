package repeatoptions

type NewExportFunc func(arg1, arg2 any, opts ...ExportOpt) error

func export(_, _ any, _ ...ExportOpt) error {
	return nil
}

func Negative1() NewExportFunc {
	// just a type convert
	return NewExportFunc(export)
}

func Negative2() {
	f := Negative1()

	_ = f(nil, nil)
}
