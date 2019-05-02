package bed2x

type Bed3 interface {
	Chr() string
	Start() int
	End() int
}
type Bed4 struct {
	chr   string
	start int
	end   int
	name  string
}

func (b Bed4) Chr() string {
	return b.chr
}

func (b Bed4) Start() int {
	return b.start
}
func (b Bed4) End() int {
	return b.end
}
func (b Bed4) Id() string {
	return b.name
}
