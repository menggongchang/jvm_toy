# jvm_toy
just a toy of jvm

----------------------------

提交版本说明

- V2：解决了java虚拟机从哪里搜索class文件，并且实现了类路径的功能，可以把class文件读取到内存中；
- V3：详细讨论class文件格式，解析class文件；



---------------------

##### 一、搜索class文件

1. Oracle的JVM实现根据**类路径**来搜索类，按照搜索顺序，类路径分类：

   - 启动类路径（bootClassPath）：默认jre\lib目录，java标准库位置；
   - 扩展类路径（extClassPath）：默认jre\lib\ext目录，使用java扩展机制的类位置；
   - 用户类路径（userClassPath）：用户实现的类或第三方类库位置；

2. 类路径可以是目录，jar包或zip包；

   ![类路径](D:\workspace\go\src\jvmgo\resource\类路径.png)

   bootClassPath和extClassPath，利用通配符得到目录下的jar包集合CompositeEntry；

   userClassPath是一个entry，可能是目录，CompositeEntry。。；

   ​

#####二、解析class文件

![classfile结构](C:\Users\梦工厂\Downloads\classfile结构.png)

1. 虚拟机规范的class文件结构

   ```java
   ClassFile {
       u4             magic; //魔数，用于标识class文件格式，固定4字节0xCAFEBABE
       u2             minor_version; //次版本号
       u2             major_version; //主版本号
       u2             constant_pool_count; 
       cp_info        constant_pool[constant_pool_count-1];
       u2             access_flags; //类访问标志，指出类还是接口、访问级别是public还是private；
       u2             this_class; //类名索引，常量池中存字符串内容
       u2             super_class; //超类名索引，常量池中存字符串内容
       u2             interfaces_count; //接口索引表，所有的接口名字
       u2             interfaces[interfaces_count];
       u2             fields_count; //字段表
       field_info     fields[fields_count];
       u2             methods_count; //方法表
       method_info    methods[methods_count];
       u2             attributes_count; //属性表
       attribute_info attributes[attributes_count];
   }
   ```

   class文件的基本数据单位是字节，作为一个字节流来处理。

   - u1,u2,u4 是jvm规范定义的三种数据类型来表示1,2,4 字节无符号整数；
   - 对于多个连续字节构成的数据，在class文件中以**大端形式**存储；
   - 同类型的多个数据以表的形式存储，表头n+表项(n个)；

   ​

2. 解析class文件

   - **魔数**：文件开头的几个字节，用于标识文件格式。

     class文件4字节0xCAFEBABE

   - **字段信息**：`public static final char X = 'X';`   *字段指的就是变量X*

     ```java
     //字段结构定义
     field_info {
         u2             access_flags; //访问标志 public static final 
         u2             name_index;  //字段名索引 X
         u2             descriptor_index; //字段描述符索引 C->char 
         u2             attributes_count;//属性表:字段值的信息
         attribute_info attributes[attributes_count];
     }
     ```

   - **方法信息**

     ```java
     method_info {
         u2             access_flags; //访问标志 
         u2             name_index;  //方法名索引
         u2             descriptor_index; //方法描述符索引
         u2             attributes_count; //方法表
         attribute_info attributes[attributes_count];
     }
     ```

   - **常量池 ConstantPool**

     ```java
     //常量信息
     cp_info{
         u1 tag; //区分常量类型
         u1 info[];
     }
     ```

     常量池中的常量分为两类：

     - 字面量：数字字面量+字符串字面量，直接存数据；
     - 符号引用：类和接口名+字段和方法信息，通过索引直接或间接指向CONSTANT_Utf8常量；

     虚拟机规范一共定义了14种常量，利用tag区分：

     - `CONSTANT_Utf8`：字符串在class文件中是以MUTF-8方式编码，而不是UTF-8；

     - `CONSTANT_String`：表示java.lang.String字面量；

       `System.out.println("Hello, World!");` 

       CONSTANT_String并不存放字符串，只是CONSTANT_Utf8索引；

     - `CONSTANT_NameAndType`：字段或方法的名称和描述符；

       字段或方法的名字就是代码中出现的（或者编译器生成的）字段或方法的名字；

       字段描述符就是**字段类型**的描述符； `short -> S ，int[] -> [I`

       方法描述符是（分号分割的参数类型描述符）+返回值**类型描述符**；

       `int binarySearch(long[] a, long key) -> ([JJ)I` 

   - **属性表 AttributeInfo** 

     ```java
     attribute_info {
         u2 attribute_name_index; //利用名字来区分
         u4 attribute_length; 
         u1 info[attribute_length];
     }
     ```







