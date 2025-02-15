// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package chain

// testnetExtraData
/*
UTC--2021-11-15T07-19-41.272626000Z--a45d87cfc1c3ca535d6609746e21e416cb5c1247
private key: 1597a7f6a406b2180e5936797a6d0bc2396896220bd3f448daf58f6056eff1a8
public key: 048ee7932f31a81dd91c9f66775503eb366e411012748bd55d7499c9f9c1cfe8549b75c6275a8183ee6791b610e2a47f6ff489d0f60905ddb5ff2a9071906c7fd1
bls private key: d7d43c0ab866b6ffe37ce4fd37a975b4ce2c68b7df1ac94df81bada453ad6910
bls public key: f9bf370ce720a3d22c72cb3fb13b730604344fb8f4fc171388aed7a2fbbbef26b5e857aa4e553dbe693d7d2caed92b01ca2c9c9d5a10323f270bc1dd6435dca15d58038447f55ea99f830413082792d9cdd90c9e7470e33bca8c1a9babe95280

UTC--2021-11-15T07-19-43.770320000Z--d74d24cc30687cb802c059b5b26e3503fa1feb4f
private key: e05c112f7031b4a86562f3db25f9b5964026bc1f07f5696d7c46eecbeefa34fd
public key: 04a2da8be962ee625efd2572db5ff9546addf9f34795f4c054ff85459b251cf300c44aa58fc539f33e7b4fd4e5f7b429d9edb702f3a16206a7fedcda9b0ce01f54
bls private key: 61cc81f32139bb636c949b27594c096b7b860d32b41d5c6fc0bbc0b8c4da4b08
bls public key: c048189cd5e2cc5695936719f6e40153301adc284e6fc88420e6aee97004bf1cf83421801014c501a76ad397848d5b018899a8727089422af17cb78f77419385c79e3f40736b7caf085f42ced030c1988366a47526801fc108b6d511b6e53e01

UTC--2021-11-15T07-19-45.975328000Z--cf1948fa74fc8c93e77c7d57b4e696f29eb30d3c
private key: 158ac4a6d97672fc4ce3e521cfcdae98b5ca3f6c39275b9d61adde9d9e5de68f
public key: 04c9e1988255827507fa83f09edff814ef5de05498347ce1a2d59d08e5a588b15b44a27fe89d2bc03ecd68cc0d912d761c1bd469ee3c3ffa9e544f79d22ba78c35
bls private key: 6363f0bba77ac948f7154550d325f4c9594e7065edf3abfc7fa095d11c421d09
bls public key: 64aa284431b5bc99418f2c4dc756c768e53b3d26919744daecaab976c5421160afeb938b9c7732b3ad85d7bb8fb367002297a280e1ba8be2952b6c28e08ae304a440fbe61cd8b0b0774b78387599245af01998672ee279dbae2fdfb8cc08ea00

UTC--2021-11-15T07-19-48.067826000Z--b522eab434df6baea887ebd1db4ee9560397515f
private key: e6d966ccf4286a89afbbc1f0d582db3940774600b24e2e57569ad1a3debac0f0
public key: 047885baa9ba610ea1693f3b30a0786101c1c7a4119f781df0af9881c8563d9134ce1a5cd4e7866a9de72ac68e514fcb8408d3f56db19289f960b291b631ca7658
bls private key: 3a23df19aa06efcf377ae35be7dba79b443725a5f020a606d8ba57af3afab404
bls public key: 2c220b971937185b5ba7c7d30e652b76370ff45643740fb894a77e860cd60a48e4290a5605a60fb9867cdb4709ad620035cb191a931146f3f7df29cc0dc47fded20d2139f84030e7bc57dcbd800743a5004b3f5b99e37f5209442684a98d1280
*/

