// Code generated by go generate; DO NOT EDIT.

package wren

/*
#cgo CFLAGS:
#cgo LDFLAGS: -lm
#include "wren.h"

extern void f0(WrenVM* vm);
extern void f1(WrenVM* vm);
extern void f2(WrenVM* vm);
extern void f3(WrenVM* vm);
extern void f4(WrenVM* vm);
extern void f5(WrenVM* vm);
extern void f6(WrenVM* vm);
extern void f7(WrenVM* vm);
extern void f8(WrenVM* vm);
extern void f9(WrenVM* vm);
extern void f10(WrenVM* vm);
extern void f11(WrenVM* vm);
extern void f12(WrenVM* vm);
extern void f13(WrenVM* vm);
extern void f14(WrenVM* vm);
extern void f15(WrenVM* vm);
extern void f16(WrenVM* vm);
extern void f17(WrenVM* vm);
extern void f18(WrenVM* vm);
extern void f19(WrenVM* vm);
extern void f20(WrenVM* vm);
extern void f21(WrenVM* vm);
extern void f22(WrenVM* vm);
extern void f23(WrenVM* vm);
extern void f24(WrenVM* vm);
extern void f25(WrenVM* vm);
extern void f26(WrenVM* vm);
extern void f27(WrenVM* vm);
extern void f28(WrenVM* vm);
extern void f29(WrenVM* vm);
extern void f30(WrenVM* vm);
extern void f31(WrenVM* vm);
extern void f32(WrenVM* vm);
extern void f33(WrenVM* vm);
extern void f34(WrenVM* vm);
extern void f35(WrenVM* vm);
extern void f36(WrenVM* vm);
extern void f37(WrenVM* vm);
extern void f38(WrenVM* vm);
extern void f39(WrenVM* vm);
extern void f40(WrenVM* vm);
extern void f41(WrenVM* vm);
extern void f42(WrenVM* vm);
extern void f43(WrenVM* vm);
extern void f44(WrenVM* vm);
extern void f45(WrenVM* vm);
extern void f46(WrenVM* vm);
extern void f47(WrenVM* vm);
extern void f48(WrenVM* vm);
extern void f49(WrenVM* vm);
extern void f50(WrenVM* vm);
extern void f51(WrenVM* vm);
extern void f52(WrenVM* vm);
extern void f53(WrenVM* vm);
extern void f54(WrenVM* vm);
extern void f55(WrenVM* vm);
extern void f56(WrenVM* vm);
extern void f57(WrenVM* vm);
extern void f58(WrenVM* vm);
extern void f59(WrenVM* vm);
extern void f60(WrenVM* vm);
extern void f61(WrenVM* vm);
extern void f62(WrenVM* vm);
extern void f63(WrenVM* vm);
extern void f64(WrenVM* vm);
extern void f65(WrenVM* vm);
extern void f66(WrenVM* vm);
extern void f67(WrenVM* vm);
extern void f68(WrenVM* vm);
extern void f69(WrenVM* vm);
extern void f70(WrenVM* vm);
extern void f71(WrenVM* vm);
extern void f72(WrenVM* vm);
extern void f73(WrenVM* vm);
extern void f74(WrenVM* vm);
extern void f75(WrenVM* vm);
extern void f76(WrenVM* vm);
extern void f77(WrenVM* vm);
extern void f78(WrenVM* vm);
extern void f79(WrenVM* vm);
extern void f80(WrenVM* vm);
extern void f81(WrenVM* vm);
extern void f82(WrenVM* vm);
extern void f83(WrenVM* vm);
extern void f84(WrenVM* vm);
extern void f85(WrenVM* vm);
extern void f86(WrenVM* vm);
extern void f87(WrenVM* vm);
extern void f88(WrenVM* vm);
extern void f89(WrenVM* vm);
extern void f90(WrenVM* vm);
extern void f91(WrenVM* vm);
extern void f92(WrenVM* vm);
extern void f93(WrenVM* vm);
extern void f94(WrenVM* vm);
extern void f95(WrenVM* vm);
extern void f96(WrenVM* vm);
extern void f97(WrenVM* vm);
extern void f98(WrenVM* vm);
extern void f99(WrenVM* vm);
extern void f100(WrenVM* vm);
extern void f101(WrenVM* vm);
extern void f102(WrenVM* vm);
extern void f103(WrenVM* vm);
extern void f104(WrenVM* vm);
extern void f105(WrenVM* vm);
extern void f106(WrenVM* vm);
extern void f107(WrenVM* vm);
extern void f108(WrenVM* vm);
extern void f109(WrenVM* vm);
extern void f110(WrenVM* vm);
extern void f111(WrenVM* vm);
extern void f112(WrenVM* vm);
extern void f113(WrenVM* vm);
extern void f114(WrenVM* vm);
extern void f115(WrenVM* vm);
extern void f116(WrenVM* vm);
extern void f117(WrenVM* vm);
extern void f118(WrenVM* vm);
extern void f119(WrenVM* vm);
extern void f120(WrenVM* vm);
extern void f121(WrenVM* vm);
extern void f122(WrenVM* vm);
extern void f123(WrenVM* vm);
extern void f124(WrenVM* vm);
extern void f125(WrenVM* vm);
extern void f126(WrenVM* vm);
extern void f127(WrenVM* vm);

static inline WrenForeignMethodFn get_f(int i) {
	switch (i) {
		case 0: return f0;
		case 1: return f1;
		case 2: return f2;
		case 3: return f3;
		case 4: return f4;
		case 5: return f5;
		case 6: return f6;
		case 7: return f7;
		case 8: return f8;
		case 9: return f9;
		case 10: return f10;
		case 11: return f11;
		case 12: return f12;
		case 13: return f13;
		case 14: return f14;
		case 15: return f15;
		case 16: return f16;
		case 17: return f17;
		case 18: return f18;
		case 19: return f19;
		case 20: return f20;
		case 21: return f21;
		case 22: return f22;
		case 23: return f23;
		case 24: return f24;
		case 25: return f25;
		case 26: return f26;
		case 27: return f27;
		case 28: return f28;
		case 29: return f29;
		case 30: return f30;
		case 31: return f31;
		case 32: return f32;
		case 33: return f33;
		case 34: return f34;
		case 35: return f35;
		case 36: return f36;
		case 37: return f37;
		case 38: return f38;
		case 39: return f39;
		case 40: return f40;
		case 41: return f41;
		case 42: return f42;
		case 43: return f43;
		case 44: return f44;
		case 45: return f45;
		case 46: return f46;
		case 47: return f47;
		case 48: return f48;
		case 49: return f49;
		case 50: return f50;
		case 51: return f51;
		case 52: return f52;
		case 53: return f53;
		case 54: return f54;
		case 55: return f55;
		case 56: return f56;
		case 57: return f57;
		case 58: return f58;
		case 59: return f59;
		case 60: return f60;
		case 61: return f61;
		case 62: return f62;
		case 63: return f63;
		case 64: return f64;
		case 65: return f65;
		case 66: return f66;
		case 67: return f67;
		case 68: return f68;
		case 69: return f69;
		case 70: return f70;
		case 71: return f71;
		case 72: return f72;
		case 73: return f73;
		case 74: return f74;
		case 75: return f75;
		case 76: return f76;
		case 77: return f77;
		case 78: return f78;
		case 79: return f79;
		case 80: return f80;
		case 81: return f81;
		case 82: return f82;
		case 83: return f83;
		case 84: return f84;
		case 85: return f85;
		case 86: return f86;
		case 87: return f87;
		case 88: return f88;
		case 89: return f89;
		case 90: return f90;
		case 91: return f91;
		case 92: return f92;
		case 93: return f93;
		case 94: return f94;
		case 95: return f95;
		case 96: return f96;
		case 97: return f97;
		case 98: return f98;
		case 99: return f99;
		case 100: return f100;
		case 101: return f101;
		case 102: return f102;
		case 103: return f103;
		case 104: return f104;
		case 105: return f105;
		case 106: return f106;
		case 107: return f107;
		case 108: return f108;
		case 109: return f109;
		case 110: return f110;
		case 111: return f111;
		case 112: return f112;
		case 113: return f113;
		case 114: return f114;
		case 115: return f115;
		case 116: return f116;
		case 117: return f117;
		case 118: return f118;
		case 119: return f119;
		case 120: return f120;
		case 121: return f121;
		case 122: return f122;
		case 123: return f123;
		case 124: return f124;
		case 125: return f125;
		case 126: return f126;
		case 127: return f127;
		default: return (void*)(0);
	}
}
*/
import "C"
import (
	"fmt"
)

