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

func DoAny(_ ...any) {

}

func Negative3() {
	DoAny(WithWidth(1), WithHeight(1), WithWidth(1))

	_ = []any{WithWidth(1), WithHeight(1), WithWidth(1)}

	_ = []func(c *Cfg){
		WithWidth(1),
		nil,
		nil,
	}

	_ = [2]func(_, _ any){
		func(_, _ any) {},
		func(_, _ any) {},
	}
}
