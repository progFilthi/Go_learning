package main

import "fmt"

type Beat struct {
	Title    string
	BPM      float64
	Producer string
}

func (b Beat) Describe() string {
	panic("implement me")
}

type SamplePack struct {
	Name       string
	NumSamples int
	Price      float64
}

type Listable interface {
	Describe() string
}

func (s *SamplePack) Describe() string {
	return fmt.Sprintf("%s - %d samples - %.2f", s.Name, s.NumSamples, s.Price)
}

func (s *SamplePack) ApplyDiscount(pct float64) float64 {
	s.Price *= 1 - pct/100

	return s.Price
}

func PrintStoreFronts(items []Listable) {

	for _, item := range items {
		fmt.Println(item.Describe())
	}

}

func main() {

	beat := Beat{"Midnight Drive", 140, "Emma"}

	pack := SamplePack{Name: "Dark samples", NumSamples: 12, Price: 29.99}

	pack.ApplyDiscount(20)

	items := []Listable{beat, &pack}

	PrintStoreFronts(items)
}
