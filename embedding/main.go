package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

/**ejercicion*/
type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

// * Create functions to calculate averages for each dashboard component
func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for i := 0; i < len(c.temp); i++ {
		sum += int(c.temp[i])
	}
	return sum / len(c.temp)
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

//   - Using struct embedding, create a Dashboard structure that contains
//     each dashboard component
type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

type BeltSize int
type Shippinig int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shippinig) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shippinig
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shippinig {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	waight int
}

func (t *ToxicWaste) Shipe() Shippinig {
	return Ground
}

func main() {
	mail := SpamMail{4000}
	automate(&mail)

	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	//* Print out a 5-second average from each component using promoted
	//  methods on the Dashboard
	fmt.Printf("Average bandwidth usage: %v\n", dash.AverageBandwidth())
	fmt.Printf("Average temp: %v\n", dash.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", dash.AverageMemoryUsage())

	//waste := ToxicWaste{1000}

}
