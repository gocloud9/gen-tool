package parse_test

import (
	"github.com/gocloud9/gen-tool/pkg/parse"
	"github.com/google/go-cmp/cmp"
	"os"
	"path/filepath"

	"testing"
)

func TestParser_ParseDirectory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *parse.Results
		wantErr bool
	}{
		{
			name: "simple struct",
			args: args{
				path: "./_testdata/simple",
			},
			want: &parse.Results{
				Packages: map[string]*parse.PackageInfo{
					"simple": {
						Name:         "simple",
						Constants:    map[string]*parse.ConstantInfo{},
						Functions:    map[string]*parse.FuncInfo{},
						Interfaces:   map[string]*parse.InterfaceInfo{},
						Vars:         map[string]*parse.VarInfo{},
						DefinedTypes: map[string]*parse.DefinedTypeInfo{},
						Aliases:      map[string]*parse.AliasTypeInfo{},
						Structs: map[string]*parse.StructInfo{
							"User": {
								Name: "User",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								Fields: map[string]*parse.FieldInfo{
									"ID": {
										Name: "ID",
										Tags: map[string][]string{
											"json": {"id"},
										},
										TypeInfo: &parse.TypeInfo{
											TypeName: "string",
										},
										Markers: map[string]string{
											"+something:id": "true",
										},
									},
									"DisplayName": {
										Name: "DisplayName",
										Tags: map[string][]string{
											"json": {"display_name"},
										},
										TypeInfo: &parse.TypeInfo{
											TypeName:  "*string",
											IsPointer: true,
											Pointer: &parse.TypeInfo{
												TypeName: "string",
											},
										},
										Markers: map[string]string{},
									},
									"Email": {
										Name: "Email",
										Tags: map[string][]string{
											"json": {"email"},
										},
										TypeInfo: &parse.TypeInfo{
											TypeName: "string",
										},
										Markers: map[string]string{},
									},
									"Age": {
										Name: "Age",
										Tags: map[string][]string{
											"json": {"age"},
										},
										TypeInfo: &parse.TypeInfo{
											TypeName: "int",
										},
										Markers: map[string]string{},
									},
								},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
								Methods:        map[string]*parse.FuncInfo{},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "permutations with multiple packages",
			args: args{
				path: "./_testdata/permutation",
			},
			want: &parse.Results{
				Packages: map[string]*parse.PackageInfo{
					"package1": {
						Name:         "package1",
						DefinedTypes: map[string]*parse.DefinedTypeInfo{},
						Aliases:      map[string]*parse.AliasTypeInfo{},
						Structs: map[string]*parse.StructInfo{
							"AnotherUser": {
								Name: "AnotherUser",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								Fields: map[string]*parse.FieldInfo{
									"ID": {
										Name: "ID",
										Tags: map[string][]string{
											"json": {"id"},
										},
										TypeInfo: &parse.TypeInfo{
											TypeName: "string",
										},
										Markers: map[string]string{"+ID": "true"},
									},
									"DisplayName": {
										Name: "DisplayName",
										Tags: map[string][]string{
											"json": {"display_name"},
										},
										TypeInfo: &parse.TypeInfo{
											TypeName: "string",
										},
										Markers: map[string]string{},
									},
									"Duration": {
										Name:    "Duration",
										Markers: map[string]string{},
										Tags:    map[string][]string{"json": {"duration"}},
										TypeInfo: &parse.TypeInfo{
											TypeName: "time.Duration",
											ImportedType: &parse.ImportedTypeInfo{
												TypeName:            "Duration",
												ImportRaw:           "\"time\"",
												PackagePath:         "time",
												PackageDefaultAlias: "time",
											},
										},
									},
									"Email": {
										Name: "Email",
										Tags: map[string][]string{
											"json": {"email"},
										},
										TypeInfo: &parse.TypeInfo{
											TypeName: "string",
										},
										Markers: map[string]string{},
									},
									"Time": {
										Name:    "Time",
										Markers: map[string]string{},
										Tags:    map[string][]string{"json": {"time"}},
										TypeInfo: &parse.TypeInfo{
											TypeName: "time.Time",
											ImportedType: &parse.ImportedTypeInfo{
												TypeName:            "Time",
												ImportRaw:           "\"time\"",
												PackagePath:         "time",
												PackageDefaultAlias: "time",
											},
										},
									},
									"Timestamp": {
										Name:    "Timestamp",
										Markers: map[string]string{},
										Tags:    map[string][]string{"json": {"timestamp"}},
										TypeInfo: &parse.TypeInfo{
											TypeName: "timestamppb.Timestamp",
											ImportedType: &parse.ImportedTypeInfo{
												TypeName:            "Timestamp",
												ImportRaw:           "\"google.golang.org/protobuf/types/known/timestamppb\"",
												PackagePath:         "google.golang.org/protobuf/types/known/timestamppb",
												PackageDefaultAlias: "timestamppb",
											},
										},
									},
								},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
								Methods:        map[string]*parse.FuncInfo{},
							},
						},
						Constants:  map[string]*parse.ConstantInfo{},
						Functions:  map[string]*parse.FuncInfo{},
						Interfaces: map[string]*parse.InterfaceInfo{},
						Vars:       map[string]*parse.VarInfo{},
					},
					"package2": {
						Name:         "package2",
						DefinedTypes: map[string]*parse.DefinedTypeInfo{},
						Aliases:      map[string]*parse.AliasTypeInfo{},
						Structs: map[string]*parse.StructInfo{
							"SomeStruct": {
								Name:    "SomeStruct",
								Markers: map[string]string{},
								Fields: map[string]*parse.FieldInfo{
									"StringField": {
										Name: "StringField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "string",
										},
										Tags:    map[string][]string{"json": {"string_field"}, "yaml": {"stringField"}},
										Markers: map[string]string{},
									},
									"IntField": {
										Name: "IntField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "int",
										},
										Tags:    map[string][]string{"json": {"int_field"}, "yaml": {"intField"}},
										Markers: map[string]string{},
									},
									"BoolField": {
										Name: "BoolField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "bool",
										},
										Tags:    map[string][]string{"json": {"bool_field"}, "yaml": {"boolField"}},
										Markers: map[string]string{},
									},
									"ChanField": {
										Name: "ChanField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "chan int",
											IsChan:   true,
											Chan: &parse.TypeInfo{
												TypeName: "int",
											},
										},
										Tags:    map[string][]string{"json": {"chan_field"}, "yaml": {"chanField"}},
										Markers: map[string]string{},
									},
									"MapField": {
										Name: "MapField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "map[string]int",
											IsMap:    true,
											MapKey: &parse.TypeInfo{
												TypeName: "string",
											},
											MapValue: &parse.TypeInfo{
												TypeName: "int",
											},
										},
										Tags:    map[string][]string{"json": {"map_field"}, "yaml": {"mapField"}},
										Markers: map[string]string{},
									},
									"SliceField": {
										Name: "SliceField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "[]int",
											IsSlice:  true,
											Slice: &parse.TypeInfo{
												TypeName: "int",
											},
										},
										Tags:    map[string][]string{"json": {"slice_field"}, "yaml": {"sliceField"}},
										Markers: map[string]string{},
									},
									"SubStructField": {
										Name: "SubStructField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "SubStruct",
											IsStruct: true,
										},
										Tags:    map[string][]string{"json": {"sub_struct_field"}, "yaml": {"subStructField"}},
										Markers: map[string]string{},
									},
									"SubStructMapField": {
										Name: "SubStructMapField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "map[string]SubStruct",
											IsMap:    true,
											MapKey: &parse.TypeInfo{
												TypeName: "string",
											},
											MapValue: &parse.TypeInfo{
												TypeName: "SubStruct",
												IsStruct: true,
											},
										},
										Tags:    map[string][]string{"json": {"sub_struct_map_field"}, "yaml": {"subStructMapField"}},
										Markers: map[string]string{},
									},
									"SubStructSliceField": {
										Name: "SubStructSliceField",
										TypeInfo: &parse.TypeInfo{
											TypeName: "[]SubStruct",
											IsSlice:  true,
											Slice: &parse.TypeInfo{
												TypeName: "SubStruct",
												IsStruct: true,
											},
										},
										Tags:    map[string][]string{"json": {"sub_struct_slice_field"}, "yaml": {"subStructSliceField"}},
										Markers: map[string]string{},
									},
								},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
								Methods:        map[string]*parse.FuncInfo{},
							},
							"SubStruct": {
								Name: "SubStruct", Markers: map[string]string{}, Fields: map[string]*parse.FieldInfo{
									"Something": {
										Name: "Something",
										TypeInfo: &parse.TypeInfo{
											TypeName: "string",
										},
										Tags:    map[string][]string{},
										Markers: map[string]string{},
									},
								},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
								Methods:        map[string]*parse.FuncInfo{},
							},
						},
						Constants:  map[string]*parse.ConstantInfo{},
						Functions:  map[string]*parse.FuncInfo{},
						Interfaces: map[string]*parse.InterfaceInfo{},
						Vars:       map[string]*parse.VarInfo{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "functions",
			args: args{
				path: "./_testdata/functions",
			},
			want: &parse.Results{
				Packages: map[string]*parse.PackageInfo{
					"functions": {
						Name:         "functions",
						Constants:    map[string]*parse.ConstantInfo{},
						DefinedTypes: map[string]*parse.DefinedTypeInfo{},
						Aliases:      map[string]*parse.AliasTypeInfo{},
						Structs: map[string]*parse.StructInfo{
							"Field": {
								Name: "Field",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								Fields:         map[string]*parse.FieldInfo{},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
								Methods: map[string]*parse.FuncInfo{
									"Test5": {
										Name: "Test5",
										Markers: map[string]string{
											"+Foo": "true",
											"+Bar": "123",
										},
										FuncDefInfo: &parse.FuncDefInfo{
											Params: []*parse.ParamInfo{
												{
													Name: "arg",
													TypeInfo: &parse.TypeInfo{
														TypeName: "Field",
														IsStruct: true,
													},
												},
											},
											Results: []*parse.ResultInfo{},
										},
										HasReciver:  true,
										ReciverName: "Field",
									},
								},
							},
							"ImportedTypeInfo": {
								Name: "ImportedTypeInfo",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								Methods:        map[string]*parse.FuncInfo{},
								Fields:         map[string]*parse.FieldInfo{},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
							},
						},
						Interfaces: map[string]*parse.InterfaceInfo{},
						Vars: map[string]*parse.VarInfo{
							"myFunc": {
								Name: "myFunc",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								TypeInfo: &parse.TypeInfo{
									TypeName: "func()",
									IsFunc:   true,
									Func: &parse.FuncDefInfo{
										Params:  []*parse.ParamInfo{},
										Results: []*parse.ResultInfo{},
									},
								},
							},
							"myGroupedFunc": {
								Name: "myGroupedFunc",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								TypeInfo: &parse.TypeInfo{
									TypeName: "func()",
									IsFunc:   true,
									Func: &parse.FuncDefInfo{
										Params:  []*parse.ParamInfo{},
										Results: []*parse.ResultInfo{},
									},
								},
							},
						},
						Functions: map[string]*parse.FuncInfo{
							"Test1": {
								Name:    "Test1",
								Markers: map[string]string{"+Bar": "123", "+Foo": "true"},
								FuncDefInfo: &parse.FuncDefInfo{
									IsVariadic: false,
									Params:     []*parse.ParamInfo{},
									Results:    []*parse.ResultInfo{},
								},
							},
							"Test2": {
								Name:    "Test2",
								Markers: map[string]string{"+Bar": "123", "+Foo": "true"},
								FuncDefInfo: &parse.FuncDefInfo{
									IsVariadic: false,
									Params:     []*parse.ParamInfo{},
									Results: []*parse.ResultInfo{
										{
											TypeInfo: &parse.TypeInfo{
												TypeName: "error",
											},
										},
									},
								},
							},
							"Test3": {
								Name: "Test3",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								FuncDefInfo: &parse.FuncDefInfo{
									Params: []*parse.ParamInfo{
										{
											Name: "arg",
											TypeInfo: &parse.TypeInfo{
												TypeName: "string",
											},
										},
									},
									Results: []*parse.ResultInfo{},
								},
							},
							"Test4": {
								Name: "Test4",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								FuncDefInfo: &parse.FuncDefInfo{
									Params: []*parse.ParamInfo{
										{
											Name: "arg",
											TypeInfo: &parse.TypeInfo{
												TypeName: "Field",
												IsStruct: true,
											},
										},
									},
									Results: []*parse.ResultInfo{},
								},
							},
							"Test5": {
								Name: "Test5",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								FuncDefInfo: &parse.FuncDefInfo{
									Params: []*parse.ParamInfo{
										{
											Name: "arg",
											TypeInfo: &parse.TypeInfo{
												TypeName: "Field",
												IsStruct: true,
											},
										},
									},
									Results: []*parse.ResultInfo{},
								},
								HasReciver:  true,
								ReciverName: "Field",
							},
							"Test6": {
								Name: "Test6",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								FuncDefInfo: &parse.FuncDefInfo{
									Params: []*parse.ParamInfo{
										{
											Name: "arg",
											TypeInfo: &parse.TypeInfo{
												TypeName: "Field",
												IsStruct: true,
											},
										},
									},
									Results: []*parse.ResultInfo{},
								},
							},
							"Variadic": {
								Name:    "Variadic",
								Markers: map[string]string{},
								FuncDefInfo: &parse.FuncDefInfo{
									IsVariadic: true,
									Params: []*parse.ParamInfo{
										{
											Name: "vArg",
											TypeInfo: &parse.TypeInfo{
												TypeName:   "...Field",
												IsEllipsis: true,
												Ellipsis: &parse.TypeInfo{
													TypeName: "Field",
													IsStruct: true,
												},
											},
										},
									},
									Results: []*parse.ResultInfo{},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "interfaces",
			args: args{
				path: "./_testdata/interfaces",
			},
			want: &parse.Results{
				Packages: map[string]*parse.PackageInfo{
					"interfaces": {
						Name:         "interfaces",
						Constants:    map[string]*parse.ConstantInfo{},
						DefinedTypes: map[string]*parse.DefinedTypeInfo{},
						Aliases:      map[string]*parse.AliasTypeInfo{},
						Structs: map[string]*parse.StructInfo{
							"TestStruct": {
								Name:           "TestStruct",
								Markers:        map[string]string{},
								Fields:         map[string]*parse.FieldInfo{},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
								Methods:        map[string]*parse.FuncInfo{},
							},
						},
						Vars:      map[string]*parse.VarInfo{},
						Functions: map[string]*parse.FuncInfo{},
						Interfaces: map[string]*parse.InterfaceInfo{
							"MyInterface": {
								Name:    "MyInterface",
								Markers: map[string]string{"+Bar": "123", "+Foo": "true"},
								Methods: map[string]*parse.FuncInfo{
									"DoSomething": {
										Name:    "DoSomething",
										Markers: map[string]string{"+Bar": "123", "+Foo": "true"},
										FuncDefInfo: &parse.FuncDefInfo{
											IsVariadic: false,
											Params: []*parse.ParamInfo{
												{
													Name: "input",
													TypeInfo: &parse.TypeInfo{
														TypeName:  "*string",
														IsPointer: true,
														Pointer: &parse.TypeInfo{
															TypeName: "string",
														},
													},
												},
												{
													Name: "f",
													TypeInfo: &parse.TypeInfo{
														TypeName: "func([]TestInterface) map[string]TestStruct",
														IsFunc:   true,
														Func: &parse.FuncDefInfo{
															Params: []*parse.ParamInfo{
																{
																	Name: "",
																	TypeInfo: &parse.TypeInfo{
																		TypeName: "[]TestInterface",
																		IsSlice:  true,
																		Slice: &parse.TypeInfo{
																			TypeName:    "TestInterface",
																			IsInterface: true,
																		},
																	},
																},
															},
															Results: []*parse.ResultInfo{
																{
																	TypeInfo: &parse.TypeInfo{
																		TypeName: "map[string]TestStruct",
																		IsMap:    true,
																		MapKey: &parse.TypeInfo{
																			TypeName: "string",
																		},
																		MapValue: &parse.TypeInfo{
																			TypeName: "TestStruct",
																			IsStruct: true,
																		},
																	},
																},
															},
														},
													},
												},
											},
											Results: []*parse.ResultInfo{
												{
													Name: "output",
													TypeInfo: &parse.TypeInfo{
														TypeName: "string",
													},
												},
												{
													Name: "err",
													TypeInfo: &parse.TypeInfo{
														TypeName: "error",
													},
												},
											},
										},
									},
								},
								EmbeddedTypes: map[string]*parse.EmbeddedTypeInfo{},
							},
							"TestInterface": {
								Name:          "TestInterface",
								Markers:       map[string]string{},
								Methods:       map[string]*parse.FuncInfo{},
								EmbeddedTypes: map[string]*parse.EmbeddedTypeInfo{},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "globals",
			args: args{
				path: "./_testdata/globals",
			},
			want: &parse.Results{
				Packages: map[string]*parse.PackageInfo{
					"globals": {
						Name:    "globals",
						Structs: map[string]*parse.StructInfo{},
						DefinedTypes: map[string]*parse.DefinedTypeInfo{
							"MyString": {
								Name:    "MyString",
								Markers: map[string]string{},
								TypeInfo: &parse.TypeInfo{
									TypeName: "string",
								},
							},
						},
						Aliases: map[string]*parse.AliasTypeInfo{},
						Constants: map[string]*parse.ConstantInfo{
							"myConstant": {
								Name: "myConstant",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								TypeName: "string",
								Value:    `"test"`,
							},
							"myStringType": {
								Name:     "myStringType",
								Markers:  map[string]string{},
								TypeName: "MyString",
								Value:    `"test"`,
							},
						},
						Vars: map[string]*parse.VarInfo{
							"myFunc": {
								Name: "myFunc",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								TypeInfo: &parse.TypeInfo{
									TypeName: "func(string) error",
									IsFunc:   true,
									Func: &parse.FuncDefInfo{
										Params: []*parse.ParamInfo{
											{
												Name: "arg",
												TypeInfo: &parse.TypeInfo{
													TypeName: "string",
												},
											},
										},
										Results: []*parse.ResultInfo{
											{
												TypeInfo: &parse.TypeInfo{
													TypeName: "error",
												},
											},
										},
									},
								},
							},
						},
						Functions:  map[string]*parse.FuncInfo{},
						Interfaces: map[string]*parse.InterfaceInfo{},
					},
				},
			},
		},
		{
			name: "embedded",
			args: args{
				path: "./_testdata/embedded",
			},
			want: &parse.Results{
				Packages: map[string]*parse.PackageInfo{
					"embedded": {
						Name:      "embedded",
						Constants: map[string]*parse.ConstantInfo{},
						Vars:      map[string]*parse.VarInfo{},
						Functions: map[string]*parse.FuncInfo{},
						DefinedTypes: map[string]*parse.DefinedTypeInfo{
							"ParentStruct": {
								Name:    "ParentStruct",
								Markers: map[string]string{},
								TypeInfo: &parse.TypeInfo{
									TypeName: "func()",
									IsFunc:   true,
									Func: &parse.FuncDefInfo{
										Params:  []*parse.ParamInfo{},
										Results: []*parse.ResultInfo{},
									},
								},
							},
						},
						Aliases: map[string]*parse.AliasTypeInfo{},
						Interfaces: map[string]*parse.InterfaceInfo{
							"ParentInterface": {
								Name:          "ParentInterface",
								Markers:       map[string]string{},
								Methods:       map[string]*parse.FuncInfo{},
								EmbeddedTypes: map[string]*parse.EmbeddedTypeInfo{},
							},
							"ChildInterface": {
								Name: "ChildInterface",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								Methods: map[string]*parse.FuncInfo{},
								EmbeddedTypes: map[string]*parse.EmbeddedTypeInfo{
									"ParentInterface": {
										Name:     "ParentInterface",
										TypeName: "ParentInterface",
										Markers:  map[string]string{"+Bar": "123", "+Foo": "true"},
									},
									"Parent": {
										Name:     "Parent",
										TypeName: "Parent",
										Markers:  map[string]string{"+Bar": "123", "+Foo": "true"},
									},
								},
							},
						},
						Structs: map[string]*parse.StructInfo{
							"Child": {
								Name: "Child",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								Fields: map[string]*parse.FieldInfo{},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{
									"ParentInterface": {
										Name:     "ParentInterface",
										TypeName: "ParentInterface",
										Markers:  map[string]string{"+Bar": "123", "+Foo": "true"},
										Tags:     map[string][]string{"yaml": {"", "inline"}},
									},
									"Parent": {
										Name:     "Parent",
										TypeName: "Parent",
										Markers:  map[string]string{"+Bar": "123", "+Foo": "true"},
										Tags:     map[string][]string{"yaml": {"", "inline"}},
									},
								},
								Methods: map[string]*parse.FuncInfo{},
							},
							"Parent": {
								Name:           "Parent",
								Markers:        map[string]string{},
								Fields:         map[string]*parse.FieldInfo{},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
								Methods:        map[string]*parse.FuncInfo{},
							},
						},
					},
				},
			},
		},
		{
			name: "typing",
			args: args{
				path: "./_testdata/typing",
			},
			want: &parse.Results{
				Packages: map[string]*parse.PackageInfo{
					"typing": {
						Name: "typing",
						Structs: map[string]*parse.StructInfo{
							"AStruct": {
								Name: "AStruct",
								Markers: map[string]string{
									"+Foo": "true",
									"+Bar": "123",
								},
								Fields:         map[string]*parse.FieldInfo{},
								EmbeddedFields: map[string]parse.EmbeddedFieldInfo{},
								Methods:        map[string]*parse.FuncInfo{},
							},
						},
						Constants:  map[string]*parse.ConstantInfo{},
						Functions:  map[string]*parse.FuncInfo{},
						Interfaces: map[string]*parse.InterfaceInfo{},
						Vars:       map[string]*parse.VarInfo{},
						DefinedTypes: map[string]*parse.DefinedTypeInfo{
							"StringType": {
								Name:    "StringType",
								Markers: map[string]string{"+Bar": "123", "+Foo": "true"},
								TypeInfo: &parse.TypeInfo{
									TypeName: "string",
								},
							},
							"AStructType": {
								Name:     "AStructType",
								Markers:  map[string]string{"+Bar": "123", "+Foo": "true"},
								TypeInfo: &parse.TypeInfo{TypeName: "AStruct", IsStruct: true},
							},
							"IntType": {
								Name:     "IntType",
								Markers:  map[string]string{"+Bar": "123", "+Foo": "true"},
								TypeInfo: &parse.TypeInfo{TypeName: "int"},
							},
							"SliceType": {
								Name:    "SliceType",
								Markers: map[string]string{"+Bar": "123", "+Foo": "true"},
								TypeInfo: &parse.TypeInfo{
									TypeName: "[]AStruct",
									IsSlice:  true,
									Slice:    &parse.TypeInfo{TypeName: "AStruct", IsStruct: true},
								},
							},
						},
						Aliases: map[string]*parse.AliasTypeInfo{
							"AliasStringType": {
								Name:     "AliasStringType",
								Markers:  map[string]string{"+Bar": "123", "+Foo": "true"},
								TypeInfo: &parse.TypeInfo{TypeName: "string"},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parse.Parser{}
			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("failed to get working directory: %v", err)
			}

			path := filepath.Join(wd, tt.args.path)

			got, err := p.ParseDirectory(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
