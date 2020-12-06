package videogame

type VGSupport uint

const (
	DEMAT    VGSupport = iota
	PHYSICAL VGSupport = iota
)

var support = map[VGSupport]string{
	DEMAT:    "dématérialisé",
	PHYSICAL: "physique",
}

func (vgs *VGSupport) String() string {
	return support[*vgs]
}

type VideoGame struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"not null"`
	Support     VGSupport
	ExtensionOf uint `gorm:"default null"`
	Platform    uint `gorm:"foreignkey:ID;column:plateform_id"`
	Type        uint `gorm:"foreignkey:ID;column:type_id"`
}
