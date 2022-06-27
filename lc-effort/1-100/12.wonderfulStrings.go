package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/number-of-wonderful-substrings/

/*
	wonderful substring
	so brute force solution would be two loops
	for i:=range s
		for j=:i+1;j<=len(s)+1;j++
			substr = s[i:j]
			if substr is wonderful
				count++

	even I know this is slow but let me do this first anyway then think what is the better solution
*/

func isStrwWonderful(s string) bool {
	m := make(map[rune]int)

	for _, c := range s {
		if _, found := m[c]; found {
			delete(m, c)
		} else {
			m[c] = 1
		}

	}

	return len(m) <= 1
}

func _1_wonderfulSubstrings(word string) int64 {
	count := int64(0)
	for i := 0; i < len(word); i++ {
		for j := i + 1; j <= len(word); j++ {
			if isStrwWonderful(word[i:j]) {
				count++
			}
		}
	}

	return count
}

/*
	43 / 88 test cases passed.
	Status: Time Limit Exceeded
	Submitted: 0 minutes ago
	Last executed input:
	"ghffghcdigdgbfhaiejibbcidgfjahegedjecfhdjchdcaefeefiabebefjccdhdchacehfdjibcgcedjhhbadfgejhceghgjjegihhbfagciecdadhcjgdajgbhegjhicbifeefcaficcbcijciegihidacjahahjgdjdgfjcadgaiicgiiegaeifjbcjacefffgdgibjfijiiedcaahbceebheifbgffchdcfbbhfdfajggfhhhbhhgbcgfdgcfcdggihhcfcghbgbchcgaccccaabeejdhdbjicghddaaeedciigeeiehjhhhicajbddedidecjiebbggdehbgejabcbcjadicjgbabehbjjbehfegicabighbcahijjdcgehjhcbfheecgjdfbhadgbhfdefffbcbggifgdhdgahgejeegbadbifjdiaggdcjfaiigdjjehebjdabgdecfcebcfggceihffhjcjfaadhgeahiajjdghehieiidcfdjaeiehcjghjcajhgdhcfjcagiibhdiiabghcffaijehjjjiafjcgfjeghdbggbhfjdihahjedgdhhdibcdgaaafejgbdjgdfichbijjajicdfehafjbgdbieaecagdifeggbefegeajejbcfaidjbgbbgejjebidcdaeggcefjdfacachfeehajfabfhbcdjbcjdifiehdjghdhegfgababaeacdjcegffgdhadjehiebggdiececeiihfghcecahbhgefhihahhfhgdeaejgdhihcgfecgecebdjgcjicjdbcicigdbaehhjcdajbehdeahbibeejddajiacgchbcffagbbdiechcfgafehifibifhaidaeidjbeefefcaiejdhibebghbjgbbhchfjeefdfgjiijbbbgchgjhaefdgejfgeegbajgbcjejahgchfhjbgiiggjiaigiibccidcjcgbabed"

	of course, O(n^3) now let me think the time complexity
	so quickly... I notice the relations between str and its substr

	for a str C[0]..C[n-1] with m as its wonderful map... the map I built in the isStrWonderful, didn't think it can valuable

	so count(str) = 1 if isWonderful(str, m)
						+ 1 if isWonderful(str[1:n], m-right) # right substr
						+ 1 if isWonderful(str[0:n], m-left) # left substr

		m needs be pre-calculated
		m-right can be deduced from m
		m-left can be deduced from m
		and use m to tell if the str is wonderful or not

		there could be some duplicated strs.. so memorization can help a bit

*/

type PosPair struct {
	start int
	end   int
}

var pospairCount map[PosPair]int64
var strCount map[string]int64

func countWonders(start, end int, s string, m map[rune]int) int64 {

	if start >= end {
		return 0
	}

	if _, found := pospairCount[PosPair{start, end}]; found {
		return 0
	}

	c := int64(0)
	if len(m) <= 1 {
		c += 1
	}

	mR := make(map[rune]int)
	mL := make(map[rune]int)
	for k, v := range m {
		mR[k] = v
		mL[k] = v
	}

	leftRune := rune(s[start])
	if _, found := mR[leftRune]; found {
		delete(mR, leftRune)
	} else {
		mR[leftRune] = 1
	}
	c += countWonders(start+1, end, s, mR)

	rightRune := rune(s[end-1])
	if _, found := mL[rightRune]; found {
		delete(mL, rightRune)
	} else {
		mL[rightRune] = 1
	}
	c += countWonders(start, end-1, s, mL)

	pospairCount[PosPair{start, end}] = c
	strCount[s[start:end]] = c
	return c
}

func _2_wonderfulSubstrings(word string) int64 {
	m := make(map[rune]int)
	pospairCount = make(map[PosPair]int64)
	strCount = make(map[string]int64)

	for _, c := range word {
		if _, found := m[c]; found {
			delete(m, c)
		} else {
			m[c] = 1
		}

	}

	return countWonders(0, len(word), word, m)
}

