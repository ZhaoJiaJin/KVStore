package logging

import(
    "os"
    log "github.com/labstack/gommon/log"
    "io/ioutil"
    "unsafe"
    "path/filepath"
    "syscall"
)

const(
    LOGSIZE = 100
    DATSIZE = 1024
)

type DB interface{
    Apply(data []byte)error
    GetCheckPoint()([]byte, error)
    RecoverCheckPoint(data []byte)(error)
}

type Log struct{
    Data [DATSIZE]byte
    DataLen int64
    Term int64
    ID int64
    Used bool
}

type CheckPoint struct{
    Data []byte
    LastTerm int
    LastID int
}

type LogStore struct{
    db DB
    cpF string
    commitsF string
    commitsFD *os.File
    mmap []byte
    commits *[LOGSIZE]Log
    idx int
}

func NewDBLogStore(db DB, dir string)(l *LogStore,err error){
    l = &LogStore{}
    os.MkdirAll(dir,0755)
    l.db = db
    l.cpF = filepath.Join(dir,"checkpoint")
    l.commitsF = filepath.Join(dir, "wal")
    size := int(unsafe.Sizeof(Log{})) * LOGSIZE
    l.commitsFD, err = os.OpenFile(l.commitsF, os.O_RDWR | os.O_CREATE, 0644)
    if err != nil{
        return
    }
    _, err = l.commitsFD.Seek(int64(size-1), 0)
    if err != nil {
        return 
    }
    _, err = l.commitsFD.Write([]byte(" "))
    if err != nil {
        return
    }
    log.Info("binding mmap file")
    l.mmap, err = syscall.Mmap(int(l.commitsFD.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
    if err != nil {
        return 
    }
    log.Info("init commits array")
    l.commits = (*[LOGSIZE]Log)(unsafe.Pointer(&(l.mmap)[0]))

    /*log.Info("init idx")
    for l.idx = 0; l.idx < LOGSIZE; l.idx++{
        if (*(l.commits))[l.idx].Used == false{
            break
        }
    }*/

    //db recovery
    log.Info("recover from log")
    err = l.recoverdb()
    if err != nil{
        return l, err
    }
    return l,nil
}

func (l *LogStore) recoverdb()(error){
    // recover checkpoint
    cpdata, err := ioutil.ReadFile(l.cpF)
    if err == nil{
        err = l.db.RecoverCheckPoint(cpdata)
        if err != nil{
            return err
        }
    }else{
        cpdata,err = l.db.GetCheckPoint()
        if err != nil{
            log.Warnf("recoverdb :%v",err)
            return err
        }
        err = ioutil.WriteFile(l.cpF,cpdata,0644)
        if err != nil{
            log.Warnf("recoverdb :%v",err)
            return err
        }
    }
    // replay all the commits
    //for i:=0; i < l.idx; i ++{
    for l.idx = 0; l.idx < LOGSIZE; l.idx++{
        if (*(l.commits))[l.idx].Used == false{
            break
        }
        cmt := l.commits[l.idx]
        err := l.db.Apply(cmt.Data[:cmt.DataLen])
        if err != nil{
            log.Warnf("recover: replay log:%v",err)
        }
    }

    return nil
}



func (l *LogStore)Apply(term, id int64, data []byte)(error){
    l.commits[l.idx].Used = true
    l.commits[l.idx].Term = term
    l.commits[l.idx].ID = id
    l.commits[l.idx].DataLen = int64(copy(l.commits[l.idx].Data[:], data))
    l.idx ++

    //Apply the operation
    err := l.db.Apply(data)
    if err != nil{
        return err
    }
    if l.idx >= LOGSIZE{
        // take checkpoint and reset all log entries to unused
        log.Info("take a checkpoint")
        cpdata,err := l.db.GetCheckPoint()
        if err != nil{
            return err
        }
        err = ioutil.WriteFile(l.cpF,cpdata,0644)
        if err != nil{
            return err
        }
        l.idx = 0
        for i:=0; i < LOGSIZE; i++{
            l.commits[i].Used = false
        }
    }
    return nil
}

func (l *LogStore)Close()(err error){
    err = syscall.Munmap(l.mmap)
    if err != nil{
        return err
    }
    return l.commitsFD.Close()
}

func (l *LogStore)GetLastCommit()(int64,int64){
    lastidx := l.idx - 1
    if lastidx < 0{
        lastidx += LOGSIZE
    }
    cmt := l.commits[lastidx]
    return cmt.Term, cmt.ID
}

func (l *LogStore)Marshal()([]byte,error){
    cpdata, err := ioutil.ReadFile(l.cpF)
    if err != nil{
        return cpdata,err
    }
    res := make([]byte,0,len(l.mmap)+len(cpdata))
    res = append(res,l.mmap...)
    res = append(res,cpdata...)
    return res,nil
}

func (l *LogStore)Unmarshal(data []byte)(error){
    logs := data[:len(l.mmap)]
    copy(l.mmap, logs)
    err := ioutil.WriteFile(l.cpF,data[len(l.mmap):],0644)
    if err != nil{
        return err
    }
    return l.recoverdb()
}
