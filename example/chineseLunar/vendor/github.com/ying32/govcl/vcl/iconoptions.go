
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TIconOptions struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsIconOptions(obj interface{}) *TIconOptions {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TIconOptions{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsIconOptions.
func IconOptionsFromInst(inst uintptr) *TIconOptions {
    return AsIconOptions(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsIconOptions.
func IconOptionsFromObj(obj IObject) *TIconOptions {
    return AsIconOptions(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsIconOptions.
func IconOptionsFromUnsafePointer(ptr unsafe.Pointer) *TIconOptions {
    return AsIconOptions(ptr)
}

// -------------------------- Deprecated end --------------------------
// 返回对象实例指针。
// 
// Return object instance pointer.
func (i *TIconOptions) Instance() uintptr {
    return i.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (i *TIconOptions) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (i *TIconOptions) IsValid() bool {
    return i.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (i *TIconOptions) Is() TIs {
    return TIs(i.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (i *TIconOptions) As() TAs {
//    return TAs(i.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TIconOptionsClass() TClass {
    return IconOptions_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (i *TIconOptions) Assign(Source IObject) {
    IconOptions_Assign(i.instance, CheckPtr(Source))
}

// 获取类名路径。
//
// Get the class name path.
func (i *TIconOptions) GetNamePath() string {
    return IconOptions_GetNamePath(i.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (i *TIconOptions) ClassType() TClass {
    return IconOptions_ClassType(i.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (i *TIconOptions) ClassName() string {
    return IconOptions_ClassName(i.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (i *TIconOptions) InstanceSize() int32 {
    return IconOptions_InstanceSize(i.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (i *TIconOptions) InheritsFrom(AClass TClass) bool {
    return IconOptions_InheritsFrom(i.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (i *TIconOptions) Equals(Obj IObject) bool {
    return IconOptions_Equals(i.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (i *TIconOptions) GetHashCode() int32 {
    return IconOptions_GetHashCode(i.instance)
}

// 文本类信息。
//
// Text information.
func (i *TIconOptions) ToString() string {
    return IconOptions_ToString(i.instance)
}

func (i *TIconOptions) Arrangement() TIconArrangement {
    return IconOptions_GetArrangement(i.instance)
}

func (i *TIconOptions) SetArrangement(value TIconArrangement) {
    IconOptions_SetArrangement(i.instance, value)
}

func (i *TIconOptions) AutoArrange() bool {
    return IconOptions_GetAutoArrange(i.instance)
}

func (i *TIconOptions) SetAutoArrange(value bool) {
    IconOptions_SetAutoArrange(i.instance, value)
}

