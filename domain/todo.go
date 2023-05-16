package domain

type Todo struct {
	title TodoTitle
	done TodoDone
}

type TodoTitle struct {
	value string
}

type TodoDone struct {
	value bool
}