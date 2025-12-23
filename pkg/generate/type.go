package generate

const (
	Global             Type = "global"
	PerPackage         Type = "per-package"
	PerStruct          Type = "per-struct"
	PerStructMethod    Type = "per-struct-method"
	PerStructField     Type = "per-struct-field"
	PerInterfaceMethod Type = "per-interface-method"
	PerInterface       Type = "per-interface"
	PerVar             Type = "per-var"
	PerConstant        Type = "per-constant"
	PerFunc            Type = "per-func"
	PerDefinedType     Type = "per-defined-type"
	PerAlias           Type = "per-alias"
)

type Type string