const MAX_REGISTRATIONS = 128

//export f0
func f0(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 0 {
			vm.bindMap[0](vm)
		}
	}
}

//export f1
func f1(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 1 {
			vm.bindMap[1](vm)
		}
	}
}

//export f2
func f2(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 2 {
			vm.bindMap[2](vm)
		}
	}
}

//export f3
func f3(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 3 {
			vm.bindMap[3](vm)
		}
	}
}

//export f4
func f4(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 4 {
			vm.bindMap[4](vm)
		}
	}
}

//export f5
func f5(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 5 {
			vm.bindMap[5](vm)
		}
	}
}

//export f6
func f6(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 6 {
			vm.bindMap[6](vm)
		}
	}
}

//export f7
func f7(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 7 {
			vm.bindMap[7](vm)
		}
	}
}

//export f8
func f8(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 8 {
			vm.bindMap[8](vm)
		}
	}
}

//export f9
func f9(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 9 {
			vm.bindMap[9](vm)
		}
	}
}

//export f10
func f10(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 10 {
			vm.bindMap[10](vm)
		}
	}
}

//export f11
func f11(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 11 {
			vm.bindMap[11](vm)
		}
	}
}

//export f12
func f12(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 12 {
			vm.bindMap[12](vm)
		}
	}
}

