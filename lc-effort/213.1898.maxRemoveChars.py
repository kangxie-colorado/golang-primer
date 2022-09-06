"""
https://leetcode.com/problems/maximum-number-of-removable-characters/


"""


from curses.ascii import SO
from typing import List


class Solution:
    def maximumRemovals(self, s: str, p: str, removable: List[int]) -> int:
        def isOkay(m):
            s2 = ""
            for i, c in enumerate(s):
                if i not in removable[:m]:
                    s2 += c
            return ''.join(sorted(p)) in ''.join(sorted(s2))

        l, r = 0, len(removable)
        while l < r:
            m = r-(r-l)//2
            if isOkay(m):
                l = m
            else:
                r = m-1

        return l


"""
not getting too far
"qobftgcueho"
"obue"
[5,3,0,6,4,9,10,7,2,8]

okay.. the test for is b is subsequence of a is too naive and wrong..
"""


class Solution:
    def maximumRemovals(self, s: str, p: str, removable: List[int]) -> int:
        def isOkay(m):
            s2 = ""
            for i, c in enumerate(s):
                if i not in removable[:m]:
                    s2 += c

            if len(p) == 0:
                return True
            if len(s2) < len(p):
                return False
            i = j = 0
            while j < len(p) and i < len(s2):
                if s2[i] != p[j]:
                    i += 1
                else:
                    j += 1

            return j == len(p)

        l, r = 0, len(removable)
        while l < r:
            m = r-(r-l)//2
            if isOkay(m):
                l = m
            else:
                r = m-1

        return l


"""
TLE..
27 / 67 test cases passed.

>>> len(s)
6063
>>> len(p)
2451
>>> len(r)
4536


size wise, it seems this should handle it..
is it dead loop again?

debug shows 
m visited these values
2268
1134
1701
1417
1275
1346
1381
1399
1390
1394
1396
1397

so it is still log(K)
where it takes time? the sub-sequence? I think my solution is already O(m+n)..
what is the better solution?

I see the draining piece of code is here
            for i, c in enumerate(s):
                if i not in removable[:m]:
                    s2 += c


"""


class Solution:
    def maximumRemovals(self, s: str, p: str, removable: List[int]) -> int:
        def isOkay(m):
            s2 = [c for c in s]
            for i in removable[:m]:
                s2[i] = 'A'

            if len(p) == 0:
                return True
            if len(s2)-m < len(p):
                return False
            i = j = 0
            while j < len(p) and i < len(s2):
                if s2[i] == p[j]:
                    j += 1
                i += 1

            return j == len(p)

        l, r = 0, len(removable)
        while l < r:
            m = r-(r-l)//2
            if isOkay(m):
                l = m
            else:
                r = m-1

        return l


"""
63 / 67 test cases passed.

damn.. the most dreaded error is here..
oka... logic error here
            while j < len(p) and i < len(s2):
                if s2[i] != p[j]:
                    i += 1
                else:
                    i+=1 <== missed this.. when s2[i]==p[j], both shoudl advance
                    j += 1
        can be simplified to
            while j < len(p) and i < len(s2):
                if s2[i] == p[j]:
                    j += 1
                i+=1

Runtime: 6016 ms, faster than 47.49% of Python3 online submissions for Maximum Number of Removable Characters.
Memory Usage: 26.5 MB, less than 90.78% of Python3 online submissions for Maximum Number of Removable Characters.

"""


