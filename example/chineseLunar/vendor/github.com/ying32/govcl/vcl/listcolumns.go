
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

type TListColumns struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewListColumns(AOwner *TListView) *TListColumns {
    l := new(TListColumns)
    l.instance = ListColumns_Create(CheckPtr(AOwner))
    l.ptr = unsafe.Pointer(l.instance)
    setFinalizer(l, (*TListColumns).Free)
    return l
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsListColumns(obj interface{}) *TListColumns {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TListColumns{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsListColumns.
func ListColumnsFromInst(inst uintptr) *TListColumns {
    return AsListColumns(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsListColumns.
func ListColumnsFromObj(obj IObject) *TListColumns {
    return AsListColumns(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsListColumns.
func ListColumnsFromUnsafePointer(ptr unsafe.Pointer) *TListColumns {
    return AsListColumns(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (l *TListColumns) Free() {
    if l.instance != 0 {
        ListColumns_Free(l.instance)
        l.instance, l.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (l *TListColumns) Instance() uintptr {
    return l.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (l *TListColumns) UnsafeAddr() unsafe.Pointer {
    return l.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (l *TListColumns) IsValid() bool {
    return l.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (l *TListColumns) Is() TIs {
    return TIs(l.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (l *TListColumns) As() TAs {
//    return TAs(l.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TListColumnsClass() TClass {
    return ListColumns_StaticClassType()
}

func (l *TListColumns) Add() *TListColumn {
    return AsListColumn(ListColumns_Add(l.instance))
}

// 组件所有者。
//
// component owner.
func (l *TListColumns) Owner() *TWinControl {
    return AsWinControl(ListColumns_Owner(l.instance))
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (l *TListColumns) Assign(Source IObject) {
    ListColumns_Assign(l.instance, CheckPtr(Source))
}

func (l *TListColumns) BeginUpdate() {
    ListColumns_BeginUpdate(l.instance)
}

// 清除。
func (l *TListColumns) Clear() {
    ListColumns_Clear(l.instance)
}

func (l *TListColumns) Delete(Index int32) {
    ListColumns_Delete(l.instance, Index)
}

func (l *TListColumns) EndUpdate() {
    ListColumns_EndUpdate(l.instance)
}

func (l *TListColumns) FindItemID(ID int32) *TCollectionItem {
    return AsCollectionItem(ListColumns_FindItemID(l.instance, ID))
}

// 获取类名路径。
//
// Get the class name path.
func (l *TListColumns) GetNamePath() string {
    return ListColumns_GetNamePath(l.instance)
}

func (l *TListColumns) Insert(Index int32) *TCollectionItem {
    return AsCollectionItem(ListColumns_Insert(l.instance, Index))
}

// 获取类的类型信息。
//
// Get class type information.
func (l *TListColumns) ClassType() TClass {
    return ListColumns_ClassType(l.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (l *TListColumns) ClassName() string {
    return ListColumns_ClassName(l.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (l *TListColumns) InstanceSize() int32 {
    return ListColumns_InstanceSize(l.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (l *TListColumns) InheritsFrom(AClass TClass) bool {
    return ListColumns_InheritsFrom(l.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (l *TListColumns) Equals(Obj IObject) bool {
    return ListColumns_Equals(l.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (l *TListColumns) GetHashCode() int32 {
    return ListColumns_GetHashCode(l.instance)
}

// 文本类信息。
//
// Text information.
func (l *TListColumns) ToString() string {
    return ListColumns_ToString(l.instance)
}

func (l *TListColumns) Count() int32 {
    return ListColumns_GetCount(l.instance)
}

func (l *TListColumns) Items(Index int32) *TListColumn {
    return AsListColumn(ListColumns_GetItems(l.instance, Index))
}

func (l *TListColumns) SetItems(Index int32, value *TListColumn) {
    ListColumns_SetItems(l.instance, Index, CheckPtr(value))
}

