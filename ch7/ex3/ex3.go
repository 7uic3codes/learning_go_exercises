package main

import "io"

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	res := r.Ranking()
	for _, v := range res {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func main() {

}
