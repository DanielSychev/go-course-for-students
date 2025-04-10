package tagcloud

import (
	"sort"
)

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	// TODO: add fields if necessary
	a map[string]*TagStat
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
	position        int
}

// New should create a valid TagCloud instance
// TODO: You decide whether this function should return a pointer or a value
func New() *TagCloud {
	// TODO: Implement this
	tmp := new(TagCloud)
	tmp.a = map[string]*TagStat{}
	return tmp
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
// TODO: You decide whether receiver should be a pointer or a value
func (t *TagCloud) AddTag(tag string) {
	// TODO: Implement this
	if _, find := t.a[tag]; !find {
		t.a[tag] = &TagStat{Tag: tag, OccurrenceCount: 0, position: len(t.a)}
	}
	t.a[tag].OccurrenceCount++
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
// TODO: You decide whether receiver should be a pointer or a value
func (t *TagCloud) TopN(n int) []TagStat {
	// TODO: Implement this
	a := []TagStat{}
	for _, elem := range t.a {
		a = append(a, *elem)
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].OccurrenceCount != a[j].OccurrenceCount {
			return a[i].OccurrenceCount > a[j].OccurrenceCount
		}
		return a[i].position < a[j].position
	})
	return a[:min(n, len(a))]
}
