// Package lenconv performs kilometers and miles conversions
package lenconv

import "fmt"

type Kilometer float64
type Mile float64

func (km Kilometer) String() string { return fmt.Sprintf("%gkm", km) }
func (mi Mile) String() string      { return fmt.Sprintf("%gmi", mi) }
