package workitems;



type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}


function NewMatrix(rows, cols int) *matrix {
    m := new(matrix)
    m.rows = rows
    m.cols = cols
    m.elems = make([]float, rows*cols)
    return m
}

type Integer int;
func (i *Integer) String() string {
    return strconv.itoa(i)
}