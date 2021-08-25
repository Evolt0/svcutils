package sm

const RAND = 16

func Sm4Key() string {
	return RandKey(RAND)
}
