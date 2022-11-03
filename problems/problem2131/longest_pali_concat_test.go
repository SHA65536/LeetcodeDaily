package problem2131

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected int
}

var Results = []Result{
	{[]string{"lc", "cl", "gg"}, 6},
	{[]string{"ab", "ty", "yt", "lc", "cl", "ab"}, 8},
	{[]string{"cc", "ll", "xx"}, 2},
	{[]string{"ca", "bl", "xy"}, 0},
	{[]string{"oi", "iz", "ym", "no", "tt", "jo", "fh", "pc", "cu", "pq", "bf", "ue", "eo", "dk", "hv", "yc", "ri", "jq", "ex", "xu", "yl", "bz", "od", "jd", "vl", "vm", "xy", "wv", "ci", "fi", "pr", "bf", "hj", "ki", "qj", "oy", "cd", "jy", "sw", "do", "cw", "ln", "oy", "sw", "hg", "pj", "fc", "gr", "vy", "dz", "pe", "xk", "ro", "mq", "xg", "fy", "fp", "mu", "ux", "sl", "py", "zh", "in", "ks", "pg", "hq", "sj", "lx", "jj", "kn", "kg", "ps", "vi", "yg", "yw", "rc", "kz", "rh", "hc", "my", "yl", "sk", "gy", "mq", "gg", "wl", "dr", "gi", "ka", "pq", "lz", "ei", "yu", "po", "td", "ir", "dl", "ib", "yl", "ug", "kp", "iz", "fj", "dn", "av", "jh", "gs", "nz", "qs", "ki", "lg", "rq", "vt", "pe", "cw", "il", "vh", "rw", "xp", "jd", "cs", "lz", "ii", "in", "gy", "bk", "vb", "jl", "rl", "ko", "bd", "wf", "ze", "um", "yt", "hl", "bf", "oz", "cp", "aj", "ot", "mt", "du", "xf", "er", "iv", "wq", "zt", "ou", "zy", "bx", "jw", "mc", "dn", "om", "fq", "ex", "bz", "lh", "ii", "rn", "nx", "iu", "mb", "dd", "iv", "wg", "ll", "ka", "ax", "zp", "go", "xg", "tp", "ei", "te", "pp", "qp", "fj", "jd", "kd", "tc", "of", "vw", "mi", "wr", "hn", "yw", "gm", "np", "fl", "ml", "ib", "oz", "xu", "is", "cq", "qa", "to", "pf", "tw", "nx", "im", "mh", "gy", "hh", "qz", "ab", "kk", "kn", "vu", "qi", "lz", "uj", "zg", "zc", "hr", "va", "qb", "kp", "pa", "du", "ma", "mr", "bu", "kj", "eo", "mu", "pq", "ou", "zf", "pk", "ok", "pe", "cn", "mw", "ck", "ru", "wi", "gd", "vn", "qn", "gg", "pu", "bb", "mq", "uk", "un", "pt", "zq", "yo", "bh", "xk", "ur", "bh", "xk", "ym", "hg", "dn", "di", "hc", "em", "zx", "ec", "cu", "xa", "dq", "uo", "ba", "px", "kb", "hh", "yo", "ah", "uy", "xc", "ko", "do", "wi", "mv", "ai", "ps", "vc", "ix", "gk", "va", "vt", "vx", "rx", "hw", "gq", "tg", "mi", "tl", "kg", "ld", "ah", "sn", "kj", "jz", "rb", "pj", "mb", "em", "dc", "sm", "jh", "gp", "wt", "gz", "rb", "bg", "gq", "lj", "yg", "nz", "wm", "nq", "oi", "fm", "pf", "im", "pe", "qz", "ia", "vt", "lh", "hx", "vb", "av", "vr", "jz", "ye", "cz", "rs", "qg", "ia", "vx", "lv", "eu", "rp", "jm", "io", "ls", "ot", "cu", "fg", "bm", "vt", "hf", "at", "qf", "gq", "wr", "kk", "gs", "df", "xo", "le", "cy", "rd", "rj", "rp", "yz", "uq", "gr", "aj", "oz", "vn", "pr", "ud", "fv", "xb", "jc", "oz", "ly", "og", "on", "he", "iz", "mp", "vd", "rv", "de", "ci", "qu", "ol", "zy", "bz", "ij", "js", "jz", "oo", "bs", "rv", "ov", "fz", "od", "qa", "fd", "ca", "an", "jt", "mg", "or", "dc", "qm", "by", "tv", "go", "zg", "ys", "am", "zx", "sa", "dd", "vk", "bp", "ac", "rx", "lu", "oz", "hu", "dh", "ub", "jm", "co", "cd", "rw", "cg", "qs", "ih", "vb", "si", "jz", "mm", "il", "zy", "wm", "wb", "mh", "vw", "pv", "vx", "pu", "ss", "xr", "yf", "qn", "ie", "ry", "zp", "oh", "sc", "ld", "iw", "oc", "um", "tl", "iz", "kc", "uv", "lt", "ea", "xf", "bt", "ds", "uy", "zp", "ux", "wi", "wv", "uq", "vj", "tm", "xb", "sv", "hv", "uo", "xu", "eq", "ul", "dn", "ra", "jz", "om", "tl", "mg", "pc", "ea", "dv", "tc", "mw", "uh", "dj", "bf", "do", "ux", "la", "en", "wj", "sg", "pa", "uf", "at", "oh", "um", "gz", "jb", "uc", "da", "wa", "dg", "rf", "mv", "og", "mk", "kt", "sj", "ks", "sh", "jz", "hj", "ia", "hw", "ty", "jn", "oz", "rq", "jx", "wl", "hr", "rv", "rk", "kz", "pr", "cx", "uw", "zg", "zu", "ne", "yj", "yx", "px", "yw", "ne", "mz", "ll", "yi", "ol", "cc", "ek", "vu", "na", "yc", "mf", "wx", "le", "pf", "nl", "wq", "cl", "au", "gv", "sx", "bg", "jr", "qx", "vb", "wy", "mx", "jg", "tr", "pt", "tn", "yz", "ep", "cm", "ml", "xa", "ta", "vc", "nw", "xp", "qn", "qp", "jq", "ul", "so", "fk", "dn", "fn", "gp", "jk", "cf", "ua", "si", "qp", "ij", "yv", "mu", "mg", "xi", "qy", "gj", "xz", "xe", "jh", "tk", "kg", "ih", "xn", "ss", "iv", "aq", "bl", "xw", "fr", "hi", "iq", "ux", "fi", "gg", "ae", "fr", "gs", "sn", "wq", "yl", "sr", "cp", "vq", "qh", "bn", "sa", "qu", "bb", "hn", "ky", "rs", "lu", "ct", "uh", "hx", "dr", "pt", "ra", "cu", "fh", "gw", "fe", "pz", "wy", "au", "kh", "wk", "xt", "ly", "za", "le", "cq", "lc", "ge", "ir", "vf", "ps", "qb", "rk", "vg", "gr", "sc", "lo", "xv", "nx", "rm", "ef", "nd", "mo", "yf", "ao", "io", "jy", "fv", "jm", "ah", "ox", "ni", "td", "dj", "fw", "xg", "ex", "lu", "xq", "xq", "gu", "nl", "tj", "qb", "sy", "bt", "kd", "wn", "dt", "cq", "an", "xs", "zm", "dw", "pm", "vx", "pp", "dv", "cn", "ev", "pb", "me", "yy", "nm", "gu", "zh", "si", "wj", "kn", "we", "mq", "fh", "lq", "er", "ba", "fo", "rq", "oc", "yt", "rs", "xo", "bt", "nh", "hy", "xk", "oe", "yg", "wy", "yl", "lk", "nv", "dd", "dt", "kc", "ic", "wd", "zd", "of", "nk", "zj", "fj", "nl", "jf", "xq", "un", "gp", "ki", "qj", "hi", "qm", "mz", "dq", "qo", "lx", "ui", "kx", "wd", "yh", "pv", "yv", "fz", "vc", "hv", "ku", "an", "pd", "zq", "bk", "hb", "kj", "xx", "jx", "vn", "gb", "sb", "ud", "ao", "sg", "xt", "ti", "ov", "kn", "ml", "us", "dl", "br", "fo", "qd", "fl", "si", "cs", "hf", "be", "zp", "qb", "sa", "wx", "uo", "ef", "ng", "ri", "rc", "rz", "rk", "oq", "rm", "px", "vo", "jp", "li", "en", "wy", "mo", "fb", "yh", "vd", "fe", "qp", "gn", "fs", "cp", "gq", "vq", "ri", "ig", "bh", "vs", "iy", "bp", "dy", "ae", "yr", "lb", "nj", "dc", "jk", "iq", "tl", "cd", "km", "nw", "im", "rs", "kr", "jx", "ny", "qk", "yi", "zy", "ov", "lb", "ll", "dc", "sj", "xx", "tu", "fn", "gk", "ay", "ui", "xa", "qq", "yc", "su", "yo", "gh", "gu", "vx", "xv", "an", "df", "mx", "xz", "ic", "om", "yr", "wv", "vp", "yc", "bg", "om", "kl", "ij", "ma", "iq", "hn", "gw", "ce", "rz", "va", "gx", "si", "vk", "ad", "ib", "jz", "bd", "ki", "ep", "ng", "bb", "ce", "ib", "vs", "qh", "qa", "mm", "br", "yt", "ro", "fn", "ry", "cf", "of", "so", "ak", "wl", "zo", "rw", "gi", "xj", "sp", "sn", "du", "yz", "ws", "kl", "he", "pv", "vs", "bm", "hx", "ex", "qz", "ht", "ip", "pk", "yq", "mg", "xz", "zu", "df", "rh", "jy", "nx", "na", "ci", "bs", "cg", "qu", "hc", "ro", "zs", "rd", "fc", "lx", "tq", "wy", "wr", "ox", "af", "ub", "wz", "gf", "ku", "lf", "bw", "fs", "lh", "vu", "tu", "bl", "ye", "ws", "pl", "bu", "tu", "ez", "pg", "fr", "ty", "eq", "un", "qn", "im"}, 1134},
}

func TestLongestPalindromeConcat(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestPalindrome(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkLongestPalindromeConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			longestPalindrome(res.Input)
		}
	}
}

func TestLongestPalindromeConcatNoMap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestPalindromeNoMap(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkLongestPalindromeConcatNoMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			longestPalindromeNoMap(res.Input)
		}
	}
}
