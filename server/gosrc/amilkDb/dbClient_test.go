package amilkdb
import(
	. "gopkg.in/check.v1"
	_ "../mlogger"
	"testing"
)
func Test(t *testing.T) {
	TestingT(t)
}
type DbClient struct {
	testGdbClient AmilkDBClient
}
var _ = Suite(&DbClient{})

func (d *DbClient) initTest() error {
	err := d.testGdbClient.InitDBConfig("root", "123456", "localhost", "3306", "amilk", 100, 10)
	if err != nil {
		return err
	}
	return nil
}


func (d *DbClient) SetUpSuite(c *C){
	err := d.initTest()
	c.Assert(err, IsNil)
}

func (d *DbClient) TearDownSuite(c *C) {
}
func (d *DbClient) SetUpTest(c *C){
}
func (d *DbClient) TearDownTest(c *C){	
}