//export f13
func f13(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 13 {
			vm.bindMap[13](vm)
		}
	}
}

//export f14
func f14(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 14 {
			vm.bindMap[14](vm)
		}
	}
}

//export f15
func f15(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 15 {
			vm.bindMap[15](vm)
		}
	}
}

//export f16
func f16(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 16 {
			vm.bindMap[16](vm)
		}
	}
}

//export f17
func f17(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 17 {
			vm.bindMap[17](vm)
		}
	}
}

//export f18
func f18(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 18 {
			vm.bindMap[18](vm)
		}
	}
}

//export f19
func f19(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 19 {
			vm.bindMap[19](vm)
		}
	}
}

//export f20
func f20(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 20 {
			vm.bindMap[20](vm)
		}
	}
}

//export f21
func f21(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 21 {
			vm.bindMap[21](vm)
		}
	}
}

//export f22
func f22(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 22 {
			vm.bindMap[22](vm)
		}
	}
}

//export f23
func f23(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 23 {
			vm.bindMap[23](vm)
		}
	}
}

//export f24
func f24(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 24 {
			vm.bindMap[24](vm)
		}
	}
}

//export f25
func f25(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 25 {
			vm.bindMap[25](vm)
		}
	}
}

//export f26
func f26(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 26 {
			vm.bindMap[26](vm)
		}
	}
}

//export f27
func f27(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 27 {
			vm.bindMap[27](vm)
		}
	}
}

//export f28
func f28(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 28 {
			vm.bindMap[28](vm)
		}
	}
}

//export f29
func f29(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 29 {
			vm.bindMap[29](vm)
		}
	}
}

//export f30
func f30(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 30 {
			vm.bindMap[30](vm)
		}
	}
}

//export f31
func f31(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 31 {
			vm.bindMap[31](vm)
		}
	}
}

//export f32
func f32(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 32 {
			vm.bindMap[32](vm)
		}
	}
}

//export f33
func f33(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 33 {
			vm.bindMap[33](vm)
		}
	}
}

//export f34
func f34(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 34 {
			vm.bindMap[34](vm)
		}
	}
}

//export f35
func f35(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 35 {
			vm.bindMap[35](vm)
		}
	}
}

//export f36
func f36(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 36 {
			vm.bindMap[36](vm)
		}
	}
}

//export f37
func f37(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 37 {
			vm.bindMap[37](vm)
		}
	}
}

//export f38
func f38(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 38 {
			vm.bindMap[38](vm)
		}
	}
}

//export f39
func f39(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 39 {
			vm.bindMap[39](vm)
		}
	}
}

//export f40
func f40(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 40 {
			vm.bindMap[40](vm)
		}
	}
}

//export f41
func f41(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 41 {
			vm.bindMap[41](vm)
		}
	}
}

//export f42
func f42(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 42 {
			vm.bindMap[42](vm)
		}
	}
}

