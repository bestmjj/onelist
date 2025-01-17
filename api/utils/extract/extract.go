package extract

import (
	"errors"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func removeEndingOne(s string) string {
	// 去除字符串默认的数字
	//\d \d 是一个简写字符类，匹配任何数字字符（等价于 [0-9]）。
	//+ + 是一个量词，表示前面的模式（即 \d）必须出现一次或多次。
	// $ $ 是一个锚点，表示字符串的结尾。
	// \d+$ 匹配一个或多个数字，并且这些数字必须出现在字符串的末尾
	re := regexp.MustCompile(`\d+$`)
	return re.ReplaceAllString(s, "")
}

// 过滤电影文件名

func ExtractMovieName(s string) string {
	oldName := s
	// 去掉扩展名
	// \. 这是一个转义字符。.在正则表达式中有特殊含义（匹配任意单个字符），因此需要用反斜杠 \ 来转义它，使其表示一个实际的点号 .
	// [^.] 方括号 [] 表示一个字符集, ^ 在方括号内表示否定，即匹配不在这个集合中的任何字符,因此 [^.] 匹配除点号以外的任意字符
	// + 量词 + 表示前面的模式（即 [^.]）必须出现一次或多次
	// $ 锚点 $ 表示字符串的结尾
	// \.[^.]+$ 匹配从最后一个点号开始直到字符串末尾的所有字符
	reExt := regexp.MustCompile(`\.[^.]+$`)
	s = reExt.ReplaceAllString(s, "")
	//fmt.Println("After removing extension:", s)

	// 去掉(前面得空白符号以及(xxxx)的内容
	// \s* \s 匹配任何空白字符（包括空格、制表符、换行符等）, * 是一个量词，表示前面的模式（即 \s）可以出现零次或多次
	// \( 这是一个转义字符。( 在正则表达式中有特殊含义（用于分组），因此需要用反斜杠 \ 来转义它，使其表示一个实际的左括号 (
	// [^)]+ 方括号 [] 表示一个字符集, ^ 在方括号内表示否定，即匹配不在这个集合中的任何字符。因此 [^)] 匹配除右括号以外的任意字符。+ 是一个量词，表示前面的模式（即 [^)]）必须出现一次或多次
	// \) 这也是一个转义字符。) 在正则表达式中有特殊含义（用于分组），因此需要用反斜杠 \ 来转义它，使其表示一个实际的右括号 )
	// \s*\([^)]+\) 匹配从零个或多个空白字符开始，紧接着是一个左括号 (，然后是零个或多个非右括号字符，最后是一个右括号 )
	reBracket := regexp.MustCompile(`\s*\([^)]+\)`)
	s = reBracket.ReplaceAllString(s, "")
	//fmt.Println("After removing brackets:", s)

	// 去除年份
	// \b \b 是一个单词边界（word boundary）锚点。它匹配一个位置，而不是字符。这个位置位于一个单词字符（字母、数字或下划线）和非单词字符之间。
	// (\d{4}) \d 匹配任何数字字符（等价于 [0-9]）。{4} 是一个量词，表示前面的模式（即 \d）必须出现恰好四次。() 是捕获组，用于将匹配的部分提取出来。在这个例子中，捕获组将匹配四位数字。
	// \s*  \s 匹配任何空白字符（包括空格、制表符、换行符等）。* 是一个量词，表示前面的模式（即 \s）可以出现零次或多次。
	// $ $ 是一个锚点，表示字符串的结尾
	// \b(\d{4})\s*$ 匹配从一个单词边界开始的四位数字，并且这些数字后面可以跟随零个或多个空白字符，直到字符串的末尾
	reYear := regexp.MustCompile(`\b(\d{4})\s*$`)
	s = reYear.ReplaceAllString(s, "")
	//fmt.Println("After removing year:", s)

	// 保留中文的剧集名
	// [^\x00-\x7F]  方括号 [] 表示一个字符集。^ 在方括号内表示否定，即匹配不在这个集合中的任何字符。\x00-\x7F 表示从 \x00 到 \x7F 的字符范围，这是ASCII字符集的范围。因此，[^\x00-\x7F] 匹配所有非ASCII字符，通常这些字符包括中文字符、日文字符、韩文字符等
	//+  量词 + 表示前面的模式（即 [^\x00-\x7F]）必须出现一次或多次
	// [\d]* \d 匹配任何数字字符（等价于 [0-9]）。
	// * 是一个量词，表示前面的模式（即 \d）可以出现零次或多次
	// [^\x00-\x7F]+[\d]* 匹配一个或多个非ASCII字符，后面可以跟随零个或多个数字
	reChinese := regexp.MustCompile(`[^\x00-\x7F]+[\d]*`)
	matches := reChinese.FindAllString(s, -1)
	//fmt.Println("xx1:", oldName, s)
	if len(matches) > 0 {
		name := strings.Join(matches, "")
		//fmt.Println("xx2:", name)
		return name
	}

	return strings.TrimSpace(oldName)
}

// 根据文件名获取剧集季及集信息
func ExtractNumberWithFile(file string) (int, int, error) {
	p, err := filepath.Abs(file)
	if err != nil {
		return 0, 0, err
	}
	SeasonNumber := 0
	EpisodeNumber := 0
	fileName := filepath.Base(p)
	re := regexp.MustCompile(`[Ss](\d{1,2})[Ee](\d{1,4})`)
	match := re.FindStringSubmatch(fileName)
	if len(match) < 3 {
		return 0, 0, errors.New("get number error")
	}
	season := match[1]
	episode := match[2]
	SeasonNumber, err = strconv.Atoi(season)
	if err != nil {
		return 0, 0, err
	}
	EpisodeNumber, err = strconv.Atoi(episode)
	if err != nil {
		return 0, 0, err
	}
	return SeasonNumber, EpisodeNumber, nil
}
