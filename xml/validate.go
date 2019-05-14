package xml

import (
	"fmt"
	"io/ioutil"
	"unsafe"

	"github.com/jbussdieker/golibxml"
	"github.com/krolaw/xsd"
)

const xsdRef = "https://gist.githubusercontent.com/devmark88/2dda3dc85f9ed173be0c8461af77098d/raw/8c2218fdda9832c55be171e6607467a3445024dd/graph.xsd"

// Validate => validate xml against xsd
func Validate(xml string) error {
	xsdContest, err := ioutil.ReadFile("graph.xsd")
	if err != nil {
		fmt.Println(err)
		return err
	}
	xsdSchema, err := xsd.ParseSchema(xsdContest)
	if err != nil {
		fmt.Println(err)
		return err
	}

	doc := golibxml.ParseDoc(xml)
	if doc == nil {
		// TODO capture and display error - help please
		fmt.Println("Error parsing document")
		return fmt.Errorf("error parsing document")
	}
	defer doc.Free()

	// golibxml._Ctype_xmlDocPtr can't be cast to xsd.DocPtr, even though they are both
	// essentially _Ctype_xmlDocPtr.  Using unsafe gets around this.
	if err := xsdSchema.Validate(xsd.DocPtr(unsafe.Pointer(doc.Ptr))); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("XML Valid as per XSD")
	return nil
}
