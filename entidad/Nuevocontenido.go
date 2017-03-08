package entidad

import (
    "time"
    )

type Nuevocontenido struct {
        NUMCODIGOOPERACIONMAESTRA  float64   `db:"NUMCODIGOOPERACIONMAESTRA" json:"numcodigooperacionmaestra"`  
        VCHOPERACIONMAESTRA string   `db:"VCHOPERACIONMAESTRA" json:"vchoperacionmaestra"`  
        VCHCODIGOOPERACION string   `db:"VCHCODIGOOPERACION" json:"vchcodigooperacion"`  
        CHRCODIGOOPERACIONSERVICIO string   `db:"CHRCODIGOOPERACIONSERVICIO" json:"chrcodigooperacionservicio"`  
        DTEFECHA     time.Time `db:"DTEFECHA" json:"dtefecha"`
        CHRESTADO string   `db:"CHRESTADO" json:"chrestado"`  
        NUMCODIGOMASTER float64   `db:"NUMCODIGOMASTER" json:"numcodigomaster"`  
        NUMCODIGOITEM float64   `db:"NUMCODIGOITEM" json:"numcodigoitem"`  
} 
