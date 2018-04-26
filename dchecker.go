package dchecker

type (
    Disk struct {
        Size float64
        Used float64
        Avail float64
        Capacity float64
        Mounted string
    }

    Check struct {
        Path string
        LastCleanUp date.Date
        CleanRule string
    }

    Dcheck interface {
        Disk
        Check
    }
)

func init() {
    return Dcheck
}

func (Dcheck)DiskInfo() (bool, error) {
    return true, nil
}


