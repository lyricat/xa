
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    "time"
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TRegistry struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewRegistry(aAccess uint32) *TRegistry {
    r := new(TRegistry)
    r.instance = Registry_Create(uintptr(aAccess))
    r.ptr = unsafe.Pointer(r.instance)
    setFinalizer(r, (*TRegistry).Free)
    return r
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsRegistry(obj interface{}) *TRegistry {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TRegistry{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsRegistry.
func RegistryFromInst(inst uintptr) *TRegistry {
    return AsRegistry(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsRegistry.
func RegistryFromObj(obj IObject) *TRegistry {
    return AsRegistry(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsRegistry.
func RegistryFromUnsafePointer(ptr unsafe.Pointer) *TRegistry {
    return AsRegistry(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (r *TRegistry) Free() {
    if r.instance != 0 {
        Registry_Free(r.instance)
        r.instance, r.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (r *TRegistry) Instance() uintptr {
    return r.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (r *TRegistry) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (r *TRegistry) IsValid() bool {
    return r.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (r *TRegistry) Is() TIs {
    return TIs(r.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (r *TRegistry) As() TAs {
//    return TAs(r.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TRegistryClass() TClass {
    return Registry_StaticClassType()
}

func (r *TRegistry) CloseKey() {
    Registry_CloseKey(r.instance)
}

func (r *TRegistry) CreateKey(Key string) bool {
    return Registry_CreateKey(r.instance, Key)
}

func (r *TRegistry) DeleteKey(Key string) bool {
    return Registry_DeleteKey(r.instance, Key)
}

func (r *TRegistry) DeleteValue(Name string) bool {
    return Registry_DeleteValue(r.instance, Name)
}

func (r *TRegistry) HasSubKeys() bool {
    return Registry_HasSubKeys(r.instance)
}

func (r *TRegistry) KeyExists(Key string) bool {
    return Registry_KeyExists(r.instance, Key)
}

func (r *TRegistry) LoadKey(Key string, FileName string) bool {
    return Registry_LoadKey(r.instance, Key , FileName)
}

func (r *TRegistry) MoveKey(OldName string, NewName string, Delete bool) {
    Registry_MoveKey(r.instance, OldName , NewName , Delete)
}

func (r *TRegistry) OpenKey(Key string, CanCreate bool) bool {
    return Registry_OpenKey(r.instance, Key , CanCreate)
}

func (r *TRegistry) OpenKeyReadOnly(Key string) bool {
    return Registry_OpenKeyReadOnly(r.instance, Key)
}

func (r *TRegistry) ReadBool(Name string) bool {
    return Registry_ReadBool(r.instance, Name)
}

func (r *TRegistry) ReadDate(Name string) time.Time {
    return Registry_ReadDate(r.instance, Name)
}

func (r *TRegistry) ReadDateTime(Name string) time.Time {
    return Registry_ReadDateTime(r.instance, Name)
}

func (r *TRegistry) ReadFloat(Name string) float64 {
    return Registry_ReadFloat(r.instance, Name)
}

func (r *TRegistry) ReadInteger(Name string) int32 {
    return Registry_ReadInteger(r.instance, Name)
}

func (r *TRegistry) ReadString(Name string) string {
    return Registry_ReadString(r.instance, Name)
}

func (r *TRegistry) ReadTime(Name string) time.Time {
    return Registry_ReadTime(r.instance, Name)
}

func (r *TRegistry) RegistryConnect(UNCName string) bool {
    return Registry_RegistryConnect(r.instance, UNCName)
}

func (r *TRegistry) RenameValue(OldName string, NewName string) {
    Registry_RenameValue(r.instance, OldName , NewName)
}

func (r *TRegistry) ReplaceKey(Key string, FileName string, BackUpFileName string) bool {
    return Registry_ReplaceKey(r.instance, Key , FileName , BackUpFileName)
}

func (r *TRegistry) RestoreKey(Key string, FileName string) bool {
    return Registry_RestoreKey(r.instance, Key , FileName)
}

func (r *TRegistry) SaveKey(Key string, FileName string) bool {
    return Registry_SaveKey(r.instance, Key , FileName)
}

func (r *TRegistry) UnLoadKey(Key string) bool {
    return Registry_UnLoadKey(r.instance, Key)
}

func (r *TRegistry) ValueExists(Name string) bool {
    return Registry_ValueExists(r.instance, Name)
}

func (r *TRegistry) WriteBool(Name string, Value bool) {
    Registry_WriteBool(r.instance, Name , Value)
}

func (r *TRegistry) WriteDate(Name string, Value time.Time) {
    Registry_WriteDate(r.instance, Name , Value)
}

func (r *TRegistry) WriteDateTime(Name string, Value time.Time) {
    Registry_WriteDateTime(r.instance, Name , Value)
}

func (r *TRegistry) WriteFloat(Name string, Value float64) {
    Registry_WriteFloat(r.instance, Name , Value)
}

func (r *TRegistry) WriteInteger(Name string, Value int32) {
    Registry_WriteInteger(r.instance, Name , Value)
}

func (r *TRegistry) WriteString(Name string, Value string) {
    Registry_WriteString(r.instance, Name , Value)
}

func (r *TRegistry) WriteExpandString(Name string, Value string) {
    Registry_WriteExpandString(r.instance, Name , Value)
}

func (r *TRegistry) WriteTime(Name string, Value time.Time) {
    Registry_WriteTime(r.instance, Name , Value)
}

// 获取类的类型信息。
//
// Get class type information.
func (r *TRegistry) ClassType() TClass {
    return Registry_ClassType(r.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (r *TRegistry) ClassName() string {
    return Registry_ClassName(r.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (r *TRegistry) InstanceSize() int32 {
    return Registry_InstanceSize(r.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (r *TRegistry) InheritsFrom(AClass TClass) bool {
    return Registry_InheritsFrom(r.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (r *TRegistry) Equals(Obj IObject) bool {
    return Registry_Equals(r.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (r *TRegistry) GetHashCode() int32 {
    return Registry_GetHashCode(r.instance)
}

// 文本类信息。
//
// Text information.
func (r *TRegistry) ToString() string {
    return Registry_ToString(r.instance)
}

func (r *TRegistry) CurrentKey() HKEY {
    return Registry_GetCurrentKey(r.instance)
}

func (r *TRegistry) CurrentPath() string {
    return Registry_GetCurrentPath(r.instance)
}

func (r *TRegistry) LazyWrite() bool {
    return Registry_GetLazyWrite(r.instance)
}

func (r *TRegistry) SetLazyWrite(value bool) {
    Registry_SetLazyWrite(r.instance, value)
}

func (r *TRegistry) LastError() int32 {
    return Registry_GetLastError(r.instance)
}

func (r *TRegistry) LastErrorMsg() string {
    return Registry_GetLastErrorMsg(r.instance)
}

func (r *TRegistry) RootKey() HKEY {
    return Registry_GetRootKey(r.instance)
}

func (r *TRegistry) SetRootKey(value HKEY) {
    Registry_SetRootKey(r.instance, value)
}

func (r *TRegistry) Access() uint32 {
    return Registry_GetAccess(r.instance)
}

func (r *TRegistry) SetAccess(value uint32) {
    Registry_SetAccess(r.instance, value)
}

