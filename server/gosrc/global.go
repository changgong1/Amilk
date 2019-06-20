package gosrc
import(
	"./amilkdb"
)

var gDbc amilkdb.AmilkDBClient
var gLogLevel string

// 分数类型 
// NiScoerType = 0
const (
	NiScoreType = iota
	GdScoreType 
	MdScoreType
	BdScoerType
)

// 对象类型
const (
	CommentType = iota
	MovieType
	ArtistType
)