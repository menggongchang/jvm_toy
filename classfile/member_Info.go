package classfile

/*
字段结构定义
field_info {
    u2             access_flags;//访问标志
    u2             name_index; //常量池索引
    u2             descriptor_index;//字段描述符
    u2             attributes_count;//属性表
    attribute_info attributes[attributes_count];
}
方法结构定义
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	n := reader.readUint16()
	members := make([]*MemberInfo, n)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	memberInfo := &MemberInfo{}
	memberInfo.cp = cp
	memberInfo.accessFlags = reader.readUint16()
	memberInfo.nameIndex = reader.readUint16()
	memberInfo.descriptorIndex = reader.readUint16()
	memberInfo.attributes = readAttributes(reader, cp) //3.4
	return memberInfo
}

//Getter
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}
