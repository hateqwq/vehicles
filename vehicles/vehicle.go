package vehicles

import "fmt"

type Vehicle struct {
    Brand  string
    Model string
    Year  int
}

func (v *Vehicle) GetInfo() string {
    return fmt.Sprintf("%d %s %s", v.Year, v.Brand, v.Model)
}

