package main

type ParseProtocolData struct {
	Name *string `"protocol" @Ident`
}

type ParseEnumValue struct {
	Name  *string `@Ident`
	Value *int    `@Int`
}

type ParseEnum struct {
	Name   *string           `"enum" @Ident "{"`
	Values []*ParseEnumValue `{ @@ } "}"`
}

type ParseRequest struct {
	Name      *string          `"request" @Ident "("`
	Arguments []*ParseArgument `{ @@ [","] } ")"`
	Since     *int             `[ "since" @Int ]`
	Summary   *string          `[ @Comment ]`
}

type ParseEvent struct {
	Name      *string          `"event" @Ident "("`
	Arguments []*ParseArgument `{ @@ [","] } ")"`
	Since     *int             `[ "since" @Int ]`
}

type ParseDualType struct {
	Base *string `@Ident`
	Sub  *string `"["@Ident"]"`
}

type ParseType struct {
	DualType   *ParseDualType `( @@`
	SingleType *string        `| @Ident )`
}

type ParseArgument struct {
	Name *string    `@Ident`
	Type *ParseType `@@`
}

type ParseCopyright struct {
	Data *string `"copyright" @RawString`
}

type ParseInterface struct {
	Name        *string           `"interface" @Ident "{"`
	Version     *int              `"version" @Int`
	Description *ParseDescription `{ @@ }`
	Enums       []*ParseEnum      `{ @@ }`
	Request     []*ParseRequest   `{ @@ }`
	Event       []*ParseEvent     `{ @@ }`
	Ending      struct{}          `"}"`
}

type ParseDescription struct {
	Summary     *string `"description" @String`
	Description *string `@RawString`
}

type ProtocolDescription struct {
	ProtocolVersion     *ParseProtocolData `{ @@ }`
	ProtocolCopyright   *ParseCopyright    `{ @@ }`
	ProtocolDescription *ParseDescription  `{ @@ }`
	Interface           []*ParseInterface  `{ @@ }`
}
