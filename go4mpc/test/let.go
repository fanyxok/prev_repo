package test

func letItCrash(err error) {
	if err != nil {
		panic(err)
	}
}
