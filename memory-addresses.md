# Memory addresses & hex

Memory is a long row of **1-byte cells**. An address is the cell's index, and addresses are written in **hex**.

**The smallest addressable unit is 1 byte = 8 bits = 2 hex digits**, so every cell holds `0x00`–`0xFF` (0–255). A single bit has no address of its own. To change one bit, the CPU reads the whole byte, changes the bit, and writes the byte back.

```
 address (full width, 16 hex digits on 64-bit)     value (1 byte = 2 hex digits)
 0x000000C000012340  ───────────────────────────►  48
 0x000000C000012341  ───────────────────────────►  69
 Each 0 is 1 hex => which is 0000 (4bits):
 64bit architecture: 00 00 00 00 00 00 00 00    ← 8 bytes, 2 hex digits each, 8 * 2 = 16hex * 4bits each = 64bit architecture
 32bit architecture: 00 00 00 00                ← 4 bytes, 2 hex digits each, 4 * 2 = 8hex * 4bits each = 32bit architecture
```

```
binary: 0000 0000 0000 0000 0000 0000 1100 0000 0000 0000 0000 0001 0010 0011 0100 0000
hex:       0    0    0    0    0    0    C    0    0    0    0    1    2    3    4    0
         → 0x000000C000012340
```
```
p (a pointer) = 8 cells × 2 hex digits = 16 hex digits = full address width (8bytes/16hex on 64bits, 4bytes/8hex on 32bits)
64bit example:
┌────┬────┬────┬────┬────┬────┬────┬────┐
│ 00 │ 00 │ 00 │ C0 │ 00 │ 01 │ 23 │ 40 │   → 0x000000C000012340
└────┴────┴────┴────┴────┴────┴────┴────┘
 1B   1B   1B   1B   1B   1B   1B   1B
```

## The key chain: 1 hex digit = 4 bits

```
16 = 2⁴  →  1 hex digit  = 4 bits (a nibble)
            2 hex digits = 8 bits = 1 byte   (0x00 – 0xFF)
```

| Binary | Hex | | Binary | Hex |
|---|---|---|---|---|
| 0000 | 0 | | 1000 | 8 |
| 0001 | 1 | | 1001 | 9 |
| 0010 | 2 | | 1010 | A |
| 0011 | 3 | | 1011 | B |
| 0100 | 4 | | 1100 | C |
| 0101 | 5 | | 1101 | D |
| 0110 | 6 | | 1110 | E |
| 0111 | 7 | | 1111 | F |

To convert, split the bits into groups of 4 and replace each group with its hex digit. You don't need any division:

```
1100 0000 1010 1000 0000 0001 0010 1010
   C    0    A    8    0    1    2    A   →  0xC0A8012A
```

## Per architecture

| | 32-bit | 64-bit |
|---|---|---|
| Address width | 32 bits = **4 bytes** | 64 bits = **8 bytes** |
| Hex digits per address | 32bit / 4bit = **8**hex | 64bit / 4bit = **16**hex |
| Lowest address | `0x00000000` | `0x0000000000000000` |
| Highest address | `0xFFFFFFFF` (2³² − 1) | `0xFFFFFFFFFFFFFFFF` (2⁶⁴ − 1) |
| Addressable space | 2³² B = **4 GiB** | 2⁶⁴ B ≈ **16 EiB** |
| Pointer size (Go `unsafe.Sizeof(p)`) | 4bytes | 8bytes |

```
bits ÷ 4 = hex digits      bits ÷ 8 = bytes      bytes × 2 = hex digits
```

In practice, x86-64 and ARM64 use only 48 bits of the address (256 TiB). The upper bits are unused, which is why real pointers look short, for example Go heap pointers like `0xc000012345`.

## Why hex, not decimal or binary

- **Maps exactly onto bits.** Each hex digit is 4 bits, so you can see the bit pattern directly. Decimal hides it.
- **Maps exactly onto bytes.** 2 digits = 1 byte, so an address or value splits cleanly into its bytes.
- **Compact.** It's 4× shorter than binary: 8 characters instead of 32.

## In Go

```go
x := 42
fmt.Printf("%p\n", &x)      // 0xc000012345  address
fmt.Printf("%x\n", 255)     // ff
fmt.Printf("%08b\n", 0xC)   // 00001100
fmt.Printf("% x\n", []byte("Hi")) // 48 69  (one byte = two hex digits)
```