const (
	// extra data in genesis block of main net
	mainnetExtraData = "0x0000000000000000000000000000000000000000000000000000000000000000f90376f85494b4e1bc0856f70a55764fd6b3f8dd27f2162108e9947a3a26123dbd9cfefc1725fe7779580b987251cb947607c9cdd733d8cda0a644839ec2bac5fa180ed49465b3fee569bf82ff148bdded9c3793fb685f9333f90208b8801446c55bf2cd8f31d31eaa58b081b30cea50b4d8d9096682576688c2f9ae627d27f97e09b64c99b0e49e7c33c864c089ba5c03cf27e04af4ac457edb46d12ce92aa1ca438667203f7d1696549bd861bf6f11cd3cb7d67738222428137ecadac91b68426ad13c8af92a9b8dc62475fbb1617640c635b812733efc9b7d21c8ab49b88025f8387695e95f4224919051cfebe64efa2efb4ca8b82bd04cada59edb9e54920a10ea0f9064067049aaf12299592b4746fe4085516fc3f00fbc7dbdd1c545fe126f7bd8d6b5eff857e6fa227c58eea0da0984b80922f5bb33dba1832cf8722f26dc97985df5bf1b4db96f102b905c6982192feecf860ffba2d5eb4db3a64bdab8802c87c9887cd5af4e33cf09b0a7c4840f87b53bfe5a21d9e18051228577fb191615bb9e374b2625f11e47737292a8688c9ce1c8533ab8000f495a4cc2e55f5531303a6a1542be1e2b6a0e6f37fe2f3f2d60d39b7516f831cb9a3811e82f2a2f1825467b24b96006ee5f494e7469b5a599ec6ab885df02d0ead0926ca45c4f2a1bb880138cf64f20414eed4eb6e5b1b0e08b03944d368d2df78461b493c16cee52fca12741fc9a64ca55a6c08987eba75fc1d1e13a64b9e1e6f6a4413acda3b4997adf184ae5b75a5bd807d5d03d698cd0366d443786276749d2831416dd497d3e53170b14f7028a3245c0383e378130ed3c439aff3b29be9195d778c13e462e4a4afaf90108b84013524ec450b9ac611fb332a25b6c2eb436d13ac8a540f69a50d6ff8d4fe9f2492b7d0f6e80e80e9b5f9c7a9fa2d482c2e8ea6c1657057c5548b7e30412d48bc3b8400e3450c5b583e57d8fe736d276e9e4bb2ce4b38a5e9ac77b1289ba14a5e9cf581ce786f52d5bd0e77c1eacfa3dd5df0e22464888fa4bfab6eff9f29e8f86084bb8402f6dd4eda4296d9cf85064adbe2507901fcd4ece425cc996827ba4a2c111c8121e6fe59e1d18c107d480077debf3ea265a52325725a853a710f7ec3af5e32869b84005fde1416ab5b30e4b140ad4a29a52cd9bc85ca27bd4662ba842a2e22118bea60dc32694f317d886daac5419b39412a33ee89e07d39d557e4e2b0e48696ac3118080c3808080c3808080"
	// extra data in genesis block of test net
	testnetExtraData = "0x0000000000000000000000000000000000000000000000000000000000000000f90376f854941c0edab88dbb72b119039c4d14b1663525b3ac159416fdbcac4d4cc24dca47b9b80f58155a551ca2af942dc45799000ab08e60b7441c36fcc74060ccbe11946c5938b49bacde73a8db7c3a7da208846898bff5f90208b880136ef6be87de9c925869387782afb4cf19496999c2684709daeb3af8d0b59d800bbe05870789f0f9b3cadababa69f5a00a38bbcba71d99c4c35d671442232c4d3017fd6b99e8356a3e4e985bdfc60bbcb8d939c87976a1ff677d7c42989b379a0b4c0f168a544c892bd2b3ec480e3d6c58c7dddb8d83677ebee2e87ab3660b80b8800a2e37ecad6e69bfec9fec2b345d0f8441a0f63acf8b45c0131a78e5d777d52e0a39404ca85f2c08752c1d4ff8df05c82c7880779d61fe3fabcd4fd682463c0515b1f0217561a6a72bd381da19e34c5560c6eccb08ff83d7d3f4ac6da7f5d1ed15a2780f782c1fa571fa65b99694af559b9df168b1d8745ac3bbc7d3fe550b94b880086fac850f3a9f36e8a5107eab0ba79044043dc2cc6b897cbbd0d4bf805570ff270a98f28e2d2e70b7b2ecc41a4a13e453178354997aa2038852c5945f0564bb02cdf57642881a1b40417fe3620429fc087f8dee6a68e5d7193d3243c38a1f3827d0f4cb616722a1fa78a283a17589d7688a769ade77e9d6417c6e2a9adf59c3b88003fea7bc386ea24aaa19c563a4f26f38cbc2ce172ba2310587405f4f05777fb911a4c3553b7b6529ea02a9da3ae2df6f70c3409105b39e1930d6a6ae8344fc221f5dfb2e73cc8ce434d1af33d95366796bdec26ca7cfcc0a03867fabf471884206db6b9e175a131995bd0c70b93a6f2eec96d831ad0c42d13d334f780d578834f90108b84014d44a97d2fc3ea62b6dcf2bd857079bd261993152f11aef5dd001db68b20d2d1ba45f117b6530a7aec45d7d90fd4e15d2a62f62b706eaa115aa801caeee294bb84015b7bcf0accf839170a5d4621282edcf14f4a438f8e53abcead5f0528cb91cb1135fd4e82ede1493ab1209af122e1dc186c885cc96d2413cbc09a58163b91eb9b8402fd433e93187f6b3d15664ec48073bd73d57c801c4a8bfc1e0e3abd3deefc45619d45ac7ad54df7dda5b8afd6f882c9d9f879dbc6d587f1da5da1751baac729fb8401b037f39d9f8e74b608a898249cc3d156ff1f0051026388366b85a84aac43bb4068275cd909e16b29f1b3bc97e91ec0a8b95a11b8a574cbc2c9ea142d26c8a498080c3808080c3808080"
	devnetExtraData  = "0x0000000000000000000000000000000000000000000000000000000000000000f90376f85494053af2b1ccbacba47c659b977e93571c89c4965494b47adf1e504601ff7682b68ba7990410b92cd95894f655fc7c95c70a118f98b46ca5028746284349a594b243f68e8e3245464d21b79c7ceae347ecc08ea6f90208b88019587d8e318b681e8ff036af9c296b321f5066d03180705c4b244fd0555fcb6c0a4cdba7fb21dfba83a727ec4038967366af667efa211fe561a9ae53182040771189dcfaede47716b3e34cd3eea96dd791ca7947a62fb53721ad0ddc0dd15bcb2c41d5d6281528141ecbbbf836891742c9ac480165366bcc8ec79814f306aab7b880262292118774c85cb53e158bd3938ea804d7e9b28fa5a49ccbd23d18fb9a789b1dcd8922043d6563298355442be7680027cc0df3aa675a82d563c56db7c3903817e2042ceaa101dfbf5b279efe0dfb3842faa278486b1aa82202f92c6a6ec0a826ce3f717b3d1fd483ab0af83bc98fbc03caa7a97f32383100b48508a2b412bab88014f821e131ea6c273607d704a4458408776354c2451ede5ba2c1e5195ec3929a109cc2d5a9487bf1a0a7b498678da54284509df5a2f74e03f8965707273dc4c70a6f29acb7b49f2e9f629c410a0923b516d23f3858611c500304a8f5c80454140ab032a87cf62167d7f16b05d7ad644a3862d1b9bb08e3baa9c39798ba1f0efab8801acabcd4e2298c91e1752620a9e99e9c49513eaebde1aad381bf2946a40214f0071c72a02e910cf0318f954e3d09588af4d0ea5fec65a0c6511b0cb6423dc76d06e0d787fac2b683fa6f19f734dfbdd8bde1cb92924695679d81e265cded37cd07a3abe53e82f957db92a4dc279a3b7a5591c775fc1fb79a8185d95bbb93b2bbf90108b84025480e726faeaecdba3d09bd8079c17153a99914400ee7c68d6754d29d7832c12b9804718e2cb3f65221781647a8c3455cf3090519b15a34ef43b1dde7e3c287b840120bf5a2d293b4d444448304d5d04775bfff199676180111112ec0db7f8a6a692685ac2dc25dc5dd06a6b4777d542d4f4afdf92847b9b7c98f5ecaf4d908f6d7b84003dda4ec969ff7950903131caf2cc0df1d91c569be382cab67df539e94a45835156b522a45ed4a625a7b5906d64046dce1c112a1dddb72972ecb670145a16042b84028681fcac6825e2a6711b2ef0d3a22eae527c41ecccdeb4e69dfff4002219d8b131f98eaf9323bf171e947401f0e6b1951f4c8f8aa525b677f1c811c88358e378080c3808080c3808080"
)
