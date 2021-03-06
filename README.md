# jvm_toy
just a toy of jvm

----------------------------

提交版本说明

- v2.0：解决了java虚拟机从哪里搜索class文件，并且实现了类路径的功能，可以把class文件读取到内存中；
- v3.0：详细讨论class文件格式，解析class文件；
- v4.0：初步实现运行时数据区，主要是线程私有的部分；
- v5.0：编写一个简单的解释器，并实现大约150条指令；
- v6.0：实现线程共享的运行时数据区，包括方法区和运行时常量池；
- v7.0：实现方法的调用和返回，并且实现类的初始化逻辑；
- v8.0：实现数组和字符串，(拿来主义);


---------------------


#####  一、搜索class文件

1. Oracle的JVM实现根据**类路径**来搜索类，按照搜索顺序，类路径分类：

   - 启动类路径（bootClassPath）：默认jre\lib目录，java标准库位置；
   - 扩展类路径（extClassPath）：默认jre\lib\ext目录，使用java扩展机制的类位置；
   - 用户类路径（userClassPath）：用户实现的类或第三方类库位置；

2. 类路径可以是目录，jar包或zip包；

   ![类路径](D:\workspace\go\src\jvmgo\resource\类路径.png)

   bootClassPath和extClassPath，利用通配符得到目录下的jar包集合CompositeEntry；

   userClassPath是一个entry，可能是目录，CompositeEntry。。；


##### 二、解析class文件

![classfile结构](D:\workspace\go\src\jvmgo\resource\classfile结构.png)

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


##### 三、运行时数据区

1. **运行时数据区**：Java虚拟机在运行java程序时所使用的内存区域。

   多线程共享：JVM启动时创建，退出时销毁；

   线程私有：常见线程时创建，线程退出时销毁；（主要用于辅助执行java字节码）

   ![运行时数据区](D:\workspace\go\src\jvmgo\resource\运行时数据区.png)

2. **栈帧**

   每个方法执行的同时会创建一个栈帧，栈帧用于存储局部变量表、操作数栈、动态链接、方法出口等信息。
   每个方法从调用直至执行完成的过程，就对应着一个栈帧在虚拟机栈中入栈到出栈的过程。

   执行方法所需的局部变量表大小、操作数栈深度是由编译器预先计算好的，存储于class文件method_info；

##### 五、指令集和解释器

1. **字节码**

   每一个类或接口都会被java编译器编译成一个class文件，类或接口的信息就在class文件的method_info结构中。

   如果方法不是抽象的，也不是本地方法，方法的java代码就会被编译器编译成**字节码**（即使方法为空，也会生成一条return语句），存放在method_info结构的Code属性中。

   字节码中存放编码后的java虚拟机指令，每条指令都以一个单字节的操作码开头，跟零字节或多字节的操作数。

2. **指令集**

   java虚拟机定义了205条指令（单字节限制数目），可以按照用途分为11类：

   1. 常量指令：把常量推入操作数栈顶；

   2. 加载指令：从局部变量表获取变量，然后推入操作数栈顶；

   3. 存储指令：把变量从操作数栈顶弹出，存入局部变量表；

   4. 操作数栈指令：直接对操作数栈顶进行操作；

   5. 数学指令：算数指令、位移指令、布尔运算指令；

   6. 转换指令：类型转换；

   7. 比较指令：将比较结果推入操作数栈顶，或根据比较结果跳转；

   8. 控制指令

   9. 引用指令

   10. 扩展指令

   11. 保留指令：一条给调试器实现断点，另两条给JVM内部使用，不允许出现在class文件中；

3. **JVM解释器的逻辑**

     ```java
     for{
        pc:=calculatePC() //计算PC
        opCode:=bytecode[pc] //读取操作码
        inst:=createInst(opCode) //解释操作码,生成响应的指令
        inst.fetchOperands(bytecode) //指令读取操作数
        inst.execute() //执行指令
     }
     ```


##### 六、类和对象

1. 方法区

   方法区主要存放总class文件中获取的类信息。

   当JVM第1次使用某个类时，会搜索类路径，找到class文件，读取并解析class文件，将相关信息放在方法区。

   ```go
   //方法区中存放的类信息
   type Class struct {
   	accessFlags       uint16
   	name              string
   	superClassName    string
   	interfaceNames    []string
   	constantPool      *ConstantPool //运行时常量池
   	fields            []*Field
   	methods           []*Method
   	loader            *ClassLoader //类加载器指针
   	superClass        *Class
   	interfaces        []*Class
   	instanceSlotCount uint  //实体变量占据的空间大小
   	staticSlotCount   uint  //类变量占据的空间大小
   	staticVars        Slots //存放静态变量
   }
   type Field、Method struct {
   	accessFlags uint16
   	name        string
   	descriptor  string
   	class       *Class
   	constValueIndex uint
   	slotID          uint
   }

   //ConstantPool运行时常量池主要存放两类信息：字面量和符号引用
   type FieldRef struct {
     	cp        *ConstantPool
   	class     *Class //字段的类，用到时才解析
   	className string 
     	name       string
   	descriptor string
   	field *Field   //用到时才解析
   }
   ```

2. 类加载器

   1）找到class文件并把数据读取到内存；

   2）解析class文件，生成虚拟机可以使用的类数据，放入方法区；（加载父类，加载接口类）

   3）链接：类的验证+类变量分配空间并初始化；

3. `startJVM`工作流程

   1）通过classpath构造classLoader；

   2）加载类MyObject；

   3）执行main方法：（1）new thread（2）由方法构建frame（3）解析方法的code字节码，执行指令；


##### 七、方法调用和返回

1. **java虚拟机调用方法过程**

   - 方法调用指令需要N+1个操作数，其中1个是常量池中的方法索引，解析这个符号引用得到方法。剩下的N个操作数是要传递给被调用方法的参数；
   - 如果执行java方法（而非本地），创建新的帧，推到java虚拟机栈顶。传递参数后，执行新的方法；
   - 方法的最后一条指令是返回指令，将方法的返回值推入前一帧的操作数栈顶，然后将当前帧弹出；

2. **类初始化**

   类初始化就是执行类的初始化方法`<clinit>`

   类初始化触发条件：

   - 执行new指令创建实例，但类还没有被初始化；

   - 执行putstatic、getstatic指令存取静态变量，但声明该字段的类还没有被初始化；

   - 执行invokestatic调用类的静态方法，但声明该方法的类还没有被初始化；

   - 当初始化一个类时，如果类的超类还没有被初始化，要先初始化类的超类；

   - 执行某些反射操作时；

##### 八、数组和字符串

数组和普通类的区别：

- 普通类从class文件中加载，但是数组类由java虚拟机在运行时生成；
- 创建数组和创建普通对象的方式不同，指令不一样；
- 数组和普通对象存放的数据不同，普通对象存放实例变量，数组对象中存放数组元素；