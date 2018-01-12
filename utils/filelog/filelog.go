// Copyright 2009 Igor Lesik. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filelog

import (
   //"io"
   "log" // https://golang.org/pkg/log/
   "os"  // https://golang.org/pkg/os/
)

type FileLog struct {
    logger *log.Logger
    file *os.File
}

func NewDefault(path string) (*FileLog, error) {
    return New(path, "", log.LstdFlags | log.Lshortfile)
}

func New(path string, prefix string, flags int) (*FileLog, error) {
    fl := new(FileLog)

    if file, err := os.Create(path); err != nil {
        return nil, err
    } else {
        fl.file = file
    }

    fl.logger = log.New(fl.file, prefix, flags)

    return fl, nil
}

func (fl FileLog) Close() {
    fl.logger.Println("LOG IS CLOSED. THE END.")
    fl.file.Close()
}