/*
	Submission Detail
	55 / 88 test cases passed.
	Status: Time Limit Exceeded
	Submitted: 0 minutes ago
	Last executed input:
	"faecdbcjjbjccjgcadeeaeddbibjiiecggbchebffdchbdicacdbcgdchdgeciceghcdhfbdhcaagghhhjichcgfdaijicahcihgdejhdafjghhjjeaeajdbiheegghhbfgdidcbbfgbdihedfgghjejafdbeedadcagcfejaaihihfdgdedibiiaaaajeifcfcjeeahhcgidcigjhafibejcadaegbeejbficdjgcjfjcgefeifcijibdhaaccdigjeafcebjaiciicghbheifebdadhdeceidcaijihhecjcfcdagdjhfabejjaijajdeaehfghehggjhaidibdidggfghejbggcadhfgjfaeehhicbgjeibbhbijhfghiejadefebigahfeheageagejehgbfajjadicggiegbcbeigdjbcijagajgfaihedcfcjgibiicbjegjgadahdeefajjbajdhcchdeiediecefacdabeegedhcfbiefjagacbfbgbdhejieffhjhcifijjfbccdfgdhhjfbdcejefedgihdjcahaggcjdfafciecibacijddbebccefchjbjfhgeghbhgaacehebafbdghdjdjdbiajeaiiccifcdcjhiaadchfagjdbhfhhiddhaehcjeafediehghbjaafiicgfejadcahfdgagfadbaiehaeaaiifcbcegihfbiacaadjjgahibbfijdbhbcfhhgccihdahchdeeahjeabeiegdgjfdgdfabfiajhagfdhhjaafacaihiagacgijfjdfafehicbeaicafciegfbefehddbfbgfhdcdchabgdigbceeiegabggeiafejbbdgjgijdfjiibadgfjehhaehjdhcdghadeejdfaicjdajjbiachaiadcghjfbefhgdfagbjibjfhegcbeiiecaaifiidgidbcbgidfgibhjcgjcifaegfgbifcjdgddadgedhbeidgbagiddfbaedhjefgjgfdbcabjcigahfbjcjhdhccjjhhffjcbhjhaeaiifajhjcfaejehjabeeggfaebjeiaacaabijjbeiggjfbdfcfgjcebbehbhcjjggaefgbdhfaigaagfafegbbahfiaabijcffbjddjefcfhhdbbgdjiiabdeaihfccbbfdaifcjhidfijdjdbaebaibfcijhhcddachacfbgddfejeafeddjabbdgcfibgijbcjjijhagddidhhaijiafbaajgejihadjdcdjagbfghjabdhhbbicgceifddgbjeceejbigjfgdedeejbjhdibjghigegjigghghhedhdadgedgbhebedjgjbjjhfedgbhhaebgaeebiajijdhdcacffbfedcbdebidbhieijdcbcdfjgdejehbgecgjcdcfhihgajjggffbaaheiabbbebjcbhdhaajbhicibgajddceaaacgaighhbaacbcbdbdbfacgiaaaaaiedfhcgfgfhhgeeidhgghfbbaadfegchahjcebacbhibbbjbhechefigefhcdcagjjhbhfggjgejfaceeebfibaegfchdabdjgfdjjcacgiifcdcigcfjbegddafefaiaihcbdgdccgiijejdadehgibbagjeajafaggehdajahfcbadjibdgjjcigccbdeiadijejhcjiihcbjijgabhcjfgaiajjigefiddhjcbjjhijbbgbjjeffedffdefbdbhifcidajichgcdchafcfhffifgcfedgafgdhbcedafabhgaejbhdejadiagibiijbdeeggejigadcedbcfbbbcijhjdhdhjceecjjbbfbegegahefigfeahgjbeddaedeaefeehhdhbcdiafhediejdbgaicahgeedfcffjagfgccbbfaajadaeijegfffjiidhgiiadiecbhifeicigjegijecgdjihfbjabigfhbjeejbfghfddehechdjffdbadieibjdddeaejhgijhhihejjcibbaaagbihcagifcjhebhgcjdejgabjjabgbfbbigbjccggbegfcfebfbcjaffhgcjfadaeggcajbgegdejhbfjiehbfgacdgaedeicigeejjgfijjjcgdjaeiadgjhgcdjfgdeagfcfdaicffifgbfdjbebeeadcdbdihjfijiiehbchhbdbbjgcgaejcefacjahcgaeajajejehhbdadihbcjfihaejjbabegbiehjhiccdacjaacdgcbfabgfeifchacbegffhccdgihbbbdcjihibjbjjdbgcfjcehedfjfahajbcfbfdhcgbjdcacbbafejafdciideecfacdhcjdbaigcieghigjfhacfeaifjifchdiggijcfaihhbaifcfjiafhffebbdcfhhacgadhcdcfcieifiifdjhhgjjiicbjjjijahcjijehbghfbbigiaieheiehbdfagijaghhjbfgbbdiajiadbddhbcidahiejehcgcjdegfidhhhcebgfhcifagjdiigicdedfjcifdbcfejifcihjgjeggbfifgbeeejcjfadadafghhhgaifegcajffcceedaejbdidhjajhabigfbaegjiejgaihhahecfccefcfgdhgaaabeiibibgdcgdeajbbjceififfdebdjcigehdgdjbgajhjhigjifejjccheceabcgbafbibiibgiicefefhiceigjhijjafdeiafffgcacjdccadcbecieegfgbfdejiafdhdhhfcjhbfjbehdfdfgcbjijiiddcafaaedcacedfebcafhcibcghhcieihdjdgbeaddgdbjheedjgihffbcbiaccbcgibjgcfagjgcbhhbbaafagefgbfcfdfajecfdhdfjifbffcacfbeccdbfgieegefjdhgdffcfjfajafhggfdjhhfhabgdgfcjhjgjjgicahghegdjddabdiechjdcbaichjijbbgehfgadibgajgdebfgahiidibahabgbcbhfddgefjgjbhbjgjdhhfbccheibbfcebjciafgbgijchchedeahagbcbegjehjgfejhcgihcbiheggeehjcigbjebddcbbedfjieijjjbdafbcdhajchgbbhebeijgheaghjhfidceiediagjhggjchefdjfhbadihaaggdchfhfbdifbibbjahbfadihefeaedgcedgdffdjecdhbhfjebihfgfbcfgieedhcijcfgaiifedebbefbgbgaihiafgafbchggdbceeiiadcbgfbfedihbjidicifdjbggebcdcjbgjfabadfgifiiehfffdfbfidghdhjieiifehiccabechjhggfbcaeccghdchccjhchgcjhdehjjecfbfbgjbabjjdjcgadfaffiicgabheadagadhfchibdbdbfeghdhhcedghdfahcgcihaaaeabfbghijgbbbfjaejhidbaaefgeiihdgdfaeijihfifjdcdjjbjfifdbhdhgeiedfcfcecjeecjcaccjjedefabaaghfbcgdagbdbjggjgjgijcijbbbeddiahahhcbjieabehcjgajcihhdjbgjibbiidigjgifhdididaifcfadcghgfceidhicfegaacideifeighghaggdacdhaiaaffafgchgaffhfjhaddgiaigjefbeabgfeaaibabffiiadgfgcddhfiijahcdccichdeebajhfacajgcihhjccajfdehbaehbdfebcaafiahibjbchcfggicbddbhhbbcfghdehjijhggiidagebfdhaifeghajabbejcgcbigifbedgaagdcccagbeccifgfbdhaidddbdehcafbefcifjjiiibb"

	better still cannot pass
	bring memorization on for str-to-count map

	gosh so many maps...
	still here... 55/88

	so think further
	now the observation on the a-to-j only chars... only 10
	then I peaked at hint 1

	For each prefix of the string, check which characters are of even frequency and which are not and represent it by a bitmask.

	aha... bit map.. it will be much quicker..
	okay... admit I am little dizzy right now.. let me see other people's solution


*/

