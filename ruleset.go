package opts

type RuleSet struct {
	Name       string
	Dependents []*RuleSet
	Conflicts  []*RuleSet
}

//Cria uma nova ruleset de Dependents e Conflicts com slices vazios
func NewRuleSet(name string) *RuleSet {
	var deps []*RuleSet
	var conflicts []*RuleSet

	return &RuleSet{
		Name:       name,
		Dependents: deps,
		Conflicts:  conflicts,
	}
}

//Add uma dependencia RuleSet caso ela não exista	
func (r *RuleSet) AddDep(dep *RuleSet) {
	for _, v := range r.Dependents {
		if v.Name == dep.Name {
			return
		}
	}
	r.Dependents = append(r.Dependents, dep)
}

//adiciona um conflito a RuleSet caso o clinflito não exista
func (r *RuleSet) AddConflict(dep *RuleSet) {

	for _, v := range r.Conflicts {
		if v.Name == dep.Name {
			return
		}
	}

	r.Conflicts = append(r.Conflicts, dep)
}

//retorna true se RuleSet for coerente verifcando em Dependents e Conflicts
func (r *RuleSet) IsCoherent() bool {

	for _, v := range r.Dependents {
		for _, va := range r.Conflicts {
			if v.Name == va.Name {
				return false
			}
		}
	}

	return true
}