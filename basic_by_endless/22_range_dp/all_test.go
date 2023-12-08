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

	fmt.Printf("max palindorme seq len of %s , is :%d\n", word, p[0][len(word)])
}
func Test1771(t *testing.T) {
	//word1 := "afaaadacb"
	//word2 := "ca"
	word1 := "afaaadacb"
	word2 := "ca"
	fmt.Println(longestPalindrome(word1, word2))
}
