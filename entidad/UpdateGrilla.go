package entidad

type UpdateGrilla struct {
    VCHCODIGOOPERACION string `db:"VCHCODIGOOPERACION" json:"vchcodigooperacion"`
    CHRCODIGOOPERACIONSERVICIO string `db:"CHRCODIGOOPERACIONSERVICIO" json:"chrcodigooperacionservicio"`
    NUMHORASHOMBRE   float64 `db:"NUMHORASHOMBRE" json:"numhorashombre"`
    NUMPRECIOSUGERIDO float64 `db:"NUMPRECIOSUGERIDO" json:"numpreciosugerido"`
    NUMDESCUENTO      float64 `db:"NUMDESCUENTO" json:"numdescuento"`
    NUMTOTAL float64 `json:"numtotal"`
}
