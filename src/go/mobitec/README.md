```
go get github.com/prefixFelix/mobitec-flipdot/src/go/mobitec
```

Reimplementation of [Nosen92/maskin-flipdot](https://github.com/Nosen92/maskin-flipdot) in Go.

## Fonts

ToDo: Ist es bold/bolder/boldest oder wide/wider? Wahrscheinlich eher bold

| Size       | Addr | Description | Tested (y/n) | Notes (To Test)                                                                                 |
|------------|------|-------------|--------------|-------------------------------------------------------------------------------------------------|
| 5px        | 0x72 |             |              | "Large Letters Only" (bjarnekvae)                                                               |
| 6px        | 0x66 |             |              | 6 oder 7px?                                                                                     |
| 7px        | 0x60 |             |              | Laut Nosen92                                                                                    |
| 7px        | 0x65 |             |              | 13px laut nosen92                                                                               |
| 7px (*)    | 0x64 | bold        |              | bold (bjarnekvae) or wide (Nosen92) oder "small F" (mqtt fd)? 7px oder 5 px? 13 px laut Nosen92 |
| 9px        | 0x75 |             |              | 9 oder 13?                                                                                      |
| 9px (*)    | 0x70 | bold        |              | 9 oder 13?                                                                                      |
| 9px (**)   | 0x62 | bolder      |              | "bolder"? (bjarnekvae), 7px bold laut nosen92                                                   |
| 13px       | 0x73 |             |              |                                                                                                 |
| 13px (*)   | 0x69 | bold        |              |                                                                                                 |
| 13px (**)  | 0x61 | bolder      |              | "bolder"? (bjarnekvae)                                                                          |
| 13px (***) | 0x79 | boldest     |              | "boldest"? (bjarnekvae)                                                                         |
| 15px       | 0x71 |             |              |                                                                                                 |
| 16px       | 0x68 |             |              | Nur Numbers laut Nosen92                                                                        |
| 16px (*)   | 0x6a | bold        |              | ? laut Nosen92                                                                                  | 
| 16px (*)   | 0x78 | bold        |              | ? laut bjarnekvae                                                                               |
| 16px (**)  | 0x74 | bolder      |              | "bolder"? (bjarnekvae). Laut mqtt-fd funktioniert nur das A?                                    |
| 19px       | 0x63 |             |              | laut "mqtt-flipdot-driver", 12px laut nosen92                                                   |
| ?          | 0x76 |             |              | Unbekannt ob hier was ist                                                                       |

(*) "bold" (bjarnekvae) / wide (Nosen92)

(**) "bolder" (bjarnekvae)

(***) "boldest" (bjarnekvae)

## Others

| Font    | Addr | Description | Tested (y/n) | Test Notes |
|---------|------|-------------|--------------|------------|
| Symbols | 0x67 |             |              |            |
| Bitwise | 0x77 | Matrix Font |              |            |

## Symbols

| Addr    | Symbol   | 
|---------|----------|
| 0       | Factory  |
| 1       | Church   |
| 2       | Soccer1  |
| 3       | Soccer2  |
| 4       | Soccer3  |
| 5       | Horse1   |
| 6       | Horse2   |
| 7       | Horse3   |
| 8       | Horse4   |
| 9       | Horse5   |
| A       | K?       |
| B       | Swimmer  |
| C       | Stripes? |
| todo... | ...      |
