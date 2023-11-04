package cmd

func Execute() {
	err := rootCMD.Execute()
	if err != nil {
		panic(err)
	}
}
