```
                                                                                                                       
                                caller                                                                                 
                          +------------------+                                                                         
                          |                  |                                                                         
|---------------------->  --------------------                                                                         
|                         |                  |                                                                         
|                         | caller parent BP |                                                                         
|           BP(pseudo SP) --------------------                                                                         
|                         |                  |                                                                         
|                         |   Local Var0     |                                                                         
|                         --------------------                                                                         
|                         |                  |                                                                         
|                         |   .......        |                                                                         
|                         --------------------                                                                         
|                         |                  |                                                                         
|                         |   Local VarN     |                                                                         
|                         --------------------                                                                         
|  caller stack frame     |                  |                                                                         
|                         |   callee arg2    |                                                                         
|                         |------------------|                                                                         
|                         |                  |                                                                         
|                         |   callee arg1    |                                                                         
|                         |------------------|                                                                         
|                         |                  |                                                                         
|                         |   callee arg0    |                                                                         
|                         ----------------------------------------------+   FP(virtual register)                       
|                         |                  |                          |                                              
|                         |   return addr    |  parent return address   |                                              
|--------------------->   +------------------+--------------------------- <----------------------|                     
                                             |  caller BP               |                        |                     
                                             |  (caller frame pointer)  |                        |                     
                              BP(pseudo SP)  ----------------------------                        |                     
                                             |                          |                        |                     
                                             |     Local Var0           |                        |                     
                                             ----------------------------                        |                     
                                             |                          |                        |   callee stack frame
                                             |     Local Var1           |                        |                     
                                             ----------------------------                        |                     
                                             |                          |                        |                     
                                             |       .....              |                        |                     
                                             ----------------------------                        |                     
                                             |                          |                        |                     
                                             |     Local VarN           |                        |                     
                           SP(Real Register) ---------------------------- <----------------------|                     
                                             |                          |                                              
                                             |                          |                                              
                                       -     |                          |                                              
                                             |                          |                                              
                                             |                          |                                              
                                             +--------------------------+                                              
                                                                                                                       
                                                       callee                                                          
```
