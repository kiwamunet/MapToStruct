# MapToStruct



result

```
main.go:65: map is  map[P_B:1 P_C:true CA_A:ChildStruct CA_A string G_A:GrandsonStruct G_A string P_F:432 P_H:56 CB_A:*ChildStructB CB_A string P_A:ParentStruct P_A string]
main.go:70: [before] struct is  &{ 0 false { 0 false { 0 false}} <nil> <nil> <nil> <nil>}
main.go:75: [after] struct is  &{ParentStruct P_A string 1 true {ChildStruct CA_A string 0 false {GrandsonStruct G_A string 0 false}} 0xc4200ac480 0xc4200ac528 0xc4200ac430 value:56 }
main.go:76: pointer struct is  ParentStruct P_E *string 432 {56}
main.go:77: [[        Success        ]]
```