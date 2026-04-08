package upload

import (
	"blog_r/internal/global"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Local struct {
}

func randomName(base string) string {
	base = strings.TrimSpace(base)
	if base == "" {
		base = "file"
	}
	return fmt.Sprintf("%s_%d", base, time.Now().UnixNano())
}

func Allow(target string, array []string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	if index < len(array) && array[index] == target {
		return true
	}
	return false
}
func (l *Local) Uploads(file *multipart.FileHeader) (fileName, filePath string, err error) {
	extOriginal := filepath.Ext(file.Filename)
	base := strings.TrimSuffix(file.Filename, extOriginal)
	ext := strings.ToLower(extOriginal)

	maxBytes := int64(global.Conf.Upload.Size) * 1024 * 1024
	if file.Size > maxBytes {
		return "", "", fmt.Errorf("文件大小不能超过%dMB", global.Conf.Upload.Size)
	}

	allowArray := []string{".png", ".jpeg", ".jpg", ".gif", ".webp"}

	if !Allow(ext, allowArray) {
		return "", "", fmt.Errorf("图片类型不支持")
	}
	storePath := global.Conf.Upload.StorePath
	Path := global.Conf.Upload.Path
	filename := randomName(base) + ext

	DiskPath := filepath.Join(storePath, filename)
	urlPrefix := "/" + filepath.ToSlash(filepath.Join(Path, filename))

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()
	err = os.MkdirAll(storePath, 0755)
	if err != nil {
		return "", "", err
	}
	disk, err := os.Create(DiskPath)
	if err != nil {
		return "", "", err
	}

	defer disk.Close()
	if _, err := io.Copy(disk, src); err != nil {
		os.Remove(DiskPath)
		return "", "", err
	}

	return filename, urlPrefix, nil
}
