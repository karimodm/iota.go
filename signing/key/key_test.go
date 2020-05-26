package key_test

import (
	"encoding/hex"

	. "github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/kerl"
	"github.com/iotaledger/iota.go/signing/key"
	sponge "github.com/iotaledger/iota.go/signing/utils"
	. "github.com/iotaledger/iota.go/trinary"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var nullHashTrits = make(Trits, HashTrinarySize)

var _ = Describe("Key", func() {

	Context("Shake()", func() {
		Context("null entropy", func() {
			var (
				// compare against official eXtended Keccak Code Package https://github.com/XKCP/XKCP
				// KeccakSum --shake256 --hex --outputbits 31104 <(printf '\x00%.0s' {1..48})
				shakeHex = "eda313c95591a023a5b37f361c07a5753a92d3d0427459f34c7895d727d62816b3aa2224eb9d823127d4f9f8a30fd7a1a02c6483d9c0f1fd41957b9ae4dfc63a3191da3442686282b3d5160f25cf162a517fd2131f83fbf2698a58f9c46afc5df74a686f5cd121749cd41fdac5ef39ec074e8a53f1d0b42314deb6ffcc684b7c28e19cb6c0b2decd25b8d226546360fd28f77ebf1f1f07701e50f7e61faa22d972b142ad9f33e64c1f01fc6b9079d16de6dedbfe0271eadd71fc7d4b76a073cb2f6793ed2d0b1e43461eac06b510cb4ab74bd3c24ff19c6c732239241a2b402f3635effeb90a36a1b7fe118152ff8cc7c8f01188c60b35122d91065ac723403a0181779ced44465c1fbc3ae14ea85c0cef31a1f28901e9dc940f4d1ea3fb16bbc90db380c463498acbe7b49217258e8f9ef28d5db7c0fec70669f56798040c165b85a293c048f3a19e30635d234d48e0cd30158cd4a5f75ae6ae2f1d98e6499cffd2fe014a17b23d8c617f9aceb11ec80c13d0c428f8a06e341c36561193b23e918457f40f8a7076abf2bdb8d180e078fb104dab3c7a733a0a3f59acc19e1a1b81e9c627545f681c2cec900d631ef7c2eb757131c059a4c37f782f945c341fa8ec7159f22f92fbd83cacab42d1bdc1efd723190ed2449748fc20be11d416158526c326d9bcbc51f74f06244fa87aa44bfdebdd256fc677a96098f30edaa5f982be91730336b0fd645491bf00a65d510f84122c543668147c61e9062c617eb4203059c2e408516ad349d9f09668117970377c2c09ed791a14753a9de885b2423dd13e6f815934897cd44771797e4bd9457dcedde5d28043b7155ce36b9c00aaad00ef2623428660407957c5dba2918193786386ed23e7213fb9de4d67eda5fd64b4a1fbcc65df2e98fc0e932b85032586cfc9b281b9136cdd439a96e3baa3b1e312085a860ad6151a53e42b523ea9b003651237628a74a5051bfbaf8b31093b09fbdea1f8ba26abb1a5cd3e08e844b5c3ba35e10f150922942fbe3939458df00a969f8ef01abff315a9391ecc658cad3c45bf5c6dc024de2c702074e5b6ec1b104cf831b6bb996b6af451587776f48f9fa51ebc4ce7f4c30c4d0c047e271840897a4c6242422a9e105b488d21236d705de9fe54df3e034ae616df4c073eeaa041dcec462cb0cc512372dc402496863c75eb9c7090067d35f23d416b164d85855b2d26384ab6f42d3cd5a4a114c6cb9c63362d2698e92851b71cf68052a2b1cdcaf122eb86c59af021b413e95d94c5a366b9f7417ef17a219ecdc21c78aaeeb532433eebce6b41ab65a4ea57323a37ae9e2dc00d9b52a026c524c59327ff1c0cd7319ad529d8feb5f820a203eaccc81a6cca4b15d8bb6f9e647768f67f1ceb8f1606cf351571d537c7439637bc770f9f9685d02537863c945f1828ff96cac880dc2bc042bc239cc18d189fe65681b1a2b2df59bf80ce62c50a7fb767540e9ccf916a687ee2abb5e8b6b778e704b79dc4b8515992dd62399a7b0d8f86df3d37ea3a2a0097375bc67161c3f31ca4a8c1b06dc43923772a0951a07ec4c3c7e099c72a7aac33af19863ef031f89aea5c5b782531e150ab0da9bf86b9d6d2d2ac3be77fb248f6b34fe0929de9ff6c918a27d48ba3fef75293cfa8ef652665bab2257d336b94e8fb9e86e20bbf5694eae1de3455ec63bd5dc884a9a27b6d633712669fed7ec7bfef46e5d6753ca818e08714137bf04b2222a4373ecc4b37f3d1af91585d485e2a96fb3a92b6920410fc1814d24cd1c3dee61e5a1af4f64872da9ad09b88f6f1f651fe8216ac54b4885cce044ac98f1ee93a440b46ffb7cbc966009c528809dcb9344b209b7f67a3acd6c8af3aedf8c6929c389d2383d4d74fff8375c42a97c0a8aa4f50696abc004dcd0c643271771e6698922bb2151705e231ee4fab6c190b98c840280dae70214e06058e1587749a3ef098760cb5c3c52cbdc5b2def31f3256b50fa0eb2595849dd0ade986ecc803c0b709d5d97f6bf547a8073cfee5138eef6198bad10a8fafccf32698e0cbb3cc9265124cc860a021904c8e47e2c395cb4184f447710e24dc08f9188acbf653439c038a2955d8a6f69a79a06b81e7853cc5c720eee880ed383a0c347d2ee4743eb04a1bd4b96a0485a4dc04348511a006a37e90b69e47e58ed9b1ee52ab917e14e98206ecf43c550394601ae1deb81fac67c061544894c6d25faabac97934ce79126ec1bf22c2d912f624e52b30a80cafb597bd0a81dbdf2ef4b610c46edf60ce41eddb8785c645e8c7791848515904ac0d641b60b3e35ce4d31152db44a0c670433cd112e9348decb84b0820691548133cff2c3ba76d86aa0da5db1c10151290bdc8c1cc113775d86c9e92eb3de8128afb44bcb50a41dd2a2a4f876a8b1fa11b98bcfa5e265c722133faf618bc924ea585b7c1b2284f02cf7be33fcb2a8526ca5dcdf0204d1b04636544e4786d1113a811f30a9fe651b1183f78d8b139d1a5629d5e7624f5ba1b67ce1b48f9b5dd57af0376d7f08e1238e071a0e7f8a5529267f889246636cb74754bcb212863876ccb55deea96569c4d0ab4bf15b21284587cc78c5ad3bdf585e96ad07cedaa7a16aa28dec2c12f7a0833135868a06bb945fc46586c4a80a15558c8b7268fa269566b5ab98bd12c05c0f5fa72aaf1ac0fd74953acb3586fdffd4841691fece077ff779b9855775aa1ad244706adfbfc64c9c09f6d6625580fb96fc694919e39f458a1de0cca13817c89bb5d3f592a1945cd4390fb48444c7a52e100f6a02475486279f096ea1867e4880d78e1aeb6f357d618e71c95220d46d9933cae02295d3d511dad0136386274d46aa4a726d00a986bece9bb6e73231325c20d218512cd0e243d8a82b6df216098ba2ce222b2960a98678835b3113b353d63c5db58a33d4b1ff062f17b98e28df01e0444d5ffd94e14190a9c3c7fbfe8b9d9c259e9bb565475a729d49bc4267e84867ba2c883e751782dd6a83a11bb184f30b9f88a1ed984e0990df0a112e00bd536c08ac526b3cd8885a003e224ef1548e4b1548f6b6f37c94db2cb32592eeabba086e96464137a325ca56faea554643de480bdc6d9952aca1149ada63fe9db89531c484787d6fa1abdb1ab8f04d5337ef9907f589f943235fc450b52a708ac6f69e3e30fbc376c046c84c46c5946e3225eb350fea3520292cb5d1540fb9394bf6cbd36f499a0f62fe5961c26d727bb14ea9b2bde5d0470822e8dbbcfa6cb9ebee736d1236712b63c6f8f6eb5c404771c9f645b248429bdf86bdf7ab572bde01fa67382f26b27a7d6e4db3f3b7dd471a8c8ce5287dc9948658c237f321922e6dcb7c821ecb1bea5dee0f56f67104b0483f15d6819ea87021d81dda708836cd9cf33b4c7e6ceb5d93b307187257425ea2eabbdb91635e52f82bdf2dc48f1abe044e51b107b47b1f3051aab382b5bfa345ba2934a4f9f099fad7711e3d50afe66e6db8aacefed812c788177bd8a3250f371d20c09bfbe634ff373045980b84398c93b13e7e53810ea5b234043ca58970625a7199ee54d3a09cf19aa459f5913a7944ef1e67adc0330484d2234d16eb856f989b7284f73bae9c6b0e65415906b203dab795fc1474e76ed13369f253d9e6319683611dad7e121615140dbd67e37299e2cbbe5d497dc8727f2ab3e5dcffc1896f607b9a92dba7935cef13b4cfd460148324b10b271bd7a427980d16a4aa3f82368329fb416d40cf2865567302d092e46b2207c0df4fca49820eea641c2cc2ab0a9240915b1d5665268d211c5cf95338d08d7f79fd46c0cdb28a04271a56a078e7420e35cea337646d06d406bc746c443fe1ddb388eb49bce99efc263f1f0226fb2a117e3c6aa50cf0d15b7bd93b57a1507b5dc871a633bcc171d2482f95dc4a97fc89e0efd3adbf1730f09ba3220df03b7052e4822ca0ab0038edf7831a76fea71e5f2eaa8dd740d4ca70c6c7b5888a79f7fb3f439eb1aff8a7da4668ce2344ad2a35da3d326b3bf8d19de71e07bb9c3361388c07c9a3e4dd1f75402f2710b921453e40e65078e22f6265707f11fff9b829046c0ad657b89eb39fc6ee49ed73d0a60b810aeb42d64f6d2ecc9f8ac678056ff68b20b301324d1c2cb3a26af6233ec6461bd8a89f00dcdc5e1f33fefd6f7f39a39c85a9e85ce2cbf7d8c3d3c59c4970a538d9c71f4468ed33bfd761bba8717b64f7c59c4993446ebb28307f41d752458282264c228445cf5a7e92f038aa8ac7da7748a9eaa3df2a3482a099e66cf15f3b9ebfa4cd9494e930db1afc78ba8d403b2db27a0797cd26cec80443896e4f9368b77d6bb6d5035ee4169c4aaaf2fdf11bcec5ab774ae5c8cda850e4db8ddd7dd84d9313e8e768b59e01a6bfc3fedcf7af021d6477a0410ab9e69b0948e5799737189f48772240fa9c0714b42848d4abadc5d53547092638e7f1156717af13dae199860f8a6ab0801e14ea3722b555522804c6f2c098b534b9c66babb052dacb74518e363832cb667790d0175aed6bd9bd792116cb8f87640852ab552d21796b302135ad8baee04ea3f810e6cc5c38e8e4f5c7fbc237345d647995a0d98011ef44849862c0f2150c15d80174561043fded6d37a818c9dac9656f184020b8fa48bea99e15724ff028b0d9665f773ad8274267006da8d53309ffd56844e88b6956971c86c1fee845457643eafcc352f9fe957d6ea2cda38993a3437b4684a8ea906f10054d78fbe95c1c575f9800e346bc1896bcf5c24388ef20b389b5ace876613d36a073800086417c95b8f4a4c27580429a207f7507a76b0402bd2ecd3b848dc1bb2dd5474da3bf603bfa1f7a75a98a8ce50122e9a9499a939174c2cc9592f54ddd0a51926352ac62ad84c005cdad52ea22cf7030ee42435474192ec177aece3a2abc3c004bddc7001b5f26839a39cb9c58abfa219dee6c0410b4b67b6251012f191f0f6539e3e08bdd710711311e3e5f7a1af80e97c8eecf9443bd19a362c83b8e31444d0a6dc7830af1191ca59b1984469791678c9335a06b3f5ce8d5cc015494171452a4abe69aae29aec2836b056a89f7bcb0f01ef11a9a4e93691abc9ba40f3ed28603bc00262ff42465ebfef5bd2576173a5ba72716a15efcc03e285b1773f37026fd6c7bc0c97a088d847669147684802cb9753a91e5f6fbea5d4b30a4d2334f83d38a3b8c90450cd6ca7d52a6c0c5e16a21fdf291da633782d43798b5030408a811f63ef2811e25f5a988b4dd0825d9342d93a867ffaab57849ef66ad8a2d19a4b672cdc0192cb8bbe7994a7737936535436536a529c1ddbe5dccaf2ed1ca21abba698179d85ec0dfefaa69f811db6502630502711ecf804b179a844b129e32421d01c27029cb85e7127e0a2aa1a2451cb250d0322e00ee658550db37b9c04331787122fd85a58d8f899e49885dba8b90f5c16ea57a79ae84f5a83a6079e649ce22b13653a77a0366d30b187b7dfc7af75"
				// the SHAKE256 sum converted to trits
				shakeTrits Trits
			)

			BeforeSuite(func() {
				shakeSum, _ := hex.DecodeString(shakeHex)
				for i := 0; i < len(shakeSum); i += HashBytesSize {
					trits, _ := kerl.KerlBytesToTrits(shakeSum[i : i+HashBytesSize])
					shakeTrits = append(shakeTrits, trits...)
				}
			})

			DescribeTable("derives the correct key",
				func(secLvl SecurityLevel) {
					key, err := key.Shake(nullHashTrits, secLvl)
					Expect(err).ToNot(HaveOccurred())
					Expect(key).To(And(
						HaveLen(int(secLvl)*SignatureMessageFragmentTrinarySize),
						Equal(shakeTrits[:len(key)])),
					)
				},
				Entry("SecurityLevelLow", SecurityLevelLow),
				Entry("SecurityLevelMedium", SecurityLevelMedium),
				Entry("SecurityLevelHigh", SecurityLevelHigh),
			)
		})

		Context("invalid entropy", func() {
			It("returns error for empty entropy", func() {
				_, err := key.Shake(nil, SecurityLevelMedium)
				Expect(err).To(HaveOccurred())
			})

			It("returns error for invalid trits ", func() {
				trits := make(Trits, HashTrinarySize)
				trits[0] = MaxTritValue + 1

				_, err := key.Shake(trits, SecurityLevelMedium)
				Expect(err).To(HaveOccurred())
			})

			It("returns error when last trit is non-zero", func() {
				trits := make(Trits, HashTrinarySize)
				trits[HashTrinarySize-1] = 1

				_, err := key.Shake(trits, SecurityLevelMedium)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("Sponge()", func() {
		var s sponge.SpongeFunction

		Context("Kerl", func() {
			kerlTrytes := "9NGBYIGJTUTYPACOHYWUGLWO9OASWBNWCIADXRWRSZPOSRYJTHDANSCVG9KULYERRBPBPLZHA9BEONKZWMMQX9QPKDYNDOZQOZRAFDXHXMILKDOSNEEAVUFRMWCGNOVHGMVR9IZC9PTQHZOVTYGGGYOSVHFDKKDHQCENWQFEUMCXUDJDXY9JOIBDCWLUZX9999AOUHUFBAJNLVIDPSKNJEWCJKJMFGAB9QQLAPFQCOKRZWKSXUYUPAEKMJYQYFMNKBAWBU9X9KIANXDV9RQDGFTPHIVBCCAILFXOYSJWTVPBPLNBNVTNQLFBBMVPCTFMGKDAVCXTTDMFKEJUTDSHDDR9TPTFDNGSXPRTAOPMOGE9SCWMCELFCIJMMNSJIUYSNJUPPMIULOQFQBNUHIZFBUUJNTSIPQBOXFUJFRAQXXHMQJITAYGGFFMOCPHLPQPBBWFEIZTK9MVRI9DEJLAQQPYTYWW9RGZ9OMOUSBQMBG99WDGLJFTZGGVBFPYBHMTDGDODWLEPJBJXYYALQUS9QXXWAMPH9BRAPMVMDJWSBNGIOVUBWUSHWQDRPJBECLZVAEZTHXQRSDUPPJN9SNEGRJFZTMUAJENZRVKYUIQEVFYDBMTAXPNTKMPJBRJTLGBKYGKFOMYYRLEXNQCPLZHFDXEANCLRGSKYWXJCOAPGTRXNPXWD9VTASXBHIWRPONLSNCVRUUZGQ9AJAJKCPVLKUR9UYTRZIAMLBAIHZWYDIKAUMSYIIMOPSHRAYKYONMMUKFDEQDKYPICVGABOABTYHXMFHGDIBIXQEPZVKWTWDYLKIAKSGTAAQQODXI9ZPBOVRXRSLD9KZUTOSPGGRSSPDCJZSHMOSSSCFFSRUGPQEQGVNERH9KBHMXYPWIDKD9YKYRUQEJOSMHZGUALHRTMEDDDR9SBNL9VYMXQTMLBMDFJXIGEEUNCWXSBGBDQJERHKKUMYROMNPMVCPUEGAJESEAVRMHUVARCXCLRYTKIWYZKOAWKIUEBGBMEHSIAJFCXTTPUQGZGGHUNQLAI99JUVWPRBDXDRXOUTOURDKOR9SJNCZAVLUQUVYLCUGMWGZURNPNPRAHBOQWLRXO9AFGSQJEWGEMNNCHFEPKBWDJSSHHZLR9DABDMNBUEL9KXX9JQAGCJSAEEOFUTNIKKJBVQVKDRCBTMLAUNMUNY9DXAEWNYXFELNPBPQNVGZCJQMKLZWGCNYLCYOCYNYXNGYSIXWMZPCVIKSOFFNSOCFKOBVN9UNSMIQAFORBCIR9LDYEQUXYRJVXIRPXFHLRDA9BYOTT9LYCDBKSUBNTLNFFAYUIFHBXGZHPVQMXLNEXOXOKCSYXJGARFAMMWIKGEOMEBYW9WRPZQL9DRSDYSZLMEI9IQKMYTUWB9CIOLHIRBWDLHVTTNXVUIJMQZGUV9RQQSLWZM9PK9UOXPYHLJNBHET99OOA9PFQSXWM9SZ9YTIQHBZCGZFCVQEJHZNJPTMZSKJTSJNHFJFB9WHLAGXFHDVRKBVLZINGYW9HDFPBSIWFBWKIUJATKCWFQVMETWPGW9BOGBEZHKNJNBBODMLRAUKAEAHKVLILMRTJXGTEVYVYYGABPOCR9DLFALNHYXDEVPGWR9CJZ9LZTAD9WW9KUGYJYC9HSQUDQYSAYBSXMJWPIFVUKNEBVOKERQUOEJNRRRGMIJYHILBWAJIBCWPHY99GQOMZURFBWPHFLLRZAYGYMFJIZZKGBSCFHFACXZGVLNUUERPIELPQSRKJKOSBQMQEVYOGZDZMVQWUWCCZSHSRIXMFNTMLXLXEG9UNADZGFTMXPRCSRRPBXKXKERAYCWBSZFRCDBAJOLUS9WXIFSCJHKSVLMQLFU9LLGQRTRCAPBPFRVHNGEDLM9YQOKBFUHVGSTYIMWCSNBYFLZPZYAWNKNBWBRBNRARFWHMLPMALGNECNBCCE9SVATELSPNJCCCRRFXYASAILNTOMFBGMETPXBXGMGLHPUMXUARJSPEDOEZVCUWWODVZPUJKWRAESFISBJQOLAJADXDPXKUPJTUBKXRQQRANAABEOSXH9WZCFXYB9FWNPOFCVKGYTXDNHPIDAGWOMZQMTRBPVTMYACNHRQHEKE9ZSRGQPXLGYSUXJDL9JYLSLOBFNFSQMKISVYSKCFOIGVQRNYCI9FWWJL9PV9VVZILTLMZCXBWVQSQPYJYONKBXWWZOPVE9YVMQBYZOFNVBU9PZETGVQE9MQYWIAONGRTKKQDGUQPUTEYDNYWHVRYFWE9GZZLYLFZNSNW9KQYRVVSOLUOU9RBZLSVLCQKGBVFOXOHUFSJMQYHSRMMHRPMUTCKDGM9E9JWC9S9NPJAVECCFUOLWHWIVIVSBVQTORYNEMPGZPP9BA9ABAIWPMWAVOVLKQVHJXAJBAMFWN9UMXTUVQQCT9XIEWUJVRHBBIYR9WAYJIXDTWOIVQWXFCKOZNOFLJXNHMRQRDFAKRGXGKEZRZSLLNEKJQWIXVUWHAVDUJGKLJZWURTSBRFPREMSNUFXCSKFIYZXCNIGMOQICJGIIAKYOTAKAHPWCCJVDISWLQTBKUWWACUUK9IXGZDRQCMKOWRBKBJHHUNZIHI9YNIE9HBBTZH9XLIMFPFL9RFPJSQALXRPYZYWVRCWYMLUPIJCZGONN9QEAJATASADEU9GAYZEHXVYRLVO9YGATSWSQRNHYGSLOO9KBOTIJVJSAWTOY9FYHVHHBAHRHIWGAKMTFJHCHM9LY9ARGJSMHY9BFDTQLMB9NRPXDRWI9KWXHZFWIQME9LVDJPZSPUMUFZMQARAOYYCQHNMM9OQD9Z9GWUQHXONJJSIBILIWZVFCJGDJDFKYBLXPAUSJSLARTAUWHRROQGGKAWVTDO9JZKUKIEUK9JZXRXZ9VHGFAZOCLRALBLYNM9IYYCGOMDEUDUWPGOYHZKGXTC99RBOOLJBHSCEHJTBTSDAXKAMLHDZBYRVDKQHMFZGBSCFO9AQXCSCCKAQMMSWJA9CAPPTSLMGGXRKYKLVMIGPUQCRQLSMFVWQIKILLZAXZAVEE9SEYWZVCNUSNRYTSTRNGQBIAXJ9OIHPYWYVLUCSW99UQTZJSOZFBIK9RVQJBMJQVR99XAKDZIXIWPNBQWJQWHGFPJXDCKKAAWJFTSJGTUWHQ9UHCBPAUJGFEN9GDR9MCYEWKQMVRLPNAZTVFKRGULYXBSOISPNYCLUHHJXEECMAAHZRUDERWGRKEBDEBDWYXDWPNOGTUDAXHXBFXGWJADOLLUIN9FBIMCKDGNUDIJVIIOLISLZOSVWBFVBKFMHNNULNTPSZGXHTOW9S9VCAPETLHRGIXAHJTRRKFIOBOMDLZKUPMCQYMWDLWIZZHNEQOYOPAPCZMAAXNSUKXDG9MNYIJYNBMKI9LRJQXKWUXVAPFC99PS9OMGQVVXWZVRNOXUKLVGRLDNHFXAUDXBEVQLKXTECZUNQTGRZCJOORDLXMJDCOZACWGCDRVLUMMMNAVBCSGHNQJHWBAYAYHVRSBLSVKXQGBDKRVMSRCLHUR9V9OIFUOGOKMHVEMNLTGLUJHZCBTMBCMQYVPYZYHFEGOSJDLSSTBWTCSSUEALNU9MMEGBDCLTLYRRQUUPXJ9HTXJHD9KTDKWWUQIHTMQMVLSWRANYOHXGNUJLVKQZMZXIHVPAFAENEBDXBMM9LLRKJYCDTJQSPXYA9KADNBAJUGXQNYHCFYUCKDQTLCBFHKDHOLAQMPUCB9TWCMEBXQISAMWYHEXTXSMATLYDKPNCB9APKHPEVISGGTEOKPJXWELLFGYJDEXUWBASPLFXNBIRPAMUOEAZJYDFNRXFL9MHNGEJBOGZRMBIWRCSL9QQZXA9HDHXFICPMRJXHJHQEGJWVWDKBIJBNYSD9FFLZRPQACOEKHXRSGYQYMVLAAVIIGNXLMFSM9TXQCDEUPOUBWSAUBDHCPZJINTMASQMPJZFGISGQVGTQARDVNNDRUSXHXOSQFUCVHQWTALTJTVDGVCPHAGFCKXZIFMBZJXBLIKWNERDYVDH9LOFQFCHZJPHCZL9AJYFDMVKVUXAPRQTQIQFFGJPBOANZOCG9URDXFNZDABDJWXQJXCHDAHISTJAEFV9YZQMG9ZW9AGRFSBWRMTZXXL9IRKXEORSMWJEEMAJSAOLFQHVTZMJUPOHUJFKEEFRBZHXKPWBRGWQBMLDHAGGXUOINUHNPUGSCFUZZXKWSULZJTBSABT9JZEJTBANMKKCXXXGSXZJAMWCSFKJJYVGMAHSIKZTYBSPITYFQNCUSNOT9ISHQVNHWXSGGOQ9HXHVXFAVSTYDYNTHUDY9LCNQQFLLXYAOYLZYKZX9PUZTQTWUPWBTPBRBLEGNPIOSBEBXEUJXGWXCJ9IDNPJSRTGPJRDPZ9TVIWOSEWT9KRXUIEJKBTMXENKMNKAR9JRFPVA9KNOELWIDRPUKSIOZTFEGDU9WDOZAOZQGBETIVQBYJUKHOQVJITFHOBPCDNMBRWZDEITBDZPTX9YOVSAL99DIVNZSPRFGPST9XCQDLFLH9LKS9JYESWYBXIRSKQMGNDLHVFOINUTBGDMYUGFSSFUKPAUWDTOYIUTSOBVQDXIUYWBTVCQ9JCDYUFQTYXRWAJMLAVBWQAURJDJFY9KWXBREMSYQF9LNRREAVZTMJBHTQUDFTTNEH9LXOIF9UHGCBAPTKYFJKR9KFNDDMFBAFMDVWOLYBWFRHJXFGNGAUAIOJ9ZRFDKSPKZYOXFDUKPFPONQJJPSQAZLPDIAHOXKTQDSZVIFZCBYBAMRGVEPVORFZSZCVTESUVDDBJYUMISUDPVGNKAYGWHHIILPTSATDHBVPU9FJAJRSQGICLFUPDQVABAJDQGSSGAVSF9MXT9ZQJIYTACSZCOLX9UXSSEVERKGIQJNPDM9QGYLRSJLEKDHACLAYGIQHAIC9DNJOLAMT9WXT9WVFBLTXDHFEJHH9UTOLSOGQOAVYBVQOCVPZTN9K9XOF9WXFDSFAGJYSADKMKMBUVETZWRPPMZXDBYUVXLMSSURSNOZAQGZWPMESQYFOAQVFVUKOZHRUTNDPCSEQUNAXUKI9ULJZLGBYYJEOBV9VBQGHYHANFFCKZGDJLLBWCIMIWLBGO9LDQPOZRTRKREMVYPHVKADHRRIDGTQKXNVRZJHWJQMQDALZUJOPPUTGA9NA99C9SWAWWGMKMQFIKHRZJU9XBLCAEXPJFYWHEKIJJIWYJSIRARBCHTKSOVEXHDZQPTNUJBITNJNSSLONPMCVXWICFFHLVLGCWHHYAJOLIZTKCLOCUGSLCYBCUGIDJGGQDEEOBKFPCGTXENCHDKZWGAUMRKJIHUZAABCFGAIOBXCNTCUNXQFNRRCSFHQRVTVDPAVKIOLQXXBZVWV9DCYZUHGHVAMSBRHCBJGHFICTHRSMDQBHXVDPPPXZVFHHXQLRVIPTAIQLWDSJTQMEQMBUJBHTMUXFNBSNJNXEKRHGKYBVRJOFYDSOXNBKOMZCABUYLJSOFI9DHNLIILEZJVYDHIL9HGP9SG9YLCYQBIVRD9CITXTWKZNVUKNJFWHAGELCSBXSTZIYIBLUIOBNNQHPDMBYYIQCRDZKGSWKAKSLDXHBOOHVQPQXOTSFXRIWAHFORYYKGSKSKTXNLTJPTPYYRQIPCYHOIPFOLOCNTXITWDNUSASNJMDNTNQDWKTTR9BNBRMUQNMRBVCKDUK9SNLAEQAKVEOAVPFTITZHXQGWYRBZTZRYMMQTNLTTTNWSJFRVOTFITSYGOITGSDJQUTCMYBRVAI9TTBTUEYQDPMCVUZSBWDLEJYHAFMPJMJPSBFVSHMLAEATVCGSYSRNDNFVSPLLKNWFYLVWIPKXOEKJKFBRFRTERUZNRHIMQLZCZWIMONGCV9C9LRBNHQHFILJDAUQGGKGYIDCWDKJJEXJPE9NGIIQGEFODVARUNHIFXKVDZXMYCLJDGUBXJLZYGHFVXFX9KMHFHVR9ISLUNMYBZGPENICYZFTNGVOZFGEGUUMYJISAECN9GHG9JZPTRYHHN9UWFQCXLWUGRKAXKXLKDNXKMENDPTUBRVOAR9ZNNGVWH9BGSD9YONIWDMIATEPWIDYNC9K9VKKIBDIMWI9UG9ILZPZTKDM9DJHKIZFRWKMXDTVDZKPKXAWNBRLFZGLTZHETEQQZSZXVLTJHNTNNRTTSSAJUAWFTFKRNAZEJIOSBDNYMO9ESA9IILNZKKKSUZTXNRYQVPZWAVWLSUGJX9FTRJBOCIREQREMAUANGUGXBRSYWV9TJTGVRLQNLTCWVDKXD99W9MUQIZUOZVQZKAFHWNEHFQ9BRYUSTLKZKACMGDUVOVVZZVGZRLPREXYWUQTULMHKONLBWOJFEIXVNLES9RLQZBJXGPBGWBXN9DQFYDBZDMKTKPXWZUCZQAXYRXNLEFMJKWDNTU9QUOWMXRKJZXLOZHVPSLIIOMWVRMZGAMLAE9QOVQOPNAPACPOFYDFDTWGAQZL9CVF9VKNMMIAMNXZIRYDGWU9VMOVQIFSEXWMQ9KVVMNJTLYZ9CSYMPYBPVNNGP9UJIJEWYMY"

			BeforeEach(func() {
				s = kerl.NewKerl()
			})

			DescribeTable("derives the correct key",
				func(secLvl SecurityLevel) {
					key, err := key.Sponge(nullHashTrits, secLvl, s)
					Expect(err).ToNot(HaveOccurred())
					Expect(MustTritsToTrytes(key)).To(And(
						HaveLen(int(secLvl)*SignatureMessageFragmentSizeInTrytes),
						Equal(kerlTrytes[:len(key)/TritsPerTryte])),
					)
				},
				Entry("SecurityLevelLow", SecurityLevelLow),
				Entry("SecurityLevelMedium", SecurityLevelMedium),
				Entry("SecurityLevelHigh", SecurityLevelHigh),
			)
		})

		Context("invalid entropy", func() {
			It("returns error for empty entropy", func() {
				_, err := key.Sponge(nil, SecurityLevelMedium, nil)
				Expect(err).To(HaveOccurred())
			})

			It("returns error for invalid trits ", func() {
				trits := make(Trits, HashTrinarySize)
				trits[0] = MaxTritValue + 1

				_, err := key.Sponge(trits, SecurityLevelMedium, nil)
				Expect(err).To(HaveOccurred())
			})

			It("returns error when Kerl is used and the last trit is non-zero", func() {
				trits := make(Trits, HashTrinarySize)
				trits[HashTrinarySize-1] = 1

				_, err := key.Sponge(trits, SecurityLevelMedium, kerl.NewKerl())
				Expect(err).To(HaveOccurred())
			})
		})

		It("SpongeFunction is reset", func() {
			h := kerl.NewKerl()
			_, err := key.Sponge(nullHashTrits, SecurityLevelMedium, h)
			Expect(err).ToNot(HaveOccurred())
			// absorb is not possible after a squeeze
			Expect(h.Absorb(nullHashTrits)).ToNot(HaveOccurred())
			// compare against new Kerl instance
			newKerl := kerl.NewKerl()
			_ = newKerl.Absorb(nullHashTrits)
			Expect(h.MustSqueeze(HashTrinarySize)).To(Equal(newKerl.MustSqueeze(HashTrinarySize)))
		})
	})
})
