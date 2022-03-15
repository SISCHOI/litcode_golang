func findRepeatNumber(nums []int) int {
    nmap:=make(map[int]bool)
    for _,num:=range nums {
        if _,ok:=nmap[num];ok{
            return num
        }
        nmap[num]=true
    }
    return 0
}