# Graph

## Sample input

This is input sample which will process by `submit_handler`.

``` xml

<graph>
    <id>123</id>
    <name>The Graph Name</name>
    <nodes>
        <node>
            <id>1</id>
            <name>1</name>
        </node>
        <node>
            <id>2</id>
            <name>2</name>
        </node>
        <node>
            <id>3</id>
            <name>3</name>
        </node>
        <node>
            <id>4</id>
            <name>4</name>
        </node>
        <node>
            <id>5</id>
            <name>5</name>
        </node>
        <node>
            <id>6</id>
            <name>6</name>
        </node>
    </nodes>
    <edges>
        <node>
            <id>1</id>
            <from>1</from>
            <to>2</to>
            <cost>1</cost>
        </node>
        <node>
            <id>2</id>
            <from>2</from>
            <to>3</to>
            <cost>1</cost>
        </node>
        <node>
            <id>3</id>
            <from>1</from>
            <to>3</to>
            <cost>1</cost>
        </node>
        <node>
            <id>4</id>
            <from>4</from>
            <to>5</to>
            <cost>3</cost>
        </node>
        <node>
            <id>5</id>
            <from>3</from>
            <to>6</to>
            <cost>1</cost>
        </node>
        <node>
            <id>6</id>
            <from>6</from>
            <to>5</to>
            <cost>1</cost>
        </node>
    </edges>
</graph>
```

This is output model of `find` endpoint

```json
{
    "answers": [
        {
            "paths": {
                "from": 1,
                "to": 6,
                "path": [
                    [
                        1,
                        3,
                        6
                    ],
                    [
                        1,
                        2,
                        3,
                        6
                    ],
                    null
                ]
            },
            "cheapest": {
                "from": 1,
                "to": 6,
                "path": null
            }
        }
    ]
}
```

The `XSD` to validate input XML

```xsd
<?xml version="1.0" encoding="UTF-8"?>
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
```

## Endpoints

### Create new graph in Database

``` curl
curl -X POST \
  http://localhost:9999/ \
  -H 'Content-Type: application/xml' \
  -H 'cache-control: no-cache' \
  -d '
<graph>
    <id>123</id>
    <name>The Graph Name</name>
    <nodes>
        <node>
            <id>1</id>
            <name>1</name>
        </node>
        <node>
            <id>2</id>
            <name>2</name>
        </node>
        <node>
            <id>3</id>
            <name>3</name>
        </node>
        <node>
            <id>4</id>
            <name>4</name>
        </node>
        <node>
            <id>5</id>
            <name>5</name>
        </node>
        <node>
            <id>6</id>
            <name>6</name>
        </node>
    </nodes>
    <edges>
        <node>
            <id>1</id>
            <from>1</from>
            <to>2</to>
            <cost>1</cost>
        </node>
        <node>
            <id>2</id>
            <from>2</from>
            <to>3</to>
            <cost>1</cost>
        </node>
        <node>
            <id>3</id>
            <from>1</from>
            <to>3</to>
            <cost>1</cost>
        </node>
        <node>
            <id>4</id>
            <from>4</from>
            <to>5</to>
            <cost>3</cost>
        </node>
        <node>
            <id>5</id>
            <from>3</from>
            <to>6</to>
            <cost>1</cost>
        </node>
        <node>
            <id>6</id>
            <from>6</from>
            <to>5</to>
            <cost>1</cost>
        </node>
    </edges>
</graph>
```


### Get every possible path from __point_1__ to __point_2__

```curl
curl -X GET \
  'http://localhost:9999/find?graphId=123' \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
    "queries": [
        {
            "paths": {
                "start": 1,
                "end": 6
            },
            "cheapest": {
                "start": 1,
                "end": 6
            }
        }
    ]
}
```

## Build

## ENV Variables

| Name                     | Default Value |
|--------------------------|---------------|
| UNIREG_DEBUG             | false         |
| UNIREG_PORT              | 9999          |
| UNIREG_DATABASE_HOST     | 127.0.0.1     |
| UNIREG_DATABASE_PORT     | 5432          |
| UNIREG_DATABASE_PASSWORD | 123456        |
| UNIREG_DATABASE_NAME     | unireg        |

### From source

Do following steps:

- git clone `git@github.com:devmark88/graph.git`
- cd `./graph`
- go build (it will install dependencies)
- ./unireg

### Docker

simply run `docker-compose up` if you have docker installed on your system

## Want more?

You can download the `Postman Collection` here => https://www.getpostman.com/collections/21b21083d5766e3138f4

## TODO

- [ ] Add test
- [ ] Add `CI/CD`
- [ ] Add Deployment config for `Kubernetes`