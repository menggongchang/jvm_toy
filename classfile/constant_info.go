package classfile

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
//14中常量，区分不同的常量类型
const (
	CONSTANT_Integer = 3
	CONSTANT_Float   = 4
	CONSTANT_Long    = 5
	CONSTANT_Double  = 6
	CONSTANT_Utf8    = 1
	CONSTANT_String  = 8

	CONSTANT_Class              = 7  //类或接口的符号引用
	CONSTANT_Fieldref           = 9  //字段符号引用
	CONSTANT_Methodref          = 10 //（非接口）方法符号引用
	CONSTANT_InterfaceMethodref = 11 //接口方法符号引用

	CONSTANT_NameAndType   = 12 //字段或方法的名称和描述符
	CONSTANT_MethodType    = 16 //以下三个支持invokeDynamic
	CONSTANT_MethodHandle  = 15
	CONSTANT_InvokeDynamic = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader) //读取数据，初始化过程
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8() //读出tag值
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

//创建具体的常量
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}

	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberRefInfo{cp: cp}}

	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
