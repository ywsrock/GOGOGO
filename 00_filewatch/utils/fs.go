package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

type FS struct{}

var zlog *zap.Logger

func NewFS() (*FS, error) {
	var err error
	zlog, err = NewLogger()
	if err != nil {
		return nil, err
	}
	fs := &FS{}
	return fs, nil
}

// ディレクトリコピー処理
func (f *FS) DirCopy(src string, dst string) error {
	var err error

	var fileInfo os.FileInfo
	var dInfos []os.FileInfo

	//フォルダ情報
	if fileInfo, err = os.Stat(src); err != nil {
		zlog.Error(err.Error())
		return err
	}

	zlog.Info("フォルダ作成")
	if err = os.MkdirAll(dst, fileInfo.Mode()); err != nil {
		zlog.Error(err.Error())
		return err
	}

	zlog.Info("フォルダリード")
	if dInfos, err = ioutil.ReadDir(src); err != nil {
		zlog.Error(err.Error())
		return err
	}

	for _, info := range dInfos {
		zlog.Info(info.Name())
		srcf := filepath.Join(src, info.Name())
		dstf := filepath.Join(dst, info.Name())

		if info.IsDir() {
			//再帰
			err = f.DirCopy(srcf, dstf)
		} else {
			err = f.FileCopy(srcf, dstf)
		}
	}
	return err
}

// File コピー処理
func (f *FS) FileCopy(src string, dst string) error {
	zlog.Info("ファイルコピー")

	srcFile, err := os.Open(src)
	if err != nil {
		zlog.Error(fmts("open error"))
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		if os.IsExist(err) {
			os.Remove(dst)
		}
		zlog.Error("ファイル作成エラー")
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		zlog.Error(fmts(",", "ファイルコピー失敗", src, dst))
		return err
	}

	srcFileSts, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.Chmod(dst, srcFileSts.Mode()); err != nil {
		return err
	}
	return nil
}

func (f *FS) FileReName(fileName string) error {
	var err error
	if _, err = os.Stat(fileName); err != nil {
		zlog.Error(err.Error())
		return err
	}

	zlog.Info(fmts(":", "RemoveFileName", fileName))
	return os.RemoveAll(fileName)
}

//文字連結処理
func fmts(delims string, str ...string) string {
	sb := strings.Builder{}
	for _, s := range str {
		sb.WriteString(s)
		sb.WriteString(delims)
	}
	return sb.String()
}
