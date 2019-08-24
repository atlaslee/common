/*
MIT License

Copyright (c) 2019 Atlas Lee, 4859345@qq.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package common

import (
	"strconv"
	"strings"
	"unsafe"
)

const (
	SIZEOF_VERSION = 3
)

// 版本
type Version struct {
	Main, Milestone, Minor byte
}

func atou8(str string) byte {
	i, err := strconv.Atoi(str)
	if err != nil {
		i = 0
	}
	return byte(i)
}

func u8toa(u byte) string {
	return strconv.Itoa(int(u))
}

// 与另一个版本比较
// 大于0则本版本更新
// 等于0则两者版本一致
// 小于0则本版本更旧
func (this *Version) Cmp(v *Version) (i int) {
	i = int(this.Main) - int(v.Main)
	if i != 0 {
		return
	}

	i = int(this.Milestone) - int(v.Milestone)
	if i != 0 {
		return
	}

	i = int(this.Minor) - int(v.Minor)
	return
}

func (this *Version) Newer(v *Version) bool {
	return this.Cmp(v) > 0
}

func (this *Version) NotNewer(v *Version) bool {
	return this.Cmp(v) <= 0
}

func (this *Version) Older(v *Version) bool {
	return this.Cmp(v) < 0
}

func (this *Version) NotOlder(v *Version) bool {
	return this.Cmp(v) >= 0
}

func (this *Version) Equal(v *Version) bool {
	return this.Cmp(v) == 0
}

func (this *Version) Bytes() []byte {
	return (*[SIZEOF_VERSION]byte)(unsafe.Pointer(this))[:]
}

func (this *Version) SetBytes(bytes []byte) *Version {
	copy(this.Bytes(), bytes)
	return this
}

func (this *Version) MainMilestone() string {
	return "v" + u8toa(this.Main) + "." + u8toa(this.Milestone)
}

func (this *Version) String() string {
	return "v" + u8toa(this.Main) + "." + u8toa(this.Milestone) + "." + u8toa(this.Minor)
}

func (this *Version) SetString(str string) *Version {
	strs := strings.Split(str, ".")
	num := len(strs)

	if num > 0 {
		this.Main = atou8(strings.Trim(strs[0], "v"))
	} else {
		return this
	}

	if num > 1 {
		this.Milestone = atou8(strs[1])
	} else {
		return this
	}

	if num > 2 {
		this.Minor = atou8(strs[2])
	}

	return this
}

func VersionNew(b ...byte) (ver *Version) {
	ver = &Version{}
	copy((*(*[SIZEOF_VERSION]byte)(unsafe.Pointer(ver)))[:], b)
	return
}
