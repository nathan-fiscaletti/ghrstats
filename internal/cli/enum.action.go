package cli

type Action string

const (
	ActionAggregateTotal    = "total"
	ActionAggregateItemized = "itemized"
)

var Actions = []Action{
	ActionAggregateTotal,
	ActionAggregateItemized,
}
