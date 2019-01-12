# Performance Library Comparison

## Purpose

The purpose of this project is to compare a few libraries that pull OS performance stats:

gopsutil
nmon (binary)

## Comparison

### CPU

#### Linux & MacOS
|  Config | gopsutil  | nmon  |   |   |
|---|---|---|---|---|
| vendorId |  x |   |   |   |
| family | x  |   |   |   |
| model | x  |   |   |   |
| stepping | x  |   |   |   |
| physicalId | x  |   |   |   |
| coreId |  x |  |   |   |
| cores (#) | x  |   |   |   |
| mhz |  x |   |   |   |
| cacheSize | x  |   |   |   |
| flags | x  |   |   |   |
| Microcode | x  |   |   |   |
| cpu family | x  |   |   |   |
| siblings |  | x  |   |   |
| apicid |   | x |   |   |
| fpu |  | x |   |   |
| wp |  | x |   |   |
| bugs |  | x |   |   |
| bogomips |  | x |   |   |
| clflush_size |  | x |   |   |
| cache alignment |   | x |   |   |
| address sizes |  | x  |   |   |

|  Metric | gopsutil  | nmon  |   |   |
|---|---|---|---|---|
| User |  x | x |   |   |
| System | x  | x |   |   |
| Idle | x  | x |   |   |
| Nice | x  |   |   |   |
| iowait | x  |   |   |   |
| irq | x |  |   |   |
| steal | x  |   |   |   |
| guest | x |   |   |   |
| guestNice | x  |   |   |   |
| stolen | x  |   |   |   |


    |  Metric | gopsutil  | nmon  |   |   |
|---|---|---|---|---|
|  |  x |   |   |   |
|  | x  |   |   |   |
|  | x  |   |   |   |
|  | x  |   |   |   |
|  | x  |   |   |   |
|  | x |  |   |   |
|  | x  |   |   |   |
|  | x |   |   |   |
|  | x  |   |   |   |
|  | x  |   |   |   |
|  | x  |   |   |   |
### Memory
### Network
### Disk
### Other