//export f43
func f43(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 43 {
			vm.bindMap[43](vm)
		}
	}
}

//export f44
func f44(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 44 {
			vm.bindMap[44](vm)
		}
	}
}

//export f45
func f45(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 45 {
			vm.bindMap[45](vm)
		}
	}
}

//export f46
func f46(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 46 {
			vm.bindMap[46](vm)
		}
	}
}

//export f47
func f47(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 47 {
			vm.bindMap[47](vm)
		}
	}
}

//export f48
func f48(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 48 {
			vm.bindMap[48](vm)
		}
	}
}

//export f49
func f49(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 49 {
			vm.bindMap[49](vm)
		}
	}
}

//export f50
func f50(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 50 {
			vm.bindMap[50](vm)
		}
	}
}

//export f51
func f51(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 51 {
			vm.bindMap[51](vm)
		}
	}
}

//export f52
func f52(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 52 {
			vm.bindMap[52](vm)
		}
	}
}

//export f53
func f53(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 53 {
			vm.bindMap[53](vm)
		}
	}
}

//export f54
func f54(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 54 {
			vm.bindMap[54](vm)
		}
	}
}

//export f55
func f55(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 55 {
			vm.bindMap[55](vm)
		}
	}
}

//export f56
func f56(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 56 {
			vm.bindMap[56](vm)
		}
	}
}

//export f57
func f57(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 57 {
			vm.bindMap[57](vm)
		}
	}
}

//export f58
func f58(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 58 {
			vm.bindMap[58](vm)
		}
	}
}

//export f59
func f59(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 59 {
			vm.bindMap[59](vm)
		}
	}
}

//export f60
func f60(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 60 {
			vm.bindMap[60](vm)
		}
	}
}

//export f61
func f61(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 61 {
			vm.bindMap[61](vm)
		}
	}
}

//export f62
func f62(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 62 {
			vm.bindMap[62](vm)
		}
	}
}

//export f63
func f63(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 63 {
			vm.bindMap[63](vm)
		}
	}
}

//export f64
func f64(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 64 {
			vm.bindMap[64](vm)
		}
	}
}

//export f65
func f65(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 65 {
			vm.bindMap[65](vm)
		}
	}
}

//export f66
func f66(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 66 {
			vm.bindMap[66](vm)
		}
	}
}

//export f67
func f67(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 67 {
			vm.bindMap[67](vm)
		}
	}
}

//export f68
func f68(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 68 {
			vm.bindMap[68](vm)
		}
	}
}

//export f69
func f69(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 69 {
			vm.bindMap[69](vm)
		}
	}
}

//export f70
func f70(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 70 {
			vm.bindMap[70](vm)
		}
	}
}

//export f71
func f71(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 71 {
			vm.bindMap[71](vm)
		}
	}
}

//export f72
func f72(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 72 {
			vm.bindMap[72](vm)
		}
	}
}

//export f73
func f73(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 73 {
			vm.bindMap[73](vm)
		}
	}
}

//export f74
func f74(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 74 {
			vm.bindMap[74](vm)
		}
	}
}

//export f75
func f75(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 75 {
			vm.bindMap[75](vm)
		}
	}
}

//export f76
func f76(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 76 {
			vm.bindMap[76](vm)
		}
	}
}

//export f77
func f77(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 77 {
			vm.bindMap[77](vm)
		}
	}
}

//export f78
func f78(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 78 {
			vm.bindMap[78](vm)
		}
	}
}

//export f79
func f79(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 79 {
			vm.bindMap[79](vm)
		}
	}
}

//export f80
func f80(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 80 {
			vm.bindMap[80](vm)
		}
	}
}

//export f81
func f81(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 81 {
			vm.bindMap[81](vm)
		}
	}
}

//export f82
func f82(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 82 {
			vm.bindMap[82](vm)
		}
	}
}

//export f83
func f83(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 83 {
			vm.bindMap[83](vm)
		}
	}
}

//export f84
func f84(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 84 {
			vm.bindMap[84](vm)
		}
	}
}

//export f85
func f85(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 85 {
			vm.bindMap[85](vm)
		}
	}
}

//export f86
func f86(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 86 {
			vm.bindMap[86](vm)
		}
	}
}

//export f87
func f87(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 87 {
			vm.bindMap[87](vm)
		}
	}
}

