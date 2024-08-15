package connectToDB

import (
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/orm"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type mysqlConnectionConfig struct {
	cf         *mysql.Config
	Alias      string
	RetryCount int
	RetryDelay time.Duration
	TimeZone   *time.Location
	DebugFlag  bool
}

func ConnectToDatabase(keys ...string) error {
	fmt.Println("Connection to DB")

	conString := "root:password@tcp(localhost:3306)/testing"
	fmt.Println("ConString:", conString)
	//deriving cf as type myqlcoontconfig interface
	cf := mysqlConnectionConfig{}
	//parseDSN is used to parse the constring into the struct
	err := cf.ParseDSN(conString)
	if err != nil {
		return errors.New("failed to parse con string into struct")
	}
	//sets all the values of the struct
	ISTLocation, _ := time.LoadLocation("Asia/Kolkata")
	cf.TimeZone = ISTLocation
	cf.Alias = "default"
	cf.RetryDelay = 500 * time.Millisecond
	cf.RetryCount = 1
	cf.DebugFlag = true
	err = beegoRegisterDB(cf)
	if err != nil {
		fmt.Println("beegoRegisterDB", err)
		return err
	}
	fmt.Println("DB connected")
	return nil
}

// this func register the DB using ORM frame work
func beegoRegisterDB(cf mysqlConnectionConfig) error {
	var err error
	for breaker := cf.RetryCount; breaker > 0; breaker-- {
		if breaker < cf.RetryCount {
			time.Sleep(cf.RetryDelay)
		}
		//check if database exsits
		_, err := orm.GetDB(cf.Alias)
		if err != nil {
			err = orm.RegisterDataBase(cf.Alias, "mysql", cf.FormatDSN())
		}
		if err != nil {
			fmt.Println("failed to register db")
			continue
		}
		//test newly registered database
		err = MySqlTest(cf.Alias)
		if err != nil {
			fmt.Println("failed to query")
			continue
		}
		orm.DefaultTimeLoc = cf.TimeZone
		orm.Debug = cf.DebugFlag
		break
	}
	//checks if any error occur,prints the value of alias ,dsn,error
	if err != nil {
		fmt.Println(fmt.Sprintf("beegoRegisterDB(alias: %s, dsn: %s) failed, error: %v", cf.Alias, cf.String(), err))
	}
	return err
}

func (c *mysqlConnectionConfig) ParseDSN(s string) error {
	var err error
	//invoking PARSEDSN method wiht string argument init
	c.cf, err = mysql.ParseDSN(s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func MySqlTest(alias string) error {
	o := orm.NewOrm()
	o.Using(alias)
	_, err := o.Raw("SELECT 1").Exec()
	return err
}

func (c *mysqlConnectionConfig) String() string {
	t := c.cf.Clone()
	t.Passwd = "root"
	return t.FormatDSN()
}

func (c *mysqlConnectionConfig) FormatDSN() string {
	return c.cf.FormatDSN()
}
