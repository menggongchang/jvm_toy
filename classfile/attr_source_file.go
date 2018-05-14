package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	courceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.courceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.courceFileIndex)
}
