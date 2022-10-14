|aturan slice   | penjelasan                                                 |
|---------------|------------------------------------------------------------|
|array[low:high]|membuat slice dari array index low sampai index sebelum high|
|array[low:]    |membuat slice dari array index low sampai index akhir       |
|array[:high]   |membuat slice dari array index 0 sampai index sebelum high  |
|array[:]       |membuat slice dari array index 0 sampau index akhir         |

array = [12]int {
    0,
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    9,
    10,
    11,
}

array[4:7] {
    |slice       |
    |------------|
    |pointer = 4 |
    |length = 3  |
    |capacity = 8|
    
} 



