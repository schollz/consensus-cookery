package meanrecipe

import (
	"fmt"
	"strings"
	"testing"

	log "github.com/cihub/seelog"
	"github.com/stretchr/testify/assert"
)

func TestIngredientLines(t *testing.T) {
	defer log.Flush()
	lines, err := GetIngredientLines("testing/chocolate_chips.gz")
	assert.Nil(t, err)
	fmt.Println(lines)
	assert.Equal(t, `-   1 cup nut butter of choice (i recommend chocolate hazelnut butter)`, strings.TrimSpace(lines[0]))
}

func TestIngredientLinesURL(t *testing.T) {
	url := "http://thepioneerwoman.com/cooking/knock-you-naked-brownies/"
	fname, err := DownloadOne(".", url)
	assert.Nil(t, err)
	lines, err := GetIngredientLines(fname)
	assert.Nil(t, err)
	fmt.Println(strings.Join(lines,
		"\n"))

}
func TestIngredientLinesURL2(t *testing.T) {
	url := "http://pureella.com/gluten-free-vegan-banana-bread-with-figs-and-walnuts/"
	fname, err := DownloadOne(".", url)
	assert.Nil(t, err)
	lines, err := GetIngredientLines(fname)
	assert.Nil(t, err)
	fmt.Println(strings.Join(lines, "\n"))
	for _, line := range lines {
		fmt.Println(parseIngredientFromLine(line))
	}
}