if __name__ == '__main__':
    s = Solution()
    S = "vxcwjdelswupsguxcplzkycblrklmwuwctuflwupnaixdjhyslbqnmxdaiwsepbdhmubnlcwevqxplwclbgwohuiaphbkzhtagbjpkwsprgwbiwgvctvbzvgqzabcewykrnpkbnhaeqcqxnwjwhynwrbjukxwxcgezgbglhgwagwmjscfxqdchwnyxecdaxccribcisvtqxojbfxbzasacjirjptkatkyyzshbhfobligthxepvgjlboywwdyqayxaicijjiyibfesdvwfvoevtzhvltwjaqhmryrguqeytlqpqxxpstdipqzcmjnzvafgcrhmcwtkwhhongiijdmxdhqlsvhmcyzvcdvmlclwjdavmcgtpmbmkprjgdhcandqdvapvvnpkrenharnqeaeeylralfhyewbfcankotvhytuexbxbjwihchmghsrwvfpznojazpgbdhkhzygwttlbiqqmjipcpgugswnmcxnnpeiywduohwdjphahkazcpyuzicaazikothynwyfcenevmbyyizmjrqbckypgaewyzgsewrzewkmhjsnxrbnaocnykptdubdsxptmnpehkcukqsompafmdnvfofpzqydjsfyxetsdssrkzwtlildywzffkzkvhdxzsssiveblmlhrzavnanwreenqzegfeuglguamngbkzxgwmhcobsbpciaikymtxtwzrweoeljbbmjyruwqhrsbpkeybtriqhwurakauipzkxwzdgzjnqmdzmrpwuqfpswpqvoebaihozbtczrnylsbgykyildktocgvrxwgwlbzelbcsyvuvsfegfmcamunqniklkqfcceraqfrdkdcvvllqjguuqffsdzvehfkjyhlqkhhecghujlzfqqylvdfdpmtrgrydvvnfdoqhncwetfziqlfyeybgvnglroydcuqfsazqamzwvymhfkyttumuucukjxaylknmnrzybknehtlmeuzbpemsuwnbgkqeliwliaimatxliafccxxhuaxdfzlvpxbglcjjgufdnhjnyfjktgrzlyoivqccnsdmcxkcmeqjptjxwvemvrpyikmacumqvhbrumldkheptkicxgjphuilzyawrqguwtcywzcimuzpepjuwnwflarpbvyeucynswzckdyssnxkmklsosddbtyxkohfdffvyrjidgcqdrgpuszkvkyvlndjghmxfcyaiznzjryzhijrfcpfmzptulkvtaoilyweywxtxctavmrrjocpkktpljtwumisttjvvihhyvpxchhnkcvcpewdkubmcukzhycwruixuuaagbcaqeeribenzenwlyeyzydqnyzxplwtnrbsuhbrlzgfdotjywdzwiccefdedwexnyphclyfyqywolfhtnicqmixzkesqnfibktuzamejhynejvoeigpqfiiwtcwoznicuwuuksuppfmtgxtnpbpgtabxuxdrseotufwkdbdzpfajpzwvgrbhgytvvmjazbpsznnglwegchmytypkhbozfneybqxclrxhryacrkvubywjgwezwtcqycbontaiaogytrmlozymlmcowyclasibngjfrcrplitxivqokcpukayywxubksgvzhgvhohwxxwatqntalyllnrhnacppyriwmzyerlmneaydvcsanrzpwkdoecbtpiewatytcqpseehyvrxbizctwzfgjoopxikabpxsqaoszzhvuhixpbzhvfuxjdlvrwiinpjpnhgghdcyeengevrucurepfyqnlcucksnefrzjcattljzfgfxmxnesdywuzsbiqlyzcoxcrqllyrqapolrzpujtrwsfxfhfytzjhtbggwvqssxqjiefncdfkccllryftikedmfxpdrduhdrzfdlbjlvcnmkioqzhjkpjrhfycrrjdhbjgflxvzmlrttejwywabrlykshabppceotluaptmrysrtfzapaizkvaagscwpeehcckzxoqnjtetibhmmrbbpdrcpybpgfjvnyendpsqlhmaedyqwuukbekrvvuzncibkcjspplucprnbduxrrcwpgrtxpoqafyybwujtiquvirzbvlrczyjxhdwnyjjdtvttpawoecuzzyicnirrivqeqfxuqjcjbgxkgsubzyaetxhajesabdzeuzumimaguujnhqysrmufouriftjncssetaixmpjtlrxfkklfnvncwkquqzuinpnrvxqeueqakpnzkarjuyxolafwyvxacvatbtxyeuljqtxcnbdingocbmombxbrsjngpswislgkugbhjbumptqamyyziijqnbhqirwmoeqnyhvdpzwwztacrevyxfmfzaweiltwhijpjovsyeodaeyfwbzxyprdoyatyxisenqvervdmqqmoafnpghqahbzqiohxnydeahzherksopctmweunihjmqtdteajkcxrtfhkqeyzcdqxpfxrijwdvtyzckxaayhhbytbnoybgszrnwtthyiaweqhfpvsjmvidmyjwpqfamaoqtqjbgcnxywdapdmdkxapvzhycnbcntlmpvipvoodacishpjodwqiqzmatuyfxlxouzumjqeuqfsdmdiaokcqnxtogihrszjvnnhygwgzfcyaniasauggmdhysyxokyfzzfrjcjokwnmswtgqokeyiyxoxluwmajdjzuglwbxnyeakejhyubveqdvdfwjlvxezvofqhkxopligprdnvlphauqyodkbbakiqklgytadiucvsrcfdjxgvbjedoammwywwxuytgkzcwofyqucmmrtdmjibryrncofjdzsfcwjfwdkpxubkzaxjjgfjojnvzrasaxpwhosgnmwpqusvfmlrivpeblmcozxpyrqreixtrokpfzlvoepycbtohmbpolueinvbilliptbguufjzrrlfioimyuadzhncxnphqmovtbyncgwvthbmmsatmlnsyjojjkwezsowahwlzetmhcqlqcuvtavtxbnmkvvjgfzinzgafovcdymgzlkzfxagmnriulaqkbcajrnqnubmtjmirzwvgtzsbgnmnihwwmkidgunvatepxyhcmquxtpwpotrgunapkmjiozutqcjyqoycapdizcritxvnqvunhhvzwgfnykgnzlwvsconsarkbqcmnzcheamtagezdgtfjxpavqihhejgcofsxstnlkpsgrntdypmueomrysodvxxxlofwqpaczudfggfmdrkuuncxnypnuqyfvtfzwgitlixqhaofdgiddhwahmayatjwswdjpkksglilgoiuklhzwqqkmtbzgwmjselayjenivkoswcviwxrwxizjjdxexibzjpbwtiyvhamoisvejjodjrdwmzwlnufdaczglicgsavmlvnrscridqjkhatlvzimolovsjczerimiwtlsquevpoijwmifutzeaogzlbirpdexsoptvmeawaofrpyxuagyyynzhwiujyhcfocyycuxvdwngytooybleormacscykxlamtqqpilsselsdudufuvmwtkfyboqcffwwoqawrgtwarslnoispjpsecjfziloqlrewxtojjyybhgfzjkakfncesqtqwpylcnnngupvuhkndwiewkqeidzcatjholhogmncksctgzjksqbrvhadneejduhayqcezapbonajxmxvsqdymuewprooxfovkugkxpxtaetqbsfyueqmrmbcznuvdotfznkfyztctuavkvvmfynnyoddypuyutopvzzpoykofcraqahqmojbxxhugbnbdfmmrsnjjsanrobfhsxwwrpnbabfzdtkidqrmnsipafjguujjjwbmkywfunnzilxjhylxalrwygqxzupqsqcfiqyrekplhdalvayynmzwtwgookymgghgnnevzcuhfqmvjrcwjnfkksnprygnwuvrdtwkyuapvblsbxgmvrzrrbfuudgkqzkwlkfeocpapcnzxulqehpbbouxymqusqkclnfcsxorzwopbyvbnhncpmxgcihloihbnbxxzyuqprfeslwfgtfhcycwtalippasbdlsferfzailmlvezvgtcakqeubvifosadnkqkjpciyartzuvzyglskepwjhnhzlhuzrfhwdqcdhpkxnsossrqqjtsuvpiwgrdvqvgqfdymhvrlexlhczutukxywzhpczuuoghjztppchhcrekdtbnfhmrwxtywudlwmqwwopinxbwkgrrnedideqrfselyikqhxlgnadyscnyvaqkfoosfwtheudezchtpkdwqawpjlsxnjnlfswogapxmttoovmayasybkhtglxtzsjeouxoyxcpwkeulgcfpayzeopwcfuxfquvjwsztmvztalrrjwsrmrmxcsequfxrsfjxzbfuftzkkgasxzikvklsiyqyfershnylutkuswnyryhpqxhqwnhcxobgdhmorpucucnveujxeyesplbjhlmjzgnvbjfjzqzpqcqjocpcaurwwzxyhwnpzghhmqpzpuclkuriwoavbtjrwoyoxpdbsyhjeaqqztbcumxuzwgametnlawkoesrylybzilswqljpuhkzrvrcpqlhllnclmyippejsnhavpxboviqbucurfciangnaicahedfdswxhijdiigagmikxbqzqddwyntlevyxjshzoukyxbhmcdsqybqwzqguviqgzyoshwgieutqbwrsjssjgmsjnkkezkpfqcmdjoelidhfumospvawfgebzvgmwchruikybddlymiuuinqkbmcmabobfsutytnycghvgbgfptbnjvtdlkkpdqctmdflekoemkmqpqyryizgjvnpgdgehqrjcsiljktidhcsnwfsfbyzzzzxbkvyaqxrrhpfwykqxvwvhgtfdvhxbwzewzasilpodkfjggdrmskegdfsfcxjdziaubrqgtnutffervzutokvqhozkpdlnfstlybeqnnpqdtxyfmzzlfjjaluhyduexycxmjrxlyvwwphvlozpynrdszubkhhtrgiwyvaeiystrzrtvwhaeoztnpidwicnfxnftcbqesmtqxvbysnozbhsqhrlcauavsqhehulozaarlagqmtohrruheryhxkocfkcqqoanfrhemkohzjdcwwihkxucsfkfabdshkzbmvrumowvpdpijbtpwkxyhaowfliwuydmzdicbrxygjhujtqypmrwndqcxlgewpfguuorcxveetpwqikpxrjwteeqlpjpfwrtovradfxotyyulefctpqwlxinpshcyhvkmaplahicoooesqbcenzcgbpygtuuielzpjxrgesbrbozqfhukxkhjmbfvkwgjgswhztfiwxsupxegwebppqwjnfdcfdqfmgxjihjrrqpzlselfokzonropbcwuupwnktgkedattpafuamlvtzmiuoafxbosfzasutroznjhumcuqhslzhuquiijlinpiikeuwifskkegsawltgzmwukxoidhtbvvkprbvqsmqihflxmxipablikuvfarojkldwsmqhaxovxuwwggttrqnnyfoaxeuzwddzuntmzudxjvrjmtooonyfwfsczxaaksybvosatucfkjqltkrduxsomlsqbcpivslyi"
    P = "cwdlwuguxclblwcflwpixjhlbqxipmceqxplwcubztbjpwsgivaberkbeqcnwwnwbukxcbggwmjsfxcaxccvtjbbsirfoblihywdqayxjjysevtltwjmqetpxpqnfgcrctohqshcyvllwjdctmkprjhndqvvnkaneehywbcoybbjwichmfznoagbdhwiqcpumxiduodhhkapzcaithfcenevjqbckpwgsezhxrnnykdspnecksofmdfzxdkdwkdsilmlnwnqegeglamkhobsmtxwzwebbrqserwkapkxdgjnmzuqwpobryiogwgbelvvsfegankqdkcvvqusdvfjykhchujqdtgydnoqhnwqlyebgryaqawvymfkttmuaylmzkhbpeuwgliimtlafxuflpbgljgfdhnyfjyqccnsccmwempkacuqvhrdhekixgphuzgtwzimpuflbvwzdynxkmlddbyxodvjicuszvnjmfjyzhifcpfmtukaolxarktljtihychvcpewkumuhcxgbcqebezdnyxpsuhbrloywdcceeexphcyqliekamjniwtoisumtgpbpguxdrskdbzpwvbhgvvjpnngletykhfncruyzcqbntiaotmlzowycsbjrcrltivqcwxsgzhgvhtalrnyrmzyerlydcsarwcwaycqsyxbzctjpboshhxpbvfuxdlvnphdynfqlucksnzjajffxmxsdwzqcocqrplrjrshftztggqxifdkclymfxudjvnioqhjkjhfjhbzlreaypcomyraiascwpecckzjebbpcpbjypslmedyqekvjslucbrgtpqabujtirbcjdtpweziirrveqfxujgxkbzaetxhajezuumjuuftcsamtlkfvnkvxueajuyowacxyelqtiombjgpslbbtamyijioqvptacrvyfmzawelhpeybzroyatyinqvvdqqmgihxdzerscuxhkqydqxpxidzxaayhhyysznwqpsmidjqftgydkvzhcntmvoacihmuyfxumeuqmdcnohrvnhygcniagmdyyxfzfrjcjnmsqylumadjzlbxyeakhudfjlehxlprlpaudkbakklytuvfdgbjomywxgkzcfyqucmdmirjdcwkpkjfjoswgmwmivemzxyroflvoepbtomuenviiptgjzrloimyucxmtbycvbsaokzowawzmqcvavkzgfovcygkzagnruaqkbrqutmirvgsbnmihwmkidvehcmwnpqcocpixqvzwfnyglwconsrkbhmfxhegsstsndmmsovwaugnnuqilhfdhatwkloiuhwbzwseyjnikvxrzxwimejrmnagvmncridhvzzrmsqjwtzogbrpdxsoptvawaryzhccyxvwytoylracscykmqseludmtfofawtwarospsjltjjyktqycnundiewkiclhccjsbradeehqceapoaxxvsqwrxfxptaeqfrudtkyzctuavvvfnopytpzzkcqoxbbfrsswzkrnsjgjjmkzxylxawfrephlvnzwgoezcfqvjrcfpyuvdtkpvlxkqzkeapnzqpbuqlnsrzwvnnxhihxyprewgtywtafamvkuakpcazsewhnzlhcdhxosjspwgdvgfdyvelcutygtchretnfxtywdwmwwpbwreddeqrfseyikqhdskozanlgmoaybtsjxyxcecaxvttlwmmxurbukgaxikvkyqerylkuswyrpqwchruuuyesnvjfcqawwwpzghppkuroavjrwysyhjezuugatnesylyziqjuzvqlclmeavbuufngnicahefhigkxntlvhzyxbdzvizoieuqbmskzpcoelhfsvawfbmchrbdyiuinkmmbynghgptvkpmdfekeprizgdlknfsfzbyqxrfvvgtdhzioksgdfjdiargtutffrvzuvhnlypqdtzljuyxmjlyhvznsbhhgwyyszrwopidwifxftmzqravsqhahreyhokqonrmohzdhkxffhkbvawiuycyrncegurxeeqwetovaxyeftpqwnshyvkmapobncgbytuuieprofxkbfkgfipwepqfihrelfonobwnkgketutzmafxsuojumhijlinpeuwegalgmkvkpbqslivroklwsqhovuwwgnfuznmxvjtnfzxaassfkxqbpiy"
    R = [3747, 3188, 3480, 1083, 3009, 3593, 1647, 3698, 4134, 1588, 4087, 3250, 899, 4191, 790, 3603, 3596, 1630, 643, 4046, 3812, 2534, 2019, 4144, 422, 4303, 5411, 3442, 1390, 3030, 1710, 1536, 1829, 52, 3589, 1216, 3236, 5049, 862, 5281, 2001, 3634, 3914, 1391, 5690, 5152, 2614, 4013, 3085, 662, 1932, 1869, 3298, 187, 2116, 1669, 3195, 1288, 3200, 3093, 4714, 3763, 2686, 1892, 2969, 3322, 813, 3976, 2501, 4514, 1897, 490, 926, 3869, 4559, 1446, 2656, 3934, 3627, 3233, 1816, 4463, 3326, 4693, 5649, 359, 178, 2034, 4810, 2532, 5413, 4778, 4257, 407, 1534, 2570, 2999, 5286, 4806, 4268, 5472, 4682, 4264, 5439, 972, 2053, 1566, 4351, 5616, 3426, 4031, 3168, 1322, 5376, 2077, 3658, 2665, 2132, 517, 1256, 3883, 421, 5251, 3090, 3311, 1343, 3467, 5330, 519, 5184, 1691, 3266, 4238, 2855, 2480, 347, 2048, 2339, 1449, 5148, 4280, 2838, 263, 4394, 3987, 3881, 3462, 4725, 4866, 4980, 1393, 318, 5154, 2576, 0, 227, 5126, 5156, 5605, 2391, 1422, 3608, 2520, 2667, 2491, 4946, 3289, 2867, 954, 3867, 1414, 3169, 802, 2761, 4809, 4915, 5600, 1651, 402, 3378, 966, 5669, 2941, 5037, 1682, 4192, 1844, 822, 768, 635, 2791, 2473, 5496, 560, 2470, 1262, 5215, 1250, 5185, 4415, 1905, 5282, 1938, 4472, 3675, 2390, 3047, 1774, 876, 4609, 1202, 221, 2294, 4345, 4938, 3244, 2093, 4307, 337, 2609, 3782, 5588, 4988, 4445, 858, 2402, 218, 4015, 3759, 2463, 3422, 4868, 1150, 1683, 1447, 3464, 297, 4887, 1571, 4416, 2106, 3248, 4384, 3654, 3229, 884, 993, 382, 2467, 4596, 3701, 4441, 2268, 3857, 1659, 2063, 3393, 242, 584, 4926, 4602, 1902, 3348, 2503, 901, 995, 5382, 1329, 3811, 2263, 3628, 353, 5074, 3489, 2202, 1028, 314, 844, 489, 5044, 2247, 1702, 1000, 631, 5477, 3257, 2387, 2649, 5050, 4701, 925, 2087, 4622, 2167, 3411, 908, 475, 420, 5265, 2552, 3282, 4339, 1608, 5661, 66, 4876, 992, 1981, 1466, 963, 2670, 4558, 3824, 3394, 230, 5505, 4167, 1866, 330, 1864, 819, 1748, 3575, 4428, 1433, 2289, 3904, 505, 1848, 2355, 3452, 1578, 971, 5711, 2354, 973, 195, 1320, 1925, 1397, 11, 1180, 3834, 835, 2586, 183, 3104, 4822, 5567, 4217, 1626, 825, 5539, 1444, 4350, 3969, 2125, 4683, 313, 1128, 4346, 4127, 2335, 4815, 117, 4634, 4002, 2860, 5038, 4886, 2287, 5089, 1168, 272, 4993, 5104, 3302, 2210, 5163, 2487, 4828, 1943, 4454, 4620, 2959, 335, 3533, 288, 1641, 688, 5302, 4301, 1846, 4053, 5550, 3231, 477, 5657, 1808, 5071, 940, 4058, 4054, 2878, 2798, 4881, 400, 2713, 1781, 3615, 2638, 492, 4469, 1304, 1636, 2897, 1036, 503, 671, 3288, 1629, 5575, 3821, 3527, 4644, 1432, 955, 4924, 4131, 888, 5165, 595, 5702, 1788, 1071, 3466, 3405, 5609, 4202, 552, 2859, 1923, 2573, 1119, 2259, 1007, 3941, 3179, 3215, 4182, 3569, 1831, 4999, 4909, 2016, 4640, 2455, 2726, 207, 4153, 1503, 604, 1965, 4624, 1034, 656, 2062, 3153, 5602, 3965, 5610, 4174, 4612, 1291, 1408, 5395, 2952, 1463, 4300, 2551, 169, 1794, 3931, 3935, 3521, 258, 2025, 3605, 1885, 865, 1458, 4775, 4488, 2632, 179, 4553, 1717, 3808, 5283, 273, 1952, 3286, 667, 3157, 1556, 4085, 5434, 798, 3637, 2598, 1347, 3561, 3870, 2305, 3806, 453, 3204, 5107, 4841, 378, 3310, 411, 5650, 4168, 3706, 3110, 947, 1526, 5664, 5058, 1768, 4870, 425, 5209, 2404, 5618, 34, 1510, 5198, 627, 4214, 3145, 4608, 2896, 1324, 5318, 5659, 4462, 1605, 3856, 1204, 934, 2742, 460, 5704, 2080, 3855, 4036, 2346, 726, 2561, 3859, 5469, 910, 5195, 1770, 2985, 1915, 5222, 1195, 979, 4681, 4983, 3717, 4927, 1575, 5231, 4996, 1350, 739, 3722, 3217, 1113, 1145, 5414, 4147, 364, 784, 5170, 4717, 870, 181, 2086, 837, 4419, 3690, 1186, 3905, 3646, 290, 3874, 244, 2366, 5046, 946, 3786, 5250, 4126, 3314, 2907, 5123, 5478, 5560, 4032, 4982, 2265, 1747, 3100, 4139, 812, 2075, 687, 3191, 1631, 5157, 1130, 2627, 1112, 948, 4297, 846, 887, 869, 3843, 4329, 3293, 2737, 2121, 1173, 2800, 2504, 608, 5296, 2146, 5339, 2795, 5147, 1014, 4274, 583, 2089, 628, 4832, 3678, 1574, 2581, 1559, 4023, 5484, 1505, 3356, 3353, 580, 2925, 5476, 3409, 564, 4934, 5509, 785, 871, 3849, 1019, 5306, 3471, 4545, 3071, 2849, 2725, 744, 2511, 903, 1944, 2816, 1106, 2073, 5361, 3751, 1026, 63, 5435, 26, 3775, 1581, 1424, 1253, 3267, 4969, 1687, 2100, 5146, 2067, 5059, 4396, 2129, 1516, 3107, 2873, 4003, 4805, 5205, 5039, 58, 340, 769, 724, 1151, 3841, 847, 852, 719, 4706, 5462, 678, 1327, 3445, 1416, 5652, 4844, 5608, 3866, 5675, 3438, 3011, 2619, 4356, 2746, 1476, 5145, 5009, 5677, 5419, 2887, 203, 633, 3702, 5642, 236, 2139, 2333, 589, 4519, 5543, 1063, 780, 1239, 3649, 62, 1507, 2788, 3007, 725, 5177, 2436, 393, 1411, 3395, 2898, 412, 701, 536, 4626, 1496, 1545, 56, 2589, 136, 4439, 2152, 3317, 240, 2854, 2664, 5349, 4597, 410, 3228, 1261, 5345, 5284, 1610, 565, 4919, 4180, 2636, 4586, 161, 4756, 403, 1189, 198, 4989, 2752, 2582, 2756, 2412, 3686, 1753, 2336, 2329, 1418, 5117, 2637, 2796, 712, 3287, 4018, 4184, 5700, 575, 4700, 1723, 2344, 4532, 3556, 3512, 5448, 1889, 2506, 1661, 4424, 1974, 3774, 3986, 4665, 3036, 1201, 1744, 5710, 5721, 229, 2822, 1366, 3664, 188, 2995, 1136, 547, 2290, 3397, 197, 3795, 2429, 2698, 2544, 2755, 1445, 2175, 372, 1599, 4686, 3891, 620, 1325, 4657, 3219, 5436, 4882, 4458, 974, 4009, 3482, 2316, 212, 4742, 2018, 3885, 222, 2044, 1617, 1697, 3028, 924, 3895, 2024, 499, 2194, 3126, 4719, 1890, 3943, 5102, 3488, 5073, 2923, 2138, 1229, 2764, 5101, 2555, 3720, 4047, 2771, 4112, 2712, 4489, 108, 210, 5238, 2275, 4548, 1139, 4260, 5671, 3275, 5098, 3724, 4330, 3316, 4953, 87, 5632, 4728, 3602, 4410, 3693, 3626, 1894, 5087, 707, 38, 1928, 3519, 3225, 4082, 2850, 2741, 4633, 3253, 4963, 1172, 2766, 141, 3017, 2642, 1072, 3964, 1919, 5190, 5494, 302, 2842, 1515, 2988, 3835, 1633, 4055, 4119, 2364, 1487, 1942, 1519, 4090, 162, 5651, 1960, 5631, 2408, 4571, 4004, 5363, 2401, 315, 157, 5362, 1920, 5485, 2137, 3207, 5293, 2966, 5287, 1511, 4947, 5021, 1182, 2434, 2451, 1572, 3691, 4211, 104, 5001, 531, 5113, 3687, 1680, 1813, 804, 4064, 3357, 2360, 4531, 1795, 1038, 4638, 1300, 2606, 4340, 4075, 3522, 5529, 3141, 2388, 1159, 2279, 4678, 279, 2597, 3243, 2900, 4715, 5324, 2708, 2624, 4751, 1550, 1377, 4952, 3669, 357, 470, 1402, 1555, 985, 4325, 445, 1598, 5161, 2599, 515, 4504, 53, 4888, 3469, 1269, 1024, 4250, 5179, 764, 554, 4727, 4022, 981, 3388, 3997, 310, 3853, 3863, 1825, 1984, 1332, 2396, 3500, 3049, 3053, 2983, 5656, 506, 83, 190, 1348, 20, 3441, 4444, 594, 2909, 4122, 500, 219, 3922, 5701, 2944, 4077, 1469, 1705, 5325, 2317, 1394,
         3631, 5288, 2865, 4012, 1701, 4769, 3872, 2435, 1708, 1714, 5581, 1436, 4898, 5601, 1413, 4764, 969, 3283, 695, 2924, 738, 5354, 863, 4762, 2538, 4955, 1456, 2040, 306, 381, 3878, 1719, 251, 4507, 5433, 1986, 1727, 429, 2096, 4570, 3588, 525, 1375, 4365, 2836, 5294, 5431, 2107, 4698, 5663, 3920, 3490, 4278, 193, 4550, 3661, 3148, 2471, 2674, 5088, 4331, 2912, 1769, 1741, 4381, 4109, 6, 4433, 4354, 3005, 1237, 4768, 2128, 3406, 3988, 4537, 619, 2027, 1426, 3532, 144, 5592, 4349, 3692, 5261, 4434, 5429, 3494, 3336, 3583, 2187, 1973, 224, 287, 771, 548, 4205, 4363, 2647, 3671, 2310, 1482, 4163, 5290, 2225, 3954, 5213, 2377, 3177, 5385, 1696, 1882, 1144, 1147, 1265, 5526, 904, 827, 4574, 5386, 3906, 4010, 110, 609, 2230, 3255, 4375, 1833, 3925, 4536, 783, 82, 1594, 4833, 3629, 3616, 5032, 4007, 5176, 4130, 2651, 2613, 4246, 4359, 2327, 1492, 300, 2208, 9, 4409, 4619, 5368, 1003, 2635, 1078, 4738, 2437, 5227, 3476, 3323, 4557, 4094, 3568, 59, 1597, 4891, 1766, 704, 1258, 5099, 1004, 4034, 1737, 4639, 4547, 3919, 2349, 3284, 3828, 2962, 4959, 5384, 239, 4016, 3294, 1839, 5263, 5072, 3579, 4116, 1143, 150, 379, 521, 2447, 3349, 5611, 5196, 1521, 3840, 2946, 2010, 1142, 1379, 4061, 2450, 3972, 694, 742, 5640, 4937, 603, 5515, 3550, 2236, 2890, 732, 3947, 5224, 1703, 917, 3291, 5065, 4453, 1964, 1990, 5520, 4588, 4839, 2681, 4801, 4883, 4917, 5397, 5569, 4287, 5720, 866, 1694, 1191, 4459, 960, 28, 3704, 1735, 2749, 3809, 1535, 1552, 3982, 541, 257, 4198, 5045, 3734, 1031, 5597, 3960, 4170, 528, 134, 1359, 4787, 2914, 3520, 2628, 3511, 2277, 3461, 4408, 2903, 4998, 791, 2219, 2060, 2588, 3832, 1951, 1317, 4567, 4672, 1178, 2641, 5681, 3363, 5012, 2921, 2300, 5691, 3837, 5667, 3790, 1227, 1207, 2787, 720, 5551, 5655, 5167, 1783, 1579, 3683, 1811, 4289, 1438, 4857, 5562, 237, 5350, 4150, 5492, 4113, 2699, 623, 5525, 2351, 3973, 2249, 1231, 1344, 12, 5457, 3648, 3752, 5521, 2536, 3860, 4399, 206, 1646, 25, 4295, 4242, 1856, 2445, 792, 1346, 5553, 3223, 5583, 90, 4194, 4035, 4845, 135, 5240, 1030, 2255, 4601, 4228, 119, 4660, 5692, 715, 2358, 1312, 1286, 944, 4106, 5245, 4074, 1319, 2846, 1504, 5706, 4315, 3696, 4506, 4621, 2474, 638, 3456, 2220, 1352, 3614, 2475, 991, 1362, 1956, 464, 3041, 223, 4575, 4987, 3084, 854, 1215, 3930, 5590, 3714, 3907, 2728, 3951, 3325, 2483, 1127, 228, 1425, 1591, 1163, 2529, 3572, 2707, 559, 1275, 5573, 2070, 5712, 4643, 2397, 4892, 3632, 1230, 4171, 1160, 3161, 3939, 2213, 799, 5076, 2643, 2927, 4834, 611, 1673, 636, 5589, 2757, 5404, 1801, 5674, 1999, 2569, 2618, 2495, 396, 576, 3543, 1146, 2575, 1878, 1968, 5270, 3770, 424, 3181, 1219, 5128, 3190, 592, 1847, 3420, 2937, 2023, 5407, 4197, 5678, 4669, 5528, 430, 5584, 4044, 5537, 2704, 3882, 1678, 2881, 1199, 2929, 2577, 2844, 868, 1080, 1285, 1837, 3366, 1058, 2493, 248, 4819, 40, 5625, 542, 3335, 2714, 2703, 629, 5545, 2879, 341, 5243, 1093, 4797, 3159, 2382, 1767, 5512, 1441, 1021, 4137, 2763, 842, 2862, 4065, 4546, 3927, 3570, 1217, 209, 3220, 5643, 1099, 2245, 2361, 1124, 2254, 2508, 3636, 3662, 280, 5410, 201, 3966, 2309, 4949, 1404, 3944, 4655, 2114, 990, 1805, 5379, 3014, 2482, 2393, 483, 243, 4997, 260, 3792, 2861, 509, 4922, 4610, 5460, 3484, 1156, 2775, 708, 4155, 5031, 951, 1483, 2218, 683, 4296, 3269, 4, 3137, 2238, 5416, 5326, 4716, 1084, 4733, 4623, 2521, 4702, 5646, 4724, 5151, 3304, 5409, 914, 3475, 1834, 2319, 3079, 1982, 5086, 338, 220, 3446, 5421, 621, 3705, 1812, 345, 2204, 3272, 4965, 2301, 5536, 2248, 2061, 1963, 1341, 4510, 5713, 118, 2224, 3051, 175, 3528, 1600, 4380, 3424, 2885, 1985, 806, 1074, 274, 3002, 4549, 4021, 4193, 3309, 4515, 4791, 3268, 4279, 1995, 2459, 3142, 1650, 4123, 2657, 2307, 3899, 3673, 1681, 8, 2033, 3327, 5644, 5463, 2626, 1409, 3044, 982, 1936, 4765, 1012, 1632, 2266, 2227, 3573, 2352, 1814, 3668, 159, 3031, 4652, 779, 1364, 4746, 1246, 2743, 5095, 1857, 855, 5534, 731, 5486, 4400, 4763, 4908, 3710, 1429, 4630, 1065, 1850, 4668, 617, 2190, 3373, 4362, 4423, 3301, 2410, 1790, 5299, 718, 3796, 4020, 1051, 4533, 1988, 1900, 5703, 1356, 4369, 3936, 5456, 2395, 2587, 3412, 3622, 4461, 4196, 2462, 4627, 4929, 1330, 2105, 2876, 5016, 92, 551, 2283, 127, 2882, 2949, 4511, 1876, 2784, 4637, 4544, 131, 4784, 4476, 480, 5276, 4262, 355, 3403, 5264, 3064, 1709, 2148, 3684, 4467, 1865, 5424, 4420, 5603, 4141, 5033, 1549, 126, 4814, 3332, 5155, 4352, 4861, 2026, 3033, 2205, 5307, 581, 1670, 2303, 3587, 467, 4590, 3990, 55, 894, 3319, 761, 1699, 3736, 1198, 5635, 1674, 3138, 4272, 1450, 3479, 2801, 3008, 459, 2770, 5204, 2324, 1992, 3656, 1543, 5335, 4426, 3864, 684, 4521, 5548, 2046, 3732, 5513, 2197, 2893, 447, 4935, 833, 4159, 3910, 3536, 1149, 4344, 4690, 2267, 3174, 5508, 4443, 1909, 2442, 395, 4470, 3149, 67, 5381, 4254, 642, 4671, 4928, 2653, 4465, 3338, 2541, 2196, 1728, 1546, 281, 259, 329, 2554, 4789, 2565, 3119, 2676, 5110, 1272, 124, 1842, 1164, 4347, 897, 4098, 3745, 3006, 3865, 1967, 4450, 697, 1376, 4566, 840, 4259, 3237, 2444, 31, 71, 2466, 450, 5552, 5322, 461, 2718, 3769, 1386, 2098, 174, 5078, 327, 787, 267, 4564, 81, 4958, 3408, 2759, 2385, 2418, 2839, 1459, 4089, 4353, 1105, 5570, 1540, 2595, 5108, 3459, 2085, 1384, 702, 4808, 4790, 1959, 476, 4096, 2970, 2772, 5641, 4961, 693, 801, 3066, 1589, 4068, 3128, 562, 3339, 1858, 1095, 510, 2510, 3123, 4371, 4233, 3060, 19, 2479, 2133, 563, 2441, 4271, 5687, 797, 5305, 3604, 4104, 3324, 5490, 3789, 3657, 3042, 5426, 2047, 3645, 2886, 3981, 4651, 33, 1123, 5417, 5530, 3303, 3749, 4931, 3871, 4072, 3538, 3340, 342, 4050, 2468, 1240, 1339, 5256, 3554, 5253, 2109, 587, 2425, 3738, 2709, 317, 4132, 293, 4052, 4383, 307, 4378, 2200, 2119, 3651, 4095, 1057, 485, 2812, 5546, 484, 1918, 4176, 5114, 2767, 4086, 130, 5075, 5153, 2216, 292, 1912, 1049, 5182, 5124, 5142, 4885, 2004, 2856, 4893, 817, 4092, 1628, 3313, 3800, 5041, 2102, 5254, 325, 1730, 2095, 4760, 3114, 5694, 4079, 5247, 4528, 5275, 4483, 3245, 3595, 3571, 5313, 906, 213, 4212, 2871, 1485, 3221, 160, 3621, 3413, 824, 3154, 5225, 2920, 5444, 4358, 2244, 3779, 1054, 4840, 5178, 4382, 3900, 2677, 680, 4793, 3184, 366, 4435, 3729, 5042, 535, 3151, 4437, 4256, 4195, 5069, 2454, 3780, 1994, 4803, 2773, 3433, 4923, 1777, 1092, 2996, 5255, 1884, 3305, 3499, 3050, 4389, 895, 751, 3879, 1462, 3205, 3523, 5149, 1563, 1484, 4579, 3497, 60, 111, 431, 5218, 3949, 44, 1103, 4641, 3227, 415, 5272]

    print(s.maximumRemovals(S, P, R))

    print(s.maximumRemovals('abcacb', "ab", [3, 1, 0]))
    print(s.maximumRemovals('qobftgcueho', "obue",
          [5, 3, 0, 6, 4, 9, 10, 7, 2, 8]))
