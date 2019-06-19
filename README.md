# RVm
VM for RuixueLanguage.

create a file named "test.rvm" under the same path of RVM.
you can do:
> * nano test.rvm
> * ./rvm_x64.exe //on windows

Enjoy it!

这是一个被弃用的版本，先说说为什么放弃Go语言编写RVM吧。
## 首先是因为Go这门语言过于简单（没错，简单有时候也会带来坏处）导致很多实现需要绕很大的圈子（也可能是我技术太差） 所以我不得不切换到了C++语言。无论如何RVM项目还是在继续的。
## 第二 GC开销大，在高密度的内存操作下，GC会耗费大量的时间，可能是我编写的代码对GC并不友好所以导致这个问题，但是我实在没有精力去顾及太多。
## 第三 Go语言过于依赖系统，包括GC等，加大了对操作系统平台移植的困难，很难支持不在go编译器支持范围内的系统。
## 所以无奈之下，我只能放弃了这一版的项目
## 事实上，在尝试了Go语言之后，我也曾经尝试过Rust，不过Rust的过于安全令我很讨厌，多个库共用一个变量需要使用一些包装，所以最后转到了C++。现在C++版本是可以正常运行。
