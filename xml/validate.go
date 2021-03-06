package xml

import (
	"fmt"
	"unsafe"

	"github.com/jbussdieker/golibxml"
	"github.com/krolaw/xsd"
)

// XML => Manage your xml content
type XML struct {
	Body string
}

const xsdRef = "https://gist.githubusercontent.com/devmark88/2dda3dc85f9ed173be0c8461af77098d/raw/8c2218fdda9832c55be171e6607467a3445024dd/graph.xsd"
const xsdContent = `<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    <xs:element name="graph">
        <xs:complexType>
            <xs:sequence>
                <xs:element name="id" type="xs:string" minOccurs="1" maxOccurs="1"></xs:element>
                <xs:element name="name" type="xs:string" minOccurs="1" maxOccurs="1"></xs:element>
                <xs:element name="nodes">
                    <xs:complexType>
                        <xs:sequence>
                            <xs:element name="node" maxOccurs="unbounded" minOccurs="1">
                                <xs:complexType>
                                    <xs:sequence>
                                        <xs:element name="id" type="xs:string" minOccurs="1" maxOccurs="1"></xs:element>
                                        <xs:element name="name" type="xs:string" minOccurs="1" maxOccurs="1"></xs:element>
                                    </xs:sequence>
                                </xs:complexType>
                            </xs:element>
                        </xs:sequence>
                    </xs:complexType>
                    <xs:unique name="nodeId">
                        <xs:selector xpath="node" />
                        <xs:field xpath="id" />
                    </xs:unique>
                </xs:element>
                <xs:element name="edges">
                    <xs:complexType>
                        <xs:sequence>
                            <xs:element name="node" maxOccurs="unbounded">
                                <xs:complexType>
                                    <xs:sequence>
                                        <xs:element name="id" type="xs:string" minOccurs="1" maxOccurs="1"></xs:element>
                                        <xs:element name="from" type="xs:string" minOccurs="1" maxOccurs="1"></xs:element>
                                        <xs:element name="to" type="xs:string" minOccurs="1" maxOccurs="1"></xs:element>
                                        <xs:element name="cost" type="xs:float" minOccurs="0" maxOccurs="1"></xs:element>
                                    </xs:sequence>
                                </xs:complexType>
                            </xs:element>
                        </xs:sequence>
                    </xs:complexType>
                </xs:element>
            </xs:sequence>
        </xs:complexType>
    </xs:element>
</xs:schema>
`

// Validate => validate xml against xsd
func (x *XML) Validate() error {
	xsdSchema, err := xsd.ParseSchema([]byte(xsdContent))
	if err != nil {
		fmt.Println(err)
		return err
	}

	doc := golibxml.ParseDoc(x.Body)
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
