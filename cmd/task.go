package cmd

func addTask(s string) error {
	if err := AppendData(s); err != nil {
		return err
	}
	return nil
}
