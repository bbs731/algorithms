package dp

import (
	"fmt"
	"testing"
)

func Test516(t *testing.T) {

	//fmt.Println(longestPalindromeSubseq("aabaaba"))
	//fmt.Println(longestPalindromeSubseq_dfs("aabaaba"))
}

//func Test1039(t *testing.T) {
//	fmt.Println(minScoreTriangulation([]int{2, 1, 4, 4}))
//}
//
//func Test132(t *testing.T) {
//	content := "fiefhgdcdcgfeibggchibffahiededbbegegdfibdbfdadfbdbceaadeceeefiheibahgececggaehbdcgebaigfacifhdbecbebfhiefchaaheiichgdbheacfbhfiaffaecicbegdgeiaiccghggdfggbebdaefcagihbdhhigdgbghbahhhdagbdaefeccfiaifffcfehfcdiiieibadcedibbedgfegibefagfccahfcbegdfdhhdgfhgbchiaieehdgdabhidhfeecgfiibediiafacagigbhchcdhbaigdcedggehhgdhedaebchcafcdehcffdiagcafcgiidhdhedgaaegdchibhdaegdfdaiiidcihifbfidechicighbcbgibadbabieaafgeagfhebfaheaeeibagdfhadifafghbfihehgcgggffgbfccgafigieadfehieafaehaggeeaaaehggffccddchibegfhdfafhadgeieggiigacbfgcagigbhbhefcadafhafdiegahbhccidbeeagcgebehheebfaechceefdiafgeddhdfcadfdafbhiifigcbddahbabbeedidhaieagheihhgffbfbiacgdaifbedaegbhigghfeiahcdieghhdabdggfcgbafgibiifdeefcbegcfcdihaeacihgdchihdadifeifdgecbchgdgdcifedacfddhhbcagaicbebbiadgbddcbagbafeadhddaeebdgdebafabghcabdhdgieiahggddigefddccfccibifgbfcdccghgceigdfdbghdihechfabhbacifgbiiiihcgifhdbhfcaiefhccibebcahidachfabicbdabibiachahggffiibbgchbidfbbhfcicfafgcagaaadbacddfiigdiiffh"
//	fmt.Println(minCut(content))
//}

func TestMaxSubPalindrom(t *testing.T) {
	//word := "faaadacb"
	word := "aa"

	p := maxSubsequencePalindrom(word)

	fmt.Printf("max palindorme seq len of %s , is :%d\n", word, p[0][len(word)-1])
}
func Test1771(t *testing.T) {
	//word1 := "afaaadacb"
	//word2 := "ca"
	word1 := "afaaadacb"
	word2 := "ca"
	fmt.Println(longestPalindrome(word1, word2))
}

func Test1547(t *testing.T) {
	n := 2
	cuts := []int{3, 2, 4, 1}
	//n := 7
	//cuts := []int{1, 3, 4, 5}

	fmt.Println(mergeStones_dfs2(cuts, n))
}