func calcBitmap(word string) int {
	bitmap := 0
	for _, c := range word {
		bitmap ^= 1 << int(c-'a')
	}

	return bitmap
}

func _3_wonderfulSubstrings(word string) int64 {

	count := int64(0)

	for k := 0; k < len(word); k++ {
		bitmap := 0
		for i := k; i < len(word); i++ {
			bitmap ^= 1 << int(word[i]-'a')

			if bitmap == 0 || int(math.Pow(2, float64(int(math.Log2(float64(bitmap)))))) == bitmap {
				// if bitmap is 0: all even freqeuncies
				// if bitmap has only 1: it will be wonderful too
				// how to tell if it has only 1, it will be 1/2/4/8/16....
				// log2(bitmap) will be an integer so pow2 back they should equal
				// log2(0b1010) will be 3 but pow(2,3) will be 8
				count++

			}

		}
	}

	return count
}

/*
57 / 88 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
"iibhicfjiejfcejhaiebdcefjdcjfafcdcaiecddhfjaffidihdcbacfgjacbdjddfcjhaibhebfijbabheifjeejfgbgccfgibabbejaieafiejibghicdfhbcgjceejgijiebjdbbhhicdjggegacdcbiaacdijgbhdjacdbddgbfccfhiiddjadhebagedhjcfgaajcdfedfjfbhaaehedfggeiddjcccefbbfjejfcjbgcjcchijigfgadfcgfjchjadfihbdbdibbbbdgjiidhajggbhdfgccgegbgehibahfiebiefhdghbbabagjbdfieehdcjchadjghbbgcdgfgjifeggicefaciiigehiehdchcjfdghfbbfdihjhajgbddggjfafjdcbfhhgbaididcaadcafhabgdabdgedafhhehbjbefedjjfejbbbhaeeceejifdecbjhjihafbfigbjhgedaeiebjadjbdhbhhhhabchiedbieabeegfabijfeccbdjefggajiejhigdhbfjbifjchhbdfffbddgcjadbaahbagedacefjfchadjdabcaehahbjahfchbcgchjbddjcbebjcgjabiiibdeaedajdhiagebaichdciiggedbigeibgadjbgabafjiifbhceaaiefcdgfcbcahhhicegfiiggbbabfabajeeddgfdeafccibaihahcbifeeeabdicihhahfjjidhbfecgehhggcebjbcbihccfdejfjdjcbfegiajfciijdfdfaecbbccdbhechgeedbigefdjhcdbghgafebahaceghgbaheheigjhgchhjafjdacegfcjbidbajejhghfdccebiaaehcegbadgbicjcbhjbdfigceajhbaafgebfaabfjibedeghicbdgeeeccbadebdcffhcafdcfgcdiiheidbaadaiifefeijhhebiebgahjbefejbiebebcjjacifcbaagdbddigbacdfeidajefahbfcidiaffbaaaddhgcjbcebaafcbechbaagidgeegcaahihedaheacbahbbeigchgdhachcffdbjejehebbahcfdaabiefecabdffgfdbjbeehfgabebifcjdajjgbfdihbagbidbefdjdbedchhdaigghehfdfajcefecdhdghhbdicidedeeicciehcjecehfjiaebecdaefdgdjahgfdjgbegiaejcghjcabcifbaigjbdgfiaaffgddeccedcgecfeeiceajhcfejhecdcfdhjjeiafgbbebfdacjcebjebahgbbbajedeidbaeeijafdaidgccfdcdgjeeibdbdgejgbjeegegjjgabdgjghicfeecdhegfhdhhbihjefhfadiigefefjcdijhcheehbggbhhcgchchafgdgedjdhdiijcieaeddbhedfeidhbifjddgahjafjfaiehabfhejhjgigfhcfdjfdfhggjeaiijeiajgghjijjicibfedafgciffgifbeajgddbejffhcdafchbecabiecedaehggaacabhgfeihaabehggbfjcgajgjfceaifcigcfhbbbdcieiihadabhhfaejjbdffhabehhefadcaabdffeggfhgcgifcgaceacccgaaiagaafcdijdggeaeahafhjichecgddbiicgdjbaajcciehfihdgdihiicehdafdadegabbifddggahgbiedfdbhfebiggdcijajgedefjbfcjedijjbagcifjfcacjfhahjjbdgfaifehfhhbfbceecccigfbfjhcgbdgjhadfcjicdaggjbhfaaeahihiehaabbgfebcdgdbihbhfheajedigfigeebdhigaicgbgbhjgigddfagieachdfijbbbcigebhhheaiicjdaabjaidjjibcjdjbbgccecdbdfdachaeigdeihbdijceggjccgfcigcefgebehbhfdjccjadfhdeejgaddhdfihjcbibbjcabbcgddfffgcbhehjggjafdjddcebfajjccfhedbiecbbbjbceiifcfddiaiffacedgfgjbjhcbahjeedfghbjijcceedcdgaacccgdaichjhaffcfcfijjhchfiecgbajjjgacifffajhadgaicgdhdffijhjeaddgdjeagdihjbfiffdbeheadbaicgcbhjgbddhbcagehffgecijdjejacafheheefhjfeihigjafddccgbcgjaehgjefidbhbifedgejcjhddfbfchegjagbagafihfceihijidejgegcfcdejfddchdfcdhhfifjefifaddcggaheiiicahjihhehjbiebiiaegbbbgfehiifgiddgdfgghjiehddgfgdgigbfeaefccfegajiaegaheffaiiejacjiifhgjcahbjdidiffbbdhjdifihfbhjaeehghbccjbahbeifabcjiibejcjcdgceiigjcacjbicafadbejiecchfbajjhadcfhgbcbfeccjiiiejjhijdhidbadaagbihaeafiafibejcjdhbgibiagdhijhjeefeccdfaheejcdiiadfjeeaejeaaidiggjghggagcgcbffehadafhgegfcciegjbdiabagcffbfadecjaaeifgfidcfbhhibgbhfggjfgccdccabfaacbgjfijbgaaejjbhgifedcbafddiabfhehdabdjjiadgfjdjeaecahhhgijdcfaajjjeegcfidfbbggcbbafeddjbieeahfaccjdfhjhdhddbfhhafagddhhgbgbhcfiefdhfdgjbbbjdddcdfjdagcichfijedffigeijjgchdiihadagjbgidcegfhaeaiheaacjhfhbdbfgajagececefddcadfbhihgjciaddgjiijbadehiddaighheigfghggcgagehjiaahcgcfecheddehcijffgiidjjcfidjeifaicbaaiddjedhcihhabcidggjeaceehiiieecigfcibcffjdjbbbgieefhdieddjchegicigajdebieiiabdbadfdcjigegdcgdjbdbggjidajhbfcigihdefgfcfjidhdeejahfachiaaiiecfcgdbcdiiigbgbfbghfgfifhbeehegigfchafjbjgieeiechdbidcijfjeciechidfcfgfciaaafcbdgbhdbgdcijeddfbgchbgjeeggidiigaijacejijfiejbhebicagcgjdbjefefegejdachfjjeaifeacadfcffeegbidaieicedjjihdigbjffcfedchhfhjcgjeceiaajgecifgdahcacjchjgafcfibjicegfahdgicbagaedacgagcijddgbbhiafeiaijcidebfbhjgaegdajjgiiediaejijfcdbjieaficdbddjgheageiffjfeggdghheijbdcbcfcicaccdhhcehfbiafcgdajeefjajijgbefjebehaffjhfgjifhheadgijijgfefbgifbabjjbdfgfbhcgdgdcjcgdfjacebdhcacjfchhhegedceggdjjdaigbbcajhiahffjbiajeahgibjecfdbffgjgeehacggjiejfhhjchcghicgggdbhacdfdadiggdcfifcbfaaecihagdeffhdhgehbjbabdadcgicdhjeecibeicjbbhahebjjgjjjjbhfcibcgehjbbgiddgciiaaaijadbaechjdjbfaiffbabjhggjecgggbdihifibadbafhfcbdaddaccjigjcdgghbbbcajccifecfbabaedaiaccdhihbffdegahhhaaiidhdijidfghcdddfgfihfdecijhjdiejjbhjjbeghhhejafhegjjhbecjgbcdbcjcgcdfedcehhecfcefbhagcchegabacedfdcagfbiebhfajhdedgfiebihdbfachfaaheichaffcdeefhacdiafdddhfhbbjjbifjegcefjegfbjdgfedfcbcddfeajjdadbhajjjacgahcjeeibiedhgedhccaiadadjcjjbaddiefbcegcgcjbegfadijicgjcedcjhhbgjijgjhbhhhjifdhddhgcefjjjihebffdfejjaajicfcdfdabjbfefhjcjahidagbchddibjdgdcjfhhagfacfghdeagddbieeeifhbdjjdgafcfdbbgfadeigigefijjghggbecaidgaagefefcebgebeaajadjcfcfhgcdjcbbiegdhchcaheijcggajjbjfcicahafjbbbahfeecejibgbfjeichffcdhfhfefciahjcbddbhaaijjfjchhbjicdhcccdccifdgfjehhjdecdcagijdddhggfhfbgagfefdfbabcdbiaeeiibgabfjheaajccfjjhjbijchafjfibdedegchdcbhieaiccegbacgdihdjigdcagaffhjghhejceaeaeiaadfhbgbdbdgjccjiacjijfjbecedeahcfbabjhiidihcjbgeefdiabgaiffabiajifcigfjgehcbjahbefjcadijhhhbgdfdhhdgbabjagfbcjfagffjdgaadfjibedjcjfhjbccdhbfdeaegefachbebjbiijcdeddbeiacdigbgbfaaiajhcccagehadjiibjfejdeacihhbgajcagbacabhcfdeifeiaegfgieahbeifccahheifjjbfejgiadjbifebbaidaaaiiddibhaidbabhgfeijdafjafhcfbhdhjgdigddaeheheecbiighabcabjgbgbaiiccjgcgejafdhaaiadcdebegjaidhghdjahbihabaifcdjhhaicecaibiciejchiajfiegddejaehfbgcfaadegcjhjheecigjdfhiehcihjdiaiehcgcgdabcfedbcgbbecghidcejecbeffddhhfbidhbhdcfhcbcjbefdahifihfhdhgjbiaejfafjfijdbiagfiijjhaiadehjgghcaajjefjddhbdhcbjahhaehiddffcfbdeggegdfeieafigehhjccbjijegbcaahbejhahaiacghibhdjbiefgihbgecdajdfdjfdbjjceajdadhfgdfjibeeajibhijbfegfhiijjggcbcagffbcgieeechbghgbaagagehbghafhbideahghchfdjhfdbiebhigjedjibdcjfeedbgjbgcibbbhdigihjjcegcebejahefifigebefcdgegeabbcegffcdbafijfbabefecgafjdfhbiggeffbcfbjiciceebhgbeahghihahgefigdfjbcdbjgiabcegdhhaiiiabcbcdhcjccgggbadefieagcaegbdddcihciagagjeafdafeedjegaccgghifbhedagdgdidbgfcjadegcaebbbddaebjifbidfefbgbajdfacecaddgfebdiaagijedhjjfiggjcbjhhbciigfdecfebdddejjiafbbfchdbaehfaeaefbfjbcebheahdjafccaibbfiagfccidjhbicfachbahcecbgdieideaajhcfcgifebedeffdiibhceidaihjejjdacccgbfeggghgebbebeadjahbhcjidjfcibeaehhbbbdfchbhjbichdfghgbejadffadhbjicijdjeidhbfbaacbjbiegifdfbeajajaejiafhfghjegbjafdgfiggjecaeehaicegcbfghfeeechaebdjbihdacjdebjjaigdajdbghhbjiigffedfdadfchcibaibgdfhhchgcaigbhbefgdcafhhghiadifedagedjbgcfcbigeafjffajhcdddchiidhhgdjijgcccgejcgdfagbceeeiidcbdfaigdeiafegihdjgdiijbjifihaifbddibdjbhfejjjcdhabafdfbcajjdfjfhacadfbcbjhafbjgdgeiijjhffgdifjegfhehfbabhhjhadediggdaabibbhjajbajbddegjjagihbcefcfdegjbgaafdbahbjgabfhdbheebbajccfbaeiefbghjcefjjjhihiacagajfjaajicfggfbfggaheieddfdihbdfchecdchcggabchbcahdiahdghhfabcdejbfgjgdaajhfdcbecbcbjbbjgchbgcdjgbhiaicccbhaihfbebaijiichhfgfahhhefieheihfcbcegiffhfjfgfjcdejiijhhedjdccdeabjahjjceecihbbfhdcdgccbdiffiiiacajdheadbeihcgabeaeegecjgdjafbjeggbafahffdcjafdbjaifeagefjdcfabibbgdbfejgahjhibbdacbiggjaacihidjdfaidhjiihbcbfcbbcfijgcchbgedaiicjfiiicbbbbdihgdbgbjfgibbfaacjfhcehfadbabbgjcdijiidehjjgciiceedibbjidacbchcihiifcihccjgcfahaecgjdijgeegdgcgejjiadeijgdeedjdbdebjhddgciiegddfabfaccbdcjggiidjfehdgcegdahecejdifcedijajigcdhdjhbdiacehiegbdbhhihjejgahjecgdciabjgcegjiegaiibfjgafaacbjabcddghcegjeibhjfcgahdedeifbdbhiiiffjbhieaadfcebdficfabbifjbjjfdaahchighcfjgjhchcghejjdbghhecgdahjhajdcdfidgdehjjeehidhgabeidejbcjcjbdigeegabahfcjjfdgiaadidfjigdfhbgejfifhiddcahagdabdaaeagifgjahjhdgajfhfffijjfgdibdafjhiiihbddfbgjfjbddbiddcihhdhdadhfhbjbhghbadecfiegfichjcfdedbiaeigcbbgdcicajbfifjdbbdfhddfegiiebcccbjfgibjgjhahihcacchefhddbefgchcecjhaebgabfjiciichecggidcgjggdcagaibdichdfcchhbjgciijbbjihcdhgcijeibaggiifiiafhhebbabhhhccigbaafjjjjfaeeebcgjcfjhjibbeihfbfbbiajafbjiciibfbgfeaacfcfgbbddfagdehbbdabbbgdfiagicgehaffhhiggdeafjajfejcbjfjchigfhheegbhjhiccaghehjdeihgdaibdfdfeaefaeijhhabbjebcfgdaegejhhdedagehebfebjigiijfiibdcgadecgfgdceiihiebjijhhdhhjheaffhchggeiajiggefaheacibejefdjfajjccjfaeifabaicigfhdaeddehcdeahgjfbaffcchjajdgicdbeiebiedafjcifegadecigahhaggabddfgadcbgdhaechadgbjefajdbcjiddjhefjiiadjifcfiiifgjejjeihbhciajfjifadejcdbdgagagchcbfaehcdhcafdeegbgfdbdjedefchcbiiihfhejgiijffbajcfgdgdhegicbefcbahhfgfbdegjefjejbeefahdggibhaaigibbedghhgcchjfbgaicdaejjcijebifbbacehdfbbaaajciejgeecgjdhieiaadeeejieagecjijbgdejecbcccjbgjadabiiidgegjhchaiacegaiaeeagbfjdficbgifefjagideiibbcfgcacdcgichdiaejejahgiiaiidfhfadaagjcchiajchhcghjgfiddffegdcfgbegiaideebffiiabjjhbcecgcagbfcchcbcgcbaijfdbfghagdffbhbjajdcjidiihejdjbcaaaeaaibdagjfagaajhiijcbiibifgdfihfdahgggecagbgghhdjgjgjhbiihhjgiabdibfidfecgfjhghebbaadbihceeicdcahheddddebdjfigcgdejjdbbcijijfchdedibjegaeahccfhhcfcejhiihaacgdbigedebadbiafadchjcjciibaaafgacadfgbaebcbceedaihhebfjfefjigdfhhbdbhcahdbjigahafeidehgbdiihhdaahjddiafghagdfabcdafdjgdahhehjheibghjhbgjchcbgdahafaiebajgggbdgbfjdjghjfefhhabjfcbgbabedfidafiheigjjgjicajaeigcbafcbijhhbgddhggbechbjgcadffhcdffgfbgffdbbhdfgggcbgjdfbggefhbbaccahfdaibdcggaddacchghhbdafjaeafjeiefejddihhccfdaddbhibjaddehjbeibjidedajigjeeefhehejjjdffggfjdbhgcdigfejabjjicbcheiffigdgdeejjecbajdffecagfaejdhicaagbfjjddggagahjdfgcjiihibfbiahhcgggjagdeiifhccdbbafeefgfdadgbchcdicefhejidfaeaddfichdebbdeeajfdhffagchebhhhbbeaahefaeccfdfdagahhadfifgbdfjbffbfjehfieddbhhbfdccahiadgbfijhcbjjgcchhdecafafbefbfeahjefbhhgffjagijbdehiiefdfejfbfjeeiacgajabgafdfiejfffdhbgdhhfehbehfiiieehbegjhijaidcifcdceibibgjfjbdifdhdciadhagdgdgebhjabgbdafcaajeiiifaddaicaideiebhfehehebfcbaaebhcejfechaiiifbbbbdfdfcdfdbheiidffhbffdaeabbgjgfgcghchigeaehbfcdfgcfhgifjjejcdjehjhjjibbjbbageheihgeeheeejajhafhahcddiajgbbhhhgfafhccdjjajeccbijiaadhecfejgdjfacghfjjdciiagdgecfifbefifhdcjiigchcjggffghcbcefhafhadccahbahdgeeafjfdiehhdijjcjccjcbeeacciebffjeaedacfdihgcdiihegefbeeedeeddihahhjhafgiccjhjjfifecacajdegbjfeaiceihhabdhjfdhjjcgdaedccbefhddicdgbeigdhhcgegghijdeedfgaigeijejgegeahifabbbcihjfcechcjibfdghgbhhccfaafaifaebfcjbafgfjaiiedecefgighehjihhjdhhgbaijfjebcbejgfjfchdacadjcibaghgeiifdjdcecbbgbeaefbcjegecebdfjajfhjbejejdeadaccfgbfbhbchehfiiijdhdhbfcjjdjbjhijdfdecdeecjcbiacbjdcaahahaceejhaifgefcdhahcgbdbgecefgbdjfcdgbaehhdeddcbbjdhahadjfdefffehahageggbcbgjgfaajjeagbbeehhiie"

also the solution #2 is not right...
hmm... remove the strcount map, it returns the right answer indeed

but why?


91067 solution 2 vs 3

real    0m7.280s
user    0m6.602s
sys     0m0.632s

real    0m0.838s
user    0m0.680s
sys     0m0.285s

the performance diff is great!

Hint 2
Find the other prefixes whose masks differs from the current prefix mask by at most one bit.

this should be talking about the sliding...
so what else hint should come out of this?

*/