//export f88
func f88(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 88 {
			vm.bindMap[88](vm)
		}
	}
}

//export f89
func f89(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 89 {
			vm.bindMap[89](vm)
		}
	}
}

//export f90
func f90(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 90 {
			vm.bindMap[90](vm)
		}
	}
}

//export f91
func f91(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 91 {
			vm.bindMap[91](vm)
		}
	}
}

//export f92
func f92(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 92 {
			vm.bindMap[92](vm)
		}
	}
}

//export f93
func f93(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 93 {
			vm.bindMap[93](vm)
		}
	}
}

//export f94
func f94(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 94 {
			vm.bindMap[94](vm)
		}
	}
}

//export f95
func f95(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 95 {
			vm.bindMap[95](vm)
		}
	}
}

//export f96
func f96(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 96 {
			vm.bindMap[96](vm)
		}
	}
}

//export f97
func f97(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 97 {
			vm.bindMap[97](vm)
		}
	}
}

//export f98
func f98(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 98 {
			vm.bindMap[98](vm)
		}
	}
}

//export f99
func f99(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 99 {
			vm.bindMap[99](vm)
		}
	}
}

//export f100
func f100(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 100 {
			vm.bindMap[100](vm)
		}
	}
}

//export f101
func f101(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 101 {
			vm.bindMap[101](vm)
		}
	}
}

//export f102
func f102(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 102 {
			vm.bindMap[102](vm)
		}
	}
}

//export f103
func f103(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 103 {
			vm.bindMap[103](vm)
		}
	}
}

//export f104
func f104(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 104 {
			vm.bindMap[104](vm)
		}
	}
}

//export f105
func f105(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 105 {
			vm.bindMap[105](vm)
		}
	}
}

//export f106
func f106(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 106 {
			vm.bindMap[106](vm)
		}
	}
}

//export f107
func f107(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 107 {
			vm.bindMap[107](vm)
		}
	}
}

//export f108
func f108(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 108 {
			vm.bindMap[108](vm)
		}
	}
}

//export f109
func f109(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 109 {
			vm.bindMap[109](vm)
		}
	}
}

//export f110
func f110(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 110 {
			vm.bindMap[110](vm)
		}
	}
}

//export f111
func f111(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 111 {
			vm.bindMap[111](vm)
		}
	}
}

//export f112
func f112(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 112 {
			vm.bindMap[112](vm)
		}
	}
}

//export f113
func f113(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 113 {
			vm.bindMap[113](vm)
		}
	}
}

//export f114
func f114(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 114 {
			vm.bindMap[114](vm)
		}
	}
}

//export f115
func f115(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 115 {
			vm.bindMap[115](vm)
		}
	}
}

//export f116
func f116(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 116 {
			vm.bindMap[116](vm)
		}
	}
}

//export f117
func f117(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 117 {
			vm.bindMap[117](vm)
		}
	}
}

//export f118
func f118(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 118 {
			vm.bindMap[118](vm)
		}
	}
}

//export f119
func f119(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 119 {
			vm.bindMap[119](vm)
		}
	}
}

//export f120
func f120(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 120 {
			vm.bindMap[120](vm)
		}
	}
}

//export f121
func f121(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 121 {
			vm.bindMap[121](vm)
		}
	}
}

//export f122
func f122(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 122 {
			vm.bindMap[122](vm)
		}
	}
}

//export f123
func f123(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 123 {
			vm.bindMap[123](vm)
		}
	}
}

//export f124
func f124(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 124 {
			vm.bindMap[124](vm)
		}
	}
}

//export f125
func f125(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 125 {
			vm.bindMap[125](vm)
		}
	}
}

//export f126
func f126(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 126 {
			vm.bindMap[126](vm)
		}
	}
}

//export f127
func f127(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > 127 {
			vm.bindMap[127](vm)
		}
	}
}

type MaxBindingsReached struct {
	VM *VM
}

func (err *MaxBindingsReached)Error() string {
	return fmt.Sprintf("Cannot bind more that %v functions or classes", MAX_REGISTRATIONS)
}

func (vm *VM) registerFunc(fn ForeignMethodFn) error {
	if len(vm.bindMap) >= MAX_REGISTRATIONS {
		return &MaxBindingsReached{VM: vm}
	}

	vm.bindMap = append(vm.bindMap, fn)
	return nil
}
