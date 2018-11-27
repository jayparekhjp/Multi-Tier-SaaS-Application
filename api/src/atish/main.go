package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

type Emp struct {
	restraunt_id string
	restraunt_name string
	restraunt_address string
	zip_code  int
}

/*type Emp struct {
	id        string
	firstName string
	lastName  string
	age       int
}*/

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "cmpe281"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}

func main() {
	// emp1 := Emp{"1399a878-7e0d-4439-9dd4-32ad58dd0cd2", "Anupam", "Raj", 20}
	// emp2 := Emp{"E-2", "Rahul", "Anand", 30}
	// createEmp(emp1)
	// fmt.Println(getEmps())
	// createEmp(emp2)
	fmt.Println(getEmps())
	// emp3 := Emp{"E-1", "Rahul", "Anand", 30}
	// updateEmp(emp3)
	// fmt.Println(getEmps())
	// deleteEmp("E-2")
	// fmt.Println(getEmps())
}

/*func getEmps() []Emp {
	fmt.Println("Getting all Employees")
	var emps []Emp
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM emps").Iter()
	for iter.MapScan(m) {
		emps = append(emps, Emp{
			id:        m["empid"].(string),
			firstName: m["first_name"].(string),
			lastName:  m["last_name"].(string),
			age:       m["age"].(int),
		})
		m = map[string]interface{}{}
	}

	return emps
}*/

func getEmps() []Emp {
	var emps []Emp
	// m := map[string]interface{}{}
	var restraunt_name string
	var restraunt_address string
	iter := Session.Query("SELECT * FROM restraunts").Iter()
	for {
		// New map each iteration
		row := map[string]interface{}{
			restraunt_name: &restraunt_name,
			restraunt_address:  &restraunt_address,	
		}
		if !iter.MapScan(row) {
			break
		}
		fmt.Printf("First: %s Address: %s\n", restraunt_name, restraunt_address)
	}
	for iter.Scan(&restraunt_address,&restraunt_name) {
	   // fmt.Println("number:", number_num)
	   /*var  data = Identification{
	    number : number_num,
	    name : common_name,
	   }*/
	   fmt.Println(restraunt_name)
	   fmt.Println(restraunt_address)
	   // json.NewEncoder(w).Encode(data)
	  }
	/*for iter.MapScan(m) {
		emps = append(emps, Emp{
			// restraunt_id:        m["restraunt_id"].(string),
			restraunt_name: m["restraunt_name"].(string),
			restraunt_address:  m["restraunt_address"].(string),
			zip_code:       m["zip_code"].(int),
		})
		m = map[string]interface{}{}
	}*/
	fmt.Println(iter.NumRows())
	return emps
}

/*func createEmp(emp Emp) {
	fmt.Println(" **** Creating new emp ****\n", emp)
	if err := Session.Query("INSERT INTO restraunt(empid, first_name, last_name, age) VALUES(?, ?, ?, ?)",
		emp.id, emp.firstName, emp.lastName, emp.age).Exec(); err != nil {
		fmt.Println("Error while inserting Emp")
		fmt.Println(err)
	}
}*/