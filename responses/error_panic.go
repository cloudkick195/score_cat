package responses

func ErrorServer(err error) {
	if err != nil {
		panic(err)
	}
	return
}