func wonderfulSubstrings(word string) int64 {
	cnt := make([]int64, 1024)
	cnt[0] = 1
	bitmap := 0
	res := int64(0)

	for _, ch := range word {
		bitmap ^= 1 << (ch - 'a')
		res += cnt[bitmap]

		for i := 0; i < 10; i++ {
			// not changing bitmap only ^ to flip that bit and see if after flipping
			// the same bitmap appears before or not
			res += cnt[bitmap^1<<i]
		}

		cnt[bitmap] += 1
	}

	return res
}

func testIsStrWonderful() {
	fmt.Println(isStrwWonderful("aba"))
	fmt.Println(isStrwWonderful("ccjjc"))

	fmt.Println(isStrwWonderful("abc"))
	fmt.Println(isStrwWonderful("ab"))

}

func testWnderfulSubstrings() {
	fmt.Println(wonderfulSubstrings("aba")) // 4
	fmt.Println(wonderfulSubstrings("aabb"))
	fmt.Println(wonderfulSubstrings("he"))

	// "faecdbcjjbjccjgcadeeaeddbibjiiecggbchebffdchbdicacdbcgdchdgeciceghcdhfbdhcaagghhhjichcgfdaijicahcihgdejhdafjghhjjeaeajdbiheegghhbfgdidcbbfgbdihedfgghjejafdbeedadcagcfejaaihihfdgdedibiiaaaajeifcfcjeeahhcgidcigjhafibejcadaegbeejbficdjgcjfjcgefeifcijibdhaaccdigjeafcebjaiciicghbheifebdadhdeceidcaijihhecjcfcdagdjhfabejjaijajdeaehfghehggjhaidibdidggfghejbggcadhfgjfaeehhicbgjeibbhbijhfghiejadefebigahfeheageagejehgbfajjadicggiegbcbeigdjbcijagajgfaihedcfcjgibiicbjegjgadahdeefajjbajdhcchdeiediecefacdabeegedhcfbiefjagacbfbgbdhejieffhjhcifijjfbccdfgdhhjfbdcejefedgihdjcahaggcjdfafciecibacijddbebccefchjbjfhgeghbhgaacehebafbdghdjdjdbiajeaiiccifcdcjhiaadchfagjdbhfhhiddhaehcjeafediehghbjaafiicgfejadcahfdgagfadbaiehaeaaiifcbcegihfbiacaadjjgahibbfijdbhbcfhhgccihdahchdeeahjeabeiegdgjfdgdfabfiajhagfdhhjaafacaihiagacgijfjdfafehicbeaicafciegfbefehddbfbgfhdcdchabgdigbceeiegabggeiafejbbdgjgijdfjiibadgfjehhaehjdhcdghadeejdfaicjdajjbiachaiadcghjfbefhgdfagbjibjfhegcbeiiecaaifiidgidbcbgidfgibhjcgjcifaegfgbifcjdgddadgedhbeidgbagiddfbaedhjefgjgfdbcabjcigahfbjcjhdhccjjhhffjcbhjhaeaiifajhjcfaejehjabeeggfaebjeiaacaabijjbeiggjfbdfcfgjcebbehbhcjjggaefgbdhfaigaagfafegbbahfiaabijcffbjddjefcfhhdbbgdjiiabdeaihfccbbfdaifcjhidfijdjdbaebaibfcijhhcddachacfbgddfejeafeddjabbdgcfibgijbcjjijhagddidhhaijiafbaajgejihadjdcdjagbfghjabdhhbbicgceifddgbjeceejbigjfgdedeejbjhdibjghigegjigghghhedhdadgedgbhebedjgjbjjhfedgbhhaebgaeebiajijdhdcacffbfedcbdebidbhieijdcbcdfjgdejehbgecgjcdcfhihgajjggffbaaheiabbbebjcbhdhaajbhicibgajddceaaacgaighhbaacbcbdbdbfacgiaaaaaiedfhcgfgfhhgeeidhgghfbbaadfegchahjcebacbhibbbjbhechefigefhcdcagjjhbhfggjgejfaceeebfibaegfchdabdjgfdjjcacgiifcdcigcfjbegddafefaiaihcbdgdccgiijejdadehgibbagjeajafaggehdajahfcbadjibdgjjcigccbdeiadijejhcjiihcbjijgabhcjfgaiajjigefiddhjcbjjhijbbgbjjeffedffdefbdbhifcidajichgcdchafcfhffifgcfedgafgdhbcedafabhgaejbhdejadiagibiijbdeeggejigadcedbcfbbbcijhjdhdhjceecjjbbfbegegahefigfeahgjbeddaedeaefeehhdhbcdiafhediejdbgaicahgeedfcffjagfgccbbfaajadaeijegfffjiidhgiiadiecbhifeicigjegijecgdjihfbjabigfhbjeejbfghfddehechdjffdbadieibjdddeaejhgijhhihejjcibbaaagbihcagifcjhebhgcjdejgabjjabgbfbbigbjccggbegfcfebfbcjaffhgcjfadaeggcajbgegdejhbfjiehbfgacdgaedeicigeejjgfijjjcgdjaeiadgjhgcdjfgdeagfcfdaicffifgbfdjbebeeadcdbdihjfijiiehbchhbdbbjgcgaejcefacjahcgaeajajejehhbdadihbcjfihaejjbabegbiehjhiccdacjaacdgcbfabgfeifchacbegffhccdgihbbbdcjihibjbjjdbgcfjcehedfjfahajbcfbfdhcgbjdcacbbafejafdciideecfacdhcjdbaigcieghigjfhacfeaifjifchdiggijcfaihhbaifcfjiafhffebbdcfhhacgadhcdcfcieifiifdjhhgjjiicbjjjijahcjijehbghfbbigiaieheiehbdfagijaghhjbfgbbdiajiadbddhbcidahiejehcgcjdegfidhhhcebgfhcifagjdiigicdedfjcifdbcfejifcihjgjeggbfifgbeeejcjfadadafghhhgaifegcajffcceedaejbdidhjajhabigfbaegjiejgaihhahecfccefcfgdhgaaabeiibibgdcgdeajbbjceififfdebdjcigehdgdjbgajhjhigjifejjccheceabcgbafbibiibgiicefefhiceigjhijjafdeiafffgcacjdccadcbecieegfgbfdejiafdhdhhfcjhbfjbehdfdfgcbjijiiddcafaaedcacedfebcafhcibcghhcieihdjdgbeaddgdbjheedjgihffbcbiaccbcgibjgcfagjgcbhhbbaafagefgbfcfdfajecfdhdfjifbffcacfbeccdbfgieegefjdhgdffcfjfajafhggfdjhhfhabgdgfcjhjgjjgicahghegdjddabdiechjdcbaichjijbbgehfgadibgajgdebfgahiidibahabgbcbhfddgefjgjbhbjgjdhhfbccheibbfcebjciafgbgijchchedeahagbcbegjehjgfejhcgihcbiheggeehjcigbjebddcbbedfjieijjjbdafbcdhajchgbbhebeijgheaghjhfidceiediagjhggjchefdjfhbadihaaggdchfhfbdifbibbjahbfadihefeaedgcedgdffdjecdhbhfjebihfgfbcfgieedhcijcfgaiifedebbefbgbgaihiafgafbchggdbceeiiadcbgfbfedihbjidicifdjbggebcdcjbgjfabadfgifiiehfffdfbfidghdhjieiifehiccabechjhggfbcaeccghdchccjhchgcjhdehjjecfbfbgjbabjjdjcgadfaffiicgabheadagadhfchibdbdbfeghdhhcedghdfahcgcihaaaeabfbghijgbbbfjaejhidbaaefgeiihdgdfaeijihfifjdcdjjbjfifdbhdhgeiedfcfcecjeecjcaccjjedefabaaghfbcgdagbdbjggjgjgijcijbbbeddiahahhcbjieabehcjgajcihhdjbgjibbiidigjgifhdididaifcfadcghgfceidhicfegaacideifeighghaggdacdhaiaaffafgchgaffhfjhaddgiaigjefbeabgfeaaibabffiiadgfgcddhfiijahcdccichdeebajhfacajgcihhjccajfdehbaehbdfebcaafiahibjbchcfggicbddbhhbbcfghdehjijhggiidagebfdhaifeghajabbejcgcbigifbedgaagdcccagbeccifgfbdhaidddbdehcafbefcifjjiiibb"
	//fmt.Println(_2_wonderfulSubstrings("faecdbcjjbjccjgcadeeaeddbibjiiecggbchebffdchbdicacdbcgdchdgeciceghcdhfbdhcaagghhhjichcgfdaijicahcihgdejhdafjghhjjeaeajdbiheegghhbfgdidcbbfgbdihedfgghjejafdbeedadcagcfejaaihihfdgdedibiiaaaajeifcfcjeeahhcgidcigjhafibejcadaegbeejbficdjgcjfjcgefeifcijibdhaaccdigjeafcebjaiciicghbheifebdadhdeceidcaijihhecjcfcdagdjhfabejjaijajdeaehfghehggjhaidibdidggfghejbggcadhfgjfaeehhicbgjeibbhbijhfghiejadefebigahfeheageagejehgbfajjadicggiegbcbeigdjbcijagajgfaihedcfcjgibiicbjegjgadahdeefajjbajdhcchdeiediecefacdabeegedhcfbiefjagacbfbgbdhejieffhjhcifijjfbccdfgdhhjfbdcejefedgihdjcahaggcjdfafciecibacijddbebccefchjbjfhgeghbhgaacehebafbdghdjdjdbiajeaiiccifcdcjhiaadchfagjdbhfhhiddhaehcjeafediehghbjaafiicgfejadcahfdgagfadbaiehaeaaiifcbcegihfbiacaadjjgahibbfijdbhbcfhhgccihdahchdeeahjeabeiegdgjfdgdfabfiajhagfdhhjaafacaihiagacgijfjdfafehicbeaicafciegfbefehddbfbgfhdcdchabgdigbceeiegabggeiafejbbdgjgijdfjiibadgfjehhaehjdhcdghadeejdfaicjdajjbiachaiadcghjfbefhgdfagbjibjfhegcbeiiecaaifiidgidbcbgidfgibhjcgjcifaegfgbifcjdgddadgedhbeidgbagiddfbaedhjefgjgfdbcabjcigahfbjcjhdhccjjhhffjcbhjhaeaiifajhjcfaejehjabeeggfaebjeiaacaabijjbeiggjfbdfcfgjcebbehbhcjjggaefgbdhfaigaagfafegbbahfiaabijcffbjddjefcfhhdbbgdjiiabdeaihfccbbfdaifcjhidfijdjdbaebaibfcijhhcddachacfbgddfejeafeddjabbdgcfibgijbcjjijhagddidhhaijiafbaajgejihadjdcdjagbfghjabdhhbbicgceifddgbjeceejbigjfgdedeejbjhdibjghigegjigghghhedhdadgedgbhebedjgjbjjhfedgbhhaebgaeebiajijdhdcacffbfedcbdebidbhieijdcbcdfjgdejehbgecgjcdcfhihgajjggffbaaheiabbbebjcbhdhaajbhicibgajddceaaacgaighhbaacbcbdbdbfacgiaaaaaiedfhcgfgfhhgeeidhgghfbbaadfegchahjcebacbhibbbjbhechefigefhcdcagjjhbhfggjgejfaceeebfibaegfchdabdjgfdjjcacgiifcdcigcfjbegddafefaiaihcbdgdccgiijejdadehgibbagjeajafaggehdajahfcbadjibdgjjcigccbdeiadijejhcjiihcbjijgabhcjfgaiajjigefiddhjcbjjhijbbgbjjeffedffdefbdbhifcidajichgcdchafcfhffifgcfedgafgdhbcedafabhgaejbhdejadiagibiijbdeeggejigadcedbcfbbbcijhjdhdhjceecjjbbfbegegahefigfeahgjbeddaedeaefeehhdhbcdiafhediejdbgaicahgeedfcffjagfgccbbfaajadaeijegfffjiidhgiiadiecbhifeicigjegijecgdjihfbjabigfhbjeejbfghfddehechdjffdbadieibjdddeaejhgijhhihejjcibbaaagbihcagifcjhebhgcjdejgabjjabgbfbbigbjccggbegfcfebfbcjaffhgcjfadaeggcajbgegdejhbfjiehbfgacdgaedeicigeejjgfijjjcgdjaeiadgjhgcdjfgdeagfcfdaicffifgbfdjbebeeadcdbdihjfijiiehbchhbdbbjgcgaejcefacjahcgaeajajejehhbdadihbcjfihaejjbabegbiehjhiccdacjaacdgcbfabgfeifchacbegffhccdgihbbbdcjihibjbjjdbgcfjcehedfjfahajbcfbfdhcgbjdcacbbafejafdciideecfacdhcjdbaigcieghigjfhacfeaifjifchdiggijcfaihhbaifcfjiafhffebbdcfhhacgadhcdcfcieifiifdjhhgjjiicbjjjijahcjijehbghfbbigiaieheiehbdfagijaghhjbfgbbdiajiadbddhbcidahiejehcgcjdegfidhhhcebgfhcifagjdiigicdedfjcifdbcfejifcihjgjeggbfifgbeeejcjfadadafghhhgaifegcajffcceedaejbdidhjajhabigfbaegjiejgaihhahecfccefcfgdhgaaabeiibibgdcgdeajbbjceififfdebdjcigehdgdjbgajhjhigjifejjccheceabcgbafbibiibgiicefefhiceigjhijjafdeiafffgcacjdccadcbecieegfgbfdejiafdhdhhfcjhbfjbehdfdfgcbjijiiddcafaaedcacedfebcafhcibcghhcieihdjdgbeaddgdbjheedjgihffbcbiaccbcgibjgcfagjgcbhhbbaafagefgbfcfdfajecfdhdfjifbffcacfbeccdbfgieegefjdhgdffcfjfajafhggfdjhhfhabgdgfcjhjgjjgicahghegdjddabdiechjdcbaichjijbbgehfgadibgajgdebfgahiidibahabgbcbhfddgefjgjbhbjgjdhhfbccheibbfcebjciafgbgijchchedeahagbcbegjehjgfejhcgihcbiheggeehjcigbjebddcbbedfjieijjjbdafbcdhajchgbbhebeijgheaghjhfidceiediagjhggjchefdjfhbadihaaggdchfhfbdifbibbjahbfadihefeaedgcedgdffdjecdhbhfjebihfgfbcfgieedhcijcfgaiifedebbefbgbgaihiafgafbchggdbceeiiadcbgfbfedihbjidicifdjbggebcdcjbgjfabadfgifiiehfffdfbfidghdhjieiifehiccabechjhggfbcaeccghdchccjhchgcjhdehjjecfbfbgjbabjjdjcgadfaffiicgabheadagadhfchibdbdbfeghdhhcedghdfahcgcihaaaeabfbghijgbbbfjaejhidbaaefgeiihdgdfaeijihfifjdcdjjbjfifdbhdhgeiedfcfcecjeecjcaccjjedefabaaghfbcgdagbdbjggjgjgijcijbbbeddiahahhcbjieabehcjgajcihhdjbgjibbiidigjgifhdididaifcfadcghgfceidhicfegaacideifeighghaggdacdhaiaaffafgchgaffhfjhaddgiaigjefbeabgfeaaibabffiiadgfgcddhfiijahcdccichdeebajhfacajgcihhjccajfdehbaehbdfebcaafiahibjbchcfggicbddbhhbbcfghdehjijhggiidagebfdhaifeghajabbejcgcbigifbedgaagdcccagbeccifgfbdhaidddbdehcafbefcifjjiiibb"))
	//fmt.Println(_2_wonderfulSubstrings("faecdbcj"))
	fmt.Println(_3_wonderfulSubstrings("faecdbcj"))

}