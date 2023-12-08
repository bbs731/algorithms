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
	n := 24811
	cuts := []int{409, 8398, 9521, 15901, 13345, 12723, 15849, 23078, 9522, 16862, 2255, 21622, 8351, 9870, 8069, 10200, 21779, 17694, 11383, 2188, 16705, 13192, 1675, 6011, 2598, 22470, 8164, 2642, 3391, 596, 21537, 4668, 4524, 13209, 24249}
	//n := 7
	//cuts := []int{1, 3, 4, 5}

	fmt.Println(minCost(n, cuts))
}
