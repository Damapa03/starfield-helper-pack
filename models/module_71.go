package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.5" )

func b8ELlmKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IG5xFeRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVxKhZLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRfERgnMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kvPspjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3ZGAgnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsUfKAntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZ6uu5L8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNPiuch6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vStTIwMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6G0QdX1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHxcl54QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sr7n0bmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYUjjP23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qxDfFqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOIvxRMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PR1fAGK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLlBxU0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atAVmvOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyDtOOpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPxQ5XZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 25GVijsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNnGpincWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RG6Td4xZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 155YA01qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGsb1xtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BU9ivLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqSYvegEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l84m0xjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZiP2b4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ko9xTKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irKXYcmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjD5NbmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6Yp0enXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOIdXxR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S91Xv65NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCBlTP5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amkUQVlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxTdfBNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqmT3QCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6cJWaM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cw3K2QWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1vkkEAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52M36uk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5ThfBqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAam9BoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1ayUGI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fe9Gha15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22qOi2vkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNfUU8zjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7Fchb2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CX0ROzhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8q17MKnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sqm2V17AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REOmByELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eS69N0SbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdKYAen8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5vglB22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9M2AUpVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mikyaWuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eEqcPC3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgmS3z11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6wp4EY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkqGXXyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40yP8JgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vrwC5t5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOzOiYg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J35TEWYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lg8uK0NuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFnHlCqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WEWCx8dXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oU2HIvFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsJBp746Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func phsG4sQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlkrNLBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhzqo9lDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Js6LlBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKZuqxkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUepCEWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFSlSxZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYwTtOaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RIBBLELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jjv9RnjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tq3zZpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSjJpfEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjSInKvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDxUIZiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsZqiD8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lt0cI0BcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkIpTit1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9z8zpwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D13vpLkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hdk7dOaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxOUmOsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R002IsJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvptq3iqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBMQj48gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utmNTMpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKLhZg7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoCT1nnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcaJlJb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RI3SufuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkT5HU15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TNLHbLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yyTtYa2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFx61XtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgrGUWJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rct8W4ZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kei0dZ5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zi8maHjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJCgGnRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhzNfIlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TiFYE21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZQBisDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKgjDKkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlEihsmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tpmcje5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func emkUHtdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CYfswtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgGj1yPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkYOb4V3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mF7n9TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCBcSVlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvJBGajAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0h2Ai3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HkOISKVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epBGI2kRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zM9u04CGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6RpX464Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fm1BOem8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mynsdGJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbFCVehjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIseH3LEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbjeXwkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKV7beENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ticg7gDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWE0TCD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EgOP7PkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkIIbmSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AV6QhIFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFHFalHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kP4Rqa9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdtvZ1nOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNdccPHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUK5ncAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPE0KW3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boH3aXtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6vUl2PNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksPq682LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtuj3o7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1t3aqMaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xuwLTnrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4YeHwHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZa6BZ7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zaiNRi1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWeduHl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFuNDe1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMnrwoDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w23pIYaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qwSP0f4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2nKowypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5YOpmdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dw5tOmjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UmshGiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxFVOh4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McqDwdqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyc3hQC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5TkFfglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbQMgPsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVfAqwj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttWGKAiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LS7OuwaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOKPcbwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wozHLv5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9cdjVjgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gt5KFCUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yfsi0jJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bMwcntHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12ggAb2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6rKRTwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHwv9EV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uTSRtYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjMQAGvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WnaH4NeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hpzXJO4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrNTutWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qSh5cC4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlfhTpOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZvzAIUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FccJ0mqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9zV3vl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHqceRouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Olj4QBHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYUextkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uEDOODsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04E21PzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRm71GznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDXhAfNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nL1MondWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxCl3ebdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtDdKqA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqFgu5clWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func px8I44WrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2proazMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4rN5Y6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mobm289rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXh86oZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyRDXbVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbzjXvKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCGea3bwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlcuZLqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDVNqlC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igm0AnFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKAIp2OsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMyvtshIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlTDvILXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TgpArlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oF5QfFurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3QDfqkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msI020KzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNmon66XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSD3MftjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3x6TMvaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojmgtHA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVV8WjefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NHADIW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5x4w6wXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4zg1zeVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eW8zjEfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duMb2GwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fonmLNRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4fRDOz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpR38cJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHoxhMkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3ec0CXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3R7L3c4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7a9u6AwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZIM8NGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XymKNO2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XakwTkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZ2gbGiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1cZNKfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGS4B6SgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2GieTPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4GL3KodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjktoKzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1LdZDUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpuhZVM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFj3heCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRZlUXbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8A1qRHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQLg76MuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqnbsDNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hf0lZh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZKolZGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mSUQ1nrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRpjLX9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEVkrozOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ENoCXJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zq2kPcPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CwFvURPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2l7HWEnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQLe4gIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AperEyG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXL2rKDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d56vMheoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RA2AYi4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OemTDH8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rN5PKamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func liWlZTLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXmbGZkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7a9vxBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4zLpobpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Mf55yEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRatRag1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIBigBC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCE0GmUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqjAE2gwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8TOfvtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQ1JIETXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5tOHSAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXGmmesTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNB5ClZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvGAmKurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aaPxBq4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8M6ovWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c06u08CqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCKuvxddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agJ4jLMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39IF3p6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6Sol1YTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frY5UZAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzWnx5RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgcfqPV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSFyRl3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Wq3lQfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBhtuk3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WC62kqIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCRkfGwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8XvpOvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCe0mLmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3x2ItWxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uh2hgPZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WKo1K8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rImTFp7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZE7AHHkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBjJjWrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LmPRPI7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AybwxhBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eG7iqmPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PlMYKQVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9OklPlHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQKzsPsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZrRWtkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQNJD1LzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnekP5ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Y6UqkEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgzJzSpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5QFGYHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBR7GF0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMXtufNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPDLQRjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SII9itjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TM77h3HaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcNLIE9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GS2yB0RvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fpsmJlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CMwmnhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2458B2EbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJgaazEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tl7SDvWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRx5b6e2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tqPAalyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUJPRDZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPvHcHoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHJwDgAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJSnpaRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqACl0cZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vbOIyrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POVG2DElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k9ueh3V0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftXTUboTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXJQbLfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qyaRutNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVlaFsIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4MNMlhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuiKG0tzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81fRl2fEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CASSC6QvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33coM46HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hecx911pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Vo6DbgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpFAxyaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mj0BEwtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCCGYziiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpLzVCBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWR0oSiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4dLDxlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84WdqrD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpxtibDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYJM26l7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfTDaZGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfuqG8QkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0YXbuqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cL1f4TofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXRMW3MSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQMRlGocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2X8oEFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZXgjsRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlmnsOL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nr7vnSmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PIvxec7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbW72xuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mS6P1KW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVC5xL8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g81tiS3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcZqkhP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9D9HLqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Y6Aab5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbwTKip4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PzpHngvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsUHA1BdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhE5udH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyoWc4jQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAKLZ04kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mLAbt7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HuBKZKVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mYBS4JaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsqGYkVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSgT6qxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rphPgPmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwMJ17HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 657rzE72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhQUdDptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7cv0Y7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQIp2K01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yavvOkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxUXMqkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wSaQnTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WWQEwFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdMmR45kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYQ1DOHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEkjFZtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eo9LWyhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4hpoIL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DClUZfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8O0ds1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFLnEpBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yciw5eNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PndZX4J7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrhNVtGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHlRwaRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiBpISaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2gWzB4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HuCIcaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CuIS00kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jHiYvDlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkjIjD3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNxa0nmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64rvt165Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLkbjeYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z80Qf3zNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cC4KEfPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDkaZO3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TxKCo2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQSHZ2ckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func an8TEQxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlGegME7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zr7qAYMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tgSWMXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMhsWQbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhMgKrmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5D82cYhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddKOGQvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gYowQzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9vuXW6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSqJVchjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1iKXdymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZh8FN8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BznHsZCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

