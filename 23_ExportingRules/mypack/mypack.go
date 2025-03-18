package mypack

import (
	"fmt"
)

func SayHello(){
	// Exported function (Accessible outside the package)
	//✅ Names starting with an uppercase letter are exported (public) and can be accessed outside the package.
	fmt.Println("hello from mypack")
}

func privatefunc() {
	//// Unexported function (Only accessible inside mypackage)
	//✅ Names starting with an uppercase letter are exported (public) and can be accessed outside the package.
	fmt.Println("private func()")
}

func RunPrivate(){
	privatefunc()
}

// Exported struct
type Person struct {
	Name string // ✅ Exported field (public)
	age int  // ❌ Unexported field (private)
}
