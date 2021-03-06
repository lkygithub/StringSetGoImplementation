package sets

// An Empty struct does not take up any memory
type Empty struct{}

// sets.String is a set of strings, implemented via map[string]struct{} for minimal memory consumption.
type String map[string]Empty


func NewString(items ...string) String {
	s := String{}
	s.Insert(items...)
	return s
}

func (s String) Insert(items ...string) {
	for _, item := range items {
		s[item] = Empty{}
	}
}

func (s String) Delete(items ...string) {
	for _, item := range items {
		delete(s, item)
	}
}

func (s String) Has(item string) bool {
	_, res := s[item]
	return res
}

// HasAll returns true if and only if all items are contained in the set.
func (s String) HasAll(items ...string) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s String) HasAny(items ...string) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Difference returns a set of objects that are not in s2
func (s String) Difference(s2 String) String {
	res := NewString()
	for item := range s {
		if !s2.Has(item) {
			res.Insert(item)
		}
	}
	return res
}

// Union returns a new set which includes items in either s1 or s2.
func (s1 String) Union(s2 String) String {
	res := NewString()
	for item := range s1 {
		res.Insert(item)
	}
	for item := range s2 {
		res.Insert(item)
	}
	return res
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
func (s1 String) Intersection(s2 String) String {
	res := NewString()
	if s1.Len() < s2.Len() {
		for item := range s1 {
			if s2.Has(item) {
				res.Insert(item)
			}
		}
	} else {
		for item := range s2 {
			if s1.Has(item) {
				res.Insert(item)
			}
		}
	}
	return res
}

func (s String) Len() int {
	return len(s)
}

func (s String) List() []string {
	res := make([]string, 0, len(s))
	for item := range s {
		res = append(res, item)
    }
	return res
}

