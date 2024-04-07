package cmd


type invoker struct {
	repos []Report
}


func(i *invoker) Schedule(cmd Report) {
	i.repos = append(i.repos, cmd)
}

func(i *invoker) Run() {
	for _, r := range i.repos {
		r.Execute()
	}
}